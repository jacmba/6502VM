package cpu

//========================================================================
// This file contains implementation of LOAD instructions
//========================================================================

/*
LoadImmediateA sets a numeric value into accumulator register
*/
func (c *CPU) LoadImmediateA(b byte) {
	c.a = b
}

/*
LoadZeroPageA loads accoumulator from zero page address
*/
func (c *CPU) LoadZeroPageA(b byte) {
	c.a = c.mem.ReadByte(uint16(b))
}

/*
LoadAbsoluteA loads accumulator from absoulte 16bit address
*/
func (c *CPU) LoadAbsoluteA(addr uint16) {
	c.a = c.mem.ReadByte(addr)
}

/*
LoadImmediateX sets numeric value into index register X
*/
func (c *CPU) LoadImmediateX(b byte) {
	c.x = b
}

/*
LoadZeroPageX sets in register X value from zero page address
*/
func (c *CPU) LoadZeroPageX(b byte) {
	c.x = c.mem.ReadByte(uint16(b))
}

/*
LoadAbsoluteX sets in register X value from absolute 16bit address
*/
func (c *CPU) LoadAbsoluteX(addr uint16) {
	c.x = c.mem.ReadByte(addr)
}

/*
LoadImmediateY sets numeric value into index register Y
*/
func (c *CPU) LoadImmediateY(b byte) {
	c.y = b
}

/*
LoadZeroPageY sets in register Y value from zero page address
*/
func (c *CPU) LoadZeroPageY(b byte) {
	c.y = c.mem.ReadByte(uint16(b))
}

/*
LoadAbsoluteY sets in register Y value from absolute 16bit address
*/
func (c *CPU) LoadAbsoluteY(addr uint16) {
	c.y = c.mem.ReadByte(addr)
}
