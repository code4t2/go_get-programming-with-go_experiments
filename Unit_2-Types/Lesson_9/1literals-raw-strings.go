package main

import "fmt"

func main() {
	peace1 := "peace"
	var peace2 = "peace"
	var peace3 string = "peace"
	var blank string
	fmt.Println(peace1, peace2, peace3, blank)
	fmt.Println("peace be upon you\nupon you be peace")
	fmt.Println(`strings can span multiple lines with the \n escape sequence`)
	fmt.Println(`
    	peace be upon you
    upon you be peace`)
	fmt.Printf("%v is a %[1]T\n", "literal string")
	fmt.Printf("%v is a %[1]T\n", `raw string literal`)
}
