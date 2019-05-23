package main

/***
信道可以想像成 Go 协程之间通信的管道。如同管道中的水会从一端流到另一端，通过使用信道，数据也可以从一端发送，在另一端接收。
所有信道都关联了一个类型。信道只能运输这种类型的数据，而运输其他类型的数据都是非法的。
chan T 表示 T 类型的信道。
信道的零值为 nil。信道的零值没有什么用，应该像对 map 和切片所做的那样，用 make 来定义信道。

信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。
ch <- v    // 将 v 发送至信道 ch。
v := <-ch  // 从 ch 接收值并赋予 v。
箭头就是数据流的方向。
通过信道发送和接收数据。
data := <- chan // 读取信道 chan
chan <- data // 写入信道 chan

默认情况下，发送和接收操作在另一端准备好之前都会阻塞。
这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。

发送与接收默认是阻塞的。这是什么意思？
当把数据发送到信道时，程序控制会在发送数据的语句处发生阻塞，直到有其它 Go 协程从信道读取到数据，才会解除阻塞。
与此类似，当读取信道的数据时，如果没有其它的协程把数据写入到这个信道，那么读取过程就会一直阻塞着。
信道的这种特性能够帮助 Go 协程之间进行高效的通信，不需要用到其他编程语言常见的显式锁或条件变量。

简单理解 就是，发送的时候或者读取的时候，如果没有消费掉，信道就会造成阻塞，协程通过信道不管是读还是写。结束后，信道才停止阻塞，进入主进程。
一个信道可以被多个协程消费。

协程之间使用信道来进行通信。多个协程通过共享信道来共享某些数据和处理数据。
***/

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

//执行完毕后，关闭信道。
func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() { //声明
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int) //使用前创建
		//简短声明通常也是一种定义信道的简洁有效的方法。
		//a := make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	}

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c) //协程计算写入信道
	go sum(s[len(s)/2:], c) //协程计算写入信道，会占用信道，阻塞信道，知道协程结束
	x, y := <-c, <-c        // 从 c 中读取数据。 主协程，这个步骤 要等到两个协程执行完毕后，信道才会结束阻塞。不阻塞才能执行下一个语句。
	fmt.Println(x, y, x+y)

	//死锁
	//使用信道需要考虑的一个重点是死锁。
	//当 Go 协程给一个信道发送数据时，照理说会有其他 Go 协程来接收数据。如果没有的话，程序就会在运行时触发 panic，形成死锁。
	//同理，当有 Go 协程等着从一个信道接收数据时，我们期望其他的 Go 协程会向该信道写入数据，要不然程序就会触发 panic。
	//由于它超出了信道的容量，因此这次写入发生了阻塞。现在想要这次写操作能够进行下去，必须要有其它协程来读取这个信道的数据。
	/***
	信道可以是 带缓冲的。将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的信道：
	ch := make(chan int, 100) 仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。

	ch1 := make(chan int) 会发生死锁，因为信道空间占满，阻塞。

	无缓冲信道的发送和接收过程是阻塞的。
	创建一个有缓冲（Buffer）的信道。
	只在缓冲已满的情况，才会阻塞向缓冲信道（Buffered Channel）发送数据。同样，只有在缓冲为空的时候，才会阻塞从缓冲信道接收数据。
	通过向 make 函数再传递一个表示容量的参数（指定缓冲的大小），可以创建缓冲信道。
	ch := make(chan type, capacity) 无缓冲信道的容量默认为 0

	***/
	ch1 := make(chan int, 2) // 不会发生死锁，因为信道空间未满，未阻塞
	ch1 <- 5
	fmt.Println(<-ch1)

	//单向信道
	//前讨论的信道都是双向信道，即通过信道既能发送数据，又能接收数据。其实也可以创建单向信道，这种信道只能发送或者接收数据。
	//var read_test <-chan int 单项通道 只支持 读
	//var write_test chan<- int 单项通道 只支持 写
	//var chan_test chan 双项通道

	/***关闭信道 数据发送方可以关闭信道，通知接收方这个信道不再有数据发送过来。
	当从信道接收数据时，接收方可以多用一个变量来检查信道是否已经关闭。
	v, ok := <- ch
	只有发送者才能关闭信道，而接收者不能。
	信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有需要发送的值时才有必要关闭，例如终止一个 range 循环。

	***/

	ch2 := make(chan int)
	go producer(ch2) //往信道里面一直循环写入数据。 和主协程for循环并行。主会一直读数据，知道写入端终止。
	for {
		v, ok := <-ch2
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}

	//range 循环从信道 ch 接收数据，直到该信道关闭。一旦关闭了 ch，循环会自动结束。
	ch3 := make(chan int)
	go producer(ch3)

	for v := range ch3 {
		fmt.Println("Received", v)
	}

	//长度 vs 容量
	//缓冲信道的容量是指信道可以存储的值的数量。我们在使用 make 函数创建缓冲信道的时候会指定容量大小。
	//缓冲信道的长度是指信道中当前排队的元素个数。

	ch4 := make(chan string, 3)
	ch4 <- "naveen"
	ch4 <- "paul"
	fmt.Println("capacity is", cap(ch4))
	fmt.Println("length is", len(ch4))
	fmt.Println("read value", <-ch4) //读取一个就处理一下，缓冲隧道里面排队的长度就减少一个。
	fmt.Println("new length is", len(ch4))
}

/***
结果集:
root@e7939faf8694:/go/src/LessonGo/class_tour/test4concurrency# go run channels.go
channel a is nil, going to define it
Type of a is chan int
-5 17 12
5
Received  0 true
Received  1 true
Received  2 true
Received  3 true
Received  4 true
Received  5 true
Received  6 true
Received  7 true
Received  8 true
Received  9 true
Received 0
Received 1
Received 2
Received 3
Received 4
Received 5
Received 6
Received 7
Received 8
Received 9
capacity is 3
length is 2
read value naveen
new length is 1
root@e7939faf8694:/go/src/LessonGo/class_tour/test4concurrency#

***/
