package utils

import (
	"math/rand"
	"time"
)

// RandomInt ...
func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(max-min) + min
	return random
}
