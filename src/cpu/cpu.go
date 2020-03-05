package cpu

import (
	"fmt"

	"../mem"
)

const (
	pageZeroBegin = 0x0000
	pageZeroEnd   = 0x00FF
	pageOneBegin  = 0x0100
	pageOneEnd    = 0x01FF
	pageTwoBegin  = 0x0200
)

/*
CPU Data structure to hold CPU registers and flags
*/
type CPU struct {
	x   byte     // X general purpose register
	y   byte     // Y general purpose register
	a   byte     // Accumulator register
	p   byte     // CPU flags register
	s   uint16   // Stack Pointer register
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
		s:   pageOneBegin,
		p:   0,
		irq: 0,
		pc:  pageTwoBegin,
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

/*
RunOnce - Runs one iteration resetting PC to 0x00
*/
func (c *CPU) RunOnce() {
	c.pc = pageTwoBegin
	c.exec()
}
