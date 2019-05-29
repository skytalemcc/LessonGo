package main

/*
time包提供了时间的显示和测量用的函数。日历的计算采用的是公历。
时间可分为时间点与时间段：时间点(Time) 和时间段(Duration)
满足特定业务：时区(Location) ，Ticker(周期性定时器) ，Timer(一次性定时器)
业务，TTL会话管理、锁、定时任务(闹钟)或更复杂的状态切换等等
*/
import (
	"fmt"
	"sync"
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

	//如果只是想单纯的等待的话，可以使用 time.Sleep 来实现
	/*
			周期性时间任务 NewTicker
			一次性时间任务 NewTimer
			Ticker 类型包含一个 channel，有时我们会遇到每隔一段时间执行的业务(比如设置心跳时间等)，就可以用它来处理，这是一个重复的过程
			NewTicker返回一个新的Ticker，该Ticker包含一个通道字段，并会每隔时间段d就向该通道发送当时的时间。
			它会调整时间间隔或者丢弃tick信息以适应反应慢的接收者。如果d<=0会panic。关闭该Ticker可以释放相关资源。
			Stop关闭一个Ticker。在关闭后，将不会发送更多的tick信息。Stop不会关闭通道t.C，以避免从该通道的读取不正确的成功。func (t *Ticker) Stop()
			Tick是NewTicker的封装，只提供对Ticker的通道的访问。如果不需要关闭Ticker，本函数就很方便。
			主要用来按照指定的时间周期来调用函数或者计算表达式，通常的使用方式是利用go新开一个协程使用，它是一个断续器


			type Ticker struct {
		    C <-chan Time // 周期性传递时间信息的通道
		    // 内含隐藏或非导出字段
		}


	*/
	tick := time.Tick(3 * time.Second) //Tick是NewTicker的封装，只提供对Ticker的通道的访问。如果不需要关闭Ticker，本函数就很方便。
	for range tick {                   //每隔时间段来去循环 做事
		// do something
		fmt.Println("Golang1")
		break
	}

	ticker := time.NewTicker(2 * time.Second) //func NewTicker(d Duration) *Ticker //定义时钟周期,每隔2秒写入时间
	i := 0
	for {
		nowTime := <-ticker.C ////阻塞，形成间隔
		//golang的time.NewTicker创建定时任务时，是阻塞同步的。如果不想因为同步阻塞了main线程，可以给每个定时函数分配一个goroutine协程。
		//<-ticker.C 没有接收的化 ，这样代表阻塞，等待消耗
		i++
		fmt.Println("ninhao + ", i)
		if nowTime.Hour() == 3 && nowTime.Minute() == 1 { //定时周期 ，特定时间做什么
			fmt.Println("Golang")
			break
		}

		if i == 3 {
			ticker.Stop() //关闭后，不会创建新的计时器
			break         //利用break 可以提前退出循环，break 终止当前的循环 ,for循环，if不是循环
		}
	}
	/*
		Timer 类型用来代表一个单独的事件，当设置的时间过期后，发送当前的时间到 channel。一次性
		NewTimer创建一个Timer，它会在最少过去时间段d后到期，向其自身的C字段发送当时的时间。 func NewTimer(d Duration) *Timer
		AfterFunc另起一个go程等待时间段d过去，然后调用f。它返回一个Timer，可以通过调用其Stop方法来取消等待和对f的调用。 func AfterFunc(d Duration, f func()) *Timer
		Reset使t重新开始计时，（本方法返回后再）等待时间段d过去后到期。如果调用时t还在等待中会返回真；如果t已经到期或者被停止了会返回假。 func (t *Timer) Reset(d Duration) bool
		Stop停止Timer的执行。如果停止了t会返回真；如果t已经被停止或者过期了会返回假。Stop不会关闭通道t.C，以避免从该通道的读取不正确的成功。 func (t *Timer) Stop() bool
		初始化一个到期时间据此时的间隔为2s的定时器
		timer.Reset 如果明确time已经expired，并且t.C已经被取空，那么可以直接使用Reset；如果程序之前没有从t.C中读取过值，这时需要首先调用Stop()，
		如果返回true，说明timer还没有expire，stop成功删除timer，可直接reset；如果返回false，说明stop前已经expire，需要显式drain channel。
	*/

	t := time.NewTimer(2 * time.Second) //一次性的工作，等待结束后发送当前时间到chan
	expire := <-t.C
	fmt.Printf("Expiration time: %v.\n", expire)
	fmt.Println(t.Reset(0)) //当Timer已经停止或者超时，返回false。当定时器未超时时，返回true。
	//当Reset返回false时，我们并不能认为一段时间之后，超时不会到来，实际上可能会到来，定时器已经生效了。
	t.Stop()

	/*
		ticker只要定义完成，从此刻开始计时，不需要任何其他的操作，每隔固定时间都会触发。
		timer定时器，是到固定时间后会执行一次
		如果timer定时器要每隔间隔的时间执行，实现ticker的效果，使用 func (t *Timer) Reset(d Duration) bool
	*/
	var wg sync.WaitGroup
	wg.Add(2)
	//NewTimer 创建一个 Timer，它会在最少过去时间段 d 后到期，向其自身的 C 字段发送当时的时间
	timer1 := time.NewTimer(2 * time.Second)

	/*
		NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，并会每隔时间段 d 就向该通道发送当时的时间。它会调
		整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。如果d <= 0会触发panic。关闭该 Ticker 可
		以释放相关资源。
		额外说明：
			time.NewTicker定时触发执行任务，当下一次执行到来而当前任务还没有执行结束时，会等待当前任务执行完毕后再执行下一次任务。查阅go官网的文档和经过代码验证。
			time.NewTimer和Reset()函数实现定时触发，Reset()函数可能失败，经测试。
	*/
	ticker1 := time.NewTicker(2 * time.Second)

	go func(t *time.Ticker) { //自带循环
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get ticker1", time.Now().Format("2006-01-02 15:04:05"))
		}
	}(ticker1)

	go func(t *time.Timer) { //一次性 ，需要使用reset来完成循环
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get timer", time.Now().Format("2006-01-02 15:04:05"))
			//Reset 使 t 重新开始计时，（本方法返回后再）等待时间段 d 过去后到期。如果调用时t
			//还在等待中会返回真；如果 t已经到期或者被停止了会返回假。
			t.Reset(2 * time.Second)
		}
	}(timer1)

	wg.Wait()

}

/*
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# go fmt time.go
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# go run time.go
返回当前的本地操作系统的时间   2019-05-29 03:01:28.2440235 +0000 UTC m=+0.000201901
返回Unix时间   1559098888
Wed May 29 03:01:28 UTC 2019
Wed May 29 03:01:28 2019
按照指定格式返回格式化的时间  2019-05-29 03:01:28
Tue May 28 13:24:51 2019
2019 May 29
3 1 28
返回年   2019
返回月   May
返回日   29
返回时   3
返回分   1
返回秒   28
返回星期几   Wednesday
Local
UTC 0
使用IN函数，传入参数loc 返回时间  Wed May 29 11:01:28 2019
2019-05-29 03:00:00 +0000 UTC 2019-05-29 03:00:00 +0000 UTC 1.5 90 5400 1h30m0s
13h36m39.2454423s
2019-05-28 13:25:01 +0000 UTC
计算时间差 dt与当前时间的间隔  -13h36m39.2456252s
true
false
Golang1
ninhao +  1
Golang
Expiration time: 2019-05-29 03:01:37.2472702 +0000 UTC m=+9.003449101.
false
get ticker1 2019-05-29 03:01:39
get timer 2019-05-29 03:01:39
get timer 2019-05-29 03:01:41
get ticker1 2019-05-29 03:01:41
get ticker1 2019-05-29 03:01:43
get timer 2019-05-29 03:01:43
get ticker1 2019-05-29 03:01:45
get timer 2019-05-29 03:01:45
^Csignal: interrupt
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage#
*/
