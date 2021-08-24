## 函数和方法

- 错误
    - 文件结尾错误（EOF）：o包保证任何由文件结束引起的读取失败都返回同一个错误——io.EOF
    -

- 函数值：在Go中，函数被看作第一类值（first-class values）：函数像其他值一样，拥有类型，可以被赋值给其他变量，传递给函数，从函数返回。

    - 函数类型的零值是nil。调用值为nil的函数值会引起panic错误
    ```
    var f func(int) int
    f(3) // 此处f的值为nil, 会引起panic错误
    ```
    - 但是函数值之间是不可比较的，也不能用函数值作为map的key。

- 匿名函数
    - 警告：捕获迭代变量:
      使用go语句（第八章）或者defer语句（5.8节）会经常遇到此类问题。这不是go或defer本身导致的，而是因为它们都会等待循环结束后，再执行函数值。

    ```
	// 捕获迭代变量
	var rmdirs []func()
	dirs := []string{"1","2","3"}
	for i := 0; i < len(dirs)-1; i++ {
		//m:=dirs[i]
		println(dirs[i]) // OK
		//println(m) // OK
		rmdirs = append(rmdirs, func() {
			//println(m) // ok
			println(dirs[i]) // NOTE: incorrect!
		})
	}
	for _,v:=range rmdirs {
		v()
    }

```