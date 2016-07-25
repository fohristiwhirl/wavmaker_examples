package main

import (
	"github.com/fohristiwhirl/wavmaker"
	"math"
)

const TWO_PI = 6.28318530718

func make_sine(freq float64, frames uint32) *wavmaker.WAV {
	wav := wavmaker.New(frames)
	for n := uint32(0); n < frames; n++ {
		raw_val := 32767 * math.Sin((float64(n) / (wavmaker.PREFERRED_FREQ / TWO_PI)) * freq)
		val := int16(raw_val)
		wav.Set(n, val, val)
	}
	return wav
}

func main() {
	wav := make_sine(440, 88200)
	wav.Save("sine.wav")
}
