package i4004

import "fmt"

func (c *CPU) NOP() uint8 {
	// No Operation
	return 1 // 1 cycle
}

func (c *CPU) JCN(condition uint8, address uint8) uint8 {
	// Jump Conditional
	return 2 // 2 cycles
}

func (c *CPU) FIM(rpair uint8, data uint8) uint8 {
	// Fetch Immediate
	c.SetRegisterPair(rpair, data)
	return 2 // 2 cycles
}

func (c *CPU) SRC(rpair uint8) uint8 {
	// Send Register Control
	c.RAMAddressRegister = c.GetRegisterPair(rpair)
	return 1 // 1 cycle
}

func (c *CPU) FIN() uint8 {
	// Fetch Indirect
	return 1 // 1 cycle
}

func (c *CPU) JIN() uint8 {
	// Jump Indirect
	return 1 // 1 cycle
}

func (c *CPU) JUN() uint8 {
	// Jump Unconditional
	return 2 // 1 cycle
}

func (c *CPU) JMS(addr1 uint8, addr2 uint8) uint8 {
	// JMP: Jump to Subroutine
	address := addr1 | addr2
	if c.StackPointer < 3 {
		c.StackPointer++
		for i := c.StackPointer; i > 0; i-- {
			if i == 1 {
				c.PCStack[i] = c.PCStack[i-1] - 1 //c.PCStack[0] as already been incremented, so we need to use one less
			} else {
				c.PCStack[i] = c.PCStack[i-1]
			}
		}
		c.PCStack[0] = address
	} else {
		panic("Stack Overflow!")
	}
	return 2 // 1 cycle
}

func (c *CPU) INC(register uint8) uint8 {
	// Increment index register
	c.Registers[register]++
	if c.Registers[register]&0xf0 == 1 {
		c.Registers[register] = 0
	}
	return 1 // 1 cycle
}

func (c *CPU) ISZ(register uint8, address uint8) uint8 {
	// Increment and Skip
	c.Registers[register] = (c.Registers[register] + 1) & 0xF
	if c.Registers[register] > 0 {
		c.PCStack[0] = (c.PCStack[0] & 0xF0) | address
	}
	return 2 // 2 cycles
}

func (c *CPU) ADD(register uint8) uint8 {
	// Add
	c.Accumulator += uint8(c.Registers[register]) + c.Carry
	c.Carry = 0
	if c.Accumulator&0xF0 == 1 {
		c.Accumulator = c.Accumulator & 0xF
		c.Carry = 1
	}
	return 1 // 1 cycle
}

func (c *CPU) SUB() uint8 {
	// Subtract
	return 1 // 1 cycle
}

func (c *CPU) LD(register uint8) uint8 {
	// LD: Load. Load index register to Accumulator.
	c.Accumulator = c.Registers[register]
	return 1 // 1 cycle
}

func (c *CPU) XCH(register uint8) uint8 {
	// XCH: Exchange. Exchange index register and accumulator.
	temp := c.Registers[register]
	c.Registers[register] = c.Accumulator
	c.Accumulator = temp
	return 1 // 1 cycle
}

func (c *CPU) BBL(data uint8) uint8 {
	// Branch Back and Load
	if c.StackPointer > 0 {
		for i := uint8(0); i < c.StackPointer; i++ {
			fmt.Println(i)
			c.PCStack[i] = c.PCStack[i+1]
		}
		c.PCStack[c.StackPointer] = 0
		c.StackPointer--
		c.Accumulator = data
	}
	return 1 // 1 cycle
}

func (c *CPU) LDM(data uint8) uint8 {
	// Load Immediate. Load data to Accumulator)
	c.Accumulator = data
	return 1 // 1 cycle
}

func (c *CPU) WRM() uint8 {
	// Write Main Memory. Write accumulator into RAM character.
	c.Accumulator = c.RAMAddressRegister //???
	return 1                             // 1 cycle
}

func (c *CPU) WMP() uint8 {
	// Write RAM Port
	c.RAMData[0] = c.Accumulator
	return 1 // 1 cycle
}

func (c *CPU) WRR() uint8 {
	// Write ROM Port
	c.ROMPort = c.Accumulator
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
