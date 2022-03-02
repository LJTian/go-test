package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "TLJ_test测试商户#as##sada"

	i := strings.Split(str1, "#")
	fmt.Printf("str1 is [%s]", i[0])
}
