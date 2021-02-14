package nes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRom(t *testing.T) {
	c := ParseRom("./roms/full_palette.nes")
	cSize := len(c.characterRom)
	pSize := len(c.programRom)

	assert.Equal(t, 8192, cSize)
	assert.Equal(t, 32768, pSize)

	assert.Equal(t, uint8(0xff), c.characterRom[len(c.characterRom)-1])
	assert.Equal(t, uint8(0x82), c.programRom[len(c.programRom)-1])
}

func TestParseEmptyCharacterRom(t *testing.T) {
	c := ParseRom("./roms/official_only.nes")
	cSize := len(c.characterRom)
	pSize := len(c.programRom)

	assert.Equal(t, 0, cSize)
	assert.Equal(t, 262144, pSize)

	assert.Equal(t, uint8(0xe2), c.programRom[len(c.programRom)-1])
}
