package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"k8s.io/klog/v2"
)

type Config struct {
	ConcurrentNum int
	Type          string
	Path          struct {
		File  string
		Dir   string
		Files []string
		Dirs  []string
	}
}

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "demo/demo.yaml", "config path")
}

func Parse() (Config, error) {
	flag.Parse()
	klog.Info(configPath)
	config, err := ioutil.ReadFile(configPath)
	if err != nil {
		klog.Error(err)
		return Config{}, err
	}
	c := Config{}
	err = yaml.Unmarshal([]byte(config), &c)
	if err != nil {
		klog.Error(err)
		return Config{}, err
	}
	return c, err
}
