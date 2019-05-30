package main

/*
log包实现了简单的日志服务。本包定义了Logger类型，该类型提供了一些格式化输出的方法。
本包也提供了一个预定义的“标准”Logger，可以通过辅助函数Print[f|ln]、Fatal[f|ln]和Panic[f|ln]访问，比手工创建一个Logger对象更容易使用。
Logger会打印每条日志信息的日期、时间，默认输出到标准错误。
Fatal系列函数会在写入日志信息后调用os.Exit(1)。
Panic系列函数会在写入日志信息后panic。

log模块主要提供了3类接口。分别是 “Print 、Panic 、Fatal ”。
对每一类接口其提供了3中调用方式，分别是 "Xxxx 、 Xxxxln 、Xxxxf"，基本和fmt中的相关函数类似。
const (
    // 字位共同控制输出日志信息的细节。不能控制输出的顺序和格式。
    // 在所有项目后会有一个冒号：2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒分辨率：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件无路径名+行号：d.go:23（会覆盖掉Llongfile）
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
)
这些选项定义Logger类型如何生成用于每条日志的前缀文本。
*/
import (
	//"fmt"
	"log"
	"os"
)

func main() {

	//不生成日志文件的。
	//标准日志输出，这种是不生成日志文件打印带时间撮的结果。
	arr := []int{2, 3}
	log.Print("Print array ", arr, "\n")
	log.Println("Println array", arr)
	log.Printf("Printf array with item [%d,%d]\n", arr[0], arr[1])

	/*
				对于log.Panic接口，该函数把日志内容刷到标准错误后调用 panic 函数
				log.Panicln("test for defer Panic")

				root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# go run log.go
				2019/05/30 12:01:30 Print array [2 3]
				2019/05/30 12:01:30 Println array [2 3]
				2019/05/30 12:01:30 Printf array with item [2,3]
				2019/05/30 12:01:30 test for defer Panic
				panic: test for defer Panic
				goroutine 1 [running]:
				log.Panicln(0xc000097f28, 0x1, 0x1)
		        /usr/local/go/src/log/log.go:347 +0xac
				main.main()
		        /go/src/LessonGo/class_tour/test6selfpackage/log.go:37 +0x23a
				exit status 2
				root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# echo $?
				1
				root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage



				对于 log.Fatal 接口，会先将日志内容打印到标准输出，接着调用系统的 os.exit(1) 接口，退出程序并返回状态 1 。
				但是有一点需要注意，由于是直接调用系统接口退出，defer函数不会被调用。
				defer func() {
					fmt.Println("--first--")
				}()
				log.Fatalln("test for defer Fatal")

				root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# go run log.go
				2019/05/30 11:51:39 Print array [2 3]
				2019/05/30 11:51:39 Println array [2 3]
				2019/05/30 11:51:39 Printf array with item [2,3]
				2019/05/30 11:51:39 test for defer Fatal
				exit status 1
				root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage#

				日志格式打印并抛出异常。直接系统退出，返回错误1，defer也无法执行了。
	*/

	//生成日志文件的。
	/*
				type Logger struct {
		    		// contains filtered or unexported fields
				}
				Logger类型表示一个活动状态的记录日志的对象，它会生成一行行的输出写入一个io.Writer接口。
				每一条日志操作会调用一次io.Writer接口的Write方法。
				Logger类型的对象可以被多个线程安全的同时使用，它会保证对io.Writer接口的顺序访问。

				func New(out io.Writer, prefix string, flag int) *Logger
				New创建一个Logger。参数out设置日志信息写入的目的地。参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）。
				该函数一共有三个参数：

				1，输出位置out，是一个io.Writer对象，该对象可以是一个文件也可以是实现了该接口的对象。通常我们可以用这个来指定日志输出到哪个文件。
				2，prefix 我们在前面已经看到，就是在日志内容前面的东西。我们可以将其置为 "[Info]" 、 "[Warning]"等来帮助区分日志级别。
				3，flags 是一个选项，显示日志开头的东西，可选的值有：
				Ldate         = 1 << iota     // 形如 2009/01/23 的日期
				Ltime                         // 形如 01:23:23   的时间
				Lmicroseconds                 // 形如 01:23:23.123123   的时间
				Llongfile                     // 全路径文件名和行号: /a/b/c/d.go:23
				Lshortfile                    // 文件名和行号: d.go:23
				LstdFlags     = Ldate | Ltime // 日期和时间
	*/

	file, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("fail to create test.log file!") //创建文件就直接失败没有成功，则直接返回即可
	}
	defer file.Close()
	//var logger *log.Logger
	logger := log.New(file, "[Info]", log.LstdFlags|log.Llongfile) //拼接 前缀，日期时间，文件路径 ，加错误打印。
	log.Println("正常输出日志到命令行终端")
	logger.Println("正常输出到日志写入文件")
	logger.SetPrefix("[Debug]") //更改前缀
	log.Println("调试输出日志到命令行终端")
	logger.Println("调试输出到日志写入文件")
	logger.SetFlags(log.LstdFlags | log.Lshortfile) // 设置日志格式
	//log.Fatalln("在命令行终端输出日志并执行os.exit(1)")
	logger.Fatalln("在日志文件中写入日志并执行os.exit(1)") //打印并退出
	//log.Panicln("在命令行终端输出panic，并中断程序执行")
	logger.Printf("Printf array with item [%d,%d]\n", arr[0], arr[1]) //格式化真正的日志
	//logger.Fatalf("Printf array with item [%d,%d]\n", arr[0], arr[1])  //Fatalf等价于{l.Printf(v...); os.Exit(1)}
	logger.Panicln("在日志文件中写入panic，并中断程序执行")

	//SetOutput设置标准logger的输出目的地，默认是标准错误输出。
	file2, err2 := os.OpenFile("test2.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err2 != nil {
		log.Fatalln("fail to create test2.log file!") //创建文件就直接失败没有成功，则直接返回即可
	}
	defer file2.Close()

	logger.SetOutput(file2) //可以设置多个文件输入，
	log.Println("开辟新的文件来写入")
	logger.Println("开辟新的文件来写入，可以定义INFO,ERROR,DEBUG三个日志文件")
}

/*
结果集:
控制台输出
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# go run log.go
2019/05/30 12:55:18 Print array [2 3]
2019/05/30 12:55:18 Println array [2 3]
2019/05/30 12:55:18 Printf array with item [2,3]
2019/05/30 12:55:18 正常输出日志到命令行终端
2019/05/30 12:55:18 调试输出日志到命令行终端
2019/05/30 12:55:18 开辟新的文件来写入
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage#

test.log
[Info]2019/05/30 12:37:09 log.go:111: 将日志写入文件
[Info]2019/05/30 12:37:48 /go/src/LessonGo/class_tour/test6selfpackage/log.go:111: 将日志写入文件
[Info]2019/05/30 12:40:03 /go/src/LessonGo/class_tour/test6selfpackage/log.go:111: 正常输出到日志写入文件
[Debug]2019/05/30 12:40:03 /go/src/LessonGo/class_tour/test6selfpackage/log.go:114: 调试输出到日志写入文件
[Debug]2019/05/30 12:42:48 log.go:117: 在日志文件中写入日志并执行os.exit(1)
[Info]2019/05/30 12:43:35 /go/src/LessonGo/class_tour/test6selfpackage/log.go:111: 正常输出到日志写入文件
[Debug]2019/05/30 12:43:35 /go/src/LessonGo/class_tour/test6selfpackage/log.go:114: 调试输出到日志写入文件
[Debug]2019/05/30 12:43:35 log.go:117: 在日志文件中写入日志并执行os.exit(1)
[Info]2019/05/30 12:44:15 /go/src/LessonGo/class_tour/test6selfpackage/log.go:111: 正常输出到日志写入文件
[Debug]2019/05/30 12:44:15 /go/src/LessonGo/class_tour/test6selfpackage/log.go:114: 调试输出到日志写入文件
[Debug]2019/05/30 12:44:30 log.go:119: 在日志文件中写入panic，并中断程序执行
[Info]2019/05/30 12:46:51 /go/src/LessonGo/class_tour/test6selfpackage/log.go:111: 正常输出到日志写入文件
[Debug]2019/05/30 12:46:51 /go/src/LessonGo/class_tour/test6selfpackage/log.go:114: 调试输出到日志写入文件
[Debug]2019/05/30 12:46:51 log.go:119: Printf array with item [2,3]
[Debug]2019/05/30 12:46:51 log.go:120: 在日志文件中写入panic，并中断程序执行

test2.log
[Debug]2019/05/30 12:54:30 log.go:131: 开辟新的文件来写入，可以定义INFO,ERROR,DEBUG三个日志文件
[Debug]2019/05/30 12:55:18 log.go:132: 开辟新的文件来写入，可以定义INFO,ERROR,DEBUG三个日志文件
*/
