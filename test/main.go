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

	/*
	   000 20 00 FIM P0,$00
	   002 22 00 FIM P1,$00
	   004 DC    LDM 12
	   005 B2    XCH R2
	   006 21    SRC P0
	   007 E0    WRM
	   008 F2    IAC
	   009 71 06 ISZ R1,$06
	   00B 60    INC R0
	   00C 72 06 ISZ R2,$06
	   00E 20 00 FIM P0,$00
	   010 22 00 FIM P1,$00
	   012 DC    LDM 12
	   013 B2    XCH R2
	   014 21    SRC P0
	   015 E4    WR0
	   016 F2    IAC
	   017 E5    WR1
	   018 F2    IAC
	   019 E6    WR2
	   01A F2    IAC
	   01B E7    WR3
	   01C F2    IAC
	   01D 60    INC R0
	   01E 72 14 ISZ R2,$14
	   020 40 20 JUN $020
	*/

	c.PROM[0] = 0x20
	c.PROM[1] = 0x00
	c.PROM[2] = 0x22
	c.PROM[3] = 0x00
	c.PROM[4] = 0xDC
	c.PROM[5] = 0xB2
	c.PROM[6] = 0x21
	c.PROM[7] = 0xE0
	c.PROM[8] = 0xF2
	c.PROM[9] = 0x71
	c.PROM[10] = 0x06
	c.PROM[11] = 0x60
	c.PROM[12] = 0x72
	c.PROM[13] = 0x06
	c.PROM[14] = 0x20
	c.PROM[15] = 0x22
	c.PROM[16] = 0x00
	c.PROM[17] = 0xDC
	c.PROM[18] = 0xB2
	c.PROM[19] = 0x21
	c.PROM[20] = 0xE4
	c.PROM[21] = 0xF2
	c.PROM[22] = 0xE5
	c.PROM[23] = 0xF2
	c.PROM[24] = 0xE6
	c.PROM[25] = 0xF2
	c.PROM[26] = 0xE7
	c.PROM[27] = 0xF2
	c.PROM[28] = 0x60
	c.PROM[29] = 0x72
	c.PROM[30] = 0x14
	c.PROM[31] = 0x40
	c.PROM[32] = 0x20

	c.Run()
}
