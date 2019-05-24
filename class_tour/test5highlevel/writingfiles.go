package main

/*
使用 Go 语言将数据写到文件里面。并且还要学习如何同步的写到文件里面。
将字符串写入文件。
将字节写入文件。
将数据一行一行的写入文件。
追加到文件里。
并发写文件。
*/

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func filewriter(i int, wg *sync.WaitGroup, f3 *os.File) { //协程，每个协程对文件进行写入动作

	newLine := "Start worker  " + strconv.Itoa(i)
	fmt.Println("Start worker ", i)
	_, err := fmt.Fprintln(f3, newLine) //将内容写入文件
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Close the worker ", i)
	newLine = "Close the worker  " + strconv.Itoa(i)
	_, err = fmt.Fprintln(f3, newLine) //将关闭动作写入文件
	if err != nil {
		fmt.Println(err)
	}
	wg.Done() //关闭协程

}

func main() {

	f, err := os.Create("test.txt") //os.Create用来创建文件，mode为0666。如果有此文件 置空。无路径，则当前文件夹进行创建。
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Type of f is %T\n", f) //f为*os.File  文件类型

	//将字符串写入文件
	//最常见的写文件就是将字符串写入文件。1，创建文件 2，将字符串写入文件
	l, err := f.WriteString("Hello World\n") //os.File.WriteString(专门处理string) 向file文件类型写入字符串 。l返回字符串的字节数
	if err != nil {
		fmt.Println(err)
		f.Close() //写入报错，然后打印错误，并关闭文件。
		return
	}
	fmt.Printf("Type of l is %T\n", l)
	fmt.Println(l, "bytes written successfully")

	/* err = f.Close() //正确写入文件的话，要关闭文件。如果这里关闭的话，下面的写入字节的事情 就会报错说文件关闭了。
	if err != nil {
		fmt.Println(err)
		return
	} */

	//将字节写入文件
	//将字节写入文件和写入字符串非常的类似。我们将使用 Write(专门处理[]byte) 方法将字节写入到文件。
	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 101, 10}
	n2, err := f.Write(d2) //写入字节
	if err != nil {
		fmt.Println(err, "222")
		f.Close() //错误的时候 关闭文件链接
		return
	}
	fmt.Println(n2, "bytes written successfully")
	/*err = f.Close() //正确的话 关闭文件链接
	if err != nil {
		fmt.Println(err)
		return
	}*/

	//将字符串一行一行的写入文件
	d := []string{"Welcome to the world of Go1.", "Go is a compiled language.",
		"It is easy to learn Go."}

	for _, v := range d {
		fmt.Fprintln(f, v) //func Fprintln(w io.Writer, a ...interface{}) (n int, err error)格式化写入文件
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")

	//追加到文件。这个文件将以追加和写的方式打开。这些标志将通过 Open 方法实现。当文件以追加的方式打开，我们添加新的行到文件里。
	//flag代表的意思 o_rdonly read only 只读 o_wronly write only 只写 o_rdwr read write 可读可写 o_trunc 若文件存在则长度被截为0(属性不变)
	/*
		const (
		        O_RDONLY int = syscall.O_RDONLY // 只读打开文件和os.Open()同义
		        O_WRONLY int = syscall.O_WRONLY // 只写打开文件
		        O_RDWR   int = syscall.O_RDWR   // 读写方式打开文件
		        O_APPEND int = syscall.O_APPEND // 当写的时候使用追加模式到文件末尾
		        O_CREATE int = syscall.O_CREAT  // 如果文件不存在，此案创建
		        O_EXCL   int = syscall.O_EXCL   // 和O_CREATE一起使用, 只有当文件不存在时才创建
		        O_SYNC   int = syscall.O_SYNC   // 以同步I/O方式打开文件，直接写入硬盘.
		        O_TRUNC  int = syscall.O_TRUNC  // 如果可以的话，当打开文件时先清空文件
		)
	*/

	f2, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644) //当写的时候使用追加模式到文件末尾，只写打开文件
	if err != nil {
		fmt.Println(err)
		return
	}

	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f2, newLine)
	if err != nil {
		fmt.Println(err)
		f2.Close()
		return
	}
	err = f2.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")

	//并发写文件
	//当多个 goroutines 同时（并发）写文件时，我们会遇到竞争条件(race condition)。因此，当发生同步写的时候需要一个 channel 作为一致写入的条件。

	f3, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	var wg sync.WaitGroup //开启工作池计数器

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go filewriter(i, &wg, f3) //遍历打开协程  传入参数

	}
	wg.Wait() //等待所有工作池关闭
	fmt.Println("Close the file write")
	fmt.Fprintln(f3, "Close the file write")
	fmt.Fprintln(f3, "File appended successfully")
	err = f3.Close() //关闭打开文件
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test5highlevel# go run writingfiles.go
Type of f is *os.File
Type of l is int
12 bytes written successfully
13 bytes written successfully
file written successfully
file appended successfully
Start worker  9
Start worker  0
Start worker  1
Start worker  2
Start worker  3
Start worker  4
Start worker  5
Start worker  6
Start worker  7
Start worker  8
Close the worker  9
Close the worker  0
Close the worker  4
Close the worker  8
Close the worker  5
Close the worker  1
Close the worker  2
Close the worker  3
Close the worker  6
Close the worker  7
Close the file write
root@e7939faf8694:/go/src/LessonGo/class_tour/test5highlevel#
*/
