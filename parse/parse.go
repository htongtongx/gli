package parse

import (
	"fmt"
	"log"
	"strconv"
)

func ToString(v interface{}) string {
	var result = ""
	switch v.(type) {
	case int:
		result = strconv.Itoa(v.(int))
	case int32:
		result = strconv.FormatInt(int64(v.(int32)), 10)
	case int64:
		result = strconv.FormatInt(v.(int64), 10)
	case float32:
		result = strconv.FormatFloat(float64(v.(float32)), 'f', 2, 32)
	case float64:
		result = strconv.FormatFloat(v.(float64), 'f', 2, 64)
	}
	return result
}

func ToInt(s string, defaultValue int) int {
	if s == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println(err.Error())
		return defaultValue
	}
	return i
}

func ToInt64(s string, defaultValue int64) int64 {
	if s == "" {
		return defaultValue
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return defaultValue
	}
	return i
}

//float64保留两位小数,四舍五入
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
