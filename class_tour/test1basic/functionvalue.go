package main

/***
函数值
函数也是值。它们可以像其它值一样传递。
函数值可以用作函数的参数或返回值。
在Go语言中，函数是一种类型，而且是第一类型(first-class)。他的地位和int string等类型是一样的。
我们经常会声明一个值类型为int或者string类型的变量，现在我们可以声明一个值类型为某个函数的变量，
这种变量叫做函数变量也就是说，函数可以被当做一个值类型赋值给变量
***/

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

/***
Go 函数可以是一个闭包。
闭包是一个函数值，它引用了其函数体之外的变量。该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。

***/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

//斐波那契闭包
// 返回一个“返回int的函数”
func fibonacci() func(i int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

//例如，函数 adder 返回一个闭包。每个闭包都被绑定在其各自的 sum 变量上。
func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}

/***
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic# go run functionvalue.go
13
5
81
0 0
1 -2
3 -6
6 -12
10 -20
15 -30
21 -42
28 -56
36 -72
45 -90
0
1
3
6
10
15
21
28
36
45
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic#
***/
