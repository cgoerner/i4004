package i4004

import (
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
	R10
	R11
	R12
	R13
	R14
	R15
)

type CPU struct {
	Manufacturer   string
	Model          string
	Speed          string
	Memory         [4096]uint8
	IndexRegisters [16]uint8

	A   uint8 //Accumulator
	C   uint8 //Carry flag
	PC  uint8 //Program counter
	PC1 uint8 //Push-down address call stack level 1
	PC2 uint8 //Push-down address call stack level 2
	PC3 uint8 //Push-down address call stack level 3

	ClockTime time.Duration
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
