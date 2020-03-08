package cpu

import (
	"testing"
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

func TestSetZeroFlag(t *testing.T) {
	c := Make(nil)
	c.setZeroFlag()
	if c.p != 0x02 {
		t.Fatal("Flags register should be 00000010")
	}
}

func TestSetNegativeFlag(t *testing.T) {
	c := Make(nil)
	c.setNegativeFlag()
	if c.p != 0x80 {
		t.Fatal("Flags register should be 10000000")
	}
}

func TestSetNegativeAndZeroFlag(t *testing.T) {
	c := Make(nil)
	c.setNegativeFlag()
	c.setZeroFlag()
	if c.p != 0x82 {
		t.Fatal("Flags register should be 10000010")
	}
}

func TestSetRegister(t *testing.T) {
	c := Make(nil)
	c.setRegister(&c.x, 0xFF)
	if c.x != 0xFF {
		t.Fatal("Register X should be 0xFF")
	}
	if c.p != 0x00 {
		t.Fatal("Flags register should remain 00000000")
	}
}

func TestSetRegisterRaiseZeroFlag(t *testing.T) {
	c := Make(nil)
	c.setRegister(&c.a, 0x00)
	if c.a != 0x00 {
		t.Fatal("Accumulator should be 0x00")
	}
	if c.p != 0x02 {
		t.Fatal("Flags register should be 00000010")
	}
}
