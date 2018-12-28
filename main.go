package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

//源文件
const sourcePath = "/Volumes/PersonalFiles/Home2/Downloads/video/1/"

//目标文件
const targetFile = "/Volumes/PersonalFiles/Home2/Downloads/target1.ts"

func main() {
	files, _ := ioutil.ReadDir(sourcePath)
	fpTarget, err := os.OpenFile(targetFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		fpSource, _ := os.OpenFile(path.Join(sourcePath, file.Name()), os.O_RDONLY, 0644)
		b, _ := ioutil.ReadAll(fpSource)
		fpTarget.Write(b)
		fpSource.Close()
		fmt.Printf("%s copied\n", file.Name())
	}
	fpTarget.Close()
}
