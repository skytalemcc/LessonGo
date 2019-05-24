/***
函数是一块执行特定任务的代码。一个函数是在输入源基础上，通过执行一系列的算法，生成预期的输出。
函数可以没有参数，或者接受多个参数，注意类型在变量名之后。可以有返回，也可以没有返回。
当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。
函数可以返回任意数量的返回值。
Go 的返回值可以被命名，并且像变量那样使用。
返回值的名称应当具有一定的意义，可以作为文档使用。
没有参数的 return 语句返回结果的当前值。也就是`直接`返回。
***/

package main

import "fmt"

/*
 函数首字母大写。在 Go 中这具有特殊意义。在 Go 中，任何以大写字母开头的变量或者函数都是被导出的名字。其它包只能访问被导出的函数和变量。
*/

/*

所有包都可以包含一个 init 函数。init 函数不应该有任何返回值类型和参数，在我们的代码中也不能显式地调用它。
init 函数可用于执行初始化任务，也可用于在开始执行之前验证程序的正确性。
包的初始化顺序如下：
1,首先初始化包级别（Package Level）的变量。外部。
2,紧接着调用 init 函数。包可以有多个 init 函数（在一个文件或分布于多个文件中），它们按照编译器解析它们的顺序进行调用。

如果一个包导入了另一个包，会先初始化被导入的包。
尽管一个包可能会被导入多次，但是它只会被初始化一次。
*/
func init() {
	fmt.Println("rectangle package initialized")
}

func add(x, y int) int { //函数中的参数列表和返回值并非是必须的
	return x + y
}

func print(x int, y int) {
	fmt.Println(x, y)
}

func swap(x, y string) (string, string) { //Go 语言支持一个函数可以有多个返回值
	return y, x
}

/*命名返回值:
Go 的返回值可以被命名，并且像变量那样使用。没有参数的 return 语句返回结果的当前值。也就是`直接`返回。
从函数中可以返回一个命名值。一旦命名了返回值，可以认为这些值在函数第一行就被声明为变量了。
*/
func swap2(x, y string) (m, n string) {
	m = y
	n = x
	return //不需要明确指定返回值，默认返回m,n 。m,n相当于函数中已经生成好的。
}

func main() {
	fmt.Println(add(3, 4))
	print(5, 6)
	fmt.Println(swap("hello", "world"))
	a, b := swap("world", "hello")
	//_ 在 Go 中被用作空白符，可以用作表示任何类型的任何值。
	//_, b := swap("world", "hello") 空白符表示我们并不需要它。
	fmt.Println(a, b)
	fmt.Println(swap2("hello", "world"))

	//什么是头等函数？
	//支持头等函数（First Class Function）的编程语言，可以把函数赋值给变量，也可以把函数作为其它函数的参数或者返回值。Go 语言支持头等函数的机制。

	p := func() { //func看看，由于没有名称，这类函数称为匿名函数（Anonymous Function）。
		fmt.Println("hello world first class function")
	}
	p()
	fmt.Printf("%T\n", p)

	//要调用一个匿名函数，可以不用赋值给变量。
	func() {
		fmt.Println("hello world first class function 2")
	}() // 使用 () 立即调用了该函数
	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers") //向匿名函数传递了一个字符串参数 并打印。

	/*高阶函数（Hiher-order Function）定义为：满足下列条件之一的函数：
	接收一个或多个函数作为参数
	返回值是一个函数
	*/
}

/***
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic# go run function.go
rectangle package initialized
7
5 6
world hello
hello world
world hello
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic#
***/
