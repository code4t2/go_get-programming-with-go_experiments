package main

import (
	"fmt"
	"math"
)

func main() {
	var days1 float64 = 365.2425
	var days2 = 365.2425
	days3 := 365.2425
	fmt.Println(days1, days2, days3)

	// Pi
	const pi64 = math.Pi
	const pi32 float32 = math.Pi

	fmt.Println("Pi64", pi64, "Pi32\n", pi32)
}
