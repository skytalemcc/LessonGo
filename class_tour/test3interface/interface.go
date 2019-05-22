package main

/***
什么是接口：
在面向对象的领域里，接口一般这样定义：接口定义一个对象的行为。
接口只指定了对象应该做什么，至于如何实现这个行为（即实现细节），则由对象本身去确定。

在 Go 语言中，接口就是方法签名（Method Signature）的集合。当一个类型定义了接口中的所有方法，我们称它实现了该接口。
接口指定了一个类型应该具有的方法，并由该类型决定如何实现这些方法。
***/
import (
	"fmt"
)

//定义接口，方法签名的集合. 里面看似有两个方法，实际上方法为结构体定义方法。接口连同结构体一起继承了。

//类型通过实现一个接口的所有方法来实现该接口。既然无需专门显式声明，也就没有“implements”关键字。
//隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。
//因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。

//实现接口：指针接受者与值接受者
type Person interface {
	DisplaySalary()
	AddSalary()
}

type Person2 interface {
	DisplaySalary2()
	AddSalary2()
}

type Employee struct {
	name     string
	salary   int
	currency string
}

func (e Employee) DisplaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func (e Employee) AddSalary() {
	e.salary = e.salary + 1000
	fmt.Printf("New Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func (e *Employee) DisplaySalary2() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func (e *Employee) AddSalary2() {
	e.salary = e.salary + 1000
	fmt.Printf("New Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

//定义一个空接口类型的函数
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	//普通接口，里面含有方法和方法的接收器结构体 实现接口：值接受者
	var emp1 Person = Employee{ //给接口 里面方法的接收器 结构体赋值。
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	fmt.Println(emp1)
	emp1.DisplaySalary()

	//接口也是值。它们可以像其它值一样传递。接口值可以用作函数的参数或返回值。当结构体指针的时候
	//因为方法是指针接收者  . 实现接口：指针接受者
	//即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。
	var emp2 Person2 = &Employee{ // 接口中 的方法为 指针接收者声明方法。所以这里要声明指针接收者。
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	fmt.Println(emp2)
	emp2.DisplaySalary2()

	//指定了零个方法的接口值被称为 空接口,interface{}。空接口可保存任何类型的值。(因为每个类型都至少实现了零个方法。)
	//空接口被用来处理未知类型的值。
	var i interface{} //定义空接口类型
	describe(i)
	i = 42 //保存任意类型的值
	describe(i)
	i = "hello"
	describe(i)

}

/***
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test3interface# go run interface.go
{Sam Adolf 5000 $}
Salary of Sam Adolf is $5000
&{Sam Adolf 5000 $}
Salary of Sam Adolf is $5000
(<nil>, <nil>)
(42, int)
(hello, string)
root@e7939faf8694:/go/src/LessonGo/class_tour/test3interface#
***/
