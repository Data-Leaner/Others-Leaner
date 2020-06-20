package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"数据接口/02栈/StackArray"
)

func GetPath(path string, files []string)([]string, error) {
	Reader, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("文件夹不可读取")
	}

	for _, filename := range Reader {
		if filename.IsDir() {
			// 构造新的路径
			fullDir := path + "\\" + filename.Name()
			files = append(files, fullDir)
			// 该部分是赋值，不能使用类型推断
			files, err = GetPath(fullDir, files)
			if err != nil {
				fmt.Println("错误信息")
			}
		} else {
			// 构造新的路径
			fullDir := path + "\\" + filename.Name()
			files = append(files, fullDir)
		}
	}
	return files, nil
}

func main() {
	var files []string
	var err error
	files, err = GetPath("D:\\beego", files)
	if err != nil {
		fmt.Println("错误信息： %V", err)
	}
	//for _, file := range files  {
	//	fmt.Println(file)
	//}

	var tmpFiles []string
	path := "D:\\beego"
	stack := StackArray.NewStack()
	stack.Push(path)
	for ! stack.IsEmpty() {
		path := stack.Pop().(string)
		tmpFiles = append(tmpFiles, path)
		reader, err := ioutil.ReadDir(path)
		if err != nil {
			break
		}
		for _, filename := range reader {
			if filename.IsDir() {
				stack.Push(path + "\\" + filename.Name())
				tmpFiles = append(tmpFiles, path + "\\" +  filename.Name())
			} else {
				tmpFiles = append(tmpFiles, path + "\\" +  filename.Name())
			}
		}
	}
	for _, filename := range tmpFiles {
		fmt.Println(filename)
	}
}
