package settings

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

var Conf = new(AppConfig)

type AppConfig struct {
	// 必须大写开头
	ThreadPools int        `yaml:"ThreadPools"`
	Confing     []DropInfo `yaml:"confing"`
}

type DropInfo struct {
	Path     string `yaml:"path"`
	KeepTime int64  `yaml:"keepTime"`
}

func Init() (err error) {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fileCPath := path.Join(dir, "conf.yaml")
	data, err := ioutil.ReadFile(fileCPath)
	fmt.Println(fileCPath)
	err = yaml.Unmarshal(data, Conf)
	if err != nil {
		log.Printf("conf.yaml 配置文件出错! %v \n", err)
	}
	return
}
