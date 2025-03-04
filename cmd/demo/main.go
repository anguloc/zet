package main

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"log"
	"math"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/anguloc/zet/pkg/application"
	"github.com/anguloc/zet/pkg/console"
	"github.com/go-vgo/robotgo"
	"github.com/gocolly/colly/v2"
	"github.com/looplab/fsm"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"golang.org/x/image/webp"
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

	// src := "C:\\Users\\anguloc\\Desktop\\t1/00002.webp"
	// src = "C:\\Users\\anguloc\\Desktop\\t1/1231.jpg"
	// dst := "C:\\Users\\anguloc\\Desktop\\t1/00002_bak.jpeg"
	// transFile(src, dst)
	//
	// return

	testFile(ctx)

	return

	testFsm(ctx)
	return

	aesDemo()
	return

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

func aesDemo() {
	origData := []byte(strings.Repeat("a", 16))       // 待加密的数据
	key := []byte("ABCDEFGHIJKLMNOPABCDEFGHIJKLMNOP") // 加密的密钥
	log.Println("原文：", string(origData))

	log.Println("------------------ CBC模式 --------------------")
	encrypted := AesEncryptCBC(origData, key)

	// fmt.Println(key)
	// fmt.Println(encrypted)
	//
	// return

	fmt.Println("数据长度:", len(origData))
	fmt.Println("加密后长度:", len(encrypted))
	// fmt.Println("str:", string(encrypted))

	// log.Println("明文(hex)：", hex.EncodeToString(origData))
	log.Println("明文1(hex)：", len(hex.EncodeToString(origData)))
	// log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文1(hex)：", len(hex.EncodeToString(encrypted)))
	log.Println("密文1(base64)：", len(base64.StdEncoding.EncodeToString(encrypted)))

	// log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := AesDecryptCBC(encrypted, key)
	// log.Println("解密结果：", string(decrypted))
	return

	log.Println("------------------ ECB模式 --------------------")
	encrypted = AesEncryptECB(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = AesDecryptECB(encrypted, key)
	log.Println("解密结果：", string(decrypted))

	log.Println("------------------ CFB模式 --------------------")
	encrypted = AesEncryptCFB(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = AesDecryptCFB(encrypted, key)
	log.Println("解密结果：", string(decrypted))
}

// =================== CBC ======================
func AesEncryptCBC(origData []byte, key []byte) (encrypted []byte) {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()               // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize) // 补全码
	fmt.Println("补全后长度：", len(origData))
	fmt.Println("aa-", origData)
	fmt.Println("aa-", string(origData))
	// fmt.Println("aaaa-", string(key[:blockSize]))
	// fmt.Println("bbbb-", key[:blockSize])
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted = make([]byte, len(origData))                     // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return encrypted
}
func AesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)                              // 分组秘钥
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	decrypted = pkcs5UnPadding(decrypted)                       // 去除补全码
	return decrypted
}
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// =================== ECB ======================
func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}
func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// =================== CFB ======================
func AesEncryptCFB(origData []byte, key []byte) (encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted
}
func AesDecryptCFB(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted
}

type A struct {
	Data string
}

func NewA(a int) (*A, error) {
	if a == 1 {
		return nil, fmt.Errorf("aa")
	}
	return &A{}, nil
}

func testFsm(ctx context.Context) {

	var a int
	a = 1
	m, _ := NewA(a)
	fmt.Println(m.Data)

	return
	f := fsm.NewFSM("init", fsm.Events{
		{
			Name: "create",
			Src:  []string{"init"},
			Dst:  "create",
		},
		{
			Name: "init",
			Src:  []string{"create"},
			Dst:  "init",
		},
	}, fsm.Callbacks{
		"create": func(ctx context.Context, event *fsm.Event) {
			fmt.Printf("执行create事件，当前状态:%s\n", event.Event)
		},
		"init": func(ctx context.Context, event *fsm.Event) {
			fmt.Printf("执行init事件，当前状态:%s\n", event.Event)
		},

		"before_create": func(ctx context.Context, event *fsm.Event) {
			fmt.Printf("create事件之前，当前状态:%s\n", event.Event)
		},
		"after_create": func(ctx context.Context, event *fsm.Event) {
			fmt.Printf("create事件之后，当前状态:%s\n", event.Event)
		},

		"before_event": func(ctx context.Context, event *fsm.Event) {
			fmt.Printf("before_event触发，当前状态:%s\n", event.Event)
		},
		"leave_state": func(ctx context.Context, event *fsm.Event) {
			fmt.Printf("leave_state触发，当前状态:%s\n", event.Event)
		},
		"enter_state": func(ctx context.Context, event *fsm.Event) {
			fmt.Printf("enter_state触发，当前状态:%s\n", event.Event)
		},
		"after_event": func(ctx context.Context, event *fsm.Event) {
			fmt.Printf("after_event触发，当前状态:%s\n", event.Event)
		},
	})

	err := f.Event(ctx, "create")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("---------------")

	err = f.Event(ctx, "init")
	if err != nil {
		fmt.Println(err)
	}

}

func copyAndRenameFile(src, dst string) error {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("无法打开源文件: %v", err)
	}
	defer srcFile.Close()

	// 创建目标文件
	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("无法创建目标文件: %v", err)
	}
	defer dstFile.Close()

	// 复制文件内容
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("复制文件失败: %v", err)
	}

	// 确保文件内容已刷新到磁盘
	err = dstFile.Sync()
	if err != nil {
		return fmt.Errorf("刷新文件失败: %v", err)
	}

	fmt.Printf("文件已复制并重命名: %s -> %s\n", src, dst)
	return nil
}

func testFile(ctx context.Context) {
	dir := "C:\\Users\\anguloc\\Desktop\\t2/"
	dir1 := "C:\\Users\\anguloc\\Desktop\\t2/"

	dir2 := `C:\Users\anguloc\Desktop\aa\result\a`
	dir3 := `C:\Users\anguloc\Desktop\aa\result\b`

	files1, err := os.ReadDir(dir2)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range files1 {
		tmp1 := strings.TrimRight(f.Name(), ".jpeg")
		id, _ := strconv.Atoi(tmp1)
		id += 434
		fileName := fmt.Sprintf("%05d.jpeg", id)

		copyAndRenameFile(fmt.Sprintf("%s/%s", dir2, f.Name()), fmt.Sprintf("%s/%s", dir3, fileName))
	}

	return

	// 读取目录下所有文件
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	// var hf []string

	type zas struct {
		Src string
		Dst string
	}
	var hf1 []zas
	// 遍历文件和子目录
	for _, file := range files {
		filePath := fmt.Sprintf("C:\\Users\\anguloc\\Desktop\\2/%s", file.Name())
		if file.IsDir() {
			fmt.Println("存在文件夹")
			// readFilesInDirectory(filePath) // 递归调用
		} else {
			// 处理文件
			// fmt.Println(filePath)
			if strings.HasSuffix(filePath, ".webp") && strings.HasPrefix(file.Name(), "00") {
				// hf = append(hf, filePath)
				hf1 = append(hf1, zas{filePath, dir1 + file.Name()})
			}
		}
	}

	for _, s := range hf1 {
		transFile(s.Src, s.Dst)
		fmt.Println(s)
	}

	fmt.Println(len(hf1))
}

func readFilesInDirectory(directory string) {
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println(err)
	}

	// 遍历文件和子目录
	for _, file := range files {
		filePath := fmt.Sprintf("%s/%s", directory, file.Name())
		if file.IsDir() {
			readFilesInDirectory(filePath) // 递归调用
		} else {
			// 处理文件
			fmt.Println(filePath)
		}
	}
}

func transFile(src, dst string) error {
	// 打开源文件
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// 图片解码
	// img, _, err := image.Decode(sourceFile)
	img, err := webp.Decode(sourceFile)
	if err != nil {
		fmt.Println("解码图片", err)
		return err
	}

	// imgResource := image.NewRGBA(img.Bounds())
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	result := image.NewRGBA(image.Rect(0, 0, width, height))
	// preImgHeight := height / 4
	// draw.Draw(result, image.Rect(0, 0*preImgHeight, width, 1*preImgHeight), img, image.Point{X: 0,
	// 	Y: 3 * preImgHeight}, draw.Src)
	// draw.Draw(result, image.Rect(0, 1*preImgHeight, width, 2*preImgHeight), img, image.Point{X: 0,
	// 	Y: 2 * preImgHeight}, draw.Src)
	// draw.Draw(result, image.Rect(0, 2*preImgHeight, width, 3*preImgHeight), img, image.Point{X: 0,
	// 	Y: 1 * preImgHeight}, draw.Src)
	// draw.Draw(result, image.Rect(0, 3*preImgHeight, width, 4*preImgHeight), img, image.Point{X: 0,
	// 	Y: 0 * preImgHeight}, draw.Src)
	//
	// // 创建或覆盖目标文件
	// destFile, err := os.Create(dst)
	// if err != nil {
	// 	fmt.Println("创建dst失败")
	// 	return err
	// }
	// defer destFile.Close()
	//
	// err = jpeg.Encode(destFile, result, &jpeg.Options{Quality: 100})
	// if err != nil {
	// 	fmt.Println("写目标失败")
	// }
	// return err

	// 1 -> 4
	// 2 -> 3
	// 3 -> 2
	// 4 -> 1

	// https://github.com/jiayaoO3O/18-comic-finder/blob/master/src/main/java/io/github/jiayaoO3O/finder/service/TaskService.java
	rule := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	chapterId := 448545

	piece := 10
	if chapterId >= 268850 {
		// src 文件名并去掉后缀
		name := filepath.Base(src)
		var photoId string
		if strings.HasSuffix(name, ".webp") {
			photoId = strings.TrimRight(name, ".webp")
		}
		// md5
		s := fmt.Sprintf("%d%s", chapterId, photoId)
		fmt.Println(s)
		h := md5.Sum([]byte(s))
		m := hex.EncodeToString(h[:])
		c := m[len(m)-1]
		mod := 10
		if chapterId >= 421926 {
			mod = 8
		}
		piece = rule[int(c)%mod]
	}
	// fmt.Println(c)
	// fmt.Println(piece)

	// piece := 4
	preImgHeight := height / piece
	for i := 0; i < piece; i++ {
		var (
			item  image.Rectangle
			point image.Point
		)
		// 从上到下第几块
		if i == piece-1 {
			// 漫画的高度除以块数时,不一定是整数,此时漫画的第一块高度要算上剩余的像素.
			item = image.Rect(0, i*preImgHeight, width, height)
			point = image.Point{X: 0, Y: 0}
		} else {
			item = image.Rect(0, i*preImgHeight, width, (i+1)*preImgHeight)
			point = image.Point{X: 0, Y: (piece - i - 1) * preImgHeight}
		}
		draw.Draw(result, item, img, point, draw.Src)
	}

	// 创建或覆盖目标文件
	destFile, err := os.Create(dst)
	if err != nil {
		fmt.Println("创建dst失败")
		return err
	}
	defer destFile.Close()

	err = jpeg.Encode(destFile, result, &jpeg.Options{Quality: 100})
	if err != nil {
		fmt.Println("写目标失败")
	}
	return err

	// 复制内容从源文件到目标文件
	_, err = io.Copy(destFile, sourceFile)
	return err
}
