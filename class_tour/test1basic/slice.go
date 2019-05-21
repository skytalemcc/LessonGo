package main

/***
Go 语言切片是对数组的抽象。
Go 数组的长度不可改变 在特定场景中这样的集合就不太适用
功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
定义切片
你可以声明一个未指定大小的数组来定义切片：
var identifier []type
切片不需要说明长度。或使用make()函数来创建切片:

var slice1 []type = make([]type, len)
也可以简写为
slice2 := make([]type, len)

也可以指定容量，其中capacity为可选参数。
make([]T, length, capacity)

这里 len 是数组的长度并且也是切片的初始长度。

***/

import "fmt"

func main() {

	//切片初始化
	s := []int{1, 2, 3} //直接初始化切片，[]int表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
	s1 := s[:]          //对切片s的引用
	s2 := s[2:]         //对切片s的引用 ，从索引2开始
	s3 := s[:2]         //对切片s的引用 ，到索引2结束
	fmt.Println(s, s1, s2, s3)

	//s :=make([]int,len,cap)  通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片
	var numbers = make([]int, 3, 5) //len代表目前有几个值，cap计算切片可以最长到多少
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)

	//空(nil)切片 一个切片在未初始化之前默认为 nil，长度为 0
	//var numbers []int  切片是空的 len=0 cap=0 slice=[]
	//创建切片
	numints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(numints[2:5], len(numints), cap(numints))

	//如果想增加切片的容量，必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
	var nums []int
	nums = append(nums, 1) //向一个空的切片组追加新元素
	fmt.Println(nums)
	nums = append(nums, 2, 3, 4) //追加多个新元素
	fmt.Println(nums)
	numbers1 := make([]int, len(nums), (cap(nums))*2)
	//创建切片 numbers1 是之前切片的两倍容量
	copy(numbers1, nums) //拷贝nums到numbers1内容
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers1), cap(numbers1), numbers1)

}

/***
结果集:
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic# go run slice.go
[1 2 3] [1 2 3] [3] [1 2]
len=3 cap=5 slice=[0 0 0]
[2 3 4] 9 9
[1]
[1 2 3 4]
len=4 cap=8 slice=[1 2 3 4]
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic#

***/
