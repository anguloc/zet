package jmcomic

import (
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/anguloc/zet/internal/app/jmcomic"
	"github.com/anguloc/zet/pkg/console"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	imgPath, err := cmd.Flags().GetString("img_path")
	if err != nil {
		console.Errorf("img_path参数错误[%s]", err)
		return
	}
	if len(imgPath) > 0 {
		chapterId, _ := cmd.Flags().GetUint64("chapter_id")
		single(ctx, chapterId, imgPath)
		return
	}
	dir, err := cmd.Flags().GetString("dir")
	if err != nil || len(dir) == 0 {
		dir, err = os.Getwd()
		if err != nil {
			console.Errorf("获取执行目录文件失败:[%s]\n", err)
			return
		}
	}

	// 判断是否为目录
	if !isDir(dir) {
		console.Errorf("[%s]不是目录或无法访问\n", dir)
		return
	}

	jm := jmcomic.NewJmComic()

	resultDir := filepath.Join(dir, "result")
	fileList, err := jm.ReadDirJmHtml(ctx, dir)
	if err != nil {
		console.Errorf("读取文件列表失败[%s]\n", err)
		return
	}
	if len(fileList) == 0 {
		console.Info("没有获取到文件列表")
		return
	}
	for _, fd := range fileList {
		photoList, pErr := jm.FindImageFiles(ctx, fd)
		if pErr != nil {
			console.Errorf("读取图片失败[%s]\n", fd.Name)
			console.Errorf("错误内容[%s]\n", pErr)
			continue
		}
		if len(photoList) == 0 {
			console.Infof("没有对应图片[%s]\n", fd.Name)
			continue
		}

		if err = genDir(resultDir); err != nil {
			console.Errorf("生成结果文件夹[%s]失败[%s]\n", resultDir, err)
			return
		}
		resPath := filepath.Join(resultDir, filepath.Base(fd.Name))
		if err = genDir(resPath); err != nil {
			console.Errorf("生成结果文件夹[%s]失败[%s]\n", resultDir, err)
			return
		}
		for _, pd := range photoList {
			imageResource, err1 := jm.TransFile(ctx, fd.Id, pd)
			if err1 != nil {
				console.Errorf("生成图片失败[%s][%s]\n]", pd.Path, err1)
				return
			}
			filePath := filepath.Join(resPath, fmt.Sprintf("%s.jpeg", pd.NameNotExt))

			if err = writeImage(filePath, imageResource, 75); err != nil {
				console.Errorf("生成图片失败[%s]", err)
				return
			}
		}
	}
}

// 单张图片
func single(ctx context.Context, chapterId uint64, imgPath string) {
	if chapterId <= 0 {
		console.Errorf("章节id异常\n")
		return
	}
	if !fileExists(imgPath) {
		console.Errorf("[%s]文件不存在\n", imgPath)
		return
	}
	jm := jmcomic.NewJmComic()
	pd, err := jm.SingleImage(imgPath)
	if err != nil {
		console.Errorf("解析失败[%s]", err)
		return
	}
	imageResource, err := jm.TransFile(ctx, chapterId, pd)
	if err != nil {
		console.Errorf("生成图片失败[%s]\n]", err)
		return
	}

	filePath := filepath.Join(filepath.Dir(imgPath), fmt.Sprintf("%s.jpeg", pd.NameNotExt))
	if err = writeImage(filePath, imageResource, 100); err != nil {
		console.Errorf("生成图片失败[%s]", err)
		return
	}
}

func writeImage(dst string, result image.Image, quality int) error {
	// 创建或覆盖目标文件
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() { _ = destFile.Close() }()

	err = jpeg.Encode(destFile, result, &jpeg.Options{Quality: quality})
	if err != nil {
		fmt.Println("写目标失败")
	}
	return nil
}

func genDir(dir string) error {
	// 判断文件夹是否存在，不存在就创建
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		// 文件夹不存在，创建文件夹
		if err = os.MkdirAll(dir, 0666); err != nil {
			return fmt.Errorf("创建文件夹失败: %w", err)
		}
		console.Infof("创建结果文件夹[%s]\n", dir)
	} else if err != nil {
		// 其他错误
		return fmt.Errorf("检查文件夹失败: %w", err)
	} else if !info.IsDir() {
		// 路径存在，但不是文件夹
		return fmt.Errorf("[%s]已存在，但不是文件夹", dir)
	}
	return nil
}

func isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err == nil {
		return !info.IsDir()
	}
	return false
}
