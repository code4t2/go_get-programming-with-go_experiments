package main

import "fmt"

func main() {
	fmt.Print("My weight on the surface of Mars is ") // ';' are optional
	fmt.Print(149.0 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Print(" years old.")
	fmt.Println("")
	// multiple strings, values or expressions
	fmt.Println("My weight on the surface of Mars is", 149.0*0.3783, "lbs, and I would be", 41*365.2425/687, "years old.")
	fmt.Println("")
	// Substituting values
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 149.0*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 41*365/687)
	fmt.Println("")
	// Substituting multiple values
	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 149.0)
	fmt.Println("")
	// Padding and alignment
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
	fmt.Println("")

}
