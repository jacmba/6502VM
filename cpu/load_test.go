package cpu

import (
	"testing"

	"6502VM/mem"
)

func TestLoadImmediateAccumulator(t *testing.T) {
	c := Make(nil)
	c.LoadImmediateA(0x05)
	if c.a != 5 {
		t.Fatal("Accumulator should have value 5")
	}
}

func TestLoadZeroPageA(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x0A, 0xEA)
	c.LoadZeroPageA(0x0A)
	if c.a != 0xEA {
		t.Fatal("Accumulator should have value 0xEA")
	}
}

func TestLoadAbsoluteA(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x0300, 0xAA)
	c.LoadAbsoluteA(0x300)
	if c.a != 0xAA {
		t.Fatal("Accumulator should have value 0xAA")
	}
}

func TestLoadZeropageXA(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0xEE, 0xAA)
	c.x = 0x01
	c.LoadZeroPageXA(0xED)
	if c.a != 0xAA {
		t.Fatal("Accumulator should have 0xAA")
	}
}

func TestLoadAbsoluteXA(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x1A0, 0xAA)
	c.x = 0x05
	c.LoadAbsoluteXA(0x19B)
	if c.a != 0xAA {
		t.Fatal("Accumulator should have 0xAA")
	}
}

func TestLoadIndirectX(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0xFB, 0x00)
	m.SetByte(0xFC, 0x01)
	m.SetByte(0x100, 0xAA)
	c.x = 0x01
	c.LoadIndirectX(0xFA)
	if c.a != 0xAA {
		t.Fatal("Accumulator should have 0xAA")
	}
}

func TestLoadIndirectY(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0xFE, 0xAA)
	m.SetByte(0xFF, 0x01)
	m.SetByte(0x1AF, 0xAB)
	c.y = 0x05
	c.LoadIndirectY(0x0FE)
	if c.a != 0xAB {
		t.Fatal("Accumulator should have 0xAB")
	}
}

func TestLoadImmediateX(t *testing.T) {
	c := Make(nil)
	c.LoadImmediateX(0x07)
	if c.x != 7 {
		t.Fatal("Register X should have value 7")
	}
}

func TestLoadZeroPageX(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x0B, 0xAD)
	c.LoadZeroPageX(0x0B)
	if c.x != 0xAD {
		t.Fatal("Register X should have value 0xAD")
	}
}

func TestLoadZeroPageYX(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0xEE, 0x11)
	c.y = 0x02
	c.LoadZeroPageYX(0xEC)
	if c.x != 0x11 {
		t.Fatal("Register X should have value 0x11")
	}
}

func TestLoadAbsoluteX(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x301, 0xAB)
	c.LoadAbsoluteX(0x301)
	if c.x != 0xAB {
		t.Fatal("Register X should have value 0xAB")
	}
}

func TestLoadAbsoluteYX(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x3B0, 0xBB)
	c.y = 0x03
	c.LoadAbsoluteYX(0x3AD)
	if c.x != 0xBB {
		t.Fatal("Register X should have value 0xBB")
	}
}

func TestLoadImmediateY(t *testing.T) {
	c := Make(nil)
	c.LoadImmediateY(0x0A)
	if c.y != 10 {
		t.Fatal("Register Y should have value 10")
	}
}

func TestLoadZeroPageY(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x0C, 0xDB)
	c.LoadZeroPageY(0x0C)
	if c.y != 0xDB {
		t.Fatal("Register Y should have value 0xDB")
	}
}

func TestLoadZeropageXY(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x0C, 0xBF)
	c.x = 0x03
	c.LoadZeroPageXY(0x09)
	if c.y != 0xBF {
		t.Fatal("Register Y should have value 0xDB")
	}
}

func TestLoadAbsoluteY(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x302, 0xAC)
	c.LoadAbsoluteY(0x302)
	if c.y != 0xAC {
		t.Fatal("Register Y should have value 0xAC")
	}
}

func TestLoadAbsoluteXY(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x2FF, 0xCC)
	c.x = 0x0A
	c.LoadAbsoluteXY(0x2F5)
	if c.y != 0xCC {
		t.Fatal("Register Y should have value 0xCC")
	}
}

//==============================================================================
// Opcode execution tests
//==============================================================================

func TestLoadImmediateAOpcodeExecution(t *testing.T) {
	mem := mem.Make(1)
	c := Make(mem)
	mem.SetByte(0x200, 0xA9)
	mem.SetByte(0x201, 0xBB)
	c.exec()
	if c.a != 0xBB {
		t.Fatal("Accumulator should have value 0xBB")
	}
	if c.pc != 0x202 {
		t.Fatal("Program counter should hold positon 0x202")
	}
}

func TestLoadZeroPageAOpcodeExecution(t *testing.T) {
	mem := mem.Make(1)
	c := Make(mem)
	mem.SetByte(0x200, 0xA5)
	mem.SetByte(0x201, 0xBB)
	mem.SetByte(0xBB, 0xFA)
	c.exec()
	if c.a != 0xFA {
		t.Fatal("Accumulator should have value 0xFA")
	}
}

func TestLoadZeropageXAOpcodeExecution(t *testing.T) {
	mem := mem.Make(1)
	c := Make(mem)
	mem.SetByte(0x200, 0xB5)
	mem.SetByte(0x201, 0xAA)
	mem.SetByte(0xAB, 0xFA)
	c.x = 0x01
	c.exec()
	if c.a != 0xFA {
		t.Fatal("Accumulator should have value 0xFA")
	}
}

func TestLoadAbsoluteAOpcodeExecution(t *testing.T) {
	mem := mem.Make(1)
	c := Make(mem)
	mem.SetByte(0x200, 0xAD)
	mem.SetByte(0x201, 0xCC)
	mem.SetByte(0x202, 0x01)
	mem.SetByte(0x1CC, 0xFB)
	c.exec()
	if c.a != 0xFB {
		t.Fatal("Accumulator should have value 0xFB")
	}
}

func TestLoadAbsoluteXAOpcodeExecution(t *testing.T) {
	mem := mem.Make(1)
	c := Make(mem)
	mem.SetByte(0x200, 0xBD)
	mem.SetByte(0x201, 0xA8)
	mem.SetByte(0x202, 0x01)
	mem.SetByte(0x1AA, 0xA1)
	c.x = 0x02
	c.exec()
	if c.a != 0xA1 {
		t.Fatal("Accumulator should have value 0xA1")
	}
}

func TestLoadIndirectXOpcodeExecution(t *testing.T) {
	mem := mem.Make(1)
	c := Make(mem)
	mem.SetByte(0x200, 0xA1)
	mem.SetByte(0x201, 0xFA)
	mem.SetByte(0xFE, 0xFF)
	mem.SetByte(0xFF, 0x01)
	mem.SetByte(0x01FF, 0xAA)
	c.x = 0x04
	c.exec()
	if c.a != 0xAA {
		t.Fatal("Accumulator should have value 0xAA")
	}
}

func TestLoadIndirectYOpcodeExecution(t *testing.T) {
	mem := mem.Make(1)
	c := Make(mem)
	mem.SetByte(0x200, 0xB1)
	mem.SetByte(0x201, 0x00)
	mem.SetByte(0x00, 0x00)
	mem.SetByte(0x01, 0x01)
	mem.SetByte(0x010E, 0xAB)
	c.y = 0x0E
	c.exec()
	if c.a != 0xAB {
		t.Fatal("Accumulator should have value 0xAB")
	}
}

func TestLoadImmediateXOpcodeExecution(t *testing.T) {
	mem := mem.Make(1)
	c := Make(mem)
	mem.SetByte(0x200, 0xA2)
	mem.SetByte(0x201, 0xCC)
	c.exec()
	if c.x != 0xCC {
		t.Fatal("Register X should have value 0xCC")
	}
	if c.pc != 0x202 {
		t.Fatal("Program counter should hold position 0x202")
	}
}

func TestLoadZeropageXOpcodeExecution(t *testing.T) {
	mem := mem.Make(1)
	c := Make(mem)
	mem.SetByte(0x200, 0xA6)
	mem.SetByte(0x201, 0xAA)
	mem.SetByte(0xAA, 0xAC)
	c.exec()
	if c.x != 0xAC {
		t.Fatal("Register X should have value 0xAC")
	}
}

func TestLoadZeropageYXOpcodeExecution(t *testing.T) {
	mem := mem.Make(1)
	c := Make(mem)
	mem.SetByte(0x200, 0xB6)
	mem.SetByte(0x201, 0xBB)
	mem.SetByte(0xBC, 0xFB)
	c.y = 0x01
	c.exec()
	if c.x != 0xFB {
		t.Fatal("Register X should have value 0xFB")
	}
}

func TestLoadAbsoluteXOpcodeExecution(t *testing.T) {
	mem := mem.Make(1)
	c := Make(mem)
	mem.SetByte(0x200, 0xAE)
	mem.SetByte(0x201, 0x00)
	mem.SetByte(0x202, 0x03)
	mem.SetByte(0x300, 0xFC)
	c.exec()
	if c.x != 0xFC {
		t.Fatal("Register X should have value 0xFC")
	}
}

func TestLoadAbsoluteYXOpcodeExecution(t *testing.T) {
	mem := mem.Make(1)
	c := Make(mem)
	mem.SetByte(0x200, 0xBE)
	mem.SetByte(0x201, 0xBB)
	mem.SetByte(0x202, 0x01)
	mem.SetByte(0x01BC, 0xBB)
	c.y = 0x01
	c.exec()
	if c.x != 0xBB {
		t.Fatal("Register X should have value 0xBB")
	}
}

func TestLoadImmediateYOpcodeExecution(t *testing.T) {
	mem := mem.Make(1)
	c := Make(mem)
	mem.SetByte(0x200, 0xA0)
	mem.SetByte(0x201, 0xDD)
	c.exec()
	if c.y != 0xDD {
		t.Fatal("Register Y should have value 0xDD")
	}
	if c.pc != 0x202 {
		t.Fatal("Program counter should hold position 0x202")
	}
}

func TestLoadZeroPageYOpcodeExecution(t *testing.T) {
	mem := mem.Make(1)
	c := Make(mem)
	mem.SetByte(0x200, 0xA4)
	mem.SetByte(0x201, 0x1F)
	mem.SetByte(0x1F, 0xEB)
	c.exec()
	if c.y != 0xEB {
		t.Fatal("Register Y should have value 0xEB")
	}
}

func TestLoadZeropageXYOpcodeExecution(t *testing.T) {
	mem := mem.Make(1)
	c := Make(mem)
	mem.SetByte(0x200, 0xB4)
	mem.SetByte(0x201, 0xCC)
	mem.SetByte(0xCD, 0xFF)
	c.x = 0x01
	c.exec()
	if c.y != 0xFF {
		t.Fatal("Register Y should have value 0xFF")
	}
}

func TestLoadAbsoluteYOpcodeExecution(t *testing.T) {
	mem := mem.Make(1)
	c := Make(mem)
	mem.SetByte(0x200, 0xAC)
	mem.SetByte(0x201, 0xFF)
	mem.SetByte(0x202, 0x03)
	mem.SetByte(0x3FF, 0xEC)
	c.exec()
	if c.y != 0xEC {
		t.Fatal("Register Y should have value 0xEC")
	}
}

func TestLoadAbsoluteXYOpcodeExecution(t *testing.T) {
	mem := mem.Make(1)
	c := Make(mem)
	mem.SetByte(0x200, 0xBC)
	mem.SetByte(0x201, 0xCC)
	mem.SetByte(0x202, 0x01)
	mem.SetByte(0x01CF, 0xCA)
	c.x = 0x03
	c.exec()
	if c.y != 0xCA {
		t.Fatal("Register Y should have value 0xCA")
	}
}
