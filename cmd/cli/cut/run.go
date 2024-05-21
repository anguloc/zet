package cut

import (
	"fmt"

	"github.com/anguloc/zet/pkg/console"
	"github.com/anguloc/zet/pkg/safe"
	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
)

var (
	displaysNum   int    // 屏幕数量
	width, height int    // 屏幕分辨率
	x, y, w, h    int    // 截图数据
	file          string // 保存文件名
)

func Run(cmd *cobra.Command, args []string) {
	displaysNum = robotgo.DisplaysNum()
	width, height = robotgo.GetScreenSize()

	console.Infof("屏幕数量:%d,宽:%d,高:%d\n", displaysNum, width, height)
	console.Infof("截图数据,x:%d,y:%d,w:%d,h:%d\n", x, y, w, h)

	if file == "" {
		file = fmt.Sprintf("cut_%d_%d_%d_%d.png", x, y, w, h)
	}

	img := robotgo.CaptureImg(x, y, w, h)
	err := robotgo.Save(img, safe.Path(""))
	if err != nil {
		console.Error("保存文件失败:", err)
		return
	}
	console.Info("保存文件成功,path:", file)
}
