package main

/***
https://studygolang.com/articles/5024  os 比较详细
大多数的程序都是处理输入，产生输出；这也正是“计算”的定义。但是, 程序如何获取要处理的输入数据呢？
一些程序生成自己的数据，但通常情况下，输入来自于程序外部：文件、网络连接、其它程序的输出、敲键盘的用户、命令行参数或其它类似输入源。
os包以跨平台的方式，提供了一些与操作系统交互的函数和变量。程序的命令行参数可从os包的Args变量获取；os包外部使用os.Args访问该变量。
os.Args变量是一个字符串（string）的切片（slice）
os.Args[0], 是命令本身的名字；其它的元素则是程序启动时传给它的参数
***/

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func echo() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	//优化，使用range的方式
	s2, sep2 := "", ""
	for _, arg := range os.Args[1:] {
		s2 += sep2 + arg //这种拼接的方式代价比较高
		sep2 = " "
	}
	fmt.Println(s2)
	//优化，直接使用stings来进行拼接
	fmt.Println(strings.Join(os.Args[1:], " "))

}

/***
对文件做拷贝、打印、搜索、排序、统计或类似事情的程序都有一个差不多的程序结构：一个处理输入的循环，在每个元素上执行计算处理，在处理的同时或最后产生输出。
bufio.Scanner、ioutil.ReadFile和ioutil.WriteFile都使用*os.File的Read和Write方法，但是，大多数程序员很少需要直接调用那些低级（lower-level）函数。
高级（higher-level）函数，像bufio和io/ioutil包中所提供的那些，用起来要容易点。
***/
func uniq() {

	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename) //
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") { //ReadFile函数返回一个字节切片（byte slice），必须把它转换为string，才能用strings.Split分割。
			counts[line]++ //统计
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line) //打印重复的行数和内容
		}
	}
}

func main() {
	//echo()
	//uniq()
	/*
			os包提供了操作系统函数的不依赖平台的接口。设计为Unix风格的，虽然错误处理是go风格的；失败的调用会返回错误值而非错误码。
			os包的接口规定为在所有操作系统中都是一致的。非公用的属性可以从操作系统特定的syscall包获取。
			此包含有的功能非常多 ，包含操作系统查询，文件操作，进程操作等。
			const (
			    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
			    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
			    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
			    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
			    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
			    O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
			    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
			    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
				)
			const (
		 		PathSeparator     = '/' // 操作系统指定的路径分隔符
				PathListSeparator = ':' // 操作系统指定的表分隔符
				)
				const DevNull = "/dev/null"
				var (
			   		Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
			    	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
			   		Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
					)
				Stdin、Stdout和Stderr是指向标准输入、标准输出、标准错误输出的文件描述符。
				var Args []string Args保管了命令行参数，第一个是程序名。
	*/
	file, err := os.Open("go.go") // For read access.
	if err != nil {
		log.Println(err)
		//log.Fatal(err)  打印并执行os.exit(1)
	}
	defer file.Close()

	fmt.Println(os.Hostname())                 //获得主机名字 func Hostname() (name string, err error)
	fmt.Println(os.Getpagesize())              //Getpagesize返回底层的系统内存页的尺寸。
	fmt.Println(os.Environ(), os.Environ()[1]) //func Environ() []string Environ返回表示环境变量的格式为"key=value"的字符串的切片拷贝。
	fmt.Println(os.Getenv("OLDPWD"))           //func Getenv(key string) string Getenv检索并返回名为key的环境变量的值。如果不存在该环境变量会返回空字符串。
	fmt.Println(os.Setenv("chenchen", "man"))  //func Setenv(key, value string) error Setenv设置名为key的环境变量。如果出错会返回该错误。貌似只对此进程起作用。
	//os.Clearenv()                              //Clearenv删除所有环境变量。 环境变量的更改都只是对当前进程的添加删除环境变量。
	//Exit让当前程序以给出的状态码code退出。一般来说，状态码0表示成功，非0表示出错。程序会立刻终止，defer的函数不会被执行。
	//os.Exit(0) //func Exit(code int)
	//函数os.Expand()这个其实就是一个回调函数替换的方法 func Expand(s string, mapping func(string) string) string 使用函数来进行替换 Expand函数替换s中的${var}或$var为mapping(var)。
	fmt.Println(os.Getuid())    //Getuid返回调用者的用户ID。
	fmt.Println(os.Geteuid())   //Geteuid返回调用者的有效用户ID。
	fmt.Println(os.Getgid())    //Getgid返回调用者的组ID。
	fmt.Println(os.Getegid())   //Getegid返回调用者的有效组ID。
	fmt.Println(os.Getgroups()) //Getgroups返回调用者所属的所有用户组的组ID。
	fmt.Println(os.Getpid())    //Getpid返回调用者所在进程的进程ID。
	fmt.Println(os.Getppid())   //Getppid返回调用者所在进程的父进程的进程ID。
	/*
			文件操作
			type FileMode uint32 FileMode代表文件的模式和权限位。这些字位在所有的操作系统都有相同的含义，因此文件的信息可以在不同的操作系统之间安全的移植。
			不是所有的位都能用于所有的系统，唯一共有的是用于表示目录的ModeDir位。
			const (
		    // 单字符是被String方法用于格式化的属性缩写。
		    ModeDir        FileMode = 1 << (32 - 1 - iota) // d: 目录
		    ModeAppend                                     // a: 只能写入，且只能写入到末尾
		    ModeExclusive                                  // l: 用于执行
		    ModeTemporary                                  // T: 临时文件（非备份文件）
		    ModeSymlink                                    // L: 符号链接（不是快捷方式文件）
		    ModeDevice                                     // D: 设备
		    ModeNamedPipe                                  // p: 命名管道（FIFO）
		    ModeSocket                                     // S: Unix域socket
		    ModeSetuid                                     // u: 表示文件具有其创建者用户id权限
		    ModeSetgid                                     // g: 表示文件具有其创建者组id的权限
		    ModeCharDevice                                 // c: 字符设备，需已设置ModeDevice
		    ModeSticky                                     // t: 只有root/创建者能删除/移动文件
		    // 覆盖所有类型位（用于通过&获取类型位），对普通文件，所有这些位都不应被设置
		    ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
		    ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
			)

	*/
	filename, _ := os.Stat("os.go") // func Stat(name string) (FileInfo, error)
	//Stat返回一个描述name指定的文件对象的FileInfo。如果指定的文件对象是一个符号链接，返回的FileInfo描述该符号链接指向的文件的信息，本函数会尝试跳转该链接。
	fm := filename.Mode()       // type FileInfo interface  Mode() FileMode     // file mode bits
	fmt.Println(fm.IsDir())     //func (m FileMode) IsDir() bool  IsDir报告m是否是一个目录。
	fmt.Println(fm.IsRegular()) //IsRegular报告m是否是一个普通文件。
	fmt.Println(fm.Perm())      //Perm方法返回m的Unix权限位。    -rw-r--r--
	fmt.Println(fm.String())    //func (m FileMode) String() string
	/*
		FileInfo用来描述一个文件对象。
				type FileInfo interface {
			    Name() string       // 文件的名字（不含扩展名）
			    Size() int64        // 普通文件返回值表示其大小；其他文件的返回值含义各系统不同
			    Mode() FileMode     // 文件的模式位
			    ModTime() time.Time // 文件的修改时间
			    IsDir() bool        // 等价于Mode().IsDir()
			    Sys() interface{}   // 底层数据来源（可以返回nil）
				}
	*/

	//os.lstat() 方法用于类似 stat() 返回文件的信息,但是没有符号链接。
	filename2, _ := os.Lstat("os.go")
	fmt.Println(filename)
	fmt.Println(filename2)
	//所有的限定都只是当前进程中的修改，不涉及外部真实切换。
	//windows 下两个都是true IsPathSeparator返回字符c是否是一个路径分隔符。
	fmt.Println(os.IsPathSeparator('/'))
	fmt.Println(os.IsPathSeparator('\\'))
	fmt.Println(os.IsPathSeparator('.'))
	//func IsExist(err error) bool 返回一个布尔值说明该错误是否表示一个文件或目录已经存在。ErrExist和一些系统调用错误会使它返回真。
	//func IsNotExist(err error) bool 返回一个布尔值说明该错误是否表示一个文件或目录不存在。ErrNotExist和一些系统调用错误会使它返回真。
	//func IsPermission(err error) bool 返回一个布尔值说明该错误是否表示因权限不足要求被拒绝。ErrPermission和一些系统调用错误会使它返回真。
	fmt.Println(os.Chdir("/root")) //func Chdir(dir string) error  Chdir将当前工作目录修改为dir指定的目录。
	fmt.Println(os.Getwd())        // func Getwd() (dir string, err error) Getwd返回一个对应当前工作目录的根路径。如果当前目录可以经过多条路径抵达（因为硬链接），Getwd会返回其中一个。
	os.Chmod("os.go", 0777)        //Chmod修改name指定的文件对象的mode。如果name指定的文件是一个符号链接，它会修改该链接的目的地文件的mode。
	os.Chown("os.go", 5, 60)       // func Chown(name string, uid, gid int) error Chmod修改name指定的文件对象的用户id和组id。如果name指定的文件是一个符号链接，它会修改该链接的目的地文件的用户id和组id。
	os.Lchown("os.go", 5, 60)      //Chmod修改name指定的文件对象的用户id和组id。如果name指定的文件是一个符号链接，它会修改该符号链接自身的用户id和组id。
	fmt.Println(filename2.Mode())
	//func Chtimes(name string, atime time.Time, mtime time.Time) error Chtimes修改name指定的文件对象的访问时间和修改时间，类似Unix的utime()或utimes()函数。底层的文件系统可能会截断/舍入时间单位到更低的精确度。
	fmt.Println(os.Mkdir("abc", os.ModePerm)) // Mkdir使用指定的权限和名称创建一个目录。func Mkdir(name string, perm FileMode) error
	//func MkdirAll(path string, perm FileMode) error MkdirAll使用指定的权限和名称创建一个目录，包括任何必要的上级目录，并返回nil，否则返回错误。权限位perm会应用在每一个被本函数创建的目录上。如果path指定了一个已经存在的目录，MkdirAll不做任何操作并返回nil。
	fmt.Println(os.MkdirAll("abc", os.ModePerm))
	fmt.Println(os.Rename("abc", "efge")) //Rename修改一个文件的名字，移动一个文件。
	//func Truncate(name string, size int64) error  Truncate修改name指定的文件的大小。如果该文件为一个符号链接，将修改链接指向的文件的大小。
	fmt.Println(os.Remove("abc")) //func Remove(name string) error Remove删除name指定的文件或目录。如果出错，会返回*PathError底层类型的错误。
	//func RemoveAll(path string) error   RemoveAll删除path指定的文件，或目录及它包含的任何下级对象。它会尝试删除所有东西，除非遇到错误并返回。
	//func Readlink(name string) (string, error) Readlink获取name指定的符号链接文件指向的文件的路径。
	//func Symlink(oldname, newname string) error Symlink创建一个名为newname指向oldname的符号链接。
	//func Link(oldname, newname string) error Link创建一个名为newname指向oldname的硬链接。
	//func SameFile(fi1, fi2 FileInfo) bool  SameFile返回fi1和fi2是否在描述同一个文件。例如，在Unix这表示二者底层结构的设备和索引节点是相同的；在其他系统中可能是根据路径名确定的。SameFile应只使用本包Stat函数返回的FileInfo类型值为参数，其他情况下，它会返回假。
	fmt.Println(os.TempDir()) //TempDir返回一个用于保管临时文件的默认目录。
	//File代表一个打开的文件对象。
	//func Create(name string) (file *File, err error) Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）。如果成功，返回的文件对象可用于I/O；对应的文件描述符具有O_RDWR模式。
	//func Open(name string) (file *File, err error) Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符具有O_RDONLY模式。
	/*
			func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
			OpenFile是一个更一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数。它会使用指定的选项（如O_RDONLY等）、指定的模式（如0666等）打开指定名称的文件。如果操作成功，返回的文件对象可用于I/O。

			func NewFile(fd uintptr, name string) *File NewFile使用给出的Unix文件描述符和名称创建一个文件。
			Stdin  := os.NewFile(uintptr(syscall.Stdin), "/dev/stdin")
		    Stdout := os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
			Stderr := os.NewFile(uintptr(syscall.Stderr), "/dev/stderr")
			func Pipe() (r *File, w *File, err error) Pipe返回一对关联的文件对象。从r的读取将返回写入w的数据。本函数会返回两个文件对象和可能的错误。
	*/
	//先file类型 再FileInfo类型
	fmt.Println(file.Name()) //Name方法返回（提供给Open/Create等方法的）文件名称。
	fmt.Println(file.Fd())   //Fd返回与文件f对应的整数类型的Unix文件描述符。
	/*
			func (f *File) Chdir() error Chdir将当前工作目录修改为f，f必须是一个目录。
			func (f *File) Chmod(mode FileMode) error Chmod修改文件的模式。
			func (f *File) Chown(uid, gid int) error  Chown修改文件的用户ID和组ID。
			func (f *File) Readdir(n int) (fi []FileInfo, err error)  Readdir读取目录f的内容，返回一个有n个成员的[]FileInfo，这些FileInfo是被Lstat返回的，采用目录顺序。对本函数的下一次调用会返回上一次调用剩余未读取的内容的信息。
			func (f *File) Readdirnames(n int) (names []string, err error) Readdir读取目录f的内容，返回一个有n个成员的[]string，切片成员为目录中文件对象的名字，采用目录顺序。对本函数的下一次调用会返回上一次调用剩余未读取的内容的信息。
			func (f *File) Truncate(size int64) error Truncate改变文件的大小，它不会改变I/O的当前位置。 如果截断文件，多出的部分就会被丢弃。
			func (f *File) Read(b []byte) (n int, err error) Read方法从f中读取最多len(b)字节数据并写入b。它返回读取的字节数和可能遇到的任何错误。文件终止标志是读取0个字节且返回值err为io.EOF。
			b := make([]byte, 100)
		    ff, _ := os.Open("main.go")
			n, _ := ff.Read(b)
			fmt.Println(n)
			fmt.Println(string(b[:n]))
			func (f *File) ReadAt(b []byte, off int64) (n int, err error) ReadAt从指定的位置（相对于文件开始位置）读取len(b)字节数据并写入b。它返回读取的字节数和可能遇到的任何错误。当n<len(b)时，本方法总是会返回错误；如果是因为到达文件结尾，返回值err会是io.EOF。
			func (f *File) WriteString(s string) (ret int, err error)  WriteString类似Write，但接受一个字符串参数。
			func (f *File) WriteAt(b []byte, off int64) (n int, err error) WriteAt在指定的位置（相对于文件开始位置）写入len(b)字节数据。它返回写入的字节数和可能遇到的任何错误。如果返回值n!=len(b)，本方法会返回一个非nil的错误。
			func (f *File) Seek(offset int64, whence int) (ret int64, err error) Seek设置下一次读/写的位置。offset为相对偏移量，而whence决定相对位置：0为相对文件开头，1为相对当前位置，2为相对文件结尾。它返回新的偏移量（相对开头）和可能的错误。
			func (f *File) Sync() (err error) Sync递交文件的当前内容进行稳定的存储。一般来说，这表示将文件系统的最近写入的数据在内存中的拷贝刷新到硬盘中稳定保存。//同步操作，将当前存在内存中的文件内容写入硬盘．
			file, _ := os.Create("tmp.txt")
			a := "hellobyte"
			file.WriteAt([]byte(a), 10)      //在文件file偏移量10处开始写入hellobyte
			file.WriteString("string")　　//在文件file偏移量0处开始写入string
			file.Write([]byte(a))                //在文件file偏移量string之后开始写入hellobyte，这个时候就会把开始利用WriteAt在offset为10处开始写入的hellobyte进行部分覆盖
			b := make([]byte, 20)
			file.Seek(0, 0)          　　　　//file指针指向文件开始位置
			n, _ := file.Read(b)
			fmt.Println(string(b[:n]))  //stringhellobytebyte，这是由于在写入过程中存在覆盖造成的
			func (f *File) Close() error Close关闭文件f，使文件不能用于读写。它返回可能出现的错误。
			type ProcAttr struct  ProcAttr保管将被StartProcess函数用于一个新进程的属性。
			type Process struct Process保管一个被StarProcess创建的进程的信息。
	*/
	fmt.Println(os.FindProcess(324)) //FindProcess根据进程id查找一个运行中的进程。函数返回的进程对象可以用于获取其关于底层操作系统进程的信息。
	//func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error) //StartProcess使用提供的属性、程序名、命令行参数开始一个新进程。StartProcess函数是一个低水平的接口。os/exec包提供了高水平的接口，应该尽量使用该包。
	//func (p *Process) Signal(sig Signal) error  Signal方法向进程发送一个信号。在windows中向进程发送Interrupt信号尚未实现。
	//func (p *Process) Kill() error  Kill让进程立刻退出。
	//func (p *Process) Wait() (*ProcessState, error)  Wait方法阻塞直到进程退出，然后返回一个描述ProcessState描述进程的状态和可能的错误。Wait方法会释放绑定到进程p的所有资源。在大多数操作系统中，进程p必须是当前进程的子进程，否则会返回错误。
	//func (p *Process) Release() error Release释放进程p绑定的所有资源， 使它们（资源）不能再被（进程p）使用。只有没有调用Wait方法时才需要调用本方法。
	//ProcessState保管Wait函数报告的某个已退出进程的信息。
	//func (p *ProcessState) Pid() int Pi返回一个已退出的进程的进程id。
	//func (p *ProcessState) Exited() bool  Exited报告进程是否已退出。
	//func (p *ProcessState) Success() bool Success报告进程是否成功退出，如在Unix里以状态码0退出。
	//func (p *ProcessState) SystemTime() time.Duration SystemTime返回已退出进程及其子进程耗费的系统CPU时间。
	//func (p *ProcessState) UserTime() time.Duration UserTime返回已退出进程及其子进程耗费的用户CPU时间。
	//func (p *ProcessState) Sys() interface{} Sys返回该已退出进程系统特定的退出信息。需要将其类型转换为适当的底层类型，如Unix里转换为*syscall.WaitStatus类型以获取其内容。
	//func (p *ProcessState) SysUsage() interface{} SysUsage返回该已退出进程系统特定的资源使用信息。需要将其类型转换为适当的底层类型，如Unix里转换为*syscall.Rusage类型以获取其内容。
	//func (p *ProcessState) String() string
	/*
		attr := &os.ProcAttr{
		        Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}, //其他变量如果不清楚可以不设定
		    }
		    p, err := os.StartProcess("/usr/bin/vim", []string{"/usr/bin/vim", "tmp.txt"}, attr) //vim 打开tmp.txt文件
		    if err != nil {
		        fmt.Println(err)
		    }
		    go func() {
		        p.Signal(os.Kill) //kill process
		    }()

		    pstat, err := p.Wait()
		    if err != nil {
		        fmt.Println(err)
		    }

		    fmt.Println(pstat) //signal: killed
	*/
}

/*
结果集:
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage# go run os.go
e7939faf8694 <nil>
4096
[GOLANG_VERSION=1.12.4 LANG=zh_CN.UTF-8 HOSTNAME=e7939faf8694 AMD_ENTRYPOINT=vs/agent/remoteExtensionHostProcess OLDPWD=/root APPLICATION_INSIGHTS_NO_DIAGNOSTIC_CHANNEL=true GOPATH=/go PWD=/go/src/LessonGo/class_tour/test6selfpackage HOME=/root TERM_PROGRAM=vscode TERM_PROGRAM_VERSION=1.35.0-insider VSCODE_IPC_HOOK_CLI=/tmp/vscode-ipc-b6ed9f6a-1ea8-4493-b223-71aca652ce1b.sock TERM=xterm-256color SHLVL=1 PIPE_LOGGING=true sss=1111 PATH=/root/.vscode-server-insiders/bin/a429fd13565d3392c87e31228ec9e8a2bf3708d5/bin:/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin VERBOSE_LOGGING=true _=/usr/local/go/bin/go] LANG=zh_CN.UTF-8
/root
<nil>
0
0
0
0
[] <nil>
7434
7389
false
true
-rw-r--r--
-rw-r--r--
&{os.go 21275 420 {930031600 63695169240 0x57c140} {2049 1707780 1 33188 0 0 0 0 21275 4096 48 {1559572441 40031600} {1559572440 930031600} {1559572440 930031600} [0 0 0]}}
&{os.go 21275 420 {930031600 63695169240 0x57c140} {2049 1707780 1 33188 0 0 0 0 21275 4096 48 {1559572441 40031600} {1559572440 930031600} {1559572440 930031600} [0 0 0]}}
true
false
false
<nil>
/root <nil>
-rw-r--r--
<nil>
<nil>
rename abc efge: file exists
<nil>
/tmp
go.go
3
&{324 0 0 {{0 0} 0 0 0 0}} <nil>
root@e7939faf8694:/go/src/LessonGo/class_tour/test6selfpackage#

*/
