# 控制语句

## if
```go
package main

func main(){
	a := 1
    if a == 0 {
        // a is 0
    }else if a == 1{
        // a is 1
    }else{
        // a not is 0 or 1
    }   
}

```

## switch
```go
package main

func main()  {
	a := 1
	switch {
	case a == 0:
		// a is 0
	case a == 1:
		// a is 1 
	case a==2:
		// a is 2
		fallthrough	
	default:
		// a not is 0 or 1
	}
}
```
默认情况： case 执行完成后会跳出本次判断，如果不想跳出需要使用fallthrough 关键字

## for

```go
package main

import "fmt"

func main() {
    
	//zero one two three four five six seven eight nine
	var iArr = []string{ "zero", "one", "two", "three", "four", 
		"five", "six", "seven", "eight", "nine"}
	
	for index, value := range iArr {
		fmt.Printf("number [%d] english is [%s]", index,value)
	}
}
```
index 是索引值，value 是内部值
***注意:*** value 是局部变量，修改它不会导致 iArr被修改

# 数据结构

## var

```go
package main

import "fmt"

func main() {

	// 声明类型，赋值
	var a int
	a = 1

	// 直接赋值
	var b = 2

	// 缩写
	c := 3

	// 类型推导
	d, e, f := 4, "5", '6'

	//输出结果{ 1 2 3 4 5 54 }
	fmt.Println(a,b,c,d,e,f)

}

```

## const

```go
package main

import "fmt"

func main() {
	// 单个常量
	const a = -1

	// 常量组 类似于 其它语言的枚举
	const (
		b = iota
		c
		d
		e
		f
	)

	// 输出结果{ -1 0 1 2 3 4 }
	fmt.Println(a, b, c, d, e, f)
}
```
iota 是一个自增变量

## [num]type

```go
package main

import "fmt"

func main() {

	// var 变量名 [长度]类型
	var iArr [5]int

	for i := 0; i < 5; i++ {
		iArr[i] = i
	}

	// 输出结果{ 0 1 2 3 4 }
	fmt.Println(iArr)

	// 自动推导计算长度
	var sArr = []string{"one", "two", "three", "four", "five"}

	//  输出结果{ one two three four five }
	fmt.Println(sArr)
}
```

## string

```go
package main

import "fmt"

func main() {
	// 字符串1
	str1 := "My name is "

	// 字符串2
	str2 := "string"

	// 字符串3 = 字符串1 与字符串2 的拼接
	str3 := str1 + str2

	// 输出结果{ My name is string }
	fmt.Println(str3)
}
```

## slice

```go
package main

import "fmt"

func main() {

	// 切片是对数组的引用
	// 创建类型不开辟空间，一个一个往里面加元素
	var slice1 []int
	slice1 = make([]int, 0)
	slice1 = append(slice1, 1)
	// 输出结果{ 1 }
	fmt.Println(slice1)

	// 创建类型并开辟空间，逐个位置设置元素值
	var slice2 []int
	slice2 = make([]int, 4)
	slice2[0] = 1
	slice2[1] = 2
	slice2[2] = 3
	slice2[3] = 4
	// 输出结果{ 1 2 3 4 }
	fmt.Println(slice2)

	// 两个切片进行拼接
	slice3 := append(slice1, slice2[2:3]...) // slice2[2:3] 含前不含后
	// 输出结果{ 1 3 }
	fmt.Println(slice3)
}
```

## map

```go
package main

import "fmt"
func main() {

	// 创建类型
	var mStr map[int]string

	// 开辟空间
	mStr = make(map[int]string)

	// 设置值
	mStr[1] = "one"
	mStr[2] = "two"
	mStr[3] = "three"
	mStr[4] = "four"

	// 输出结果{map[1:one 2:two 3:three 4:four]}
	fmt.Println(mStr)
}
```

## struct

```go
package main
import "fmt"
func main() {
	// 定义结构体类型
	type strInfo struct {
		typeName string
	}

	// 声明变量
	var objStrInfo strInfo

	// 结构体属性赋值
	objStrInfo.typeName = "strInfo"
	
	// 输出结果{ 结构体类型为: strInfo }
	fmt.Println("结构体类型为:", objStrInfo.typeName)
}
```

## class
```go
package main

import "fmt"
// 定义类
type strInfo struct {
	typeName string  //类的属性
}

// 通过类的方法设置属性
func (this *strInfo) SetTypeName(str1 string) {
	this.typeName = str1
}

// 通过类的方法获取属性
func (this *strInfo) GetTypeName() string {
	return this.typeName
}

func main() {

	// 声明变量
	var objStrInfo strInfo

	// 结构体属性赋值
	objStrInfo.SetTypeName("strInfo")

	// 输出结果{ 结构体类型为: strInfo }
	fmt.Println("结构体类型为:", objStrInfo.GetTypeName())
}
```

## interface
```go
package main
import "fmt"
// 接口
type objIF interface {
	SetName(string)
	Speak()
}

// 实现接口给的类1
type obj1 struct {
	myName string
}
func (this *obj1) SetName(name string) {
	this.myName = name
}
func (this *obj1) Speak() {
	fmt.Printf("My name is %s, skill is 无敌\n", this.myName)
}

// 实现接口给的类2
type obj2 struct {
	myName string
}
func (this *obj2) SetName(name string) {
	this.myName = name
}
func (this *obj2) Speak() {
	fmt.Printf("My name is %s, skill is 隐身\n", this.myName)
}

func main() {
	IF := []objIF{}
	// 开辟空间设置属性
	people1 := new(obj1)
	people1.SetName("张三")
	IF = append(IF, people1)
	// 开辟空间设置属性
	people2 := new(obj2)
	people2.SetName("李四")
	IF = append(IF, people2)
	// 让他们进行发言
	for _, v := range IF {
		v.Speak()
	}
}
```

## 管道
- 普通管道
```go
package main

import (
	"fmt"
	"time"
)
// 普通管道
func main() {
  // 创建缓存管道
  ch1 := make(chan int, 5)
  // 创建生产者
  go func() {
    for i := 0; i < 5; i++ {
        ch1 <- i
	}
    close(ch1) // 切记：需要关闭管道
  }()
  time.Sleep(100)
  // 创建消费者
  for {
    if v, ok := <-ch1; !ok {
        break
    } else {
        fmt.Println(v)
    }
  }
  fmt.Println("结束")
}
```
- 匿名管道
```go
package main
import (
  "fmt"
  "time"
)
// 匿名管道
func main() {
	// 创建没有缓存的管道:这种管道主要用于通知
	ch1 := make(chan int) // 通知消费者打印
	ch2 := make(chan int) // 通知主进程退出
	var arrt []int        // 局部容器
	// 消费者
	go func() {
		//当管道里面有了数据才会进行执行，没有数据会一直阻塞在这里
		fmt.Println("我在等通知")
		<-ch1
		fmt.Println("开始消费")
		for _, v := range arrt {
			fmt.Printf("%v\t ", v)
		}
		fmt.Println("消费结束")
		ch2 <- 1
		close(ch2)
	}()
	time.Sleep(10) // 等待
	fmt.Println("开始生产")
	// 生产者
	for i := 0; i < 5; i++ {
		arrt = append(arrt, i)
	}
	fmt.Println("生产完成")
	close(ch1)
	<-ch2 // 接受消费者通知退出
	fmt.Println("结束")
}
```
结果：  
```shell
我在等通知
开始生产
生产完成
开始消费
0        1       2       3       4       消费结束
结束
```

# 函数
## 没有接受者的普通函数
```go
package main
import "fmt"

// 函数
func func1(inPara1, inPara2 string) string {
	return fmt.Sprintf(" %s，%s", inPara1, inPara2)
}

func main() {

	// 我没有接受者，所以我是普通函数
	fmt.Println(func1("我没有接受者", "所以我是普通函数"))
}
```
## 匿名函数
- 匿名函数
```go
package main
import "fmt"
func main() {

	// 没有设置具体的函数名称
	func1 := func(inPara1, inPara2 string) string {
		return fmt.Sprintf(" %s，%s", inPara1, inPara2)
	}

	// 我的函数名为变量名，没有具体的名字，所以我是匿名函数
	fmt.Println(func1("我的函数名为变量名，没有具体的名字", "所以我是匿名函数"))
}
```
- 闭包  
  匿名函数内部可以直接使用外部的变量，并可对其进行修改。称之为闭包
```go
package main
import "fmt"

func main() {
	a := 0
	for i := 0; i < 5; i++ {
		// 匿名函数内部可以直接使用外部的变量，并可对其进行修改。称之为闭包
		func() {
			a++
		}()
		// 输出 a is 1  a is 2  a is 3  a is 4  a is 5
		fmt.Printf("a is %d \t", a)
	}
}
```
## 有接收者的函数(方法)
```go
package main
import "fmt"

// 自定义类型
type myStr string

// 方法
func (this *myStr) mod1() {
	fmt.Println("我有接受者，所以我是方法，不是普通函数")
}

func main() {

	// 初始化接收者
	var objMyStr myStr
	// 通过接收者调用方法
	objMyStr.mod1()
	// 输出为{我有接受者，所以我是方法，不是普通函数}
}

```

# 内置函数

