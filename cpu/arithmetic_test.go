package cpu

import (
	"6502VM/mem"
	"testing"
)

func TestAdcImmediate(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	c.a = 0x1F
	c.AdcImmediate(0xE0)
	if c.a != 0xFF {
		t.Fatalf("Accumulator should have value 0xFF. Found %X\n", c.a)
	}
}

func TestAdcZeropage(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	c.a = 0x0B
	c.mem.SetByte(0x10, 0xA0)
	c.AdcZeropage(0x10)
	if c.a != 0xAB {
		t.Fatalf("Accumulator should have value 0xAB. Found %X\n", c.a)
	}
}

func TestAdcZeropageIndexed(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	c.a = 0x02
	c.x = 0x0F
	c.mem.SetByte(0x1F, 0x03)
	c.AdcZeropageIndexed(0x10)
	if c.a != 0x05 {
		t.Fatalf("Accumulator should have value 0x05. Found %X\n", c.a)
	}
}

func TestAdcAbsolute(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	c.a = 0x05
	c.mem.SetByte(0x01FF, 0x05)
	c.AdcAbsolute(0x01FF)
	if c.a != 0x0A {
		t.Fatalf("Accumulator should have value 0x0A. Found %X\n", c.a)
	}
}

func TestAdcAbsoluteIndexedX(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	c.a = 0xA0
	c.x = 0xFF
	c.mem.SetByte(0x01FF, 0x0A)
	c.AdcAbsoluteIndexedX(0x0100)
	if c.a != 0xAA {
		t.Fatalf("Accumulator should have value 0xAA. Found %X\n", c.a)
	}
}

func TestAdcAbsoluteIndexedY(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	c.a = 0xB0
	c.y = 0xEE
	c.mem.SetByte(0x02EE, 0x0B)
	c.AdcAbsoluteIndexedY(0x0200)
	if c.a != 0xBB {
		t.Fatalf("Accumulator should have value 0xBB. Found %X\n", c.a)
	}
}
