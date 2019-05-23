package main

/***
并发入门：
Go 是并发式语言，而不是并行式语言。
并发是指立即处理多个任务的能力。例如处理一个任务停下来处理别的任务，处理完毕再回来处理之前的任务。
并行是指同时处理多个任务。例如处理一个任务，同时处理另外一个任务。一般指的是多核。
多核的并行系统上，并行运行的组件之间可能需要相互通信。所以，并行不一定会加快运行速度。

Go 编程语言原生支持并发。Go 使用 Go 协程（Goroutine） 和信道（Channel）来处理并发。
goroutines go协程
Go 协程是与其他函数或方法一起并发运行的函数或方法。Go 协程可以看作是轻量级线程。
Go 协程相比于线程的优势：
1，相比线程而言，Go 协程的成本极低。堆栈大小只有若干 kb，并且可以根据应用的需求进行增减。而线程必须指定堆栈的大小，其堆栈是固定不变的。
2，Go 协程会复用（Multiplex）数量更少的 OS 线程。即使程序有数以千计的 Go 协程，也可能只有一个线程。
如果该线程中的某一 Go 协程发生了阻塞（比如说等待用户输入），那么系统会再创建一个 OS 线程，并把其余 Go 协程都移动到这个新的 OS 线程。
所有这一切都在运行时进行。
3，Go 协程使用信道（Channel）来进行通信。信道用于防止多个协程访问共享内存时发生竞态条件（Race Condition）。
信道可以看作是 Go 协程之间通信的管道。

每个go 命令的都是子协程 ，而主函数则为主协程。可以这么称呼。并行跑任务，即无感知下切换处理任务。
***/

import (
	"fmt"
	"time"
)

func hello() {

	fmt.Println("This is hello world go routine")
}

//启动多个 Go 协程
func hello2() {

	fmt.Println("This is hello world another go routine")
}

func main() {
	// 主函数和协程会并发执行，但是主函数不会等待协程是否完成。
	//启动一个新的协程时，协程的调用会立即返回。与函数不同，程序控制不会去等待 Go 协程执行完毕。
	//在调用 Go 协程之后，程序控制会立即返回到代码的下一行，忽略该协程的任何返回值。
	//如果希望运行其他 Go 协程，Go 主协程必须继续运行着。如果 Go 主协程终止，则程序终止，于是其他 Go 协程也不会继续运行。
	go hello()
	go hello2()
	//在 Go 主协程中使用休眠，以便等待其他协程执行完毕，这种方法只是用于理解 Go 协程如何工作的技巧。信道可用于在其他协程结束执行之前，阻塞 Go 主协程。
	time.Sleep(2 * time.Second)
	fmt.Println("This is the hello main function")

}

/***
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test4concurrency# go run goroutines.go
This is hello world go routine
This is hello world another go routine
This is the hello main function
root@e7939faf8694:/go/src/LessonGo/class_tour/test4concurrency#

***/
