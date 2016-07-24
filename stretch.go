// Stretch - stretch or squash the given WAV file by the multiplier.
// This changes its pitch, of course.

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fohristiwhirl/wavmaker"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <filename> <multiplier>\n", os.Args[0])
		return
	}

	wav, err := wavmaker.Load(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	multiplier, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	new_frames_f := float64(wav.FrameCount()) * multiplier

	wav = wav.Stretched(uint32(new_frames_f))
	wav.Save(os.Args[1] + ".stretched.wav")
}
