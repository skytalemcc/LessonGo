package main

import "fmt"

//同一个目录下面不能有个多 package main ，否则IDE会提示错误 ： main redeclared in this block
func main() {

	fmt.Println("hello world")
}
