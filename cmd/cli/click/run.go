package click

import (
	"fmt"
	"time"

	"github.com/anguloc/zet/internal/pkg/console"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/spf13/cobra"
)

var run bool

func Run(cmd *cobra.Command, args []string) {
	go switchRun()
	console.Info("start")
	for ; ; time.Sleep(time.Millisecond) {
		if run {
			println(time.Now().Format("15:04:05"))

			robotgo.KeyDown("e")
			robotgo.KeyDown("d")

			time.Sleep(time.Millisecond * 1256)

			robotgo.KeyUp("e")
			robotgo.KeyUp("d")

			robotgo.KeyDown("e")
			robotgo.KeyDown("a")

			time.Sleep(time.Millisecond * 1200)

			robotgo.KeyUp("e")
			robotgo.KeyUp("a")

			//s = time.Now()
			//robotgo.MouseClick("left")
			//robotgo.KeyDown("j")
			//e := time.Since(s)
			//println(e.Milliseconds())
		}
	}
}

func switchRun() {
	_ = hook.MouseDown
	for ; ; time.Sleep(time.Second / 2) {
		if robotgo.AddEvent("f4") {
			if run {
				fmt.Println("停止")
			} else {
				fmt.Println("开始")
			}
			run = !run
		}
	}
}
