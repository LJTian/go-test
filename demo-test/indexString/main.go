package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "TLJ_test测试商户#as##sada"

	i := strings.Split(str1, "#")
	fmt.Printf("str1 is [%s]\n", i[0])

	imageName := "registry.uniontech.com/utccp-components/docker-builder:1.2.1-rc9"
	fmt.Println(imageName[:strings.LastIndexByte(imageName, '/')] + "/" + "abc" + imageName[strings.LastIndexByte(imageName, ':'):])
}
