package demo

import (
	"fmt"
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

