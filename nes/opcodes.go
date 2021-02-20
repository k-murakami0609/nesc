package nes

type AddressingMode int
type InstructionName string

// addressing Modes
const (
	_ AddressingMode = iota
	ModeAbsolute
	ModeAbsoluteX
	ModeAbsoluteY
	ModeAccumulator
	ModeImmediate
	ModeImplied
	ModeIndexedIndirect
	ModeIndirect
	ModeIndirectIndexed
	ModeRelative
	ModeZeroPage
	ModeZeroPageX
	ModeZeroPageY
)

const (
	ADC = InstructionName("ADC")
	AND = InstructionName("AND")
	ASL = InstructionName("ASL")
	BCC = InstructionName("BCC")
	BCS = InstructionName("BCS")
	BEQ = InstructionName("BEQ")
	BIT = InstructionName("BIT")
	BMI = InstructionName("BMI")
	BNE = InstructionName("BNE")
	BPL = InstructionName("BPL")
	BRK = InstructionName("BRK")
	BVC = InstructionName("BVC")
	BVS = InstructionName("BVS")
	CLC = InstructionName("CLC")
	CLD = InstructionName("CLD")
	CLI = InstructionName("CLI")
	CLV = InstructionName("CLV")
	CMP = InstructionName("CMP")
	CPX = InstructionName("CPX")
	CPY = InstructionName("CPY")
	DEC = InstructionName("DEC")
	DEX = InstructionName("DEX")
	DEY = InstructionName("DEY")
	EOR = InstructionName("EOR")
	INC = InstructionName("INC")
	INX = InstructionName("INX")
	INY = InstructionName("INY")
	JMP = InstructionName("JMP")
	JSR = InstructionName("JSR")
	LDA = InstructionName("LDA")
	LDX = InstructionName("LDX")
	LDY = InstructionName("LDY")
	LSR = InstructionName("LSR")
	NOP = InstructionName("NOP")
	ORA = InstructionName("ORA")
	PHA = InstructionName("PHA")
	PHP = InstructionName("PHP")
	PLA = InstructionName("PLA")
	PLP = InstructionName("PLP")
	ROL = InstructionName("ROL")
	ROR = InstructionName("ROR")
	RTI = InstructionName("RTI")
	RTS = InstructionName("RTS")
	SBC = InstructionName("SBC")
	SEC = InstructionName("SEC")
	SED = InstructionName("SED")
	SEI = InstructionName("SEI")
	STA = InstructionName("STA")
	STX = InstructionName("STX")
	STY = InstructionName("STY")
	TAX = InstructionName("TAX")
	TAY = InstructionName("TAY")
	TSX = InstructionName("TSX")
	TXA = InstructionName("TXA")
	TXS = InstructionName("TXS")
	TYA = InstructionName("TYA")
)

type Opcode struct {
	Name      InstructionName
	Mode      AddressingMode
	Code      int
	Cycle     int
	Size      int
	PageCycle int
}

var Opcodes = CreateOpcodes()

// http://obelisk.me.uk/6502/reference.html
// http://wiki.nesdev.com/w/index.php/Programming_with_unofficial_opcodes
func CreateOpcodes() map[byte]Opcode {
	opcodes := map[byte]Opcode{}

	opcodes[0x69] = Opcode{Name: ADC, Mode: ModeImmediate, Code: 0x69, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0x65] = Opcode{Name: ADC, Mode: ModeZeroPage, Code: 0x65, Cycle: 3, Size: 2, PageCycle: 0}
	opcodes[0x75] = Opcode{Name: ADC, Mode: ModeZeroPageX, Code: 0x75, Cycle: 4, Size: 2, PageCycle: 0}
	opcodes[0x6D] = Opcode{Name: ADC, Mode: ModeAbsolute, Code: 0x6D, Cycle: 4, Size: 3, PageCycle: 0}
	opcodes[0x7D] = Opcode{Name: ADC, Mode: ModeAbsoluteX, Code: 0x7D, Cycle: 4, Size: 3, PageCycle: 1}
	opcodes[0x79] = Opcode{Name: ADC, Mode: ModeAbsoluteY, Code: 0x79, Cycle: 4, Size: 3, PageCycle: 1}
	opcodes[0x61] = Opcode{Name: ADC, Mode: ModeIndexedIndirect, Code: 0x61, Cycle: 6, Size: 2, PageCycle: 0}
	opcodes[0x71] = Opcode{Name: ADC, Mode: ModeIndirectIndexed, Code: 0x71, Cycle: 5, Size: 2, PageCycle: 1}
	opcodes[0x29] = Opcode{Name: AND, Mode: ModeImmediate, Code: 0x29, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0x25] = Opcode{Name: AND, Mode: ModeZeroPage, Code: 0x25, Cycle: 3, Size: 2, PageCycle: 0}
	opcodes[0x35] = Opcode{Name: AND, Mode: ModeZeroPageX, Code: 0x35, Cycle: 4, Size: 2, PageCycle: 0}
	opcodes[0x2D] = Opcode{Name: AND, Mode: ModeAbsolute, Code: 0x2D, Cycle: 4, Size: 3, PageCycle: 0}
	opcodes[0x3D] = Opcode{Name: AND, Mode: ModeAbsoluteX, Code: 0x3D, Cycle: 4, Size: 3, PageCycle: 1}
	opcodes[0x39] = Opcode{Name: AND, Mode: ModeAbsoluteY, Code: 0x39, Cycle: 4, Size: 3, PageCycle: 1}
	opcodes[0x21] = Opcode{Name: AND, Mode: ModeIndexedIndirect, Code: 0x21, Cycle: 6, Size: 2, PageCycle: 0}
	opcodes[0x31] = Opcode{Name: AND, Mode: ModeIndirectIndexed, Code: 0x31, Cycle: 5, Size: 2, PageCycle: 1}
	opcodes[0x0A] = Opcode{Name: ASL, Mode: ModeAccumulator, Code: 0x0A, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0x06] = Opcode{Name: ASL, Mode: ModeZeroPage, Code: 0x06, Cycle: 5, Size: 2, PageCycle: 0}
	opcodes[0x16] = Opcode{Name: ASL, Mode: ModeZeroPageX, Code: 0x16, Cycle: 6, Size: 2, PageCycle: 0}
	opcodes[0x0E] = Opcode{Name: ASL, Mode: ModeAbsolute, Code: 0x0E, Cycle: 6, Size: 3, PageCycle: 0}
	opcodes[0x1E] = Opcode{Name: ASL, Mode: ModeAbsoluteX, Code: 0x1E, Cycle: 7, Size: 3, PageCycle: 0}
	opcodes[0x90] = Opcode{Name: BCC, Mode: ModeRelative, Code: 0x90, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0xB0] = Opcode{Name: BCS, Mode: ModeRelative, Code: 0xB0, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0xF0] = Opcode{Name: BEQ, Mode: ModeRelative, Code: 0xF0, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0x24] = Opcode{Name: BIT, Mode: ModeZeroPage, Code: 0x24, Cycle: 3, Size: 2, PageCycle: 0}
	opcodes[0x2C] = Opcode{Name: BIT, Mode: ModeAbsolute, Code: 0x2C, Cycle: 4, Size: 3, PageCycle: 0}
	opcodes[0x30] = Opcode{Name: BMI, Mode: ModeRelative, Code: 0x30, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0xD0] = Opcode{Name: BNE, Mode: ModeRelative, Code: 0xD0, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0x10] = Opcode{Name: BPL, Mode: ModeRelative, Code: 0x10, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0x00] = Opcode{Name: BRK, Mode: ModeImplied, Code: 0x00, Cycle: 7, Size: 1, PageCycle: 0}
	opcodes[0x50] = Opcode{Name: BVC, Mode: ModeRelative, Code: 0x50, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0x70] = Opcode{Name: BVS, Mode: ModeRelative, Code: 0x70, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0x18] = Opcode{Name: CLC, Mode: ModeImplied, Code: 0x18, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0xD8] = Opcode{Name: CLD, Mode: ModeImplied, Code: 0xD8, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0x58] = Opcode{Name: CLI, Mode: ModeImplied, Code: 0x58, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0xB8] = Opcode{Name: CLV, Mode: ModeImplied, Code: 0xB8, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0xC9] = Opcode{Name: CMP, Mode: ModeImmediate, Code: 0xC9, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0xC5] = Opcode{Name: CMP, Mode: ModeZeroPage, Code: 0xC5, Cycle: 3, Size: 2, PageCycle: 0}
	opcodes[0xD5] = Opcode{Name: CMP, Mode: ModeZeroPageX, Code: 0xD5, Cycle: 4, Size: 2, PageCycle: 0}
	opcodes[0xCD] = Opcode{Name: CMP, Mode: ModeAbsolute, Code: 0xCD, Cycle: 4, Size: 3, PageCycle: 0}
	opcodes[0xDD] = Opcode{Name: CMP, Mode: ModeAbsoluteX, Code: 0xDD, Cycle: 4, Size: 3, PageCycle: 1}
	opcodes[0xD9] = Opcode{Name: CMP, Mode: ModeAbsoluteY, Code: 0xD9, Cycle: 4, Size: 3, PageCycle: 1}
	opcodes[0xC1] = Opcode{Name: CMP, Mode: ModeIndexedIndirect, Code: 0xC1, Cycle: 6, Size: 2, PageCycle: 0}
	opcodes[0xD1] = Opcode{Name: CMP, Mode: ModeIndirectIndexed, Code: 0xD1, Cycle: 5, Size: 2, PageCycle: 1}
	opcodes[0xE0] = Opcode{Name: CPX, Mode: ModeImmediate, Code: 0xE0, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0xE4] = Opcode{Name: CPX, Mode: ModeZeroPage, Code: 0xE4, Cycle: 3, Size: 2, PageCycle: 0}
	opcodes[0xEC] = Opcode{Name: CPX, Mode: ModeAbsolute, Code: 0xEC, Cycle: 4, Size: 3, PageCycle: 0}
	opcodes[0xC0] = Opcode{Name: CPY, Mode: ModeImmediate, Code: 0xC0, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0xC4] = Opcode{Name: CPY, Mode: ModeZeroPage, Code: 0xC4, Cycle: 3, Size: 2, PageCycle: 0}
	opcodes[0xCC] = Opcode{Name: CPY, Mode: ModeAbsolute, Code: 0xCC, Cycle: 4, Size: 3, PageCycle: 0}
	opcodes[0xC6] = Opcode{Name: DEC, Mode: ModeZeroPage, Code: 0xC6, Cycle: 5, Size: 2, PageCycle: 0}
	opcodes[0xD6] = Opcode{Name: DEC, Mode: ModeZeroPageX, Code: 0xD6, Cycle: 6, Size: 2, PageCycle: 0}
	opcodes[0xCE] = Opcode{Name: DEC, Mode: ModeAbsolute, Code: 0xCE, Cycle: 6, Size: 3, PageCycle: 0}
	opcodes[0xDE] = Opcode{Name: DEC, Mode: ModeAbsoluteX, Code: 0xDE, Cycle: 7, Size: 3, PageCycle: 0}
	opcodes[0xCA] = Opcode{Name: DEX, Mode: ModeImplied, Code: 0xCA, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0x88] = Opcode{Name: DEY, Mode: ModeImplied, Code: 0x88, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0x49] = Opcode{Name: EOR, Mode: ModeImmediate, Code: 0x49, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0x45] = Opcode{Name: EOR, Mode: ModeZeroPage, Code: 0x45, Cycle: 3, Size: 2, PageCycle: 0}
	opcodes[0x55] = Opcode{Name: EOR, Mode: ModeZeroPageX, Code: 0x55, Cycle: 4, Size: 2, PageCycle: 0}
	opcodes[0x4D] = Opcode{Name: EOR, Mode: ModeAbsolute, Code: 0x4D, Cycle: 4, Size: 3, PageCycle: 0}
	opcodes[0x5D] = Opcode{Name: EOR, Mode: ModeAbsoluteX, Code: 0x5D, Cycle: 4, Size: 3, PageCycle: 1}
	opcodes[0x59] = Opcode{Name: EOR, Mode: ModeAbsoluteY, Code: 0x59, Cycle: 4, Size: 3, PageCycle: 1}
	opcodes[0x41] = Opcode{Name: EOR, Mode: ModeIndexedIndirect, Code: 0x41, Cycle: 6, Size: 2, PageCycle: 0}
	opcodes[0x51] = Opcode{Name: EOR, Mode: ModeIndirectIndexed, Code: 0x51, Cycle: 5, Size: 2, PageCycle: 1}
	opcodes[0xE6] = Opcode{Name: INC, Mode: ModeZeroPage, Code: 0xE6, Cycle: 5, Size: 2, PageCycle: 0}
	opcodes[0xF6] = Opcode{Name: INC, Mode: ModeZeroPageX, Code: 0xF6, Cycle: 6, Size: 2, PageCycle: 0}
	opcodes[0xEE] = Opcode{Name: INC, Mode: ModeAbsolute, Code: 0xEE, Cycle: 6, Size: 3, PageCycle: 0}
	opcodes[0xFE] = Opcode{Name: INC, Mode: ModeAbsoluteX, Code: 0xFE, Cycle: 7, Size: 3, PageCycle: 0}
	opcodes[0xE8] = Opcode{Name: INX, Mode: ModeImplied, Code: 0xE8, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0xC8] = Opcode{Name: INY, Mode: ModeImplied, Code: 0xC8, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0x4C] = Opcode{Name: JMP, Mode: ModeAbsolute, Code: 0x4C, Cycle: 3, Size: 3, PageCycle: 0}
	opcodes[0x6C] = Opcode{Name: JMP, Mode: ModeIndirect, Code: 0x6C, Cycle: 5, Size: 3, PageCycle: 0}
	opcodes[0x20] = Opcode{Name: JSR, Mode: ModeAbsolute, Code: 0x20, Cycle: 6, Size: 3, PageCycle: 0}
	opcodes[0xA9] = Opcode{Name: LDA, Mode: ModeImmediate, Code: 0xA9, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0xA5] = Opcode{Name: LDA, Mode: ModeZeroPage, Code: 0xA5, Cycle: 3, Size: 2, PageCycle: 0}
	opcodes[0xB5] = Opcode{Name: LDA, Mode: ModeZeroPageX, Code: 0xB5, Cycle: 4, Size: 2, PageCycle: 0}
	opcodes[0xAD] = Opcode{Name: LDA, Mode: ModeAbsolute, Code: 0xAD, Cycle: 4, Size: 3, PageCycle: 0}
	opcodes[0xBD] = Opcode{Name: LDA, Mode: ModeAbsoluteX, Code: 0xBD, Cycle: 4, Size: 3, PageCycle: 1}
	opcodes[0xB9] = Opcode{Name: LDA, Mode: ModeAbsoluteY, Code: 0xB9, Cycle: 4, Size: 3, PageCycle: 1}
	opcodes[0xA1] = Opcode{Name: LDA, Mode: ModeIndexedIndirect, Code: 0xA1, Cycle: 6, Size: 2, PageCycle: 0}
	opcodes[0xB1] = Opcode{Name: LDA, Mode: ModeIndirectIndexed, Code: 0xB1, Cycle: 5, Size: 2, PageCycle: 1}
	opcodes[0xA2] = Opcode{Name: LDX, Mode: ModeImmediate, Code: 0xA2, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0xA6] = Opcode{Name: LDX, Mode: ModeZeroPage, Code: 0xA6, Cycle: 3, Size: 2, PageCycle: 0}
	opcodes[0xB6] = Opcode{Name: LDX, Mode: ModeZeroPageY, Code: 0xB6, Cycle: 4, Size: 2, PageCycle: 0}
	opcodes[0xAE] = Opcode{Name: LDX, Mode: ModeAbsolute, Code: 0xAE, Cycle: 4, Size: 3, PageCycle: 0}
	opcodes[0xBE] = Opcode{Name: LDX, Mode: ModeAbsoluteY, Code: 0xBE, Cycle: 4, Size: 3, PageCycle: 1}
	opcodes[0xA0] = Opcode{Name: LDY, Mode: ModeImmediate, Code: 0xA0, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0xA4] = Opcode{Name: LDY, Mode: ModeZeroPage, Code: 0xA4, Cycle: 3, Size: 2, PageCycle: 0}
	opcodes[0xB4] = Opcode{Name: LDY, Mode: ModeZeroPageX, Code: 0xB4, Cycle: 4, Size: 2, PageCycle: 0}
	opcodes[0xAC] = Opcode{Name: LDY, Mode: ModeAbsolute, Code: 0xAC, Cycle: 4, Size: 3, PageCycle: 0}
	opcodes[0xBC] = Opcode{Name: LDY, Mode: ModeAbsoluteX, Code: 0xBC, Cycle: 4, Size: 3, PageCycle: 1}
	opcodes[0x4A] = Opcode{Name: LSR, Mode: ModeAccumulator, Code: 0x4A, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0x46] = Opcode{Name: LSR, Mode: ModeZeroPage, Code: 0x46, Cycle: 5, Size: 2, PageCycle: 0}
	opcodes[0x56] = Opcode{Name: LSR, Mode: ModeZeroPageX, Code: 0x56, Cycle: 6, Size: 2, PageCycle: 0}
	opcodes[0x4E] = Opcode{Name: LSR, Mode: ModeAbsolute, Code: 0x4E, Cycle: 6, Size: 3, PageCycle: 0}
	opcodes[0x5E] = Opcode{Name: LSR, Mode: ModeAbsoluteX, Code: 0x5E, Cycle: 7, Size: 3, PageCycle: 0}
	opcodes[0xEA] = Opcode{Name: NOP, Mode: ModeImplied, Code: 0xEA, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0x09] = Opcode{Name: ORA, Mode: ModeImmediate, Code: 0x09, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0x05] = Opcode{Name: ORA, Mode: ModeZeroPage, Code: 0x05, Cycle: 3, Size: 2, PageCycle: 0}
	opcodes[0x15] = Opcode{Name: ORA, Mode: ModeZeroPageX, Code: 0x15, Cycle: 4, Size: 2, PageCycle: 0}
	opcodes[0x0D] = Opcode{Name: ORA, Mode: ModeAbsolute, Code: 0x0D, Cycle: 4, Size: 3, PageCycle: 0}
	opcodes[0x1D] = Opcode{Name: ORA, Mode: ModeAbsoluteX, Code: 0x1D, Cycle: 4, Size: 3, PageCycle: 1}
	opcodes[0x19] = Opcode{Name: ORA, Mode: ModeAbsoluteY, Code: 0x19, Cycle: 4, Size: 3, PageCycle: 1}
	opcodes[0x01] = Opcode{Name: ORA, Mode: ModeIndexedIndirect, Code: 0x01, Cycle: 6, Size: 2, PageCycle: 0}
	opcodes[0x11] = Opcode{Name: ORA, Mode: ModeIndirectIndexed, Code: 0x11, Cycle: 5, Size: 2, PageCycle: 1}
	opcodes[0x48] = Opcode{Name: PHA, Mode: ModeImplied, Code: 0x48, Cycle: 3, Size: 1, PageCycle: 0}
	opcodes[0x08] = Opcode{Name: PHP, Mode: ModeImplied, Code: 0x08, Cycle: 3, Size: 1, PageCycle: 0}
	opcodes[0x68] = Opcode{Name: PLA, Mode: ModeImplied, Code: 0x68, Cycle: 4, Size: 1, PageCycle: 0}
	opcodes[0x28] = Opcode{Name: PLP, Mode: ModeImplied, Code: 0x28, Cycle: 4, Size: 1, PageCycle: 0}
	opcodes[0x2A] = Opcode{Name: ROL, Mode: ModeAccumulator, Code: 0x2A, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0x26] = Opcode{Name: ROL, Mode: ModeZeroPage, Code: 0x26, Cycle: 5, Size: 2, PageCycle: 0}
	opcodes[0x36] = Opcode{Name: ROL, Mode: ModeZeroPageX, Code: 0x36, Cycle: 6, Size: 2, PageCycle: 0}
	opcodes[0x2E] = Opcode{Name: ROL, Mode: ModeAbsolute, Code: 0x2E, Cycle: 6, Size: 3, PageCycle: 0}
	opcodes[0x3E] = Opcode{Name: ROL, Mode: ModeAbsoluteX, Code: 0x3E, Cycle: 7, Size: 3, PageCycle: 0}
	opcodes[0x6A] = Opcode{Name: ROR, Mode: ModeAccumulator, Code: 0x6A, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0x66] = Opcode{Name: ROR, Mode: ModeZeroPage, Code: 0x66, Cycle: 5, Size: 2, PageCycle: 0}
	opcodes[0x76] = Opcode{Name: ROR, Mode: ModeZeroPageX, Code: 0x76, Cycle: 6, Size: 2, PageCycle: 0}
	opcodes[0x6E] = Opcode{Name: ROR, Mode: ModeAbsolute, Code: 0x6E, Cycle: 6, Size: 3, PageCycle: 0}
	opcodes[0x7E] = Opcode{Name: ROR, Mode: ModeAbsoluteX, Code: 0x7E, Cycle: 7, Size: 3, PageCycle: 0}
	opcodes[0x40] = Opcode{Name: RTI, Mode: ModeImplied, Code: 0x40, Cycle: 6, Size: 1, PageCycle: 0}
	opcodes[0x60] = Opcode{Name: RTS, Mode: ModeImplied, Code: 0x60, Cycle: 6, Size: 1, PageCycle: 0}
	opcodes[0xE9] = Opcode{Name: SBC, Mode: ModeImmediate, Code: 0xE9, Cycle: 2, Size: 2, PageCycle: 0}
	opcodes[0xE5] = Opcode{Name: SBC, Mode: ModeZeroPage, Code: 0xE5, Cycle: 3, Size: 2, PageCycle: 0}
	opcodes[0xF5] = Opcode{Name: SBC, Mode: ModeZeroPageX, Code: 0xF5, Cycle: 4, Size: 2, PageCycle: 0}
	opcodes[0xED] = Opcode{Name: SBC, Mode: ModeAbsolute, Code: 0xED, Cycle: 4, Size: 3, PageCycle: 0}
	opcodes[0xFD] = Opcode{Name: SBC, Mode: ModeAbsoluteX, Code: 0xFD, Cycle: 4, Size: 3, PageCycle: 1}
	opcodes[0xF9] = Opcode{Name: SBC, Mode: ModeAbsoluteY, Code: 0xF9, Cycle: 4, Size: 3, PageCycle: 1}
	opcodes[0xE1] = Opcode{Name: SBC, Mode: ModeIndexedIndirect, Code: 0xE1, Cycle: 6, Size: 2, PageCycle: 0}
	opcodes[0xF1] = Opcode{Name: SBC, Mode: ModeIndirectIndexed, Code: 0xF1, Cycle: 5, Size: 2, PageCycle: 1}
	opcodes[0x38] = Opcode{Name: SEC, Mode: ModeImplied, Code: 0x38, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0xF8] = Opcode{Name: SED, Mode: ModeImplied, Code: 0xF8, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0x78] = Opcode{Name: SEI, Mode: ModeImplied, Code: 0x78, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0x85] = Opcode{Name: STA, Mode: ModeZeroPage, Code: 0x85, Cycle: 3, Size: 2, PageCycle: 0}
	opcodes[0x95] = Opcode{Name: STA, Mode: ModeZeroPageX, Code: 0x95, Cycle: 4, Size: 2, PageCycle: 0}
	opcodes[0x8D] = Opcode{Name: STA, Mode: ModeAbsolute, Code: 0x8D, Cycle: 4, Size: 3, PageCycle: 0}
	opcodes[0x9D] = Opcode{Name: STA, Mode: ModeAbsoluteX, Code: 0x9D, Cycle: 5, Size: 3, PageCycle: 0}
	opcodes[0x99] = Opcode{Name: STA, Mode: ModeAbsoluteY, Code: 0x99, Cycle: 5, Size: 3, PageCycle: 0}
	opcodes[0x81] = Opcode{Name: STA, Mode: ModeIndexedIndirect, Code: 0x81, Cycle: 6, Size: 2, PageCycle: 0}
	opcodes[0x91] = Opcode{Name: STA, Mode: ModeIndirectIndexed, Code: 0x91, Cycle: 6, Size: 2, PageCycle: 0}
	opcodes[0x86] = Opcode{Name: STX, Mode: ModeZeroPage, Code: 0x86, Cycle: 3, Size: 2, PageCycle: 0}
	opcodes[0x96] = Opcode{Name: STX, Mode: ModeZeroPageY, Code: 0x96, Cycle: 4, Size: 2, PageCycle: 0}
	opcodes[0x8E] = Opcode{Name: STX, Mode: ModeAbsolute, Code: 0x8E, Cycle: 4, Size: 3, PageCycle: 0}
	opcodes[0x84] = Opcode{Name: STY, Mode: ModeZeroPage, Code: 0x84, Cycle: 3, Size: 2, PageCycle: 0}
	opcodes[0x94] = Opcode{Name: STY, Mode: ModeZeroPageX, Code: 0x94, Cycle: 4, Size: 2, PageCycle: 0}
	opcodes[0x8C] = Opcode{Name: STY, Mode: ModeAbsolute, Code: 0x8C, Cycle: 4, Size: 3, PageCycle: 0}
	opcodes[0xAA] = Opcode{Name: TAX, Mode: ModeImplied, Code: 0xAA, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0xA8] = Opcode{Name: TAY, Mode: ModeImplied, Code: 0xA8, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0xBA] = Opcode{Name: TSX, Mode: ModeImplied, Code: 0xBA, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0x8A] = Opcode{Name: TXA, Mode: ModeImplied, Code: 0x8A, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0x9A] = Opcode{Name: TXS, Mode: ModeImplied, Code: 0x9A, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0x98] = Opcode{Name: TYA, Mode: ModeImplied, Code: 0x98, Cycle: 2, Size: 1, PageCycle: 0}
	opcodes[0xEA] = Opcode{Name: NOP, Mode: ModeImplied, Code: 0xEA, Cycle: 2, Size: 1, PageCycle: 0}

	return opcodes
}
