package main

/***
select 语句使一个协程可以等待多个通信操作。
select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。
select 语句用于在多个发送/接收信道操作中进行选择。select 语句会一直阻塞，直到发送/接收操作准备就绪。
如果有多个信道操作准备完毕，select 会随机地选取其中之一执行。该语法与 switch 类似，所不同的是，这里的每个 case 语句都是信道操作。

谁先符合条件，先选择执行谁。其他的则抛弃掉。select case只作用于信道上，对信道的值进行选择。
应用的场景：例如，多路选择，超时，选择执行。

标准格式:带for 就一直循环，每次都检查信道的值是否符合case。不带for，则一次性执行。
for{
	select {
		case i := <-c:
    	// 使用 i
		default:
    	// 从 c 中接收会阻塞时执行
	}
}

***/

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}

func process(ch chan string) {
	time.Sleep(10000 * time.Millisecond) //沉睡10秒
	ch <- "process successful"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)

	//两个协程 谁先符合条件，选择执行谁。选择执行。
	select {
	case s1 := <-output1: //6秒钟后才有值。
		fmt.Println(s1)
	case s2 := <-output2: //3秒钟就有值，先执行。
		fmt.Println(s2)
	}

	//当 select 中的其它分支都没有准备好时，default 分支就会执行。
	//为了在尝试发送或者接收时不发生阻塞，可使用 default 分支。
	ch := make(chan string)
	go process(ch) //等待10秒，信道有值写入
	for {
		time.Sleep(1000 * time.Millisecond) //1秒后进入循环
		select {
		case v := <-ch: //检查信道是否有值，无，跳过
			fmt.Println("received value: ", v)
			return
		default: // 在没有 case 准备就绪时，可以执行 select 语句中的默认情况（Default Case）。这通常用于防止 select 语句一直阻塞。
			fmt.Println("no value received")
		}
	}

	/*死锁，一般是读取信道的时候，信道并没有数据，select会一直阻塞，导致死锁。触发panic。
			ch := make(chan string)
	    	select {
			case <-ch:

			default: //如果存在默认情况，就不会发生死锁，因为在没有其他 case 准备就绪时，会执行默认情况。
	        fmt.Println("default case executed")

	    	}

	*/
	/*
		当 select 由多个 case 准备就绪时，将会随机地选取其中之一去执行。
		select {}  如果只有空select，由于没有case执行，select语句一直会阻塞，造成死锁。

	*/
}

/***
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test4concurrency# go run select.go
from server2
no value received
no value received
no value received
no value received
no value received
no value received
no value received
no value received
no value received
received value:  process successful
root@e7939faf8694:/go/src/LessonGo/class_tour/test4concurrency#

***/
