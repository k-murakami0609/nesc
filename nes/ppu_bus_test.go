package nes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameTables(t *testing.T) {
	console := NewConsole("./roms/sample1.nes")

	for i := 0; i < 100; i++ {
		console.Run()
	}

	assert.Equal(t, console.PpuBus.nameTables[457], uint8(72))
	assert.Equal(t, console.PpuBus.nameTables[458], uint8(69))
	assert.Equal(t, console.PpuBus.nameTables[459], uint8(76))
	assert.Equal(t, console.PpuBus.nameTables[460], uint8(76))
	assert.Equal(t, console.PpuBus.nameTables[461], uint8(79))
	assert.Equal(t, console.PpuBus.nameTables[462], uint8(44))
	assert.Equal(t, console.PpuBus.nameTables[463], uint8(32))
	assert.Equal(t, console.PpuBus.nameTables[464], uint8(87))
	assert.Equal(t, console.PpuBus.nameTables[465], uint8(79))
	assert.Equal(t, console.PpuBus.nameTables[466], uint8(82))
	assert.Equal(t, console.PpuBus.nameTables[467], uint8(76))
	assert.Equal(t, console.PpuBus.nameTables[468], uint8(68))
	assert.Equal(t, console.PpuBus.nameTables[469], uint8(33))
}
