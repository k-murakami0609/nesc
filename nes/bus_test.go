package nes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead16(t *testing.T) {
	console := NewConsole("./roms/full_palette.nes")
	console.CpuBus.Write(0xFFFC, 0x10)
	console.CpuBus.Write(0xFFFD, 0x20)

	assert.Equal(t, uint16(0x2010), console.CpuBus.Read16(0xFFFC))
}
