package main

import (
	"go_test/pkgInPkg/pkg1"
	"go_test/pkgInPkg/pkg1/pkg2"
)

func main() {
	pkg1.PrintfLog()
	pkg2.PrintfLog1()
}
