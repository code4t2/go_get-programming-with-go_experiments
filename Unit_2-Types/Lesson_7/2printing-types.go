package main

import "fmt"

func main() {
	year := 2018
	fmt.Printf("Type %T for %v \n", year, year) // %T for type and %v for (decimal?) value
	days := 365.2425
	fmt.Printf("Type %T for %[1]v \n", days) //use the first argument [1] for the second format verb
	var red, green, blue uint8 = 0, 0x8d, 0xd5
	fmt.Printf("The RGB values are %v %x %b \n", red, green, blue) //%x or %X prints hexadecimal, %b binary
	fmt.Println("Number formats can vary, here decimal 10 equal hex 10", 10 == 0x0A)
	var value uint8 = 3
	fmt.Printf("%08b\n", value)  // the preceding 08 left-pads with zeros upto 8 character width
	value++
	fmt.Printf("%8b\n", value)  // The preceding 8 left pads with spaces upto 8 character width
	value = value + 30 // To bring it to a suitable Unicode code point
	fmt.Printf("%c\n", value) // %c prints the Unicode character
}
