package traffic

import (
	"fmt"
	"time"
)

type SlideWindow struct {

	Count int64 // 当前流量

	MaxLimit int32 // 每秒最大限制
	
}

func SlideWindowLimit() bool  {
	curTime:=time.Now().Unix()
	time.Sleep(time.Second)
	enTime := time.Now().Unix()
	fmt.Printf("%+d",enTime-curTime)
	return true
}
