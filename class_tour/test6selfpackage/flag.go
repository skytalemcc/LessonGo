package main

/*
flag包提供了一系列解析命令行参数的功能接口。
如果用户在命令行输入了一个无效的标志参数，或者输入-h或-help参数，那么将打印所有标志参数的名字、默认值和描述信息。
三个属性 第一个是的命令行标志参数的名字，然后是该标志参数的默认值（这里是false），最后是该标志参数对应的描述信息。
在使用标志参数对应的变量之前先调用flag.Parse函数，用于更新每个标志参数对应变量的值（之前是默认值）。
如果在flag.Parse函数解析命令行参数时遇到错误，默认将打印相关的提示信息，然后调用os.Exit(2)终止程序。
命令行 flag 的语法有如下三种形式：
-flag // 只支持bool类型
-flag=x
-flag x // 只支持非bool类型

*/

import (
	"flag"
	"fmt"
	"os"
)

// 实际中应该用更好的变量名
var (
	h bool

	v, V bool
	t, T bool
	q    *bool

	s string
	p string
	c string
	g string
)

func init() {
	//flag.XxxVar()，将 flag 绑定到一个变量上 第一个参数 ：接收flagname的实际值的 第二个参数 ：flag名称为flagname 第三个参数 ：flagname默认值为1234 第四个参数 ：flagname的提示信息
	flag.BoolVar(&h, "h", false, "this help")

	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.BoolVar(&V, "V", false, "show version and configure options then exit")

	flag.BoolVar(&t, "t", false, "test configuration and exit")
	flag.BoolVar(&T, "T", false, "test configuration, dump it and exit")

	// 另一种绑定方式
	q = flag.Bool("q", false, "suppress non-error messages during configuration testing")

	// 注意 `signal`。默认是 -s string，有了 `signal` 之后，变为 -s signal
	flag.StringVar(&s, "s", "", "send `signal` to a master process: stop, quit, reopen, reload")
	flag.StringVar(&p, "p", "/usr/local/nginx/", "set `prefix` path")
	flag.StringVar(&c, "c", "conf/nginx.conf", "set configuration `file`")
	flag.StringVar(&g, "g", "conf/nginx.conf", "set global `directives` out of configuration file")

	// 改变默认的 Usage，flag包中的Usage 其实是一个函数类型。这里是覆盖默认函数实现，具体见后面Usage部分的分析
	flag.Usage = usage
}

func main() {
	flag.Parse() //在所有的 flag 定义完成之后，可以通过调用 flag.Parse() 进行解析。

	if h {
		flag.Usage()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
`)
	flag.PrintDefaults()
}

/***
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# go run flag.go
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# go run flag.go  -h
nginx version: nginx/1.10.0
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
  -T    test configuration, dump it and exit
  -V    show version and configure options then exit
  -c file
        set configuration file (default "conf/nginx.conf")
  -g directives
        set global directives out of configuration file (default "conf/nginx.conf")
  -h    this help
  -p prefix
        set prefix path (default "/usr/local/nginx/")
  -q    suppress non-error messages during configuration testing
  -s signal
        send signal to a master process: stop, quit, reopen, reload
  -t    test configuration and exit
  -v    show version and exit
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# go run flag.go  -t
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# go run flag.go  -c 123
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# go run flag.go  -c=123
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage#
*/
