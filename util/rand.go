package util

import (
	"encoding/base64"
	"math/rand"
	"time"
)

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func RandInt(len int) int {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len)
	return i
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GenerateSubId(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(letterRunes))]
	}
	return string(b)
}

var nums = []rune("0123456789")

func RandNumByCount(count int) string {
	b := make([]rune, count)
	for i := range b {
		b[i] = nums[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(nums))]
	}
	return string(b)
}
