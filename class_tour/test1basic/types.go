package main

/***
golang的基本类型 ：
bool
string
int int8  int16  int32  int64 有符号整型
uint uint8 uint16 uint32 uint64 uintptr  无符号整型
byte // uint8 的别名
rune // int32 的别名 代表一个Unicode码    byte与rune都属于别名类型。byte是uint8的别名类型，而rune是int32的别名类型。
float32 float64 浮点型
complex64 complex128 复数类型

***/
import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14 //常量的定义与变量类似，只不过使用 const 关键字。 常量可以是字符、字符串、布尔或数字类型的值。 常量不能使用 := 语法定义。

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
	//变量在定义时没有明确的初始化时会赋值为_零值_。零值是：数值类型为 `0`，布尔类型为 `false`，字符串为 `""`（空字符串）。
	var i int
	var l float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, l, b, s)

	//类型转换 Go 的在不同类型之间的项目赋值时需要显式转换
	var z int = 42
	var x float64 = float64(z) //需要使用float64(z)来代表显示转换，不带是不能转换的
	fmt.Println(z, x)
	//类型推导 在定义一个变量但不指定其类型时（使用没有类型的 var 或 := 语句）， 变量的类型由右值推导得出。
	var g int
	v := g
	fmt.Printf("%T,%T\n", g, v)
	fmt.Println(Pi)
}
