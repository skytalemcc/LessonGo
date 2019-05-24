package main

/***
数组是同一类型元素的集合。
Go 语言中不允许混合不同类型的元素，例如包含字符串和整数的数组。
当然，如果是 interface{} 类型数组，可以包含任意类型。

数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。
Go 语言数组声明需要指定元素类型及元素个数
数组定义：
数组声明需要指定元素类型及元素个数
数组里面的值 是有顺序的值，可以根据索引来进行取值。
一维数组  ，就像excel 中的一行
var variable_name [SIZE] variable_type
多维数组 ，就是excel中的多行
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type

数组元素可以通过索引（位置）来读取（或者修改），索引从 0 开始，第一个元素索引为 0，第二个索引为 1，以此类推。




***/

import "fmt"

var balance [4]int //声明长度为4的整型数组
var depth []int    //如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小

var excels [3][4]int //三行四列的数组 数组中的所有元素都被自动赋值为数组类型的零值

//向函数传递数组，数组作为形参
func getAverage(arr []int, size int) float32 {
	var i int
	var avg, sum float32
	for i = 0; i < size; i++ {
		sum += float32(arr[i])
	}
	avg = sum / float32(size)
	return avg
}

func main() {

	//初始化数组
	balance = [4]int{3, 4, 1, 2}     //给数组赋值 用{}大括号进行赋值
	fmt.Println(balance, balance[1]) //访问数组中的某个元素
	depth = []int{1, 2, 3, 4, 5, 6}  // 不声明长度的数组，进行赋值
	fmt.Println(depth)
	balance2 := [4]int{3, 4, 1, 2} //使用 简略声明 来创建相同的数组。
	fmt.Println(balance2)
	excels = [3][4]int{ //三行四列 多维数组
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 0, 7, 8}, //这种换行的大括号的话，这里需要加上逗号。如果不换行的，不需要加逗号。
	}
	fmt.Println(excels, excels[0][2]) //索引从0开始，多维数组通过指定坐标来访问，如数组中的行索引和列索引
	fmt.Println(getAverage(depth, 5)) //数组作为形参

	//Go 中的数组是值类型而不是引用类型。这意味着当数组赋值给一个新的变量时，该变量会得到一个原始数组的一个副本。
	//如果对新变量进行更改，则不会影响原始数组。
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a             // 只是把值赋给了a。
	b[0] = "Singapore" //对b进行修改值，不会影响a的值。
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)

	//数组的长度 通过将数组作为参数传递给 len 函数，可以得到数组的长度。
	a1 := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of a is", len(a1))
	//for 循环可用于遍历数组中的元素。
	for i := 0; i < len(a1); i++ { // looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %.2f\n", i, a1[i])
	}
	//Go 提供了一种更好、更简洁的方法，通过使用 for 循环的 range 方法来遍历数组。range 返回索引和该索引处的值。
	for i, v := range a1 { //range可以返回索引和值
		fmt.Printf("%d the element of a is %.2f\n", i, v)
	}
}

/***
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic# go run array.go
[3 4 1 2] 4
[1 2 3 4 5 6]
[3 4 1 2]
[[1 2 3 4] [5 6 7 8] [9 0 7 8]] 3
3
a is  [USA China India Germany France]
b is  [Singapore China India Germany France]
length of a is 4
0 th element of a is 67.70
1 th element of a is 89.80
2 th element of a is 21.00
3 th element of a is 78.00
0 the element of a is 67.70
1 the element of a is 89.80
2 the element of a is 21.00
3 the element of a is 78.00
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic#

***/
