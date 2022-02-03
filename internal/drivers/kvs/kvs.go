package kvs

import (
	"log"
	"sync"
)

var store *KeyValueStore

type KeyValueStore struct {
	sync.RWMutex
	Data map[string]interface{}
}

func New() *KeyValueStore {
	if store == nil {
		store = &KeyValueStore{
			Data: make(map[string]interface{}),
		}
	}
	return store
}

func (kvs *KeyValueStore) GetAll() map[string]interface{} {
	kvs.RLock()
	defer kvs.RUnlock()
	log.Println(kvs.Data)
	r := kvs.Data
	return r
}

func (kvs *KeyValueStore) Get(key string) interface{} {
	kvs.RLock()
	defer kvs.RUnlock()
	return kvs.Data[key]
}

func (kvs *KeyValueStore) Set(key string, value interface{}) {
	kvs.Lock()
	defer kvs.Unlock()
	kvs.Data[key] = value
}

func (kvs *KeyValueStore) Delete(key string) {
	kvs.Lock()
	defer kvs.Unlock()
	delete(kvs.Data, key)
}
