package main

import (
	"fmt"
	"os"

	"github.com/cgoerner/i4004"
)

func main() {
	c := i4004.New()
	c.Debug = true

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
		fmt.Println(len)
		panic(err)
	}

	f.Close()

	for index, element := range bytes {
		c.PROM[index] = element
	}

	c.Run()
}
