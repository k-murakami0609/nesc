package nes

import "fmt"

type CpuBus struct {
	console *Console
}

func NewCpuBus(console *Console) *CpuBus {
	return &CpuBus{console: console}
}

// http://nesdev.com/NESDoc.pdf#page=11
func (cpuBus *CpuBus) Write(address uint16, value byte) {
	switch {
	case address < 0x2000:
		cpuBus.console.RAM[address%0x0800] = value
		return
	case address < 0x4000:
		cpuBus.console.PpuBus.Write(address, value)
		return
	case address < 0x4013:
		// ???
		return
	case address < 0x4014:
		cpuBus.console.PpuBus.Write(address, value)
		return
	case address >= 0x8000:
		cpuBus.console.Cartridge.WriteProgramRom(address-0x8000, value)
		return
	}

	fmt.Printf("cant write address: %04X", address)
	panic("memory write")
}

func (cpuBus *CpuBus) Read(address uint16) byte {
	switch {
	case address < 0x2000:
		return cpuBus.console.RAM[address%0x0800]
	case address < 0x4000:
		return cpuBus.console.PpuBus.Read(address)
	case address >= 0x8000:
		return cpuBus.console.Cartridge.ReadProgramRom(address - 0x8000)
	}

	fmt.Printf("cant read address: %04X", address)
	panic("memory read")
}

func (cpuBus *CpuBus) Read16(address uint16) uint16 {
	low := uint16(cpuBus.Read(address))
	high := uint16(cpuBus.Read(address + 1))

	return high<<8 | low
}

// https://github.com/fogleman/nes/blob/8c4b9cf54c354137c37e8ae17caf4b1b1405313b/nes/cpu.go#L318
func (cpuBus *CpuBus) Read16bug(address uint16) uint16 {
	a := address
	b := (a & 0xFF00) | uint16(byte(a)+1)
	low := uint16(cpuBus.Read(a))
	high := uint16(cpuBus.Read(b))
	return uint16(high)<<8 | uint16(low)
}
