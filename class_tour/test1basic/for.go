package main

/***
Go语言中 只有一种循环结构 for 循环
基本的 for 循环除了没有了 `( )` 之外 而 `{ }` 是必须的。 没有小括号，大括号是必须的。
***/
import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ { //常见的循环
		sum += i
	}
	fmt.Println(sum)
	//golang没有while关键字，但可以让前置后置语句为空，就像其他语言的while一样.基于此可以省略分号：C 的 while 在 Go 中叫做 `for`。
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	//怎么来使用死循环 如果省略了循环条件，循环就不会结束，因此可以用更简洁地形式表达死循环。
	for {
		fmt.Println("print for ever") //会一直打印下去
	}

}
