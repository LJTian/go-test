/*
	原题：
	var m map[int]int
	fmt.Println(m[1])

	a := []int{1,2,3,4,5,6,7}
	println(a[7:])

	代码的问题是什么？
	for _, item := range arr {
		slic = append(slic, &item)
	}
*/
package main

import "fmt"

func test(v interface{}) {
	fmt.Println(v == nil)
}

func main() {

	var m map[int]int
	//m[1] = 1 // 直接set painc
	fmt.Println(&m) // get 为 0

	a := []int{1, 2, 3, 4}
	fmt.Println(cap(a))
	//fmt.Println(a[1:3])

	//fmt.Println(a[4])  // 越界
	fmt.Println(a[4:]) // 结果是空切片

	var b *string
	fmt.Println(b == nil)
	test(a)
	test(nil)

	// 代码的问题是什么？ - 这个题的Get 点很重要。
	// 换一种问法，slic最终的值是什么？(也就能明白了) 局部变量 只保留了最终的值也就是最后的值
	slic := []*int{}
	arr := []int{1, 2, 3}
	for _, item := range arr {
		slic = append(slic, &item)
	}
	fmt.Println(slic)
	//
	//for _, k := range slic {
	//	fmt.Println(*k)
	//}
}
