package helper

import "math/rand"

func GenerateId() int {
	id := rand.Intn(10000)
	return id
}