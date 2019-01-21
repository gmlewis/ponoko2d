package main

import (
	"flag"
	"log"
	"math"
	"os"

	"github.com/gmlewis/ponoko2d/float"
)

var (
	step  = flag.Float64("step", 0.1, "Angle step for greating spiral in radians")
	n     = flag.Int("n", 100, "Number of turns in each coil")
	gap   = flag.Float64("gap", 0.1, "Gap between traces in mm")
	trace = flag.Float64("trace", 0.1, "Trace width in mm")

	canvas        = ponoko2d.New(os.Stdout)
	width, height float64
	start, end    float64
)

func background(v int) { canvas.Rect(0, 0, width, height, canvas.RGB(v, v, v)) }

func main() {
	flag.Parse()

	start = 2 * math.Pi
	end = start + float64(*n)*2*math.Pi
	x, _ := genPt(end, *trace, 0)
	x = 2 * math.Abs(x)
	xl, _ := genPt(end, *trace, math.Pi)
	xl = 2 * math.Abs(xl)
	if xl > x {
		x = xl
	}
	width = x
	height = width
	log.Printf("n=%v: (%.2f,%.2f)", *n, width, height)

	canvas.Start(width, height)
	background(255)

	genSpiral(0)
	genSpiral(math.Pi)

	canvas.Grid(0, 0, width, height, 10, "stroke:black;opacity:0.1")
	canvas.End()
}

func genPt(angle, halfTW, offset float64) (float64, float64) {
	r := (angle + *trace + *gap) / (4 * math.Pi)
	x := width/2 + (r+halfTW)*math.Cos(angle+offset)
	y := height/2 - (r+halfTW)*math.Sin(angle+offset)
	return x, y
}

func genSpiral(offset float64) {
	halfTW := 0.5 * *trace
	var xs, ys []float64
	steps := int64(math.Ceil((end - start) / *step))
	for i := int64(0); i < steps; i++ {
		angle := start + *step*float64(i)
		x, y := genPt(angle, halfTW, offset)
		xs = append(xs, x)
		ys = append(ys, y)
	}
	x, y := genPt(end, halfTW, offset)
	xs = append(xs, x)
	ys = append(ys, y)
	x, y = genPt(end, -halfTW, offset)
	xs = append(xs, x)
	ys = append(ys, y)
	for i := steps - 1; i >= 0; i-- {
		angle := start + *step*float64(i)
		x, y := genPt(angle, -halfTW, offset)
		xs = append(xs, x)
		ys = append(ys, y)
	}
	xs = append(xs, xs[0])
	ys = append(ys, ys[0])
	canvas.Polygon(xs, ys, "stroke:none; fill:black")
}
