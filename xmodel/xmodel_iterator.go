package xmodel

import (
	"github.com/xuperchain/xuperunion/kv/kvdb"
	"github.com/xuperchain/xuperunion/pb"
	"github.com/xuperchain/xuperunion/xmodel/pb"
)

// XMIterator data structure for XModel Iterator
type XMIterator struct {
	iter  kvdb.Iterator
	model *XModel
	err   error
}

// Data get data pointer to VersionedData for XMIterator
func (di *XMIterator) Data() *xmodel_pb.VersionedData {
	version := di.iter.Value()
	verData, err := di.model.fetchVersionedData(string(version))
	di.err = err
	return verData
}

// Next check if next element exist
func (di *XMIterator) Next() bool {
	return di.iter.Next()
}

// First ...
func (di *XMIterator) First() bool {
	return di.iter.First()
}

// Key get key for XMIterator
func (di *XMIterator) Key() []byte {
	tablePrefixLen := len(pb.ExtUtxoTablePrefix)
	kvdbKey := di.iter.Key()
	return kvdbKey[tablePrefixLen:len(kvdbKey)]
}

// Error return error info for XMIterator
func (di *XMIterator) Error() error {
	kverr := di.iter.Error()
	if kverr != nil {
		return kverr
	}
	return di.err
}

// Release release XMIterator
func (di *XMIterator) Release() {
	di.iter.Release()
}
