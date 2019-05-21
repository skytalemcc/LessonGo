package main

/***
使用range可以进行遍历
Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对的 key 值。
当使用 for 循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。
***/
import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow { //range专门来解决遍历，返回索引和值 。range返回的i 即为pow数组自己的索引
		fmt.Printf("index=%d value=%d\n", i, v)
	}
	//可以将下标或值赋予 _ 来忽略它。 for i, _ := range pow , for _, value := range pow 忽略值或者索引

	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs { //range的k 即为map对应的key
		fmt.Printf("%s -> %s\n", k, v)
	}

	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "goa" {
		fmt.Println(i, c)
	}

}

/***
结果集:
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic# go run range.go
index=0 value=1
index=1 value=2
index=2 value=4
index=3 value=8
index=4 value=16
index=5 value=32
index=6 value=64
index=7 value=128
a -> apple
b -> banana
0 103
1 111
2 97
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic#

***/
