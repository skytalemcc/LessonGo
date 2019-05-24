package main

/***
Go 语言切片是对数组的抽象。
Go 数组的长度不可改变 在特定场景中这样的集合就不太适用
功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

切片是由数组建立的一种方便、灵活且功能强大的包装（Wrapper）。切片本身不拥有任何数据。它们只是对现有数组的引用。


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

切片就像数组的引用
切片并不存储任何数据，它只是描述了底层数组中的一段。
更改切片的元素会修改其底层数组中对应的元素。
与它共享底层数组的切片都会观测到这些修改。

切片拥有 长度 和 容量。
切片的长度就是它所包含的元素个数。 已经有的。
切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。可以有的。
切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。
***/

import (
	"fmt"
	"strings"
)

func main() {

	//切片初始化
	s := []int{1, 2, 3} //直接初始化切片，[]int表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
	s1 := s[:]          //对切片s的引用
	s2 := s[2:]         //对切片s的引用 ，从索引2开始 创建切片
	s3 := s[:1]         //对切片s的引用 ，到索引2结束 创建切片
	fmt.Println(s, s1, s2, s3)
	//切片的长度是切片中的元素数。切片的容量是从创建切片索引开始的底层数组中元素数。
	fmt.Println(len(s), cap(s), len(s3), cap(s3))

	//s :=make([]int,len,cap)  通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片
	//make 函数创建一个数组，并返回引用该数组的切片。
	var numbers = make([]int, 3, 5) //len代表目前有几个值，cap计算切片可以最长到多少
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)

	//空(nil)切片 一个切片在未初始化之前默认为 nil，长度为 0
	//var numbers []int  切片是空的 len=0 cap=0 slice=[]  nil 切片的长度和容量为 0 且没有底层数组。
	//创建切片
	numints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(numints[2:5], len(numints), cap(numints))

	//如果想增加切片的容量，必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
	var nums []int
	nums = append(nums, 1) //向一个空的切片组追加新元素
	fmt.Println(nums)
	nums = append(nums, 2, 3, 4) //追加多个新元素
	fmt.Println(nums)

	//切片可以用内建函数 make 来创建，这也是你创建动态数组的方式。make 函数会分配一个元素为零值的数组并返回一个引用了它的切片
	numbers1 := make([]int, len(nums), (cap(nums))*2)
	//创建切片 numbers1 是之前切片的两倍容量

	//切片持有对底层数组的引用。只要切片在内存中，数组就不能被垃圾回收。在内存管理方面，这是需要注意的。
	//让我们假设我们有一个非常大的数组，我们只想处理它的一小部分。然后，我们由这个数组创建一个切片，并开始处理切片。
	//这里需要重点注意的是，在切片引用时数组仍然存在内存中。
	//一种解决方法是使用 copy 函数 func copy(dst，src[]T)int 来生成一个切片的副本。这样我们可以使用新的切片，原始数组可以被垃圾回收。
	copy(numbers1, nums) //拷贝nums到numbers1内容
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers1), cap(numbers1), numbers1)

	//切片文法 切片文法类似于没有长度的数组文法。
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	//在进行切片时，你可以利用它的默认行为来忽略上下界。切片下界的默认值为 0，上界则是该切片的长度。
	//a[0:10]和a[:10] 等价切片

	//切片的切片  切片可包含任何类型，甚至包括其它的切片。
	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X" //切片自己不拥有任何数据。它只是底层数组的一种表示。对切片所做的任何修改都会反映在底层数组中。
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " ")) //字符串拼接，更高效
	}
}

/***
结果集:
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic# go run slice.go
[1 2 3] [1 2 3] [3] [1]
3 3 1 3
len=3 cap=5 slice=[0 0 0]
[2 3 4] 9 9
[1]
[1 2 3 4]
len=4 cap=8 slice=[1 2 3 4]
[true false true true false true]
X _ X
O _ X
_ _ O
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic#

***/
