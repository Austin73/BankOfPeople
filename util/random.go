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

// Randomly generate the random number between min and max

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		sb.WriteByte(alphabet[rand.Intn(k)])
	}
	return sb.String()
}

// RandomOwner name will generate random owner
func RandomOwner() string {
	return RandomString(6)
}

//  RandomMoney function will return random money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency will generate the random currency
func RandomCurrency() string {
	currencies := []string{"USD", "INR", "EUR"}
	return currencies[rand.Intn(len(currencies))]
}
