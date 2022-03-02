package main

import "fmt"

func main() {

	var ch chan interface{}

	ch = make(chan interface{})

	go func(ch chan interface{}) {
		fmt.Printf("ch is [%s]", <-ch)
		return
	}(ch)
	fmt.Printf("len(ch) is [%d] ", len(ch))
	ch <- "123"

}
