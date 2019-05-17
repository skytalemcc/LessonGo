/***
函数可以没有参数，或者接受多个参数，注意类型在变量名之后。可以有返回，也可以没有返回。
当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。
函数可以返回任意数量的返回值。
Go 的返回值可以被命名，并且像变量那样使用。
返回值的名称应当具有一定的意义，可以作为文档使用。
没有参数的 return 语句返回结果的当前值。也就是`直接`返回。
***/

package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func print(x int, y int) {
	fmt.Println(x, y)
}

func swap(x, y string) (string, string) {
	return y, x
}

//Go 的返回值可以被命名，并且像变量那样使用。没有参数的 return 语句返回结果的当前值。也就是`直接`返回。
func swap2(x, y string) (m, n string) {
	m = y
	n = x
	return
}

func main() {
	fmt.Println(add(3, 4))
	print(5, 6)
	fmt.Println(swap("hello", "world"))
	a, b := swap("world", "hello")
	fmt.Println(a, b)
	fmt.Println(swap2("hello", "world"))
}
