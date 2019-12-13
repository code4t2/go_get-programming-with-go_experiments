package main

import (
	"fmt"
	"math/big"
)

func main() {
	var distanceI64 int64 = 41.3e12 // The representation of the number isn't as important as the ability of the variable to hold such a large value
	fmt.Printf("The distance between the Sun and Alpha Centauri is %v kilometers", distanceI64)

	// Using big
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)

	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Println("Andromeda Galaxy is", distance, "km away.")

	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)

	fmt.Printf("That is %v days of travel at light speed. \n", days)

	// Using large literals directly (see readme)
	fmt.Println("Andromeda Galaxy is", 24000000000000000000/299792/86400, "light-speed days away.")

	// all as constants :-
	const distanceC = 24000000000000000000
	const lightSpeedC = 299792
	const secondsPerDayC = 86400

	const daysC = distanceC / lightSpeedC / secondsPerDayC

	fmt.Println("Andromeda Galaxy is", daysC, "light-speed days away. (using const this time)")
}
