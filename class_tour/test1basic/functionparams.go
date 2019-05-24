package main

/***
什么是可变参数函数?
可变参数函数是一种参数个数可变的函数。
如果函数最后一个参数被记作 ...T ，这时函数可以接受任意个 T 类型参数作为最后一个参数。
请注意只有函数的最后一个参数才允许是可变的。

***/
import (
	"fmt"
)

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}
func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)
	//给可变参数函数传入切片
	//有一个可以直接将切片传入可变参数函数的语法糖，你可以在在切片后加上 ... 后缀。如果这样做，切片将直接传入函数，不再创建新的切片
	nums := []int{89, 90, 95}
	find(89, nums...) //给可变参数函数传入切片 要加语法糖
}
