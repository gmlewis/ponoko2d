// magnets2 is an example of using the ponoko2d library to create a 2D SVG
// design that can be sent to Ponoko.com for fabrication.
//
// To run this program, type:
// go run magnets2.go -magnets > magnets2.svg
//
package main

import (
	"flag"
	"math"
	"os"

	"github.com/gmlewis/ponoko2d/float"
)

var (
	canvas  *ponoko2d.SVG
	magnets = flag.Bool("magnets", false, "Print design with magnet cutouts.")
)

// Actual dimensions as a result of the initial cutting:
// gap = ~ 0.467 in +/- 0.001 in (delta = 0.092 in !!! = 2.34mm) - calculations must have been wrong.
// bearingOD = 1.234 in (delta = 0.109 in!!! = 2.77mm)
// resulting bearingOD cutout = 1.2280 in (delta = 0.103 in = 2.62mm)
// Laser width = 0.15mm
// Magnet length = 2.475in (delta = .475 in!!! WAY OFF!)
// Magnet length cutout = 2.488in - should have been 2.0in !!!  Laser diff=0.013in/2 = 0.33mm/2 = 0.165mm
//
// The actual full circle diameter is 13.75 in !!!  Should have been 11.0 in = 279.4 mm
// Adobe Illustrator said that it was 349.562 mm => 13.7622834646

const (
	in        = ponoko2d.In
	mm        = ponoko2d.Mm
	tolerance = 0.5 * ponoko2d.LaserWidth

	gap        = 3.0 / 8.0 * in
	numMagnets = 16
	degDelta   = 360.0 / numMagnets
	halfDelta  = 0.5 * degDelta
	magnetID   = 4.0 * in
	magnetOD   = 8.0 * in
	rDiff      = gap * numMagnets / (2.0 * math.Pi)
	fullD      = 11.0 * in
	bearingOD  = (1.0 + 1.0/8.0) * in
	zipTieW    = 5.0 * mm
	zipTieH    = 2.0 * mm

	degToRad  = math.Pi / 180.0
	magnetIR  = 0.5 * magnetID
	magnetOR  = 0.5 * magnetOD
	fullR     = 0.5 * fullD
	bearingOR = 0.5 * bearingOD
)

// cutZipTieAtRadius draws a single cut for one zip tie.
func cutZipTieAtRadius(r, deg float64) {
	x1i := (r - 0.5*zipTieW) * math.Cos(deg*degToRad-0.5*zipTieH/r)
	y1i := (r - 0.5*zipTieW) * math.Sin(deg*degToRad-0.5*zipTieH/r)
	x1o := (r + 0.5*zipTieW) * math.Cos(deg*degToRad-0.5*zipTieH/r)
	y1o := (r + 0.5*zipTieW) * math.Sin(deg*degToRad-0.5*zipTieH/r)
	x2i := (r - 0.5*zipTieW) * math.Cos(deg*degToRad+0.5*zipTieH/r)
	y2i := (r - 0.5*zipTieW) * math.Sin(deg*degToRad+0.5*zipTieH/r)
	x2o := (r + 0.5*zipTieW) * math.Cos(deg*degToRad+0.5*zipTieH/r)
	y2o := (r + 0.5*zipTieW) * math.Sin(deg*degToRad+0.5*zipTieH/r)
	cx := canvas.CenterX
	cy := canvas.CenterY
	canvas.Line(cx+x1i, cy+y1i, cx+x2i, cy+y2i)
	canvas.Line(cx+x2i, cy+y2i, cx+x2o, cy+y2o)
	canvas.Line(cx+x2o, cy+y2o, cx+x1o, cy+y1o)
	canvas.Line(cx+x1o, cy+y1o, cx+x1i, cy+y1i)
}

// cutZipTie draws all the cuts for the zip ties.
func cutZipTie(deg float64) {
	cutZipTieAtRadius(fullR-0.25*in, deg-2.0)
	cutZipTieAtRadius(fullR-0.25*in, deg+2.0)
	cutZipTieAtRadius(magnetIR+rDiff-0.25*in, deg-3.0)
	cutZipTieAtRadius(magnetIR+rDiff-0.25*in, deg+3.0)
}

// cutWedgeMagnet draws the cuts for a wedge magnet.
func cutWedgeMagnet(deg, off float64) {
	x1i := math.Cos((deg-halfDelta)*degToRad + tolerance/magnetIR)
	y1i := math.Sin((deg-halfDelta)*degToRad + tolerance/magnetIR)
	x1o := math.Cos((deg-halfDelta)*degToRad + tolerance/magnetOR)
	y1o := math.Sin((deg-halfDelta)*degToRad + tolerance/magnetOR)
	xoff := off * math.Cos(deg*degToRad)
	yoff := off * math.Sin(deg*degToRad)
	tx := tolerance * math.Cos(deg*degToRad)
	ty := tolerance * math.Sin(deg*degToRad)
	x2i := math.Cos((deg+halfDelta)*degToRad - tolerance/magnetIR)
	y2i := math.Sin((deg+halfDelta)*degToRad - tolerance/magnetIR)
	x2o := math.Cos((deg+halfDelta)*degToRad - tolerance/magnetOR)
	y2o := math.Sin((deg+halfDelta)*degToRad - tolerance/magnetOR)
	cx := canvas.CenterX + xoff
	cy := canvas.CenterY + yoff
	ilx := cx + magnetIR*x1i + tx
	ily := cy + magnetIR*y1i + ty
	olx := cx + magnetOR*x1o - tx
	oly := cy + magnetOR*y1o - ty
	irx := cx + magnetIR*x2i + tx
	iry := cy + magnetIR*y2i + ty
	orx := cx + magnetOR*x2o - tx
	ory := cy + magnetOR*y2o - ty
	canvas.Arc(ilx, ily, magnetIR, magnetIR, 0, false, true, irx, iry)
	canvas.Line(ilx, ily, olx, oly)
	canvas.Arc(olx, oly, magnetOR, magnetOR, 0, false, true, orx, ory)
	canvas.Line(irx, iry, orx, ory)
}

// cutDiskMagnet cuts a circular hole for a disk magnet.
func cutDiskMagnet(deg float64) {
	r := 0.5 * 1.25 * in
	xoff := (magnetIR + rDiff - 2.0*r) * math.Cos(deg*degToRad)
	yoff := (magnetIR + rDiff - 2.0*r) * math.Sin(deg*degToRad)
	x := canvas.CenterX + xoff
	y := canvas.CenterY + yoff
	canvas.InnerCircle(x, y, r)
}

func main() {
	flag.Parse()
	canvas = ponoko2d.NewP2(os.Stdout)
	canvas.GCut()
	canvas.OuterCircle(canvas.CenterX, canvas.CenterY, fullR)
	canvas.InnerCircle(canvas.CenterX, canvas.CenterY, bearingOR)
	for i := 0; i < numMagnets; i++ {
		a := 360.0 * float64(i) / numMagnets
		if *magnets {
			cutWedgeMagnet(a, rDiff)
		}
		cutZipTie(a)
	}
	canvas.Gend()
	canvas.End()
}
