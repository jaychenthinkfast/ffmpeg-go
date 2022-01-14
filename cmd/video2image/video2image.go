package main

import (
	"github.com/jaychenthinkfast/ffmpeg-go/cmd/video2image/config"
	"github.com/jaychenthinkfast/ffmpeg-go/pkg/video2image"
	"io/ioutil"
	"k8s.io/klog/v2"
	"strings"
)

func readDir(dir string) (fileSlice []string) {
	dir = strings.TrimSuffix(dir, "/") + "/"
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		isDir := file.IsDir()
		if isDir {
			_fileSlice := readDir(dir + file.Name() + "/")
			fileSlice = append(fileSlice, _fileSlice...)
		} else {
			fileSlice = append(fileSlice, dir+file.Name())
		}
	}
	return fileSlice
}

func main() {
	c, err := config.Parse()
	if err != nil {
		klog.Error(err)
	}
	if c.Path.Dir != "" {
		c.Path.Dirs = append(c.Path.Dirs, c.Path.Dir)
	}
	c.Path.Files = append(c.Path.Files, c.Path.File)
	klog.Info(c.Path)
	for _, dir := range c.Path.Dirs {
		c.Path.Files = append(c.Path.Files, readDir(dir)...)
	}
	klog.Info(c.Path.Files)
	video2image.Init(c.Type, c.ConcurrentNum, c.FrameRate)
	for _, filePath := range c.Path.Files {
		video2image.Add(filePath)
	}
	video2image.End()
}
