package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Making big.ints
	// 1. With NewInt
	secondsPerDay1 := big.NewInt(86400)
	fmt.Println("Number of seconds in a day", secondsPerDay1)
	// using 'new' and SetString()
	secondsPerDay2 := new(big.Int)
	secondsPerDay2.SetString("86400", 10)
}
