package main

/***
switch 是一个条件语句，用于将一个表达式的求值结果与可能的值的列表进行匹配，
并根据匹配结果执行相应的代码。可以认为 switch 语句是编写多个 if-else 子句的替代方式。
除非以 fallthrough 语句结束，否则分支会自动终止。
正常情况下 是一个个进行条件匹配工作，
***/
import (
	"fmt"
	"runtime"
)

func main() {
	//case 是进行顺序匹配，当没有匹配的case的时候 则进入default了。
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os { //os变量 只在switch存在
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")

	case "freedsb", "openbsd", "plan9": //多个值中选择的时候 ，可以在一个 case 中包含多个表达式，每个表达式用逗号分隔。
		fmt.Println("others")
	default: //这个是没有匹配的时候走的路径
		//fmt.Printf("%s.", os)
		fmt.Println(os)
	}

	//没有表达式的switch ，则相当于 switch true 这种情况下会将每一个 case 的表达式的求值结果与 true 做比较，如果相等，则执行相应的代码
	//这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。
	num := 75
	switch {
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")

		//在 Go 中执行完一个 case 之后会立即退出 switch 语句。fallthrough语句用于标明执行完当前 case 语句之后按顺序执行下一个case 语句。
		fallthrough
	case num >= 101:
		fmt.Println("num is greater than 100")
	}

}

/***
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic# go run switch.go
Go runs on Linux.
num is greater than 51 and less than 100
num is greater than 100
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic#

***/
