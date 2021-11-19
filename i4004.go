package i4004

import (
	"fmt"
	"os"
	"time"
)

const (
	SpeedDebug      = 10 * time.Millisecond
	Speed740kHz     = 1350 * time.Nanosecond
	CPUManufacturer = "Intel"
	CPUModel        = "4004"
	CPUSpeed        = "741 kHz"
)

const (
	R0 uint8 = iota
	R1
	R2
	R3
	R4
	R5
	R6
	R7
	R8
	R9
	RA
	RB
	RC
	RD
	RE
	RF
)

type CPU struct {
	Manufacturer string
	Model        string
	Speed        string

	PROM      [4096]uint8 //4KB
	RAMData   [4096]uint8 //4KB
	RAMStatus [16][4][4]uint8
	Registers [16]uint8

	Accumulator        uint8 //Accumulator (4 bits)
	Carry              uint8 //Carry flag (1 bit)
	Test               uint8 //Test pin (1 bit)
	PC                 uint8 //Program counter
	PC1                uint8 //Push-down address call stack level 1
	PC2                uint8 //Push-down address call stack level 2
	PC3                uint8 //Push-down address call stack level 3
	RAMAddressRegister uint8 //???

	ClockTime    time.Duration
	CycleCounter uint64
}

func New() (c *CPU) {
	c = &CPU{
		Manufacturer: CPUManufacturer,
		Model:        CPUModel,
		Speed:        CPUSpeed,
		ClockTime:    Speed740kHz,
	}

	return
}

const tickBudget = 10 * time.Millisecond

func (c *CPU) Run() (cycles uint64) {

	fmt.Println(tickBudget)

	debug := true

	ticker := time.NewTicker(tickBudget)

	defer func() {
		f := recover()
		if f != nil && f != "hlt" {
			panic(f)
		}
	}()
	defer ticker.Stop()

	var TimeUsed time.Duration

	nops := 0

	for {
		<-ticker.C // wait for next tick

		for TimeUsed < tickBudget {

			op := c.FetchOpCode()
			cycles := c.PerformOp(op)

			c.CycleCounter += uint64(cycles)
			TimeUsed += time.Duration(cycles) * c.ClockTime

			if debug {
				fmt.Printf("OP=%02x %s cycles=%d timeUsed=%s \r\n", op, c.Debug(), c.CycleCounter, TimeUsed)
			}

			if op == 0x00 {
				nops++
			} else {
				nops = 0
			}

			if c.CycleCounter > 100 {
				os.Exit(2)
			}

		}

		TimeUsed -= tickBudget
	}

}

func (c *CPU) FetchOpCode() uint8 {
	op := c.PROM[c.PC]
	c.PC++

	return op
}

func (c *CPU) PerformOp(opcode uint8) uint8 {

	operand := opcode % 16

	//fmt.Println("opcode and operand")
	//PrintAll(opcode)
	//PrintAll(operand)

	if opcode == 0x00 {
		return c.NOP()
	} else if opcode >= 0x10 && opcode <= 0x1F {
		return c.JCN(opcode, operand)
	} else if opcode >= 0x20 && opcode <= 0x2F {
		if !(operand%2 == 1) {
			nextcode := c.FetchOpCode()
			return c.FIM(nextcode)
		} else {
			return c.SRC((operand - 1) / 2)
		}
	} else if opcode >= 0x60 && opcode <= 0x6F {
		return c.INC(operand)
	} else if opcode >= 0x70 && opcode <= 0x7F {
		nextcode := c.FetchOpCode()
		return c.ISZ(operand, nextcode)
	} else if opcode >= 0xB0 && opcode <= 0xBF {
		return c.XCH(operand)
	} else if opcode >= 0xD0 && opcode <= 0xDF {
		return c.LDM(operand)
	} else if opcode == 0xE0 {
		return c.WRM()
	} else if opcode == 0xF2 {
		return c.IAC()
	} else {
		fmt.Println("OOPS")
		return 1
	}

}

func (c *CPU) SetRegisterPair(index uint8, data uint8) {
	c.Registers[(index)+1] = uint8(data & 0xF)
	c.Registers[(index)] = uint8(data >> 0x4)
}

func (c *CPU) GetRegisterPair(index uint8) uint8 {
	fmt.Println()
	return ((uint8(c.Registers[index])<<4)&0xF0 | uint8(c.Registers[index+1])&0xF)
}
