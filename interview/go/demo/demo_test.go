package demo

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

// Go 语言的 switch 语句虽然没有"break"，但如果 case 完成程序会默认 break，可以在 case 语句后面加上关键字 fallthrough，这样就会接着走下一个 case 语句（不用匹配后续条件表达式）。
// 或者，利用 case 可以匹配多个值的特性。
func TestName1(t *testing.T) {
	isMatch := func(i int) bool {
		switch i {
		case 1:
			fallthrough
		case 2:
			return true
		default:
			return false
		}
		return false
	}

	fmt.Println(isMatch(1)) // true
	fmt.Println(isMatch(2)) // true

	match := func(i int) bool {
		switch i {
		case 1, 2:
			return true
		default:
			return false
		}
	}
	fmt.Println(match(1))
	fmt.Println(match(2))
}

// --
var o = fmt.Print

func TestName2(t1 *testing.T) {
	c := make(chan int, 1)
	for range [3]struct{}{} {
		select {
		default:
			o(1)
		case <-c:
			o(2)
			c = nil
		case c <- 1:
			o(3)
		}
	}
}

// 输出 321
// 第一次循环，写操作已经准备好，执行 o(3)，输出 3；
//第二次，读操作准备好，执行 o(2)，输出 2 并将 c 赋值为 nil；
//第三次，由于 c 为 nil，走的是 default 分支，输出 1。

// ---- 溢出
func TestName3(t *testing.T) {
	var x int8 = -128
	fmt.Println(x)
	var y = x / -1
	fmt.Println(y)
	var n int8 = 127
	m := n + 2
	fmt.Println(m)
}

// ------
func Fu(i int) func() int {
	return func() int {
		i++
		return i
	}
}
func TestName4(t *testing.T) {
	F := Fu(1)
	defer func() {
		fmt.Println(F())
	}()
	defer fmt.Println(F()) // 优先计算 i= i + 1 :2 // 第二个输出 2
	i := F()               // i :3
	fmt.Println(i)         // 第一个输出 3
}

// 324
// defer() 后面的函数如果带参数，会优先计算参数，并将结果存储在栈中，到真正执行 defer() 的时候取出。

// next---
// 将 Mutex 作为匿名字段时，相关的方法必须使用指针接收者，否则会导致锁机制失效。
type data struct {
	//*sync.Mutex // 修复方法二
	sync.Mutex
}

func (d *data) test(s string) { // 修复方法1
	//func (d data) test(s string) {    // 值方法
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func TestName5(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	var d data
	//d := data{new(sync.Mutex)}   // 初始化 修复方法二

	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()
}

//var (
//	size := 1024 // 只能在函数内部使用简短模式
//)

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
	l
)

func TestName6(t *testing.T) {
	fmt.Println(x, y, z, k, p, l)
}

// iota初始值为0，所以x为0，_表示不赋值，但是iota是从上往下加1的，所以y是2，z是“zz”,k和上面一个同值也是“zz”,p是iota,从上0开始数他是5

// next -------
// init()函数
// init()函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等；
//一个包可以出现多个init()函数，一个源文件也可以包含多个init()函数；
//同一个包中多个init()函数的执行顺序没有明确的定义，但是不同包的init函数是根据包导入的依赖关系决定的；
//init函数在代码中不能被显示调用、不能被引用（赋值给函数变量），否则出现编译失败；
//一个包被引用多次，如A import B，C import B，A import C，B被引用多次，但B包只会初始化一次；
//引入包，不可出现死循环。即A import B，B import A，这种情况下编译失败；

// next-----

func TestName7(t *testing.T) {
	str := "hello"
	//str[0] = 'x' // Go 语言中的字符串是只读的。
	fmt.Println(str)

	// nil 切片和 nil 相等，一般用来表示一个不存在的切片；空切片和 nil 不相等，表示一个空的集合
	var s1 []int
	//var s2 = []int{}
	if s1 == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println("no nil")
	}

	//i := 65
	//fmt.Println(string(i))

	s := [3]int{1, 2, 3}
	d := s[1:2:cap(s)]
	println(len(d))
	println(cap(d))
	// 截取操作有带 2 个或者 3 个参数，形如：[i:j] 和 [i:j:k]，假设截取对象的底层数组长度为 l。在操作符 [i:j] 中，如果 i 省略，默认 0，如果 j 省略，默认底层数组的长度，截取得到的切片长度和容量计算方法是 j-i、l-i。
	// 操作符 [i:j:k]，k 主要是用来限制切片的容量，但是不能大于数组的长度 l，截取得到的切片长度和容量计算方法是 j-i、k-i。
}

// next-------

func TestName8(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	fmt.Println(s1)
	s2 = append(s2, 5, 6, 7)
	//s1 = append(s1, 8, 9)
	//fmt.Println(s2)
	fmt.Println(s1)
}

// golang 中切片底层的数据结构是数组。当使用 s1[1:] 获得切片 s2，和 s1 共享同一个底层数组，这会导致 s2[1] = 4 语句影响 s1。
//而 append 操作会导致底层数组扩容，生成新的数组，因此追加数据后的 s2 不会影响 s1。
//但是为什么对 s2 赋值后影响的却是 s1 的第三个元素呢？这是因为切片 s2 是从数组的第二个元素开始，s2 索引为 1 的元素对应的是 s1 索引为 2 的元素。

// next------
func TestName9(t *testing.T) {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// 程序执行到 main() 函数三行代码的时候，会先执行 calc() 函数的 b 参数，即：calc("10",a,b)，输出：10 1 2 3，得到值 3，因为 defer 定义的函数是延迟函数，故 calc("1",1,3) 会被延迟执行；
//程序执行到第五行的时候，同样先执行 calc("20",a,b) 输出：20 0 2 2 得到值 2，同样将 calc("2",0,2) 延迟执行；
//程序执行到末尾的时候，按照栈先进后出的方式依次执行：calc("2",0,2)，calc("1",1,3)，则就依次输出：2 0 2 2，1 1 3 4。

// next-----
// 重点介绍下这个操作符 &^，按位置零，例如：z = x &^ y，表示如果 y 中的 bit 位为 1，则 z 对应 bit 位为 0，否则 z 对应 bit 位等于 x 中相应的 bit 位的值。
// 很多语言都是采用 ~ 作为按位取反运算符，Go 里面采用的是 ^ 。按位取反之后返回一个每个 bit 位都取反的数，对于有符号的整数来说，是按照补码进行取反操作的（快速计算方法：对数 a 取反，结果为 -(a+1) ），对于无符号整数来说就是按位取反。
// 或操作符 | ，表达式 z = x | y，如果 y 中的 bit 位为 1，则 z 对应 bit 位为 1，否则 z 对应 bit 位等于 x 中相应的 bit 位的值，与 &^ 完全相反
func TestName10(t *testing.T) {
	var x uint8 = 214
	var y uint8 = 92
	fmt.Printf("x: %08b\n", x)
	fmt.Printf("y: %08b\n", y)
	fmt.Printf("x ^ y: %08b\n", x^y)
	fmt.Printf("x | y: %08b\n", x|y)
	fmt.Printf("x &^ y: %08b\n", x&^y)
}

func TestName11(t *testing.T) {
	var a int = 64
	//var b string = string(a)
	//println(b)
	c:=strconv.Itoa(a)
	println(c)
}
