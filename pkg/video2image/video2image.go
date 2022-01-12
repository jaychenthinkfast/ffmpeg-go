package video2image

import (
	"github.com/jaychenthinkfast/ffmpeg-go/pkg/util/workqueue"
	"k8s.io/klog/v2"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type Item struct {
	FilePath string
}

var typ, fr string
var concurrentNum int
var q *workqueue.Type
var consumerWG sync.WaitGroup

func Init(t string, num int, r string) {
	typ = t
	fr = r
	concurrentNum = num
	if concurrentNum > 0 {
		q = workqueue.NewQueue()
		consumerWG = sync.WaitGroup{}
		consumerWG.Add(concurrentNum)
		for i := 0; i < concurrentNum; i++ {
			go func(i int) {
				defer consumerWG.Done()
				for {
					item, quit := q.Get()
					if quit {
						return
					}
					v := item.(*Item)
					klog.Infof("Worker %v: begin processing %v", i, v.FilePath)
					v.Run()
					klog.Infof("Worker %v: done processing %v", i, v.FilePath)
					q.Done(item)
				}
			}(i)
		}
	}
}

func Add(path string) {
	v := &Item{
		FilePath: path,
	}
	if concurrentNum == 0 {
		v.Run()
	} else {
		q.Add(v)
	}
}

func End() {
	if concurrentNum > 0 {
		q.ShutDownWithDrain()
		consumerWG.Wait()
	}
}

func (item *Item) Run() {
	if strings.HasSuffix(item.FilePath, typ) {
		path := strings.TrimSuffix(item.FilePath, "."+typ)
		err := os.Mkdir(path, 0777)
		if err != nil && !strings.Contains(err.Error(), "file exists") {
			klog.Error(err)
		}
		cmd := exec.Command("ffmpeg", "-i", item.FilePath, "-r", fr, path+"/%06d.jpg")
		klog.Info(cmd.String())
		//cmd.Stdout = os.Stdout
		//cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			klog.Errorf("failed to call cmd.Run(): %v", err)
		}
	}
}
