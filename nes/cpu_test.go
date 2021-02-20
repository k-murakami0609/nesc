package nes

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetProcessorStatus(t *testing.T) {
	console := NewConsole("./roms/full_palette.nes")
	cpu := console.CPU

	check := func(expteced []bool) {
		status := []bool{
			cpu.Register.C,
			cpu.Register.Z,
			cpu.Register.I,
			cpu.Register.D,
			cpu.Register.B,
			cpu.Register.R,
			cpu.Register.V,
			cpu.Register.N,
		}
		for i := 0; i < len(status); i++ {
			assert.Equal(t, expteced[i], status[i])
		}
	}

	expects := []bool{false, false, false, false, false, false, false, false}
	cpu.Register.SetProcessorStatus(0)
	check(expects)

	expects = []bool{false, false, false, false, true, false, false, false}
	cpu.Register.SetProcessorStatus(uint8(math.Pow(2, 4)))
	check(expects)

	expects = []bool{true, true, true, true, true, true, true, true}
	cpu.Register.SetProcessorStatus(uint8(math.Pow(2, 8)) - 1)
	check(expects)
}

// func TestOprations(t *testing.T) {
// 	console := NewConsole("./roms/nestest.nes")

// 	console.CPU.Register.PC = 0xC000
// 	console.CPU.Register.A = 0
// 	console.CPU.Register.X = 0
// 	console.CPU.Register.Y = 0
// 	console.CPU.Register.SP = 0xFD
// 	console.CPU.Register.SetProcessorStatus(0x24)

// 	file, err := os.Create("../tmp/nestest.log")
// 	if err != nil {
// 		panic(err)
// 	}

// 	for i := 0; i < 5259; i++ {
// 		_, _ = file.WriteString(console.CPU.Debug())
// 		console.CPU.Step()
// 	}
// }
