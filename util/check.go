package util

import "regexp"

func IsIP(text string) (matched bool, err error) {
	matched, err = regexp.MatchString("((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}", text)
	return
}
