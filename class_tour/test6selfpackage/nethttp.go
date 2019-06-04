package main

/***
对于很多现代应用来说，访问互联网上的信息和访问本地文件系统一样重要。
使用net包可以更简单地用网络收发信息，还可以建立更底层的网络连接，编写服务器程序。
http包提供了HTTP客户端和服务端的实现。
Get、Head、Post和PostForm函数发出HTTP/ HTTPS请求。
***/

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func curl() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url) //http.Get函数是创建HTTP请求的函数  resp这个结构体中得到访问的请求结果
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			//break和continue语句会改变控制流。os.Exit(1)则是退出。
			//和其它语言中的break和continue一样，break会中断当前的循环，并开始执行循环之后的内容，而continue会跳过当前循环，并开始执行下一次循环。
			os.Exit(1) //异常的时候 进行退出
		}
		b, err := ioutil.ReadAll(resp.Body) //resp的Body字段包括一个可读的服务器响应流  ioutil.ReadAll函数从response中读取到全部内容
		resp.Body.Close()                   //resp.Body.Close关闭resp的Body流，防止资源泄露
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("The context of fetching are %s", b)
	}
}

//搭建一个web service使用自建库 一个简单的服务端例子
func webservice() {
	http.HandleFunc("/", handler) // 每一个请求都会发给handler来进行业务处理
	// http.HandleFunc("/count", count) 可以添加别的请求，根据请求的url不同会调用不同的函数
	http.ListenAndServe("localhost:8000", nil) //启动服务，占用端口 。最后外部要套一层 log最好，出现异常就捕获退出。

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path) //通过request 来获得访问路径, 标准输出流的fmt.Fprintf作为返回ResponseWriter
	//这里我可以定义任何一种返回方式，例如返回数字，字符，图片，文件内容等等。这样curl访问地址的时候 就能进行解析
}

func nethttp() {
	/*Get、Head、Post和PostForm函数发出HTTP/ HTTPS请求。
		resp, err := http.Get("http://example.com/")
		resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf) //func Post(url, contentType string, body io.Reader) (resp *Response, err error)
		resp, err := http.PostForm("http://example.com/form",url.Values{"key": {"Value"}, "id": {"123"}})

		程序在使用完回复后必须关闭回复的主体。
		if err != nil {
		// handle error
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		要管理HTTP客户端的头域、重定向策略和其他设置，创建一个Client： 相当于封装了一层再去发起请求。对request的内容进行了封装。
		client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
		}
		resp, err := client.Get("http://example.com")

		req, err := http.NewRequest("GET", "http://example.com", nil)  //func NewRequest(method, url string, body io.Reader) (*Request, error)

		req.Header.Add("If-None-Match", `W/"wyzzy"`)
		resp, err := client.Do(req)

		要管理代理、TLS配置、keep-alive、压缩和其他设置，创建一个Transport： 也是相当于在request的内容进行了封装，加了一些请求需要的内容。
		tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
		}
		client := &http.Client{Transport: tr}
		resp, err := client.Get("https://example.com")
		Client和Transport类型都可以安全的被多个go程同时使用。出于效率考虑，应该一次建立、尽量重用。
		ListenAndServe使用指定的监听地址和处理器启动一个HTTP服务端。处理器参数通常是nil，这表示采用包变量DefaultServeMux作为处理器。Handle和HandleFunc函数可以向DefaultServeMux添加处理器。


		http.Handle("/foo", fooHandler)
		http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		})
		log.Fatal(http.ListenAndServe(":8080", nil))  出现异常捕捉错误 进行退出

		如果要管理服务器的行为，需要创建一个自定义的server：
		s := &http.Server{
		Addr:           ":8080",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		}
		log.Fatal(s.ListenAndServe())

	const (
	    StatusContinue           = 100
	    StatusSwitchingProtocols = 101
	    StatusOK                   = 200
	    StatusCreated              = 201
	    StatusAccepted             = 202
	    StatusNonAuthoritativeInfo = 203
	    StatusNoContent            = 204
	    StatusResetContent         = 205
	    StatusPartialContent       = 206
	    StatusMultipleChoices   = 300
	    StatusMovedPermanently  = 301
	    StatusFound             = 302
	    StatusSeeOther          = 303
	    StatusNotModified       = 304
	    StatusUseProxy          = 305
	    StatusTemporaryRedirect = 307
	    StatusBadRequest                   = 400
	    StatusUnauthorized                 = 401
	    StatusPaymentRequired              = 402
	    StatusForbidden                    = 403
	    StatusNotFound                     = 404
	    StatusMethodNotAllowed             = 405
	    StatusNotAcceptable                = 406
	    StatusProxyAuthRequired            = 407
	    StatusRequestTimeout               = 408
	    StatusConflict                     = 409
	    StatusGone                         = 410
	    StatusLengthRequired               = 411
	    StatusPreconditionFailed           = 412
	    StatusRequestEntityTooLarge        = 413
	    StatusRequestURITooLong            = 414
	    StatusUnsupportedMediaType         = 415
	    StatusRequestedRangeNotSatisfiable = 416
	    StatusExpectationFailed            = 417
	    StatusTeapot                       = 418
	    StatusInternalServerError     = 500
	    StatusNotImplemented          = 501
	    StatusBadGateway              = 502
	    StatusServiceUnavailable      = 503
	    StatusGatewayTimeout          = 504
	    StatusHTTPVersionNotSupported = 505
	)
	*/

}

func main() {
	go webservice()

	fmt.Println(http.StatusText(404)) //func StatusText(code int) string StatusText返回HTTP状态码code对应的文本，如220对应"OK"。如果code是未知的状态码，会返回""。
	//type ConnState int ConnState代表一个客户端到服务端的连接的状态。本类型用于可选的Server.ConnState回调函数。
	/*
				const (
				    // StateNew代表一个新的连接，将要立刻发送请求。
				    // 连接从这个状态开始，然后转变为StateAlive或StateClosed。
				    StateNew ConnState = iota
				    // StateActive代表一个已经读取了请求数据1到多个字节的连接。
				    // 用于StateAlive的Server.ConnState回调函数在将连接交付给处理器之前被触发，
				    // 等到请求被处理完后，Server.ConnState回调函数再次被触发。
				    // 在请求被处理后，连接状态改变为StateClosed、StateHijacked或StateIdle。
				    StateActive
				    // StateIdle代表一个已经处理完了请求、处在闲置状态、等待新请求的连接。
				    // 连接状态可以从StateIdle改变为StateActive或StateClosed。
				    StateIdle
				    // 代表一个被劫持的连接。这是一个终止状态，不会转变为StateClosed。
				    StateHijacked
				    // StateClosed代表一个关闭的连接。
				    // 这是一个终止状态。被劫持的连接不会转变为StateClosed。
				    StateClosed
				)

			func (c ConnState) String() string
			type Header map[string][]string  Header代表HTTP头域的键值对。
			func (h Header) Get(key string) string Get返回键对应的第一个值，如果键不存在会返回""。如要获取该键对应的值切片，请直接用规范格式的键访问map。
			func (h Header) Set(key, value string) Set添加键值对到h，如键已存在则会用只有新值一个元素的切片取代旧值切片。
			func (h Header) Add(key, value string)  Add添加键值对到h，如键已存在则会将新的值附加到旧值切片后面。

			type Cookie struct {
		    Name       string
		    Value      string
		    Path       string
		    Domain     string
		    Expires    time.Time
		    RawExpires string
		    // MaxAge=0表示未设置Max-Age属性
		    // MaxAge<0表示立刻删除该cookie，等价于"Max-Age: 0"
		    // MaxAge>0表示存在Max-Age属性，单位是秒
		    MaxAge   int
		    Secure   bool
		    HttpOnly bool
		    Raw      string
		    Unparsed []string // 未解析的“属性-值”对的原始文本
		}
		func (c *Cookie) String() string String返回该cookie的序列化结果。如果只设置了Name和Value字段，序列化结果可用于HTTP请求的Cookie头或者HTTP回复的Set-Cookie头；如果设置了其他字段，序列化结果只能用于HTTP回复的Set-Cookie头。

	*/
	/*
			type Request struct {
		    // Method指定HTTP方法（GET、POST、PUT等）。对客户端，""代表GET。
		    Method string
		    // URL在服务端表示被请求的URI，在客户端表示要访问的URL。
		    //
		    // 在服务端，URL字段是解析请求行的URI（保存在RequestURI字段）得到的，
		    // 对大多数请求来说，除了Path和RawQuery之外的字段都是空字符串。
		    // （参见RFC 2616, Section 5.1.2）
		    //
		    // 在客户端，URL的Host字段指定了要连接的服务器，
		    // 而Request的Host字段（可选地）指定要发送的HTTP请求的Host头的值。
		    URL *url.URL
		    // 接收到的请求的协议版本。本包生产的Request总是使用HTTP/1.1
		    Proto      string // "HTTP/1.0"
		    ProtoMajor int    // 1
		    ProtoMinor int    // 0
		    // Header字段用来表示HTTP请求的头域。如果头域（多行键值对格式）为：
		    //	accept-encoding: gzip, deflate
		    //	Accept-Language: en-us
		    //	Connection: keep-alive
		    // 则：
		    //	Header = map[string][]string{
		    //		"Accept-Encoding": {"gzip, deflate"},
		    //		"Accept-Language": {"en-us"},
		    //		"Connection": {"keep-alive"},
		    //	}
		    // HTTP规定头域的键名（头名）是大小写敏感的，请求的解析器通过规范化头域的键名来实现这点。
		    // 在客户端的请求，可能会被自动添加或重写Header中的特定的头，参见Request.Write方法。
		    Header Header
		    // Body是请求的主体。
		    //
		    // 在客户端，如果Body是nil表示该请求没有主体买入GET请求。
		    // Client的Transport字段会负责调用Body的Close方法。
		    //
		    // 在服务端，Body字段总是非nil的；但在没有主体时，读取Body会立刻返回EOF。
		    // Server会关闭请求的主体，ServeHTTP处理器不需要关闭Body字段。
		    Body io.ReadCloser
		    // ContentLength记录相关内容的长度。
		    // 如果为-1，表示长度未知，如果>=0，表示可以从Body字段读取ContentLength字节数据。
		    // 在客户端，如果Body非nil而该字段为0，表示不知道Body的长度。
		    ContentLength int64
		    // TransferEncoding按从最外到最里的顺序列出传输编码，空切片表示"identity"编码。
		    // 本字段一般会被忽略。当发送或接受请求时，会自动添加或移除"chunked"传输编码。
		    TransferEncoding []string
		    // Close在服务端指定是否在回复请求后关闭连接，在客户端指定是否在发送请求后关闭连接。
		    Close bool
		    // 在服务端，Host指定URL会在其上寻找资源的主机。
		    // 根据RFC 2616，该值可以是Host头的值，或者URL自身提供的主机名。
		    // Host的格式可以是"host:port"。
		    //
		    // 在客户端，请求的Host字段（可选地）用来重写请求的Host头。
		    // 如过该字段为""，Request.Write方法会使用URL字段的Host。
		    Host string
		    // Form是解析好的表单数据，包括URL字段的query参数和POST或PUT的表单数据。
		    // 本字段只有在调用ParseForm后才有效。在客户端，会忽略请求中的本字段而使用Body替代。
		    Form url.Values
		    // PostForm是解析好的POST或PUT的表单数据。
		    // 本字段只有在调用ParseForm后才有效。在客户端，会忽略请求中的本字段而使用Body替代。
		    PostForm url.Values
		    // MultipartForm是解析好的多部件表单，包括上传的文件。
		    // 本字段只有在调用ParseMultipartForm后才有效。
		    // 在客户端，会忽略请求中的本字段而使用Body替代。
		    MultipartForm *multipart.Form
		    // Trailer指定了会在请求主体之后发送的额外的头域。
		    //
		    // 在服务端，Trailer字段必须初始化为只有trailer键，所有键都对应nil值。
		    // （客户端会声明哪些trailer会发送）
		    // 在处理器从Body读取时，不能使用本字段。
		    // 在从Body的读取返回EOF后，Trailer字段会被更新完毕并包含非nil的值。
		    // （如果客户端发送了这些键值对），此时才可以访问本字段。
		    //
		    // 在客户端，Trail必须初始化为一个包含将要发送的键值对的映射。（值可以是nil或其终值）
		    // ContentLength字段必须是0或-1，以启用"chunked"传输编码发送请求。
		    // 在开始发送请求后，Trailer可以在读取请求主体期间被修改，
		    // 一旦请求主体返回EOF，调用者就不可再修改Trailer。
		    //
		    // 很少有HTTP客户端、服务端或代理支持HTTP trailer。
		    Trailer Header
		    // RemoteAddr允许HTTP服务器和其他软件记录该请求的来源地址，一般用于日志。
		    // 本字段不是ReadRequest函数填写的，也没有定义格式。
		    // 本包的HTTP服务器会在调用处理器之前设置RemoteAddr为"IP:port"格式的地址。
		    // 客户端会忽略请求中的RemoteAddr字段。
		    RemoteAddr string
		    // RequestURI是被客户端发送到服务端的请求的请求行中未修改的请求URI
		    // （参见RFC 2616, Section 5.1）
		    // 一般应使用URI字段，在客户端设置请求的本字段会导致错误。
		    RequestURI string
		    // TLS字段允许HTTP服务器和其他软件记录接收到该请求的TLS连接的信息
		    // 本字段不是ReadRequest函数填写的。
		    // 对启用了TLS的连接，本包的HTTP服务器会在调用处理器之前设置TLS字段，否则将设TLS为nil。
		    // 客户端会忽略请求中的TLS字段。
		    TLS *tls.ConnectionState
			}
			Request类型代表一个服务端接受到的或者客户端发送出去的HTTP请求。 Request各字段的意义和用途在服务端和客户端是不同的。除了字段本身上方文档，还可参见Request.Write方法和RoundTripper接口的文档。

			func NewRequest(method, urlStr string, body io.Reader) (*Request, error) NewRequest使用指定的方法、网址和可选的主题创建并返回一个新的*Request。
			func ReadRequest(b *bufio.Reader) (req *Request, err error) ReadRequest从b读取并解析出一个HTTP请求。（本函数主要用在服务端从下层获取请求）
			func (r *Request) UserAgent() string UserAgent返回请求中的客户端用户代理信息（请求的User-Agent头）。
			func (r *Request) Referer() string Referer返回请求中的访问来路信息。（请求的Referer头）
			func (r *Request) AddCookie(c *Cookie) AddCookie向请求中添加一个cookie。
			func (r *Request) SetBasicAuth(username, password string) SetBasicAuth使用提供的用户名和密码，采用HTTP基本认证，设置请求的Authorization头。HTTP基本认证会明码传送用户名和密码。
			func (r *Request) Write(w io.Writer) error Write方法以有线格式将HTTP/1.1请求写入w（用于将请求写入下层TCPConn等）。
			func (r *Request) ParseForm() error ParseForm解析URL中的查询字符串，并将解析结果更新到r.Form字段。对于POST或PUT请求，ParseForm还会将body当作表单解析，并将结果既更新到r.PostForm也更新到r.Form。解析结果中，POST或PUT请求主体要优先于URL查询字符串。

			type Response struct {
		    Status     string // 例如"200 OK"
		    StatusCode int    // 例如200
		    Proto      string // 例如"HTTP/1.0"
		    ProtoMajor int    // 例如1
		    ProtoMinor int    // 例如0
		    // Header保管头域的键值对。
		    // 如果回复中有多个头的键相同，Header中保存为该键对应用逗号分隔串联起来的这些头的值
		    // （参见RFC 2616 Section 4.2）
		    // 被本结构体中的其他字段复制保管的头（如ContentLength）会从Header中删掉。
		    //
		    // Header中的键都是规范化的，参见CanonicalHeaderKey函数
		    Header Header
		    // Body代表回复的主体。
		    // Client类型和Transport类型会保证Body字段总是非nil的，即使回复没有主体或主体长度为0。
		    // 关闭主体是调用者的责任。
		    // 如果服务端采用"chunked"传输编码发送的回复，Body字段会自动进行解码。
		    Body io.ReadCloser
		    // ContentLength记录相关内容的长度。
		    // 其值为-1表示长度未知（采用chunked传输编码）
		    // 除非对应的Request.Method是"HEAD"，其值>=0表示可以从Body读取的字节数
		    ContentLength int64
		    // TransferEncoding按从最外到最里的顺序列出传输编码，空切片表示"identity"编码。
		    TransferEncoding []string
		    // Close记录头域是否指定应在读取完主体后关闭连接。（即Connection头）
		    // 该值是给客户端的建议，Response.Write方法的ReadResponse函数都不会关闭连接。
		    Close bool
		    // Trailer字段保存和头域相同格式的trailer键值对，和Header字段相同类型
		    Trailer Header
		    // Request是用来获取此回复的请求
		    // Request的Body字段是nil（因为已经被用掉了）
		    // 这个字段是被Client类型发出请求并获得回复后填充的
		    Request *Request
		    // TLS包含接收到该回复的TLS连接的信息。 对未加密的回复，本字段为nil。
		    // 返回的指针是被（同一TLS连接接收到的）回复共享的，不应被修改。
		    TLS *tls.ConnectionState
			} Response代表一个HTTP请求的回复。
			func ReadResponse(r *bufio.Reader, req *Request) (*Response, error) ReadResponse从r读取并返回一个HTTP 回复。req参数是可选的，指定该回复对应的请求（即是对该请求的回复）。如果是nil，将假设请求是GET请求。客户端必须在结束resp.Body的读取后关闭它。读取完毕并关闭后，客户端可以检查resp.Trailer字段获取回复的trailer的键值对。
			func (r *Response) Write(w io.Writer) error Write以有线格式将回复写入w（用于将回复写入下层TCPConn等）。

			type ResponseWriter interface {
		    // Header返回一个Header类型值，该值会被WriteHeader方法发送。
		    // 在调用WriteHeader或Write方法后再改变该对象是没有意义的。
		    Header() Header
		    // WriteHeader该方法发送HTTP回复的头域和状态码。
		    // 如果没有被显式调用，第一次调用Write时会触发隐式调用WriteHeader(http.StatusOK)
		    // WriterHeader的显式调用主要用于发送错误码。
		    WriteHeader(int)
		    // Write向连接中写入作为HTTP的一部分回复的数据。
		    // 如果被调用时还未调用WriteHeader，本方法会先调用WriteHeader(http.StatusOK)
		    // 如果Header中没有"Content-Type"键，
		    // 本方法会使用包函数DetectContentType检查数据的前512字节，将返回值作为该键的值。
		    Write([]byte) (int, error)
			}  ResponseWriter接口被HTTP处理器用于构造HTTP回复。


			type Transport struct {
		    // Proxy指定一个对给定请求返回代理的函数。
		    // 如果该函数返回了非nil的错误值，请求的执行就会中断并返回该错误。
		    // 如果Proxy为nil或返回nil的*URL置，将不使用代理。
		    Proxy func(*Request) (*url.URL, error)
		    // Dial指定创建TCP连接的拨号函数。如果Dial为nil，会使用net.Dial。
		    Dial func(network, addr string) (net.Conn, error)
		    // TLSClientConfig指定用于tls.Client的TLS配置信息。
		    // 如果该字段为nil，会使用默认的配置信息。
		    TLSClientConfig *tls.Config
		    // TLSHandshakeTimeout指定等待TLS握手完成的最长时间。零值表示不设置超时。
		    TLSHandshakeTimeout time.Duration
		    // 如果DisableKeepAlives为真，会禁止不同HTTP请求之间TCP连接的重用。
		    DisableKeepAlives bool
		    // 如果DisableCompression为真，会禁止Transport在请求中没有Accept-Encoding头时，
		    // 主动添加"Accept-Encoding: gzip"头，以获取压缩数据。
		    // 如果Transport自己请求gzip并得到了压缩后的回复，它会主动解压缩回复的主体。
		    // 但如果用户显式的请求gzip压缩数据，Transport是不会主动解压缩的。
		    DisableCompression bool
		    // 如果MaxIdleConnsPerHost!=0，会控制每个主机下的最大闲置连接。
		    // 如果MaxIdleConnsPerHost==0，会使用DefaultMaxIdleConnsPerHost。
		    MaxIdleConnsPerHost int
		    // ResponseHeaderTimeout指定在发送完请求（包括其可能的主体）之后，
		    // 等待接收服务端的回复的头域的最大时间。零值表示不设置超时。
		    // 该时间不包括获取回复主体的时间。
		    ResponseHeaderTimeout time.Duration
		    // 内含隐藏或非导出字段
			}

			type Client struct {
		    // Transport指定执行独立、单次HTTP请求的机制。
		    // 如果Transport为nil，则使用DefaultTransport。
		    Transport RoundTripper
		    // CheckRedirect指定处理重定向的策略。
		    // 如果CheckRedirect不为nil，客户端会在执行重定向之前调用本函数字段。
		    // 参数req和via是将要执行的请求和已经执行的请求（切片，越新的请求越靠后）。
		    // 如果CheckRedirect返回一个错误，本类型的Get方法不会发送请求req，
		    // 而是返回之前得到的最后一个回复和该错误。（包装进url.Error类型里）
		    //
		    // 如果CheckRedirect为nil，会采用默认策略：连续10此请求后停止。
		    CheckRedirect func(req *Request, via []*Request) error
		    // Jar指定cookie管理器。
		    // 如果Jar为nil，请求中不会发送cookie，回复中的cookie会被忽略。
		    Jar CookieJar
		    // Timeout指定本类型的值执行请求的时间限制。
		    // 该超时限制包括连接时间、重定向和读取回复主体的时间。
		    // 计时器会在Head、Get、Post或Do方法返回后继续运作并在超时后中断回复主体的读取。
		    //
		    // Timeout为零值表示不设置超时。
		    //
		    // Client实例的Transport字段必须支持CancelRequest方法，
		    // 否则Client会在试图用Head、Get、Post或Do方法执行请求时返回错误。
		    // 本类型的Transport字段默认值（DefaultTransport）支持CancelRequest方法。
		    Timeout time.Duration
			} Client类型代表HTTP客户端。它的零值（DefaultClient）是一个可用的使用DefaultTransport的客户端。
			Client的Transport字段一般会含有内部状态（缓存TCP连接），因此Client类型值应尽量被重用而不是每次需要都创建新的。Client类型值可以安全的被多个go程同时使用。
			Client类型的层次比RoundTripper接口（如Transport）高，还会管理HTTP的cookie和重定向等细节。
			func (c *Client) Do(req *Request) (resp *Response, err error) Do方法发送请求，返回HTTP回复。它会遵守客户端c设置的策略（如重定向、cookie、认证）。
			如果客户端的策略（如重定向）返回错误或存在HTTP协议错误时，本方法将返回该错误；如果回应的状态码不是2xx，本方法并不会返回错误。

			如果返回值err为nil，resp.Body总是非nil的，调用者应该在读取完resp.Body后关闭它。如果返回值resp的主体未关闭，c下层的RoundTripper接口（一般为Transport类型）可能无法重用resp主体下层保持的TCP连接去执行之后的请求。
			请求的主体，如果非nil，会在执行后被c.Transport关闭，即使出现错误。
			一般应使用Get、Post或PostForm方法代替Do方法。
			func (c *Client) Head(url string) (resp *Response, err error) Head向指定的URL发出一个HEAD请求
			func (c *Client) Get(url string) (resp *Response, err error) Get向指定的URL发出一个GET请求，如果回应的状态码如下，Get会在调用c.CheckRedirect后执行重定向
			func (c *Client) Post(url string, bodyType string, body io.Reader) (resp *Response, err error)  Post向指定的URL发出一个POST请求。bodyType为POST数据的类型， body为POST数据，作为请求的主体。如果参数body实现了io.Closer接口，它会在发送请求后被关闭。调用者有责任在读取完返回值resp的主体后关闭它。
			func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error) PostForm向指定的URL发出一个POST请求，url.Values类型的data会被编码为请求的主体。POST数据的类型一般会设为"application/x-www-form-urlencoded"。如果返回值err为nil，resp.Body总是非nil的，调用者应该在读取完resp.Body后关闭它。
			type Handler interface {
		    ServeHTTP(ResponseWriter, *Request)
			} 实现了Handler接口的对象可以注册到HTTP服务端，为特定的路径及其子树提供服务。 ServeHTTP应该将回复的头域和数据写入ResponseWriter接口然后返回。返回标志着该请求已经结束，HTTP服务端可以转移向该连接上的下一个请求。

			type HandlerFunc func(ResponseWriter, *Request) HandlerFunc type是一个适配器，通过类型转换让我们可以将普通的函数作为HTTP处理器使用。如果f是一个具有适当签名的函数，HandlerFunc(f)通过调用f实现了Handler接口。
			func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) ServeHTTP方法会调用f(w, r)
			type ServeMux struct {
		    // 内含隐藏或非导出字段
			} ServeMux类型是HTTP请求的多路转接器。它会将每一个接收的请求的URL与一个注册模式的列表进行匹配，并调用和URL最匹配的模式的处理器。
			模式是固定的、由根开始的路径，如"/favicon.ico"，或由根开始的子树，如"/images/"（注意结尾的斜杠）。较长的模式优先于较短的模式，因此如果模式"/images/"和"/images/thumbnails/"都注册了处理器，后一个处理器会用于路径以"/images/thumbnails/"开始的请求，前一个处理器会接收到其余的路径在"/images/"子树下的请求。
			注意，因为以斜杠结尾的模式代表一个由根开始的子树，模式"/"会匹配所有的未被其他注册的模式匹配的路径，而不仅仅是路径"/"。
			模式也能（可选地）以主机名开始，表示只匹配该主机上的路径。指定主机的模式优先于一般的模式，因此一个注册了两个模式"/codesearch"和"codesearch.google.com/"的处理器不会接管目标为"http://www.google.com/"的请求。
			ServeMux还会注意到请求的URL路径的无害化，将任何路径中包含"."或".."元素的请求重定向到等价的没有这两种元素的URL。
			func NewServeMux() *ServeMux     NewServeMux创建并返回一个新的*ServeMux
			func (mux *ServeMux) Handle(pattern string, handler Handler) Handle注册HTTP处理器handler和对应的模式pattern。如果该模式已经注册有一个处理器，Handle会panic。
			mux := http.NewServeMux()
			mux.Handle("/api/", apiHandler{})
			mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		    // The "/" pattern matches everything, so we need to check
		    // that we're at the root here.
		    if req.URL.Path != "/" {
		        http.NotFound(w, req)
		        return
		    }
		    fmt.Fprintf(w, "Welcome to the home page!")
			})

			func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) HandleFunc注册一个处理器函数handler和对应的模式pattern。
			func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string) Handler根据r.Method、r.Host和r.URL.Path等数据，返回将用于处理该请求的HTTP处理器。它总是返回一个非nil的处理器。如果路径不是它的规范格式，将返回内建的用于重定向到等价的规范路径的处理器。
			func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)  ServeHTTP将请求派遣到与请求的URL最匹配的模式对应的处理器。
			type Server struct {
		    Addr           string        // 监听的TCP地址，如果为空字符串会使用":http"
		    Handler        Handler       // 调用的处理器，如为nil会调用http.DefaultServeMux
		    ReadTimeout    time.Duration // 请求的读取操作在超时前的最大持续时间
		    WriteTimeout   time.Duration // 回复的写入操作在超时前的最大持续时间
		    MaxHeaderBytes int           // 请求的头域最大长度，如为0则用DefaultMaxHeaderBytes
		    TLSConfig      *tls.Config   // 可选的TLS配置，用于ListenAndServeTLS方法
		    // TLSNextProto（可选地）指定一个函数来在一个NPN型协议升级出现时接管TLS连接的所有权。
		    // 映射的键为商谈的协议名；映射的值为函数，该函数的Handler参数应处理HTTP请求，
		    // 并且初始化Handler.ServeHTTP的*Request参数的TLS和RemoteAddr字段（如果未设置）。
		    // 连接在函数返回时会自动关闭。
		    TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
		    // ConnState字段指定一个可选的回调函数，该函数会在一个与客户端的连接改变状态时被调用。
		    // 参见ConnState类型和相关常数获取细节。
		    ConnState func(net.Conn, ConnState)
		    // ErrorLog指定一个可选的日志记录器，用于记录接收连接时的错误和处理器不正常的行为。
		    // 如果本字段为nil，日志会通过log包的标准日志记录器写入os.Stderr。
		    ErrorLog *log.Logger
		    // 内含隐藏或非导出字段
			} Server类型定义了运行HTTP服务端的参数。Server的零值是合法的配置。
			func (s *Server) SetKeepAlivesEnabled(v bool) SetKeepAlivesEnabled控制是否允许HTTP闲置连接重用（keep-alive）功能。默认该功能总是被启用的。只有资源非常紧张的环境或者服务端在关闭进程中时，才应该关闭该功能。
			func (srv *Server) Serve(l net.Listener) error Serve会接手监听器l收到的每一个连接，并为每一个连接创建一个新的服务go程。该go程会读取请求，然后调用srv.Handler回复请求。
			func (srv *Server) ListenAndServe() error ListenAndServe监听srv.Addr指定的TCP地址，并且会调用Serve方法接收到的连接。如果srv.Addr为空字符串，会使用":http"。

			func FileServer(root FileSystem) Handler
			FileServer返回一个使用FileSystem接口root提供文件访问服务的HTTP处理器。要使用操作系统的FileSystem接口实现，可使用http.Dir：
			http.Handle("/", http.FileServer(http.Dir("/tmp")))
			func Head(url string) (resp *Response, err error) Head向指定的URL发出一个HEAD请求，如果回应的状态码如下，Head会在调用c.CheckRedirect后执行重定向

			func Get(url string) (resp *Response, err error) Get向指定的URL发出一个GET请求，如果回应的状态码如下，Get会在调用c.CheckRedirect后执行重定向。
			func Post(url string, bodyType string, body io.Reader) (resp *Response, err error) Post向指定的URL发出一个POST请求。bodyType为POST数据的类型， body为POST数据，作为请求的主体。如果参数body实现了io.Closer接口，它会在发送请求后被关闭。调用者有责任在读取完返回值resp的主体后关闭它。
			func PostForm(url string, data url.Values) (resp *Response, err error)  PostForm向指定的URL发出一个POST请求，url.Values类型的data会被编码为请求的主体。如果返回值err为nil，resp.Body总是非nil的，调用者应该在读取完resp.Body后关闭它。

			func Handle(pattern string, handler Handler) Handle注册HTTP处理器handler和对应的模式pattern（注册到DefaultServeMux）。如果该模式已经注册有一个处理器，Handle会panic。ServeMux的文档解释了模式的匹配机制。

			func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) HandleFunc注册一个处理器函数handler和对应的模式pattern（注册到DefaultServeMux）。ServeMux的文档解释了模式的匹配机制。
			func Serve(l net.Listener, handler Handler) error  Serve会接手监听器l收到的每一个连接，并为每一个连接创建一个新的服务go程。该go程会读取请求，然后调用handler回复请求。handler参数一般会设为nil，此时会使用DefaultServeMux。

			func ListenAndServe(addr string, handler Handler) error ListenAndServe监听TCP地址addr，并且会使用handler参数调用Serve函数处理接收到的连接。handler参数一般会设为nil，此时会使用DefaultServeMux。

			func ListenAndServeTLS(addr string, certFile string, keyFile string, handler Handler) error  ListenAndServeTLS函数和ListenAndServe函数的行为基本一致，除了它期望HTTPS连接之外。此外，必须提供证书文件和对应的私钥文件。如果证书是由权威机构签发的，certFile参数必须是顺序串联的服务端证书和CA证书。如果srv.Addr为空字符串，会使用":https"。
			一个简单的服务端例子：
			func handler(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("This is an example server.\n"))
			}
			func main() {
			http.HandleFunc("/", handler)
			log.Printf("About to listen on 10443. Go to https://127.0.0.1:10443/")
			err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
			if err != nil {
				log.Fatal(err)
			}
			}
			程序员可以使用crypto/tls包的generate_cert.go文件来生成cert.pem和key.pem两个文件。
	*/

	for {
		time.Sleep(2 * time.Second)
		curl() //死循环 ，让程序挂起 来校验
	}

}

/***
结果集 ：启动一个web服务可以 使用go run net.go &

当你发起访问的时候
root@e7939faf8694:/go/src/LessonGo# curl http://127.0.0.1:8000
URL.Path = "/"
root@e7939faf8694:/go/src/LessonGo#

root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# go run net.go  http://localhost:8000/hello
The context of fetching are URL.Path = "/hello"
The context of fetching are URL.Path = "/hello"
The context of fetching are URL.Path = "/hello"
The context of fetching are URL.Path = "/hello"
The context of fetching are URL.Path = "/hello"
time.Sleep(2 * time.Second)The context of fetching are URL.Path = "/hello"
The context of fetching are URL.Path = "/hello"
The context of fetching are URL.Path = "/hello"
^Z
[1]+  Stopped                 go run net.go http://localhost:8000/hello
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage#


***/
