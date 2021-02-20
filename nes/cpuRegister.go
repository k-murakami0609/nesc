package nes

type CpuRegister struct {
	PC uint16
	SP byte
	A  byte
	X  byte
	Y  byte
	C  bool
	Z  bool
	I  bool
	D  bool
	B  bool
	R  bool
	V  bool
	N  bool
}

func (reg *CpuRegister) SetProcessorStatus(s uint8) {
	reg.C = (s >> 0 & 1) == 1
	reg.Z = (s >> 1 & 1) == 1
	reg.I = (s >> 2 & 1) == 1
	reg.D = (s >> 3 & 1) == 1
	reg.B = (s >> 4 & 1) == 1
	reg.R = (s >> 5 & 1) == 1
	reg.V = (s >> 6 & 1) == 1
	reg.N = (s >> 7 & 1) == 1
}

// 7, 6, 5, 4, 3, 2, 1, 0
// N, V, R, B, D, I, Z, C
func (reg *CpuRegister) processorStatus() uint8 {
	var result uint8 = 0
	bits := []bool{
		reg.C,
		reg.Z,
		reg.I,
		reg.D,
		reg.B,
		reg.R,
		reg.V,
		reg.N,
	}

	for i := 0; i < 8; i++ {
		if bits[i] {
			result |= 1 << i
		}
	}

	return result
}
