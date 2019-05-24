package main

/*
文件读取是所有编程语言中最常见的操作之一
1,将整个文件读取到内存
  使用绝对文件路径 或者使用命令行标记来传递文件路径 或者将文件绑定在二进制文件中
2,分块读取文件
3,逐行读取文件
*/

/*
将整个文件读取到内存
将整个文件读取到内存是最基本的文件操作之一。这需要使用 ioutil 包中的 ReadFile 函数。

*/

/*
总结一下：
将整个文件读取到内存
ioutil.ReadFile("test.txt")
分块读取文件
r := bufio.NewReader(f)
b := make([]byte, 3)
for{
	_, err := r.Read(b)
}

逐行读取文件
s := bufio.NewScanner(f)
for s.Scan() {
    fmt.Println(s.Text())
}
err = s.Err()

*/

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// 将整个文件读取到内存
	// 使用绝对文件路径 ，避免编译后的二进制文件再更换环境后无法执行。
	//看似这是一个简单的方法，但它的缺点是：文件必须放在程序指定的路径中，否则就会出错。
	data, err := ioutil.ReadFile("test.txt") //ReadFile 方法 返回结果为[]byte，需要进行转换不然生成的是ASIC码
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data)) //对字节码进行转换为字符串

	//使用命令行标记来传递文件路径，这样就可以 传入参数来作为路径。通过flag包可以从输入的命令行获取到文件路径，接着读取文件内容。

	fptr := flag.String("fpath", "test.txt", "file path to read from") //三个参数 标记名，默认值，标记的简短描述。这个fptr是字符串变量的指针。
	flag.Parse()
	fmt.Println("value of fpath is", *fptr) // 命令的时候 go run readingfiles.go  -fpath=/root/test.txt ，返回结果value of fpath is /root/test.txt
	//传递路径的方法很好 ，但是有一个更好的解决方法，将文件绑定再二进制文件中
	/*
			安装 go get -u github.com/gobuffalo/packr/...
			packr 会把静态文件（例如 .txt 文件）转换为 .go 文件，接下来，.go 文件会直接嵌入到二进制文件中。
			import "github.com/gobuffalo/packr"
			box := packr.NewBox("../filehandling") 这是定义一个文件夹，其内容会被嵌入到二进制中
		    data := box.String("test.txt")
			fmt.Println("Contents of file:", data)

			使用packr 可以将很多项目目录下面的静态文件等，打包到二进制里面取。需要好好研究。暂时略。静态文件其实可以一块打包到二进制也可以放在二进制外面来交付。
			或者使用makefile的时候来负责打包。
	*/

	//分块读取文件
	//当文件非常大的时候 ，内存是无法承受的，整个文件读入内存没有意义。更好的方法是分块读取文件。这可以使用 bufio 包来完成。

	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	f2, err2 := os.Open("test.txt")
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	r := bufio.NewReader(f) //读取文件 返回*Reader
	b := make([]byte, 20)   //长度为3的字节组
	for {
		_, err := r.Read(b) // *Reader.Read方法支持按字节数来读取
		if err != nil {
			fmt.Println("Error reading file And end read by []bytes:", err) //循环结束，没有内容了通过break来退出所在for循环。然后下一步defer来关闭
			break
		}
		fmt.Println(string(b)) //不断循环读取，打印字符
	}

	//逐行读取文件 这可以使用 bufio 来实现。
	s := bufio.NewScanner(f2) //使用它来扫描文件并且逐行读取
	for s.Scan() {            //使用scan方法来读取下一行
		fmt.Println(s.Text()) //使用text来打印内容
	}
	err2 = s.Err() //当 Scan 返回 false 时，除非已经到达文件末尾（此时 Err() 返回 nil），否则 Err() 就会返回扫描过程中出现的错误。
	if err2 != nil {
		fmt.Println("error happened when scan")
	}

	defer func() { //匿名函数 闭包。
		if err = f.Close(); err != nil {
			fmt.Println(err)
			f.Close()

		} else {
			f.Close()
		}

		if err = f2.Close(); err2 != nil {
			fmt.Println(err2)
			f2.Close()

		} else {
			f2.Close()
		}
	}()

}

/*
结果集:
root@e7939faf8694:/go/src/LessonGo/class_tour/test5highlevel# go run readingfiles.go
Contents of file: Hello World
hello worlde
Welcome to the world of Go1.
Go is a compiled language.
It is easy to learn Go.
File handling is easy.
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
Close the worker  2
Close the worker  9
Close the worker  0
Close the worker  1
Close the worker  4
Close the worker  3
Close the worker  5
Close the worker  6
Close the worker  8
Close the worker  7
Close the file write
File appended successfully

value of fpath is test.txt
Hello World
hello wo
rlde
Welcome to the
world of Go1.
Go is
a compiled language.

It is easy to learn
 Go.
File handling i
s easy.
Start worker
  9
Start worker  0

Start worker  1
Star
t worker  2
Start wo
rker  3
Start worker
  4
Start worker  5

Start worker  6
Star
t worker  7
Start wo
rker  8
Close the wo
rker  2
Close the wo
rker  9
Close the wo
rker  0
Close the wo
rker  1
Close the wo
rker  4
Close the wo
rker  3
Close the wo
rker  5
Close the wo
rker  6
Close the wo
rker  8
Close the wo
rker  7
Close the fi
le write
File append
ed successfully
pend
Error reading file And end read by []bytes: EOF
Hello World
hello worlde
Welcome to the world of Go1.
Go is a compiled language.
It is easy to learn Go.
File handling is easy.
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
Close the worker  2
Close the worker  9
Close the worker  0
Close the worker  1
Close the worker  4
Close the worker  3
Close the worker  5
Close the worker  6
Close the worker  8
Close the worker  7
Close the file write
File appended successfully
root@e7939faf8694:/go/src/LessonGo/class_tour/test5highlevel#
*/
