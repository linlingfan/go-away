package main

import (
	"fmt"
)

func main() {
	//BitTest()
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"} // 指定索引初始化
	fmt.Println(RMB, symbol[RMB])                                 // "3 ￥"

	//var m map[string]int // nil
	m := make(map[string]int)
	m["123"] = 123
	println(m["123"])

	p := Point{1, 2}
	q := Point{1, 2}
	fmt.Println(p.X == q.X && p.Y == q.Y) // "true"
	fmt.Println(p == q)                   // "true"
	fmt.Println(&p == &q)                 // "false"
}

func BitTest() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("&:%08b\n", x&y)   // "00000010", the intersection {1}
	fmt.Printf("|:%08b\n", x|y)   // "00100110", the union {1, 2, 5}
	fmt.Printf("^:%08b\n", x^y)   // "00100100", the symmetric difference {2, 5}
	fmt.Printf("&^:%08b\n", x&^y) // "00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
}

type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

type Point struct{ X, Y int }
