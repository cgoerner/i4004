package main

import (
	"fmt"

	"github.com/cgoerner/i4004/pkg/cpu"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/html/v2"
)

var c = cpu.New()

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

func main() {

	engine := html.New("./web/views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", getRoot)
	app.Post("/load", postLoad)
	app.Post("/step", postStep)

	log.Fatal(app.Listen(":8080"))
}

func getRoot(ctx *fiber.Ctx) error {

	ram := ""
	for index := range c.RAMData[0:64] {
		ram += fmt.Sprintf("%01x", c.RAMData[index])
	}

	return ctx.Render("index", fiber.Map{})

}

func postLoad(ctx *fiber.Ctx) error {

	//c.Debug = true

	// Populate PROM with instructions
	c.LoadFileIntoROM("examples/fulltest.rom")

	ram := ""
	for index := range c.RAMData[0:64] {
		ram += fmt.Sprintf("%01x", c.RAMData[index])
	}

	return ctx.Render("cpuinfo", fiber.Map{
		"CPUInfo":            c.CPUInfo(),
		"Model":              c.Model,
		"Accumulator":        fmt.Sprintf("%01x", c.Accumulator),
		"Carry":              fmt.Sprintf("%01x", c.Carry),
		"RAMAddressRegister": fmt.Sprintf("%01x", c.RAMAddressRegister),
		"StackPointer":       fmt.Sprintf("%01x", c.StackPointer),
		"PCStack":            fmt.Sprintf("%01x", c.PCStack[0]),
		"Level1Stack":        fmt.Sprintf("%01x", c.PCStack[1]),
		"Level2Stack":        fmt.Sprintf("%01x", c.PCStack[2]),
		"Level3Stack":        fmt.Sprintf("%01x", c.PCStack[3]),
		"RAM":                ram,
		"PROM":               fmt.Sprintf("%01x", c.PROM[:]),
	})

}

func postStep(ctx *fiber.Ctx) error {

	c.Step()

	ram := ""
	for index := range c.RAMData[0:64] {
		ram += fmt.Sprintf("%01x", c.RAMData[index])
	}

	regs := fmt.Sprintf("REG=%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x",
		c.Registers[R0], c.Registers[R1], c.Registers[R2], c.Registers[R3], c.Registers[R4], c.Registers[R5], c.Registers[R6], c.Registers[R7],
		c.Registers[R8], c.Registers[R9], c.Registers[RA], c.Registers[RB], c.Registers[RC], c.Registers[RD], c.Registers[RE], c.Registers[RF])

	return ctx.Render("cpuinfo", fiber.Map{
		"CPUInfo":            c.CPUInfo(),
		"Model":              c.Model,
		"Accumulator":        fmt.Sprintf("%01x", c.Accumulator),
		"Carry":              fmt.Sprintf("%01x", c.Carry),
		"RAMAddressRegister": fmt.Sprintf("%01x", c.RAMAddressRegister),
		"StackPointer":       fmt.Sprintf("%01x", c.StackPointer),
		"PCStack":            fmt.Sprintf("%01x", c.PCStack[0]),
		"Level1Stack":        fmt.Sprintf("%01x", c.PCStack[1]),
		"Level2Stack":        fmt.Sprintf("%01x", c.PCStack[2]),
		"Level3Stack":        fmt.Sprintf("%01x", c.PCStack[3]),
		"RAM":                ram,
		"Registers":          regs,
		"PROM":               fmt.Sprintf("%01x", c.PROM[:]),
	})

}
