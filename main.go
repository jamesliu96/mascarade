package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"math/rand"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func fraction(n int) (x, y int) {
	if n == 0 {
		return
	}
	x = int(math.Sqrt(float64(n)))
	for n%x != 0 {
		x--
	}
	y = n / x
	return
}

func randBool() bool {
	return rand.Int63()&(1<<62) == 0
}

func main() {
	switch {
	case len(os.Args) > 1:
		m, err := png.Decode(os.Stdin)
		check(err)
		r := m.Bounds()
		w, h := r.Dx(), r.Dy()
		data := make([]byte, w*h)
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				data[x+y*w] = color.GrayModel.Convert(m.At(x, y)).(color.Gray).Y
			}
		}
		_, err = os.Stdout.Write(data)
		check(err)
	default:
		data, err := io.ReadAll(os.Stdin)
		check(err)
		n := len(data)
		w, h := fraction(n)
		if randBool() {
			w, h = h, w
		}
		m := image.NewGray(image.Rect(0, 0, w, h))
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				m.SetGray(x, y, color.Gray{data[x+y*w]})
			}
		}
		check(png.Encode(os.Stdout, m))
	}
}
