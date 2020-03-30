package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println(rand.Intn(5))
	}
}
