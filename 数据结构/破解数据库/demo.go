package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"os"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

func checkPassWord(password string) bool  {
	_, err := sql.Open("mysql", "root:"+password+"@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		return false
	} else {
		return true
	}
}

func main() {
	path := ""
	resultPath := ""
	reader, err := os.Open(path)
	createFile, err := os.Create(resultPath)
	defer createFile.Close()
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	if err != nil {
		panic(err)
	}
	newReader := bufio.NewReader(reader)
	writer := bufio.NewWriter(createFile)
	for  {
		line, _, err := newReader.ReadLine()
		if err == io.EOF {
			break
		}
		lineStr := string(line)
		split := strings.Split(lineStr, ",")
		if len(split) == 2 {
			password := split[0]
			if checkPassWord(password) {
				fmt.Println("密码成功" + password)
			} else {
				fmt.Println("密码失败")
			}
			fmt.Fprintln(writer, password)
		}
	}
	writer.Flush()
}
