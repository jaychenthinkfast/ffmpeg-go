package video2image

import (
	"k8s.io/klog/v2"
	"os"
	"os/exec"
	"strings"
)

type Item struct {
	FilePath string
	Type     string
}

func New(path, typ string) *Item {
	return &Item{
		FilePath: path,
		Type:     typ,
	}
}

func (item *Item) Run() {
	path := strings.TrimSuffix(item.FilePath, "."+item.Type)
	err := os.Mkdir(path, 0777)
	if err != nil {
		klog.Error(err)
	}
	cmd := exec.Command("ffmpeg", "-i", item.FilePath, "-r", "1", path+"/%06d.jpg")
	klog.Info(cmd.String())
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		klog.Errorf("failed to call cmd.Run(): %v", err)
	}
}
