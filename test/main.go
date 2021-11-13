package main

import (
	"fmt"

	"github.com/cgoerner/i4004"
)

func main() {
	c := i4004.New()
	fmt.Println("CPU Initialised")
	fmt.Println("Manufacturer: " + c.Manufacturer)
	fmt.Println("Model: " + c.Model)
	fmt.Println("Speed: " + c.Speed)
	fmt.Print("Tick length: ")
	fmt.Println(c.ClockTime)
}
