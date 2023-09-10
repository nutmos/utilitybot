package random

import "math/rand"

func randomNumber(n int) int {
	return rand.Intn(n)
}
