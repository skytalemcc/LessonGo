//每个go程序都是由包组成的
/*
程序的运行入口的包main函数
用圆括号来打包多个的方式导入语句
在包中，引用时候，首写字母大写则是可以导出和引用的。小写字母是不能的。

包用于组织 Go 源代码，提供了更好的可重用性与可读性。由于包提供了代码的封装，因此使得 Go 应用程序易于维护。
所有可执行的 Go 程序都必须包含一个 main 函数。这个函数是程序运行的入口。main 函数应该放置于 main 包中。
*/

//package packagename 这行代码指定了某一源文件属于一个包。它应该放在每一个源文件的第一行。

package main

import (
	"fmt"
	"math"
	"math/rand"
	//"geometry/rectangle" 导入自定义包,将其他目录下的文件导入进来，即可参与编译。

	//导入了包，却不在代码中使用它，这在 Go 中是非法的。当这么做时，编译器是会报错的。其原因是为了避免导入过多未使用的包，从而导致编译时间显著增加。
	//在程序开发的活跃阶段，又常常会先导入包，而暂不使用它。遇到这种情况就可以使用空白标识符 _。
	//有时候我们导入一个包，只是为了确保它进行了初始化，而无需使用包中的任何函数或变量。
	_ "image/jpeg" //导入但不使用，只使用被引用包的init函数
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
