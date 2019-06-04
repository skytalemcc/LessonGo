### 0.0.1 Gin是什么
Gin 是一个用 Go (Golang) 编写的 HTTP web 框架。它是一个类似于 martini 但拥有更好性能的 API 框架, 由于 httprouter，速度提高了近 40 倍。  
Gin是一个golang的微框架，封装比较优雅，API友好，源码注释比较明确。具有快速灵活，容错方便等特点。其实对于golang而言，web框架的依赖要远比Python，Java之类的要小。自身的net/http足够简单，性能也非常不错。框架更像是一些常用函数或者工具的集合。借助框架开发，不仅可以省去很多常用的封装带来的时间，也有助于团队的编码风格和形成规范。  
特性:  
快速: 基于 Radix 树的路由，小内存占用。没有反射。可预测的 API 性能。  
支持中间件：传入的 HTTP 请求可以由一系列中间件和最终操作来处理。 例如：Logger，Authorization，GZIP，最终操作 DB。  
Crash：Gin 可以 catch 一个发生在 HTTP 请求中的 panic 并 recover 它。这样，你的服务器将始终可用。例如，你可以向 Sentry 报告这个 panic！  
JSON 验证：Gin 可以解析并验证请求的 JSON，例如检查所需值的存在。  
路由组：更好地组织路由。是否需要授权，不同的 API 版本…… 此外，这些组可以无限制地嵌套而不会降低性能。  
错误管理：Gin 提供了一种方便的方法来收集 HTTP 请求期间发生的所有错误。最终，中间件可以将它们写入日志文件，数据库并通过网络发送。  
内置渲染：Gin 为 JSON，XML 和 HTML 渲染提供了易于使用的 API。  
可扩展性：新建一个中间件非常简单。  
学习前了解项目管理  
1，GOROOT golang安装路径，无需设置。  
root@e7939faf8694:/go# go env |grep GOROOT  
GOROOT="/usr/local/go"  
2，GOPATH golang中最重要的一个变量，是项目的工作空间。  
GOPATH下会有3个目录：src、bin、pkg。  
src目录：go编译时查找代码的地方；  
bin目录：go get这种bin工具的时候，二进制文件下载的目的地；  
pkg目录：编译生成的lib文件存储的地方。  
root@e7939faf8694:/go# go env |grep GOPATH  
GOPATH="/go"  
root@e7939faf8694:/go#  
3，GOBIN。go install编译存放路径。可以为空。为空时可执行文件放在各自GOPATH目录的bin文件夹中。  
如果GOBIN为空最好将GOPATH的bin目录添加到环境变量PATH中，以便在shell直接运行可执行文件名。如安装glide后键入glide。  
root@e7939faf8694:/go# go env |grep GOBIN  
GOBIN=""  
root@e7939faf8694:/go#  
root@e7939faf8694:/go# env |grep PATH  
GOPATH=/go  
PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin  
root@e7939faf8694:/go#   
### 0.0.2 快速入门
要安装 Gin 软件包，需要先安装 Go 并设置 Go 工作区。  
export GOPATH=/go/src/LessonGo/class_gin/  更改为此项目单独的GOPATH路径。这样需要的包就会被下载到此路径的pkg和src里面。  
下载并安装：go get -u -v github.com/gin-gonic/gin  下载消耗几分钟时间，请耐心等待。-v 显示操作流程的日志及信息，方便检查错误 -u 下载丢失的包，但不会更新已经存在的包。  
将 gin 引入到代码中：import "github.com/gin-gonic/gin"  
（可选）如果使用诸如 http.StatusOK 之类的常量，则需要引入 net/http 包：import "net/http"  
下载的过程中可能会出现由于墙的问题导致的无法下载的问题。例如package golang.org/x/sys/unix: unrecognized import path "golang.org/x/sys/unix" (https fetch: Get https://golang.org/x/sys/unix?go-get=1: dial tcp 216.239.37.1:443: connect: connection refused)。遇到这种问题，只能自己手动下载。https://golang.org/x/ 只能选择替换为github里面的mirror镜像。此问题要mkdir -p $GOPATH/src/golang.org/x；cd $GOPATH/src/golang.org/x；然后进入x目录下使用命令
git clone git@github.com:golang/sys.git 即可把需要的包下载下来。这样就解决了某个包无法下载的问题。如下过程可以看到已经存在的包被跳过了。相当于校验补充。  
```  
root@e7939faf8694:/go/src/LessonGo/class_gin/src# go get -u -v github.com/gin-gonic/gin
github.com/gin-gonic/gin (download)
github.com/gin-contrib/sse (download)
github.com/golang/protobuf (download)
github.com/ugorji/go (download)
Fetching https://gopkg.in/go-playground/validator.v8?go-get=1
Parsing meta tags from https://gopkg.in/go-playground/validator.v8?go-get=1 (status code 200)
get "gopkg.in/go-playground/validator.v8": found meta tag get.metaImport{Prefix:"gopkg.in/go-playground/validator.v8", VCS:"git", RepoRoot:"https://gopkg.in/go-playground/validator.v8"} at https://gopkg.in/go-playground/validator.v8?go-get=1
gopkg.in/go-playground/validator.v8 (download)
Fetching https://gopkg.in/yaml.v2?go-get=1
Parsing meta tags from https://gopkg.in/yaml.v2?go-get=1 (status code 200)
get "gopkg.in/yaml.v2": found meta tag get.metaImport{Prefix:"gopkg.in/yaml.v2", VCS:"git", RepoRoot:"https://gopkg.in/yaml.v2"} at https://gopkg.in/yaml.v2?go-get=1
gopkg.in/yaml.v2 (download)
github.com/mattn/go-isatty (download)
Fetching https://golang.org/x/sys/unix?go-get=1
https fetch failed: Get https://golang.org/x/sys/unix?go-get=1: dial tcp 216.239.37.1:443: i/o timeout
golang.org/x/sys (download)
github.com/gin-gonic/gin/internal/json
github.com/gin-contrib/sse
github.com/golang/protobuf/proto
github.com/ugorji/go/codec
gopkg.in/go-playground/validator.v8
gopkg.in/yaml.v2
golang.org/x/sys/unix
github.com/mattn/go-isatty
github.com/gin-gonic/gin/render
github.com/gin-gonic/gin/binding
github.com/gin-gonic/gin
root@e7939faf8694:/go/src/LessonGo/class_gin/src#
```  
其它 golang.org/x 下的包获取皆可使用该方法。很多go的软件在编译时都要使用tools里面的内容，使用下面方法获取：进入上面的x目录下，输入：git clone https://github.com/golang/tools.git。注意，一定要保持与go get获取的目录结构是一致的，否则库就找不到了。  
### 0.0.3 使用 Govendor 工具创建项目
1，go get govendor   $ go get -u -v github.com/kardianos/govendor  
2，创建项目并且 cd 到项目目录中 $ mkdir -p $GOPATH/src/github.com/myusername/project && cd "$_"  
3，使用 govendor 初始化项目，并且引入 gin $ govendor init $ govendor fetch github.com/gin-gonic/gin@v1.3  
4，复制启动文件模板到项目目录中 $ curl https://raw.githubusercontent.com/gin-gonic/examples/master/basic/main.go > main.go  
5，启动项目 $ go run main.go  
### 0.0.4 手动编写和管理项目
基准测试 Gin 使用了自定义版本的 HttpRouter 具有高总调用数，低次操作耗时（ns/op），低堆内存分配 （B/op），低每次操作的平均内存分配次数（allocs/op）。  
使用 jsoniter 编译 Gin 使用 encoding/json 作为默认的 json 包，但是你可以在编译中使用标签将其修改为 jsoniter。  
$ go build -tags=jsoniter . 在src目录下进行编译的时候使用 ，指定使用tag来进行编译。  

