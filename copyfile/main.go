package main

import (
	"fmt"
	"io"
	"os"
)

//define function for scpfile, src is source files, dst is destination file
func ScpFile(src, dst string) error {
	//open source files
	srcfile, err := os.Open(src)
	if err != nil{
		return err
	}
	//delay closing source files
	defer srcfile.Close()

	//create destination file
	dstfile, err := os.Create(dst)
	if err != nil{
		return err
	}
	//delay closing destination file
	defer dstfile.Close()

	//Copy source files to destination file
	//Define 'tmpdata' to store temporarily read file contents
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
