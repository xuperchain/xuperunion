package p2pv1

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"sync"

	"github.com/pkg/errors"
	log "github.com/xuperchain/log15"
	p2pPb "github.com/xuperchain/xuperchain/core/p2p/pb"
	p2p_pb "github.com/xuperchain/xuperchain/core/p2p/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Conn maintain the connection of a node
type Conn struct {
	// addr:"IP:Port"
	id     string
	lg     log.Logger
	conn   *grpc.ClientConn
	quitCh chan bool
	lock   *sync.RWMutex
}

func NewConn(lg log.Logger, addr string, certPath, serviceName string, maxMsgSize int) (*Conn, error) {
	conn := &Conn{
		id:     addr,
		lg:     lg,
		quitCh: make(chan bool, 1),
		lock:   &sync.RWMutex{},
	}
	if err := conn.NewGrpcClient(maxMsgSize, certPath, serviceName); err != nil {
		lg.Error("NewConn error", "error", err.Error())
		return nil, err
	}
	return conn, nil
}

func genCreds(certPath, serviceName string) (credentials.TransportCredentials, error) {
	bs, err := ioutil.ReadFile(certPath + "/cacert.pem")
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	ok := certPool.AppendCertsFromPEM(bs)
	if !ok {
		return nil, err
	}

	certificate, err := tls.LoadX509KeyPair(certPath+"/cert.pem", certPath+"/private.key")
	if err != nil {
		return nil, err
	}
	creds := credentials.NewTLS(
		&tls.Config{
			ServerName:   serviceName,
			Certificates: []tls.Certificate{certificate},
			RootCAs:      certPool,
			ClientCAs:    certPool,
			ClientAuth:   tls.RequireAndVerifyClientCert,
		})
	return creds, nil
}

func (c *Conn) NewGrpcClient(maxMsgSize int, certPath string, serviceName string) error {
	conn := &grpc.ClientConn{}
	creds, err := genCreds(certPath, serviceName)
	if err != nil {
		return err
	}
	conn, err = grpc.Dial(c.id, grpc.WithTransportCredentials(creds), grpc.WithMaxMsgSize(maxMsgSize))
	if err != nil {
		return errors.New("New grpcs conn error")
	}
	c.conn = conn
	return nil
}

func (c *Conn) newClient(ctx context.Context) (p2p_pb.P2PService_SendP2PMessageClient, error) {
	client := p2p_pb.NewP2PServiceClient(c.conn)
	return client.SendP2PMessage(ctx)
}

// SendMessage send message to a peer
func (c *Conn) SendMessage(ctx context.Context, msg *p2pPb.XuperMessage) error {
	client, err := c.newClient(ctx)
	if err != nil {
		c.lg.Error("SendMessage new client error")
		return err
	}
	return client.Send(msg)
}

// SendMessageWithResponse send message to a peer with responce
func (c *Conn) SendMessageWithResponse(ctx context.Context, msg *p2pPb.XuperMessage) (*p2pPb.XuperMessage, error) {
	client, err := c.newClient(ctx)
	if err != nil {
		c.lg.Error("SendMessageWithResponse new client error")
		return nil, err
	}
	err = client.Send(msg)
	if err != nil {
		c.lg.Error("SendMessageWithResponse error", "error", err.Error())
		return nil, err
	}
	res := &p2pPb.XuperMessage{}
	waitc := make(chan struct{})
	go func() {
		for {
			res, err = client.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				c.lg.Error("SendMessageWithResponse Recv error", "error", err.Error())
				return
			}
		}
	}()
	client.CloseSend()
	<-waitc
	return res, err
}

func (c *Conn) Close() {
	c.conn.Close()
}

func (c *Conn) GetConnID() string {
	return c.id
}
