package main

/*
tar包实现了tar格式压缩文件的存取。 打包、解包
tar包实现了文件的打包功能,可以将多个文件或者目录存储到单一的.tar压缩文件中
tar本身不具有压缩功能,只能打包文件或目录
tar是一种打包格式，但不对文件进行压缩，所以打包后的文档一般远远大于zip和tar.gz，因为不需要压缩的原因，所以打包的速度是非常快的，打包时CPU占用率也很低。
ar的目的是什么？
方便文件的管理（帮助理解：就是你存在很多文件的时候，但是你很多要很长时间不去接触的话，你想要变得更加简洁，可以进行tar操作，就可以变得更简洁。
比如就像生活中，有很多小箱子分散在不同的房间里，可以将小箱子叠起来放在一个房间里，tar可以类似这样）

Header代表tar档案文件里的单个头。Header类型的某些字段可能未使用。
type Header struct {
    Name       string    // 记录头域的文件名
    Mode       int64     // 权限和模式位
    Uid        int       // 所有者的用户ID
    Gid        int       // 所有者的组ID
    Size       int64     // 字节数（长度）
    ModTime    time.Time // 修改时间
    Typeflag   byte      // 记录头的类型
    Linkname   string    // 链接的目标名
    Uname      string    // 所有者的用户名
    Gname      string    // 所有者的组名
    Devmajor   int64     // 字符设备或块设备的major number
    Devminor   int64     // 字符设备或块设备的minor number
    AccessTime time.Time // 访问时间
    ChangeTime time.Time // 状态改变时间
    Xattrs     map[string]string
}

// 有四个变量，分别是写内容太多，头部信息太长，关闭错误，以及无效tar头部信息
var (
    ErrWriteTooLong    = errors.New("archive/tar: write too long")
    ErrFieldTooLong    = errors.New("archive/tar: header field too long")
    ErrWriteAfterClose = errors.New("archive/tar: write after close")
)

var (
    ErrHeader = errors.New("archive/tar: invalid tar header")
)

*/
import (
	"archive/tar"
	"io"
	"log"
	"os"
)

func tarfile() {
	/*
		打包操作：
		1、生成打包后的目标文件
		2、获取要打包的文件集
		3、往目标文件写入文件
	*/

	unTarFileNames := []string{"log.go", "net.go"} //如果多个文件打包 就使用[]string ，并使用for循环来在最后写入的时候 多次读取文件写入。

	//destfile, err := os.OpenFile("test.tar", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)  //这个是存在文件就继续追加的动作。多次执行就连续写入tar文件
	destfile, err := os.Create("test.tar") //建议用这个，创建文件，如果存在文件就清空。
	if err != nil {
		log.Fatalln("fail to create test.tar file!") //创建文件就直接失败没有成功，则直接返回即可
	}
	defer destfile.Close() //先入后出

	// 通过destfile创建一个tar.Writer
	tw := tar.NewWriter(destfile) //func NewWriter(w io.Writer) *Writer 创建一个*writer
	// 如果关闭失败会造成tar包不完整
	defer func() { //后入先出
		if err := tw.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	//创建目标文件后，开始写入数据

	for _, unTarFileName := range unTarFileNames {
		//理论上 有error返回的必须加error判断语句来返回错误。
		sfileInfo, _ := os.Stat(unTarFileName)      //获得文件信息 func Stat(name string) (FileInfo, error)
		hdr, _ := tar.FileInfoHeader(sfileInfo, "") //FileInfoHeader返回一个根据fi填写了部分字段的Header。 输入os.FileInfo 获得*Header
		// 将tar的文件信息hdr写入到tw
		tw.WriteHeader(hdr) //func (tw *Writer) WriteHeader(hdr *Header) error   将获得的Header写入到Writer里面。

		// 将源文件数据写入目标文件
		fs, _ := os.Open(unTarFileName) //func Open(name string) (*File, error) os.File 同时实现了io.Reader和io.Writer

		if _, err = io.Copy(tw, fs); err != nil { //func Copy(dst Writer, src Reader) (written int64, err error) 将源文件写入到目标文件
			log.Fatalln(err)
		}
		fs.Close()
	}
}

/*
结果集:
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# go run archive.go
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# tar -tvf test.tar
-rw-r--r-- root/root      9003 2019-05-30 13:04 log.go
-rw-r--r-- root/root      2956 2019-05-27 08:06 net.go
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage#
*/

func untarfile() {
	/*
		解包操作：
		1、打开tar文件
		2、遍历tar中文件信息
		3、创建文件，写入，保存，关闭文件
	*/
	srcFile := "test.tar"
	// 打开 tar 包
	fr, err := os.Open(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fr.Close()

	tr := tar.NewReader(fr) //NewReader创建一个从r读取的Reader。

	//Reader提供了对一个tar档案文件的顺序读取。一个tar档案文件包含一系列文件。
	//Next方法返回档案中的下一个文件（包括第一个），返回值可以被视为io.Reader来获取文件的数据。
	for hdr, err := tr.Next(); err != io.EOF; hdr, err = tr.Next() { //Next 转入tar档案文件下一记录，它会返回下一记录的头域。
		if err != nil {
			log.Println(err)
			continue
		}
		// 读取文件信息
		fi := hdr.FileInfo()

		// 创建一个空文件，用来写入解包后的数据
		fw, err := os.Create(fi.Name())
		if err != nil {
			log.Println(err)
			continue
		}

		if _, err := io.Copy(fw, tr); err != nil {
			log.Println(err)
		}
		os.Chmod(fi.Name(), fi.Mode().Perm())
		fw.Close()
	}
}

/*
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# go  run archive.go
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# ls -ltr
结果集:
-rw-r--r-- 1 root root  4789 May 31 13:15 archive.go
-rw-r--r-- 1 root root  2956 May 31 13:15 net.go
-rw-r--r-- 1 root root  9003 May 31 13:15 log.go
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage#
*/

func main() {
	tarfile()
	untarfile()
}

/*
tar包深入学习
打包和解包的原理和实现
1、打包实现原理
先创建一个文件x.tar，然后向x.tar写入tar头部信息。打开要被tar的文件，向x.tar写入头部信息，然后向x.tar写入文件信息。
重复第二步直到所有文件都被写入到x.tar中，关闭x.tar，整个过程就这样完成了 。--关闭源文件，关闭writer,关闭目标文件。
2、解包实现原理
先打开tar文件，然后从这个tar头部中循环读取存储在这个归档文件内的文件头信息，从这个文件头里读取文件名，以这个文件名创建文件，然后向这个文件里写入数据。

打包：
// 接下来对底层实现进行分析
tr := tar.NewReader(fr)
hdr, err := tar.FileInfoHeader(fi, "")
// 将tar的文件信息hdr写入到tw
err = tw.WriteHeader(hdr)

解包：
    fr, err := os.Open(srcFile)
    tr := tar.NewReader(fr)
    hdr, err := tr.Next()
    fi := hdr.FileInfo()
    fw, err := os.Create(fi.Name())
    io.Copy(fw, tr)
	os.Chmod(fi.Name(), fi.Mode().Perm())


func NewReader(r io.Reader) *Reader NewReader创建一个从r读取的Reader。
func (tr *Reader) Next() (*Header, error) 转入tar档案文件下一记录，它会返回下一记录的头域。
func (tr *Reader) Read(b []byte) (n int, err error) 从档案文件的当前记录读取数据，到达记录末端时返回(0, EOF)，直到调用Next方法转入下一记录。
func NewWriter(w io.Writer) *Writer NewWriter创建一个写入w的*Writer。
func (tw *Writer) WriteHeader(hdr *Header) error WriteHeader写入hdr并准备接受文件内容。
func (tw *Writer) Write(b []byte) (n int, err error) Write向tar档案文件的当前记录中写入数据。如果写入的数据总数超出上一次调用WriteHeader的参数hdr.Size字节，返回ErrWriteTooLong错误。
func (tw *Writer) Flush() error Flush结束当前文件的写入。
func (tw *Writer) Close() error Close关闭tar档案文件，会将缓冲中未写入下层的io.Writer接口的数据刷新到下层。

*/
