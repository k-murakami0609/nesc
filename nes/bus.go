package nes

import "fmt"

type CpuBus struct {
	console *Console
}

func NewCpuBus(console *Console) *CpuBus {
	return &CpuBus{console: console}
}

func (cpuBus *CpuBus) Read16(address uint16) uint16 {
	low := uint16(cpuBus.Read(address))
	high := uint16(cpuBus.Read(address + 1))

	return high<<8 | low
}

func (cpuBus *CpuBus) Write(address uint16, value byte) {
	switch {
	case address < 0x2000:
		cpuBus.console.RAM[address%0x0800] = value
		return
	case address >= 0x8000:
		cpuBus.console.Cartridge.Write(address-0x8000, value)
		return
	}

	fmt.Printf("cant write address: %04X", address)
	panic("memory write")
}

func (cpuBus *CpuBus) Read(address uint16) byte {
	switch {
	case address < 0x2000:
		return cpuBus.console.RAM[address%0x0800]
	case address >= 0x8000:
		return cpuBus.console.Cartridge.Read(address - 0x8000)
	}

	fmt.Printf("cant read address: %04X", address)
	panic("memory read")
}
