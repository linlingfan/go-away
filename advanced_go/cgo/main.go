package main

import "C"

/*
#include <stdio.h>
static void SayHello(const char* s) {
    puts(s);
}
*/
import "C"

func main() {
	println("hello cgo")
	C.puts(C.CString("Hello, World\n"))
	C.SayHello(C.CString("Hello, World\n"))
}
