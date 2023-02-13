package cos

// Author: Hanyongyan

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// SaveUploadedFile 上传文件，返回文件名
// @param1 传入的视频数据
// @param2 视频类型后缀
func SaveUploadedFile(ctx context.Context, data []byte) (flag bool, saveVideoPath, savePhotoPath string) {
	// 设置文件保存路径
	// 文件路径为  根目录/上传时间戳 + uuid
	filename := uuid.New()
	// video/{filename}.mp4 保存到 cos 的路径
	saveVideoPath = fmt.Sprintf("%s%s.mp4", viper.GetString("cos.videoPath"), filename)
	// photo/{filename}.mp4 保存到 cos 的路径
	savePhotoPath = fmt.Sprintf("%s%s.jpg", viper.GetString("cos.photoPath"), filename)
	// os.TempDir()\{filename}.mp4 临时保存视频的路径
	tempVideoPath := os.TempDir() + "\\" + strings.Split(saveVideoPath, "/")[1]

	// 进行临时保存 视频数据
	file, err := os.OpenFile(tempVideoPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return false, "", ""
	}
	// 进行删除临时视频文件
	defer os.Remove(tempVideoPath)
	// 进行关闭文件流
	defer file.Close()
	writer := bufio.NewWriter(file)
	// 将数据写入缓存
	_, err = writer.Write(data)
	// 写入缓存失败
	if err != nil {
		return false, "", ""
	}
	// 将数据从缓存中写入硬盘
	writer.Flush()
	// 写入数据失败
	if err != nil {
		return false, "", ""
	}
	var wg sync.WaitGroup
	wg.Add(2)
	//开启一个协程进行上传视频文件
	go func() {
		Cos.Object.Put(ctx, saveVideoPath, bytes.NewReader(data), nil)
		wg.Done()
	}()

	// 使用 ffmeg 生成截图
	// 临时封面图片文件路径
	tempPhotoPath := fmt.Sprintf("%s\\%s.jpg", os.TempDir(), filename)
	// 进行生成封面图片
	cmd := exec.Command("ffmpeg", "-i", tempVideoPath, tempPhotoPath,
		"-ss", "00:00:00", "-r", "1", "-vframes", "1", "-an", "-vcodec", "mjpeg")
	_ = cmd.Run()
	defer os.Remove(tempPhotoPath)
	go func() {
		Cos.Object.PutFromFile(ctx, savePhotoPath, tempPhotoPath, nil)
		wg.Done()
	}()
	wg.Wait()
	return true, saveVideoPath, savePhotoPath
}
