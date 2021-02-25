package nes

func BoolArrayToUint8(bits [8]bool) uint8 {
	var result uint8 = 0
	for i := 0; i < 8; i++ {
		if bits[i] {
			result |= 1 << i
		}
	}

	return result
}
