package inmem

import (
	"errors"
	"sync"
)

type storage struct {
	items map[string]string
	mu    sync.RWMutex
}

func NewStorage() *storage {
	return &storage{
		items: make(map[string]string),
		mu:    sync.RWMutex{},
	}
}

//Insert creates a new key value pair in in-memory db
func (p *storage) Insert(key, value string) (*Pair, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if _, ok := p.items[key]; ok {
		return nil, errors.New("key already exists")
	}
	p.items[key] = value
	return &Pair{
		Key:   key,
		Value: value,
	}, nil

}

//Get fetches the key value pair based on given key
func (p *storage) Get(key string) (*Pair, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if value, ok := p.items[key]; ok {
		return &Pair{
			Key:   key,
			Value: value,
		}, nil
	}
	return nil, errors.New("key not found")
}
