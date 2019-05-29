package main

/***
大多数的程序都是处理输入，产生输出；这也正是“计算”的定义。但是, 程序如何获取要处理的输入数据呢？
一些程序生成自己的数据，但通常情况下，输入来自于程序外部：文件、网络连接、其它程序的输出、敲键盘的用户、命令行参数或其它类似输入源。
os包以跨平台的方式，提供了一些与操作系统交互的函数和变量。程序的命令行参数可从os包的Args变量获取；os包外部使用os.Args访问该变量。
os.Args变量是一个字符串（string）的切片（slice）
os.Args[0], 是命令本身的名字；其它的元素则是程序启动时传给它的参数
***/

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func echo() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	//优化，使用range的方式
	s2, sep2 := "", ""
	for _, arg := range os.Args[1:] {
		s2 += sep2 + arg //这种拼接的方式代价比较高
		sep2 = " "
	}
	fmt.Println(s2)
	//优化，直接使用stings来进行拼接
	fmt.Println(strings.Join(os.Args[1:], " "))

}

/***
对文件做拷贝、打印、搜索、排序、统计或类似事情的程序都有一个差不多的程序结构：一个处理输入的循环，在每个元素上执行计算处理，在处理的同时或最后产生输出。
bufio.Scanner、ioutil.ReadFile和ioutil.WriteFile都使用*os.File的Read和Write方法，但是，大多数程序员很少需要直接调用那些低级（lower-level）函数。
高级（higher-level）函数，像bufio和io/ioutil包中所提供的那些，用起来要容易点。
***/
func uniq() {

	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename) //
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") { //ReadFile函数返回一个字节切片（byte slice），必须把它转换为string，才能用strings.Split分割。
			counts[line]++ //统计
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line) //打印重复的行数和内容
		}
	}
}

func main() {
	echo()
	uniq()
}
