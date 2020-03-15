package cmd

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xuperchain/xuperchain/core/cmd/xdev/internal/mkfile"
)

var (
	defaultCxxFlags = []string{
		"-std=c++11",
		"-Os",
		"-I/usr/local/include",
		"-Isrc",
		"-Werror=vla",
	}
	defaultLDFlags = []string{
		"-Oz",
		"-s TOTAL_STACK=256KB",
		"-s TOTAL_MEMORY=1MB",
		"-s DETERMINISTIC=1",
		"-s EXTRA_EXPORTED_RUNTIME_METHODS=[\"stackAlloc\"]",
		"-L/usr/local/lib",
		"-lprotobuf-lite",
		"-lpthread",
	}
)

type buildCommand struct {
	cxxFlags []string
	ldflags  []string
	builder  *mkfile.Builder
	entryPkg *mkfile.Package

	genCompileCommand bool
	makeFileOnly      bool
	output            string
	compiler          string
}

func newBuildCommand() *cobra.Command {
	c := &buildCommand{}
	cmd := &cobra.Command{
		Use:   "build",
		Short: "build command builds a project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.build(args)
		},
	}
	cmd.Flags().BoolVarP(&c.makeFileOnly, "makefile", "m", false, "generate makefile and exits")
	cmd.Flags().BoolVarP(&c.genCompileCommand, "compile_command", "p", false, "generate compile_commands.json for IDE")
	cmd.Flags().StringVarP(&c.output, "output", "o", "", "output file name")
	cmd.Flags().StringVarP(&c.compiler, "compiler", "", "docker", "compiler env docker|host")
	return cmd
}

func (c *buildCommand) parsePackage(root, xcache, xroot string) error {
	abspath, err := filepath.Abs(root)
	if err != nil {
		return err
	}

	addons, err := addonModules(xroot, abspath)
	if err != nil {
		return err
	}
	loader := mkfile.NewLoader().WithXROOT(xroot)
	pkg, err := loader.Load(abspath, addons)
	if err != nil {
		return err
	}

	b := mkfile.NewBuilder().
		WithCxxFlags(c.cxxFlags).
		WithLDFlags(c.ldflags).
		WithCacheDir(xcache)

	err = b.Parse(pkg)
	if err != nil {
		return err
	}
	c.builder = b
	c.entryPkg = pkg
	return nil
}

func (c *buildCommand) xdevRoot() (string, error) {
	xroot := os.Getenv("XDEV_ROOT")
	if xroot != "" {
		return filepath.Abs(xroot)
	}
	xchainRoot := os.Getenv("XCHAIN_ROOT")
	if xchainRoot == "" {
		return "", errors.New("can not found XDEV_ROOT")
	}
	xroot = filepath.Join(xchainRoot, "contractsdk", "cpp")
	return filepath.Abs(xroot)
}

func (c *buildCommand) xdevCacheDir() (string, error) {
	xcache := os.Getenv("XDEV_CACHE")
	if xcache != "" {
		return filepath.Abs(xcache)
	}
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, ".xdev-cache"), nil
}

func (c *buildCommand) initCompileFlags(xroot string) error {
	c.cxxFlags = append([]string{}, defaultCxxFlags...)
	c.ldflags = append([]string{}, defaultLDFlags...)

	exportJsPath := filepath.Join(xroot, "src", "xchain", "exports.js")
	c.ldflags = append(c.ldflags, "--js-library "+exportJsPath)
	return nil
}

func (c *buildCommand) build(args []string) error {
	if len(args) == 0 {
		root, err := findPackageRoot()
		if err != nil {
			return err
		}
		output := c.output
		if output == "" {
			output = filepath.Base(root) + ".wasm"
		}
		out, err := c.buildPackage(root)
		if err != nil {
			return err
		}
		if out == "" {
			return nil
		}
		if strings.HasSuffix(out, ".a") {
			return nil
		}
		return cpfile(output, out)
	}

	return c.buildFiles(args)
}

func xchainModule(xroot string) mkfile.DependencyDesc {
	return mkfile.DependencyDesc{
		Name: "xchain",
		Path: xroot,
		Modules: []string{
			"xchain",
		},
	}
}

func addonModules(xroot, pkgpath string) ([]mkfile.DependencyDesc, error) {
	desc, err := mkfile.ParsePackageDesc(pkgpath)
	if err != nil {
		return nil, err
	}
	if desc.Package.Name != mkfile.MainPackage {
		return nil, nil
	}
	return []mkfile.DependencyDesc{xchainModule(xroot)}, nil
}

func (c *buildCommand) buildPackage(root string) (string, error) {
	wd, _ := os.Getwd()
	err := os.Chdir(root)
	if err != nil {
		return "", err
	}
	defer os.Chdir(wd)

	xroot, err := c.xdevRoot()
	if err != nil {
		return "", err
	}

	xcache, err := c.xdevCacheDir()
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(xcache, 0755)
	if err != nil {
		return "", err
	}

	err = c.initCompileFlags(xroot)
	if err != nil {
		return "", err
	}

	err = c.parsePackage(".", xcache, xroot)
	if err != nil {
		return "", err
	}

	if c.makeFileOnly {
		return "", c.builder.GenerateMakeFile(os.Stdout)
	}

	if c.genCompileCommand {
		cfile, err := os.Create("compile_commands.json")
		if err != nil {
			return "", err
		}
		c.builder.GenerateCompileCommands(cfile)
		cfile.Close()
	}

	makefile, err := os.Create(".Makefile")
	if err != nil {
		return "", err
	}
	err = c.builder.GenerateMakeFile(makefile)
	if err != nil {
		makefile.Close()
		return "", err
	}
	makefile.Close()
	defer os.Remove(".Makefile")

	runner := mkfile.NewRunner().
		WithEntry(c.entryPkg).
		WithCacheDir(xcache).
		WithXROOT(xroot)

	if c.compiler != "docker" {
		runner = runner.WithoutDocker()
	}

	err = runner.Make(".Makefile")
	if err != nil {
		return "", err
	}
	return c.builder.OutputPath(), nil
}

func convertWasmFileName(fname string) string {
	idx := strings.LastIndex(fname, ".")
	if idx == -1 {
		return fname + ".wasm"
	}
	return fname[:idx] + ".wasm"
}

// 拷贝文件构造一个工程的目录结构，编译工程
func (c *buildCommand) buildFiles(files []string) error {
	basedir, err := ioutil.TempDir("", "xdev-build")
	if err != nil {
		return err
	}
	defer os.RemoveAll(basedir)

	output := c.output
	if output == "" {
		output = convertWasmFileName(filepath.Base(files[0]))
	}

	pkgDescFile := filepath.Join(basedir, mkfile.PkgDescFile)
	err = ioutil.WriteFile(pkgDescFile, []byte(`[package]
	name = "main"
	`), 0644)
	if err != nil {
		return err
	}

	srcdir := filepath.Join(basedir, "src")
	err = os.Mkdir(srcdir, 0755)
	if err != nil {
		return err
	}

	for _, file := range files {
		destfile := filepath.Join(srcdir, filepath.Base(file))
		err = cpfile(destfile, file)
		if err != nil {
			return err
		}
	}

	out, err := c.buildPackage(basedir)
	if err != nil {
		return err
	}
	return cpfile(output, out)
}

func cpfile(dest, src string) error {
	srcf, err := os.Open(src)
	if err != nil {
		return err
	}

	destf, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destf.Close()

	_, err = io.Copy(destf, srcf)
	return err
}

func uniq(list []string) []string {
	var result []string
	m := make(map[string]bool)
	for _, str := range list {
		if !m[str] {
			result = append(result, str)
			m[str] = true
		}
	}
	sort.Strings(result)
	return result
}

func findPackageRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if wd == "/" {
			return "", errors.New("can't find " + mkfile.PkgDescFile)
		}
		xcfile := filepath.Join(wd, mkfile.PkgDescFile)
		if _, err := os.Stat(xcfile); err == nil {
			return wd, nil
		}
		wd = filepath.Dir(wd)
	}
}

func init() {
	addCommand(newBuildCommand)
}
