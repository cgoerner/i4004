package main

import (
	"fmt"
	"os"

	"github.com/cgoerner/i4004"
)

func main() {
	// Create a new CPU object
	c := i4004.New()
	c.Debug = true

	fmt.Println(c.CPUInfo())

	c.LoadFileIntoROM(os.Args[1])

	c.Run()
}
