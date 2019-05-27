package main

/***
golang的基本类型 ：
bool 表示一个布尔值，值为 true 或者 false。
string 在 Golang 中，字符串是字节的集合
int int8  int16  int32  int64 有符号整型
uint uint8 uint16 uint32 uint64 uintptr  无符号整型
byte // uint8 的别名
rune // int32 的别名 代表一个Unicode码    byte与rune都属于别名类型。byte是uint8的别名类型，而rune是int32的别名类型。
float32 float64 浮点型
complex64 complex128 复数类型
Go 有着非常严格的强类型特征。Go 没有自动类型提升或类型转换
***/

/***
什么是字符串？
Go 语言中的字符串是一个字节切片。把内容放在双引号""之间，我们可以创建一个字符串。
由于字符串是一个字节切片，所以我们可以获取字符串的每一个字节。 使用for 循环遍历字符串。

rune 是 Go 语言的内建类型，它也是 int32 的别称。在 Go 语言中，rune 表示一个代码点。代码点无论占用多少个字节，都可以用一个 rune 来表示。

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

//常量的定义与变量类似，只不过使用 const 关键字。 常量可以是字符、字符串、布尔或数字类型的值。 常量不能使用 := 语法定义。
//常量的值会在编译的时候确定。因为函数调用发生在运行时，所以不能将函数的返回值赋值给常量。
const Pi = 3.14

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	runes := []rune(s) //rune 表示一个代码点
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

//修改字符串中的字符  为了修改字符串，可以把字符串转化为一个 rune 切片。然后这个切片可以进行任何想要的改变，然后再转化为一个字符串。
func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

/*
这种是无法修改字符串的。
func mutate(s string)string {
    s[0] = 'a'//any valid unicode character within single quote is a rune
    return s
}
*/

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

	name := "Hello World"
	printBytes(name) //打印字节码，ASCII码
	fmt.Println("")
	printChars(name) //打印字符 ，当需要遍历字符串的时候 使用rune来打印字符.
	fmt.Println("")
	//使用for range来遍历字符串更好
	for index, rune := range name {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}

	//用字节切片构造字符串
	byteSlice := []byte{67, 97, 102, 195, 169} //10 进制值
	//byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9} 16 进制换
	str := string(byteSlice)
	fmt.Println(str)

	//用 rune 切片构造字符串
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072} //16 进制的 Unicode 代码点
	str2 := string(runeSlice)
	fmt.Println(str2)

	//字符串的内容可以用类似于数组下标的方式获取，但与数组不同，字符串的内容不能在初始化后被修改。
	//修改字符串中的字符，使用rune
	h := "hello"
	fmt.Println(mutate([]rune(h)))

	//另一个创建变量的方法是调用用内建的new函数。表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为*T。
	m := new(int) //m实际上为指针，它的值为int初始值0
	//用new创建变量和普通变量声明语句方式创建变量没有什么区别,除了不需要声明一个临时变量的名字外，我们还可以在表达式中使用new(T)。
	//换言之，new函数类似是一种语法糖，而不是一个新的基础概念。每次调用new函数都是返回一个新的变量的地址。
	//new函数使用通常相对比较少，因为对于结构体来说，直接用字面量语法创建新变量的方法会更灵活
	//由于new只是一个预定义的函数，它并不是一个关键字，因此我们可以将new名字重新定义为别的类型。
	fmt.Println(m, *m)

}

/***
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic# go run types.go
bool(false)
uint64(18446744073709551615)
complex128((2+3i))
0 0 false ""
42 42
int,int
3.14
48 65 6c 6c 6f 20 57 6f 72 6c 64
H e l l o   W o r l d
H starts at byte 0
e starts at byte 1
l starts at byte 2
l starts at byte 3
o starts at byte 4
  starts at byte 5
W starts at byte 6
o starts at byte 7
r starts at byte 8
l starts at byte 9
d starts at byte 10
Café
Señor
aello
0xc000074158 0
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic#

***/
