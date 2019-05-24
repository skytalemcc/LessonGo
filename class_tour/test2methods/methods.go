package main

/***
Go语言没有类。你可以为结构体类型定义方法。
方法就是一类带特殊的 接收者 参数的函数。
方法其实就是一个函数，在 func 这个关键字和方法名中间加入了一个特殊的接收器类型。
接收器可以是结构体类型或者是非结构体类型。接收器是可以在方法的内部访问的。


方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
func (t Type) methodName(parameter list) {
}

为什么不直接方法，而使用结构体类型定义方法
1,Go 不是纯粹的面向对象编程语言，而且Go不支持类。因此，基于类型的方法是一种实现和类相似行为的途径。
2,相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的。

当一个函数有一个值参数，它只能接受一个值参数。带指针参数的函数必须接受一个指针。
当一个方法有一个值接收器，它可以接受值接收器和指针接收器。而以指针为接收者的方法被调用时，接收者既能为值又能为指针。

为了在一个类型上定义一个方法，方法的接收器类型定义和方法的定义应该在同一个包中。
到目前为止，我们定义的所有结构体和结构体上的方法都是在同一个 main 包中，因此它们是可以运行的。



Go 支持面向对象吗？
Go 并不是完全面向对象的编程语言。
虽然 Go 有类型和方法，支持面向对象的编程风格，但却没有类型的层次结构。Go 中的“接口”概念提供了一种不同的方法，我们认为它易于使用，也更为普遍。
Go 也可以将结构体嵌套使用，这与子类化（Subclassing）类似，但并不完全相同。
此外，Go 提供的特性比 C++ 或 Java 更为通用：子类可以由任何类型的数据来定义，甚至是内建类型（如简单的“未装箱的”整型）。这在结构体（类）中没有受到限制。


Go 使用结构体，而非类
Go 不支持类，而是提供了结构体。结构体中可以添加方法。这样可以将数据和操作数据的方法绑定在一起，实现与类相似的效果。

***/

import (
	"fmt"
)

type Employee struct {
	name     string
	salary   int
	currency string
}

/*
  displaySalary() 方法将 Employee 做为接收器类型
*/
//记住：方法只是个带接收者参数的函数。
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary) //接收器是可以在方法的内部访问的
}

//普通函数
func displaySalaryfunc(e Employee) {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary) //接收器是可以在方法的内部访问的
}

//你也可以为非结构体类型声明方法，例如某个变量带方法。就是接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。

/***指针接收者
你可以为指针接收者声明方法。 指针接受 代表可以修改值。

这意味着对于某类型 T，接收者的类型可以用 *T 的文法。
指针接收者的方法可以修改接收者指向的值 。
于方法经常需要修改它的接收者，指针接收者比值接收者更常用。

如果不带*指针的话 ，方法只是对接收者的副本进行了操作，并不影响原值。

值接收器和指针接收器之间的区别在于，在指针接收器的方法内部的改变对于调用者是可见的，然而值接收器的情况不是这样的


那么什么时候使用指针接收器，什么时候使用值接收器？
一般来说，指针接收器可以使用在：对方法内部的接收器所做的改变应该对调用者可见时。
指针接收器也可以被使用在如下场景：当拷贝一个结构体的代价过于昂贵时。考虑下一个结构体有很多的字段。
在方法内使用这个结构体做为值接收器需要拷贝整个结构体，这是很昂贵的。
在这种情况下使用指针接收器，结构体不会被拷贝，只会传递一个指针到方法内部使用。
在其他的所有情况，值接收器都可以被使用。

指针接收器使用场合：
方法能够修改其接收者指向的值。
这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。

***/
//比较 带指针和不带指针 对值的影响
func (e Employee) ChangeSalary1() { //使用值接收器的方法。
	e.salary = e.salary + 1000

}

func (e *Employee) ChangeSalary2() { //使用指针接收器的方法。
	e.salary = e.salary + 1000

}

//属于结构体的匿名字段的方法可以被直接调用，就好像这些方法是属于定义了匿名字段的结构体一样。
//都是内部方法才行。  结构体内的结构体方法 可以被直接使用。
type person struct {
	name string
	age  int
	location
}

type location struct {
	city string
	loc  string
}

func (e location) displaylocation() {
	fmt.Printf("location of %s is %s\n", e.city, e.loc)
}

func main() {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary() // 调用 Employee 类型的 displaySalary() 方法

	emp1.ChangeSalary1() //调用 Employee 类型的 ChangeSalary1()方法,无指针 。函数不会对结构体的值造成影响。只会操作结构体的副本。
	fmt.Printf("New Salary of %s is %s%d\n", emp1.name, emp1.currency, emp1.salary)
	emp1.ChangeSalary2() //调用 Employee 类型的 ChangeSalary1()方法,指针接收者 。函数可以直接修改结构体的值。只会操作结构体的本体。
	fmt.Printf("New Salary of %s is %s%d\n", emp1.name, emp1.currency, emp1.salary)

	p := person{
		name: "james",
		age:  30,
		location: location{
			city: "Pairs",
			loc:  "Road Shanlije 11th",
		},
	}

	fmt.Println(p)
	p.location.displaylocation() //结构体内的结构体的方法，写完整的 .如果是公有字段，只能这么用，不能下面这种。
	p.displaylocation()          // 结构体内的结构体的方法，直接可以被上一层结构体拿来直接使用，但是只有私有(即匿名字段可以这么用)

	//方法与指针重定向 接受一个值作为参数的函数必须接受一个指定类型的值
	//普通函数
	displaySalaryfunc(emp1) //返回OK
	//displaySalaryfunc(&emp1)  //编译报错。接受一个值作为参数的函数必须接受一个指定类型的值。传参是啥类型就是啥类型。

	//而以值为接收者的方法被调用时，接收者既能为值又能为指针：
	emp1.displaySalary()    //接收者为值
	(&emp1).displaySalary() //接收者为指针

	/*
	   	使用 New() 函数，而非构造器
	   如果 Employee为零值，如果对变量零值使用
	   Go 并不支持构造器。如果某类型的零值不可用，需要程序员来隐藏该类型，避免从其他包直接访问。
	   程序员应该提供一种名为 NewT(parameters) 的 函数，按照要求来初始化 T 类型的变量。
	   按照 Go 的惯例，应该把创建 T 类型变量的函数命名为 NewT(parameters)。这就类似于构造器了。

	   虽然 Go 不支持类，但结构体能够很好地取代类，而以 New(parameters) 签名的方法可以替代构造器。
	   e := employee.New("Sam", "Adolf", 30, 20)
	   e.LeavesRemaining()

	   func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
	       e := employee {firstName, lastName, totalLeave, leavesTaken}
	       return e
	   }
	*/

	/*组合取代继承 。Go 不支持继承，但它支持组合（Composition）。组合一般定义为“合并在一起”。
	通过嵌套结构体进行组合。在 Go 中，通过在结构体内嵌套结构体，可以实现组合。
	多态
	Go 通过接口来实现多态。
	一个类型如果定义了接口所声明的全部方法，那它就实现了该接口。现在我们来看看，利用接口，Go 是如何实现多态的。



	*/

}

/***
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test2methods# go run methods.go
Salary of Sam Adolf is $5000
New Salary of Sam Adolf is $5000
New Salary of Sam Adolf is $6000
{james 30 {Pairs Road Shanlije 11th}}
location of Pairs is Road Shanlije 11th
location of Pairs is Road Shanlije 11th
Salary of Sam Adolf is $6000
Salary of Sam Adolf is $6000
Salary of Sam Adolf is $6000
root@e7939faf8694:/go/src/LessonGo/class_tour/test2methods#
***/
