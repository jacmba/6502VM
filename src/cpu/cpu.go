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

	negativeFlagMask = 0x80
	zeroFlagMask     = 0x02
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
exec parses opcode and executes instruction at current PC position
*/
func (c *CPU) exec() {
	currentPos := c.pc
	opcode := c.getNextByte()

	switch OpcodeMap[opcode] {

	//============================================================================
	// Load opcodes
	//============================================================================
	case OpLoadAI:
		val := c.getNextByte()
		c.LoadImmediateA(val)
	case OpLoadXI:
		val := c.getNextByte()
		c.LoadImmediateX(val)
	case OpLoadYI:
		val := c.getNextByte()
		c.LoadImmediateY(val)
	case OpLoadAZP:
		val := c.getNextByte()
		c.LoadZeroPageA(val)
	case OpLoadXZP:
		val := c.getNextByte()
		c.LoadZeroPageX(val)
	case OpLoadYZP:
		val := c.getNextByte()
		c.LoadZeroPageY(val)
	case OpLoadAAbs:
		val := c.getNextWord()
		c.LoadAbsoluteA(val)
	case OpLoadXAbs:
		val := c.getNextWord()
		c.LoadAbsoluteX(val)
	case OpLoadYAbs:
		val := c.getNextWord()
		c.LoadAbsoluteY(val)
	case OpLoadXAZP:
		val := c.getNextByte()
		c.LoadZeroPageXA(val)
	case OpLoadYXZP:
		val := c.getNextByte()
		c.LoadZeroPageYX(val)
	case OpLoadXYZP:
		val := c.getNextByte()
		c.LoadZeroPageXY(val)
	case OpLoadXAAbs:
		val := c.getNextWord()
		c.LoadAbsoluteXA(val)
	case OpLoadYXAbs:
		val := c.getNextWord()
		c.LoadAbsoluteYX(val)
	case OpLoadXYAbs:
		val := c.getNextWord()
		c.LoadAbsoluteXY(val)
	case OpLoadXIndirect:
		val := c.getNextByte()
		c.LoadIndirectX(val)
	case OpLoadYIndirect:
		val := c.getNextByte()
		c.LoadIndirectY(val)

	//============================================================================
	// Store opcodes
	//============================================================================
	case OpStoreAZP:
		addr := c.getNextByte()
		c.StoreZeropageA(addr)
	case OpStoreIAZP:
		addr := c.getNextByte()
		c.StoreZeropageIA(addr)
	case OpStoreXZP:
		addr := c.getNextByte()
		c.StoreZeropageX(addr)
	case OpStoreIXZP:
		addr := c.getNextByte()
		c.StoreZeropageIX(addr)
	case OpStoreYZP:
		addr := c.getNextByte()
		c.StoreZeropageY(addr)
	case OpStoreIYZP:
		addr := c.getNextByte()
		c.StoreZeropageIY(addr)

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
getNextWord - Gets current 16bit word from PC position
*/
func (c *CPU) getNextWord() uint16 {
	low := uint16(c.getNextByte())
	high := uint16(c.getNextByte()) << 8

	return high | low
}

/*
RunOnce - Runs one iteration resetting PC to 0x00
*/
func (c *CPU) RunOnce() {
	c.pc = pageTwoBegin
	c.exec()
}

/*
setZeroFlag - Sets the zero flag in flags register
*/
func (c *CPU) setZeroFlag() {
	c.p |= zeroFlagMask
}

/*
setNegativeFlag - Sets the negative flag in flags register
*/
func (c *CPU) setNegativeFlag() {
	c.p |= negativeFlagMask
}

/*
setRegister - Put byte value in given reference register
*/
func (c *CPU) setRegister(reg *byte, val byte) {
	if val == 0x00 {
		c.setZeroFlag()
	}
	*reg = val
}
