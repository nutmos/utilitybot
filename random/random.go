package random

import (
	"log"
	"math/rand"
)

func RandomNumber(n int) int {
	log.Printf("%d\n", n)
	return rand.Intn(n)
}
