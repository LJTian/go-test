package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}

	if n > 0 {
		nums1 = append(nums1, nums2...)
	}
	return
}

func main() {

	a := make([]int, 0)
	b := make([]int, 0)

	a = append(a, 1, 2, 3, 0, 0, 0)
	b = append(b, 2, 5, 6)

	for _, v := range a {
		fmt.Println(v)
	}
	//fmt.Println(b[1])

	fmt.Println(a)
	fmt.Println(b)
	merge(a, 3, b, 3)

	fmt.Println(a)
}
