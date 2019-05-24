package main

/***
Go语言中 只有一种循环结构 for 循环
基本的 for 循环除了没有了 `( )` 之外 而 `{ }` 是必须的。 没有小括号，大括号是必须的。
循环语句是用来重复执行某一段代码。
for 是 Go 语言唯一的循环语句。Go 语言中并没有其他语言比如 C 语言中的 while 和 do while 循环。
语法：
for initialisation; condition; post {
}

***/
import (
	"fmt"
	"time"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ { //常见的循环
		sum += i
		if i == 5 {
			break //break语句用于在完成正常执行之前突然终止 for 循环，之后程序将会在 for 循环下一行代码开始执行。
		}

		if i == 3 {
			//continue语句用来跳出 for 循环中当前循环。在 continue 语句后的所有的 for 循环语句都不会在本次循环中执行。循环体会在一下次循环中继续执行。
			continue
			fmt.Println("will not show forever")
		}

	}
	fmt.Println(sum)
	//golang没有while关键字，但可以让前置后置语句为空，就像其他语言的while一样.基于此可以省略分号：C 的 while 在 Go 中叫做 `for`。
	for sum < 1000 { //这个格式的 for 循环可以看作是 for while 循环。
		sum += sum
	}
	fmt.Println(sum)
	//怎么来使用死循环 如果省略了循环条件，循环就不会结束，因此可以用更简洁地形式表达死循环。
	for {
		fmt.Println("print for ever")               //会一直打印下去
		time.Sleep(time.Duration(20) * time.Second) //20秒sleep时间
	}

}

/***
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic# go run for.go
15
1920
print for ever
print for ever
print for ever
***/
