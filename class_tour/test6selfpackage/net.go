package main

/*
方法太多 ，建议有需要的时候 直接现查。
net包提供了可移植的网络I/O接口，包括TCP/IP、UDP、域名解析和Unix域socket。
虽然本包提供了对网络原语的访问，大部分使用者只需要Dial、Listen和Accept函数提供的基本接口；
以及相关的Conn和Listener接口。crypto/tls包提供了相同的接口和类似的Dial和Listen函数。

实际上dial.go这个文件中并没有实际发起连接的部分，基本上是在为真正发起连接做一系列的准备，
比如：解析网络类型、从addr解析ip地址。。。实际发起连接的函数在tcpsock_posix.go、udpsock_posix.go。
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

type PacketConn interface {
    // ReadFrom方法从连接读取一个数据包，并将有效信息写入b
    // ReadFrom方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
    // 返回写入的字节数和该数据包的来源地址
    ReadFrom(b []byte) (n int, addr Addr, err error)
    // WriteTo方法将有效数据b写入一个数据包发送给addr
    // WriteTo方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
    // 在面向数据包的连接中，写入超时非常罕见
    WriteTo(b []byte, addr Addr) (n int, err error)
    // Close方法关闭该连接
    // 会导致任何阻塞中的ReadFrom或WriteTo方法不再阻塞并返回错误
    Close() error
    // 返回本地网络地址
    LocalAddr() Addr
    // 设定该连接的读写deadline
    SetDeadline(t time.Time) error
    // 设定该连接的读操作deadline，参数t为零值表示不设置期限
    // 如果时间到达deadline，读操作就会直接因超时失败返回而不会阻塞
    SetReadDeadline(t time.Time) error
    // 设定该连接的写操作deadline，参数t为零值表示不设置期限
    // 如果时间到达deadline，写操作就会直接因超时失败返回而不会阻塞
    // 即使写入超时，返回值n也可能>0，说明成功写入了部分数据
    SetWriteDeadline(t time.Time) error
}
PacketConn接口代表通用的面向数据包的网络连接。多个线程可能会同时调用同一个Conn的方法。
func ListenPacket(net, laddr string) (PacketConn, error) ListenPacket函数监听本地网络地址laddr。网络类型net必须是面向数据包的网络类型："ip"、"ip4"、"ip6"、"udp"、"udp4"、"udp6"、或"unixgram"。
type Dialer struct {
    // Timeout是dial操作等待连接建立的最大时长，默认值代表没有超时。
    // 如果Deadline字段也被设置了，dial操作也可能更早失败。
    // 不管有没有设置超时，操作系统都可能强制执行它的超时设置。
    // 例如，TCP（系统）超时一般在3分钟左右。
    Timeout time.Duration
    // Deadline是一个具体的时间点期限，超过该期限后，dial操作就会失败。
    // 如果Timeout字段也被设置了，dial操作也可能更早失败。
    // 零值表示没有期限，即遵守操作系统的超时设置。
    Deadline time.Time
    // LocalAddr是dial一个地址时使用的本地地址。
    // 该地址必须是与dial的网络相容的类型。
    // 如果为nil，将会自动选择一个本地地址。
    LocalAddr Addr  //真正dial时的本地地址，兼容各种类型(TCP、UDP...),如果为nil，则系统自动选择一个地址
    // DualStack允许单次dial操作在网络类型为"tcp"，
    // 且目的地是一个主机名的DNS记录具有多个地址时，
    // 尝试建立多个IPv4和IPv6连接，并返回第一个建立的连接。
    DualStack bool //双协议栈，即是否同时支持ipv4和ipv6.当network值为tcp时，dial函数会向host主机的v4和v6地址都发起连接
    // KeepAlive指定一个活动的网络连接的生命周期；如果为0，会禁止keep-alive。
    // 不支持keep-alive的网络连接会忽略本字段。
    KeepAlive time.Duration
}

Examples:
Dial("tcp", "golang.org:http")
Dial("tcp", "192.0.2.1:http")
Dial("tcp", "198.51.100.1:80")
Dial("udp", "[2001:db8::1]:domain")
Dial("udp", "[fe80::1%lo0]:53")
Dial("tcp", ":80")

Dialer类型包含与某个地址建立连接时的参数。

每一个字段的零值都等价于没有该字段。因此调用Dialer零值的Dial方法等价于调用Dial函数。
func (d *Dialer) Dial(network, address string) (Conn, error)  Dial在指定的网络上连接指定的地址。


*/
import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func connserver() { //客户端发起连接到服务端
	conn, err := net.Dial("tcp", "fund.eastmoney.com:8080")
	if err != nil {
		// handle error
	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status, err)
	defer conn.Close()
}

func buildserver() { //服务端提供连接等待客户端访问
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
	conn2, err := net.DialTimeout("tcp", "fund.eastmoney.com:8080", 2*time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn2.Close() //如果没有建立连接 而使用close的话，会控台抛出错误 。所以最好在conn里面抛出来。
	/*
		func Pipe() (Conn, Conn) Pipe创建一个内存中的同步、全双工网络连接。连接的两端都实现了Conn接口。一端的读取对应另一端的写入，直接将数据在两端之间作拷贝；没有内部缓冲。
	*/
	/*
			type Listener interface {
				// Addr返回该接口的网络地址
				Addr() Addr
				// Accept等待并返回下一个连接到该接口的连接
				Accept() (c Conn, err error)
				// Close关闭该接口，并使任何阻塞的Accept操作都会不再阻塞并返回错误。
				Close() error
			}

			Listener是一个用于面向流的网络协议的公用的网络监听器接口。多个线程可能会同时调用一个Listener的方法。
			我们用这个方法来创建服务端 ，对外提供连接服务 来做一些远程调用的网络通信，提供某些返回给客户端来进行通信。
			func Listen(net, laddr string) (Listener, error) 返回在一个本地网络地址laddr上监听的Listener。网络类型参数net必须是面向流的网络："tcp"、"tcp4"、"tcp6"、"unix"或"unixpacket"。

			type IPAddr struct {  //IPAddr代表一个IP终端的地址。
		    IP   IP
		    Zone string // IPv6范围寻址域
			}

			func ResolveIPAddr(net, addr string) (*IPAddr, error) ResolveIPAddr将addr作为一个格式为"host"或"ipv6-host%zone"的IP地址来解析。 函数会在参数net指定的网络类型上解析，net必须是"ip"、"ip4"或"ip6"。
			func (a *IPAddr) Network() string Network返回地址的网络类型："ip"。
			type TCPAddr struct {
		    IP   IP
		    Port int
		    Zone string // IPv6范围寻址域
			}TCPAddr代表一个TCP终端地址。

			func ResolveTCPAddr(net, addr string) (*TCPAddr, error) ResolveTCPAddr将addr作为TCP地址解析并返回。参数addr格式为"host:port"或"[ipv6-host%zone]:port"，解析得到网络名和端口名；net必须是"tcp"、"tcp4"或"tcp6"。
			func (a *TCPAddr) Network() string 返回地址的网络类型，"tcp"。
			type UDPAddr struct {
		    IP   IP
		    Port int
		    Zone string // IPv6范围寻址域
			} UDPAddr代表一个UDP终端地址。
			func ResolveUDPAddr(net, addr string) (*UDPAddr, error) ResolveTCPAddr将addr作为TCP地址解析并返回。参数addr格式为"host:port"或"[ipv6-host%zone]:port"，解析得到网络名和端口名；net必须是"udp"、"udp4"或"udp6"。
			func (a *UDPAddr) Network() string  返回地址的网络类型，"udp"。

			type UnixAddr struct {
		    Name string
		    Net  string
			} UnixAddr代表一个Unix域socket终端地址。
			func ResolveUnixAddr(net, addr string) (*UnixAddr, error) ResolveUnixAddr将addr作为Unix域socket地址解析，参数net指定网络类型："unix"、"unixgram"或"unixpacket"。

			func (a *UnixAddr) Network() string 返回地址的网络类型，"unix"，"unixgram"或"unixpacket"。
			type IPConn struct {
		    // 内含隐藏或非导出字段
			} IPConn类型代表IP网络连接，实现了Conn和PacketConn接口。
			func DialIP(netProto string, laddr, raddr *IPAddr) (*IPConn, error) DialIP在网络协议netProto上连接本地地址laddr和远端地址raddr，netProto必须是"ip"、"ip4"或"ip6"后跟冒号和协议名或协议号。
			func ListenIP(netProto string, laddr *IPAddr) (*IPConn, error) ListenIP创建一个接收目的地是本地地址laddr的IP数据包的网络连接，返回的*IPConn的ReadFrom和WriteTo方法可以用来发送和接收IP数据包。
			func (c *IPConn) LocalAddr() Addr LocalAddr返回本地网络地址
			func (c *IPConn) RemoteAddr() Addr RemoteAddr返回远端网络地址
			func (c *IPConn) SetReadBuffer(bytes int) error SetReadBuffer设置该连接的系统接收缓冲
			func (c *IPConn) SetWriteBuffer(bytes int) error  SetWriteBuffer设置该连接的系统发送缓冲
			func (c *IPConn) SetDeadline(t time.Time) error  SetDeadline设置读写操作绝对期限，实现了Conn接口的SetDeadline方法
			func (c *IPConn) SetReadDeadline(t time.Time) error  SetReadDeadline设置读操作绝对期限，实现了Conn接口的SetReadDeadline方法
			func (c *IPConn) SetWriteDeadline(t time.Time) error SetWriteDeadline设置写操作绝对期限，实现了Conn接口的SetWriteDeadline方法
			func (c *IPConn) Read(b []byte) (int, error)  Read实现Conn接口Read方法
			func (c *IPConn) ReadFrom(b []byte) (int, Addr, error) ReadFrom实现PacketConn接口ReadFrom方法。注意本方法有bug，应避免使用。
			func (c *IPConn) ReadFromIP(b []byte) (int, *IPAddr, error)  ReadFromIP从c读取一个IP数据包，将有效负载拷贝到b，返回拷贝字节数和数据包来源地址。ReadFromIP方法会在超过一个固定的时间点之后超时，并返回一个错误。注意本方法有bug，应避免使用。
			func (c *IPConn) ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error) ReadMsgIP从c读取一个数据包，将有效负载拷贝进b，相关的带外数据拷贝进oob，返回拷贝进b的字节数，拷贝进oob的字节数，数据包的flag，数据包来源地址和可能的错误。
			func (c *IPConn) Write(b []byte) (int, error) Write实现Conn接口Write方法
			func (c *IPConn) WriteTo(b []byte, addr Addr) (int, error) WriteTo实现PacketConn接口WriteTo方法
			func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error) WriteToIP通过c向地址addr发送一个数据包，b为包的有效负载，返回写入的字节。WriteToIP方法会在超过一个固定的时间点之后超时，并返回一个错误。在面向数据包的连接上，写入超时是十分罕见的。
			func (c *IPConn) WriteMsgIP(b, oob []byte, addr *IPAddr) (n, oobn int, err error) WriteMsgIP通过c向地址addr发送一个数据包，b和oob分别为包有效负载和对应的带外数据，返回写入的字节数（包数据、带外数据）和可能的错误。
			func (c *IPConn) Close() error Close关闭连接
			func (c *IPConn) File() (f *os.File, err error) File方法设置下层的os.File为阻塞模式并返回其副本。使用者有责任在用完后关闭f。关闭c不影响f，关闭f也不影响c。返回的os.File类型文件描述符和原本的网络连接是不同的。试图使用该副本修改本体的属性可能会（也可能不会）得到期望的效果。
			type TCPConn struct {
		    // 内含隐藏或非导出字段
			} TCPConn代表一个TCP网络连接，实现了Conn接口。
			func DialTCP(net string, laddr, raddr *TCPAddr) (*TCPConn, error) DialTCP在网络协议net上连接本地地址laddr和远端地址raddr。net必须是"tcp"、"tcp4"、"tcp6"；如果laddr不是nil，将使用它作为本地地址，否则自动选择一个本地地址。
			func (c *TCPConn) LocalAddr() Addr LocalAddr返回本地网络地址
			func (c *TCPConn) RemoteAddr() Addr RemoteAddr返回远端网络地址
			func (c *TCPConn) SetReadBuffer(bytes int) error  SetReadBuffer设置该连接的系统接收缓冲
			func (c *TCPConn) SetWriteBuffer(bytes int) error SetWriteBuffer设置该连接的系统发送缓冲
			func (c *TCPConn) SetDeadline(t time.Time) error SetDeadline设置读写操作期限，实现了Conn接口的SetDeadline方法
			func (c *TCPConn) SetReadDeadline(t time.Time) error SetReadDeadline设置读操作期限，实现了Conn接口的SetReadDeadline方法
			func (c *TCPConn) SetWriteDeadline(t time.Time) error SetWriteDeadline设置写操作期限，实现了Conn接口的SetWriteDeadline方法
			func (c *TCPConn) SetKeepAlive(keepalive bool) error SetKeepAlive设置操作系统是否应该在该连接中发送keepalive信息
			func (c *TCPConn) SetKeepAlivePeriod(d time.Duration) error SetKeepAlivePeriod设置keepalive的周期，超出会断开
			func (c *TCPConn) SetLinger(sec int) error SetLinger设定当连接中仍有数据等待发送或接受时的Close方法的行为。 如果sec < 0（默认），Close方法立即返回，操作系统停止后台数据发送；如果 sec == 0，Close立刻返回，操作系统丢弃任何未发送或未接收的数据；如果sec > 0，Close方法阻塞最多sec秒，等待数据发送或者接收，在一些操作系统中，在超时后，任何未发送的数据会被丢弃。
			func (c *TCPConn) SetNoDelay(noDelay bool) error SetNoDelay设定操作系统是否应该延迟数据包传递，以便发送更少的数据包（Nagle's算法）。默认为真，即数据应该在Write方法后立刻发送。
			func (c *TCPConn) Read(b []byte) (int, error) Read实现了Conn接口Read方法
			func (c *TCPConn) Write(b []byte) (int, error) Write实现了Conn接口Write方法
			func (c *TCPConn) ReadFrom(r io.Reader) (int64, error) ReadFrom实现了io.ReaderFrom接口的ReadFrom方法
			func (c *TCPConn) Close() error Close关闭连接
			func (c *TCPConn) CloseRead() error CloseRead关闭TCP连接的读取侧（以后不能读取），应尽量使用Close方法。
			func (c *TCPConn) CloseWrite() error CloseWrite关闭TCP连接的写入侧（以后不能写入），应尽量使用Close方法。
			func (c *TCPConn) File() (f *os.File, err error) File方法设置下层的os.File为阻塞模式并返回其副本。 使用者有责任在用完后关闭f。关闭c不影响f，关闭f也不影响c。返回的os.File类型文件描述符和原本的网络连接是不同的。试图使用该副本修改本体的属性可能会（也可能不会）得到期望的效果。
			type UDPConn struct {
		    // 内含隐藏或非导出字段
			}UDPConn代表一个UDP网络连接，实现了Conn和PacketConn接口。
			func DialUDP(net string, laddr, raddr *UDPAddr) (*UDPConn, error) DialTCP在网络协议net上连接本地地址laddr和远端地址raddr。net必须是"udp"、"udp4"、"udp6"；如果laddr不是nil，将使用它作为本地地址，否则自动选择一个本地地址。
			func ListenUDP(net string, laddr *UDPAddr) (*UDPConn, error) ListenUDP创建一个接收目的地是本地地址laddr的UDP数据包的网络连接。net必须是"udp"、"udp4"、"udp6"；如果laddr端口为0，函数将选择一个当前可用的端口，可以用Listener的Addr方法获得该端口。返回的*UDPConn的ReadFrom和WriteTo方法可以用来发送和接收UDP数据包（每个包都可获得来源地址或设置目标地址）。
			func (c *UDPConn) Read(b []byte) (int, error) Read实现Conn接口Read方法
			func (c *UDPConn) Write(b []byte) (int, error) Write实现Conn接口Write方法
			func (c *UDPConn) Close() error Close关闭连接

			type UnixConn struct {
		    // 内含隐藏或非导出字段
			}UnixConn代表Unix域socket连接，实现了Conn和PacketConn接口。
			func DialUnix(net string, laddr, raddr *UnixAddr) (*UnixConn, error) DialUnix在网络协议net上连接本地地址laddr和远端地址raddr。net必须是"unix"、"unixgram"、"unixpacket"，如果laddr不是nil将使用它作为本地地址，否则自动选择一个本地地址。
			func ListenUnixgram(net string, laddr *UnixAddr) (*UnixConn, error) ListenUnixgram接收目的地是本地地址laddr的Unix datagram网络连接。net必须是"unixgram"，返回的*UnixConn的ReadFrom和WriteTo方法可以用来发送和接收数据包（每个包都可获取来源址或者设置目标地址）。
			func (c *UnixConn) LocalAddr() Addr  LocalAddr返回本地网络地址
			func (c *UnixConn) RemoteAddr() Addr RemoteAddr返回远端网络地址
			func (c *UnixConn) Read(b []byte) (int, error)  Read实现了Conn接口Read方法
			func (c *UnixConn) Write(b []byte) (int, error) Write实现了Conn接口Write方法
			func (c *UnixConn) Close() error   Close关闭连接
			func (c *UnixConn) CloseRead() error CloseRead关闭TCP连接的读取侧（以后不能读取），应尽量使用Close方法
			func (c *UnixConn) CloseWrite() error CloseWrite关闭TCP连接的写入侧（以后不能写入），应尽量使用Close方法
			type TCPListener struct {
		    // 内含隐藏或非导出字段
			} TCPListener代表一个TCP网络的监听者。使用者应尽量使用Listener接口而不是假设（网络连接为）TCP。
			func ListenTCP(net string, laddr *TCPAddr) (*TCPListener, error) ListenTCP在本地TCP地址laddr上声明并返回一个*TCPListener，net参数必须是"tcp"、"tcp4"、"tcp6"，如果laddr的端口字段为0，函数将选择一个当前可用的端口，可以用Listener的Addr方法获得该端口。
			func (l *TCPListener) Addr() Addr Addr返回l监听的的网络地址，一个*TCPAddr。
			func (l *TCPListener) SetDeadline(t time.Time) error 设置监听器执行的期限，t为Time零值则会关闭期限限制。
			func (l *TCPListener) Accept() (Conn, error) Accept用于实现Listener接口的Accept方法；他会等待下一个呼叫，并返回一个该呼叫的Conn接口。
			func (l *TCPListener) AcceptTCP() (*TCPConn, error) AcceptTCP接收下一个呼叫，并返回一个新的*TCPConn。
			func (l *TCPListener) Close() error Close停止监听TCP地址，已经接收的连接不受影响。
			type UnixListener struct {
		    // 内含隐藏或非导出字段
			}UnixListener代表一个Unix域scoket的监听者。使用者应尽量使用Listener接口而不是假设（网络连接为）Unix域scoket。
			func ListenUnix(net string, laddr *UnixAddr) (*UnixListener, error) ListenTCP在Unix域scoket地址laddr上声明并返回一个*UnixListener，net参数必须是"unix"或"unixpacket"。
			func (l *UnixListener) Accept() (c Conn, err error) Accept用于实现Listener接口的Accept方法；他会等待下一个呼叫，并返回一个该呼叫的Conn接口。
			func (l *UnixListener) Close() error Close停止监听Unix域socket地址，已经接收的连接不受影响。
			func FileConn(f *os.File) (c Conn, err error) FileConn返回一个下层为文件f的网络连接的拷贝。调用者有责任在结束程序前关闭f。关闭c不会影响f，关闭f也不会影响c。本函数与各种实现了Conn接口的类型的File方法是对应的。
			func FileListener(f *os.File) (l Listener, err error) FileListener返回一个下层为文件f的网络监听器的拷贝。调用者有责任在使用结束后改变l。关闭l不会影响f，关闭f也不会影响l。本函数与各种实现了Listener接口的类型的File方法是对应的。
			type MX struct {
		    Host string
		    Pref uint16
			} MX代表一条DNS MX记录（邮件交换记录），根据收信人的地址后缀来定位邮件服务器。
			type NS struct {
		    Host string
			}NS代表一条DNS NS记录（域名服务器记录），指定该域名由哪个DNS服务器来进行解析。
			type SRV struct {
		    Target   string
		    Port     uint16
		    Priority uint16
		    Weight   uint16
			}SRV代表一条DNS SRV记录（资源记录），记录某个服务由哪台计算机提供。

	*/
	fmt.Println(net.LookupPort("fund.eastmoney.com", "8080")) //func LookupPort(network, service string) (port int, err error) LookupPort函数查询指定网络和服务的（默认）端口。
	fmt.Println(net.LookupHost("e7939faf8694"))               //func LookupHost(host string) (addrs []string, err error)  LookupHost函数查询主机的网络地址序列。
	fmt.Println(net.LookupIP("e7939faf8694"))                 //LookupIP函数查询主机的ipv4和ipv6地址序列。
	//func LookupAddr(addr string) (name []string, err error) LookupAddr查询某个地址，返回映射到该地址的主机名序列，本函数和LookupHost不互为反函数。

}

/*
结果集：
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# go run net.go
HTTP/1.1 200 OK
 <nil>
127.0.0.1 8080 <nil>
127.0.0.1:8080
00:50:56:c0:00:01 <nil>
10.253.0.119
10.253.0.118
true
false
false
false
false
false
false
ff000000
10.116.34.3 10.116.34.0/25 <nil>
true
ip+net
10.116.34.0/25
8080 <nil>
[172.17.0.2] <nil>
[172.17.0.2] <nil>
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage#
*/
