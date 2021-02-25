package nes

type Console struct {
	CPU       *CPU
	Cartridge *Cartridge
	RAM       [0x0800]byte
	CpuBus    *CpuBus
	PPU       *PPU
	PpuBus    *PpuBus
}

func NewConsole(path string) *Console {
	console := Console{
		CPU:       nil,
		Cartridge: ParseRom(path),
	}
	console.CpuBus = NewCpuBus(&console)
	console.CPU = NewCPU(console.CpuBus)
	console.PPU = NewPPU(&console)
	return &console
}
