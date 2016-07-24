package main

import (
	"github.com/fohristiwhirl/wavmaker"
	"math"
)

const FRAMES = 88200
const FREQ = 440
const TWO_PI = 6.28318530718

func main() {
	wav := wavmaker.NewWAV(FRAMES)

	for n := uint32(0); n < FRAMES; n++ {
		raw_val := 32767 * math.Sin((float64(n) / (wavmaker.PREFERRED_FREQ / TWO_PI)) * FREQ)
		val := int16(raw_val)
		wav.Set(n, val, val)
	}

	wav.Save("sine.wav")
}
