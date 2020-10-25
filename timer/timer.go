package timer

import (
	"fmt"
	"time"
)

func GetTimestamp() int {
	return int(time.Now().Unix())
}

func TimestampToString(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

// format "2006-01-02 15:04:05"
func StringToTimestamp(value, format string) (timestamp int64) {
	parsedTime, err := time.Parse(format, value)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return parsedTime.Unix()
}
