package main

import (
	"encoding/json"
	"fmt"
	"time"
)

//func main() {
//
//	var a int
//	a = 1
//
//	var b = 2
//
//	c := 3
//
//	d, e, f := 4, "5", '6'
//
//	fmt.Println(a, b, c, d, e, f)
//
//}

/*--------------------------------------------------------------------------------------------------------------------*/
//func main() {
//
//	// 单个常量
//	const a = -1
//
//	// 常量组 类似于 其它语言的枚举
//	const (
//		b = iota
//		c
//		d
//		e
//		f
//	)
//
//	// 输出结果{ -1 0 1 2 3 4 }
//	fmt.Println(a, b, c, d, e, f)
//}

/*--------------------------------------------------------------------------------------------------------------------*/
//func main() {
//
//	//var 变量名 [长度]类型
//	var iArr [5]int
//
//	for i := 0; i < 5; i++ {
//		iArr[i] = i
//	}
//
//	// 输出结果{ 0 1 2 3 4 }
//	fmt.Println(iArr)
//
//	// 自动推导计算长度
//	var sArr = []string{"one", "two", "three", "four", "five"}
//
//	//  输出结果{ one two three four five }
//	fmt.Println(sArr)
//}

/*--------------------------------------------------------------------------------------------------------------------*/
//func main() {
//
//	// 字符串1
//	str1 := "My name is "
//
//	// 字符串2
//	str2 := "string"
//
//	// 字符串3 = 字符串1 与字符串2 的拼接
//	str3 := str1 + str2
//
//	// 输出结果{ My name is string }
//	fmt.Println(str3)
//}

/*--------------------------------------------------------------------------------------------------------------------*/
//func main() {
//
//	// 创建类型
//	var mStr map[int]string
//
//	// 开辟空间
//	mStr = make(map[int]string)
//
//	// 设置值
//	mStr[1] = "one"
//	mStr[2] = "two"
//	mStr[3] = "three"
//	mStr[4] = "four"
//
//	// 输出结果{map[1:one 2:two 3:three 4:four]}
//	fmt.Println(mStr)
//}

/*--------------------------------------------------------------------------------------------------------------------*/
//func main() {
//
//	// 定义结构体类型
//	type strInfo struct {
//		typeName string
//	}
//
//	// 声明变量
//	var objStrInfo strInfo
//
//	// 结构体属性赋值
//	objStrInfo.typeName = "strInfo"
//
//	fmt.Println("结构体类型为:", objStrInfo.typeName)
//}

/*--------------------------------------------------------------------------------------------------------------------*/
//
//// 定义类
//type strInfo struct {
//	typeName string // 属性
//}
//
//// 通过类的方法设置属性
//func (this *strInfo) SetTypeName(str1 string) {
//	this.typeName = str1
//}
//
//// 通过类的方法获取属性
//func (this *strInfo) GetTypeName() string {
//	return this.typeName
//}
//
//func main() {
//
//	// 声明变量
//	var objStrInfo strInfo
//
//	// 结构体属性赋值
//	objStrInfo.SetTypeName("strInfo")
//
//	// 输出结果{ 结构体类型为: strInfo }
//	fmt.Println("结构体类型为:", objStrInfo.GetTypeName())
//}

/*--------------------------------------------------------------------------------------------------------------------*/

//// 接口
//type objIF interface {
//	SetName(string)
//	GetName() string
//	Speak()
//}
//
//// 实现接口给的类1
//type obj1 struct {
//	myName string
//}
//
//func (this *obj1) SetName(name string) {
//	this.myName = name
//}
//
//func (this *obj1) GetName() string {
//	return this.myName
//}
//
//func (this *obj1) Speak() {
//	fmt.Printf("My name is %s, skill is 无敌\n", this.myName)
//}
//
//// 实现接口给的类2
//type obj2 struct {
//	myName string
//}
//
//func (this *obj2) SetName(name string) {
//	this.myName = name
//}
//
//func (this *obj2) GetName() string {
//	return this.myName
//}
//
//func (this *obj2) Speak() {
//	fmt.Printf("My name is %s, skill is 隐身\n", this.myName)
//}
//
//func main() {
//	IF := []objIF{}
//
//	// 开辟空间设置属性
//	people1 := new(obj1)
//	people1.SetName("张三")
//	IF = append(IF, people1)
//
//	// 开辟空间设置属性
//	people2 := new(obj2)
//	people2.SetName("李四")
//	IF = append(IF, people2)
//
//	// 让他们进行发言
//	for _, v := range IF {
//		v.Speak()
//	}
//}

/*--------------------------------------------------------------------------------------------------------------------*/
//func func1(inPara1, inPara2 string) string {
//	return fmt.Sprintf(" %s，%s", inPara1, inPara2)
//}
//
//func main() {
//
//	// 我没有接受者，所以我是普通函数
//	fmt.Println(func1("我没有接受者", "所以我是普通函数"))
//}
/*--------------------------------------------------------------------------------------------------------------------*/

//func main() {
//
//	// 没有设置具体的函数名称
//	func1 := func(inPara1, inPara2 string) string {
//		return fmt.Sprintf(" %s，%s", inPara1, inPara2)
//	}
//
//	// 我的函数名为变量名，没有具体的名字，所以我是匿名函数
//	fmt.Println(func1("我的函数名为变量名，没有具体的名字", "所以我是匿名函数"))
//}
/*--------------------------------------------------------------------------------------------------------------------*/
//
//func main() {
//
//	a := 0
//	for i := 0; i < 5; i++ {
//		// 匿名函数内部可以直接使用外部的变量，并可对其进行修改。称之为闭包
//		func() {
//			a++
//		}()
//		// 输出 a is 1  a is 2  a is 3  a is 4  a is 5
//		fmt.Printf("a is %d \t", a)
//	}
//}
/*--------------------------------------------------------------------------------------------------------------------*/
//// 自定义类型
//type myStr string
//
//// 方法
//func (this *myStr) mod1() {
//	fmt.Println("我有接受者，所以我是方法，不是普通函数")
//}
//
//func main() {
//
//	// 初始化接收者
//	var objMyStr myStr
//	// 通过接收者调用方法
//	objMyStr.mod1()
//	// 输出为{我有接受者，所以我是方法，不是普通函数}
//}
/*-------package main

//func main() {
//
//	var a int
//	a = 1
//
//	var b = 2
//
//	c := 3
//
//	d, e, f := 4, "5", '6'
//
//	fmt.Println(a, b, c, d, e, f)
//
//}

/*--------------------------------------------------------------------------------------------------------------------*/
//func main() {
//
//	// 单个常量
//	const a = -1
//
//	// 常量组 类似于 其它语言的枚举
//	const (
//		b = iota
//		c
//		d
//		e
//		f
//	)
//
//	// 输出结果{ -1 0 1 2 3 4 }
//	fmt.Println(a, b, c, d, e, f)
//}

/*--------------------------------------------------------------------------------------------------------------------*/
//func main() {
//
//	//var 变量名 [长度]类型
//	var iArr [5]int
//
//	for i := 0; i < 5; i++ {
//		iArr[i] = i
//	}
//
//	// 输出结果{ 0 1 2 3 4 }
//	fmt.Println(iArr)
//
//	// 自动推导计算长度
//	var sArr = []string{"one", "two", "three", "four", "five"}
//
//	//  输出结果{ one two three four five }
//	fmt.Println(sArr)
//}

/*--------------------------------------------------------------------------------------------------------------------*/
//func main() {
//
//	// 字符串1
//	str1 := "My name is "
//
//	// 字符串2
//	str2 := "string"
//
//	// 字符串3 = 字符串1 与字符串2 的拼接
//	str3 := str1 + str2
//
//	// 输出结果{ My name is string }
//	fmt.Println(str3)
//}

/*--------------------------------------------------------------------------------------------------------------------*/
//func main() {
//
//	// 创建类型
//	var mStr map[int]string
//
//	// 开辟空间
//	mStr = make(map[int]string)
//
//	// 设置值
//	mStr[1] = "one"
//	mStr[2] = "two"
//	mStr[3] = "three"
//	mStr[4] = "four"
//
//	// 输出结果{map[1:one 2:two 3:three 4:four]}
//	fmt.Println(mStr)
//}

/*--------------------------------------------------------------------------------------------------------------------*/
//func main() {
//
//	// 定义结构体类型
//	type strInfo struct {
//		typeName string
//	}
//
//	// 声明变量
//	var objStrInfo strInfo
//
//	// 结构体属性赋值
//	objStrInfo.typeName = "strInfo"
//
//	fmt.Println("结构体类型为:", objStrInfo.typeName)
//}

/*--------------------------------------------------------------------------------------------------------------------*/
//
//// 定义类
//type strInfo struct {
//	typeName string // 属性
//}
//
//// 通过类的方法设置属性
//func (this *strInfo) SetTypeName(str1 string) {
//	this.typeName = str1
//}
//
//// 通过类的方法获取属性
//func (this *strInfo) GetTypeName() string {
//	return this.typeName
//}
//
//func main() {
//
//	// 声明变量
//	var objStrInfo strInfo
//
//	// 结构体属性赋值
//	objStrInfo.SetTypeName("strInfo")
//
//	// 输出结果{ 结构体类型为: strInfo }
//	fmt.Println("结构体类型为:", objStrInfo.GetTypeName())
//}

/*--------------------------------------------------------------------------------------------------------------------*/

//// 接口
//type objIF interface {
//	SetName(string)
//	GetName() string
//	Speak()
//}
//
//// 实现接口给的类1
//type obj1 struct {
//	myName string
//}
//
//func (this *obj1) SetName(name string) {
//	this.myName = name
//}
//
//func (this *obj1) GetName() string {
//	return this.myName
//}
//
//func (this *obj1) Speak() {
//	fmt.Printf("My name is %s, skill is 无敌\n", this.myName)
//}
//
//// 实现接口给的类2
//type obj2 struct {
//	myName string
//}
//
//func (this *obj2) SetName(name string) {
//	this.myName = name
//}
//
//func (this *obj2) GetName() string {
//	return this.myName
//}
//
//func (this *obj2) Speak() {
//	fmt.Printf("My name is %s, skill is 隐身\n", this.myName)
//}
//
//func main() {
//	IF := []objIF{}
//
//	// 开辟空间设置属性
//	people1 := new(obj1)
//	people1.SetName("张三")
//	IF = append(IF, people1)
//
//	// 开辟空间设置属性
//	people2 := new(obj2)
//	people2.SetName("李四")
//	IF = append(IF, people2)
//
//	// 让他们进行发言
//	for _, v := range IF {
//		v.Speak()
//	}
//}

/*--------------------------------------------------------------------------------------------------------------------*/
//func func1(inPara1, inPara2 string) string {
//	return fmt.Sprintf(" %s，%s", inPara1, inPara2)
//}
//
//func main() {
//
//	// 我没有接受者，所以我是普通函数
//	fmt.Println(func1("我没有接受者", "所以我是普通函数"))
//}
/*--------------------------------------------------------------------------------------------------------------------*/

//func main() {
//
//	// 没有设置具体的函数名称
//	func1 := func(inPara1, inPara2 string) string {
//		return fmt.Sprintf(" %s，%s", inPara1, inPara2)
//	}
//
//	// 我的函数名为变量名，没有具体的名字，所以我是匿名函数
//	fmt.Println(func1("我的函数名为变量名，没有具体的名字", "所以我是匿名函数"))
//}
/*--------------------------------------------------------------------------------------------------------------------*/
//
//func main() {
//
//	a := 0
//	for i := 0; i < 5; i++ {
//		// 匿名函数内部可以直接使用外部的变量，并可对其进行修改。称之为闭包
//		func() {
//			a++
//		}()
//		// 输出 a is 1  a is 2  a is 3  a is 4  a is 5
//		fmt.Printf("a is %d \t", a)
//	}
//}
/*--------------------------------------------------------------------------------------------------------------------*/
//// 自定义类型
//type myStr string
//
//// 方法
//func (this *myStr) mod1() {
//	fmt.Println("我有接受者，所以我是方法，不是普通函数")
//}
//
//func main() {
//
//	// 初始化接收者
//	var objMyStr myStr
//	// 通过接收者调用方法
//	objMyStr.mod1()
//	// 输出为{我有接受者，所以我是方法，不是普通函数}
//}
/*--------------------------------------------------------------------------------------------------------------------*/

//// 普通管道
//func main() {
//	// 创建缓存管道
//	ch1 := make(chan int, 5)
//	// 创建生产者
//	go func() {
//		for i := 0; i < 5; i++ {
//			ch1 <- i
//		}
//		close(ch1) // 切记：需要关闭管道
//	}()
//	time.Sleep(100)
//	// 创建消费者
//	for {
//		if v, ok := <-ch1; !ok {
//			break
//		} else {
//			fmt.Println(v)
//		}
//	}
//	fmt.Println("结束")
//}
/*-------------------------------------------------------------------------------------------------------------*/
//// 匿名管道
//func main() {
//	// 创建没有缓存的管道:这种管道主要用于通知
//	ch1 := make(chan int) // 通知消费者打印
//	ch2 := make(chan int) // 通知主进程退出
//	var arrt []int        // 局部容器
//	// 消费者
//	go func() {
//		//当管道里面有了数据才会进行执行，没有数据会一直阻塞在这里
//		fmt.Println("我在等通知")
//		<-ch1
//		fmt.Println("开始消费")
//		for _, v := range arrt {
//			fmt.Printf("%v\t ", v)
//		}
//		fmt.Println("消费结束")
//		ch2 <- 1
//		close(ch2)
//	}()
//	time.Sleep(10) // 等待
//	fmt.Println("开始生产")
//	// 生产者
//	for i := 0; i < 5; i++ {
//		arrt = append(arrt, i)
//	}
//	fmt.Println("生产完成")
//	close(ch1)
//	<-ch2 // 接受消费者通知退出
//	fmt.Println("结束")
//}

//// map 测试取不存在的值，string 返回的值是什么
//func main() {
//
//	map1 := make(map[string]string)
//	map1["123"] = "123"
//
//	if map1["321"] == "" {
//		fmt.Println("is ")
//	}
//	// system Out"is "
//}

//func worker(ch chan struct{}) {
//	<-ch // 阻塞等待
//	fmt.Println("do something")
//	close(ch)
//}
//
//func main() {
//	ch := make(chan struct{})
//	go worker(ch)
//	ch <- struct{}{}             // 通知协程执行
//	time.Sleep(10 * time.Second) // 这需要睡眠等待，否则，子协程还没有执行，就进程就退出了
//}

// 关闭一个没有缓存的chennel, 对这个管道再读，会读到 0 或者 nil
//func main() {
//
//	ch1 := make(chan (int))
//	close(ch1)
//
//	fmt.Println("%v", <-ch1)
//
//	//strings.Split()
//}

// json 序列化-反序列化
//type Name1 struct {
//	Name_string []string `json:"nameString"`
//	Age         int      `json:"age"`
//}
//
//func main() {
//
//	name1 := Name1{
//		[]string{
//			"tian",
//			"li",
//			"jun",
//		},
//		28,
//	}
//
//	strBuff, _ := json.Marshal(&name1)
//	fmt.Println(string(strBuff))
//
//	var name2 Name1
//	json.Unmarshal(strBuff, &name2)
//	fmt.Printf("%+v", name2)
//}

type ResultCluster struct {
	Id           int       `json:"id" `
	CreatedAt    time.Time `json:"created_at" `
	UpdatedAt    time.Time `json:"updated_at" `
	ClusterName  string    `json:"clustername" `
	Version      string    `json:"version" `
	MasterNum    int       `json:"masternum" `
	WorkerNum    int       `json:"workernum" `
	MasterIp     string    `json:"masterip" `
	WorkerIp     string    `json:"workerip" `
	MasterPasswd string    `json:"masterpasswd" `
	WorkerPasswd string    `json:"workerpasswd" `
	MasterUser   string    `json:"masteruser" `
	WorkerUser   string    `json:"workeruser" `
}

func main() {
	data := "{\n    \"version\":\"kubernetes\",\n    \"clustername\":\"myclust0000\",\n    \"masterip\":\"[10.10.15.74]\",\n    \"workerip\":\"[10.10.15.51,10.10.15.64]\",\n    \"masteruser\":\"root\",\n    \"masterpasswd\":\"1\",\n    \"workeruser\":\"root\",\n    \"workerpasswd\":\"1\"\n}"

	var buff ResultCluster
	err := json.Unmarshal([]byte(data), &buff)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("%+v", buff)
}
