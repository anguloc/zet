package kv

import (
	"errors"

	"github.com/anguloc/zet/pkg/conf"
)

type IKV interface {
	Put(key string, value []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
}

var (
	ErrRecordNotFound = errors.New("key record not found")
)

// NewKv 返回kv存储接口
func New() IKV {
	g := conf.Conf().Global
	if g.KvType == "file" {
		return NewFile()
	}
	return NewFile()
	return NewBigCache()
}
