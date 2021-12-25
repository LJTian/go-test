#Tutorial: Getting started with generics(教程：泛型入门)

##Table of Contents(目录)

- 1、Prerequisites(前提)
- 2、Create a folder for your code(创建代码文件夹)
- 3、Add non-generic functions(添加非泛型函数)
- 4、Add a generic function to handle multiple types(添加一个泛型函数 用来处理多种类型)
- 5、Remove type arguments when calling the generic function(调用泛型函数时移除类型参数)
- 6、Declare a type constraint(声明类型范围/约束)
- 7、Conclusion(结束语)
- 8、Completed code(最终代码)

####Note: This is beta content.(重点: 这只是测试办的内容)
####原文:
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
####翻译:
        本教程介绍了 Go 中泛型的基础知识。使用泛型，可以声明和使用为使用调用代码提供的任何一组类型而编写的函数或类型。【注者：使用泛型可以为调用者提供
    不限制参数类型的函数和自定义的结构体/类/类型】
        在本教程中，有需要声明两个非泛型函数，然后提取这两个函数的相同处理逻辑生成新的泛型函数。
        您将学习以下部分：
            1、创建代码文件夹/目录
            2、添加非泛型函数
            3、添加一个泛型函数，用来处理多种类型
            4、调用泛型函数，移除类型参数
            5、声明类型范围/约束

####Note: For other tutorials, see Tutorials.(注意：有关其他教程，请参阅教程。)
####Note: If you prefer, you can use the Go playground in “Go dev branch” mode to edit and run your program instead.(如果您愿意，您可以在“Go dev branch”模式下使用 Go playground 来编辑和运行您的程序。)

##1、Prerequisites(前提)
- **An installation of Go 1.18 Beta 1 or later.** For installation instructions, see Installing and using the beta.(**Go 1.18 Beta 1 或更高版本的安装。** 有关安装说明，请参阅安装和使用测试版。)
- **A tool to edit your code.** Any text editor you have will work fine.(**编辑代码的工具.** 任何文本编辑器都可以正常工作。)
- **A command terminal.** Go works well using any terminal on Linux and Mac, and on PowerShell or cmd in Windows. (**命令终端。** Go 适用于 Linux 和 Mac 上的任何终端，以及 Windows 中的 PowerShell 或 cmd。)

###Installing and using the beta(安装和使用测试版本)

####原文:
    This tutorial requires the generics feature available in Beta 1. To install the beta, following these steps:
        1、Run the following command to install the beta.
            $ go install golang.org/dl/go1.18beta1@latest
        2、Run the following command to download updates.
            $ go1.18beta1 download
        3、Run go commands using the beta instead of a released version of Go (if you have one).
            You can run commands with the beta either by using the beta name or by aliasing the beta to another name.
                · Using the beta name, you can run commands by invoking go1.18beta1 instead of go:
                    