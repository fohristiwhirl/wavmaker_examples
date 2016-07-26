// Note: the Add function in wavmaker is routinely getting updated; this example may fail
// if it's not up to date with the most recent changes.

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

	output := wavmaker.New(44100 * 3)

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

	output.Add(11025 * 0, c4, 0, 44100 * 4, 1.0)
	output.Add(11025 * 1, d4, 0, 44100 * 4, 1.0)
	output.Add(11025 * 2, e4, 0, 44100 * 4, 1.0)
	output.Add(11025 * 3, f4, 0, 44100 * 4, 1.0)
	output.Add(11025 * 4, g4, 0, 44100 * 4, 1.0)
	output.Add(11025 * 5, a4, 0, 44100 * 4, 1.0)
	output.Add(11025 * 6, b4, 0, 44100 * 4, 1.0)
	output.Add(11025 * 7, c5, 0, 44100 * 4, 1.0)

	output.Save("scale.wav")
}
