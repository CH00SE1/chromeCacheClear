package Mb

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

/**
 * @title 统计文件夹下面每个文件的大小
 * @author xiongshao
 * @date 2022-09-02 09:40:28
 */

var waitGroup sync.WaitGroup
var ch = make(chan struct{}, 255)

func dirents(path string) ([]os.FileInfo, bool) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}
	return entries, true
}

// 递归计算目录下所有文件
func walkDir(path string, fileSize chan<- int64) {
	defer waitGroup.Done()
	ch <- struct{}{} //限制并发量
	entries, ok := dirents(path)
	<-ch
	if !ok {
		log.Fatal("can not find this dir path!!")
		return
	}
	for _, e := range entries {
		if e.IsDir() {
			waitGroup.Add(1)
			go walkDir(filepath.Join(path, e.Name()), fileSize)
		} else {
			fileSize <- e.Size()
		}
	}
}

func all_file(dir_path string) {

	//文件大小chennel
	fileSize := make(chan int64)
	//文件总大小
	var sizeCount int64
	//文件数目
	var fileCount int

	//计算目录下所有文件占的大小总和
	waitGroup.Add(1)
	go walkDir(dir_path, fileSize)

	go func() {
		defer close(fileSize)
		waitGroup.Wait()
	}()

	for size := range fileSize {
		fileCount++
		sizeCount += size
	}

	fmt.Printf("-该目录大小为 %.1fM-文件总数为 %d\n", float64(sizeCount)/1024/1024, fileCount)
}

func DirSizeMain(dir_path string) {
	t := time.Now()
	files, err := ioutil.ReadDir(dir_path)
	if err != nil {
		log.Fatal(err)
	} else {
		for _, fi := range files {
			if fi.IsDir() {
				fmt.Printf("%s", dir_path+"\\"+fi.Name())
				all_file(dir_path + "/" + fi.Name())
			} else {
				fmt.Printf("%s 大小为 %.1fM \n", fi.Name(), float64(fi.Size())/1024/1024)
			}
		}
	}
	fmt.Println("花费的时间为 " + time.Since(t).String())
}
