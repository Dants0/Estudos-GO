package mymathpackage

import (
	"fmt"
	"math/rand"
)

func RandomNumber() int{
	random := rand.Intn(25)
	fmt.Println(random)
	return int(random)
}