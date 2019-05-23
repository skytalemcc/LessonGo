package main

/***
工作池(worker pools) ，WaitGroup用于实现工作池。
WaitGroup 用于等待一批 Go 协程执行结束。程序控制会一直阻塞，直到这些协程全部执行完毕。
假设我们有 3 个并发执行的 Go 协程（由 Go 主协程生成）。Go 主协程需要等待这 3 个协程执行结束后，才会终止。这就可以用 WaitGroup 来实现。

工作池的实现：
缓冲信道重要应用之一就是实现工作池。
一般而言，工作池就是一组等待任务分配的线程。一旦完成了所分配的任务，这些线程可继续等待任务的分配。
就是一些线程 专门用来处理协程的任务。

首先说明下面这个例子
定义了100个jobs 和10个works。
一个写入chan 一个读取chan
allocate协程 负责写入100个任务到写入chan
createWorkerPool用来启动10个worker 来处理写入chan，将数据传给读取chan
result协程 负责将读取chan的数据消费掉

共12个协程同时并行着，写入数据到写入chan，拷贝数据从写入chan到读取chan，打印消费读取chan的内容。
因为是并行的，所以打印会打印不同协程的内容

***/
import (
	"fmt"
	//"reflect"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) { //声明wg 是waitgroup类型的指针
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() //要减少计数器，可以调用 WaitGroup 的 Done() 方法。 解引用。
}

var jobs = make(chan int, 10)    //定义写入chan 来接收作业 ，生产者
var results = make(chan int, 10) //定义读取chan 来处理作业，消费者

func allocate(noOfJobs int) { //此方法 用来将任务添加到jobs信道
	for i := 0; i < noOfJobs; i++ {
		jobs <- i
	}
	close(jobs) //添加完毕后，关闭信道
}

func createWorkerPool(noOfWorkers int) { //根据传参 并行开启n个worker 作为工作池
	var wg sync.WaitGroup //需要对worker的状态进行监控
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg) //循环开启worker

		fmt.Println("open the worker ", i)
	}
	wg.Wait()      //等待worker全部关闭后才能执行下一步 ，wg 要等待到wg的计数器为0 才行。所以worker里面要有计数器减1。
	close(results) //关闭接收chan
	fmt.Println("close the pool")
}

func worker(i int, wg *sync.WaitGroup) { //协程将写入chan的数据 传给读取chan
	for job := range jobs { //会不断从信道接收值，直到它被关闭。
		time.Sleep(2 * time.Second)
		results <- job //将写入信道数据传给 读取信道
		//fmt.Println( reflect.TypeOf(i), reflect.TypeOf(job)) 用来显示变量的类型
		fmt.Printf("worker %d send %d to chan result \n", i, job)
	}
	wg.Done() //等待jobs chan关闭后，worker的for循环会结束进入下一步到这一行，这一行会将wg工作池计数器减1。

	fmt.Println("close the worker ", i) //打印关闭worker
}

func result(done chan bool) { //对读取chan 进行消费来打印数据
	for result := range results {
		fmt.Println("The result is ", result)
	} //当读取chan result 关闭后，range循环才会结束进入下一行。
	done <- true //将成功标记赋给done
}

func main() {
	no := 3
	var wg sync.WaitGroup     //创建了 WaitGroup 类型的变量，其初始值为零值 。WaitGroup 使用计数器来工作。
	for i := 0; i < no; i++ { //主协程 循环三次 生成三个并发的协程。
		wg.Add(1) //当我们调用 WaitGroup 的 Add 并传递一个 int 时，WaitGroup 的计数器会加上 Add 的传参。

		//传递 wg 的地址给协程，是很重要的。如果没有传递 wg 的地址，那么每个 Go 协程将会得到一个 WaitGroup 值的拷贝。
		//因而当它们执行结束时，main 函数并不会知道。
		go process(i, &wg) //传递指针地址
	}
	wg.Wait() //Wait() 方法会阻塞调用它的 Go主协程，直到计数器变为 0 后才会停止阻塞。然后跳到下一行。
	fmt.Println("All go routines finished executing")

	/*工作池的核心功能如下：
	创建一个 Go 协程池，监听一个等待作业分配的输入型缓冲信道。
	将作业添加到该输入型缓冲信道中。
	作业完成后，再将结果写入一个输出型缓冲信道。
	从输出型缓冲信道读取并打印结果。
	*/

	startTime := time.Now()
	noOfJobs := 100   //定义任务的数目为100
	noOfWorkers := 10 //定义消费的worker为10

	go allocate(noOfJobs)   //并行。将100个任务 传入到大小为10的写入信道jobs
	done := make(chan bool) //done默认值为false
	//并行。监控读取chan results ，大小为10,100个任务不断进入，不断被消费掉，打印结果。直到result chan关闭。
	//go result 函数会去消费results ,worker拷贝工作完毕后，会关闭results chan。 关闭后，results chan消费完毕后，会将true传给done。
	go result(done)
	createWorkerPool(noOfWorkers) //启动工作池函数，创建10个worker来去消费，直到jobs chan先关闭，results chan后关闭。
	<-done                        //消费掉这个标记chan的内容
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
	fmt.Println(startTime, endTime)
}

/***
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test4concurrency# go run waitgroup.go
started Goroutine  0
started Goroutine  2
started Goroutine  1
Goroutine 2 ended
Goroutine 1 ended
Goroutine 0 ended
All go routines finished executing
open the worker  0
open the worker  1
open the worker  2
open the worker  3
open the worker  4
open the worker  5
open the worker  6
open the worker  7
open the worker  8
open the worker  9
worker 8 send 9 to chan result
worker 9 send 0 to chan result
worker 0 send 1 to chan result
worker 1 send 2 to chan result
worker 2 send 3 to chan result
worker 3 send 4 to chan result
worker 4 send 5 to chan result
worker 5 send 6 to chan result
worker 6 send 7 to chan result
worker 7 send 8 to chan result
The result is  9
The result is  0
The result is  1
The result is  2
The result is  3
The result is  4
The result is  5
The result is  6
The result is  7
The result is  8
worker 7 send 19 to chan result
worker 8 send 10 to chan result
worker 9 send 11 to chan result
worker 0 send 12 to chan result
worker 1 send 13 to chan result
worker 2 send 14 to chan result
worker 3 send 15 to chan result
worker 4 send 16 to chan result
worker 5 send 17 to chan result
worker 6 send 18 to chan result
The result is  19
The result is  10
The result is  11
The result is  12
The result is  13
The result is  14
The result is  15
The result is  16
The result is  17
The result is  18
worker 6 send 29 to chan result
worker 2 send 25 to chan result
worker 3 send 26 to chan result
worker 4 send 27 to chan result
worker 5 send 28 to chan result
The result is  29
The result is  25
The result is  26
The result is  27
The result is  28
worker 9 send 22 to chan result
worker 7 send 20 to chan result
worker 8 send 21 to chan result
The result is  22
The result is  20
The result is  21
worker 0 send 23 to chan result
The result is  23
worker 1 send 24 to chan result
The result is  24
worker 1 send 39 to chan result
worker 6 send 30 to chan result
worker 3 send 32 to chan result
worker 4 send 33 to chan result
worker 5 send 34 to chan result
worker 9 send 35 to chan result
worker 7 send 36 to chan result
worker 8 send 37 to chan result
worker 0 send 38 to chan result
The result is  39
The result is  30
The result is  32
The result is  33
The result is  34
The result is  35
The result is  36
The result is  37
The result is  38
worker 2 send 31 to chan result
The result is  31
worker 2 send 49 to chan result
worker 3 send 42 to chan result
The result is  49
The result is  48
The result is  40
The result is  41
The result is  42
worker 6 send 41 to chan result
worker 0 send 48 to chan result
worker 9 send 45 to chan result
worker 4 send 43 to chan result
worker 5 send 44 to chan result
The result is  45
The result is  46
The result is  43
The result is  44
worker 8 send 47 to chan result
The result is  47
worker 7 send 46 to chan result
worker 1 send 40 to chan result
worker 8 send 57 to chan result
worker 1 send 59 to chan result
worker 3 send 51 to chan result
worker 7 send 58 to chan result
worker 0 send 53 to chan result
worker 2 send 50 to chan result
worker 6 send 52 to chan result
worker 4 send 55 to chan result
worker 9 send 54 to chan result
worker 5 send 56 to chan result
The result is  57
The result is  59
The result is  51
The result is  58
The result is  53
The result is  50
The result is  52
The result is  55
The result is  54
The result is  56
worker 8 send 60 to chan result
The result is  60
worker 5 send 69 to chan result
worker 2 send 65 to chan result
worker 1 send 61 to chan result
worker 3 send 62 to chan result
worker 4 send 67 to chan result
worker 9 send 68 to chan result
The result is  69
The result is  64
The result is  65
The result is  61
The result is  66
The result is  67
The result is  68
The result is  62
The result is  63
worker 7 send 63 to chan result
worker 6 send 66 to chan result
worker 0 send 64 to chan result
worker 0 send 79 to chan result
worker 8 send 70 to chan result
worker 5 send 71 to chan result
worker 2 send 72 to chan result
worker 3 send 74 to chan result
worker 4 send 75 to chan result
worker 6 send 78 to chan result
worker 1 send 73 to chan result
worker 9 send 76 to chan result
worker 7 send 77 to chan result
The result is  79
The result is  70
The result is  71
The result is  72
The result is  74
The result is  75
The result is  78
The result is  73
The result is  76
The result is  77
worker 7 send 89 to chan result
worker 6 send 86 to chan result
worker 0 send 80 to chan result
worker 8 send 81 to chan result
worker 5 send 82 to chan result
worker 2 send 83 to chan result
worker 3 send 84 to chan result
worker 4 send 85 to chan result
worker 1 send 87 to chan result
worker 9 send 88 to chan result
The result is  89
The result is  80
The result is  86
The result is  81
The result is  82
The result is  83
The result is  84
The result is  85
The result is  87
The result is  88
worker 6 send 91 to chan result
worker 2 send 95 to chan result
close the worker  2
worker 0 send 92 to chan result
close the worker  0
worker 5 send 94 to chan result
close the worker  5
worker 9 send 99 to chan result
close the worker  9
worker 4 send 97 to chan result
close the worker  4
worker 1 send 98 to chan result
close the worker  1
worker 7 send 90 to chan result
The result is  91
The result is  99
The result is  90
The result is  92
The result is  93
The result is  94
The result is  95
The result is  97
The result is  96
The result is  98
worker 8 send 93 to chan result
close the worker  8
close the worker  7
close the worker  6
worker 3 send 96 to chan result
close the worker  3
close the pool
total time taken  20.0053994 seconds
2019-05-23 10:05:26.4764284 +0000 UTC m=+2.001164801 2019-05-23 10:05:46.4817751 +0000 UTC m=+22.006564201
root@e7939faf8694:/go/src/LessonGo/class_tour/test4concurrency#

***/
