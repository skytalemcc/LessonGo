//每个go程序都是由包组成的
/*
程序的运行入口的包main函数
用圆括号来打包多个的方式导入语句
在包中，引用时候，首写字母大写则是可以导出和引用的。小写字母是不能的
*/

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {

	fmt.Println("My favorite number is ", rand.Intn(10))
	fmt.Printf("Now you have %g problem \n", math.Nextafter(2, 3))
	fmt.Println(math.Pi) // Pi是首字母大写,pi则无法引用
}

/***
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic# go run package.go
My favorite number is  1
Now you have 2.0000000000000004 problem
3.141592653589793
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic#

***/
