package main

func main() {
	x, y := 0, 1
	n := 10
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	println(x)
}
