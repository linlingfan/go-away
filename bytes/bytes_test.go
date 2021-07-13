package bytes

import (
	"bytes"
	"fmt"
	"testing"
)

// char ASCII码 集合
// 字符串操作方法的底实现参照bytes实现
func TestBytes(t *testing.T) {
	var a = []byte{65, 97, 98, 255}
	var b = []byte{97, 98}
	result := bytes.Contains(a, b)
	println(result)

	var m = []byte("ABC")
	fmt.Printf("%+v \n", m)
	// 但凡有一个字符 包含在 m中
	println(fmt.Sprintf("ContainsAny:%+v", bytes.ContainsAny(m, "AD")))
	println(fmt.Sprintf("ContainsRune:%+v", bytes.ContainsRune(m, 65)))
	println(fmt.Sprintf("ContainsRune:%+v", bytes.ContainsRune([]byte("你是猪"), '猪')))

	var c = []byte{2, 3, 65, 97, 1}
	result2 := bytes.IndexAny(c, "aA") // 第一个字符所在的字节数组下标（不存在返回-1）
	println(result2)

	reader := bytes.NewReader(c)
	println(reader.Size())
	println(reader.Len())
}

func TestByteReader(t *testing.T) {
	var c = []byte{2, 3, 65, 97, 1}
	reader := bytes.NewReader(c)
	println(reader.Size())
	println(reader.Len())
}

func TestByteBuffer(t *testing.T) {
	var c = []byte{65, 97, 98}
	buffer := bytes.NewBuffer(c)
	println(buffer.String())
	println(buffer.Len())

	l, e := buffer.WriteString("ABCDE")
	if e != nil {
		fmt.Printf("%+v", e)
	}
	println(l)
}

func TestTowStrSum(t *testing.T){
	m:=[]byte("923")
	println(m[0])
	println(m[1])
	println('0')

	println(m[0]+m[1] - '0' - '0')

	println(byte('9') + byte(1))
}