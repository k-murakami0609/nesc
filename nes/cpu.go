package nes

import (
	"fmt"
)

type CPU struct {
	bus      *CpuBus
	Register CpuRegister
	Cycles   int
}

func NewCPU(bus *CpuBus) *CPU {
	cpu := CPU{bus: bus}
	cpu.Reset()

	return &cpu
}

func (cpu *CPU) Reset() {
	cpu.Register.PC = cpu.bus.Read16(0xFFFC)
	cpu.Register.SP = 0xFD
	cpu.Register.SetProcessorStatus(0x24)
}

func (cpu *CPU) Step() int {
	opcode := cpu.nextOpcode()

	cycles := ExecuteOpration(cpu.bus, &cpu.Register, opcode)
	cpu.Cycles += cycles

	return cycles
}

func (cpu *CPU) nextOpcode() Opcode {
	return Opcodes[cpu.bus.Read(cpu.Register.PC)]
}

func (cpu *CPU) Debug() string {
	opcode := cpu.nextOpcode()
	w0 := fmt.Sprintf("%02X", cpu.bus.Read(cpu.Register.PC+0))
	w1 := fmt.Sprintf("%02X", cpu.bus.Read(cpu.Register.PC+1))
	w2 := fmt.Sprintf("%02X", cpu.bus.Read(cpu.Register.PC+2))
	r := cpu.Register

	if opcode.Size < 2 {
		w1 = "  "
	}
	if opcode.Size < 3 {
		w2 = "  "
	}
	return fmt.Sprintf("%4X  %s %s %s  %s %28s"+
		"A:%02X X:%02X Y:%02X P:%02X SP:%02X CYC:%3d\n",
		r.PC, w0, w1, w2, opcode.Name, "",
		r.A, r.X, r.Y, r.processorStatus(), r.SP, (cpu.Cycles*3)%341)

	// address, _ := Address(cpu.bus, &cpu.Register, opcode)
	// return fmt.Sprintf("%4X  %s %s %s  %s %28s"+
	// 	"A:%02X X:%02X Y:%02X P:%02X SP:%02X CYC:%3d addr:%04X\n",
	// 	r.PC, w0, w1, w2, opcode.Name, "",
	// 	r.A, r.X, r.Y, r.processorStatus(), r.SP, (cpu.Cycles*3)%341, address)
}
