package main

//此框架经常会结合html模板使用，具体请研究html模板引擎的语法。

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/sync/errgroup"
	"gopkg.in/go-playground/validator.v8"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"reflect"
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
	/*
		HTTPS配置步骤:
		首先在阿里云搞定ICP域名备案
		添加一个子域名
		给子域名申请免费 SSL 证书, 然后下载证书对应的 pem 和 key 文件.
	*/
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

//XML/JSON/YAML/ProtoBuf 渲染
func xmljsonyamlprotobuf() {
	r := gin.Default()

	// gin.H 是 map[string]interface{} 的一种快捷方式
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// 你也可以使用一个结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// 注意 msg.Name 在 JSON 中变成了 "user"
		// 将输出：{"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	/*ProtoBuf 这种需要专门研究
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf 的具体定义写在 testdata/protoexample 文件中。
		data := &Test{
			Label: &label,
			Reps:  reps,
		}
		// 请注意，数据在响应中变为二进制数据
		// 将输出被 protoexample.Test protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data)
	})
	*/
	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")

}

/*
结果集：
root@e7939faf8694:/go/src/LessonGo# curl http://127.0.0.1:8080/someJSON
{"message":"hey","status":200}
root@e7939faf8694:/go/src/LessonGo# curl http://127.0.0.1:8080/moreJSON
{"user":"Lena","Message":"hey","Number":123}
root@e7939faf8694:/go/src/LessonGo# curl http://127.0.0.1:8080/someXML
<map><message>hey</message><status>200</status></map>root@e7939faf8694:/go/src/LessonGo# curl http://127.0.0.1:8080/someYAML
message: hey
status: 200
root@e7939faf8694:/go/src/LessonGo#
*/

//上传单个文件
func singlefileupload() {
	r := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// r.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// 上传文件至指定目录
		c.SaveUploadedFile(file, "/go/src/LessonGo/class_gin/src/")

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	r.Run(":8080")

}

/*
结果集：
root@e7939faf8694:/go/src/LessonGo/class_gin/src#
curl -X POST http://localhost:8080/upload  -F "file=@/go/src/LessonGo/LICENSE"   -H "Content-Type: multipart/form-data"
'LICENSE' uploaded!
root@e7939faf8694:/go/src/LessonGo/class_gin/src#
*/

//上传多个文件
func multiplefilesupload() {
	r := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// 上传文件至指定目录
			c.SaveUploadedFile(file, "/go/src/LessonGo/class_gin/src/") //上传文件到指定的目录
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	r.Run(":8080")

}

/*
结果集：
root@e7939faf8694:/go#
curl -X POST http://localhost:8080/upload  -F "upload[]=@/go/1.txt"  -F "upload[]=@/go/2.txt"   -H "Content-Type: multipart/form-data"
2 files uploaded

[GIN] 2019/06/09 - 14:30:20 | 200 |      1.2997ms |       127.0.0.1 | POST     /upload
2019/06/09 14:30:22 1.txt
2019/06/09 14:30:22 2.txt
*/

//不使用默认的中间件
//使用r := gin.New() 代替 r := gin.Default()  Default 使用 Logger 和 Recovery 中间件

//从 reader 读取数据，从一个文件，从一个网络地址等，都是reader中读取内容，然后进行处理。
func servingdatafromreader() {
	r := gin.Default()
	r.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("http://fund.eastmoney.com/f10/F10DataApi.aspx?type=lsjz&code=000001&per=3")
		//response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders) //会把整个抓取的内容全部打印出来。
	})
	r.Run(":8080")

}

/*
优雅地重启或停止
有很多替代的三方包来帮助进行优雅的停止。如果使用Go1.8以上，则不需要这些库。考虑使用 http.Server 内置的 Shutdown() 方法优雅地关机。
*/

func shutdowngraceful() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server\n") //直接返回一个字符串，而不是结构体。
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt) //按ctrl-c来终止 ，系统会捕捉到中断信号。Notify函数让signal包将输入信号转发到c。如果没有列出要传递的信号，会将所有输入信号传递到c；否则只传递列出的输入信号。
	<-quit                            //然后当有信号进来，然后被消耗后，阻塞解除，进入下一步。
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //超时自动取消，是多少时间后自动取消Context上下文。
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil { //执行shutdown函数
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

//使用 BasicAuth 中间件  ，使用中间件来进行授权。限制用户访问。
//Basic Auth是一种开放平台认证方式，简单的说就是需要你输入用户名和密码才能继续访问。Bath Auth是其中一种认证方式，另一种是OAuth。
func usingbasicauthmiddleware() {
	// 模拟一些私人数据
	var secrets = gin.H{
		"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
		"austin": gin.H{"email": "austin@example.com", "phone": "666"},
		"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
	}
	r := gin.Default()

	// 路由组使用 gin.BasicAuth() 中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	//Group函数注册了一个群组路由，gin.BasicAuth就是中间件，它的参数gin.Accounts其实是一个map[string]string类型的映射，这里是用来记录用户名和密码。
	//然后在路由/admin之下，又注册了/secrets路由，所以他的完整路由应该是/admin/secrets。
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{ //Accounts defines a key/value for user/pass list of authorized logins.
		"foo":    "bar", //用户名：密码   type Accounts map[string]string
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets 端点
	// 触发 "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string) //AuthUserKey is the cookie name for user credential in basic auth. const AuthUserKey = "user"
		if secret, ok := secrets[user]; ok {        //只有符合条件的用户 才有正确返回。首先users用户才会返回200，其次只有存在在MustGet里面的才能返回OK。
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")

}

/*
结果集:
root@e7939faf8694:/go/src/LessonGo# curl  lena:hello2@localhost:8080/admin/secrets  gin返回HTTP 200
{"secret":{"email":"lena@guapa.com","phone":"523443"},"user":"lena"}
root@e7939faf8694:/go/src/LessonGo# curl  manu:4321@localhost:8080/admin/secrets  gin返回HTTP 200 ，只是此user不在secrects里面。
{"secret":"NO SECRET :(","user":"manu"}
root@e7939faf8694:/go/src/LessonGo# curl  man4u:43221@localhost:8080/admin/secrets 没有这个用户。gin直接返回 HTTP 401 未授权: (Unauthorized)
root@e7939faf8694:/go/src/LessonGo#
*/

//使用 HTTP 方法
func httpmethods() {
	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	// 使用默认中间件（logger 和 recovery 中间件）创建 gin 路由
	router := gin.Default()

	//router.GET("/someGet", getting)
	//router.POST("/somePost", posting)
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	// 默认在 8080 端口启动服务，除非定义了一个 PORT 的环境变量。
	router.Run()
	// router.Run(":3000") hardcode 端口号

}

//使用中间件
/*
func usingmiddleware() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())

	// 你可以为每个路由添加任意数量的中间件。
	//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)


	// 认证路由组
	//authorized := r.Group("/", AuthRequired())
	// 和使用以下两行代码的效果完全一样:
	authorized := r.Group("/")
	// 路由组中间件! 在此例中，我们在 "authorized" 路由组中使用自定义创建的
	// AuthRequired() 中间件
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		// 嵌套路由组
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")

}
*/
//只绑定 url 查询字符串 ShouldBindQuery 对应ShouldBind
//ShouldBindQuery 函数只绑定 url 查询参数而忽略 post 数据。
type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}
func onlybindquerystring() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":8080")

}

/*
结果集：只绑定 url 查询字符串 ，对form进行忽略。
root@e7939faf8694:/go/src/LessonGo# curl  --form name=man --form address=hangzhou   localhost:8080/testing?name=man\&address=hangzhou
Success
root@e7939faf8694:/go/src/LessonGo#
[GIN-debug] Listening and serving HTTP on :8080
2019/06/10 08:07:57 ====== Only Bind By Query String ======
2019/06/10 08:07:57 man
2019/06/10 08:07:57 hangzhou
[GIN] 2019/06/10 - 08:07:57 | 200 |       489.6µs |       127.0.0.1 | POST     /testing?name=man&address=hangzhou
*/

//在中间件中使用 Goroutine
//当在中间件或 handler 中启动新的 Goroutine 时，不能使用原始的上下文，必须使用只读副本。
func goroutineinmiddleware() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) { //Context异步，curl已经返回，gin控制台等待协程执行完毕输出结果。
		// 创建在 goroutine 中使用的副本
		cCp := c.Copy()
		go func() {
			// 用 time.Sleep() 模拟一个长任务。
			time.Sleep(5 * time.Second)

			// 请注意您使用的是复制的上下文 "cCp"，这一点很重要
			log.Println("Done! in path " + cCp.Request.URL.Path) //log是输出到gin控制台的。
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) { //Context同步，curl等待整个执行完毕返回结果。
		// 用 time.Sleep() 模拟一个长任务。
		time.Sleep(5 * time.Second)

		// 因为没有使用 goroutine，不需要拷贝上下文
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")

}

//Gin 默认允许只使用一个 html 模板。 使用https://github.com/gin-contrib/multitemplate 来完成多模板渲染。

//如何记录日志
func writelog() {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	//gin.DisableConsoleColor()

	// 记录到文件。
	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)  //创建写入动作可以到多个文件

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) //只是将gin的控台标准输出到文件里面，个人定义的log不会捕捉到文件。

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
		log.Println("success") //自定义的打印 只会出现在控台上，不属于Gin的标准打印日志，不会输出到指定文件。
	})

	router.Run(":8080")

}

//定义路由日志的格式 ，更改Gin的默认打印格式，按照你想要的打印格式进行输出。
//如果你想要以指定的格式（例如 JSON，key values 或其他格式）记录信息，则可以使用 gin.DebugPrintRouteFunc 指定格式。
//在下面的示例中，我们使用标准日志包记录所有路由，但你可以使用其他满足你需求的日志工具。

func defineformatlogs() {
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.POST("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "foo")
	})

	r.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run()

}

//将 request body 绑定到不同的结构体中。一般通过调用 c.Request.Body 方法绑定数据，但不能多次调用这个方法。
type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func SomeHandler1(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// c.ShouldBind 使用了 c.Request.Body，不可重用。
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// 因为现在 c.Request.Body 是 EOF，所以这里会报错。
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		c.String(http.StatusOK, `the body is not form either`)
	}
}

//要想多次绑定，可以使用 c.ShouldBindBodyWith。
/*c.ShouldBindBodyWith 会在绑定之前将 body 存储到上下文中。
这会对性能造成轻微影响，如果调用一次就能完成绑定的话，那就不要用这个方法。
只有某些格式需要此功能，如 JSON, XML, MsgPack, ProtoBuf。 对于其他格式, 如 Query, Form, FormPost, FormMultipart 可以多次调用 c.ShouldBind()
而不会造成任任何性能损失

func SomeHandler2(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// 读取 c.Request.Body 并将结果存入上下文。
	if errA := c.ShouldBindBodyWith(&objA, BindingBody.binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// 这时, 复用存储在上下文中的 body。
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB JSON`)
		// 可以接受其他格式
	} else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		c.String(http.StatusOK, `the body should be formB XML`)
	} else {
		c.String(http.StatusOK, `the body is not form either`)
	}
}
*/

/*
支持 Let's Encrypt
Let's Encrypt是国外一个公共的免费SSL项目 由 Linux 基金会托管 目的就是向网站自动签发和管理免费证书，以便加速互联网由HTTP过渡到HTTPS
其自动化发行证书，但是证书只有90天的有效期。适合个人使用或者临时使用，不用再忍受自签发证书不受浏览器信赖的提示。
一行代码支持 LetsEncrypt HTTPS servers 示例。
package main

import (
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
}

自定义 autocert manager 示例。
autocert 负责生成SSL证书。
autocert是使用golang完整实现的acme客户端, 相对于letsencrypt-auto,
零依赖:   二进制版的autocert不依赖其它安装包, 不需要任何配置文件, 只需传入一个ServerName参数即可自动处理签发ssl证书的所有步骤, 随用随走。
package main

import (
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(r, &m))
}


*/

//映射查询字符串或表单参数，即映射查询又映射表格。
func mapasquerystringorpostform() {
	router := gin.Default()
	router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")        //QueryMap returns a map for a given query key。func (c *Context) QueryMap(key string) map[string]string
		names := c.PostFormMap("names") //PostFormMap returns a map for a given form key。func (c *Context) PostFormMap(key string) map[string]string

		fmt.Printf("ids: %v; names: %v", ids, names)
	})
	router.Run(":8080")
}

/*
结果集：
root@e7939faf8694:/go/src/LessonGo# curl  -g --form names[first]=thinkerou --form names[second]=tianou   localhost:8080/post?ids[a]=1234&ids[b]=hello
[1] 14254
root@e7939faf8694:/go/src/LessonGo#
[1]+  Done                    curl -g --form names[first]=thinkerou --form names[second]=tianou localhost:8080/post?ids[a]=1234
root@e7939faf8694:/go/src/LessonGo#
*/

//查询字符串参数
func querystringparam() {
	router := gin.Default()

	// 使用现有的基础请求对象解析查询字符串参数。
	// 示例 URL： /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest") //默认无输入则guest
		lastname := c.Query("lastname")                   // c.Request.URL.Query().Get("lastname") 的一种快捷方式。无则为空返回。

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.Run(":8080")

}

//模型绑定和验证 要将请求体绑定到结构体中，使用模型绑定。 Gin目前支持JSON、XML、YAML和标准表单值的绑定（foo=bar＆boo=baz）。
//Gin使用 go-playground/validator.v8 进行验证。
//使用时，需要在要绑定的所有字段上，设置相应的tag。 例如，使用 JSON 绑定时，设置字段标签为 json:"fieldname"。
/*
Gin提供了两类绑定方法：
Type - Must bind
Methods - Bind, BindJSON, BindXML, BindQuery, BindYAML
Behavior - 这些方法属于 MustBindWith 的具体调用。 如果发生绑定错误，则请求终止，并触发 c.AbortWithError(400, err).SetType(ErrorTypeBind)。响应状态码被设置为 400 并且 Content-Type 被设置为 text/plain; charset=utf-8。 如果您在此之后尝试设置响应状态码，Gin会输出日志 [GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422。 如果您希望更好地控制绑定，考虑使用 ShouldBind 等效方法。
Type - Should bind
Methods - ShouldBind, ShouldBindJSON, ShouldBindXML, ShouldBindQuery, ShouldBindYAML
Behavior - 这些方法属于 ShouldBindWith 的具体调用。 如果发生绑定错误，Gin 会返回错误并由开发者处理错误和请求。
使用 Bind 方法时，Gin 会尝试根据 Content-Type 推断如何绑定。 如果你明确知道要绑定什么，可以使用 MustBindWith 或 ShouldBindWith。
*/
// 绑定 JSON
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"` //binding:"required"` 强制要有数据，否则会抛出error
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func bindingandvalidation() {
	router := gin.Default()

	// 绑定 JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// 绑定 XML (
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>user</user>
	//		<password>123</user>
	//	</root>)
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// 绑定 HTML 表单 (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// 根据 Content-Type Header 推断使用哪个绑定器。
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
}

/*
结果集:
root@e7939faf8694:/go/src/LessonGo#  curl  -X POST \
>   http://localhost:8080/loginJSON \
>   -H 'content-type: application/json' \
>   -d '{ "user": "manu","password":"123" }'
{"status":"you are logged in"}
root@e7939faf8694:/go/src/LessonGo#

*/

/*
绑定 HTML 复选框
...

type myForm struct {
    Colors []string `form:"colors[]"`  \\声明字符串数组，并进行绑定
}

...

func formHandler(c *gin.Context) {
    var fakeForm myForm
    c.ShouldBind(&fakeForm)  //绑定到context
    c.JSON(200, gin.H{"color": fakeForm.Colors})
}

...
form.html
<form action="/" method="POST">
    <p>Check some colors</p>
    <label for="red">Red</label>
    <input type="checkbox" name="colors[]" value="red" id="red">
    <label for="green">Green</label>
    <input type="checkbox" name="colors[]" value="green" id="green">
    <label for="blue">Blue</label>
    <input type="checkbox" name="colors[]" value="blue" id="blue">
    <input type="submit">
</form>

结果集:
{"color":["red","green","blue"]}
*/

//绑定 Uri
//URI（Uniform Resource Identifier，统一资源标识符）就是在IMS网络中IMS用户的“名字”，也就是IMS用户的身份标识。
type Person2 struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func bingduri() {

	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var person Person2
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
	route.Run(":8080")

}

/*
结果集：
$ curl -v localhost:8088/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
$ curl -v localhost:8088/thinkerou/not-uuid
会校验uri的格式和长度 才能正确返回200

*/

//绑定查询字符串或表单数据
type Person3 struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func bindqueryorpost() {
	route := gin.Default()
	route.GET("/testing", startPage2)
	route.Run(":8080")
}

func startPage2(c *gin.Context) {
	var person Person3
	// 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
	// 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。
	// 查看更多：https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&person) == nil { //返回值为error为空
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}
	c.String(200, "Success")
}

/*
结果集:
root@e7939faf8694:/go/src/LessonGo# curl -X GET "localhost:8080/testing?name=appleboy&address=xyz&birthday=1992-03-15"
Successroot@e7939faf8694:/go/src/LessonGo#
[GIN-debug] GET    /testing                  --> main.startPage2 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
2019/06/10 11:57:19 appleboy
2019/06/10 11:57:19 xyz
2019/06/10 11:57:19 1992-03-15 00:00:00 +0000 UTC
[GIN] 2019/06/10 - 11:57:19 | 200 |       698.9µs |       127.0.0.1 | GET      /testing?name=appleboy&address=xyz&birthday=1992-03-15

*/

//绑定表单数据至自定义结构体
//总之, 目前仅支持没有 form 的嵌套结构体。
type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func GetDataC(c *gin.Context) {
	var b StructC
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}

func GetDataD(c *gin.Context) {
	var b StructD
	c.Bind(&b)
	c.JSON(200, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}
func bindformdatarequestwithcustomstruct() {
	r := gin.Default()
	r.GET("/getb", GetDataB)
	r.GET("/getc", GetDataC)
	r.GET("/getd", GetDataD)
	r.Run()
}

/*
结果集：
root@e7939faf8694:/go/src/LessonGo# curl "http://localhost:8080/getb?field_a=hello&field_b=world"
{"a":{"FieldA":"hello"},"b":"world"}
root@e7939faf8694:/go/src/LessonGo# curl "http://localhost:8080/getc?field_a=hello&field_c=world"
{"a":{"FieldA":"hello"},"c":"world"}
root@e7939faf8694:/go/src/LessonGo# curl "http://localhost:8080/getd?field_x=hello&field_d=world"
{"d":"world","x":{"FieldX":"hello"}}
root@e7939faf8694:/go/src/LessonGo#

*/

//自定义 HTTP 配置 ，直接使用 http.ListenAndServe()
func customhttpconfig1() {

	router := gin.Default()
	http.ListenAndServe(":8080", router)

}

func customhttpconfig2() {
	router := gin.Default()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}

//自定义中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置 example 变量
		c.Set("example", "12345")

		// 请求前

		c.Next()

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}

func custommiddleware() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// 打印："12345"
		log.Println(example)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")

}

/*
结果集：
gin的日志结果完全发生了变化：
root@e7939faf8694:/go/src/LessonGo# curl http://127.0.0.1:8080/test
root@e7939faf8694:/go/src/LessonGo#
[GIN-debug] Listening and serving HTTP on :8080
2019/06/10 12:31:38 3.9µs
2019/06/10 12:31:38 404
2019/06/10 12:31:43 12345
2019/06/10 12:31:43 58.6µs
2019/06/10 12:31:43 200
2019/06/10 12:31:46 12345
2019/06/10 12:31:46 56.2µs
2019/06/10 12:31:46 200
2019/06/10 12:31:48 12345
2019/06/10 12:31:48 55.2µs
2019/06/10 12:31:48 200
*/

//自定义验证器 注册自定义验证器
// Booking 包含绑定和验证的数据。
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func customvalidators() {

	route := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate) //binding.validator来负责校验
	}

	route.GET("/bookable", getBookable)
	route.Run(":8080")
}

/*
结果集：
root@e7939faf8694:/go/src/LessonGo# curl http://127.0.0.1:8080/bookable
{"error":"Key: 'Booking.CheckIn' Error:Field validation for 'CheckIn' failed on the 'required' tag\nKey: 'Booking.CheckOut' Error:Field validation for 'CheckOut' failed on the 'required' tag"}
root@e7939faf8694:/go/src/LessonGo# curl "localhost:8080/bookable?check_in=2019-06-10&check_out=2019-06-11"
{"message":"Booking dates are valid!"}
root@e7939faf8694:/go/src/LessonGo# curl "localhost:8080/bookable?check_in=2019-06-08&check_out=2019-06-09"
{"error":"Key: 'Booking.CheckIn' Error:Field validation for 'CheckIn' failed on the 'bookabledate' tag"}
root@e7939faf8694:/go/src/LessonGo#
*/

//设置和获取 Cookie
func cookie() {

	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie") //查找cookie

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true) //SetCookie adds a Set-Cookie header to the ResponseWriter's headers
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.Run()
}

//路由参数
func paraminpath() {
	router := gin.Default()

	// 此 handler 将匹配 /user/john 但不会匹配 /user/ 或者 /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 此 handler 将匹配 /user/john/ 和 /user/john/send
	// 如果没有其他路由匹配 /user/john，它将重定向到 /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")

}

//路由组 ,主要为了防止后续路由升级，进行兼容性工作。
/*
func groupingroutes() {
	router := gin.Default()

	// 简单的路由组: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// 简单的路由组: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")

}
*/

//运行多个服务
var (
	g errgroup.Group
)

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})

	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})

	return e
}

//运行多个服务
func runmultipleservice() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

/*
结果集：
root@e7939faf8694:/go/src/LessonGo# curl http://127.0.0.1:8080
{"code":200,"error":"Welcome server 01"}
root@e7939faf8694:/go/src/LessonGo# curl http://127.0.0.1:8081
{"code":200,"error":"Welcome server 02"}
root@e7939faf8694:/go/src/LessonGo#

*/

//重定向 HTTP 重定向很容易。 内部、外部重定向均支持。
/*
r.GET("/test", func(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
})
*/

//路由重定向，使用 HandleContext： HandleContext re-enter a context that has been rewritten。
/*
r.GET("/test", func(c *gin.Context) {
    c.Request.URL.Path = "/test2"
    r.HandleContext(c)
})
r.GET("/test2", func(c *gin.Context) {
    c.JSON(200, gin.H{"hello": "world"})
})
*/

//提供静态文件服务
func servingstaticfiles() {

	r := gin.Default()
	r.Static("/assets", "/go/src/LessonGo/class_gin/src/templates/static") //逻辑名称 对应实际名称
	r.StaticFS("/more_static", http.Dir("/go/src/LessonGo/class_gin/src/templates/static"))
	r.StaticFile("/app.js", "/go/src/LessonGo/class_gin/src/templates/static/app.js")
	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

/*
结果集:
root@e7939faf8694:/go/src/LessonGo# curl http://127.0.0.1:8080/assets
<a href="/assets/">Moved Permanently</a>.

root@e7939faf8694:/go/src/LessonGo# curl http://127.0.0.1:8080/more_static
<a href="/more_static/">Moved Permanently</a>.

root@e7939faf8694:/go/src/LessonGo# curl http://127.0.0.1:8080/app.js
root@e7939faf8694:/go/src/LessonGo#
[GIN-debug] Listening and serving HTTP on :8080
[GIN-debug] redirecting request 301: /assets --> /assets/
[GIN-debug] redirecting request 301: /more_static --> /more_static/
[GIN] 2019/06/11 - 01:06:44 | 200 |      8.6593ms |       127.0.0.1 | GET      /app.js

*/

//静态资源嵌入
// loadTemplate 加载由 go-assets-builder 嵌入的模板
/*
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func bindsinglebinarywithtemplate() {
	r := gin.New()

	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/go/src/LessonGo/class_gin/src/templates/index.tmpl", nil)
	})
	r.Run(":8080")

}
请参阅 examples/assets-in-binary 目录中的完整示例。

*/

/*
怎么样编写测试用例:
怎样编写 Gin 的测试用例
HTTP 测试首选 net/http/httptest 包。
package main

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}

上面这段代码的测试用例：
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
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
	//go securejson()
	//go xmljsonyamlprotobuf()
	//go singlefileupload()
	//go multiplefilesupload()
	//go servingdatafromreader()
	//go shutdowngraceful()
	//go usingbasicauthmiddleware()
	//go httpmethods()
	//go usingmiddleware()
	//go onlybindquerystring()
	//go goroutineinmiddleware()
	//go writelog()
	//go defineformatlogs()
	//go mapasquerystringorpostform()
	//go querystringparam()
	//go bindingandvalidation()
	//go bingduri()
	//go bindqueryorpost()
	//go bindformdatarequestwithcustomstruct()
	//customhttpconfig1()
	//customhttpconfig2()
	//go custommiddleware()
	//go customvalidators()
	//go cookie()
	//go paraminpath()
	//go groupingroutes()
	//go runmultipleservice()
	//go servingstaticfiles()
}
