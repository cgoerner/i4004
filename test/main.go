package main

import (
	"fmt"
	"os"

	"github.com/cgoerner/i4004"
)

func main() {
	c := i4004.New()

	fmt.Println(c.CPUInfo())

	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}

	bytes := make([]byte, fi.Size())

	len, err := f.Read(bytes)
	if err != nil {
		panic(err)
	}

	f.Close()

	fmt.Println(len)

	for index, element := range bytes {
		c.PrintAll(element)
		c.PROM[index] = element
	}

	c.Run()
}
