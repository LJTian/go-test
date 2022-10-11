package main

import "fmt"

//func main() {
//    in := [3]string{"a", "b", "c"}
//    var out []*string
//    for _, v := range in {
//        输入：
//        out = append(out, &v) // v 是一个局部变量，循环完成之后存储的是"c"，这个操作只是把V的地址保存了，没有保存地址的值
//    }
//    fmt.Println("Values:", *out[0], *out[1], *out[2])
//}

func main() {
	in := [5]string{"a", "b", "c", "d", "e"}
	var out []*string
	for _, v := range in {
		// 输入： 可修改的位置
		var arrStrP []string // 中间加一层存储值的空间
		arrStrP = append(arrStrP, v)
		out = append(out, &arrStrP[len(arrStrP)-1])
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
}
