package main

/*
fmt包实现了类似C语言printf和scanf的格式化I/O。
通用：
%d          十进制整数
%x, %o, %b  十六进制，八进制，二进制整数。
%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
%t          布尔：true或false
%c          字符（rune） (Unicode码点) 该值对应的unicode码值
%s          字符串
%q          带双引号的字符串"abc"或带单引号的字符'c'
%v          变量的自然形式（natural format）
%T          变量的类型
%%          字面上的百分号标志（无操作数）
%U			表示为Unicode格式
%p			表示指针，十六进制

宽度和精度：
宽度和精度格式化控制的是Unicode码值的数量。
对于大多数类型的值，宽度是输出字符数目的最小数量，如果必要会用空格填充。对于字符串，精度是输出字符数目的最大数量，如果必要会截断字符串。
对于整数，宽度和精度都设置输出总长度。采用精度时表示右对齐并用0填充，而宽度默认表示用空格填充。
%f:    默认宽度，默认精度
%9f    宽度9，默认精度
%.2f   默认宽度，精度2
%9.2f  宽度9，精度2
%9.f   宽度9，精度0

输入

一系列类似的函数读取格式化的文本，生成值。Scan，Scanf和Scanln从os.Stdin读取；Fscan，Fscanf和Fscanln 从特定的io.Reader读取；Sscan，Sscanf和Sscanln 从字符串读取；
Scanln，Fscanln和Sscanln在换行时结束读取，并要求数据连续出现；
Scanf，Fscanf和Sscanf会读取一整行以匹配格式字符串；其他的函数将换行看着空格。

Scan、Scanf和Scanln从标准输入os.Stdin读取文本；Fscan、Fscanf、Fscanln从指定的io.Reader接口读取文本；Sscan、Sscanf、Sscanln从一个参数字符串读取文本。
Scanln、Fscanln、Sscanln会在读取到换行时停止，并要求一次提供一行所有条目；Scanf、Fscanf、Sscanf只有在格式化文本末端有换行时会读取到换行为止；其他函数会将换行视为空白。
Scanf、Fscanf、Sscanf会根据格式字符串解析参数，类似Printf。例如%x会读取一个十六进制的整数，%v会按对应值的默认格式读取。
*/
import (
	"fmt"
	"os"
)

func main() {

	// Print 将参数列表 a 中的各个参数转换为字符串并写入到标准输出中。
	// 非字符串参数之间会添加空格，返回写入的字节数。
	fmt.Print("a", "b", 1, 2, 3, "c", "d", "\n")
	// Println 功能类似 Print，只不过最后会添加一个换行符。
	// 所有参数之间会添加空格，返回写入的字节数。
	fmt.Println("a", "b", 1, 2, 3, "c", "d")
	// Printf 将参数列表 a 填写到格式字符串 format 的占位符中。
	// 填写后的结果写入到标准输出中，返回写入的字节数。
	fmt.Printf("ab %d %d %d cd\n", 1, 2, 3)

	// 功能同上面三个函数，只不过将转换结果写入到 w 中。 一般通过这个来赋值给io.Writer
	//func Fprint(w io.Writer, a ...interface{}) (n int, err error)
	//func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	//func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	a := "asdf"
	fmt.Fprintln(os.Stdout, a)
	fmt.Fprintf(os.Stdout, "%.2s\n", a)
	fmt.Fprint(os.Stdout, a)

	// 功能同上面三个函数，只不过将转换结果以字符串形式返回。
	//func Sprint(a ...interface{}) string
	//func Sprintln(a ...interface{}) string
	//func Sprintf(format string, a ...interface{}) string
	var progress = 2
	var target = 8
	title := fmt.Sprintf("已采集%d个药草, 还需要%d个完成任务", progress, target) //给出返回值
	fmt.Println(title)
	// 功能同 Sprintf，只不过结果字符串被包装成了 error 类型。
	//func Errorf(format string, a ...interface{}) error
	err := fmt.Errorf("数值 %d 超出范围（100）", 101)
	fmt.Println(err)

	/* Scan 从标准输入中读取数据，并将数据用空白分割并解析后存入 a 提供
	   的变量中（换行符会被当作空白处理），变量必须以指针传入。
	   当读到 EOF 或所有变量都填写完毕则停止扫描。
	   返回成功解析的参数数量。
	   func Scan(a ...interface{}) (n int, err error)

	   Scanln 和 Scan 类似，只不过遇到换行符就停止扫描。
	   func Scanln(a ...interface{}) (n int, err error)

	   Scanf 从标准输入中读取数据，并根据格式字符串 format 对数据进行解析，
	   将解析结果存入参数 a 所提供的变量中，变量必须以指针传入。
	   输入端的换行符必须和 format 中的换行符相对应（如果格式字符串中有换行
	   符，则输入端必须输入相应的换行符）。
	   占位符 %c 总是匹配下一个字符，包括空白，比如空格符、制表符、换行符。
	   返回成功解析的参数数量。
	   func Scanf(format string, a ...interface{}) (n int, err error)

	   功能同上面三个函数，只不过从 r 中读取数据。

	   Fscan从r扫描文本，将成功读取的空白分隔的值保存进成功传递给本函数的参数。换行视为空白。
	   返回成功扫描的条目个数和遇到的任何错误。如果读取的条目比提供的参数少，会返回一个错误报告原因。
	   func Fscan(r io.Reader, a ...interface{}) (n int, err error)

	   Fscanln类似Fscan，但会在换行时才停止扫描。
	   func Fscanln(r io.Reader, a ...interface{}) (n int, err error)

	   Fscanf从r扫描文本，根据format 参数指定的格式将成功读取的空白分隔的值保存进成功传递给本函数的参数。

	   func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)

	   功能同上面三个函数，只不过从 str 中读取数据。
	   func Sscan(str string, a ...interface{}) (n int, err error)
	   func Sscanln(str string, a ...interface{}) (n int, err error)
	   func Sscanf(str string, format string, a ...interface{}) (n int, err error)
	*/

	//三种输入方式 ：从控制台输入 Scanln 从IO输入 Fscanln 从字符串输入 Sscanln
	/*
		a, b, c := "", 0, false
		fmt.Scan(&a, &b, &c)  //扫描控制台输入 传完才算结束
		fmt.Println(a, b, c)
		// 在终端执行后，输入 abc 1 回车 true 回车
		// 结果 abc 1 true
	*/
	a2, b2, c2 := "", 0, false
	fmt.Scanln(&a2, &b2, &c2) //扫描控制台输入 按回车就结束传参
	fmt.Println(a2, b2, c2)
	// 在终端执行后，输入 abc 1 true 回车
	// 结果 abc 1 true

	a3, b3, c3 := "", 0, false
	fmt.Scanf("%4s%d%t", &a3, &b3, &c3) // scan format格式字符串可以指定宽度
	fmt.Println(a3, b3, c3)
	// 在终端执行后，输入 1234567true 回车
	// 结果 1234 567 true

	//注：Sscanf有固定格式去进行分割读取数值，而Sscan和Sscanln靠空格进行分割进行值存储．
	//Sscan 这三个 这些代表从字符串输入传参，按照指定格式截取 然后传参给参数
	a4, b4, c4 := "", 0, 0
	fmt.Sscan("hello 1", &a4, &b4) //hello 1   Sscan 扫描实参 string，并将连续由空格分隔的值存储为连续的实参。换行符计为空格。
	fmt.Println(a4, b4)
	fmt.Sscanf("helloworld 2 ", "hello%s%d", &a4, &c4) //world 2  Scanf 扫描实参 string，并将连续由空格分隔的值存储为连续的实参，其格式由 format 决定。
	fmt.Println(a4, c4)
}

/*
结果集:
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# go run fmt.go
ab1 2 3cd
a b 1 2 3 c d
ab 1 2 3 cd
asdf
as
asdf已采集2个药草, 还需要8个完成任务
数值 101 超出范围（100）
abc 1 true
abc 1 true
1234567true
1234 567 true
hello 1
world 2
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage#
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage#

*/
