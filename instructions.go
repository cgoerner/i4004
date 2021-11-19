package i4004

import "fmt"

func (c *CPU) NOP() uint8 {
	// No Operation
	fmt.Println("NOP")
	return 1 // 1 cycle
}

func (c *CPU) JCN(opcode uint8, condition uint8) uint8 {
	// Jump Conditional
	fmt.Println("JCN " + string(condition))

	return 2 // 2 cycles
}

func (c *CPU) FIM(rpair uint8) uint8 {
	// Fetch Immediate
	fmt.Println("FIM", rpair)
	return 2 // 2 cycles
}

func (c *CPU) SRC(rpair uint8) uint8 {
	// Send Register Control
	fmt.Println("SRC", rpair)
	c.RAMAddressRegister = c.GetRegisterPair(rpair)
	return 1 // 1 cycle
}

func (c *CPU) XCH(data uint8) uint8 {
	// Exchange
	fmt.Println("XCH", data)
	c.Registers[data] = c.Accumulator
	return 1 // 1 cycle
}

func (c *CPU) LDM(data uint8) uint8 {
	// Load Immediate
	fmt.Println("LDM", data)
	c.Accumulator = data
	return 1 // 1 cycle
}

func (c *CPU) ADD(register uint8) uint8 {
	// Add
	fmt.Println("LDM", register)
	c.Accumulator += uint8(c.Registers[register]) + c.Carry
	c.Carry = 0
	if c.Accumulator&0xF0 == 1 {
		c.Accumulator = c.Accumulator & 0xF
		c.Carry = 1
	}
	return 1 // 1 cycle
}

func (c *CPU) IAC() uint8 {
	// Increment Accumulator
	fmt.Println("IAC")
	c.Accumulator++
	c.Carry = 0
	if c.Accumulator&0xF0 == 1 {
		c.Accumulator = c.Accumulator & 0xf
		c.Carry = 1
	}
	return 1 // 1 cycle
}

func (c *CPU) WRM() uint8 {
	// Write Main Memory. Write accumulator into RAM character.
	fmt.Println("WRM")
	c.Accumulator = c.RAMAddressRegister //???
	return 1                             // 1 cycle
}
