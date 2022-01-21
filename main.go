package main

import (
	"fmt"
	"io"
	"os"
)

func scpfile(src, dst string) error {
	//打开原文件
	srcfile, err := os.Open(src)
	if err != nil{
		return err
	}
	//延迟关闭原文件
	defer srcfile.Close()

	//创建目标文件
	dstfile, err := os.Create(dst)
	if err != nil{
		return err
	}
	//延迟关闭目标文件
	defer dstfile.Close()

	//拷贝原文件至目标文件中
	//定义tmpdata用于存储临时读取的文件内容
	tmpdata := make([]byte, 1024)
	for {
		n, err := srcfile.Read(tmpdata)
		if err != nil{
			if err == io.EOF{
				break
			}
			return err
		}

		dstfile.Write(tmpdata[:n])
	}
	return nil
}

func main() {
	fmt.Println(scpfile("file01.txt", "file02.txt"))
}