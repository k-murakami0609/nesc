package nes

func ExecuteOpration(bus *CpuBus, reg *CpuRegister, opcode Opcode) int {
	address, pageCrossed := address(bus, reg, opcode)
	reg.PC += uint16(opcode.Size)

	opration := findOpration(opcode)
	isBranchSuccess := opration(bus, reg, address)

	return sumCycles(opcode, pageCrossed, isBranchSuccess, reg.PC, address)
}

// http://obelisk.me.uk/6502/reference.html
func sumCycles(opcode Opcode, pageCrossed bool, isBranchSuccess bool, pc uint16, address uint16) int {
	cycles := opcode.Cycle
	if pageCrossed {
		cycles += opcode.PageCycle
	}
	if isBranchSuccess {
		cycles += 1
		if pagesDiffer(pc, address) {
			cycles += 1
		}
	}
	return cycles
}

// https://github.com/fogleman/nes/blob/8c4b9cf54c354137c37e8ae17caf4b1b1405313b/nes/cpu.go#L288
func pagesDiffer(a, b uint16) bool {
	return a&0xFF00 != b&0xFF00
}

// http://nesdev.com/NESDoc.pdf#page=39
func address(bus *CpuBus, reg *CpuRegister, o Opcode) (uint16, bool) {
	var address uint16 = 0
	var pageCrossed bool = false
	switch o.Mode {
	case ModeAbsolute:
		address = bus.Read16(reg.PC + 1)
	case ModeAbsoluteX:
		address = bus.Read16(reg.PC+1) + uint16(reg.X)
		pageCrossed = pagesDiffer(address-uint16(reg.X), address)
	case ModeAbsoluteY:
		address = bus.Read16(reg.PC+1) + uint16(reg.Y)
		pageCrossed = pagesDiffer(address-uint16(reg.Y), address)
	case ModeAccumulator:
		// nothing
	case ModeImmediate:
		address = uint16(reg.PC + 1)
	case ModeImplied:
		// nothing
	case ModeIndexedIndirect:
		address = bus.Read16bug(uint16(bus.Read(reg.PC+1) + reg.X))
	case ModeIndirect:
		address = bus.Read16bug(bus.Read16(reg.PC + 1))
	case ModeIndirectIndexed:
		address = bus.Read16bug(uint16(bus.Read(reg.PC+1))) + uint16(reg.Y)
		pageCrossed = pagesDiffer(address-uint16(reg.Y), address)
	case ModeRelative:
		offset := uint16(bus.Read(reg.PC + 1))
		if offset < 0x80 {
			address = reg.PC + 2 + offset
		} else {
			address = reg.PC + 2 + offset - 0x100
		}
	case ModeZeroPage:
		address = uint16(bus.Read(reg.PC + 1))
	case ModeZeroPageX:
		address = uint16(bus.Read(reg.PC+1) + reg.X)
	case ModeZeroPageY:
		address = uint16(bus.Read(reg.PC+1) + reg.Y)
	default:
		panic("address")
	}
	return address, pageCrossed
}

func findOpration(opcode Opcode) func(*CpuBus, *CpuRegister, uint16) bool {
	wrap := func(instruction func(bus *CpuBus, reg *CpuRegister, address uint16)) func(bus *CpuBus, reg *CpuRegister, address uint16) bool {
		return func(bus *CpuBus, reg *CpuRegister, address uint16) bool {
			instruction(bus, reg, address)
			return false
		}
	}

	switch opcode.Name {
	case ADC:
		return wrap(adc)
	case AND:
		return wrap(and)
	case ASL:
		return wrap(func(bus *CpuBus, reg *CpuRegister, address uint16) { asl(opcode, bus, reg, address) })
	case BCC:
		return bcc
	case BCS:
		return bcs
	case BEQ:
		return beq
	case BIT:
		return wrap(bit)
	case BMI:
		return bmi
	case BNE:
		return bne
	case BPL:
		return bpl
	case BRK:
		return wrap(brk)
	case BVC:
		return bvc
	case BVS:
		return bvs
	case CLC:
		return wrap(clc)
	case CLD:
		return wrap(cld)
	case CLI:
		return wrap(cli)
	case CLV:
		return wrap(clv)
	case CMP:
		return wrap(cmp)
	case CPX:
		return wrap(cpx)
	case CPY:
		return wrap(cpy)
	case DEC:
		return wrap(dec)
	case DEX:
		return wrap(dex)
	case DEY:
		return wrap(dey)
	case EOR:
		return wrap(eor)
	case INC:
		return wrap(inc)
	case INX:
		return wrap(inx)
	case INY:
		return wrap(iny)
	case JMP:
		return wrap(jmp)
	case JSR:
		return wrap(jsr)
	case LDA:
		return wrap(lda)
	case LDX:
		return wrap(ldx)
	case LDY:
		return wrap(ldy)
	case LSR:
		return wrap(func(bus *CpuBus, reg *CpuRegister, address uint16) { lsr(opcode, bus, reg, address) })
	case NOP:
		return wrap(nop)
	case ORA:
		return wrap(ora)
	case PHA:
		return wrap(pha)
	case PHP:
		return wrap(php)
	case PLA:
		return wrap(pla)
	case PLP:
		return wrap(plp)
	case ROL:
		return wrap(func(bus *CpuBus, reg *CpuRegister, address uint16) { rol(opcode, bus, reg, address) })
	case ROR:
		return wrap(func(bus *CpuBus, reg *CpuRegister, address uint16) { ror(opcode, bus, reg, address) })
	case RTI:
		return wrap(rti)
	case RTS:
		return wrap(rts)
	case SBC:
		return wrap(sbc)
	case SEC:
		return wrap(sec)
	case SED:
		return wrap(sed)
	case SEI:
		return wrap(sei)
	case STA:
		return wrap(sta)
	case STX:
		return wrap(stx)
	case STY:
		return wrap(sty)
	case TAX:
		return wrap(tax)
	case TAY:
		return wrap(tay)
	case TSX:
		return wrap(tsx)
	case TXA:
		return wrap(txa)
	case TXS:
		return wrap(txs)
	case TYA:
		return wrap(tya)
	default:
		panic("findOpration")
	}
}

// utils
func checkMSB(target uint8) bool {
	return target&0x80 != 0
}

func checkLSB(target uint8) bool {
	return target&1 == 1
}

func boolToUInt8(b bool) uint8 {
	if b {
		return 1
	} else {
		return 0
	}
}

func push(bus *CpuBus, reg *CpuRegister, target byte) {
	bus.Write(0x100|uint16(reg.SP), target)
	reg.SP--
}

func pull(bus *CpuBus, reg *CpuRegister, address uint16) byte {
	reg.SP++
	return bus.Read(0x100 | uint16(reg.SP))
}

func updateZNStatus(bus *CpuBus, reg *CpuRegister, address uint16, target byte) {
	reg.Z = target == 0
	reg.N = checkMSB(target)
}

func compareAndUpdateCZNStatus(bus *CpuBus, reg *CpuRegister, address uint16, base uint8, target uint8) {
	reg.C = uint8(base) >= uint8(target)
	updateZNStatus(bus, reg, address, uint8(base)-uint8(target))
}

// instructions
// http://obelisk.me.uk/6502/reference.html
// http://taotao54321.hatenablog.com/entry/2017/04/09/151355
func adc(bus *CpuBus, reg *CpuRegister, address uint16) {
	A := reg.A
	M := bus.Read(address)
	C := boolToUInt8(reg.C)
	sum16 := uint16(A) + uint16(M) + uint16(C)
	sum := uint8(sum16)

	reg.A = sum
	reg.C = sum16 > 0xff
	reg.V = (A^M)&0x80 == 0 && (A^sum)&0x80 != 0
	updateZNStatus(bus, reg, address, sum)

}

func and(bus *CpuBus, reg *CpuRegister, address uint16) {
	M := bus.Read(address)
	reg.A &= M

	updateZNStatus(bus, reg, address, reg.A)

}

func asl(opcode Opcode, bus *CpuBus, reg *CpuRegister, address uint16) {
	if opcode.Mode == ModeAccumulator {
		reg.C = checkMSB(reg.A)
		reg.A <<= 1
		updateZNStatus(bus, reg, address, reg.A)
	} else {
		M := bus.Read(address)
		reg.C = checkMSB(M)
		M <<= 1
		bus.Write(address, M)
		updateZNStatus(bus, reg, address, M)
	}
}

func bcc(bus *CpuBus, reg *CpuRegister, address uint16) bool {
	isSuccess := !reg.C
	if isSuccess {
		reg.PC = address
	}
	return isSuccess
}

func bcs(bus *CpuBus, reg *CpuRegister, address uint16) bool {
	isSuccess := reg.C
	if isSuccess {
		reg.PC = address
	}
	return isSuccess
}

func beq(bus *CpuBus, reg *CpuRegister, address uint16) bool {
	isSuccess := reg.Z
	if isSuccess {
		reg.PC = address
	}
	return isSuccess
}

func bit(bus *CpuBus, reg *CpuRegister, address uint16) {
	M := bus.Read(address)

	reg.V = M&0x40 != 0
	reg.N = checkMSB(M)
	reg.Z = reg.A&M == 0
}

func bmi(bus *CpuBus, reg *CpuRegister, address uint16) bool {
	isSuccess := reg.N
	if isSuccess {
		reg.PC = address
	}
	return isSuccess
}

func bne(bus *CpuBus, reg *CpuRegister, address uint16) bool {
	isSuccess := !reg.Z
	if isSuccess {
		reg.PC = address
	}
	return isSuccess
}

func bpl(bus *CpuBus, reg *CpuRegister, address uint16) bool {
	isSuccess := !reg.N
	if isSuccess {
		reg.PC = address
	}
	return isSuccess
}

func brk(bus *CpuBus, reg *CpuRegister, address uint16) {

}

func bvc(bus *CpuBus, reg *CpuRegister, address uint16) bool {
	isSuccess := !reg.V
	if isSuccess {
		reg.PC = address
	}
	return isSuccess
}

func bvs(bus *CpuBus, reg *CpuRegister, address uint16) bool {
	isSuccess := reg.V
	if isSuccess {
		reg.PC = address
	}
	return isSuccess
}

func clc(bus *CpuBus, reg *CpuRegister, address uint16) {
	reg.C = false
}

func cld(bus *CpuBus, reg *CpuRegister, address uint16) {
	reg.D = false
}

func cli(bus *CpuBus, reg *CpuRegister, address uint16) {
	reg.I = false
}

func clv(bus *CpuBus, reg *CpuRegister, address uint16) {
	reg.V = false
}

func cmp(bus *CpuBus, reg *CpuRegister, address uint16) {
	M := bus.Read(address)
	compareAndUpdateCZNStatus(bus, reg, address, reg.A, M)
}

func cpx(bus *CpuBus, reg *CpuRegister, address uint16) {
	M := bus.Read(address)
	compareAndUpdateCZNStatus(bus, reg, address, reg.X, M)
}

func cpy(bus *CpuBus, reg *CpuRegister, address uint16) {
	M := bus.Read(address)
	compareAndUpdateCZNStatus(bus, reg, address, reg.Y, M)
}

func dec(bus *CpuBus, reg *CpuRegister, address uint16) {
	M := bus.Read(address)
	bus.Write(address, M-1)
	updateZNStatus(bus, reg, address, M-1)
}

func dex(bus *CpuBus, reg *CpuRegister, address uint16) {
	reg.X = reg.X - 1
	updateZNStatus(bus, reg, address, reg.X)
}

func dey(bus *CpuBus, reg *CpuRegister, address uint16) {
	reg.Y = reg.Y - 1
	updateZNStatus(bus, reg, address, reg.Y)
}

func eor(bus *CpuBus, reg *CpuRegister, address uint16) {
	M := bus.Read(address)
	reg.A ^= M
	updateZNStatus(bus, reg, address, reg.A)
}

func inc(bus *CpuBus, reg *CpuRegister, address uint16) {
	M := bus.Read(address)
	bus.Write(address, M+1)
	updateZNStatus(bus, reg, address, M+1)
}

func inx(bus *CpuBus, reg *CpuRegister, address uint16) {
	reg.X = reg.X + 1
	updateZNStatus(bus, reg, address, reg.X)
}

func iny(bus *CpuBus, reg *CpuRegister, address uint16) {
	reg.Y = reg.Y + 1
	updateZNStatus(bus, reg, address, reg.Y)
}

func jmp(bus *CpuBus, reg *CpuRegister, address uint16) {
	reg.PC = address
}

func jsr(bus *CpuBus, reg *CpuRegister, address uint16) {
	pc := reg.PC - 1
	push(bus, reg, byte(pc>>8))
	push(bus, reg, byte(pc&0xFF))

	reg.PC = address
}

func lda(bus *CpuBus, reg *CpuRegister, address uint16) {
	M := bus.Read(address)
	reg.A = M
	updateZNStatus(bus, reg, address, reg.A)
}

func ldx(bus *CpuBus, reg *CpuRegister, address uint16) {
	M := bus.Read(address)
	reg.X = M
	updateZNStatus(bus, reg, address, reg.X)
}

func ldy(bus *CpuBus, reg *CpuRegister, address uint16) {
	M := bus.Read(address)
	reg.Y = M
	updateZNStatus(bus, reg, address, reg.Y)
}

func lsr(opcode Opcode, bus *CpuBus, reg *CpuRegister, address uint16) {
	if opcode.Mode == ModeAccumulator {
		reg.C = checkLSB(reg.A)
		reg.A >>= 1
		updateZNStatus(bus, reg, address, reg.A)
	} else {
		M := bus.Read(address)
		reg.C = checkLSB(M)
		M >>= 1
		bus.Write(address, M)
		updateZNStatus(bus, reg, address, M)
	}
}

func nop(bus *CpuBus, reg *CpuRegister, address uint16) {
	// nop
}

func ora(bus *CpuBus, reg *CpuRegister, address uint16) {
	M := bus.Read(address)
	reg.A |= M
	updateZNStatus(bus, reg, address, reg.A)
}

func pha(bus *CpuBus, reg *CpuRegister, address uint16) {
	push(bus, reg, reg.A)
}

func php(bus *CpuBus, reg *CpuRegister, address uint16) {
	// TODO: need source
	push(bus, reg, reg.processorStatus()|0x10)
}

func pla(bus *CpuBus, reg *CpuRegister, address uint16) {
	pull := pull(bus, reg, address)
	reg.A = pull
	updateZNStatus(bus, reg, address, reg.A)
}

func plp(bus *CpuBus, reg *CpuRegister, address uint16) {
	pull := pull(bus, reg, address)
	// TODO: need source
	reg.SetProcessorStatus(pull&0xEF | 0x20)
}

func rol(opcode Opcode, bus *CpuBus, reg *CpuRegister, address uint16) {
	if opcode.Mode == ModeAccumulator {
		res := uint8(reg.A<<1) | boolToUInt8(reg.C)
		reg.C = checkMSB(reg.A)
		reg.A = res
		updateZNStatus(bus, reg, address, reg.A)
	} else {
		M := bus.Read(address)
		res := uint8(M<<1) | boolToUInt8(reg.C)
		reg.C = checkMSB(M)
		bus.Write(address, res)
		updateZNStatus(bus, reg, address, res)
	}
}

func ror(opcode Opcode, bus *CpuBus, reg *CpuRegister, address uint16) {
	if opcode.Mode == ModeAccumulator {
		res := uint8(reg.A>>1) | boolToUInt8(reg.C)<<7
		reg.C = checkLSB(reg.A)
		reg.A = res
		updateZNStatus(bus, reg, address, reg.A)
	} else {
		M := bus.Read(address)
		res := uint8(M>>1) | boolToUInt8(reg.C)<<7
		reg.C = checkLSB(M)
		bus.Write(address, res)
		updateZNStatus(bus, reg, address, res)
	}
}

func rti(bus *CpuBus, reg *CpuRegister, address uint16) {
	pulled := pull(bus, reg, address)
	reg.SetProcessorStatus(pulled&0xEF | 0x20)

	low := pull(bus, reg, address)
	high := pull(bus, reg, address)

	reg.PC = uint16(high)<<8 | uint16(low)
}

func rts(bus *CpuBus, reg *CpuRegister, address uint16) {
	low := pull(bus, reg, address)
	high := pull(bus, reg, address)
	reg.PC = (uint16(high)<<8 | uint16(low)) + 1
}

func sbc(bus *CpuBus, reg *CpuRegister, address uint16) {
	A := reg.A
	M := bus.Read(address)
	C := boolToUInt8(reg.C)
	sum16 := int16(A) - int16(M) - (1 - int16(C))
	sum := uint8(sum16)

	reg.A = sum
	reg.C = sum16 >= 0
	reg.V = (A^M)&0x80 != 0 && (A^sum)&0x80 != 0
	updateZNStatus(bus, reg, address, sum)
}

func sec(bus *CpuBus, reg *CpuRegister, address uint16) {
	reg.C = true
}

func sed(bus *CpuBus, reg *CpuRegister, address uint16) {
	reg.D = true
}

func sei(bus *CpuBus, reg *CpuRegister, address uint16) {
	reg.I = true
}

func sta(bus *CpuBus, reg *CpuRegister, address uint16) {
	bus.Write(address, reg.A)
}

func stx(bus *CpuBus, reg *CpuRegister, address uint16) {
	bus.Write(address, reg.X)
}

func sty(bus *CpuBus, reg *CpuRegister, address uint16) {
	bus.Write(address, reg.Y)
}

func tax(bus *CpuBus, reg *CpuRegister, address uint16) {
	reg.X = reg.A
	updateZNStatus(bus, reg, address, reg.X)
}

func tay(bus *CpuBus, reg *CpuRegister, address uint16) {
	reg.Y = reg.A
	updateZNStatus(bus, reg, address, reg.Y)
}

func tsx(bus *CpuBus, reg *CpuRegister, address uint16) {
	reg.X = reg.SP
	updateZNStatus(bus, reg, address, reg.X)
}

func txa(bus *CpuBus, reg *CpuRegister, address uint16) {
	reg.A = reg.X
	updateZNStatus(bus, reg, address, reg.A)
}

func txs(bus *CpuBus, reg *CpuRegister, address uint16) {
	reg.SP = reg.X
}

func tya(bus *CpuBus, reg *CpuRegister, address uint16) {
	reg.A = reg.Y
	updateZNStatus(bus, reg, address, reg.A)
}
