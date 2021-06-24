package main

import (
	"SAMDel/settings"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)


func Evens(out chan string) {
	defer close(out)
	for _,vc := range settings.Conf.Confing{
		// 判断路径是否存在
		if _, err0 := os.Stat(vc.Path); err0 != nil {
			continue
		}
		// 判断是否以根结尾
		if ! strings.HasSuffix(vc.Path, "\\")  {
			vc.Path = vc.Path + "\\"
		}
		// 设置默认保存时间
		if vc.KeepTime == 0 {
			vc.KeepTime = 200
		}
		vc.KeepTime = vc.KeepTime * 86400

		DirInfoList, err := ioutil.ReadDir(vc.Path)
		if err != nil {
			log.Fatal(err)
		}
		for _,value := range DirInfoList {
			// 查找超过保留时间的文件
			dropObjName := path.Join(vc.Path + value.Name())
			fileInfo, _ := os.Stat(dropObjName)
			diffTime := time.Now().Unix() - fileInfo.ModTime().Unix()
			if diffTime > vc.KeepTime {
				out <- dropObjName
			}
		}
	}
}

// DropFile 多线程删除
func DropFile(w int, jobs <-chan string,wg *sync.WaitGroup)  {
	for file:=range jobs {
		start := time.Now().Unix()
		if err := os.RemoveAll(file); err != nil {
			log.Println(err)
		}
		end := time.Now().Unix()
		log.Printf("%d删除%s完成，耗时%d秒！\n",w,file, end-start)
	}
	wg.Done()
}



func main() {
	start := time.Now().Unix()
	// 初始化配置文件
	if err := settings.Init(); err != nil {
		log.Println("配置文件错误, 请检查配置文件!")
	}
	jobs := make(chan string)

	wg := sync.WaitGroup{}
	// 开通线程池，执行消费jobs管道
	if settings.Conf.ThreadPools > 1 {
		settings.Conf.ThreadPools = 1
	}
	for w := 1; w <= settings.Conf.ThreadPools; w++ {
		wg.Add(1)
		go DropFile(w, jobs, &wg)
	}
	// 给jobs管道添加内容
	Evens(jobs)


	wg.Wait()

	end := time.Now().Unix()
	log.Printf("lizuqing 总耗时：%d 秒！", end-start)
	time.Sleep(time.Second * 5)


}
