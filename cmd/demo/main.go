package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	"math"
	"net"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/anguloc/zet/pkg/application"
	"github.com/anguloc/zet/pkg/console"
	"github.com/go-vgo/robotgo"
	"github.com/gocolly/colly/v2"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

type Foo struct {
	A struct {
		B struct {
			C string
		}
	}
}

func (f *Foo) Bar() string {
	res := f.A.B.C
	return res
}

func (f *Foo) Qux() string {
	return f.A.B.C
}

func printMemoryStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println("Total allocated in bytes:", m.TotalAlloc)
	fmt.Println("Currently allocated in bytes:", m.Alloc)
	fmt.Println("Number of mallocs:", m.Mallocs)
	fmt.Println("Number of frees:", m.Frees)
}

func Reverse[T any](slice []T) []T {
	for l, r := 0, len(slice)-1; l < r; l, r = l+1, r-1 {
		slice[l], slice[r] = slice[r], slice[l]
	}
	return slice
}

func netScan(addr string) string {
	conn, err := net.DialTimeout("tcp", addr, time.Second)
	if err != nil {
		if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
			return "timeout"
		} else {
			return "refuse"
		}
	}
	_ = conn.Close()
	return "success"
}

func comsumer(ich <-chan res, och chan<- res) {
	for v := range ich {
		v.result = netScan(v.addr)
		och <- v
	}
}

type res struct {
	addr   string
	result string
}

type Asdddf struct {
	Id   int
	Name string
}

func runA(wg *sync.WaitGroup, handle func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		handle()
	}()
}

func runB(id int, name string) string {
	fmt.Printf("id:%d,name:%s\n", id, name)

	time.Sleep(time.Second)
	return name
}

type Adzg struct {
	Res string
}

func start() {
	arr := []Asdddf{
		{
			Id:   1,
			Name: "a",
		},
		{
			Id:   2,
			Name: "b",
		},
		{
			Id:   3,
			Name: "c",
		},
	}
	wg := &sync.WaitGroup{}

	zva := []Adzg{}
	for _, v := range arr {
		tmp := Adzg{}
		name := v.Name
		id := v.Id
		runA(wg, func() {
			tmp.Res = runB(id, name)
		})
		zva = append(zva, tmp)
	}
	wg.Wait()

	time.Sleep(3 * time.Second)

	xcgvsdfgds, _ := json.MarshalIndent(zva, "", "\t")
	println(string(xcgvsdfgds))

	return
}

type AzgadfRsp struct {
	A     int
	B     int
	C     int
	Total int
}

var ErrRoot = errors.New("root")

var ErrA = errors.New("err a")
var ErrB = errors.New("err b")

func main() {
	ctx := context.TODO()

	windowsDemo(ctx)
	return

	ne := errors.Join(ErrRoot, ErrA)
	fmt.Println(errors.Is(ne, ErrRoot))
	fmt.Println(errors.Is(ErrA, ne))
	fmt.Println(errors.Is(ne, ErrB))
	fmt.Println(1231231)

	content := bytes.NewReader([]byte("asdsad"))
	doc, err := goquery.NewDocumentFromReader(content)
	_ = doc
	_ = err

	return

	m := map[string]*AzgadfRsp{}
	key := "a"
	m[key] = &AzgadfRsp{
		A: 1,
		B: 2,
	}

	m[key].Total = m[key].A + m[key].B

	return
	start()

	return

	ich := make(chan res, 1024)
	och := make(chan res, 1024)
	sch := make(chan struct{})

	for i := 0; i < 200; i++ {
		go func() {
			comsumer(ich, och)
		}()
	}

	fmt.Println("start")
	go func() {
		for end := 1; end < 256; end++ {
			ipStr := fmt.Sprintf("192.168.1.%d", end)
			for port := 6000; port < 10000; port++ {
				addr := fmt.Sprintf("%s:%d", ipStr, port)
				ich <- res{
					addr: addr,
				}
			}
		}
	}()

	var s []string
	go func() {
		for r := range och {
			fmt.Printf("addr:%s,res:%s\n", r.addr, r.result)
			if r.result == "success" {
				s = append(s, r.addr)
			}
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Printf("loop,time:%s,%d,%d\n", time.Now().Format(time.DateTime), len(ich), len(och))
			if len(ich) == 0 && len(och) == 0 {
				sch <- struct{}{}
				break
			}
		}
	}()

	<-sch
	for {
		time.Sleep(time.Second)
		if len(ich) == 0 && len(och) == 0 {
			close(ich)
			close(och)
			break
		}
	}

	fmt.Println("end")
	for _, v := range s {
		fmt.Println("res:" + v)
	}

	return
	tc()

	return

	f := &Foo{}
	_ = f

	for i := 0; i < 4096; i++ {
		f.A.B.C += "a"
	}

	fmt.Println("-------------------------")
	f.Bar()
	fmt.Println("-------------------------")
	f.Qux()

	return
	console.SetLevel(console.DebugLevel, console.InfoLevel, console.WarnLevel, console.ErrorLevel)
	app := application.New()

	app.RegisterWorker("script", application.NewScript(handle))

	_ = app.Init(ctx)

	if err := app.Run(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type Printable interface {
	InitPrinter() Printable
	DoPrint(first bool, num int, outCh chan<- int)
}

type LockPrinter struct {
	state int
	lock  sync.Mutex
	start bool
}

func (p *LockPrinter) InitPrinter() Printable {
	p.state = 0
	p.lock = sync.Mutex{}
	p.start = false

	return p // 固定返回p
}

func (p *LockPrinter) DoPrint(first bool, num int, outCh chan<- int) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.state+1 == num {
		outCh <- num // 在合适的时机写入到 outCh
		p.state = (p.state + 1) % 4
	}
}

func joey() {

}

func ft() {
	a := lo.Max([]int{1, 2, 3, 4})
	fmt.Println(a)
}

func tc() {
	ctx := context.TODO()
	var root = cobra.Command{
		Use:     "demo",
		Version: "0.0.1",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("adasd")
		},
	}

	root.AddCommand(&cobra.Command{
		Use: "m",
		Run: func(cmd *cobra.Command, args []string) {
		},
	})

	if err := root.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		return
	}
}

func openBroser() {

	out1 := &bytes.Buffer{}
	out2 := &bytes.Buffer{}

	var c *exec.Cmd
	if runtime.GOOS == "windows" {
		c = exec.Command("cmd", "/c", "start", "http://www.baidu.com")
	} else if runtime.GOOS == "linux" {
		c = exec.Command("cmd", "/c", "start", "http://www.baidu.com")
	} else {
		fmt.Println("no os")
		return
	}

	c.Stdout = out1
	c.Stderr = out2
	err := c.Start()
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println(out1)
	fmt.Println(out2)

	return

}

func getList() []string {
	res := make([]string, 0, 10000)
	for i := 0; i < 8000; i++ {
		res = append(res, "zcasd")
	}
	return res
}

func mockHttp(s string) string {
	time.Sleep(1 * time.Second)
	return "asd"
}

func collyDemo() {
	c := colly.NewCollector()
	_ = c
}

func handle(ctx context.Context, param *application.ScriptParam) {
	for param.IsRun() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(param.IsRun())
	}
	time.Sleep(100 * time.Millisecond)
}

func touchPanic() {
	var err error
	res, _ := func() (string, error) {
		defer func() {
			if perr := recover(); perr != nil {
				err = fmt.Errorf("asdasdasd,%+v", perr)
			}
		}()
		panic("asdasda")
	}()
	_ = res
	fmt.Println(err)
}

func limit1() {
	// wg := &sync.WaitGroup{}
	// t := semaphore.NewWeighted(3)
	//
	// list := getList()
	//
	// for _, s := range list {
	//	wg.Add(1)
	//	if aerr := t.Acquire(ctx, 1); aerr != nil {
	//		wg.Done()
	//		continue
	//	}
	//	go func(s string) {
	//		defer wg.Done()
	//		defer t.Release(1)
	//		httpRes := mockHttp(s)
	//		fmt.Printf("%s,res:%s\n", time.Now().Format(time.RFC3339), httpRes)
	//	}(s)
	// }
	//
	// wg.Wait()
}

func windowsDemo(ctx context.Context) {
	var err error
	_ = err

	console.SetLevel(console.DebugLevel, console.InfoLevel, console.WarnLevel, console.ErrorLevel)

	// fpid, err := robotgo.FindIds("Google")
	// fpid, err := robotgo.FindIds("steam")
	// if err != nil {
	// 	fmt.Println("读取窗口错误", err)
	// 	return
	// }
	// fmt.Println(fpid)

	displaysNum := robotgo.DisplaysNum()
	width, height := robotgo.GetScreenSize()
	console.Infof("屏幕数量:%d,宽:%d,高:%d\n", displaysNum, width, height)

	return

	img1 := robotgo.CaptureImg(0, 0, width, height)
	img2 := robotgo.CaptureImg(100, 0, width, height)
	// img1RGBA := image.NewRGBA(img1.Bounds())
	// img2RGBA := image.NewRGBA(img2.Bounds())
	// a, err := FastCompare(img1RGBA, img2RGBA)
	// fmt.Println(a)
	// fmt.Println(err)

	robotgo.Save(img1, "tmp/c.png")
	robotgo.Save(img1, "tmp/c.jpeg")
	// robotgo.Save(img2, "tmp/test_2.png")

	a, b, err := ImgCompare1(img1, img2)
	fmt.Println(a)
	_ = b
	fmt.Println(b)
	fmt.Println(err)
	// err = robotgo.Save(img, "tmp/test_1.png")
	// if err != nil {
	// 	fmt.Println("保存图片错误", err)
	// 	return
	// }
	fmt.Println("succ")
}

func ImgCompare(img1, img2 image.Image) (int64, image.Image, error) {
	bounds1 := img1.Bounds()
	bounds2 := img2.Bounds()
	if bounds1 != bounds2 {
		return math.MaxInt64, nil, fmt.Errorf("image bounds not equal: %+v, %+v", img1.Bounds(), img2.Bounds())
	}

	accumError := int64(0)
	resultImg := image.NewRGBA(image.Rect(
		bounds1.Min.X,
		bounds1.Min.Y,
		bounds1.Max.X,
		bounds1.Max.Y,
	))

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
				resultImg.Set(
					bounds1.Min.X+x,
					bounds1.Min.Y+y,
					color.RGBA{R: 255, A: 255})
			}
		}
	}

	return int64(math.Sqrt(float64(accumError))), resultImg, nil
}

func ImgCompare1(img1, img2 image.Image) (int64, float64, error) {
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

func FastCompare(img1, img2 *image.RGBA) (int64, error) {
	if img1.Bounds() != img2.Bounds() {
		return 0, fmt.Errorf("image bounds not equal: %+v, %+v", img1.Bounds(), img2.Bounds())
	}

	accumError := int64(0)

	for i := 0; i < len(img1.Pix); i++ {
		accumError += int64(sqDiffUInt8(img1.Pix[i], img2.Pix[i]))
	}

	return int64(math.Sqrt(float64(accumError))), nil
}

// func FastCompare1(img1, img2 image.Image) (int64, error) {
// 	bounds1 := img1.Bounds()
// 	bounds2 := img2.Bounds()
// 	if bounds1 != bounds2 {
// 		return math.MaxInt64, fmt.Errorf("image bounds not equal: %+v, %+v", img1.Bounds(), img2.Bounds())
// 	}
// }

func sqDiffUInt8(x, y uint8) uint64 {
	d := uint64(x) - uint64(y)
	return d * d
}

func sqDiffUInt32(x, y uint32) uint64 {
	d := uint64(x) - uint64(y)
	return d * d
}
