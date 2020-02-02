package main

import "fmt"

func main() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang) // Both byte and rune behave like the integer types they are aliases for
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang) // %c to display the unicode characters rather than their numeric values

	grade1 := 'A'
	var grade2 = 'A'
	var grade3 rune = 'A'
	fmt.Printf("%v %c %d\n", grade1, grade2, grade3)

	var star byte = '*'
	fmt.Printf("%c %[1]v\n", star)
	var smile rune = '😃'
	fmt.Printf("%c %[1]v\n", smile)
	acute := 'é'
	fmt.Printf("%c %[1]v\n", acute)

	peace := "shalom"
	peace = `salām`
	c := peace[5]
	fmt.Printf("\n %c \n \n", c) //%c -> UNICODE character

	for i := 0; i < len(peace); i++ {
		fmt.Printf("%c \n", peace[i])
	}
}
