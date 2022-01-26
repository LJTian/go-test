package main

import (
	"fmt"
	"time"
)

func main() {

	sTime, err := time.Parse("20060102", "20220124")
	if err != nil {
		return
	}

	sTime = sTime.AddDate(0, 0, -1)

	fmt.Println(sTime.Format("20060102"))
}
