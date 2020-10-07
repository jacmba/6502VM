package mem

import "testing"

func TestMemoryObjectIsCreated(t *testing.T) {
	mem := Make(1)
	if mem == nil {
		t.Fatal("Memory object is nil")
	}
	if len(mem.table) != 1024 {
		t.Fatal("Memory shold have a size of 1024 bytes")
	}
	if mem.table[1023] != 0 {
		t.Fatal("Memory bytes should be initialized to 0")
	}
}

func TestSetByte(t *testing.T) {
	mem := Make(16)
	mem.SetByte(10, 99)
	if mem.table[10] != 99 {
		t.Fatal("Address 0x0A should have value 99")
	}
}

func TestReadByte(t *testing.T) {
	mem := Make(2)
	mem.table[15] = 255
	val := mem.ReadByte(0x0F)
	if val != 0xFF {
		t.Fatal("Address 0x0F should have value 0xFF")
	}
}
