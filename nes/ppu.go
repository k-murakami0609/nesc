package nes

type PPU struct {
	console  *Console
	Register PpuRegister
}

func NewPPU(console *Console) *PPU {
	return &PPU{console: console}
}
