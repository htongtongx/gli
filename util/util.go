package util

import (
	"os"
	"strings"
	"time"
)

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func IsInTests() bool {
	return strings.HasSuffix(os.Args[0], ".test")
	// for _, arg := range os.Args {
	// 	if strings.HasPrefix(arg, "-test.v=") {
	// 		return true
	// 	}
	// }
	// return false
}

func ReverseString(text string) string {
	var reslut []byte
	for i := len(text) - 1; i >= 0; i-- {
		reslut = append(reslut, text[i])
	}
	return string(reslut)
}

func ParseDateToTimeStamp(str string, timeLayout string) (int64, error) {
	loc, _ := time.LoadLocation("Local") //重要：获取时区

	if theTime, err := time.ParseInLocation(timeLayout, str, loc); err != nil {
		return 0, err
	} else {
		return theTime.Unix(), nil
	}
}
