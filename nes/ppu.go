package nes

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
)

type PPU struct {
	console  *Console
	Register PpuRegister

	cycle int
	line  int

	backGrounds [30][32]struct {
		data [8][8]int
	}
}

func NewPPU(console *Console) *PPU {
	return &PPU{console: console}
}

func (ppu *PPU) Step(cycle int) bool {
	cycle = ppu.cycle + cycle
	if cycle < 341 {
		ppu.cycle = cycle
		return false
	}

	ppu.cycle = 0
	ppu.line += 1

	// visible scanline
	if ppu.line <= 240 && ppu.line%8 == 0 {
		ppu.createBackGrounds()
		// TODO: need palete
		return ppu.line == 240
	}

	// vertical blanking line
	if ppu.line == 241 {
		return false
	}

	// pre-render scanline
	if ppu.line >= 262 {
		ppu.line = 0
		return false
	}

	return false
}

func (ppu *PPU) createBackGrounds() {
	ty := ((ppu.line / 8) - 1)
	for tx := 0; tx < 32; tx++ {
		offset := ty * 32
		characterIndex := ppu.console.PpuBus.nameTables[offset+tx]
		c := ppu.createCharacter(int(characterIndex))
		ppu.backGrounds[ty][tx].data = c
	}
}

func (ppu *PPU) createCharacter(index int) [8][8]int {
	var result [8][8]int
	character1, character2 := ppu.console.PpuBus.ReadCharacter(index)
	for ty := 0; ty < 8; ty++ {
		for tx := 0; tx < 8; tx++ {
			a := int(character1[ty]) >> (7 - tx) & 1
			b := int(character2[ty]) >> (7 - tx) & 1
			result[ty][tx] = a<<1 | b
		}
	}

	return result
}

func (ppu *PPU) generateImage() *bytes.Buffer {
	width := 256
	height := 240

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x := 0; x < 32; x++ {
		for y := 0; y < 30; y++ {
			back := ppu.backGrounds[y][x]

			for tx := 0; tx < 8; tx++ {
				for ty := 0; ty < 8; ty++ {
					if back.data[ty][tx] == 0 {
						cyan := color.RGBA{0, 0, 0, 0xff}
						img.Set(x*8+tx, y*8+ty, cyan)
					} else {
						cyan := color.RGBA{0xff, 0xff, 0xff, 0xff}
						img.Set(x*8+tx, y*8+ty, cyan)
					}
				}
			}
		}
	}

	b := new(bytes.Buffer)
	err := jpeg.Encode(b, img, &jpeg.Options{Quality: 100})
	if err != nil {
		fmt.Println("encode")
		panic(err)
	}

	return b
}
