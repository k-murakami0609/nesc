package nes

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
)

func GenerateImage(selector int) *bytes.Buffer {
	width := 200
	height := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	val1 := uint8(0)
	val2 := uint8(0)
	val3 := uint8(0)

	if selector == 1 {
		val1 = uint8(200)
	}
	if selector == 2 {
		val2 = uint8(200)
	}
	if selector == 3 {
		val3 = uint8(200)
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			cyan := color.RGBA{val1, val2, val3, 0xff}
			img.Set(x, y, cyan)
		}
	}

	b := new(bytes.Buffer)
	err := jpeg.Encode(b, img, &jpeg.Options{})
	if err != nil {
		fmt.Println("encode")
		panic(err)
	}

	return b
}
