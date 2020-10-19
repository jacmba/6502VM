package cpu

// Constant values for Opcodes
const (
	OpLoadAI        = "OPCODE_LOAD_A_IMMEDIATE"
	OpLoadXI        = "OPCODE_LOAD_X_IMMEDIATE"
	OpLoadYI        = "OPCODE_LOAD_Y_IMMEDIATE"
	OpLoadAZP       = "OPCODE_LOAD_A_ZEROPAGE"
	OpLoadXZP       = "OPCODE_LOAD_X_ZEROPAGE"
	OpLoadYZP       = "OPCODE_LOAD_Y_ZEROPAGE"
	OpLoadAAbs      = "OPCODE_LOAD_A_ABSOLUTE"
	OpLoadXAbs      = "OPCODE_LOAD_X_ABSOLUTE"
	OpLoadYAbs      = "OPCODE_LOAD_Y_ABSOLUTE"
	OpLoadXAZP      = "OPCODE_LOAD_A_INDEXED_ZEROPAGE"
	OpLoadYXZP      = "OPCODE_LOAD_X_INDEXED_ZEROPAGE"
	OpLoadXYZP      = "OPCODE_LOAD_Y_INDEXED_ZEROPAGE"
	OpLoadXAAbs     = "OPCODE_LOAD_A_INDEXED_ABSOLUTE"
	OpLoadYXAbs     = "OPCODE_LOAD_X_INDEXED_ABSOLUTE"
	OpLoadXYAbs     = "OPCODE_LOAD_Y_INDEXED_ABSOLUTE"
	OpLoadXIndirect = "OPCODE_LOAD_X_INDIRECT"
	OpLoadYIndirect = "OPCODE_LOAD_Y_INDIRECT"

	OpStoreAZP       = "OPCODE_STORE_A_ZEROPAGE"
	OpStoreXZP       = "OPCODE_STORE_X_ZEROPAGE"
	OpStoreYZP       = "OPCODE_STORE_Y_ZEROPAGE"
	OpStoreIAZP      = "OPCODE_STORE_A_INDEXED_ZEROPAGE"
	OpStoreIXZP      = "OPCODE_STORE_X_INDEXED_ZEROPAGE"
	OpStoreIYZP      = "OPCODE_STORE_Y_INDEXED_ZEROPAGE"
	OpStoreAAbs      = "OPCODE_STORE_A_ABSOLUTE"
	OpStoreAXAbs     = "OPCODE_STORE_AX_INDEXED_ABSOLUTE"
	OpStoreAYAbs     = "OPCODE_STORE_AY_INDEXED_ABSOLUTE"
	OpStoreXAbs      = "OPCODE_STORE_X_ABSOLUTE"
	OpStoreYAbs      = "OPCODE_STORE_Y_ABSOLUTE"
	OpStoreXIndirect = "OPCODE_STORE_X_INDIRECT"
	OpStoreYIndirect = "OPCODE_STORE_Y_INDIRECT"
)

//OpcodeMap - map of Opcode HEX values
var OpcodeMap map[byte]string = map[byte]string{
	0xA9: OpLoadAI,
	0xA2: OpLoadXI,
	0xA0: OpLoadYI,
	0xA5: OpLoadAZP,
	0xA6: OpLoadXZP,
	0xA4: OpLoadYZP,
	0xAD: OpLoadAAbs,
	0xAE: OpLoadXAbs,
	0xAC: OpLoadYAbs,
	0xB5: OpLoadXAZP,
	0xB6: OpLoadYXZP,
	0xB4: OpLoadXYZP,
	0xBD: OpLoadXAAbs,
	0xBE: OpLoadYXAbs,
	0xBC: OpLoadXYAbs,
	0xA1: OpLoadXIndirect,
	0xB1: OpLoadYIndirect,

	0x85: OpStoreAZP,
	0x86: OpStoreXZP,
	0x84: OpStoreYZP,
	0x95: OpStoreIAZP,
	0x96: OpStoreIXZP,
	0x94: OpStoreIYZP,
	0x8D: OpStoreAAbs,
	0x9D: OpStoreAXAbs,
	0x99: OpStoreAYAbs,
	0x8E: OpStoreXAbs,
	0x8C: OpStoreYAbs,
	0x81: OpStoreXIndirect,
	0x91: OpStoreYIndirect,
}
