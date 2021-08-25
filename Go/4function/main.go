package main

func main() {
	//log.Fatalf("Fatalf error : %s","1234")

	//// 捕获迭代变量
	//var rmdirs []func()
	//dirs := []string{"1", "2", "3"}
	//for i := 0; i < len(dirs)-1; i++ {
	//	m := dirs[i]
	//	//println(dirs[i]) // OK
	//	println(m) // OK
	//	rmdirs = append(rmdirs, func() { // for 循环后才调用函数
	//		println(m) // ok
	//		//println(dirs[i]) // NOTE: incorrect!
	//	})
	//}
	//for _, v := range rmdirs {
	//	v() // debug 后发现 继续调用append(rmdirs, func() {} 的逻辑；
	//}

	p := Point{ID: 1}
	p2 := &Point{ID: 1}
	println(p.Add(1, 2))
	println(p2.Add(1, 3))

	println(uint('A'%64))
	println('0')
	println('1')
	println('a')
}

type Point struct {
	ID int
}

func (p *Point) Add(args ...int) int {
	sum := 0
	for _, v := range args {
		sum = sum + v
	}
	return sum
}
