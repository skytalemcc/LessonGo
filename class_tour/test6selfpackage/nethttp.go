package main

/***
对于很多现代应用来说，访问互联网上的信息和访问本地文件系统一样重要。
使用net包可以更简单地用网络收发信息，还可以建立更底层的网络连接，编写服务器程序。
***/

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func curl() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url) //http.Get函数是创建HTTP请求的函数  resp这个结构体中得到访问的请求结果
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			//break和continue语句会改变控制流。os.Exit(1)则是退出。
			//和其它语言中的break和continue一样，break会中断当前的循环，并开始执行循环之后的内容，而continue会跳过当前循环，并开始执行下一次循环。
			os.Exit(1) //异常的时候 进行退出
		}
		b, err := ioutil.ReadAll(resp.Body) //resp的Body字段包括一个可读的服务器响应流  ioutil.ReadAll函数从response中读取到全部内容
		resp.Body.Close()                   //resp.Body.Close关闭resp的Body流，防止资源泄露
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("The context of fetching are %s", b)
	}
}

//搭建一个web service使用自建库
func webservice() {
	http.HandleFunc("/", handler) // 每一个请求都会发给handler来进行业务处理
	// http.HandleFunc("/count", count) 可以添加别的请求，根据请求的url不同会调用不同的函数
	http.ListenAndServe("localhost:8000", nil) //启动服务，占用端口

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path) //通过request 来获得访问路径, 标准输出流的fmt.Fprintf作为返回ResponseWriter
	//这里我可以定义任何一种返回方式，例如返回数字，字符，图片，文件内容等等。这样curl访问地址的时候 就能进行解析
}

func main() {
	go webservice()

	for {
		time.Sleep(2 * time.Second)
		curl() //死循环 ，让程序挂起 来校验
	}
}

/***
结果集 ：启动一个web服务可以 使用go run net.go &

当你发起访问的时候
root@e7939faf8694:/go/src/LessonGo# curl http://127.0.0.1:8000
URL.Path = "/"
root@e7939faf8694:/go/src/LessonGo#

root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# go run net.go  http://localhost:8000/hello
The context of fetching are URL.Path = "/hello"
The context of fetching are URL.Path = "/hello"
The context of fetching are URL.Path = "/hello"
The context of fetching are URL.Path = "/hello"
The context of fetching are URL.Path = "/hello"
time.Sleep(2 * time.Second)The context of fetching are URL.Path = "/hello"
The context of fetching are URL.Path = "/hello"
The context of fetching are URL.Path = "/hello"
^Z
[1]+  Stopped                 go run net.go http://localhost:8000/hello
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage#


***/
