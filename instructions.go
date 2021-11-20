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

func (c *CPU) FIN() uint8 {
	// Fetch Indirect
	fmt.Println("FIN")
	return 1 // 1 cycle
}

func (c *CPU) JIN() uint8 {
	// Jump Indirect
	fmt.Println("JIN")
	return 1 // 1 cycle
}

func (c *CPU) JUN() uint8 {
	// Jump Unconditional
	fmt.Println("JUN")
	return 2 // 1 cycle
}

func (c *CPU) JMS() uint8 {
	// Jump to Subroutine
	fmt.Println("JMS")
	return 2 // 1 cycle
}

func (c *CPU) INC(register uint8) uint8 {
	// Increment index register
	fmt.Println("INC", register)
	c.Registers[register]++
	if c.Registers[register]&0xf0 == 1 {
		c.Registers[register] = 0
	}
	return 1 // 1 cycle
}

func (c *CPU) ISZ(register uint8, address uint8) uint8 {
	// Increment and Skip
	fmt.Println("ISZ", register, address)
	c.Registers[register] = (c.Registers[register] + 1) & 0xF
	if c.Registers[register] > 0 {
		c.PC = (c.PC & 0xF0) | address
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

func (c *CPU) SUB() uint8 {
	// Subtract
	fmt.Println("SUB")
	return 1 // 1 cycle
}

func (c *CPU) LD() uint8 {
	// Load
	fmt.Println("LD")
	return 1 // 1 cycle
}

func (c *CPU) XCH(register uint8) uint8 {
	// Exchange
	fmt.Println("XCH", register)
	c.Registers[register] = c.Accumulator
	return 1 // 1 cycle
}

func (c *CPU) BBL() uint8 {
	// Branch Back and Load
	fmt.Println("BBL")
	return 1 // 1 cycle
}

func (c *CPU) LDM(data uint8) uint8 {
	// Load Immediate. Load data to Accumulator)
	fmt.Println("LDM", data)
	c.Accumulator = data
	return 1 // 1 cycle
}

func (c *CPU) WRM() uint8 {
	// Write Main Memory. Write accumulator into RAM character.
	fmt.Println("WRM")
	c.Accumulator = c.RAMAddressRegister //???
	return 1                             // 1 cycle
}

func (c *CPU) WMP() uint8 {
	// Write RAM Port
	fmt.Println("WMP")
	return 1 // 1 cycle
}

func (c *CPU) WRR() uint8 {
	// Write ROM Port
	fmt.Println("WRR")
	return 1 // 1 cycle
}

func (c *CPU) WR0() uint8 {
	// Write Status Char 0
	fmt.Println("WR0")
	return 1 // 1 cycle
}

func (c *CPU) WR1() uint8 {
	// Write Status Char 1
	fmt.Println("WR1")
	return 1 // 1 cycle
}

func (c *CPU) WR2() uint8 {
	// Write Status Char 2
	fmt.Println("WR2")
	return 1 // 1 cycle
}

func (c *CPU) WR3() uint8 {
	// Write Status Char 0
	fmt.Println("WR3")
	return 1 // 1 cycle
}

func (c *CPU) SBM() uint8 {
	// Subtract Main Memory
	fmt.Println("SBM")
	return 1 // 1 cycle
}

func (c *CPU) RDM() uint8 {
	// Read Main Memory
	fmt.Println("RDM")
	return 1 // 1 cycle
}

func (c *CPU) RDR() uint8 {
	// Read ROM Port
	fmt.Println("RDR")
	return 1 // 1 cycle
}

func (c *CPU) RD0() uint8 {
	// Read Status Char 0
	fmt.Println("RD0")
	return 1 // 1 cycle
}

func (c *CPU) RD1() uint8 {
	// Read Status Char 1
	fmt.Println("RD1")
	return 1 // 1 cycle
}

func (c *CPU) RD2() uint8 {
	// Read Status Char 2
	fmt.Println("RD2")
	return 1 // 1 cycle
}

func (c *CPU) RD3() uint8 {
	// Read Status Char 3
	fmt.Println("RD3")
	return 1 // 1 cycle
}

func (c *CPU) CLB() uint8 {
	// Clear Both
	fmt.Println("CLB")
	return 1 // 1 cycle
}

func (c *CPU) CLC() uint8 {
	// Clear Carry
	fmt.Println("CLC")
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

func (c *CPU) CMC() uint8 {
	// Complement Carry
	fmt.Println("CMC")
	return 1 // 1 cycle
}

func (c *CPU) CMA() uint8 {
	// Complement
	fmt.Println("CMA")
	return 1 // 1 cycle
}

func (c *CPU) RAL() uint8 {
	// Rotate Left
	fmt.Println("RAL")
	return 1 // 1 cycle
}

func (c *CPU) RAR() uint8 {
	// Rotate Right
	fmt.Println("RAR")
	return 1 // 1 cycle
}

func (c *CPU) TCC() uint8 {
	// Transfer Carry and Clear
	fmt.Println("TCC")
	return 1 // 1 cycle
}

func (c *CPU) DAC() uint8 {
	// Decrement Accumulator
	fmt.Println("DAC")
	return 1 // 1 cycle
}

func (c *CPU) TCS() uint8 {
	// Transfer Carry Clear
	fmt.Println("TCS")
	return 1 // 1 cycle
}

func (c *CPU) STC() uint8 {
	// Set Carry
	fmt.Println("STC")
	return 1 // 1 cycle
}

func (c *CPU) DAA() uint8 {
	// Decimal Adjust Accumulator
	fmt.Println("DAA")
	return 1 // 1 cycle
}

func (c *CPU) KBP() uint8 {
	// Keyboard Process
	fmt.Println("KBP")
	return 1 // 1 cycle
}

func (c *CPU) DCL() uint8 {
	// Designate Command Line
	fmt.Println("DCL")
	return 1 // 1 cycle
}
