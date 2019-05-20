package main

/***
一个结构体（`struct`）就是一个字段的集合。本身是一种类型。
结构体（struct）是用户自定义的类型，它代表若干字段的集合。
有些时候将多个数据看做一个整体要比单独使用这些数据更有意义，这种情况下就适合使用结构体。
如果一个结构体类型的名称以大写字母开头，则该结构体被导出，其他包可以访问它。
同样地，如果结构体中的字段名以大写字母开头，则这些字段也可以被其他包访问。

两个结构体 的 例子 是可以进行值比较的。
结构体是值类型，如果其字段是可比较的，则该结构体就是可以比较的。如果两个结构体变量的所有非空字段都相等，则认为这两个结构体变量相等。
如果结构体包含不可比较的类型的字段，那么这两个结构体是不可比较的。
***/

import "fmt"

//声明结构体,定义具名结构体变量
//下面的结构体 Employee 是一个具名结构体（named structure）
//因为它创建了一个具有名字的结构体类型： Employee 。我们可以定义具名结构体类型的变量。
type Employee struct {
	firstName, lastName string //相同类型的变量可以合并到一行，用逗号分隔。
	age, salary         int
}

//定义一个没有类型名称的结构体，这种结构体叫做匿名结构体（anonymous structures）。
var employee struct {
	firstName, lastName string
	age                 int
}

//匿名字段 定义结构体类型时可以仅指定字段类型而不指定字段名字。这种字段叫做匿名字段（anonymous field）。
//创建了一个 Person 结构体，它有两个匿名字段，类型为 string 和 int 。
type Person struct {
	string
	int
}

//结构体的字段也可以是一个结构体。这种结构体称为嵌套结构体。

type Address struct {
	city, state string
}

type Company struct {
	name    string
	age     int
	address Address
}

func main() {

	//使用字段名的写法
	emp1 := Employee{
		firstName: "Sam",
		lastName:  "chan",
		age:       25,
		salary:    10000,
	}
	fmt.Println("emp1 is ", emp1)
	//使用点 . 操作符来访问结构体中的字段。
	fmt.Println("emp1 age is ", emp1.age)

	//不使用字段名的写法
	emp2 := Employee{"Thomas", "Paul", 29, 20000}
	fmt.Println("emp2 is ", emp2)

	//定义匿名结构体变量 这种结构体成为匿名结构体，因为它只创建了一个新的结构体变量 emp3，而没有定义新的结构体类型。
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    30000,
	}
	fmt.Println("emp3 is ", emp3)

	//当定义一个结构体变量，但是没有给它提供初始值，则对应的字段被赋予它们各自类型的0值。
	var emp4 Employee
	fmt.Println("emp4 is ", emp4)
	//可以创建一个 0 值结构体变量，稍后给它的字段一一赋值。
	emp4.firstName = "Barry"
	fmt.Println("emp4 is ", emp4)

	//可以定义结构体指针
	emp5 := &Employee{"Sam", "Anderson", 55, 60000} //emp5 是一个结构体指针，为一个内存地址。
	fmt.Println("First Name:", (*emp5).firstName)   //(*emp5) 代表解引用得到的值
	fmt.Println("Age:", (*emp5).age)
	//在 Go 中我们可以使用 emp8.firstName 替代显示解引用 (*emp8).firstName 来访问 firstName 字段。具有一样的效果。
	fmt.Println("First Name:", emp5.firstName)
	fmt.Println("Age:", emp5.age)

	//使用匿名字段 虽然匿名字段没有名字，但是匿名字段的默认名字为类型名。
	p := Person{"Naveen", 50}
	fmt.Println(p, p.string, p.int)

	//结构体嵌套 结构体的字段也可以是一个结构体。这种结构体称为嵌套结构体。
	var cm Company
	cm.name = "Yahoo"
	cm.age = 30
	cm.address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println(cm, cm.name, cm.age, cm.address, cm.address.city, cm.address.state)
}
