package main

import (
	"fmt"
	"os"

	cpu "github.com/cgoerner/i4004/pkg/cpu"
)

func main() {
	// Create a new CPU
	c := cpu.New()
	c.Debug = true

	fmt.Println(c.CPUInfo())

	// Populate PROM with instructions
	c.LoadFileIntoROM(os.Args[1])

	// Start processing instructions
	c.Run()
}
