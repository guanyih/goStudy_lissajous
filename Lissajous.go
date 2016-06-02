// Lissajous
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.RGBA{0, 128, 0, 1}}

const (
	whiteIndex = 0 //first color in the palette
	blackIndex = 1 //next color in the palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     //number of complex x oscillator revolutions
		res     = 0.001 //angular resolution
		size    = 100   //image canvas covers [-size, size]
		nframes = 64    //number of animated frames
		delay   = 8     //delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 //relative requency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference
	for i := 0; i < nframes; i++ {
		rec := image.Rect(0, 0, size*2+1, 2*size+1)
		img := image.NewPaletted(rec, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(freq*t + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(size*y+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
