package i4004

import "fmt"

func (c *CPU) Debug() string {
	//mnem, args := c.disasmPC()

	//disasm := fmt.Sprintf("ASM %04x %02x => %-8s%-8s", c.PC, c.Memory[c.PC], mnem, args)

	regs := fmt.Sprintf("REG=%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x%01x",
		c.Registers[R0], c.Registers[R1], c.Registers[R2], c.Registers[R3], c.Registers[R4], c.Registers[R5], c.Registers[R6], c.Registers[R7],
		c.Registers[R8], c.Registers[R9], c.Registers[RA], c.Registers[RB], c.Registers[RC], c.Registers[RD], c.Registers[RE], c.Registers[RF])

	ptrs := fmt.Sprintf("A=%04x C=%01x RA=%01x PC=%01x PC1=%01x PC2=%01x PC3=%01x",
		c.Accumulator, c.Carry, c.RAMAddressRegister, c.PC, c.PC1, c.PC2, c.PC3)

	return fmt.Sprintf("%s %s", ptrs, regs)
}
