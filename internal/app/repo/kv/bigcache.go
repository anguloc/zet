package kv

import (
	"context"
	"errors"
	"io"
	"log"
	"sync"

	"github.com/allegro/bigcache/v3"
)

var (
	bcLock = &sync.Mutex{}
	bc     *bigcache.BigCache
	bcInit bool
)

type BigCache struct {
}

func NewBigCache() *BigCache {
	return &BigCache{}
}

func (b *BigCache) Put(name string, value []byte) error {
	if err := initBigCache(); err != nil {
		return err
	}
	return bc.Set(name, value)
}

func (b *BigCache) Get(name string) ([]byte, error) {
	if err := initBigCache(); err != nil {
		return nil, err
	}
	value, err := bc.Get(name)
	if errors.Is(err, bigcache.ErrEntryNotFound) {
		return nil, ErrRecordNotFound
	}
	return value, err
}

func (b *BigCache) Delete(name string) error {
	if err := initBigCache(); err != nil {
		return err
	}
	return bc.Delete(name)
}

func initBigCache() error {
	if bcInit {
		return nil
	}
	bcLock.Lock()
	defer bcLock.Unlock()
	if bcInit {
		return nil
	}
	config := NewBigCacheConfig()
	c, err := bigcache.New(context.Background(), config)
	if err != nil {
		return err
	}

	bc = c
	bcInit = true

	return nil
}

func NewBigCacheConfig() bigcache.Config {
	config := bigcache.DefaultConfig(0)
	config.CleanWindow = 0
	config.Verbose = false
	config.Logger = log.New(io.Discard, "", 0)
	config.HardMaxCacheSize = 4096

	return config
}
