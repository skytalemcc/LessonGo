package main

/***
指针是一种存储变量内存地址（Memory Address）的变量。
Go 具有指针。 指针保存了变量的内存地址。
指针是存储一个变量的内存地址的变量。
类型 *T 是指向类型 T 的值的指针。其零值是 `nil`。  星号 代表 类型值的指针地址。var p *int 针对类型挂着星号，则为指针声明
& 符号会生成一个指向其作用对象的指针。&操作符用来获取一个变量的地址。&代表取指针地址。
*p 解引用指针的意思是通过指针访问被指向的值。

*int 声明类型指针
&i 取指针
*p 解指针获得值

指针的解引用:
指针的解引用可以获取指针所指向的变量的值。将 a 解引用的语法是 *a。

***/

import "fmt"

func main() {

	i := 42
	var p *int = &i //左边代表 定义一个int类型的指针为p，就是内存地址块。&i 则为取i的变量地址 。等于则相当于同类型赋值了。声明指针并赋值。
	//p := &i //获取变量i的内存指针地址，取i的地址 ，则p为指针类型。
	fmt.Println("address of p is ", p) //打印出指针的内存地址

	//解引用指针的意思是通过指针访问被指向的值。指针解引用*p 和 基于类型的* 例如*int  ，是两码事情。我们将p解引用并打印这个解引用得到的值。
	fmt.Println("value of p is ", *p)
	*p = 43 //通过对解引用的结果进行重新赋值，其实此指针地址 是i的地址。
	fmt.Println(i)

	j := 66
	p = &j         //将j的地址给指针p ，p指针从i的地址 换位了p的地址
	*p = *p / 33   //解引用得到的值可以进行日常计算
	fmt.Println(j) //指针p的值是j的地址，*p改变了j的值。

}

/***
结果集:
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic# go run pointer.go
address of p is  0xc000074010
value of p is  42
43
2
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic#

***/
