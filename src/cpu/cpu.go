package cpu

/*
CPU Data structure to hold CPU registers and flags
*/
type CPU struct {
	x   byte   // X general purpose register
	y   byte   // Y general purpose register
	a   byte   // Accumulator register
	s   byte   // Stack Pointer register
	p   byte   // CPU flags register
	irq uint16 // Interrupt routine address
	pc  uint16 // Program Counter
}

/*
Make creates a new CPU instance
*/
func Make() *CPU {
	return &CPU{
		x:   0,
		y:   0,
		a:   0,
		s:   0,
		p:   0,
		irq: 0,
		pc:  0,
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
