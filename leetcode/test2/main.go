package main

import "fmt"

func maxSubArray(nums []int) int {
	var Max int // 最大值
	var sum int // 当前累计值

	if len(nums) < 0 {
		return 0
	}

	for k, v := range nums {
		if k == 0 {
			sum = nums[0]
			Max = sum
			continue
		}
		// 判断累计值是否小于0
		if sum+v < 0 {
			if Max < v {
				Max = v
			}
			//sum += v
		} else {
			if sum < 0 {
				sum = 0
			}
			sum += v
		}
		if Max < sum {
			Max = sum
		}
		fmt.Println(v, "\t", Max, "\t", sum)
	}
	return Max
}

func main() {

	//var nums []int = []int{-2,1,-3,4,-1,2,1,-5,4}
	var nums []int = []int{8, -19, 5, -4, 20}
	//var nums []int = []int{-2,-1}
	//var nums []int = []int{-1,-2}
	fmt.Println(maxSubArray(nums))
}
