package nes

type Console struct {
	CPU    *CPU
	CpuBus *CpuBus
	PPU    *PPU
	PpuBus *PpuBus
}

func NewConsole(path string) *Console {
	cartridge := ParseRom(path)

	console := Console{
		CPU: nil,
	}

	console.CpuBus = NewCpuBus(&console, cartridge)
	console.PpuBus = NewPpuBus(&console, cartridge)

	console.CPU = NewCPU(console.CpuBus)
	console.PPU = NewPPU(&console)
	return &console
}
