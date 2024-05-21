package gbfr

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"image"
	"math"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/anguloc/zet/pkg/application"
	"github.com/anguloc/zet/pkg/console"
	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
)

var (
	displaysNum   int                // 屏幕数量
	width, height int                // 屏幕分辨率
	stopFunc      context.CancelFunc // 退出方法

	mainPid int  // 游戏主进程id
	isFocus bool // 游戏窗口是否有焦点
	display int  // 使用屏幕的序号
	running bool
)

const (
	gameTitle = "granblue_fantasy_relink" // 游戏名称
)

func Run(cmd *cobra.Command, args []string) {
	var err error
	if runtime.GOOS != "windows" {
		console.Error("只支持windows")
		return
	}
	displaysNum = robotgo.DisplaysNum()
	width, height = robotgo.GetScreenSize()

	display, err = cmd.Flags().GetInt("display")
	if err != nil || display < 0 || display > displaysNum-1 {
		console.Info("读取屏幕序号失败，使用默认值：0")
		display = 0
	}

	ctx := cmd.Context()
	ctx, stopFunc = context.WithCancel(ctx)

	app := application.New()

	app.RegisterWorker("script", application.NewScript(func(ctx context.Context, param *application.ScriptParam) {
		handle(ctx)
	}))

	_ = app.Init(ctx)

	if err = app.Run(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handle(ctx context.Context) {
	wg := &sync.WaitGroup{}

	initNode()
	help()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// 监听游戏窗口存在以及焦点
		listenGame(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// 执行动作
		playGame(ctx)
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			console.Info("业务监听到退出信号，等待进行中的逻辑完成")
			wg.Wait()
			return
		default:
			// TODO 有个bug，退出时这里会多次打印，顺序问题
			fmt.Printf(">")
			if scanner.Scan() {
				input := strings.TrimSpace(scanner.Text())
				switch input {
				case "help":
					help()
				case "quit":
					stopFunc()
				case "start":
					running = true
					console.Info("开始执行")
				case "stop":
					running = false
					console.Info("停止执行")
				default:
					console.Warn("无效输入：", input)
				}
			} else if err := scanner.Err(); err != nil {
				// 先不管
				fmt.Println("输入错误", err)
			}

		}
	}
}

func help() {
	console.Infof("屏幕数量:%d,宽:%d,高:%d\n", displaysNum, width, height)
	title()
	console.Info(`
start 开始
stop 暂停
quit 退出
help 帮助
`)
}

func title() {
	console.Infof("游戏主进程id:%d,是否获得游戏焦点:%t\n", mainPid, isFocus)
}

func listenGame(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Millisecond * 101):
			isFocus = true
			continue
			if mainPid == 0 {
				pids, err := robotgo.FindIds(gameTitle)
				if err != nil {
					console.Warn("搜索游戏进程错误:", err)
					mainPid = 0
					isFocus = false
					continue
				}
				mainPid = pids[0]
			}
			if ok, err := robotgo.PidExists(mainPid); !ok || err != nil {
				mainPid = 0
				isFocus = false
				continue
			}

			if !isFocus {
				isFocus = mainPid != 0 && robotgo.GetPid() == mainPid
				if isFocus {
					console.Info("获得游戏焦点")
				}
			}
		}
	}
}

func playGame(ctx context.Context) {
	n := 0
	for {
		select {
		case <-ctx.Done():
			return
		// case <-time.After(time.Millisecond * 97):
		case <-time.After(time.Millisecond * 97 * 10):
			n++
			if n == 30 {
				title()
			}
			// 非焦点或者游戏进程不存在，不执行操作
			if !isFocus || !running {
				continue
			}
			feature()
		}
	}
}

type Node struct {
	page       int
	title      string
	featureFn  featureFn
	featureImg image.Image
	cn         []*Node
	fn         actionFn
}

func (n *Node) Search() *Node {
	// if n.featureFn != nil {
	// 	if ok := n.compare(n, n.featureFn(), n.featureImg); ok {
	// 		return n
	// 	}
	// }
	// for _, node := range n.cn {
	// 	if res := node.Search(); res != nil {
	// 		return res
	// 	}
	// }
	for _, node := range n.cn {
		if ok := n.compare(node, node.featureFn(), node.featureImg); ok {
			return node
		}
	}
	return nil
}

func (n *Node) compare(nn *Node, img1, img2 image.Image) bool {
	_, f, _ := ImgCompare(img1, img2)
	if nn.page == PageFeature1 {
		return 900000 < f && f < 1100000
	}
	if nn.page == PageFeature2 {
		return 900000 < f && f < 2000000
	}
	if nn.page == PageFeature2On {
		return 2100000 < f && f < 2200000
	}
	if nn.page == PageFeature2Off {
		return 260000 < f && f < 280000
	}
	if nn.page == PageFeature3 {
		return 2800000 < f && f < 3000000
	}
	if nn.page == PageFeature3On {
		return 1100000 < f && f < 1500000
	}
	if nn.page == PageFeature3Off {
		return 1100000 < f && f < 1500000
	}
	return false
}

func ImgCompare(img1, img2 image.Image) (int64, float64, error) {
	// 粗暴
	bounds1 := img1.Bounds()
	bounds2 := img2.Bounds()
	if bounds1 != bounds2 {
		return math.MaxInt64, 0, fmt.Errorf("image bounds not equal: %+v, %+v", img1.Bounds(), img2.Bounds())
	}

	accumError := int64(0)

	for x := bounds1.Min.X; x < bounds1.Max.X; x++ {
		for y := bounds1.Min.Y; y < bounds1.Max.Y; y++ {
			r1, g1, b1, a1 := img1.At(x, y).RGBA()
			r2, g2, b2, a2 := img2.At(x, y).RGBA()

			diff := int64(sqDiffUInt32(r1, r2))
			diff += int64(sqDiffUInt32(g1, g2))
			diff += int64(sqDiffUInt32(b1, b2))
			diff += int64(sqDiffUInt32(a1, a2))

			if diff > 0 {
				accumError += diff
			}
		}
	}

	n := bounds1.Max.X * bounds1.Max.Y
	score := float64(accumError) / (4 * float64(n))

	return int64(math.Sqrt(float64(accumError))), score, nil
}

func sqDiffUInt32(x, y uint32) uint64 {
	d := uint64(x) - uint64(y)
	return d * d
}

type featureFn func() image.Image
type actionFn func(ctx context.Context)

// const (
// 	arCode = iota + 1
// )
//
// type actionResult struct {
// 	code int
// }
//
// func (a actionResult) isNone() bool {
// 	return a.code == arCode
// }
//
// func noneRes() actionResult {
// 	return actionResult{code: arCode}
// }

var ns *Node

func initNode() {
	ns = &Node{
		page: PageRoot,
		cn: []*Node{
			{
				page:       PageFeature1,
				title:      "第一个结算页面的特征",
				featureFn:  getPage1,
				featureImg: decodeToImg(featurePage1),
			},
			{
				page:       PageFeature2,
				title:      "第二个结算页特征",
				featureFn:  getPage2,
				featureImg: decodeToImg(featurePage2),
				cn: []*Node{
					{
						page:       PageFeature2On,
						title:      "第二个结算页重开标志特征 - 开",
						featureFn:  getPage3,
						featureImg: decodeToImg(featurePage3),
						cn:         nil,
						fn:         nil,
					},
					{
						page:       PageFeature2Off,
						title:      "第二个结算页重开标志特征 - 关",
						featureFn:  getPage3,
						featureImg: decodeToImg(featurePage4),
						cn:         nil,
						fn:         nil,
					},
				},
			},
			{
				page:       PageFeature3,
				title:      "十次弹窗特征",
				featureFn:  getPage4,
				featureImg: decodeToImg(featurePage5),
				cn: []*Node{
					{
						page:       PageFeature3On,
						title:      "十次选择 - 是",
						featureFn:  getPage5,
						featureImg: decodeToImg(featurePage6),
						cn:         nil,
						fn:         nil,
					},
					{
						page:       PageFeature3Off,
						title:      "十次选择 - 否",
						featureFn:  getPage5,
						featureImg: decodeToImg(featurePage7),
						cn:         nil,
						fn:         nil,
					},
				},
			},
		},
	}
}

func decodeToImg(b []byte) image.Image {
	img, _, _ := image.Decode(bytes.NewReader(b))
	return img
}

func feature() (int, bool) {
	// 找当前所在页面
	n := ns.Search()
	if n == nil {
		return 0, false
	}
	console.Infof("[%s]匹配\n", n.title)

	// 当前是哪个特征
	cn := n.Search()
	if cn == nil {
		return 0, false
	}
	console.Infof("[%s]匹配-特征：[%s]\n", n.title, cn.title)

	return 0, false
}

// 拿第一个结算页面特征
func getPage1() image.Image {
	return robotgo.CaptureImg(display*width+480, 150, 240, 100)
}

// 拿第二个结算页特征
func getPage2() image.Image {
	return robotgo.CaptureImg(display*width+200, 700, 130, 40)
}

// 第二个结算页重开标志特征
func getPage3() image.Image {
	return robotgo.CaptureImg(display*width+430, 750, 40, 40)
}

// 十次弹窗特征
func getPage4() image.Image {
	return robotgo.CaptureImg(display*width+850, 350, 230, 50)
}

// 十次选择
func getPage5() image.Image {
	return robotgo.CaptureImg(display*width+920, 650, 80, 100)
}
