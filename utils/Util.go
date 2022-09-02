package utils

import (
	"reflect"
	"unsafe"
)

/**
 * @title 工具类
 * @author xiongshao
 * @date 2022-09-02 11:12:47
 */

// string 转 字节数组
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// 字节数组 转 string
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
