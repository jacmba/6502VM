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

func TestStoreAbsoluteA(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	c.a = 0xAB
	c.StoreAbsoluteA(0x1FF)
	v := m.ReadByte(0x1FF)
	if v != 0xAB {
		t.Fatal("Address 0x1FF should have value 0xAB")
	}
}

func TestStoreAbsoluteXA(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	c.a = 0xAC
	c.x = 0x0F
	c.StoreAbsoluteXA(0x100)
	v := m.ReadByte(0x10F)
	if v != 0xAC {
		t.Fatal("Address 0x10F should have value 0xAC")
	}
}

func TestStoreAbsoluteYA(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	c.a = 0xAD
	c.y = 0x0F
	c.StoreAbsoluteYA(0x110)
	v := m.ReadByte(0x11F)
	if v != 0xAD {
		t.Fatal("Address 0x11F should have value 0xAD")
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

func TestAbsoluteX(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	c.x = 0xBA
	c.StoreAbsoluteX(0x1AA)
	v := m.ReadByte(0x1AA)
	if v != 0xBA {
		t.Fatal("Address 0x1AA should have value 0xBA")
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

func TestAbsoluteY(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	c.y = 0xCA
	c.StoreAbsoluteY(0x1BB)
	v := m.ReadByte(0x1BB)
	if v != 0xCA {
		t.Fatal("Address 0x1BB should have value 0xCA")
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

func TestStoreAAbsoluteOpcodeExecution(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x200, 0x8D)
	m.SetByte(0x201, 0xFF)
	m.SetByte(0x202, 0x01)
	c.a = 0xAA
	c.exec()
	v := m.ReadByte(0x01FF)
	if v != 0xAA {
		t.Fatal("Address 0x01FF should have value 0xAA")
	}
}

func TestStoreAAbsoluteXOpcodeExecution(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x200, 0x9D)
	m.SetByte(0x201, 0xF0)
	m.SetByte(0x202, 0x01)
	c.a = 0xAA
	c.x = 0x0F
	c.exec()
	v := m.ReadByte(0x01FF)
	if v != 0xAA {
		t.Fatal("Address 0x01FF should have value 0xAA")
	}
}

func TestStoreAAbsoluteYOpcodeExecution(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x200, 0x99)
	m.SetByte(0x201, 0xF0)
	m.SetByte(0x202, 0x01)
	c.a = 0xAA
	c.y = 0x0F
	c.exec()
	v := m.ReadByte(0x01FF)
	if v != 0xAA {
		t.Fatal("Address 0x01FF should have value 0xAA")
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

func TestStoreAbsoluteXOpcodeExecution(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x200, 0x8E)
	m.SetByte(0x201, 0xBF)
	m.SetByte(0x202, 0x01)
	c.x = 0xBB
	c.exec()
	v := m.ReadByte(0x01BF)
	if v != 0xBB {
		t.Fatal("Address 0x01BF should have value 0xBB")
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

func TestStoreAbsoluteYOpcodeExecution(t *testing.T) {
	m := mem.Make(1)
	c := Make(m)
	m.SetByte(0x200, 0x8C)
	m.SetByte(0x201, 0xAF)
	m.SetByte(0x202, 0x01)
	c.y = 0xCC
	c.exec()
	v := m.ReadByte(0x01AF)
	if v != 0xCC {
		t.Fatal("Address 0x01AF should have value 0xCC")
	}
}
