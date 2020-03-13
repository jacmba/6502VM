package cpu

//==============================================================================

// This file contains methods to implement memory storage instructions

//==============================================================================

/*
StoreZeropageA stores value from accumulator into Zeropage 8bit address
*/
func (c *CPU) StoreZeropageA(addr byte) {
	c.mem.SetByte(uint16(addr), c.a)
}

/*
StoreZeropageIA stores value from accumulator into indexed addr+X ZP address
*/
func (c *CPU) StoreZeropageIA(addr byte) {
	dest := uint16(addr) + uint16(c.x)
	c.mem.SetByte(dest, c.a)
}

/*
StoreZeropageX stores value from register X into Zeropage 8bit address
*/
func (c *CPU) StoreZeropageX(addr byte) {
	c.mem.SetByte(uint16(addr), c.x)
}

/*
StoreZeropageIX stores value from accumulator into indexed addr+Y ZP address
*/
func (c *CPU) StoreZeropageIX(addr byte) {
	dest := uint16(addr) + uint16(c.y)
	c.mem.SetByte(dest, c.x)
}

/*
StoreZeropageY stores value from register Y into Zeropage 8bit address
*/
func (c *CPU) StoreZeropageY(addr byte) {
	c.mem.SetByte(uint16(addr), c.y)
}

/*
StoreZeropageIY stores value from register Y into indexed addr+X ZP address
*/
func (c *CPU) StoreZeropageIY(addr byte) {
	dest := uint16(addr) + uint16(c.x)
	c.mem.SetByte(dest, c.y)
}
