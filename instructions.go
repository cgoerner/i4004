package i4004

import "fmt"

func (c *CPU) NOP() uint8 {
	// No Operation
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) JCN(condition uint8, address uint8) uint8 {
	invert := 0
	if condition&0x8 > 0 {
		invert = 1
	}
	temp := c.PCStack[0] & 0xf00
	temp |= uint16(address)
	if condition&0x4 > 0 {
		if (^int(c.Accumulator))^invert > 0 { // if accumulator is 0
			c.PCStack[0] = temp
		} else {
			c.IncrementPC()
		}
	} else if condition&0x2 > 0 {
		if (int(c.Carry))^invert > 0 { // if carry
			c.PCStack[0] = temp
		} else {
			c.IncrementPC()
		}
	} else if condition&0x1 > 0 {
		if (int(c.Test))^invert > 0 { // if test
			c.PCStack[0] = temp
		} else {
			c.IncrementPC()
		}
	} else {
		c.IncrementPC()
	}

	return 2 // 2 cycles
}

func (c *CPU) FIM(rpair uint8, data uint8) uint8 {
	// Fetch Immediate
	c.SetRegisterPair(rpair, data)
	c.IncrementPC()
	return 2 // 2 cycles
}

func (c *CPU) SRC(rpair uint8) uint8 {
	// Send Register Control
	c.RAMAddressRegister = c.GetRegisterPair(rpair)
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) FIN(rpair uint8) uint8 {
	// Fetch Indirect
	c.SetRegisterPair(rpair, c.PROM[uint8(c.PCStack[0]&0xF00)|c.GetRegisterPair(0)])
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) JIN(rpair uint8) uint8 {
	// Jump Indirect
	fmt.Println("NOT YET IMPLEMENTED!")
	return 1 // 1 cycle
}

func (c *CPU) JUN(addr1 uint16, addr2 uint8) uint8 {
	// Jump Unconditional
	c.PCStack[0] = addr1 | uint16(addr2)
	return 2 // 1 cycle
}

func (c *CPU) JMS(addr1 uint16, addr2 uint8) uint8 {
	// JMP: Jump to Subroutine
	address := addr1 | uint16(addr2)
	if c.StackPointer < 3 {
		c.StackPointer++
		for i := c.StackPointer; i > 0; i-- {
			c.PCStack[i] = c.PCStack[i-1]
		}
		c.PCStack[0] = uint16(address)
	} else {
		c.IncrementPC()
		if c.Debug {
			fmt.Println("Stack Overflow!")
		}
	}

	return 2 // 1 cycle
}

func (c *CPU) INC(register uint8) uint8 {
	// Increment index register
	c.Registers[register]++
	if c.Registers[register]&0xf0 == 1 {
		c.Registers[register] = 0
	}
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) ISZ(register uint8, address uint8) uint8 {
	// Increment and Skip
	c.Registers[register] = (c.Registers[register] + 1) & 0xF
	if c.Registers[register] > 0 {
		c.PCStack[0] = ((c.PCStack[0]) & 0xF00) | uint16(address)
	} else {
		c.IncrementPC()
	}
	return 2 // 2 cycles
}

func (c *CPU) ADD(register uint8) uint8 {
	// ADD: Add. Add index register to accumulator with carry.
	c.Accumulator += uint8(c.Registers[register]) + c.Carry
	c.Carry = 0
	if c.Accumulator&0xF0 == 1 {
		c.Accumulator = c.Accumulator & 0xF
		c.Carry = 1
	}
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) SUB(register uint8) uint8 {
	// SUB: Subtract. Subtract index register from accumulator with borrow.
	c.Accumulator += (^c.Registers[register] & 0xf) + (c.Carry & 1)
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) LD(register uint8) uint8 {
	// LD: Load. Load index register to Accumulator.
	c.Accumulator = c.Registers[register]
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) XCH(register uint8) uint8 {
	// XCH: Exchange. Exchange index register and accumulator.
	temp := c.Registers[register]
	c.Registers[register] = c.Accumulator
	c.Accumulator = temp
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) BBL(data uint8) uint8 {
	// Branch Back and Load
	if c.StackPointer > 0 {
		for i := uint8(0); i < c.StackPointer; i++ {
			c.PCStack[i] = c.PCStack[i+1]
		}
		c.PCStack[c.StackPointer] = 0
		c.StackPointer--
		c.Accumulator = data
	} else {
		if c.Debug {
			fmt.Println("Stack error")
		}
	}
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) LDM(data uint8) uint8 {
	// Load Immediate. Load data to Accumulator)
	c.Accumulator = data
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) WRM() uint8 {
	// Write Main Memory. Write accumulator into RAM character.
	c.RAMData[c.RAMAddressRegister] = c.Accumulator
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) WMP() uint8 {
	// Write RAM Port
	c.ActiveBank = c.Accumulator
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) WRR() uint8 {
	// Write ROM Port
	c.ROMPort = c.Accumulator
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) WR(n uint8) uint8 {
	// Write Status Char n
	fmt.Println("WRn")
	c.RAMStatus[c.ActiveBank][0][n] = c.Accumulator
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) SBM() uint8 {
	// Subtract Main Memory
	fmt.Println("NOT YET IMPLEMENTED!")
	return 1 // 1 cycle
}

func (c *CPU) RDM() uint8 {
	// Read Main Memory
	fmt.Println("NOT YET IMPLEMENTED!")
	return 1 // 1 cycle
}

func (c *CPU) RDR() uint8 {
	// Read ROM Port
	fmt.Println("RDR")
	c.Accumulator = c.ROMPort
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) ADM() uint8 {
	// Read ROM Port
	fmt.Println("NOT YET IMPLEMENTED!")
	return 1 // 1 cycle
}

func (c *CPU) RD(n uint8) uint8 {
	// Read Status Char n
	fmt.Println("NOT YET IMPLEMENTED!")
	return 1 // 1 cycle
}

func (c *CPU) CLB() uint8 {
	// Clear Both
	c.Accumulator = 0x0
	c.Carry = 0x0
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) CLC() uint8 {
	// Clear Carry
	c.Carry = 0
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) IAC() uint8 {
	// Increment Accumulator
	c.Accumulator++
	c.Carry = 0
	if c.Accumulator == 0x10 {
		c.Accumulator = 0x0
		c.Carry = 1
	}
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) CMC() uint8 {
	// Complement Carry
	fmt.Println("NOT YET IMPLEMENTED!")
	return 1 // 1 cycle
}

func (c *CPU) CMA() uint8 {
	// Complement
	fmt.Println("NOT YET IMPLEMENTED!")
	return 1 // 1 cycle
}

func (c *CPU) RAL() uint8 {
	// Rotate Left
	c.Accumulator = (c.Accumulator << 1) | c.Carry
	c.Carry = 0
	if (c.Accumulator & 0xF0) != 0 {
		c.Accumulator = c.Accumulator & 0xF
		c.Carry = 1
	}
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) RAR() uint8 {
	// Rotate Right
	temp := c.Accumulator & 1
	c.Accumulator = (c.Accumulator >> 1) | (c.Carry << 3)
	c.Carry = temp
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) TCC() uint8 {
	// Transfer Carry and Clear
	fmt.Println("NOT YET IMPLEMENTED!")
	return 1 // 1 cycle
}

func (c *CPU) DAC() uint8 {
	// Decrement Accumulator
	fmt.Println("NOT YET IMPLEMENTED!")
	return 1 // 1 cycle
}

func (c *CPU) TCS() uint8 {
	// Transfer Carry Clear
	fmt.Println("NOT YET IMPLEMENTED!")
	return 1 // 1 cycle
}

func (c *CPU) STC() uint8 {
	// Set Carry
	fmt.Println("STC")
	c.Carry = 1
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) DAA() uint8 {
	// Decimal Adjust Accumulator
	fmt.Println("NOT YET IMPLEMENTED!")
	return 1 // 1 cycle
}

func (c *CPU) KBP() uint8 {
	// Keyboard Process
	fmt.Println("NOT YET IMPLEMENTED!")
	return 1 // 1 cycle
}

func (c *CPU) DCL() uint8 {
	// Designate Command Line
	fmt.Println("DCL")
	switch val := c.Accumulator & 0x7; val {
	case 0:
		c.ActiveBank = 1
	case 1:
		c.ActiveBank = 2
	case 2:
		c.ActiveBank = 4
	case 3:
		c.ActiveBank = 3
	case 4:
		c.ActiveBank = 8
	case 5:
		c.ActiveBank = 10
	case 6:
		c.ActiveBank = 12
	case 7:
		c.ActiveBank = 14
	}

	c.IncrementPC()
	return 1 // 1 cycle
}
