package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("此程序是Hyper-v的开启关闭的一种快捷方式,需要使用系统管理员权限执行")
	fmt.Println("1:开启 2：关闭")
	var key int
	for true {
		fmt.Scanf("%d\n", &key)
		if key == 1 {
			cmd := exec.Command("bcdedit", "/set", "hypervisorlaunchtype", "auto")
			buf, _ := cmd.Output()
			var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(buf)
			fmt.Println(string(decodeBytes))
			break
		} else if key == 2 {
			cmd := exec.Command("bcdedit", "/set", "hypervisorlaunchtype", "off")
			buf, _ := cmd.Output()
			var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(buf)
			fmt.Println(string(decodeBytes))
			break
		} else {
			fmt.Println("输入错误")
		}
	}
	fmt.Println("操作完成，3秒后退出")
	time.Sleep(3 * time.Second)
}
