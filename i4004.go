package i4004

import (
	"fmt"
	"os"
	"time"
)

const (
	SpeedDebug      = 10 * time.Millisecond
	Speed741kHz     = 1350 * time.Nanosecond
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
	ROMPort   uint8
	RAMData   [4096]uint8 //4KB
	RAMStatus [16][4][4]uint8
	Registers [16]uint8

	Accumulator        uint8    //Accumulator (4 bits)
	Carry              uint8    //Carry flag (1 bit)
	Test               uint8    //Test pin (1 bit)
	PCStack            [4]uint8 //Program counter stack
	StackPointer       uint8
	RAMAddressRegister uint8 //???

	ClockTime    time.Duration //Time between each clock tick
	TimeUsed     time.Duration //Time spent processing operations
	CycleCounter uint64

	Debug bool
}

func New() (c *CPU) {
	c = &CPU{
		Manufacturer: CPUManufacturer,
		Model:        CPUModel,
		Speed:        CPUSpeed,
		ClockTime:    Speed741kHz,
	}
	return
}

func (c *CPU) Run() {

	ticker := time.NewTicker(c.ClockTime)

	defer func() {
		f := recover()
		if f != nil && f != "hlt" {
			panic(f)
		}
	}()
	defer ticker.Stop()

	for {
		<-ticker.C // wait for next tick

		c.Step()

		if c.CycleCounter >= 100 {
			os.Exit(2)
		}
	}
}

func (c *CPU) Step() {
	opcode := c.FetchOpCode()
	cycles := c.PerformOp(opcode)

	c.CycleCounter += uint64(cycles)
	c.TimeUsed = time.Duration(c.CycleCounter) * c.ClockTime

	if c.Debug {
		fmt.Printf("OP=%02x %s cycles=%d timeUsed=%s \r\n", opcode, c.DebugInfo(), c.CycleCounter, c.TimeUsed)
	}
}

func (c *CPU) FetchOpCode() uint8 {
	opcode := c.PROM[c.PCStack[0]]
	c.PCStack[0]++
	return opcode
}

func (c *CPU) PerformOp(opcode uint8) uint8 {

	operand := opcode % 16

	//fmt.Println("opcode and operand")
	//PrintAll(opcode)
	//PrintAll(operand)

	if opcode >= 0x0 && opcode <= 0x09 {
		return c.NOP()
	} else if opcode >= 0x10 && opcode <= 0x1F {
		return c.JCN(opcode, operand)
	} else if opcode >= 0x20 && opcode <= 0x2F {
		if !(operand%2 == 1) {
			nextcode := c.FetchOpCode()
			return c.FIM(operand, nextcode)
		} else {
			return c.SRC((operand - 1) / 2)
		}
	} else if opcode >= 0x30 && opcode <= 0x3F {
		if !(operand%2 == 1) {
			nextcode := c.FetchOpCode()
			return c.FIN(operand, nextcode)
		} else {
			return c.JIN((operand - 1) / 2)
		}
	} else if opcode >= 0x40 && opcode <= 0x4F {
		return c.JUN(operand & 0xF0)
	} else if opcode >= 0x50 && opcode <= 0x5F {
		nextcode := c.FetchOpCode()
		return c.JMS(operand, nextcode)
	} else if opcode >= 0x60 && opcode <= 0x6F {
		return c.INC(operand)
	} else if opcode >= 0x70 && opcode <= 0x7F {
		nextcode := c.FetchOpCode()
		return c.ISZ(operand, nextcode)
	} else if opcode >= 0x80 && opcode <= 0x8F {
		return c.ADD(operand)
	} else if opcode >= 0x90 && opcode <= 0x9F {
		return c.SUB(operand)
	} else if opcode >= 0xA0 && opcode <= 0xAF {
		return c.LD(operand)
	} else if opcode >= 0xB0 && opcode <= 0xBF {
		return c.XCH(operand)
	} else if opcode >= 0xC0 && opcode <= 0xCF {
		return c.BBL(operand)
	} else if opcode >= 0xD0 && opcode <= 0xDF {
		return c.LDM(operand)
	} else if opcode == 0xE0 {
		return c.WRM()
	} else if opcode == 0xE1 {
		return c.WMP()
	} else if opcode == 0xE2 {
		return c.WRR()
	} else if opcode == 0xE4 {
		return c.WR(0)
	} else if opcode == 0xE5 {
		return c.WR(1)
	} else if opcode == 0xE6 {
		return c.WR(2)
	} else if opcode == 0xE7 {
		return c.WR(3)
	} else if opcode == 0xE8 {
		return c.SBM()
	} else if opcode == 0xE9 {
		return c.RDM()
	} else if opcode == 0xEA {
		return c.RDR()
	} else if opcode == 0xEB {
		return c.ADM()
	} else if opcode == 0xEC {
		return c.RD(0)
	} else if opcode == 0xED {
		return c.RD(1)
	} else if opcode == 0xEE {
		return c.RD(2)
	} else if opcode == 0xEF {
		return c.RD(3)
	} else if opcode == 0xF0 {
		return c.CLB()
	} else if opcode == 0xF1 {
		return c.CLC()
	} else if opcode == 0xF2 {
		return c.IAC()
	} else if opcode == 0xF5 {
		return c.RAL()
	} else if opcode == 0xF6 {
		return c.RAR()
	} else if opcode == 0xF7 {
		return c.TCC()
	} else if opcode == 0xF8 {
		return c.DAC()
	} else if opcode == 0xF9 {
		return c.TCS()
	} else if opcode == 0xFA {
		return c.STC()
	} else if opcode == 0xFB {
		return c.DAA()
	} else if opcode == 0xFD {
		return c.DCL()
	} else {
		fmt.Println("Unknown operation!")
		return 1
	}

}

func (c *CPU) SetRegisterPair(index uint8, data uint8) {
	c.Registers[(index)+1] = uint8(data & 0xF)
	c.Registers[(index)] = uint8(data >> 0x4)
}

func (c *CPU) GetRegisterPair(index uint8) uint8 {
	return ((uint8(c.Registers[index])<<4)&0xF0 | uint8(c.Registers[index+1])&0xF)
}

func (c *CPU) LoadFileIntoROM(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}

	bytes := make([]byte, fi.Size())

	len, err := f.Read(bytes)
	if err != nil {
		fmt.Println(len)
		panic(err)
	}

	f.Close()

	for index, element := range bytes {
		c.PROM[index] = element
	}
}
