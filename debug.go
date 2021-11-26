package i4004

import (
	"fmt"
	"strconv"
)

func (c *CPU) DebugInfo() string {
	//mnem, args := c.disasmPC()

	//disasm := fmt.Sprintf("ASM %04x %02x => %-8s%-8s", c.PC, c.Memory[c.PC], mnem, args)

	regs := fmt.Sprintf("REG=%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x",
		c.Registers[R0], c.Registers[R1], c.Registers[R2], c.Registers[R3], c.Registers[R4], c.Registers[R5], c.Registers[R6], c.Registers[R7],
		c.Registers[R8], c.Registers[R9], c.Registers[RA], c.Registers[RB], c.Registers[RC], c.Registers[RD], c.Registers[RE], c.Registers[RF])

	ptrs := fmt.Sprintf("A=%01x C=%01x RA=%01x SP=%01x PC=%03x %03x %03x %03x",
		c.Accumulator, c.Carry, c.RAMAddressRegister, c.StackPointer, c.PCStack[0], c.PCStack[1], c.PCStack[2], c.PCStack[3])

	ram := ""
	for index := range c.RAMData[0:64] {
		ram += fmt.Sprintf("%01x", c.RAMData[index])
	}

	return fmt.Sprintf("%s %s\r\n RAM: %s", ptrs, regs, ram)

}

func (c *CPU) CPUInfo() string {
	return fmt.Sprintf("Manufacturer: %s\r\nModel: %s\r\nSpeed: %s\r\nTick length: %s", c.Manufacturer, c.Model, c.Speed, c.ClockTime)
}

func (c *CPU) PrintAll(n uint8) {
	var number int64 = int64(n)
	fmt.Print(strconv.FormatInt(number, 2), " ")
	fmt.Print(strconv.FormatInt(number, 10), " ")
	fmt.Println(strconv.FormatInt(number, 16))
}

func (c *CPU) PrintAll16(n uint16) {
	var number int64 = int64(n)
	fmt.Print(strconv.FormatInt(number, 2), " ")
	fmt.Print(strconv.FormatInt(number, 10), " ")
	fmt.Println(strconv.FormatInt(number, 16))
}
