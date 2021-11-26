package i4004

import "fmt"

func (c *CPU) NOP() uint8 {
	// No Operation
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) JCN(condition uint8, address uint8) uint8 {
	// Jump Conditional
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
	fmt.Println("FIN")
	c.PrintAll(rpair)
	c.SetRegisterPair(rpair, c.PROM[uint8(c.PCStack[0]&0xF00)|c.GetRegisterPair(0)])
	c.IncrementPC()
	return 1 // 1 cycle
}

func (c *CPU) JIN(rpair uint8) uint8 {
	// Jump Indirect
	return 1 // 1 cycle
}

func (c *CPU) JUN(address uint8) uint8 {
	// Jump Unconditional
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
		c.PCStack[0]-- // Hmm. Not sure about this. We didn't actually want to increment the PC before?
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
	return 1 // 1 cycle
}

func (c *CPU) SUB(register uint8) uint8 {
	// SUB: Subtract. Subtract index register from accumulator with borrow.
	c.Accumulator += (^c.Registers[register] & 0xf) + (c.Carry & 1)
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
	return 1 // 1 cycle
}

func (c *CPU) WRM() uint8 {
	// Write Main Memory. Write accumulator into RAM character.
	c.RAMData[c.RAMAddressRegister] = c.Accumulator
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
	return 1 // 1 cycle
}

func (c *CPU) SBM() uint8 {
	// Subtract Main Memory
	return 1 // 1 cycle
}

func (c *CPU) RDM() uint8 {
	// Read Main Memory
	return 1 // 1 cycle
}

func (c *CPU) RDR() uint8 {
	// Read ROM Port
	return 1 // 1 cycle
}

func (c *CPU) ADM() uint8 {
	// Read ROM Port
	return 1 // 1 cycle
}

func (c *CPU) RD(n uint8) uint8 {
	// Read Status Char n
	return 1 // 1 cycle
}

func (c *CPU) CLB() uint8 {
	// Clear Both
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
	if c.Accumulator&0xF0 == 1 {
		c.Accumulator = c.Accumulator & 0xf
		c.Carry = 1
	}
	return 1 // 1 cycle
}

func (c *CPU) CMC() uint8 {
	// Complement Carry
	return 1 // 1 cycle
}

func (c *CPU) CMA() uint8 {
	// Complement
	return 1 // 1 cycle
}

func (c *CPU) RAL() uint8 {
	// Rotate Left
	c.Accumulator = (c.Accumulator << 1) | c.Carry
	c.Carry = 0
	if (c.Accumulator & 0xF0) == 1 {
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
	return 1 // 1 cycle
}

func (c *CPU) TCC() uint8 {
	// Transfer Carry and Clear
	return 1 // 1 cycle
}

func (c *CPU) DAC() uint8 {
	// Decrement Accumulator
	return 1 // 1 cycle
}

func (c *CPU) TCS() uint8 {
	// Transfer Carry Clear
	return 1 // 1 cycle
}

func (c *CPU) STC() uint8 {
	// Set Carry
	return 1 // 1 cycle
}

func (c *CPU) DAA() uint8 {
	// Decimal Adjust Accumulator
	return 1 // 1 cycle
}

func (c *CPU) KBP() uint8 {
	// Keyboard Process
	return 1 // 1 cycle
}

func (c *CPU) DCL() uint8 {
	// Designate Command Line
	return 1 // 1 cycle
}
