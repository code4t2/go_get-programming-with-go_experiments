package main

import (
	"fmt"
	"math"
)

func main() {
	third := 1.0 / 3.0
	fmt.Println(third)
	fmt.Println(third + third + third)
	// Float addition
	piggyBank := 0.1
	piggyBank = piggyBank + 0.2
	fmt.Println(piggyBank == 0.3) //Just like ECMAScript
	fmt.Println(math.Abs(piggyBank-0.3) < 0.0001)

	if 0.3 == 0.1+0.2 {
		fmt.Printf("Not like ECMAScript %v", 0.1+0.2)
	} else {
		fmt.Printf("Just like ECMAScript %v", 0.1+0.2)
	}

	bankVault := 0.0 + 0.1 + 0.1 + 0.1 + 0.1 + 0.1 + 0.1 + 0.1 + 0.1 + 0.1 + 0.1 + 0.1
	fmt.Println("\n", bankVault)

	bankVault = 0.0
	for i := 0; i < 11; i++ {
		bankVault += 0.1
	}
	fmt.Println(bankVault)
}
