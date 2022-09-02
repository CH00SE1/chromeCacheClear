package main

import (
	"ClearChromeCache/Mb"
)

/**
 * @title 主程序
 * @author xiongshao
 * @date 2022-09-01 10:07:38
 */

// 启动程序
func main() {
	pathTmp := "C:\\Users\\Administrator\\Desktop\\"
	Mb.DirSizeMain(pathTmp)
	//ClearCache.FilePathDelete(pathTmp)
}
