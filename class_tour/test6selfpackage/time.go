package main

/*
time包提供了时间的显示和测量用的函数。日历的计算采用的是公历。
时间可分为时间点与时间段：时间点(Time) 和时间段(Duration)
满足特定业务：时区(Location) ，Ticker(周期性定时器) ，Timer(一次性定时器)
*/
import (
	"fmt"
	"time"
)

func main() {

	//时间常量，会预定义的版式用于Time.Format和Time.Parse函数。用在版式中的参考时间是：Mon Jan 2 15:04:05 MST 2006 这个时间是Unix time 1136239445
	//Time代表一个纳秒精度的时间点。
	//程序中应使用Time类型值来保存和传递时间，而不能用指针。就是说，表示时间的变量和字段，应为time.Time类型，而不是*time.Time.类型。
	//一个Time类型值可以被多个go程同时使用。时间点可以使用Before、After和Equal方法进行比较。Sub方法让两个时间点相减，生成一个Duration类型值（代表时间段）。
	//Add方法给一个时间点加上一个时间段，生成一个新的Time类型时间点。

	fmt.Println("返回当前的本地操作系统的时间  ", time.Now())  //返回TIME类型的结构体
	fmt.Println("返回Unix时间  ", time.Now().Unix()) //Unix将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位秒）。

	fmt.Println(time.Now().Format(time.UnixDate)) //按照自带的预定义的const UnixDate版式来输出当前时间
	fmt.Println(time.Now().Format(time.ANSIC))    //按照自带的预定义的const ANSIC版式来输出当前时间
	//Format根据layout指定的格式返回t代表的时间点的格式化文本表示。
	fmt.Println("按照指定格式返回格式化的时间 ", time.Now().Format("2006-01-02 15:04:05"))
	//Parse解析一个格式化的时间字符串并返回它代表的时间。 用来将一个字符串转换为时间格式
	dt, _ := time.Parse("2006-01-02 15:04:05", "2019-05-28 13:24:51") //第一个参数是格式layout，第二个参数是真实字符串时间值 。dt为Time结构体
	fmt.Println(dt.Format(time.ANSIC))                                //按预定于格式打印
	fmt.Println(time.Now().Date())                                    //返回年月日多个值
	fmt.Println(time.Now().Clock())                                   //返回时分秒多个值
	fmt.Println("返回年  ", time.Now().Year())
	fmt.Println("返回月  ", time.Now().Month())
	fmt.Println("返回日  ", time.Now().Day())
	fmt.Println("返回时  ", time.Now().Hour())
	fmt.Println("返回分  ", time.Now().Minute())
	fmt.Println("返回秒  ", time.Now().Second())
	fmt.Println("返回星期几  ", time.Now().Weekday())

	//golang默认采用UTC，即Unix标准时间
	fmt.Println(time.Now().Location())                                         //返回时间的地点和时区信息
	fmt.Println(time.Now().Zone())                                             //返回该时区的规范名（如"UTC"）和该时区相对于UTC的时间偏移量（单位秒）。
	loc, _ := time.LoadLocation("Asia/Shanghai")                               //如果参数是UTC，返回UTC。如果是Local，则返回Local。如果是时区数据库记录的地点名，则返回地点和时区 *Location
	fmt.Println("使用IN函数，传入参数loc 返回时间 ", time.Now().In(loc).Format(time.ANSIC)) //返回上海时间，按照格式打印出来

	//时间区段
	tp, _ := time.ParseDuration("1.5h") //ParseDuration解析一个时间段字符串。合法的单位有"s"、"m"、"h"等
	//Round(四舍五入)和Truncate(向下取整)返回的是最接近t的时间点。Hours Minutes Seconds String 对应的小时，分，秒，以及字符串格式
	fmt.Println(time.Now().Round(tp), time.Now().Truncate(tp), tp.Hours(), tp.Minutes(), tp.Seconds(), tp.String())

	//Duration表示两个时间之间经过的时间，要将整数个某时间单元表示为Duration类型值，用乘法：
	//Sleep阻塞当前go程至少d代表的时间段。After会在另一线程经过时间段d后向返回值发送当时的时间。
	time.Sleep(time.Duration(2) * time.Second)           //休眠2秒，休眠时处于阻塞状态，后续程序无法执行
	time.After(time.Duration(2) * time.Second)           // 非阻塞,可用于延迟
	fmt.Println(time.Since(dt))                          //Now().Sub(t)， 可用来计算一段业务的消耗时间，当前时间 和起始时间dt 。来计算时间差。
	fmt.Println(dt.Add(time.Duration(10) * time.Second)) //对某个时间增加一段时间 time.Second time.Day time.Hour
	fmt.Println("计算时间差 dt与当前时间的间隔 ", dt.Sub(time.Now()))
	fmt.Println(time.Now().After(dt))          // 新的为true，早的为false
	fmt.Println(time.Now().Before(dt))         //新的为false，早的为true
	time.After(time.Duration(2) * time.Second) //函数与 select 结合使用可用于处理程序超时设定

	//周期性时间任务 NewTicker
	//一次性时间任务 NewTimer
	//Ticker 类型包含一个 channel，有时我们会遇到每隔一段时间执行的业务(比如设置心跳时间等)，就可以用它来处理，这是一个重复的过程
	//NewTicker返回一个新的Ticker，该Ticker包含一个通道字段，并会每隔时间段d就向该通道发送当时的时间。
	//它会调整时间间隔或者丢弃tick信息以适应反应慢的接收者。如果d<=0会panic。关闭该Ticker可以释放相关资源。
	//Stop关闭一个Ticker。在关闭后，将不会发送更多的tick信息。Stop不会关闭通道t.C，以避免从该通道的读取不正确的成功。func (t *Ticker) Stop()
	//Tick是NewTicker的封装，只提供对Ticker的通道的访问。如果不需要关闭Ticker，本函数就很方便。
	//主要用来按照指定的时间周期来调用函数或者计算表达式，通常的使用方式是利用go新开一个协程使用，它是一个断续器
	tick := time.Tick(1 * time.Minute) //Tick是NewTicker的封装，只提供对Ticker的通道的访问。如果不需要关闭Ticker，本函数就很方便。
	for _ = range tick {               //每隔时间段来去循环 做事
		// do something
		break

	}

	// 可通过调用ticker.Stop取消
	ticker := time.NewTicker(1 * time.Minute) //func NewTicker(d Duration) *Ticker
	for _ = range ticker {
		// do something
		break
	}
	//Timer 类型用来代表一个单独的事件，当设置的时间过期后，发送当前的时间到 channel。一次性
	//NewTimer创建一个Timer，它会在最少过去时间段d后到期，向其自身的C字段发送当时的时间。 func NewTimer(d Duration) *Timer
	//AfterFunc另起一个go程等待时间段d过去，然后调用f。它返回一个Timer，可以通过调用其Stop方法来取消等待和对f的调用。 func AfterFunc(d Duration, f func()) *Timer
	//Reset使t重新开始计时，（本方法返回后再）等待时间段d过去后到期。如果调用时t还在等待中会返回真；如果t已经到期或者被停止了会返回假。 func (t *Timer) Reset(d Duration) bool
	//Stop停止Timer的执行。如果停止了t会返回真；如果t已经被停止或者过期了会返回假。Stop不会关闭通道t.C，以避免从该通道的读取不正确的成功。 func (t *Timer) Stop() bool
	//初始化一个到期时间据此时的间隔为2s的定时器
	t := time.NewTimer(2 * time.Second) //一次性的工作，等待结束后发送当前时间到chan
	expire := <-t.C
	fmt.Printf("Expiration time: %v.\n", expire)
}

/*
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# go run time.go
返回当前的本地操作系统的时间   2019-05-28 14:55:59.067563 +0000 UTC m=+0.000417101
返回Unix时间   1559055359
Tue May 28 14:55:59 UTC 2019
Tue May 28 14:55:59 2019
按照指定格式返回格式化的时间  2019-05-28 14:55:59
Tue May 28 13:24:51 2019
2019 May 28
14 55 59
返回年   2019
返回月   May
返回日   28
返回时   14
返回分   55
返回秒   59
返回星期几   Tuesday
Local
UTC 0
使用IN函数，传入参数loc 返回时间  Tue May 28 22:55:59 2019
2019-05-28 15:00:00 +0000 UTC 2019-05-28 13:30:00 +0000 UTC 1.5 90 5400 1h30m0s
1h31m10.0691535s
2019-05-28 13:25:01 +0000 UTC
计算时间差 dt与当前时间的间隔  -1h31m10.0693231s
true
false
Expiration time: 2019-05-28 14:56:03.0695953 +0000 UTC m=+4.002468801.
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage#

*/
