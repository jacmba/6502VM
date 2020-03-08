package cpu

// Constant values for Opcodes
const (
	OpLoadAI   = "OPCODE_LOAD_A_IMMEDIATE"
	OpLoadXI   = "OPCODE_LOAD_X_IMMEDIATE"
	OpLoadYI   = "OPCODE_LOAD_Y_IMMEDIATE"
	OpLoadAZP  = "OPCODE_LOAD_A_ZEROPAGE"
	OpLoadXZP  = "OPCODE_LOAD_X_ZEROPAGE"
	OpLoadYZP  = "OPCODE_LOAD_Y_ZEROPAGE"
	OpLoadAAbs = "OPCODE_LOAD_A_ABSOLUTE"
	OpLoadXAbs = "OPCODE_LOAD_X_ABSOLUTE"
	OpLoadYAbs = "OPCODE_LOAD_Y_ABSOLUTE"
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
}
