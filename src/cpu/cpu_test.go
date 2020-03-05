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

func TestLoadImmediateX(t *testing.T) {
	c := Make(nil)
	c.LoadImmediateX(0x07)
	if c.x != 7 {
		t.Fatal("Register X should have value 7")
	}
}

func TestLoadImmediateY(t *testing.T) {
	c := Make(nil)
	c.LoadImmediateY(0x0A)
	if c.y != 10 {
		t.Fatal("Register Y should have value 10")
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
