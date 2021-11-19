package i4004

import "fmt"

func (c *CPU) NOP() uint8 {
	// No Operation
	fmt.Println("NOP")
	return 1 // 1 cycle
}

func (c *CPU) JCN(condition uint8, address uint8) uint8 {
	// Jump Conditional
	fmt.Println("JCN", condition, address)

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

func (c *CPU) XCH(register uint8) uint8 {
	// Exchange
	fmt.Println("XCH", register)
	c.Registers[register] = c.Accumulator
	return 1 // 1 cycle
}

func (c *CPU) LDM(data uint8) uint8 {
	// Load Immediate
	fmt.Println("LDM", data)
	c.Accumulator = data
	return 1 // 1 cycle
}

func (c *CPU) INC(register uint8) uint8 {
	/* Mnemonic:	INC (Increment index register)
	     OPR OPA:	0110 RRRR
	     Symbolic:	(RRRR) +1 --> RRRR
	     Description:	The 4 bit content of the designated index register is
			 incremented by 1. The index register is set to zero in case of overflow.
			 The carry/link is unaffected. */

	fmt.Println("INC", register)
	c.Registers[register]++
	if c.Registers[register]&0xf0 == 1 {
		c.Registers[register] = 0
	}
	return 1 // 1 cycle
}

func (c *CPU) ISZ(register uint8, address uint8) uint8 {
	// Jump Conditional
	fmt.Println("ISZ", register, address)
	c.Registers[register] = (c.Registers[register] + 1) & 0xF
	if c.Registers[register] > 0 {
		c.PC = (c.PC & 0xF0) | address
	} else {
		// don't do anything
	}
	return 2 // 2 cycles
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
