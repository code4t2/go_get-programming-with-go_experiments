package main

import "fmt"

func main() {
	message := "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println("")
	messageInSpanish := "Hola Estación Espacial Internacional"

	for _, r := range messageInSpanish {
		if r >= 'a' && r <= 'z' {
			r = r + 13
			if r > 'z' {
				r = r - 26
			}
		}
		if r >= 'A' && r <= 'Z' {
			r = r + 13
			if r > 'Z' {
				r = r - 26
			}
		}
		fmt.Printf("%c", r)
	}

}
