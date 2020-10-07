package cpu

//========================================================================
// This file contains implementation of LOAD instructions
//========================================================================

/*
LoadImmediateA sets a numeric value into accumulator register
*/
func (c *CPU) LoadImmediateA(b byte) {
	c.setRegister(&c.a, b)
}

/*
LoadZeroPageA loads accoumulator from zero page address
*/
func (c *CPU) LoadZeroPageA(b byte) {
	val := c.mem.ReadByte(uint16(b))
	c.setRegister(&c.a, val)
}

/*
LoadZeroPageXA loads into accumulator byte from address + X reg value
*/
func (c *CPU) LoadZeroPageXA(b byte) {
	val := c.mem.ReadByte(uint16(b) + uint16(c.x))
	c.setRegister(&c.a, val)
}

/*
LoadAbsoluteA loads accumulator from absoulte 16bit address
*/
func (c *CPU) LoadAbsoluteA(addr uint16) {
	val := c.mem.ReadByte(addr)
	c.setRegister(&c.a, val)
}

/*
LoadAbsoluteXA loads accumulator from X-indexed 16bit address
*/
func (c *CPU) LoadAbsoluteXA(addr uint16) {
	val := c.mem.ReadByte(addr + uint16(c.x))
	c.setRegister(&c.a, val)
}

/*
LoadIndirectX loads accumulator from address starting at address + X
*/
func (c *CPU) LoadIndirectX(addr byte) {
	low := uint16(c.mem.ReadByte(uint16(addr) + uint16(c.x)))
	high := uint16(c.mem.ReadByte(uint16(addr)+uint16(c.x)+1)) << 8
	orig := high | low
	val := c.mem.ReadByte(orig)
	c.setRegister(&c.a, val)
}

/*
LoadIndirectY loads accumulator from indirect address + Y
*/
func (c *CPU) LoadIndirectY(addr byte) {
	low := uint16(c.mem.ReadByte(uint16(addr)))
	high := uint16(c.mem.ReadByte(uint16(addr)+1)) << 8
	orig := (high | low) + uint16(c.y)
	val := c.mem.ReadByte(orig)
	c.setRegister(&c.a, val)
}

/*
LoadImmediateX sets numeric value into index register X
*/
func (c *CPU) LoadImmediateX(b byte) {
	c.setRegister(&c.x, b)
}

/*
LoadZeroPageX sets in register X value from zero page address
*/
func (c *CPU) LoadZeroPageX(b byte) {
	val := c.mem.ReadByte(uint16(b))
	c.setRegister(&c.x, val)
}

/*
LoadZeroPageYX sets in register X Y-indexed value fom zeropage address
*/
func (c *CPU) LoadZeroPageYX(b byte) {
	val := c.mem.ReadByte(uint16(b) + uint16(c.y))
	c.setRegister(&c.x, val)
}

/*
LoadAbsoluteX sets in register X value from absolute 16bit address
*/
func (c *CPU) LoadAbsoluteX(addr uint16) {
	val := c.mem.ReadByte(addr)
	c.setRegister(&c.x, val)
}

/*
LoadAbsoluteYX sets in register X value from Y-indexed absolute 16b address
*/
func (c *CPU) LoadAbsoluteYX(addr uint16) {
	val := c.mem.ReadByte(addr + uint16(c.y))
	c.setRegister(&c.x, val)
}

/*
LoadImmediateY sets numeric value into index register Y
*/
func (c *CPU) LoadImmediateY(b byte) {
	c.setRegister(&c.y, b)
}

/*
LoadZeroPageY sets in register Y value from zero page address
*/
func (c *CPU) LoadZeroPageY(b byte) {
	val := c.mem.ReadByte(uint16(b))
	c.setRegister(&c.y, val)
}

/*
LoadZeroPageXY sets in register Y X-indexed value from zeropage address
*/
func (c *CPU) LoadZeroPageXY(b byte) {
	val := c.mem.ReadByte(uint16(b) + uint16(c.x))
	c.setRegister(&c.y, val)
}

/*
LoadAbsoluteY sets in register Y value from absolute 16bit address
*/
func (c *CPU) LoadAbsoluteY(addr uint16) {
	val := c.mem.ReadByte(addr)
	c.setRegister(&c.y, val)
}

/*
LoadAbsoluteXY sets in register Y value from X-indexed absolute 16b address
*/
func (c *CPU) LoadAbsoluteXY(addr uint16) {
	val := c.mem.ReadByte(addr + uint16(c.x))
	c.setRegister(&c.y, val)
}
