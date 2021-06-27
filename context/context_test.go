package context

import (
	"context"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	defer func() {
		println(1)
	}()
	if 1 != 1 {
		return
	}
	defer func() {
		println(2)
	}()
}

func TestContextValue(t *testing.T) {
	//c:=context.TODO()
	c := context.Background()
	cDeadLine,err:=context.WithDeadline(c,time.Now())
}
