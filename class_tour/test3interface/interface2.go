package main

/***
接口的实际用途：
接口可以定义一系列方法的集合。
方法接收者不一样，但是方法名可以一样。
当你需要扩展的时候 ，对接口不需要修改，只需要对接收者定义的方法进行修改和定义就可以。

使用值接受者声明的方法，既可以用值来调用，也能用指针调用。
不管是一个值，还是一个可以解引用的指针，调用这样的方法都是合法的。

类型可以实现多个接口。
***/

import (
	"fmt"
)

//尽管 Go 语言没有提供继承机制，但可以通过嵌套其他的接口，创建一个新接口。
type EmployeeOperations interface { //新接口里面嵌套老接口 和别的方法 来定义
	SalaryCalculator
}

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d\n", expense)
}

//类型选择 类型选择中的声明与类型断言 i.(T) 的语法相同，只是具体类型 T 被替换成了关键字 type。
func interswitch(i interface{}) {
	switch v := i.(type) { //type来做类型选择 ，固定语法
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)

	//类型断言 提供了访问接口值底层具体值的方式。 t := i.(T)
	//该语句断言接口值 i 保存了具体类型 T，并将其底层类型为 T 的值赋予变量 t。
	//若 i 并未保存 T 类型的值，该语句就会触发一个报错。
	var i interface{} = "hello"
	s := i.(string) //类型断言，接口里面存在string类型，所以不会报错
	fmt.Println(s)
	//为了 判断 一个接口值是否保存了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。
	s, ok := i.(string)
	fmt.Println(s, ok)
	//f = i.(float64) 这样写会报错(panic)，因为没有这种类型
	f, ok := i.(float64) //这样写不会报错，因为加了判断
	fmt.Println(f, ok)

	//类型选择 是一种按顺序从几个类型断言中选择分支的结构。
	//类型选择与一般的 switch 语句相似，不过类型选择中的 case 为类型（而非值）， 它们针对给定接口值所存储的值的类型进行比较。

	interswitch(21)
	interswitch("hello")
	interswitch(true)
}

/***
root@e7939faf8694:/go/src/LessonGo/class_tour/test3interface# go run interface2.go
Total Expense Per Month $14050
hello
hello true
0 false
Twice 21 is 42
"hello" is 5 bytes long
I don't know about type bool!
root@e7939faf8694:/go/src/LessonGo/class_tour/test3interface#

***/
