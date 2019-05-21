package main

/***
defer是golang的一个特色功能，被称为“延迟调用函数”。当外部函数返回后执行defer。
defer 语句会延迟函数的执行直到上层函数返回。
延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。

注意点
当defer被声明时，其参数就会被实时解析
defer执行顺序为先进后出
defer可以读取有名返回值

defer通常用来释放函数内部变量。释放文件打开。捕捉返回异常。输出日志等收尾工作。

***/
import "fmt"

//defer栈
//延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。

func defer_mutil() {
	fmt.Println("counting:")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

//简单defer和复杂defer。遵循先进后出原则
func main() {
	defer_mutil()              //函数最先执行，里面的defer 按先进后出执行。
	defer fmt.Println("world") //在最下面的defer后出。
	fmt.Println("hello")
}

/***
结果集:
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic# go run defer.go
counting:
9
8
7
6
5
4
3
2
1
0
hello
world
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic#


***/
