package main

import "fmt"

type node struct {
	data byte
	next *node
}

/*
	4. 给定一单链表的头指针，请使用一趟扫描完成链表反转。
	初始链表：A -> B -> C -> D 预期结果：D -> C -> B -> A
*/
func list() *node {

	var tail, l1 *node
	n := 4

	for i := 0; i < n; i++ {
		if l1 == nil {
			l1 = &node{byte('A' + i), nil}
			tail = l1
		} else {
			l1.next = &node{byte('A' + i), nil}
			l1 = l1.next
		}
	}

	return tail

}

func func1(tail *node) *node {

	if tail == nil || tail.next == nil {
		return tail
	}

	newNode := func1(tail.next)

	tail.next.next = tail
	tail.next = nil

	return newNode

}

func main() {

	tail := list()
	//for tail != nil {
	//	fmt.Printf("%s\t", string(tail.data))
	//	tail = tail.next
	//}
	ret := func1(tail)
	for ret != nil {
		fmt.Printf("%s\t", string(ret.data))
		ret = ret.next
	}
}
