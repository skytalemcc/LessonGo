package main

/*
net包提供了可移植的网络I/O接口，包括TCP/IP、UDP、域名解析和Unix域socket。
虽然本包提供了对网络原语的访问，大部分使用者只需要Dial、Listen和Accept函数提供的基本接口；
以及相关的Conn和Listener接口。crypto/tls包提供了相同的接口和类似的Dial和Listen函数。
常用：
const (
    IPv4len = 4  //IP address lengths (bytes).
    IPv6len = 16
)
var (
    IPv4bcast     = IPv4(255, 255, 255, 255) // 广播地址
    IPv4allsys    = IPv4(224, 0, 0, 1)       // 所有主机和路由器
    IPv4allrouter = IPv4(224, 0, 0, 2)       // 所有路由器
    IPv4zero      = IPv4(0, 0, 0, 0)         // 本地地址，只能作为源地址（曾用作广播地址）
)
const (
    FlagUp           Flags = 1 << iota // 接口在活动状态
    FlagBroadcast                      // 接口支持广播
    FlagLoopback                       // 接口是环回的
    FlagPointToPoint                   // 接口是点对点的
    FlagMulticast                      // 接口支持组播
)
type Interface struct {
    Index        int          // 索引，>=1的整数
    MTU          int          // 最大传输单元
    Name         string       // 接口名，例如"en0"、"lo0"、"eth0.100"
    HardwareAddr HardwareAddr // 硬件地址，IEEE MAC-48、EUI-48或EUI-64格式
    Flags        Flags        // 接口的属性，例如FlagUp、FlagLoopback、FlagMulticast
}
Interface类型代表一个网络接口（系统与网络的一个接点）。包含接口索引到名字的映射，也包含接口的设备信息。
func InterfaceByIndex(index int) (*Interface, error)  InterfaceByIndex返回指定索引的网络接口。
func InterfaceByName(name string) (*Interface, error) InterfaceByName返回指定名字的网络接口。
func (ifi *Interface) Addrs() ([]Addr, error)  Addrs方法返回网络接口ifi的一或多个接口地址。
func (ifi *Interface) MulticastAddrs() ([]Addr, error) MulticastAddrs返回网络接口ifi加入的多播组地址。
func Interfaces() ([]Interface, error)  Interfaces返回该系统的网络接口列表。
func InterfaceAddrs() ([]Addr, error) InterfaceAddrs返回该系统的网络接口的地址列表。

type Addr interface {
    Network() string // 网络名
    String() string  // 字符串格式的地址
}
type Conn interface {
    // Read从连接中读取数据
    // Read方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
    Read(b []byte) (n int, err error)
    // Write从连接中写入数据
    // Write方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
    Write(b []byte) (n int, err error)
    // Close方法关闭该连接
    // 并会导致任何阻塞中的Read或Write方法不再阻塞并返回错误
    Close() error
    // 返回本地网络地址
    LocalAddr() Addr
    // 返回远端网络地址
    RemoteAddr() Addr
    // 设定该连接的读写deadline，等价于同时调用SetReadDeadline和SetWriteDeadline
    // deadline是一个绝对时间，超过该时间后I/O操作就会直接因超时失败返回而不会阻塞
    // deadline对之后的所有I/O操作都起效，而不仅仅是下一次的读或写操作
    // 参数t为零值表示不设置期限
    SetDeadline(t time.Time) error
    // 设定该连接的读操作deadline，参数t为零值表示不设置期限
    SetReadDeadline(t time.Time) error
    // 设定该连接的写操作deadline，参数t为零值表示不设置期限
    // 即使写入超时，返回值n也可能>0，说明成功写入了部分数据
    SetWriteDeadline(t time.Time) error
}
Conn接口代表通用的面向流的网络连接。多个线程可能会同时调用同一个Conn的方法。


*/
import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func connserver() {
	conn, err := net.Dial("tcp", "fund.eastmoney.com:8080")
	if err != nil {
		// handle error
	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status, err)
	defer conn.Close()
}

func buildserver() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		_, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}

		//go handleConnection(conn)
	}
}

func main() {
	connserver()
	fmt.Println(net.SplitHostPort("127.0.0.1:8080"))   //func SplitHostPort(hostport string) (host, port string, err error) 函数将格式为"host:port"、"[host]:port"或"[ipv6-host%zone]:port"的网络地址分割为host或ipv6-host%zone和port两个部分。
	fmt.Println(net.JoinHostPort("127.0.0.1", "8080")) //func JoinHostPort(host, port string) string  函数将host和port合并为一个网络地址。一般格式为"host:port"；如果host含有冒号或百分号，格式为"[host]:port"。

	//type HardwareAddr []byte HardwareAddr类型代表一个硬件地址（MAC地址）。func (a HardwareAddr) String() string
	fmt.Println(net.ParseMAC("00-50-56-C0-00-01")) //ParseMAC函数使用如下格式解析一个IEEE 802 MAC-48、EUI-48或EUI-64硬件地址 func ParseMAC(s string) (hw HardwareAddr, err error)

	/*
		type IP []byte IP类型是代表单个IP地址的[]byte切片。本包的函数都可以接受4字节（IPv4）和16字节（IPv6）的切片作为输入。注意，IP地址是IPv4地址还是IPv6地址是语义上的属性，而不取决于切片的长度：16字节的切片也可以是IPv4地址。

	*/
	fmt.Println(net.IPv4(10, 253, 0, 119))                               //func IPv4(a, b, c, d byte) IP IPv4返回包含一个IPv4地址a.b.c.d的IP地址（16字节格式）。
	fmt.Println(net.ParseIP("10.253.0.118"))                             // func ParseIP(s string) IP  ParseIP将s解析为IP地址，并返回该地址。如果s不是合法的IP地址文本表示，ParseIP会返回nil。字符串可以是小数点分隔的IPv4格式（如"74.125.19.99"）或IPv6格式（如"2001:4860:0:2001::68"）格式。
	fmt.Println(net.ParseIP("10.253.0.118").IsGlobalUnicast())           //func (ip IP) IsGlobalUnicast() bool  如果ip是全局单播地址，则返回真。
	fmt.Println(net.ParseIP("10.253.0.118").IsLinkLocalUnicast())        //func (ip IP) IsLinkLocalUnicast() bool  如果ip是链路本地单播地址，则返回真。
	fmt.Println(net.ParseIP("10.253.0.118").IsInterfaceLocalMulticast()) //func (ip IP) IsInterfaceLocalMulticast() bool  如果ip是接口本地组播地址，则返回真。
	fmt.Println(net.ParseIP("10.253.0.118").IsLinkLocalMulticast())      //func (ip IP) IsLinkLocalMulticast() bool 如果ip是链路本地组播地址，则返回真。
	fmt.Println(net.ParseIP("10.253.0.118").IsMulticast())               //func (ip IP) IsMulticast() bool  如果ip是组播地址，则返回真。
	fmt.Println(net.ParseIP("10.253.0.118").IsLoopback())                //func (ip IP) IsLoopback() bool  如果ip是环回地址，则返回真。
	fmt.Println(net.ParseIP("10.253.0.118").IsUnspecified())             //func (ip IP) IsUnspecified() bool  如果ip是未指定地址，则返回真。
	fmt.Println(net.ParseIP("10.253.0.118").DefaultMask())               //func (ip IP) DefaultMask() IPMask  函数返回IP地址ip的默认子网掩码。只有IPv4有默认子网掩码；如果ip不是合法的IPv4地址，会返回nil。
	//func (ip IP) Equal(x IP) bool 如果ip和x代表同一个IP地址，Equal会返回真。代表同一地址的IPv4地址和IPv6地址也被认为是相等的。
	//func (ip IP) To16() IP To16将一个IP地址转换为16字节表示。如果ip不是一个IP地址（长度错误），To16会返回nil。
	//func (ip IP) To4() IP To4将一个IPv4地址转换为4字节表示。如果ip不是IPv4地址，To4会返回nil。
	//func (ip IP) Mask(mask IPMask) IP Mask方法认为mask为ip的子网掩码，返回ip的网络地址部分的ip。（主机地址部分都置0）
	//func (ip IP) String() string  String返回IP地址ip的字符串表示。如果ip是IPv4地址，返回值的格式为点分隔的，如"74.125.19.99"；否则表示为IPv6格式，如"2001:4860:0:2001::68"。
	//func (ip IP) MarshalText() ([]byte, error) MarshalText实现了encoding.TextMarshaler接口，返回值和String方法一样。
	//func (ip *IP) UnmarshalText(text []byte) error UnmarshalText实现了encoding.TextUnmarshaler接口。IP地址字符串应该是ParseIP函数可以接受的格式。
	//type IPMask []byte IPMask代表一个IP地址的掩码。
	//func IPv4Mask(a, b, c, d byte) IPMask IPv4Mask返回一个4字节格式的IPv4掩码a.b.c.d。
	//func CIDRMask(ones, bits int) IPMask CIDRMask返回一个IPMask类型值，该返回值总共有bits个字位，其中前ones个字位都是1，其余字位都是0。
	//func (m IPMask) Size() (ones, bits int) Size返回m的前导的1字位数和总字位数。如果m不是规范的子网掩码（字位：/^1+0+$/），将返会(0, 0)。
	//func (m IPMask) String() string  String返回m的十六进制格式，没有标点。
	/*
			type IPNet struct {  //IPNet表示一个IP网络
		    IP   IP     // 网络地址
		    Mask IPMask // 子网掩码
			}
			func ParseCIDR(s string) (IP, *IPNet, error)ParseCIDR将s作为一个CIDR（无类型域间路由）的IP地址和掩码字符串，如"192.168.100.1/24"或"2001:DB8::/48"，解析并返回IP地址和IP网络。
			本函数会返回IP地址和该IP所在的网络和掩码。例如，ParseCIDR("192.168.100.1/16")会返回IP地址192.168.100.1和IP网络192.168.0.0/16。
	*/
	fmt.Println(net.ParseCIDR("10.116.34.3/25"))
	_, ipnet, _ := net.ParseCIDR("10.116.34.3/25")
	fmt.Println(ipnet.Contains([]byte{10, 116, 34, 5})) //Contains报告该网络是否包含地址ip。func (n *IPNet) Contains(ip IP) bool
	fmt.Println(ipnet.Network())                        //func (n *IPNet) Network() string  Network返回网络类型名："ip+net"，注意该类型名是不合法的。
	fmt.Println(ipnet.String())                         //String返回n的CIDR表示，如"192.168.100.1/24"或"2001:DB8::/48"。

	/*
		func Dial(network, address string) (Conn, error)
		在网络network上连接地址address，并返回一个Conn接口。可用的网络类型有：
		"tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6"、"ip"、"ip4"、"ip6"、"unix"、"unixgram"、"unixpacket"
		对TCP和UDP网络，地址格式是host:port或[host]:port
		Dial("tcp", "12.34.56.78:80")
		Dial("tcp", "google.com:http")
		Dial("tcp", "[2001:db8::1]:http")
		Dial("tcp", "[fe80::1%lo0]:80")
		对IP网络，network必须是"ip"、"ip4"、"ip6"后跟冒号和协议号或者协议名，地址必须是IP地址字面值。
		Dial("ip4:1", "127.0.0.1") 对Unix网络，地址必须是文件系统路径。
		Dial("ip6:ospf", "::1")
	*/
	conn, _ := net.Dial("tcp", "fund.eastmoney.com:80")
	defer conn.Close()
	conn2, err := net.DialTimeout("tcp", "fund.eastmoney.com:8081", 2*time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn2.Close() //如果没有建立连接 而使用close的话，会控台抛出错误 。所以最好在conn里面抛出来。
	/*
		func Pipe() (Conn, Conn) Pipe创建一个内存中的同步、全双工网络连接。连接的两端都实现了Conn接口。一端的读取对应另一端的写入，直接将数据在两端之间作拷贝；没有内部缓冲。
	*/

}
