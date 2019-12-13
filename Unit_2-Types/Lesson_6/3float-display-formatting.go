package main

import "fmt"

func main() {
	third := 1.0 / 3
	fmt.Println("Third\n", third)
	fmt.Printf("%v\n", third) //Prints 0.3333333333333333
	//The %f verb formats the value of third with a width and with precision
	fmt.Printf("%f\n", third) //Prints 0.333333
	fmt.Printf("%.3f\n", third) //Prints 0.333
	fmt.Printf("%4.2f\n", third) //Prints 0.33
	fmt.Printf("%05.2f\n", third) //Prints 00.33
	fmt.Printf("%06.2f\n", third) //Prints 000.33
}
