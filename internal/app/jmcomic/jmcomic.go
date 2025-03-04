package jmcomic

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"image"
	"image/draw"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/image/webp"
)

type JmComic struct {
}

func NewJmComic() *JmComic {
	return &JmComic{}
}

// ReadDirJmHtml 读目录下第一层的html
func (j JmComic) ReadDirJmHtml(ctx context.Context, dir string) ([]*FileData, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	fs := make([]*FileData, 0, len(files))
	// 暂时不管软连接等特殊情况
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.Contains(file.Name(), ".html") {
			continue
		}
		item, err1 := j.parseHtmlFile(ctx, file.Name(), filepath.Join(dir, file.Name()))
		if err1 != nil {
			return nil, err1
		}
		fs = append(fs, item)
	}
	return fs, nil
}

func (j JmComic) parseHtmlFile(_ context.Context, name, f string) (*FileData, error) {
	res := &FileData{
		Name:       name,
		NameNotExt: strings.TrimRight(name, ".html"),
		Path:       f,
	}

	// 读html
	content, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	doc, err := htmlquery.Parse(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}

	// 读html中meta属性为og:url中的content值
	node := htmlquery.FindOne(doc, `//meta[@property="og:url"]`)
	if node == nil {
		return nil, fmt.Errorf("no og.url found")
	}
	url := htmlquery.SelectAttr(node, "content")
	if url == "" {
		return nil, fmt.Errorf("no og.url content found")
	}

	// 移除空格
	url = strings.TrimSpace(url)
	// 移除最后一个?以及后面的值
	if i := strings.LastIndex(url, "?"); i != -1 {
		url = url[:i]
	}
	url = strings.TrimRight(url, "/")
	// 取url的/后面的数字
	i := strings.LastIndex(url, "/")
	if i == -1 {
		return nil, fmt.Errorf("no og.url id found")
	}
	idStr := url[i+1:]
	res.Id, err = strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid og.url id parse err,%w", err)
	}

	return res, nil
}

// FindImageFiles 图片在html下同层的同名文件夹下，目前都当作webp处理
func (j JmComic) FindImageFiles(_ context.Context, fd *FileData) ([]*PhotoData, error) {
	dir := filepath.Dir(fd.Path)
	ext := filepath.Ext(fd.Path)
	filename := filepath.Base(fd.Path)

	// 取图片文件目录
	name := filename[0 : len(filename)-len(ext)]
	imagesDir := filepath.Join(dir, fmt.Sprintf("%s_files", name))
	st, err := os.Stat(imagesDir)
	if err != nil {
		return nil, fmt.Errorf("images dir stat err,%w", err)
	}
	if !st.IsDir() {
		return nil, fmt.Errorf("images dir is not a folder,[%s]", imagesDir)
	}

	files, err := os.ReadDir(imagesDir)
	if err != nil {
		return nil, fmt.Errorf("images dir read err,%w", err)
	}

	var matches []string
	// 遍历一级目录
	// webp文件名都是{00001.webp}从00001开始，先简单搞，解析html比较麻烦，先按当前规则来简单搞
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".webp") && strings.HasPrefix(file.Name(), "00") && len(file.Name()) == 10 {
			matches = append(matches, filepath.Join(imagesDir, file.Name()))
		}
	}
	// matches, err := filepath.Glob(fmt.Sprintf("%s/*.webp", imagesDir))
	// if err != nil {
	// 	return nil, fmt.Errorf("find images files,%w", err)
	// }

	res := make([]*PhotoData, 0, len(matches))
	for _, match := range matches {
		photoName := filepath.Base(match)
		photoExt := filepath.Ext(photoName)
		idStr := photoName[0 : len(photoName)-len(photoExt)]
		photoId, idErr := strconv.ParseUint(idStr, 10, 64)
		if idErr != nil {
			return nil, fmt.Errorf("photo id parse,id:[%s],err:%w", photoName, idErr)
		}
		res = append(res, &PhotoData{
			Name:       photoName,
			NameNotExt: idStr,
			Id:         photoId,
			IdStr:      idStr,
			Path:       match,
		})
	}

	return res, nil
}

// TransFile 读文件后处理反爬后转到目标文件夹
func (j JmComic) TransFile(_ context.Context, fd *FileData, pd *PhotoData, ) (*image.RGBA, error) {
	// 打开源文件
	sourceFile, err := os.Open(pd.Path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = sourceFile.Close() }()

	// 图片解码
	var img image.Image
	if filepath.Ext(pd.Path) == "webp" {
		img, err = webp.Decode(sourceFile)
	} else {
		img, _, err = image.Decode(sourceFile)
	}
	if err != nil {
		return nil, err
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	result := image.NewRGBA(image.Rect(0, 0, width, height))
	piece := calcPiece(fd.Id, pd.IdStr)
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

	return result, nil
}

func calcPiece(chapterId uint64, photoId string) int {
	// https://github.com/jiayaoO3O/18-comic-finder/blob/master/src/main/java/io/github/jiayaoO3O/finder/service/TaskService.java
	rule := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	piece := 10
	if chapterId >= 268850 {
		// md5
		s := fmt.Sprintf("%d%s", chapterId, photoId)
		h := md5.Sum([]byte(s))
		m := hex.EncodeToString(h[:])
		c := m[len(m)-1]
		mod := 10
		if chapterId >= 421926 {
			mod = 8
		}
		piece = rule[int(c)%mod]
	}
	return piece
}
