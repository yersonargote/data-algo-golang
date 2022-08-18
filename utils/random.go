package utils

import "math/rand"

var LIMIT int = 100

func Random() int {
	return rand.Intn(LIMIT)
}
