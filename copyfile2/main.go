package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//define function for copy file of buffer
func Scp(src, dst string) error{
	//打开文件
	srcfile, err := os.Open(src)
	if err != nil{
		return err
	}

	defer srcfile.Close()
	// 创建读文件对象
	reader := bufio.NewReader(srcfile)

	//创建目标文件
	dstfile, err := os.Create(dst)
	if err != nil{
		return err
	}

	//创建写文件对象 
	write := bufio.NewWriter(dstfile)
	defer write.Flush()

	//写文件
	line := make([]byte, 1024*1024)
	for {
		n, err := reader.Read(line)
		if err != nil{
			if err == io.EOF{
				break
			}
			return err
		}

		write.Write(line[:n])
		
	}

	return nil	
}


func main(){

	if err := Scp("test", "test.log"); err != nil{
		fmt.Println("copy file faild. ", err)
		return
	}
	fmt.Println("copy file sucessful.")
	
}