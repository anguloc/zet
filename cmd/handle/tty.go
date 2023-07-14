package handle

import (
	"sync"

	"github.com/mattn/go-tty"
)

type iTTY struct {
	mu  *sync.Mutex
	tty *tty.TTY
}

var it = newITTY()

func newITTY() *iTTY {
	return &iTTY{
		mu: &sync.Mutex{},
	}
}

func (i *iTTY) open() (*tty.TTY, error) {
	i.mu.Lock()
	defer i.mu.Unlock()
	if i.tty != nil {
		return i.tty, nil
	}
	tmp, err := tty.Open()
	if err != nil {
		return nil, err
	}
	i.tty = tmp
	return i.tty, nil
}

func (i *iTTY) close() {
	if i.tty == nil {
		return
	}
	i.mu.Lock()
	defer i.mu.Unlock()
	if i.tty == nil {
		return
	}
	_ = i.tty.Close()
	i.tty = nil
}
