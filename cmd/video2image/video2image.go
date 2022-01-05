package main

import (
	"github.com/jaychenthinkfast/ffmpeg-go/cmd/video2image/config"
	"github.com/jaychenthinkfast/ffmpeg-go/pkg/video2image"
	"io/ioutil"
	"k8s.io/klog/v2"
	"strings"
)

func main() {
	c, err := config.Parse()
	if err != nil {
		klog.Error(err)
	}
	c.Path.Dirs = append(c.Path.Dirs, c.Path.Dir)
	c.Path.Files = append(c.Path.Files, c.Path.File)
	klog.Info(c.Path)
	for _, dir := range c.Path.Dirs {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			klog.Error(err)
		}
		for _, file := range files {
			dir = strings.TrimSuffix(dir, "/")
			klog.Info(dir + "/" + file.Name())
			c.Path.Files = append(c.Path.Files, dir+"/"+file.Name())
		}
	}
	for _, filePath := range c.Path.Files {
		video2image.Run(filePath, c.Type)
	}
}
