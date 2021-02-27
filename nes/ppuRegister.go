package nes

// http://wiki.nesdev.com/w/index.php/PPU_registers
// http://nesdev.com/NESDoc.pdf#page=34
type PpuRegister struct {
	// $2000
	Ctrl struct {
		LN bool // low Base nametable address
		HN bool // height Base nametable address
		I  bool
		S  bool
		B  bool
		H  bool
		P  bool
		V  bool
	}
	// $2001
	Mask struct {
		GS  bool // Greyscale
		MSB bool // show background in leftmost
		MSS bool // show sprites in leftmost
		SB  bool // show background
		SS  bool // show sprites
		R   bool
		G   bool
		B   bool
	}
	// $2002
	Status struct {
		_ [5]bool
		O bool
		S bool
		V bool
	}
	// $2003 (write * 2)
	SpriteAddress byte
	// $2005
	Scroll uint16
	// $2005 (write * 2)
	PpuAddress uint16
}

func (reg *PpuRegister) SetCtrl(s uint8) {
	reg.Ctrl.LN = (s >> 0 & 1) == 1
	reg.Ctrl.HN = (s >> 1 & 1) == 1
	reg.Ctrl.I = (s >> 2 & 1) == 1
	reg.Ctrl.S = (s >> 3 & 1) == 1
	reg.Ctrl.B = (s >> 4 & 1) == 1
	reg.Ctrl.H = (s >> 5 & 1) == 1
	reg.Ctrl.P = (s >> 6 & 1) == 1
	reg.Ctrl.V = (s >> 7 & 1) == 1
}

func (reg *PpuRegister) SetMask(s uint8) {
	reg.Mask.GS = (s >> 0 & 1) == 1
	reg.Mask.MSB = (s >> 1 & 1) == 1
	reg.Mask.MSS = (s >> 2 & 1) == 1
	reg.Mask.SB = (s >> 3 & 1) == 1
	reg.Mask.SS = (s >> 4 & 1) == 1
	reg.Mask.R = (s >> 5 & 1) == 1
	reg.Mask.G = (s >> 6 & 1) == 1
	reg.Mask.B = (s >> 7 & 1) == 1
}

func (reg *PpuRegister) SetStatus(s uint8) {
	reg.Status.O = (s >> 5 & 1) == 1
	reg.Status.S = (s >> 6 & 1) == 1
	reg.Status.V = (s >> 7 & 1) == 1
}

func (reg *PpuRegister) ReadCtrl() uint8 {
	return BoolArrayToUint8([8]bool{
		reg.Ctrl.LN,
		reg.Ctrl.HN,
		reg.Ctrl.I,
		reg.Ctrl.S,
		reg.Ctrl.B,
		reg.Ctrl.H,
		reg.Ctrl.P,
		reg.Ctrl.V,
	})
}

func (reg *PpuRegister) ReadMask() uint8 {
	return BoolArrayToUint8([8]bool{
		reg.Mask.GS,
		reg.Mask.MSB,
		reg.Mask.MSS,
		reg.Mask.SB,
		reg.Mask.SS,
		reg.Mask.R,
		reg.Mask.G,
		reg.Mask.B,
	})
}

func (reg *PpuRegister) ReadStatus() uint8 {
	return BoolArrayToUint8([8]bool{
		false,
		false,
		false,
		false,
		false,
		reg.Status.O,
		reg.Status.S,
		reg.Status.V,
	})
}
