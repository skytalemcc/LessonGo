package main

/***
在 Go 语言中，程序中一般是使用错误来处理异常情况。对于程序中出现的大部分异常情况，错误就已经够用了。
但在有些情况，当程序发生异常时，无法继续运行。在这种情况下，我们会使用 panic 来终止程序。
当函数发生 panic 时，它会终止运行，在执行完所有的延迟函数后，程序控制返回到该函数的调用方。
这样的过程会一直持续下去，直到当前协程的所有函数都返回退出，然后程序会打印出 panic 信息，接着打印出堆栈跟踪（Stack Trace），最后程序终止。

当程序发生 panic 时，使用 recover 可以重新获得对该程序的控制。

什么时候使用panic
你应该尽可能地使用错误，而不是使用 panic 和 recover。只有当程序不能继续运行的时候，才应该使用 panic 和 recover 机制。

panic 有两个合理的用例:
1，发生了一个不能恢复的错误，此时程序不能继续运行。 例如web端口被占用，无法绑定的时候，要使用panic 直接退出。
2，发生了一个编程上的错误。假如我们有一个接收指针参数的方法，而其他人使用 nil 作为参数调用了它。

recover 不能恢复一个不同协程的 panic。,只有在相同的 Go 协程中调用 recover 才管用。



***/

import (
	"fmt"
	"runtime/debug"
)

//recover 是一个内建函数，用于重新获得 panic 协程的控制。它是一个接口。
//只有在延迟函数的内部，调用 recover 才有用。在延迟函数内调用 recover，可以取到 panic 的错误信息，
//并且停止 panic 续发事件（Panicking Sequence），程序运行恢复正常。
//如果在延迟函数的外部调用 recover，就不能停止 panic 续发事件。
func recoverName() { //不使用它的时候 ，panic会退出程序，同时还续发一堆错误打印。 使用recover的时候，panic会退出程序，续发的错误打印停止。
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
		//当我们恢复 panic 时，我们就释放了它的堆栈跟踪。实际上，在上述程序里，恢复 panic 之后，我们就失去了堆栈跟踪。
		//有办法可以打印出堆栈跟踪，就是使用 Debug 包中的 PrintStack 函数。
		debug.PrintStack()
	}

}

func main() {

	//发生 panic 时的 defer
	//当函数发生 panic 时，它会终止运行，在执行完所有的延迟函数后，程序控制返回到该函数的调用方。
	//这样的过程会一直持续下去，直到当前协程的所有函数都返回退出，然后程序会打印出 panic 信息，接着打印出堆栈跟踪，最后程序终止。
	defer recoverName()
	defer fmt.Println("it will be shown ,then panic exec ")

	var econtent error
	if econtent == nil {
		panic("runtime error: econtent  cannot be nil")
	}
	fmt.Println("it will not be shown ") //panic后面的不会执行

}

/***
结果集:
root@e7939faf8694:/go/src/LessonGo/class_tour/test3interface# go run error2.go
it will be shown ,then panic exec
recovered from  runtime error: econtent  cannot be nil
goroutine 1 [running]:      panic后续堆栈跟踪错误信息 使用recover不打印了。 使用debug.PrintStack() 方法来调试打印出来。
runtime/debug.Stack(0x37, 0x0, 0x0)
        /usr/local/go/src/runtime/debug/stack.go:24 +0x9d
runtime/debug.PrintStack()
        /usr/local/go/src/runtime/debug/stack.go:16 +0x22
main.recoverName()
        /go/src/LessonGo/class_tour/test3interface/error2.go:38 +0xb5
panic(0x498460, 0x4cfac0)
        /usr/local/go/src/runtime/panic.go:522 +0x1b5
main.main()
        /go/src/LessonGo/class_tour/test3interface/error2.go:53 +0xb5
root@e7939faf8694:/go/src/LessonGo/class_tour/test3interface#

***/
