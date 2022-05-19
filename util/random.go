package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomStirng generates a random string of length n
func RandomStirng(n int) string {
	var sb strings.Builder
	sb.Grow(7)
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// TestAccount
// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomStirng(7)
}

// RandomBalance generates a random amount of Money
func RandomBalance() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "GBP", "JPY"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// TestEntry
// RandomAmount generates a random amount of Money
func RandomAmount() int64 {
	return RandomInt(0, 1000)
}
