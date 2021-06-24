package settings

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)


var Conf = new(AppConfig)

type AppConfig struct {
	// 必须大写开头
	ThreadPools    int `yaml:"ThreadPools"`
	Confing   []DropInfo	`yaml:"confing"`
}

type DropInfo struct {
	Path    string	`yaml:"path"`
	KeepTime int64 `yaml:"keepTime"`
}


func Init()(err error){
	data,err := ioutil.ReadFile("conf.yaml")
	err = yaml.Unmarshal(data, Conf)
	if err != nil {
		log.Printf("conf.yaml 配置文件出错! %v \n", err)
	}
	return
}


