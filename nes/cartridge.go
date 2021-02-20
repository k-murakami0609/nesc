package nes

import (
	"encoding/binary"
	"os"
)

const headerSize = 16

// TODO: ROMによっては512bytes
const trainerSize = 0

const programRomBaseSize = 16384
const characterRomBaseSize = 8192

type header struct {
	Name    uint32 // 4 * 8 = 32
	PrgSize byte
	ChrSize byte
	// TODO: Flag6 ~ 10, unused 11 ~ 15
	_ [10]byte
}

type Cartridge struct {
	programRom   []byte
	characterRom []byte
	// TODO:
	// mapper
	// mirror
}

func (c Cartridge) Read(address uint16) byte {
	index := int(address) % len(c.programRom)
	return c.programRom[index]
}

func (c Cartridge) Write(address uint16, value byte) {
	c.programRom[address] = value
}

func ParseRom(romPath string) *Cartridge {
	file, err := os.Open(romPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	h := parseINesHeader(file)

	programRom := parseProgramRom(file, h)
	characterRom := parseCharacterRom(file, h)

	return &Cartridge{programRom: programRom, characterRom: characterRom}
}

func parseProgramRom(file *os.File, h header) []byte {
	offsetOfProgramRom, programRomSize := calcProgramRomOffsetAndSize(h)
	programRom := make([]byte, programRomSize)
	_, err := file.ReadAt(programRom, offsetOfProgramRom)
	if err != nil {
		panic(err)
	}

	return programRom
}

func parseCharacterRom(file *os.File, h header) []byte {
	offsetOfProgramRom, programRomSize := calcProgramRomOffsetAndSize(h)
	offsetOfCharacterRom := offsetOfProgramRom + programRomSize
	characterRomSize := characterRomBaseSize * int64(h.ChrSize)

	characterRom := make([]byte, characterRomSize)
	_, err := file.ReadAt(characterRom, offsetOfCharacterRom)
	if err != nil {
		panic(err)
	}

	return characterRom
}

func calcProgramRomOffsetAndSize(h header) (int64, int64) {
	offsetOfProgramRom := int64(headerSize + trainerSize)
	programRomSize := programRomBaseSize * int64(h.PrgSize)

	return offsetOfProgramRom, programRomSize
}

func parseINesHeader(file *os.File) header {
	h := header{}
	if err := binary.Read(file, binary.LittleEndian, &h); err != nil {
		panic(err)
	}

	return h
}
