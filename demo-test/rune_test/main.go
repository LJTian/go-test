package main

import "fmt"

func main() {

	str1 := "汉字123"

	fmt.Println(str1[:4])
	fmt.Println(string([]rune(str1)[:3])) // 使用 []rune 可以保证汉字不会被切错

}
