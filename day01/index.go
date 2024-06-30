// 单行注释
/**
 * 多行注释
 **/

// 导入包的字句在每个源文件的开头
// main比较特殊，他用来声明可执行文件，而不是一个库

package main

// import 语句声明了当前文件引用的包
import (
	"fmt"       // Go语言标准库中的包
	"io/ioutil" // 包含一些输入输出函数

	// 数学标准库，在此文件中别名为m
	"net/http" // 一个web服务器包
	"os"       // 系统底层函数，如文件读写
	"strconv"  // 字符串转换
)

// 函数声明：main是程序执行的入口
func main() {
	// 标准输出打印一行
	// 用包名fmt 限制打印函数
	fmt.Println("你好世界")

	// 调用当前包的另外一个函数
	beyondHello()
}

// 函数可以在括号里面加参数
// 如果没有参数的话，也需要一个空括号
func beyondHello() {
	var x int // 变量声明，变量必须在使用之前声明
	x = 3     // 变量赋值
	// 可以用 := 来偷懒，他自动把变量类型、声明合赋值都搞定了
	y := 4
	sum, prod := learnMultiple(x, y)        // 返回多个变量的函数
	fmt.Println("sum:", sum, "prod:", prod) // 简单输出
	learnTypes()                            // 少于Y分钟 学的更多
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y // 返回两个值
}

// 内置变量类型和关键字
func learnTypes() {
	// 短声明
	str := "少说话多读书" // String类型

	s2 := `这是一个
	可以换行的字符串
	` // 同样是String类型

	// 非ascii支付。GO使用utf-8编码
	g := 'Σ' // run类型，int32的别名，使用UTF-8编码

	f := 3.14159 // float64 类型， IEEE-754 64位浮点数

	c := 3 + 4i // complex128类型，内部使用两个float64表示

	// var变量可以直接初始化
	var u uint = 7 // unsigned 无符号变量，但是实现依赖init型变量的长度
	var pi float32 = 22. / 7

	// 字符转换
	n := byte('\n') // byte是uint8的别名

	// 数组（Array）类型的大小在编译时即确定
	var a4 [4]int           // 有4个init变量的数组，初始值为0
	a3 := [...]int{3, 1, 5} // 有3个int变量的数组，同时进行了初始化

	// Array 和 slice 各有所长， 但是slice可以动态的增删，所以更多时候还是使用slice
	s3 := []int{4, 5, 9}    // 相比a3没有省略号
	s4 := make([]int, 4)    // 分配4个int大小的内存并初始化为0
	var d2 [][]float64      // 这里只是声明，并未分配内存空间
	bs := []byte("a slice") // 进行类型转换

	// 切片 Slice 的大小是动态的，他的长度可以按需增长
	// 用内置函数 append() 向切片末尾添加元素
	// 要添加到的目标是 append 函数的第一个参数
	// 多数时候 数组在原内存处依次顺次增长 如
	s := []int{1, 2, 3}    // 这是个长度3的slice
	s = append(s, 4, 5, 6) // 在加3元素，长度变为6了
	fmt.Println(s)         // 更新后的数组是 [1 2 3 4 5 6]

	// 除了向append()提供一组原子元素以外
	// 我们还可以用如下方法传递一个 slice 常量或变量，并在后面加上省略号
	// 用以表示我们将引用一个 slice 解包其中的元素，并将其添加到s数组末尾
	s = append(s, []int{7, 8, 9}...) // 第二个参数是一个slice常量
	fmt.Println(s)                   // 更新后的数组是 [1 2 3 4 5 6 7 8 9]

	p, q := learnMemory() // 声明p, q为int型变量的指针
	fmt.Println(*p, *q)   // * 取值

	// Map 是动态可增长关联数组，和其他语言中的 hash 或字典类似
	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1

	// 在GO语言中未使用的变量在编译的时候会报错，而不是warning
	// 下划线 —— 可以使你”使用“一个变量，但是丢弃他的值
	_, _, _, _, _, _, _, _, _, _ = str, s2, g, f, u, pi, n, a3, s4, bs

	// 通常的用法是，在调用拥有多个返回值的函数时，
	// 用下划线抛弃其中的一个参数。
	file, _ := os.Create("output.txt")
	fmt.Fprint(file, "这句代码还展示了如何写入文件呢")
	file.Close()

	// 输出变量
	fmt.Println(s, c, a4, s3, d2, m)

	learnFlowControl() // 回到流程控制
}

// 和其他编程语言不同的是，GO支持有变量名称的变量返回值。
// 声明返回值时带上一个名字允许我们在函数内的不同位置
// 只用写return一个词就能将函数内指定名称的变量返回
func learnNamedReturns(x, y int) (z int) {
	z = x * y
	return // 隐式返回z，因为前面指定了它
}

// GO全面支持垃圾回收。GO有指针，但是不支持指针运算
// 你会因为空指针而报错，但是不会因为增加指针而报错
func learnMemory() (p, q *int) {
	// 返回int型变量指针p和q
	p = new(int) // 内置函数new分配内存
	// 自动将分配的int赋值0，p不再是空的了
	s := make([]int, 20) // 给20个int变量分配一块内存
	s[3] = 7             // 赋值
	r := -2              // 声明另外一个局部变量
	return &s[3], &r     // &取地址
}

func expensiveComputation() int {
	return 1e6
}

func learnFlowControl() {
	// if 需要花括号 括号就免了
	if true {
		fmt.Println("这句话肯定被执行")
	}

	// 用 go fmt 命令可以帮你格式化代码
	if false {
		// pout
	} else {
		// gloat
	}

	// 如果太多嵌套的 if 语句，推荐使用 switch
	x := 1
	switch 1 {
	case 0:
	case 1:
		// 隐式调用break语句，匹配上一个即停止
	case 2:
		// 不会运行
	}

	// 和 if 一样 for 也不用括号
	for x := 0; x < 3; x++ { // ++ 自增
		fmt.Println("遍历", x)
	}
	// x 在这里还是1。为什么？

	// for 是 go 里面唯一的循环关键字，不过他有很多变种
	for { // 死循环
		break    // 中断
		continue // 不会运行
	}

	// 用 range 可以枚举 array、slice、string、map、channel等不同类型
	// 对于 channel，range返回一个值
	// array、slice、string、map等其他类型返回一堆
	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		// 打印 map 中的每一个键值对
		fmt.Printf("索引：%s, 值为：%d\n", key, value)
	}

	// 如果你只是想要值，那就用前面讲的下划线扔掉没用的
	for _, name := range []string{"Bob", "Bill", "Joe"} {
		fmt.Printf("你是。。%s\n", name)
	}

	// 和for一样，if中的 := 先给y赋值，然后再和x做比较
	if y := expensiveComputation(); y > x {
		x = y
	}

	// 闭包函数
	xBig := func() bool {
		return x > 100 // x是上面声明的变量引用
	}

	fmt.Println("xBig:", xBig()) // true
	x /= 1e5                     // x 变成 10
	fmt.Println("xBig", xBig())  // false

	// 除此之外 函数体可以在其他函数定义并调用
	// 满足下列条件时，也可以作为参数传递给其他函数
	// 1、定义的函数被立即调用
	// 2、函数返回值符合调用者对类型的要求
	fmt.Println("两数相加乘二",
		func(a, b int) int {
			return (a + b) * 2
		}(10, 2))

	// 当你需要goto的时候，你会爱死的
	goto love
love:
	learnFunctionFactory() // 返回函数的函数
	learnDefer()           // 对defer关键字的简单介绍
	learnInterfaces()      // 好东西来了
}

func learnFunctionFactory() {
	// 空格分割的两个写法是相同的，不过第二个写法比较实用
	fmt.Println(sentenceFactory("原谅")("当然选择", "她！"))

	d := sentenceFactory("原谅")
	fmt.Println(d("当然选择", "她！"))
	fmt.Println(d("你什么可以", "她？"))
}

// Decorator在一些语言中很常见，在 go 语言中，
// 接受参数作为其定义的一部分的函数是修饰符的替代品
func sentenceFactory(mystring string) func(before, after string) string {
	return func(before, after string) string {
		return fmt.Sprintf("%s %s %s", before, mystring, after) // new string
	}
}

func learnDefer() (ok bool) {
	// defer 表达式在函数返回的前一刻执行
	defer fmt.Println("defer表达式执行顺序为后进先出（LIFO）")
	defer fmt.Println("\n这句话比上句话先输出，因为")
	// 关于defer的用法，例如用defer关闭一个文件
	// 就可以让关闭操作与打开操作的代码更接近一些
	return true
}

// 定义Stringer为一个接口类型，有一个方法String
type Stringer interface {
	String() string
}

// 定义pair为一个结构体，有 x 和 y两个int型变量
type pair struct {
	x, y int
}

// 定义pair类型的方法，实现Stringer接口
func (p pair) String() string { // P 叫做接收器
	// Springtf 是fmt包中的另外一个共有函数
	// 用 . 调用 p 中的元素
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func learnInterfaces() {
	// 花括号用来定义结构体变量，:= 在这里将一个结构体变量赋值给P
	p := pair{3, 4}
	fmt.Println(p.String()) // 调用 pair 类型p 的 String 方法
	var i Stringer          // 声明i为Stringer接口类型
	i = p                   // 有效 因为 p 实现了Stringer接口
	// 调用 i 的 String方法，输出和上面一样
	fmt.Println(i.String())

	// fmt 包中的 Println 函数向他对象要他们的string 输出， 实现了String方法就可以这样使用了

	fmt.Println(p) // 输出和上面一样，自动调用 String 函数
	fmt.Println(i)

	learnVariadicParams("great", "learning", "here!")
}

// 有变长参数列表的函数
func learnVariadicParams(myStrings ...interface{}) {
	// 枚举变量边长参数列表的每个参数值
	// 下划线在这里用来抛弃枚举时候返回的数组索引值
	for _, param := range myStrings {
		fmt.Println("param:", param)
	}

	fmt.Println("params:", fmt.Sprintln(myStrings...))

	learnErrorHandling()
}

func learnErrorHandling() {
	// ", ok" 用来判断有没有正常工作
	m := map[int]string{3: "three", 4: "four"}
	if x, ok := m[1]; !ok {
		fmt.Println("别找了，真没有")
	} else {
		fmt.Print(x) // 如果X在map中的话， x就是那个值
	}

	// 错误可不只是 ok， 它还可以给出关于问题的更多细节
	if _, err := strconv.Atoi("non-int"); err != nil { // _ discards value
		// 输出 “strconv.ParseInt: parsing "non-int" : invalid syntax”
		fmt.Println(err)
	}

	// 待会再说接口吧
	learnConcurrency()
}

// c 是 channel 类型， 一个并发安全的通信对象
func inc(i int, c chan int) {
	c <- i + 1 // <- 把右边的发送到左边的channel
}

// 我们将用 inc 函数来并发的增加一些数字
func learnConcurrency() {
	// 用make来声明一个slice， make会分配和初始化slice， map 和 channel
	c := make(chan int)
	// 用 go 关键字开始三个并发的 goroutine，如果机器支持的话，还可能是并行执行
	// 三个都被发送到同一个 channel
	go inc(0, c) // go is a statement that starts a new goroutine
	go inc(10, c)
	go inc(-805, c)

	// 从 chanel 中读取结果并打印
	// 打印出什么东西结果是不可预知的
	fmt.Println(<-c, <-c, <-c) // channel在右边的时候，<-是读操作

	cs := make(chan string)       // 操作 string 的 channel
	cc := make(chan chan string)  // 操作 channel 的 channel
	go func() { c <- 84 }()       // 开始给一个 goroutine 来发送一个新的数字
	go func() { cs <- "wordy" }() // 发送给cs

	// Select 类似于 switch，但是每个 case 包括一个 channel 的操作
	// 他随机选择一个通讯好的 case
	select {
	case i := <-c: // 从 channel 接收的值可以赋给其他变量
		fmt.Println("这是......", i)
	case <-cs: // 或者直接丢弃
		fmt.Println("这是个字符串！")
	case <-cc: // 空的，还没做好通信的准备
		fmt.Println("别瞎想")
	}
	// 上面 c 或者 cs 的值被取到， 其中一个 goroutine 结束，另外一个一直阻塞

	learnWebProgramming() // GO很适合web编程
}

// http包中的一个简单的函数就可以开启web服务器
func learnWebProgramming() {
	// ListenAndServe 第一个参数指定了监听端口，第二个参数是一个接口，特定是http.Handler
	go func() {
		err := http.ListenAndServe(":8080", pair{})
		fmt.Println(err) // 不要无视错误
	}()

	requestServer()
}

// 使 pair 实现 http.Handler 接口的 ServeHTTP 方法
func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 使用 http.ResponseWriter 返回数据
	w.Write([]byte("Y分钟 golang 速成"))
}

func requestServer() {
	resp, err := http.Get("http://localhost:8080")
	fmt.Println(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll((resp.Body))
	fmt.Printf("\n服务器消息：`%s`", string(body))
}
