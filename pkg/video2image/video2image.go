package video2image

import (
	"k8s.io/klog/v2"
	"os"
	"os/exec"
	"strings"
)

func Run(videoPath string, typ string) {
	path := strings.TrimSuffix(videoPath, "."+typ)
	err := os.Mkdir(path, 0777)
	if err != nil {
		klog.Error(err)
	}
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-r", "1", path+"/%06d.jpg")
	klog.Info(cmd.String())
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		klog.Errorf("failed to call cmd.Run(): %v", err)
	}
}
