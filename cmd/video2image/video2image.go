package main

import (
	"github.com/jaychenthinkfast/ffmpeg-go/pkg/video2image"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"k8s.io/klog/v2"
	"strings"
)

type Config struct {
	Type string
	Path struct {
		File  string
		Dir   string
		Files []string
		Dirs  []string
	}
}

func main() {
	config, err := ioutil.ReadFile("demo/demo.yaml")
	if err != nil {
		klog.Error(err)
	}
	c := Config{}
	err = yaml.Unmarshal([]byte(config), &c)
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
