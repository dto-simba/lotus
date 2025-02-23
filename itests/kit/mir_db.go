package kit

import (
	"context"
	"sync"

	ds "github.com/ipfs/go-datastore"

	"github.com/filecoin-project/lotus/chain/consensus/mir/db"
)

var _ db.DB = (*TestDB)(nil)

type TestDB struct {
	db   map[ds.Key][]byte
	lock sync.Mutex
}

func NewTestDB() *TestDB {
	return &TestDB{
		db: make(map[ds.Key][]byte),
	}
}

func (kv *TestDB) Get(ctx context.Context, key ds.Key) (value []byte, err error) {
	kv.lock.Lock()
	defer kv.lock.Unlock()
	v, ok := kv.db[key]
	if !ok {
		return nil, ds.ErrNotFound
	}
	return v, nil
}

func (kv *TestDB) Put(ctx context.Context, key ds.Key, value []byte) error {
	kv.lock.Lock()
	defer kv.lock.Unlock()
	kv.db[key] = value
	return nil
}

func (kv *TestDB) Delete(ctx context.Context, key ds.Key) error {
	kv.lock.Lock()
	defer kv.lock.Unlock()
	_, ok := kv.db[key]
	if !ok {
		return ds.ErrNotFound
	}
	delete(kv.db, key)
	return nil
}
