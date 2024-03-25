package main

import (
	"math/rand"
)

func random(min, max int) int {
	if min == max {
		return min
	}
	r := rand.New(rand.NewSource(rand.Int63()))
	return r.Intn(max-min+1) + min
}
