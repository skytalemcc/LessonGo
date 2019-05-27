package main

/***
golang是一个开源的编译型静态语言。
Golang 的主要关注点是使得高可用性和可扩展性的 Web 应用的开发变得简便容易。Go 的定位是系统编程语言，只是对 Web 开发支持较好。

为什么要选择 Golang 作为服务端编程语言？
1,并发是语言的一部分，所以编写多线程程序会是一件很容易的。并发是通过 Goroutines 和 channels 机制实现的。
2,Golang 是一种编译型语言。源代码会编译为二进制机器码。而在解释型语言中没有这个过程。
编译型语言在程序执行之前，有一个单独编译过程，将程序翻译成机器语言，以后执行这个程序时，就不用再进行翻译了。如果换平台 要重新编译，例如windows换linux，bsd等。
解释型语言，是在运行的时候将程序翻译成机器语言，所以运行速度相对于编译型语言要慢。例如java要依靠编译器成字节码class文件，只要windows或者linux装有java虚拟机，就可以跨平台。
3,语言规范十分简洁。
4,Go 编译器支持静态链接。所有 Go 代码都可以静态链接为一个大的二进制文件，就直接可以将此文件跨平台跨机器执行。
可以轻松部署到云服务器，而不必担心各种依赖性。

所有 Go 源文件都应该放置在工作区里的 src 目录下。请在刚添加的 go 目录下面创建目录 src。
所有 Go 项目都应该依次在 src 里面设置自己的子目录。

多行注释可以用 \/* ... *\/ 来包裹，和其它大多数语言一样。
在文件一开头的注释一般都是这种形式，或者一大段的解释性的注释文字也会被这符号包住，来避免每一行都需要加//。
在注释中//和/*是没什么意义的，所以不要在注释中再嵌入注释。

***/
import "fmt"

//同一个目录下面不能有个多 package main ，否则IDE会提示错误 ： main redeclared in this block
func main() {

	fmt.Println("hello world")
}

/***
结果集
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic# go run hello.go
hello world
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic#
***/

// 用于单行注释
/* 用于单行注释 */
/*** 用于多行注释
***/
