package cpu

//================================================================================================
// This file contains implementation of arithmetic operations: add/subtract with carry
//================================================================================================

/*
add sums given value to accumulator and handles flags
*/
func (c *CPU) add(val byte) {
	a := uint16(c.a)
	b := uint16(val)
	x := a + b
	mask := uint16(0x00FF)

	c.a = byte(x & mask)
}

/*
AdcImmediate adds given value to accumulator
*/
func (c *CPU) AdcImmediate(val byte) {
	c.add(val)
}

/*
AdcZeropage adds value in zeropage address to accumulator
*/
func (c *CPU) AdcZeropage(addr byte) {
	v := c.mem.ReadByte(uint16(addr))
	c.add(v)
}

/*
AdcZeropageIndexed adds value in zeropage + X offset address to accumulator
*/
func (c *CPU) AdcZeropageIndexed(addr byte) {
	v := c.mem.ReadByte(uint16(addr + c.x))
	c.add(v)
}

/*
AdcAbsolute adds value in absolute 16 bit address to accumulator
*/
func (c *CPU) AdcAbsolute(addr uint16) {
	v := c.mem.ReadByte(addr)
	c.add(v)
}

/*
AdcAbsoluteIndexedX adds value in absolute 16 bit address offset X to accumulator
*/
func (c *CPU) AdcAbsoluteIndexedX(addr uint16) {
	v := c.mem.ReadByte(addr + uint16(c.x))
	c.add(v)
}

/*
AdcAbsoluteIndexedY adds value in absolute 16 bit address offset Y to accumulator
*/
func (c *CPU) AdcAbsoluteIndexedY(addr uint16) {
	v := c.mem.ReadByte(addr + uint16(c.y))
	c.add(v)
}
