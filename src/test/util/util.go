package testUtil

import (
	"fmt"
	"math/rand"
	"time"
)

func MakeRandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	sources := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := 0; i < n; i++ {
		randIndex := rand.Intn(len(sources))
		b[i] = sources[randIndex]
	}
	return string(b)
}

func MakeRandomTimeString() string {
	rand.Seed(time.Now().UnixNano())

	hour := rand.Intn(25)
	minute := rand.Intn(61)
	return fmt.Sprintf("%d:%d", hour, minute)
}
