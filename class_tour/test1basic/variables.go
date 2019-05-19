package main

import "fmt"

/***
变量  var 语句定义了一个变量的列表；跟函数的参数列表一样，类型在后面。
***/
var c, python, java bool
var d, j int = 1, 2 //变量定义可以包含初始值，每个变量对应一个。

func main() {

	var i int
	k := 3                             //短声明变量，在函数中可以使用，可以替代var。函数外都必须以关键字开始，短声明不能使用在函数外部。
	fmt.Println(i, c, python, java, k) //有值则打印值，无值则打印类型的初始值。
	fmt.Println(d, j)

	var q, p, o = true, false, 2      //如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。
	fmt.Printf("%T %T %T\n", q, p, o) //Printf 格式化输出，%T代表打印变量的类型。
}
