package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Str1 struct {
	a []string
}

var a = Str1{
	a: []string{"adsadas", "asdasdas"},
}

// 知道怎么用了，但是没有对应的冲突案例
func main() {

	//data.Store(a)
	rand.Seed(time.Now().UnixNano())
	n := 100
	for i := 0; i <= n; i++ {

		go func(i int) {
			for {
				time.Sleep(time.Duration(rand.Int() % 10))
				//fmt.Printf("%v \n", data.Load().(Str1))
				fmt.Printf("%v \n", a.a)
			}
			//data.Store(a)
		}(i)
	}

	for i := 0; i <= n; i++ {
		go func(i int) {
			for {
				time.Sleep(time.Duration(rand.Int() % 10))
				a.a = []string{fmt.Sprintf("texs%d", i),
					fmt.Sprintf("texs%d", i+1)}
			}
		}(i)
	}

	time.Sleep(10 * time.Second)
}
