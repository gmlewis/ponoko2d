package ponoko2d

import "io"

const (
	header = `<?xml version="1.0" encoding="utf-8"?>
<!-- Generator: http://github.com/gmlewis/ponoko2d  -->
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
	 width="%vpx" height="%vpx" viewBox="0 0 %v %v" enable-background="new 0 0 %v %v"
	 xml:space="preserve">
<g id="Template_Background">
	<rect x="0" y="0" fill="#F6921E" width="%v" height="%v"/>
	<rect x="14.1735" y="14.1735" fill="#FFFFFF" width="%v" height="%v"/>
</g>
<g id="Your_Design">
`
	p1w = 541.417
	p1h = 541.417

	p2w = 1116.85
	p2h = 1116.85

	p3w = 2267.717
	p3h = 1116.851

	margin = 28.347

	// In is the number of points per inch.
	In = 72.0 // points == pixels
	// Mm is the number of points per millimeter.
	Mm = In / 25.4
	// LaserWidth is the width of the laser in points.
	LaserWidth = 0.15 * Mm
)

// newTemplate creates a new SVG object with a custom header that matches a Ponoko template.
func newTemplate(iow io.Writer, w, h float64) *SVG {
	s := &SVG{
		Writer:    iow,
		W:         w,
		H:         h,
		CenterX:   0.5 * w,
		CenterY:   0.5 * h,
		Decimals:  3,
		endString: "</g>\n</svg>",
	}
	s.printf(header, w, h, w, h, w, h, w, h, w-margin, h-margin)
	return s
}

// NewP1 creates a new SVG object with a custom header that matches the Ponoko P1 template.
func NewP1(w io.Writer) *SVG {
	return newTemplate(w, p1w, p1h)
}

// NewP2 creates a new SVG object with a custom header that matches the Ponoko P2 template.
func NewP2(w io.Writer) *SVG {
	return newTemplate(w, p2w, p2h)
}

// NewP3 creates a new SVG object with a custom header that matches the Ponoko P3 template.
func NewP3(w io.Writer) *SVG {
	return newTemplate(w, p3w, p3h)
}

// GCut starts a Gstyle group for laser-cut lines.  GEnd() must be called to terminate the style.
func (s *SVG) GCut() {
	s.Gstyle("fill:none;stroke:#00f;stroke-width:0.01mm") // Blue = cut
}

// GEngrave starts a Gstyle group for laser-cut engraving.  GEnd() must be called to terminate the style.
func (s *SVG) GEngrave() {
	s.Gstyle("fill:none;stroke:#f00;stroke-width:0.01mm") // Red = engrave
}

// OuterCircle laser-cuts an outer circle such that the radius is increased by the laser width.
func (s *SVG) OuterCircle(x, y, r float64) {
	s.Circle(x, y, r+0.5*LaserWidth)
}

// InnerCircle laser-cuts an inner circle such that the radius is decreased by the laser width.
func (s *SVG) InnerCircle(x, y, r float64) {
	s.Circle(x, y, r-0.5*LaserWidth)
}
