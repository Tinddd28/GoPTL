package random

import (
	//"crypto/rand"
	"encoding/hex"
	"time"

	"golang.org/x/exp/rand"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(uint64(time.Now().UnixNano()))
}

func RandomPass(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}

func RandomCountry() string {
	countries := []string{"Russia", "USA", "UK", "Canada", "Germany", "France", "Italy", "Spain", "Japan", "China", "Russia", "Australia", "Brazil", "Mexico", "Argentina", "India", "South Africa", "Nigeria", "Egypt", "Kenya", "Ghana", "Morocco", "Turkey", "Saudi Arabia", "UAE", "Qatar", "Kuwait", "Iraq", "Iran", "Pakistan", "Afghanistan", "Bangladesh", "Sri Lanka", "Nepal", "Bhutan", "Myanmar", "Thailand", "Vietnam", "Malaysia", "Indonesia", "Philippines", "Singapore", "South Korea", "North Korea", "New Zealand"}
	return countries[rand.Intn(len(countries))]
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
