package main

func main()  {
	//log.Fatalf("Fatalf error : %s","1234")

	// 捕获迭代变量
	var rmdirs []func()
	dirs := []string{"1","2","3"}
	for i := 0; i < len(dirs)-1; i++ {
		m:=dirs[i]
		//println(dirs[i]) // OK
		println(m) // OK
		rmdirs = append(rmdirs, func() { // for 循环后才调用函数
			println(m) // ok
			//println(dirs[i]) // NOTE: incorrect!
		})
	}
	for _,v:=range rmdirs {
		v() // debug 后发现 继续调用append(rmdirs, func() {} 的逻辑；
	}
}
