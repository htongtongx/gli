package timer

import "github.com/htongtongx/gli/parse"

const (

	//定义每分钟的秒数
	SecondsPerMinute = 60
	//定义每小时的秒数
	SecondsPerHour = SecondsPerMinute * 60
	//定义每天的秒数
	SecondsPerDay = SecondsPerHour * 24
)

func ResolveTime(seconds int) (day, hour, minute, second int) {
	second = seconds % 60
	//每分钟秒数
	minute = seconds / SecondsPerMinute
	if minute >= 60 {
		minute = minute - (seconds/SecondsPerHour)*60
	}
	//每小时秒数
	hour = seconds / SecondsPerHour
	if hour >= 24 {
		hour = hour - (seconds/SecondsPerDay)*24
	}
	//每天秒数
	day = seconds / SecondsPerDay
	return
}

//将秒数转成x天x时x分x秒
func GetResolveTimeStr(seconds int) string {
	result := ""
	day, hour, minute, second := ResolveTime(seconds)
	if day > 0 {
		result += parse.ToString(day) + "天"
	}
	if hour > 0 {
		result += parse.ToString(hour) + "时"
	}
	if minute > 0 {
		result += parse.ToString(minute) + "分"
	}
	if second > 0 {
		result += parse.ToString(second) + "秒"
	}
	return result
}
