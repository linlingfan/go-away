package container

import (
	"container/list"
	"fmt"
	"testing"
)

func TestContainer(t *testing.T) {
	// 双向列表
	l1 := list.New()
	e := l1.PushBack(4)
	e2 := l1.PushFront(1)
	l1.InsertBefore(3, e)
	l1.InsertAfter(2, e2)

	for e := l1.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// 遍历列表并打印其内容。
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
