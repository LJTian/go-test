package main

import "fmt"

/*
	数组交集
	给定两个数组，编写一个函数来计算它们的交集。

	示例 1:

	输入: nums1 = [1,2,2,1], nums2 = [2,2]

	输出: [2]
*/

type node struct {
	data int
	next *node
}

/* 写个链表 */
func list() {

	var tail, l1 *node
	n := 10

	for i := 0; i < n; i++ {
		if l1 == nil {
			l1 = &node{i, nil}
			tail = l1
		} else {
			l1.next = &node{i, nil}
			l1 = l1.next
		}
	}

	for tail != nil {
		fmt.Println(tail.data)
		tail = tail.next
	}
}

func test1(nums1, nums2 []int) (ret []int) {

	m1 := make(map[int]bool)

	for _, v := range nums1 {
		m1[v] = false
	}

	for _, v := range nums2 {
		if _, ok := m1[v]; ok {
			m1[v] = true
		}
	}

	for k, v := range m1 {
		if v {
			ret = append(ret, k)
		}
	}

	return
}

func main() {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:] // "b","c"
	str2[1] = "new"  // "b","new"

	fmt.Println(str1)
	str2 = append(str2, "z", "x", "y")
	fmt.Println(str1)

}
