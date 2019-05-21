package main

import "fmt"

//同一个目录下面不能有个多 package main ，否则IDE会提示错误 ： main redeclared in this block
func main() {

	fmt.Println("hello world")
}

/***
结果集
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic# go run hello.go
hello world
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic#
***/

// 用于单行注释
/* 用于单行注释 */
/*** 用于多行注释
***/
