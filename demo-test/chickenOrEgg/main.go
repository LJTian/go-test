package main

import (
	"fmt"
	"sort"
)

func main() {
	chinaArr := onChina()
	sort.Strings(chinaArr)
	USAArr := onUSA()
	sort.Strings(USAArr)
	myPrintf("中国", chinaArr)
	myPrintf("美国/英国", USAArr)
}

func onChina() []string {
	return []string{
		"鸡",
		"蛋",
	}
}

func onUSA() []string {
	return []string{
		"chicken",
		"eggs",
	}
}

func myPrintf(area string, arr []string) {
	fmt.Printf("在[%s], 先有[%s],后有[%s]\n", area, arr[0], arr[1])
}
