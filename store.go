package main

import "sync"

const ErrNoSuchKey = Error("value with specified key does not exist")

type LockableMap struct {
	sync.RWMutex
	m map[string]string
}

var store = LockableMap{m: make(map[string]string)}

func Put(key string, value string) error {
	store.Lock()
	store.m[key] = value
	store.Unlock()

	return nil
}

func Get(key string) (string, error) {
	store.RLock()
	defer store.RUnlock()

	value, ok := store.m[key]
	if !ok {
		return "", ErrNoSuchKey
	}

	return value, nil
}

func Delete(key string) error {
	store.Lock()
	defer store.Unlock()

	delete(store.m, key)

	return nil
}
