package main

import "fmt"

func main() {
	ageInt := 41
	marsAgeFloat64 := float64(ageInt)

	marsDaysInt := 687
	marsDaysFloat64 := float64(marsDaysInt)
	earthDaysFloat64 := 365.2425
	marsAgeFloat64 = marsAgeFloat64 * earthDaysFloat64 / marsDaysFloat64
	fmt.Println("I am", marsAgeFloat64, "years old on Mars.")
	fmt.Printf("I am %5.2f years old on Mars.", marsAgeFloat64)
}
