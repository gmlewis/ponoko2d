package main

import (
	"flag"
	"math"
	"os"

	"github.com/gmlewis/ponoko2d/float"
)

var (
	width  = flag.Float64("width", 500, "Width of svg in mm")
	height = flag.Float64("height", 500, "Height of svg in mm")
	step   = flag.Float64("step", 0.1, "Angle step for greating spiral in radians")
	n      = flag.Float64("n", 10, "Number of turns in each coil")
	trace  = flag.Float64("trace", 0.5, "Trace width in mm")

	canvas = ponoko2d.New(os.Stdout)
)

func background(v int) { canvas.Rect(0, 0, *width, *height, canvas.RGB(v, v, v)) }

func main() {
	canvas.Start(*width, *height)
	background(255)

	genSpiral(0)
	genSpiral(math.Pi)

	canvas.Grid(0, 0, *width, *height, 10, "stroke:black;opacity:0.1")
	canvas.End()
}

func genPt(angle, tw, offset float64) (float64, float64) {
	x := *width/2 + (angle+tw)*math.Cos(angle+offset)
	y := *height/2 - (angle+tw)*math.Sin(angle+offset)
	return x, y
}

func genSpiral(offset float64) {
	start := 2 * math.Pi
	end := start + float64(*n)*2*math.Pi
	var xs, ys []float64
	steps := int64(math.Ceil((end - start) / *step))
	for i := int64(0); i < steps; i++ {
		angle := start + *step*float64(i)
		x, y := genPt(angle, *trace, offset)
		xs = append(xs, x)
		ys = append(ys, y)
	}
	x, y := genPt(end, *trace, offset)
	xs = append(xs, x)
	ys = append(ys, y)
	x, y = genPt(end, -*trace, offset)
	xs = append(xs, x)
	ys = append(ys, y)
	for i := steps - 1; i >= 0; i-- {
		angle := start + *step*float64(i)
		x, y := genPt(angle, -*trace, offset)
		xs = append(xs, x)
		ys = append(ys, y)
	}
	xs = append(xs, xs[0])
	ys = append(ys, ys[0])
	canvas.Polyline(xs, ys, "stroke:black; fill:none; stroke-width:0.1")
}
