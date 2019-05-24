package main

/***
什么是反射？
反射就是程序能够在运行时检查变量和值，求出它们的类型。

为何需要检查变量，确定变量的类型？
在学习反射时，所有人首先面临的疑惑就是：如果程序中每个变量都是我们自己定义的。
那么在编译时就可以知道变量类型了，为什么我们还需要在运行时检查变量，求出它的类型呢？

场景：
在 Go 语言中，reflect 实现了运行时反射。reflect 包会帮助识别 interface{} 变量的底层具体类型和具体值。
首先需要了解 reflect 包中的几种类型和方法：
reflect.Type 表示 interface{} 的具体类型，而 reflect.Value 表示它的具体值。
reflect.TypeOf() 和 reflect.ValueOf() 两个函数可以分别返回 reflect.Type 和 reflect.Value。


清晰优于聪明。而反射并不是一目了然的。
反射是 Go 语言中非常强大和高级的概念，应该小心谨慎地使用它。使用反射编写清晰和可维护的代码是十分困难的。

***/

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

//定义一个大的反射函数 ，传入任意值，然后根据反射来判断类型 组成查询语句返回。
func createQuery3(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()                     //从类型获取结构体名称。
		query := fmt.Sprintf("insert into %s values(", t) //根据结构体名称 来定表名。
		v := reflect.ValueOf(q)                           //取结构体的值 ，要进行遍历。
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() { //对结构体的值的类型，进行判断。
			case reflect.Int: //判断上面结果是否为reflect.Int类型。
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int()) //i==0 代表把结构体值的第一个转为sting。拼接语句，第一个参数传入到query语句组合
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int()) //i>0以后开始循环，将上一次已经拼接的好的语句加，逗号和新的字段拼接起来。
				}
			case reflect.String: //这样相当于 新表 第一个传入值可以为int，可以为string 然后和query语句进行拼接。然后循环新的 不同类型字段进行拼接。字符串加引号，转移。
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default: //这个是没有匹配的时候走的路径，当类型值不为int，也不为string的时候要进行报错并退出。
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query) //打印
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")
}

func createQuery(q interface{}) { //当空interface的时候，可以传入任何结构和内容。所以需要判断传入的是什么类型。
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	//reflect 包中还有一个重要的类型：Kind。
	k := t.Kind() //对实际类型名 求它真正的特定类型。
	fmt.Println("Type ", t)
	fmt.Println("Kind ", k)
	fmt.Println("Value ", v)

}

//NumField() 方法返回结构体中字段的数量，而 Field(i int) 方法返回字段 i 的 reflect.Value。

func createQuery2(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField()) //这个方法只能结构体用，查看结构体有多少个字段。
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i)) //取结构体内字段的值。
		}
	}

}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	p := "hello world"
	createQuery(o)
	createQuery(p)

	q := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery2(q)

	//Int 和 String 可以帮助我们分别取出 reflect.Value 作为 int64 和 string。
	a := 56
	x := reflect.ValueOf(a).Int() //反射取值进行类型转换为 int64
	fmt.Printf("type:%T value:%v\n", x, x)
	b := "Naveen"
	y := reflect.ValueOf(b).String() //反射取值进行类型转换为 string
	fmt.Printf("type:%T value:%v\n", y, y)

	//反射，怎么取不同类型进来 进行拼接成相关sql语句。例子。
	w := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery3(w)

	g := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery3(g)
	z := 90
	createQuery3(z) //函数只选择了Struct 进行判断。否则其他类型，打印unsupported 报错。

}

/*
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test5highlevel# go run reflection.go
Type  main.order
Kind  struct
Value  {456 56}
Type  string
Kind  string
Value  hello world
Number of fields 2
Field:0 type:reflect.Value value:456
Field:1 type:reflect.Value value:56
type:int64 value:56
type:string value:Naveen
insert into order values(456, 56)
insert into employee values("Naveen", 565, "Coimbatore", 90000, "India")
unsupported type
root@e7939faf8694:/go/src/LessonGo/class_tour/test5highlevel#

*/
