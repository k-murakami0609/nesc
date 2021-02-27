package nes

// TODO: rethink name
type PpuBus struct {
	console   *Console
	cartridge *Cartridge

	nameTables       [0x1F00]byte
	backgroudPalette [0x10]byte
	spritePalette    [0x10]byte

	spriteRAM [0x100]byte
}

func NewPpuBus(console *Console, cartridge *Cartridge) *PpuBus {
	return &PpuBus{console: console, cartridge: cartridge}
}

// http://nesdev.com/NESDoc.pdf#page=34
func (bus *PpuBus) Read(address uint16) byte {
	switch address {
	case 0x2002:
		return bus.console.PPU.Register.ReadStatus()
	case 0x2004:
		return bus.readSpriteData()
	case 0x2007:
		return bus.readData()
	}

	panic("ppu read")
}

// http://nesdev.com/NESDoc.pdf#page=34
func (bus *PpuBus) Write(address uint16, value byte) {
	switch address {
	case 0x2000:
		bus.console.PPU.Register.SetCtrl(value)
		return
	case 0x2001:
		bus.console.PPU.Register.SetMask(value)
		return
	case 0x2003:
		bus.console.PPU.Register.SpriteAddress = value
		return
	case 0x2004:
		bus.writeSpriteData(value)
		return
	case 0x2005:
		bus.writeScroll(value)
		return
	case 0x2006:
		bus.writePpuAddress(value)
		return
	case 0x2007:
		bus.writeData(value)
		return
	case 0x4014:
		bus.transferDMA(value)
		return
	}
}

func (bus *PpuBus) ReadCharacter(index int) ([8]byte, [8]byte) {
	var character1 [8]byte
	var character2 [8]byte
	offset := index * 16
	for i := 0; i < 8; i++ {
		character1[i] = bus.cartridge.ReadCharacterRom(uint16(offset + i))
	}
	for i := 8; i < 16; i++ {
		character2[i-8] = bus.cartridge.ReadCharacterRom(uint16(offset + i))
	}
	return character1, character2
}

func (bus *PpuBus) readData() byte {
	addr := bus.console.PPU.Register.PpuAddress
	bus.console.PPU.Register.PpuAddress += 1

	switch {
	case addr < 0x2000:
		// TODO: confirm 0x1000?
		return bus.cartridge.ReadCharacterRom(addr % 0x1000)
	case addr < 0x3F00:
		index := addr - 0x2000
		return bus.nameTables[index]
	case addr < 0x3F10:
		index := addr - 0x3F00
		return bus.backgroudPalette[index]
	case addr < 0x3F20:
		index := addr - 0x3F10
		return bus.spritePalette[index]
	case addr < 0x4000:
		// TODO: mirror
		panic("ppu readData mirror")
	default:
		panic("ppu readData")
	}
}

func (bus *PpuBus) writeData(value byte) {
	addr := bus.console.PPU.Register.PpuAddress
	if bus.console.PPU.Register.Ctrl.I {
		bus.console.PPU.Register.PpuAddress += 32
	} else {
		bus.console.PPU.Register.PpuAddress += 1
	}

	switch {
	case addr < 0x2000:
		// TODO: confirm 0x1000?
		bus.cartridge.WriteCharacterRom(addr%0x1000, value)
		return
	case addr < 0x3F00:
		index := addr - 0x2000
		bus.nameTables[index] = value
		return
	case addr < 0x3F10:
		index := addr - 0x3F00
		bus.backgroudPalette[index] = value
	case addr < 0x3F20:
		index := addr - 0x3F10
		bus.spritePalette[index] = value
	case addr < 0x4000:
		// TODO: mirror
		panic("ppu writeData mirror")
	default:
		panic("ppu writeData mirror")
	}
}

func (bus *PpuBus) readSpriteData() byte {
	return bus.spriteRAM[int(bus.console.PPU.Register.SpriteAddress)]
}

func (bus *PpuBus) writeSpriteData(value byte) {
	bus.spriteRAM[int(bus.console.PPU.Register.SpriteAddress)] = value
	bus.console.PPU.Register.SpriteAddress++
}

func (bus *PpuBus) writeScroll(value byte) {
	bus.console.PPU.Register.Scroll <<= 8
	bus.console.PPU.Register.Scroll |= uint16(value)
}

func (bus *PpuBus) writePpuAddress(value byte) {
	bus.console.PPU.Register.PpuAddress <<= 8
	bus.console.PPU.Register.PpuAddress |= uint16(value)
}

func (bus *PpuBus) transferDMA(value byte) {
	addr := uint16(value) << 8
	for i := addr; i < addr+256; i++ {
		bus.spriteRAM[int(bus.console.PPU.Register.SpriteAddress)] = bus.console.CpuBus.Read(i)
		bus.console.PPU.Register.SpriteAddress++
	}
}
