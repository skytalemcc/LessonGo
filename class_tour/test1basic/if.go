package main

/***
if 语句除了没有了 `( )` 之外，而 `{ }` 是必须的。小括号不要，大括号必须
if 是条件语句。if 语句的语法是
if condition {  //  如果 condition 为真，则执行 { 和 } 之间的代码。
}

***/

import (
	"fmt"
	"math"
)

//求平方根
func sqrt(x float64) string {

	if x < 0 {
		return sqrt(-x) + "i"

	}
	//Print 将输入参数转换为 string 后, 写入标准输出。也就是程序运行时，我们可以在运行界面看到转换后的 string。
	//Sprint 仅完成将输入参数转换为String，不会写入标准输出。
	return fmt.Sprint(math.Sqrt(x)) //只是负责转换为String,不输出。所以Sprint适合配合return使用。
	//return fmt.Println(math.Sqrt(x)) 这种则会返回报错 too many arguments to return ,have (int, error),want (string)
	//return math.Sqrt(x)  如果使用则报错 ：cannot use math.Sqrt(x) (type float64) as type string in return argument
}

//if的便捷语句。跟 for 一样，`if` 语句可以在条件之前执行一个简单的语句。由这个语句定义的变量的作用域仅在 if 范围之内。
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { //Pow x的y次方  而变量v只限制在了if语句里面。
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim) //if else语句使用
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	num := 99
	if num <= 50 {
		fmt.Println("number is less than or equal to 50")
	} else if num >= 51 && num <= 100 { //多分支选择
		fmt.Println("number is between 51 and 100")
	} else {
		fmt.Println("number is greater than 100")
	}

}

/***
结果集:
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic# go run if.go
1.4142135623730951 2i
27 >= 20
9 20
number is between 51 and 100
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic#

***/
