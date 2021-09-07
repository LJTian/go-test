package main

import "fmt"

type ErrorTest struct {
	err error
}

func (e *ErrorTest) Error() string {
	return string("This is a Error")
}

func main() {
	e := ErrorTest{}

	fmt.Println(e.Error())
}
