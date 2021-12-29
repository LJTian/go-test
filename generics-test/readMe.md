# Tutorial: Getting started with generics(教程：泛型入门)  

## Table of Contents(目录)  

- Prerequisites(前提) 
- Create a folder for your code(创建代码文件夹) 
- Add non-generic functions(添加非泛型函数) 
- Add a generic function to handle multiple types(添加一个泛型函数 用来处理多种类型) 
- Remove type arguments when calling the generic function(调用泛型函数时移除类型参数) 
- Declare a type constraint(声明类型范围/约束) 
- Conclusion(结束语) 
- Completed code(最终代码) 

#### Note: This is beta content.(重点: 这只是测试办的内容) 
#### 原文: 
        This tutorial introduces the basics of generics in Go. With generics, you can declare and use functions or types 
    that are written to work with any of a set of types provided by calling code.
        In this tutorial, you’ll declare two simple non-generic functions, then capture the same logic in a single 
    generic function.
        You’ll progress through the following sections:
            1、Create a folder for your code.
            2、Add non-generic functions.
            3、Add a generic function to handle multiple types.
            4、Remove type arguments when calling the generic function.
            5、Declare a type constraint.
#### 翻译: 
        本教程介绍了 Go 中泛型的基础知识。使用泛型，可以声明和使用为使用调用代码提供的任何一组类型而编写的函数或类型。【注者：使用泛型可以为调用者提供
    不限制参数类型的函数和自定义的结构体/类/类型】
        在本教程中，有需要声明两个非泛型函数，然后提取这两个函数的相同处理逻辑生成新的泛型函数。
        您将学习以下部分：
            1、创建代码文件夹/目录
            2、添加非泛型函数
            3、添加一个泛型函数，用来处理多种类型
            4、调用泛型函数，移除类型参数
            5、声明类型范围/约束

#### Note: For other tutorials, see Tutorials.(注意：有关其他教程，请参阅教程。) 
#### Note: If you prefer, you can use the Go playground in “Go dev branch” mode to edit and run your program instead.(如果您愿意，您可以在“Go dev branch”模式下使用 Go playground 来编辑和运行您的程序。) 

## 1、Prerequisites(前提) 
- **An installation of Go 1.18 Beta 1 or later.** For installation instructions, see Installing and using the beta.(**Go 1.18 Beta 1 或更高版本的安装。** 有关安装说明，请参阅安装和使用测试版。) 
- **A tool to edit your code.** Any text editor you have will work fine.(**编辑代码的工具.** 任何文本编辑器都可以正常工作。) 
- **A command terminal.** Go works well using any terminal on Linux and Mac, and on PowerShell or cmd in Windows. (**命令终端。** Go 适用于 Linux 和 Mac 上的任何终端，以及 Windows 中的 PowerShell 或 cmd。) 

### Installing and using the beta(安装和使用测试版本) 

#### 原文: 
This tutorial requires the generics feature available in Beta 1. To install the beta, following these steps:
 
- 1、Run the following command to install the beta.  
  `$ go install golang.org/dl/go1.18beta1@latest`
- 2、Run the following command to download updates.  
  `$ go1.18beta1 download`
- 3、Run go commands using the beta instead of a released version of Go (if you have one).  
  You can run commands with the beta either by using the beta name or by aliasing the beta to another name. 
  - Using the beta name, you can run commands by invoking go1.18beta1 instead of go:  
`
$ go1.18beta1 version
`
  - By aliasing the beta name to another name, you can simplify the command:  
`
$ alias go=go1.18beta1 && go version
`

Commands in this tutorial will assume you have aliased the beta name.
             
#### 翻译:
本步骤需要使用beta 1.版本中的泛型功能。请安装对应版本，安装步骤如下：
- 1、执行以下命令安装beta 版本。  
    `$ go install golang.org/dl/go1.18beta1@latest`
- 2、执以下命令进行更新  
`$ go1.18beta1 download`
- 3、使用beta版本，而不是go的发布版本(如果你本地已安装发布版本)的执行命令:go  
  你可以通过go1.18beta1(使用beta名称), 或者通过给go创建别名(linux中的alias 方式)运行beta版本
  - 使用beta名称，你可以通过go1.18beata1运行，而不是go：  
  `$ go1.18beta1 version`                        
  - 也可以通过别名来运行go的测试版本：  
  `$ alias go=go1.18beta1 && go version`

本教程中的命令假定为通过别名方式来运行【译者注：后面使用到的"go"命令都为"go1.18beta1"命令】

## Create a folder for your code(创建代码文件夹)  
To begin, create a folder for the code you’ll write.(首先，你要为你编写的代码创建一个文件夹)
### 1、Open a command prompt and change to your home directory.(打开一个命令终端，并移至你的Home目录下)
#### On Linux or Mac(linux或者Mac系统上的命令为):
`$ cd`
#### On Windows(Windows系统下命令为):
`C:\> cd %HOMEPATH%`  
The rest of the tutorial will show a $ as the prompt. The commands you use will work on Windows too.(本教程的其余部分将显示
$ 作为提示。您使用的命令也适用于 Windows。)
### 2、From the command prompt, create a directory for your code called generics.(在命令终端下，创建一个名为"generics"的目录)
```shell
  $ mkdir generics
  $ cd generics    
```
### 3、Create a module to hold your code.(给你的代码创建一个mod文件进行包管理)
Run the go mod init command, giving it your new code’s module path.(运行"go mod init"命令，提供你新代码module的路径)
```shell
   $ go mod init example/generrics
   go: creating new go.mod: module example/generics ##此部分是cmd终端自动输出的
```
***Note:***  For production code, you’d specify a module path that’s more specific to your own needs. For more, be sure 
to see Managing dependencies.(注意：对于生产代码，您需要指定一个更符合您自己需求的模块路径。有关更多信息，请务必参阅管理依赖项。)

Next, you’ll add some simple code to work with maps.(接下来，你将添加一些简单的代码来处理Maps)

## Add non-generic functions(添加非泛型函数)

In this step, you’ll add two functions that each add together the values of a map and return the total.
(在此步骤中，您将添加两个函数，每个函数将map中的值相加并返回总和。)  
You’re declaring two functions instead of one because you’re working with two different types of maps: 
one that stores int64 values, and one that stores float64 values.
(你需要声明了两个函数而不是一个函数，因为您正在使用两种不同类型的map：一种存储 int64 类型值，另一种存储 float64 类型值。)

###Write the code(编写代码)

#### 原文:
- 1.Using your text editor, create a file called main.go in the generics directory. You’ll write your Go code in this file.
- 2.Into main.go, at the top of the file, paste the following package declaration.  
`package main`  
  A standalone program (as opposed to a library) is always in package main.
- 3.Beneath the package declaration, paste the following two function declarations.
```go
// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}
```

#### 翻译:
- 1.使用你的文本编辑器，在"generics"目录中创建一个名为"main.go"的文件，你将在此文件中编写go的代码.
- 2.在"mian.go"文件的头部，粘贴下面的包声明
  `package main`  
  一个独立的执行程序(与包相对)都位于main包中
- 3.在包声明下方，粘贴以下两个函数声明
```go
// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
var s int64
for _, v := range m {
s += v
}
return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
var s float64
for _, v := range m {
s += v
}
return s
}
```


## Add a generic function to handle multiple types(添加一个泛型函数 用来处理多种类型)

## Remove type arguments when calling the generic function(调用泛型函数时移除类型参数) 

## Declare a type constraint(声明类型范围/约束) 

## Conclusion(结束语)

## Completed code(最终代码) 