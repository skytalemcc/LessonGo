package main

import "fmt"

/***
变量  var 语句定义了一个变量的列表；跟函数的参数列表一样，类型在后面。
变量指定了某存储单元（Memory Location）的名称，该存储单元会存储特定类型的值。
由于 Go 是强类型（Strongly Typed）语言，因此不允许某一类型的变量赋值为其他类型的值。
***/
var c, python, java bool
var d, j int = 1, 2 //变量定义可以包含初始值，每个变量对应一个。

func main() {

	var i int                          //声明变量，没有初始化值。var i int 29  则为声明变量并初始化值
	k := 3                             //短声明变量(简短声明)，在函数中可以使用，可以替代var。函数外都必须以关键字开始，短声明不能使用在函数外部。
	fmt.Println(i, c, python, java, k) //有值则打印值，无值则打印类型的初始值。
	fmt.Println(d, j)

	var q, p, o = true, false, 2      //如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。  类型推断。
	fmt.Printf("%T %T %T\n", q, p, o) //Printf 格式化输出，%T代表打印变量的类型。
}

/***
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic# go run variables.go
0 false false false 3
1 2
bool bool int
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic#

***/
