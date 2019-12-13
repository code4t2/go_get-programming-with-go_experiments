package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var count = 0

	for count < 10 {
		var num = rand.Intn(10) + 1
		fmt.Println(num)

		count++
	}

	for count := 10; count > 0; count-- { // Different scope, due to renewed assignment
		fmt.Println(count)
	}

	fmt.Println(count)
}
