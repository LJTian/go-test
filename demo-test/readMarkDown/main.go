package main

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
	"io/ioutil"
	"strings"
)

func main() {
	// 从文件中读取 Markdown 内容
	data, err := ioutil.ReadFile("/home/ljtian/file/git/openshift/README.md")
	if err != nil {
		panic(err)
	}

	// 将 Markdown 转换为 HTML
	html := blackfriday.Run(data)
	a := string(html)

	b := strings.Split(strings.Split(a, "<h1>")[1], "</h1>")
	// 输出 HTML
	fmt.Println(b[0])
}
