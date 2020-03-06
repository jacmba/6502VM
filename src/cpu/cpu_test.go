package cpu

import (
	"testing"

	"../mem"
)

func TestCPUObjectExists(t *testing.T) {
	var c *CPU
	c = Make(nil)
	if c == nil {
		t.Fatal("CPU Object is null")
	}
}

func TestCPUHasInitializedFlags(t *testing.T) {
	c := Make(nil)
	if c.x != 0 {
		t.Fatal("CPU should have X register initalized to 0")
	}

	if c.y != 0 {
		t.Fatal("CPU should have Y register initalized to 0")
	}

	if c.a != 0 {
		t.Fatal("CPU should have Accumulator initialized to 0")
	}

	if c.s != 0x100 {
		t.Fatal("CPU should have Stack Pointer initialized to page 1 (0x0100)")
	}

	if c.p != 0 {
		t.Fatal("CPU should have flags register initialzed to 0")
	}

	if c.pc != 0x200 {
		t.Fatal("CPU should have Program Counter initialzed to Page 2 (0x200)")
	}

	if c.irq != 0 {
		t.Fatal("CPU should have IRQ initialized to 0")
	}
}

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

func TestLoadAbsoluteX(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x301, 0xAB)
	c.LoadAbsoluteX(0x301)
	if c.x != 0xAB {
		t.Fatal("Register X should have value 0xAB")
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

func TestLoadAbsoluteY(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x302, 0xAC)
	c.LoadAbsoluteY(0x302)
	if c.y != 0xAC {
		t.Fatal("Register Y should have value 0xAC")
	}
}

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
		t.Fatal("Register X should halv value 0xAC")
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
