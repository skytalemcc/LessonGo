package main

/***
map 是在 Go 中将值（value）与键（key）关联的内置类型。通过相应的键可以获取到值。
映射将键映射到值。
映射的零值为 nil 。nil 映射既没有键，也不能添加键。
make 函数会返回给定类型的映射，并将其初始化备用。
Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。

定义map 声明变量，默认 map 是 nil
var map_variable map[key_data_type]value_data_type
使用make函数
map_variable := make(map[key_data_type]value_data_type)

如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
***/

import "fmt"

type HashMap struct {
	key      string
	value    string
	hashCode int
	next     *HashMap //声明结构体类型的指针
}

func main() {

	var countryCapitalMap map[string]string     //创建map集合，形式声明
	countryCapitalMap = make(map[string]string) //实际创建，启用内存块
	countryCapitalMap["France"] = "巴黎"          //增加元素
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	for country := range countryCapitalMap { //遍历 map 中所有的元素需要用 for range 循环。
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	//value, ok := map[key] 查看元素在集合中是否存在
	capital, ok := countryCapitalMap["American"] /*获取一条元素 。通过双赋值检测某个键是否存在。 如果确定是真实的,则存在,否则不存在 */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	//插入一条元素或者修改一条元素
	countryCapitalMap["American "] = "华盛顿"
	//删除一条元素 ，如果存在的话
	delete(countryCapitalMap, "Japan")

	var hs HashMap
	fmt.Println(hs)

}

/***
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic# go run maps.go
Japan 首都是 东京
India  首都是 新德里
France 首都是 巴黎
Italy 首都是 罗马
American 的首都不存在
{  0 <nil>}
root@e7939faf8694:/go/src/LessonGo/class_tour/test1basic#

***/
