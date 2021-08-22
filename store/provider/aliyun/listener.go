package aliyun

import (
	"fmt"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

func NewListener() oss.ProgressListener {
	return &listener{}
}

type listener struct {
	bar     *progressbar.ProgressBar
	startAt time.Time
}

func (l *listener) ProgressChanged(event *oss.ProgressEvent) {
	// switch  event.EventType {
	// case oss.TransferStartedEvent:
	// 	l.bar = progressbar.DefaultBytes(
	// 		event.TotalBytes,
	// 		"文件上传中",
	// 	)
	// case oss.TransferDataEvent:
	// 	l.bar.Add64(event.RwBytes)
	// case oss.TransferCompletedEvent:
	// 	fmt.Printf("\n上传成功\n")
	// case oss.TransferFailedEvent:
	// 	fmt.Printf("\n上传失败\n")
	// default:
	// }

	// 自定义格式
	switch event.EventType {
	case oss.TransferStartedEvent:
		l.bar = progressbar.NewOptions64(event.TotalBytes,
			progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
			progressbar.OptionEnableColorCodes(true),
			progressbar.OptionShowBytes(true),
			progressbar.OptionSetWidth(30),
			progressbar.OptionSetDescription("开始上传:"),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        "=",
				SaucerHead:    ">",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			}),
		)
		l.startAt = time.Now()
		fmt.Printf("文件大小:%v\n", event.TotalBytes)
	case oss.TransferDataEvent:
		l.bar.Add64(event.RwBytes)
	case oss.TransferCompletedEvent:
		fmt.Printf("\n上传完成: 耗时%d秒\n", int(time.Since(l.startAt).Seconds()))
	case oss.TransferFailedEvent:
		fmt.Printf("\n上传失败\n")
	default:
	}
}
