package kv

import (
	"os"
	"path/filepath"

	"github.com/anguloc/zet/pkg/conf"
	"github.com/anguloc/zet/pkg/safe"
)

const (
	fileDirName = "file_cache"
)

var (
	dir string = getFileDir()
)

// File 非协程安全文件存储，用于调试
type File struct {
}

func NewFile() *File {
	return &File{}
}

func (f *File) Put(name string, value []byte) error {
	filePath := f.getFilePath(name)
	d := filepath.Dir(filePath)
	if err := os.MkdirAll(d, 0766); err != nil {
		return err
	}
	return os.WriteFile(filePath, value, 0666)
}

func (f *File) Get(name string) ([]byte, error) {
	filePath := f.getFilePath(name)
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil, ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}

	return os.ReadFile(filePath)
}

func (f *File) Delete(name string) error {
	filePath := f.getFilePath(name)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil
	}

	return os.Remove(filePath)
}

func (f *File) getFilePath(name string) string {
	return filepath.Join(dir, name)
}

func getFileDir() string {
	rootDir := filepath.Join(conf.Conf().Global.TmpDir, fileDirName)
	return safe.Path(rootDir)
}
