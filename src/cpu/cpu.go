package cpu

import (
	"fmt"

	"../mem"
)

/*
CPU Data structure to hold CPU registers and flags
*/
type CPU struct {
	x   byte     // X general purpose register
	y   byte     // Y general purpose register
	a   byte     // Accumulator register
	s   byte     // Stack Pointer register
	p   byte     // CPU flags register
	irq uint16   // Interrupt routine address
	pc  uint16   // Program Counter
	mem *mem.Mem // Access to RAM memory
}

/*
Make creates a new CPU instance
*/
func Make(mem *mem.Mem) *CPU {
	return &CPU{
		x:   0,
		y:   0,
		a:   0,
		s:   0,
		p:   0,
		irq: 0,
		pc:  0,
		mem: mem,
	}
}

/*
LoadImmediateA sets a numeric value into accumulator register
*/
func (c *CPU) LoadImmediateA(b byte) {
	c.a = b
}

/*
LoadImmediateX sets numeric value into index register X
*/
func (c *CPU) LoadImmediateX(b byte) {
	c.x = b
}

/*
LoadImmediateY sets numeric value into index register Y
*/
func (c *CPU) LoadImmediateY(b byte) {
	c.y = b
}

/*
exec parses opcode and executes instruction at current PC position
*/
func (c *CPU) exec() {
	currentPos := c.pc
	opcode := c.getNextByte()

	switch OpcodeMap[opcode] {
	case OpLoadAI:
		val := c.getNextByte()
		c.LoadImmediateA(val)
	case OpLoadXI:
		val := c.getNextByte()
		c.LoadImmediateX(val)
	case OpLoadYI:
		val := c.getNextByte()
		c.LoadImmediateY(val)
	default:
		panic(fmt.Sprintf("Invalid opcode [%#X] found at %#X!!", opcode, currentPos))
	}
}

/*
getNextByte - Gets current byte from PC position and advances PC
*/
func (c *CPU) getNextByte() byte {
	b := c.mem.ReadByte(c.pc)
	c.pc++
	return b
}
