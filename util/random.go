package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandonInt Generate random int64 based in max, min provided
func RandonInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomString Generate random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner Generate random owner
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney Generate random amount of money
func RandomMoney() int64 {
	return RandonInt(0, 1000)
}

// RandomCurrency Generate random currency
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD", "HKD", "PKR", "JPY", "KRW", "BRL"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
