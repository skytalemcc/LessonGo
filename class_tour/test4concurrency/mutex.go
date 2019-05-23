package main

/***
信道非常适合在各个 Go 协程间进行通信。
信道里面的数据可以被多个协程处理。但是如果我们不需要进行通信呢?
若我们只是想保证每次只有一个 Go 协程能够访问一个共享的变量，从而避免冲突。
这里涉及的概念叫做 *互斥（mutual*exclusion）* ，我们通常使用 *互斥锁（Mutex）* 这一数据结构来提供这种机制。
Go 标准库中提供了 sync.Mutex 互斥锁类型及其两个方法： Lock Unlock
我们可以通过在代码前调用 Lock 方法，在代码后调用 Unlock 方法来保证一段代码的互斥执行
我们也可以用 defer 语句来保证互斥锁一定会被解锁。

通过 Mutex 和信道来处理竞态条件（Race Condition）
理解并发编程中临界区（Critical Section）的概念。当程序并发地运行时，多个 Go 协程不应该同时访问那些修改共享资源的代码。这些修改共享资源的代码称为临界区。
如果在任意时刻只允许一个 Go 协程访问临界区，那么就可以避免竞态条件。而使用 Mutex 可以达到这个目的。
否则，竞态条件（Race Condition）会造成混乱，其程序的输出是由协程的执行顺序决定的。

Mutex 用于提供一种加锁机制（Locking Mechanism），可确保在某时刻只有一个协程在临界区运行，以防止出现竞态条件。
Mutex 可以在 sync 包内找到。Mutex 定义了两个方法：Lock 和 Unlock。
所有在 Lock 和 Unlock 之间的代码，都只能由一个 Go 协程执行，于是就可以避免竞态条件。

如果有一个 Go 协程已经持有了锁（Lock），当其他协程试图获得该锁时，这些协程会被阻塞，直到 Mutex 解除锁定为止。

***/

import (
	"fmt"
	"sync"
)

var x = 0 //对公共的外部变量 ，协程会使用的时候，进行加锁。给互斥锁做例子。
var y = 0 //加锁。使用信道来加锁。
func increment(i int, wg *sync.WaitGroup, m *sync.Mutex) { //解引用
	m.Lock()   //加锁
	x = x + 1  //每一个go协程 的目的 都是对x的值加1，所以哪个协程在前，无所谓，目的都是对x加1的操作。
	m.Unlock() //解锁
	fmt.Println("The work num and the x value are ", i, x)
	wg.Done() //释放

}

func increment2(wg *sync.WaitGroup, ch chan bool) {
	ch <- true //由于此信道大小为1，所以当赋值的时候 ，其他启动协程，看到此信道满了，导致处于阻塞状态，等待。
	y = y + 1  //对需要互斥的值进行处理。
	<-ch       //将信道消费掉，变为空，则其他等待的协程立马可以进行处理。
	wg.Done()  //关闭此工作池中的协程。
}

func main() {
	var w sync.WaitGroup //定义一个工作池,工作池里面可以起一堆协程来并行处理任务。
	var m sync.Mutex     // 提供互斥锁
	for i := 0; i < 10; i++ {
		w.Add(1)                //往工作池的计数器添加
		go increment(i, &w, &m) //传地址，解引用，传递的是地址而非值，这样每一个协程用的是同一个地址，而不是副本。
	}
	w.Wait()
	fmt.Println("final value of x is", x)

	//除了使用互斥锁以外，还可以使用信道处理竞态条件。
	var s sync.WaitGroup
	chy := make(chan bool, 1) //我们创建了容量为 1 的缓冲信道 ,默认为空。该缓冲信道用于保证只有一个协程访问增加 y 的临界区
	for i := 0; i < 1000; i++ {
		s.Add(1)
		go increment2(&s, chy) //启动一堆协程的工作池去读取信道。
	}
	s.Wait() //等待所有工作池计数器结束。
	fmt.Println("final value of y", y)

	/*
		互斥锁和信道：
		通过使用 Mutex 和信道，我们已经解决了竞态条件的问题。
		总体说来，当 Go 协程需要与其他协程通信时，可以使用信道。而当只允许一个协程访问临界区时，可以使用 Mutex。
		就我们上面解决的问题而言，更倾向于使用 Mutex，因为该问题并不需要协程间的通信。只是对互斥锁之间的任务加锁。
		而第二种，则是通过多个协程之间共享信道的方法，让信道阻塞来解决竞态条件的问题。
	*/
}

/*
	结果集：
	root@e7939faf8694:/go/src/LessonGo/class_tour/test4concurrency# go run mutex.go
	The work num and the x value are  9 2
	The work num and the x value are  1 1
	The work num and the x value are  0 3
	The work num and the x value are  5 4
	The work num and the x value are  2 5
	The work num and the x value are  3 6
	The work num and the x value are  4 7
	The work num and the x value are  7 8
	The work num and the x value are  6 9
	The work num and the x value are  8 10
	final value of x is 10
	final value of y 1000
	root@e7939faf8694:/go/src/LessonGo/class_tour/test4concurrency#

*/
