package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "汉字123"
	str2 := "zimu"
	str3 := "汉字zimu"

	for k, v := range strings.Split(str1, "") {
		fmt.Printf("k:%v\t v:%s\n", k, v)
	}
	for k, v := range strings.Split(str2, "") {
		fmt.Printf("k:%v\t v:%s\n", k, v)
	}
	for k, v := range strings.Split(str3, "") {
		fmt.Printf("k:%v\t v:%s\n", k, v)
	}
}
