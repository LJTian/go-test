package main

func main() {
	deferFun1()
	//deferFun2()
}

func deferFun1() {

	// 错误写法 原因：https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-panic-recover/ 评论区
	defer recover()
	panic(1)
}

func deferFun2() {
	defer func() {
		recover()
	}()
	panic(1)
}
