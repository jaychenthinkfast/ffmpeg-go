package main

import (
	"github.com/jaychenthinkfast/ffmpeg-go/pkg/video2image"
	"io/ioutil"
	"k8s.io/klog/v2"
	"os"
	"strings"
)

func main() {
	klog.Info(os.Args[1])
	if strings.HasSuffix(os.Args[1], ".mp4") {
		video2image.Run(os.Args[1])
	} else {
		files, err := ioutil.ReadDir(os.Args[1])
		if err != nil {
			klog.Error(err)
		}
		for _, file := range files {
			klog.Info(os.Args[1] + "/" + file.Name())
			video2image.Run(os.Args[1] + "/" + file.Name())
		}
	}

}
