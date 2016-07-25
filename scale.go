package main

import (
	"github.com/fohristiwhirl/wavmaker"
)

const (
	C4 = 261.625565
	D4 = 293.664768
	E4 = 329.627557
	F4 = 349.228231
	G4 = 391.995436
	A4 = 440
	B4 = 493.883301
	C5 = 523.251131
)

func main() {

	output := wavmaker.New(44100 * 5)

	c4, err := wavmaker.Load("piano.ff.C4.wav")
	if err != nil {
		panic("load failed")
	}

	d4 := c4.StretchedRelative(C4 / D4)
	e4 := c4.StretchedRelative(C4 / E4)
	f4 := c4.StretchedRelative(C4 / F4)
	g4 := c4.StretchedRelative(C4 / G4)
	a4 := c4.StretchedRelative(C4 / A4)
	b4 := c4.StretchedRelative(C4 / B4)
	c5 := c4.StretchedRelative(C4 / C5)

	output.Add(22050 * 0, c4, 0, 44100 * 4)
	output.Add(22050 * 1, d4, 0, 44100 * 4)
	output.Add(22050 * 2, e4, 0, 44100 * 4)
	output.Add(22050 * 3, f4, 0, 44100 * 4)
	output.Add(22050 * 4, g4, 0, 44100 * 4)
	output.Add(22050 * 5, a4, 0, 44100 * 4)
	output.Add(22050 * 6, b4, 0, 44100 * 4)
	output.Add(22050 * 7, c5, 0, 44100 * 4)

	output.Save("scale.wav")
}