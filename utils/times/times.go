package times

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	DateTimeTemplate      = "2006-01-02 15:04:05"
	DateTemplate          = "2006-01-02"
	DateTimeShortTemplate = "20060102150405"
)

func CompareNowDateTime(dataStr string) bool {
	now := time.Now()
	local, _ := time.LoadLocation("Local")
	dateTime, _ := time.ParseInLocation(DateTimeTemplate, dataStr, local)
	return now.After(dateTime)
}

func GetNowTimeFormatDate() string {
	return time.Now().In(time.Local).Format(DateTemplate)
}

func ParseTimeWithTemplate(timeStr, template string) (time.Time, error) {
	local, _ := time.LoadLocation("Local")
	dateTime, err := time.ParseInLocation(template, timeStr, local)
	return dateTime, err
}

func FormatTimeWithTemplate(dateTime time.Time, template string) string {
	return dateTime.In(time.Local).Format(template)
}

func GetRandCode(width int) string {
	pools := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(pools)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", pools[rand.Intn(r)])
	}
	return sb.String()
}

func GetOutOrderNo() string {
	dateTime := FormatTimeWithTemplate(time.Now(), DateTimeShortTemplate)
	randCode := GetRandCode(6)
	return dateTime + randCode
}
