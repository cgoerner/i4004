package i4004

import (
	"fmt"
	"strconv"
)

func PrintAll(n uint8) {
	var number int64 = int64(n)
	fmt.Print(strconv.FormatInt(number, 2), " ")
	fmt.Print(strconv.FormatInt(number, 10), " ")
	fmt.Println(strconv.FormatInt(number, 16))
}
