package main

import "fmt"

func main() {
	const lightSpeed = 299792 //km s^-1
	var distance = 56000000   //km
	fmt.Println(distance/lightSpeed, "seconds @ nearest")
	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds @ furthest")
	const hoursPerEarthDay = 24
	var (
		speed       = 100800   // km h^-1
		avgDistance = 96300000 // km
	)
	fmt.Println(avgDistance/speed/hoursPerEarthDay, "days")
}
