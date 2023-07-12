package handle

import (
	"sync"

	"github.com/mattn/go-tty"
)

var (
	ttyMu = &sync.Mutex{}
	exist = false
	t     *tty.TTY
)

func openTTY() error {
	ttyMu.Lock()
	defer ttyMu.Unlock()
	if exist {
		return nil
	}
	tmp, err := tty.Open()
	if err != nil {
		return err
	}
	t = tmp
	exist = true
	return nil
}
func closeTTY() {
	ttyMu.Lock()
	defer ttyMu.Unlock()
	if !exist {
		return
	}
	_ = t.Close()
	t = nil
	exist = false
}
