package main

/***
Go 程序使用 error 值来表示错误状态。error 类型是一个内建接口。
error 为 nil 时表示成功；非 nil 的 error 表示失败。
一般自建或者三方库 返回时候会自带err返回，以供判断使用。

error本身是一个接口类型 ，有个签名为Error() string 的方法
type error interface {
    Error() string
}
绝不要忽略错误。忽视错误会带来问题

***/

import (
	"errors" //引入errors 来自定义错误内容。
	"fmt"
	"math"
	"net"
	"os"
)

//自定义返回，含有error。error 关键字 是内建接口。
func Sqrt(x float64) (float64, error) { //error关键字是接口类型 。可以为nil。或者自定义
	return 0, nil
}

//使用new函数来自定义error返回内容。
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero") //使用New来自定义报错。
		//使用 Errorf 给错误添加更多信息
		//return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}
func main() {
	fmt.Println(Sqrt(2))

	_, err := os.Open("/test.txt") //方法是这么定义的 func Open(name string) (file *File, err error)
	if err != nil {
		fmt.Println(err)

	}

	//断言底层结构体类型，使用结构体字段获取更多信息 。os.Open 查看 If there is an error, it will be of type *PathError.
	if err, ok := err.(*os.PathError); ok { // err.(*os.PathError)  这个方法 为 os.Open 本来返回error的正确格式。可以查一下doc。
		fmt.Println("File at path", err.Path, "failed to open")

	}

	//对底层类型进行断言，然后通过调用该结构体类型的方法，来获取更多的信息
	//net.LookupHost的error返回类型为*net.DNSError ,它本身有很多指针定义的方法

	addr, err := net.LookupHost("www.baidu.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() { //这些方法 就是 net.LookupHost 返回error类型自带的方法。
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)

	//使用 New 函数创建自定义错误
	//在使用 New 函数 创建自定义错误之前，可以取官网看一下New函数的定义。
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}

/***
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test3interface# go run error.go
0 <nil>
open /test.txt: no such file or directory
File at path /test.txt failed to open
[183.232.231.174 183.232.231.172]
Area calculation failed, radius is less than zero
root@e7939faf8694:/go/src/LessonGo/class_tour/test3interface#
***/
