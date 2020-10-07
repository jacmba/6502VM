package cpu

import (
	"testing"

	"6502VM/mem"
)

func TestStoreZeropageA(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	c.a = 0xAA
	c.StoreZeropageA(0xFA)
	v := m.ReadByte(0xFA)
	if v != 0xAA {
		t.Fatal("Address 0xFA should have value 0xAA")
	}
}

func TestStoreZeropageIndexedA(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	c.a = 0xAA
	c.x = 0x05
	c.StoreZeropageIA(0xF5)
	v := m.ReadByte(0xFA)
	if v != 0xAA {
		t.Fatal("Address 0xFA should have value 0xAA")
	}
}

func TestStoreZeropageX(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	c.x = 0xBB
	c.StoreZeropageX(0xFB)
	v := m.ReadByte(0xFB)
	if v != 0xBB {
		t.Fatal("Address 0xFB should have value 0xBB")
	}
}

func TestStoreZeropageIndexedX(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	c.x = 0xBB
	c.y = 0x02
	c.StoreZeropageIX(0xF9)
	v := m.ReadByte(0xFB)
	if v != 0xBB {
		t.Fatal("Address 0xFB should have value 0xBB")
	}
}

func TestStoreZeropageY(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	c.y = 0xCC
	c.StoreZeropageY(0xFC)
	v := m.ReadByte(0xFC)
	if v != 0xCC {
		t.Fatal("Address 0xFC should have value 0xCC")
	}
}

func TestStoreZeropageIndexedY(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	c.y = 0xCC
	c.x = 0x0A
	c.StoreZeropageIY(0xF2)
	v := m.ReadByte(0xFC)
	if v != 0xCC {
		t.Fatal("Address 0xFC should have value 0xCC")
	}
}

//==============================================================================
// Opcode execution tests
//==============================================================================

func TestStoreZeropageAOpcodeExecution(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x200, 0x85)
	m.SetByte(0x201, 0xFA)
	c.a = 0xAA
	c.exec()
	v := m.ReadByte(0xFA)
	if v != 0xAA {
		t.Fatal("Address 0xFA should have value 0xAA")
	}
}

func TestStoreZeropageIndexedAOpcodeExecution(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x200, 0x95)
	m.SetByte(0x201, 0xF5)
	c.a = 0xAA
	c.x = 0x05
	c.exec()
	v := m.ReadByte(0xFA)
	if v != 0xAA {
		t.Fatal("Address 0xFA should have value 0xAA")
	}
}

func TestStoreZeropageXOpcodeExecution(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x200, 0x86)
	m.SetByte(0x201, 0xFB)
	c.x = 0xBB
	c.exec()
	v := m.ReadByte(0xFB)
	if v != 0xBB {
		t.Fatal("Address 0xFB should have value 0xBB")
	}
}

func TestStoreZeropageIndexedXOpcodeExecution(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x200, 0x96)
	m.SetByte(0x201, 0xF7)
	c.x = 0xBB
	c.y = 0x04
	c.exec()
	v := m.ReadByte(0xFB)
	if v != 0xBB {
		t.Fatal("Address 0xFB should have value 0xBB")
	}
}

func TestStoreZeropageYOpcodeExecution(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x200, 0x84)
	m.SetByte(0x201, 0xFC)
	c.y = 0xCC
	c.exec()
	v := m.ReadByte(0xFC)
	if v != 0xCC {
		t.Fatal("Address 0xFC should have value 0xCC")
	}
}

func TestStoreZeropageIndexedYOpcodeExecution(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x200, 0x94)
	m.SetByte(0x201, 0xFA)
	c.y = 0xCC
	c.x = 0x02
	c.exec()
	v := m.ReadByte(0xFC)
	if v != 0xCC {
		t.Fatal("Address 0xFC should have value 0xCC")
	}
}
