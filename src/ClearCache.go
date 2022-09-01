package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/**
 * @title 清除浏览器缓存
 * @author xiongshao
 * @date 2022-09-01 10:07:38
 */

// 浏览器缓存路径
const path = "/AppData/Local/Google/Chrome/User Data/Default"

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// 不删除空文件夹,只删除文件
func removeAllFiles(path string) {
	// 检查路径是否存在
	_, err := os.Stat(path)
	if err != nil {
		log.Panic(path + " 路径不存在")
	}
	// 如果该路径是一个文件，直接删除
	if IsFile(path) {
		os.Remove(path)
	}
	// 路径是一个文件夹的路径
	dir, err := os.ReadDir(path)
	if err != nil {
		log.Panic("path not found --> " + path)
	}
	for _, entry := range dir {
		tmpPath := path + entry.Name()
		if IsFile(tmpPath) {
			fmt.Println("删除文件:" + tmpPath)
			os.Remove(tmpPath)
		} else {
			removeAllFiles(tmpPath)
		}
	}
}

func clear(profile_path string) {
	// 检查路径是否存在
	_, err := os.Stat(profile_path)
	if err != nil {
		log.Panic(profile_path + " 路径不存在")
	}
	// 清除缓存的文件和图片 这个文件夹里的数据相对较大
	cacheDataFileAndPhoto := profile_path + "/Cache/Cache_Data/"
	removeAllFiles(cacheDataFileAndPhoto)
}

// 获取当前执行下的路径
func GetCurrentDirectory() string {
	//返回绝对路径 filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	prefixPath := "C:/Users/"
	//将\替换成/
	s := strings.Replace(dir, "\\", "/", -1)
	s1 := strings.Replace(s, prefixPath, "", -1)
	index := strings.Index(s1, "/")
	return prefixPath + s1[:index]
}

// 启动程序
func main() {
	fmt.Println("  ___  _  _  ____   __   _  _  ____  ___   __    ___  _  _  ____ \n / __)/ )( \\(  _ \\ /  \\ ( \\/ )(  __)/ __) / _\\  / __)/ )( \\(  __)\n( (__ ) __ ( )   /(  O )/ \\/ \\ ) _)( (__ /    \\( (__ ) __ ( ) _) \n \\___)\\_)(_/(__\\_) \\__/ \\_)(_/(____)\\___)\\_/\\_/ \\___)\\_)(_/(____)")
	fmt.Println("version: 1.0.0")
	fmt.Println("-----------------------------------------------")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		directory := GetCurrentDirectory()
		targetPath := directory + path
		fmt.Printf("目标路径:%s \n\n", targetPath)
		clear(targetPath)
		time.Sleep(1 * time.Second)
		cancel()
	}()
	<-ctx.Done()
	fmt.Println("按任意键继续...")
	var input string
	fmt.Scanln(&input)
}
