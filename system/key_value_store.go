package database

import (
	"sync"
)

type ReadMode int

const (
	Snapshot ReadMode = iota
	Committed
)

type WriteMode int

type KeyValueStore struct {
	store     map[string]string
	actions   []Transaction
	readMode  ReadMode
	writeMode WriteMode
	*sync.RWMutex
}

func NewKeyValueStore(readMode ReadMode, writeMode WriteMode) *KeyValueStore {
	return &KeyValueStore{
		store:     make(map[string]string),
		readMode:  readMode,
		writeMode: writeMode,
	}
}

func (kvs KeyValueStore) Get(key string) string {
	kvs.Lock()
	defer kvs.Unlock()

	return kvs.store[key]
}

func (kvs *KeyValueStore) Set(key string, value string) {
	kvs.Lock()
	defer kvs.Unlock()

	kvs.store[key] = value
}

func (kvs *KeyValueStore) Del(key string) {
	kvs.Lock()
	defer kvs.Unlock()

	delete(kvs.store, key)
}
