package main

//此框架经常会结合html模板使用，具体请研究html模板引擎的语法。

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"time"
)

func basic() {
	//gin.SetMode(gin.ReleaseMode) 加载后默认按生产模式跑，就没有debug模式下的日志了。
	r := gin.Default() //func Default() *Engine 返回了一个Engine instance with the Logger and Recovery middleware already attached。自带默认的。

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data) //使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON。
	})

	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务，如果不填的话，按默认8080。
	//Run attaches the router to a http.Server and starts listening and serving HTTP requests.
	//It is a shortcut for http.ListenAndServe(addr, router) 等价
	//Note: this method will block the calling goroutine indefinitely unless an error happens.

}

func loadtmpl() {

	//载入HTML模板文件
	//可以使用自定义分隔 。
	//r.Delims("{[{", "}]}")
	//r.LoadHTMLGlob("/path/to/templates") 默认是*这种

	r := gin.Default() //返回*Engine
	r.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.Run(":8080")
}

func loaddiffdirtmpl() {

	//使用不同目录下名称相同的模板
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	r.Run(":8080")
}

func htmltemplate() {
	//使用自定义的 html 模板渲染
	r := gin.Default()
	r.Delims("{[{", "}]}")
	html := template.Must(template.ParseFiles("templates/layout.html")) //可以多个文件 也可以一个文件
	r.SetHTMLTemplate(html)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"status": "success",
		})
	})
	r.Run(":8080")
}

//使用自定义模板功能。对模板内容进行函数计算。
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

func helloadd(a string) string {
	return a + " world"
}

func htmltmplfunc() {
	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.SetFuncMap(template.FuncMap{ //允许多个函数可以被调用，可以添加多个到router里面去。等待被调用。
		"formatAsDate": formatAsDate,
		"helloadd":     helloadd, //相当于传入多个函数给html模板调用。
	})
	router.LoadHTMLFiles("templates/layout.tmpl")

	router.GET("/layout", func(c *gin.Context) {
		c.HTML(http.StatusOK, "layout.tmpl", map[string]interface{}{ //后者相当于传参进入到了tmpl里面，然后被里面的内容调用
			"now":  time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
			"now2": "hello", //可以传多个参数给html模板调用。
		})
	})

	router.Run(":8080")

}

//HTTP2 server 推送 ，目前只能是https这种网站，需要证书。对于 访问一个主页 等待资源响应的时间减慢了主页面加载速度。并发传输并不能提高串行解析的资源访问体验。
//一系列资源需要被服务器主动并发推送依赖资源至客户端。这样减少了客户端去下载解析资源的串行。推送静态文件。
//推送HTTP2 只用于https:// 目前 http:// 将继续使用HTTP1
//HTTP/2 新增的另一个强大的新功能是，服务器可以对一个客户端请求发送多个响应。
//换句话说，除了对最初请求的响应外，服务器还可以向客户端推送额外资源，而无需客户端明确地请求。
func http2push() {
	var html = template.Must(template.New("http").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
  <h1>{{ .status }}</h1>
</body>
</html>
`))

	r := gin.Default()
	// router.Static("/static", "/var/www") 静态文件管理。左边是域名路径，右面是真实的服务器路径。将此域名映射到当前static路径目录。
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html) //载入进的html含有静态文件 可以被主动推送。

	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// 使用 pusher.Push() 做服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil { //可以一口气推送多个静态文件。
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(200, "http", gin.H{
			"status": "success",
		})
	})

	// 监听并在 https://127.0.0.1:8080 上启动服务 如果需要https访问，要声明公私钥
	//r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
	r.Run(":8080") //在开放互联网上HTTP 2.0将只用于https://网址，而 http://网址将继续使用HTTP/1
}

//使用 JSONP 向不同域的服务器请求数据。如果查询参数存在回调，则将回调添加到响应体中。
//Jsonp(JSON with Padding) 是 json 的一种"使用模式"，可以让网页从别的域名（网站）那获取资料，即跨域读取数据。
func Jsonp() {

	var html = template.Must(template.New("http").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
  <h1>{{ .status }}</h1>
</body>
</html>
`))

	r := gin.Default()
	r.SetHTMLTemplate(html)
	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// 如果callback 是 x ，则使用curl http://127.0.0.1:8080/JSONP?callback=x ，返回结果生成回调函数要求格式。
		// 将输出：x({\"foo\":\"bar\"})
		//JSONP serializes the given struct as JSON into the response body.
		//It add padding to response body to request data from a server residing in a different domain than the client.
		// It also sets the Content-Type as "application/javascript".
		c.JSONP(http.StatusOK, data) //返回回调函数
		//c.HTML(200, "http", gin.H{"status": "success"})

	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")

}

//Multipart/Urlencoded 绑定 ,即按form格式来传入数据。 需要客户端按form格式来传入数据，得到返回结果。
type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func formpost() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		// 你可以使用显式绑定声明绑定 multipart form：
		// c.ShouldBindWith(&form, binding.Form)
		// 或者简单地使用 ShouldBind 方法自动绑定：
		var form LoginForm
		// 在这种情况下，将自动选择合适的绑定
		if c.ShouldBind(&form) == nil {
			if form.User == "james" && form.Password == "cole" {
				c.JSON(200, gin.H{"status": "you are logged in"}) //返回纯json
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})
	r.Run(":8080")
}

/*
结果集：
root@e7939faf8694:/go/src/LessonGo# curl -v  --form user=james --form password=cole http://localhost:8080/login
{"status":"you are logged in"}  纯jason返回
root@e7939faf8694:/go/src/LessonGo#
*/

//Multipart/Urlencoded 表单 。可以get和post一起使用。
//get是从服务器上获取数据，post是向服务器传送数据。get会带一些默认的参数，post是将加工的数据传给服务器。
//Get + Post 混合 对于一个前端页面 要填一些内容 ，会使用form，但是还有一些自带的url参数需要则使用query。

func post_table() {
	r := gin.Default()

	r.POST("/form_post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		message := c.PostForm("message")               //填写前端表单可以使用 ,然后上面的query 用来放非表单的一些默认参数。
		nick := c.DefaultPostForm("nick", "anonymous") //此方法可以设置默认值

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
			"id":      id,
			"page":    page,
		})
		fmt.Printf("id: %s; page: %s; nick: %s; message: %s", id, page, nick, message) //让控台除了gin的打印日志也增加你自己要求的格式输出日志。
	})
	r.Run(":8080")
}

/*
结果集:
root@e7939faf8694:/go/src/LessonGo# curl  --form message=abc --form nick=bbb  http://localhost:8080/form_post?id=1\&page=3
{"id":"1","message":"abc","nick":"bbb","page":"3","status":"posted"}
root@e7939faf8694:/go/src/LessonGo#  可以看出来 两个参数是通过表单传给服务器，另外两个是默认自带的参数，例如固定标记等。
*/

//PureJSON 通常，JSON 使用 unicode 替换特殊 HTML 字符，例如 < 变为 \ u003c。如果要按字面对这些字符进行编码，则可以使用 PureJSON。

func pureJson() {

	r := gin.Default()
	// 提供 unicode 实体
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	// 提供字面字符
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

/*
结果集:
root@e7939faf8694:/go/src/LessonGo# curl http://127.0.0.1:8080/json
{"html":"\u003cb\u003eHello, world!\u003c/b\u003e"}
root@e7939faf8694:/go/src/LessonGo# curl http://127.0.0.1:8080/purejson
{"html":"<b>Hello, world!</b>"}
root@e7939faf8694:/go/src/LessonGo#
*/

//SecureJSON
//使用 SecureJSON 防止 json 劫持。如果给定的结构是数组值，则默认预置 "while(1)," 到响应体。
func securejson() {
	r := gin.Default()
	// 你也可以使用自己的 SecureJSON 前缀
	// r.SecureJsonPrefix(")]}',\n")
	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		// 将输出：while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})
	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

/*
结果集：
root@e7939faf8694:/go/src/LessonGo# curl http://127.0.0.1:8080/someJSON
while(1);["lena","austin","foo"]
*/

func main() {
	//go basic()
	//go loadtmpl()
	//go loaddiffdirtmpl()
	//go htmltemplate()
	//go htmltmplfunc()
	//go http2push()
	//go Jsonp()
	//go formpost()
	//go post_table()
	//go pureJson()
	securejson()
}
