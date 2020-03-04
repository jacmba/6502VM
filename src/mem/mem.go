package mem

/*
Mem Data type to emulate physical memory
*/
type Mem struct {
	table []byte // Array of bytes holding the linear memory
}

const (
	kb = 1024
)

/*
Make creates a memory instance with given memory in KBs
*/
func Make(size uint16) *Mem {
	return &Mem{
		table: make([]byte, size*kb),
	}
}

/*
SetByte set value of a byte in given address position
*/
func (m *Mem) SetByte(address uint16, val byte) {
	m.table[address] = val
}

/*
ReadByte returns value from given address
*/
func (m *Mem) ReadByte(address uint16) byte {
	return m.table[address]
}

/*
Dump returns complete array of memory
*/
func (m *Mem) Dump() []byte {
	return m.table
}
