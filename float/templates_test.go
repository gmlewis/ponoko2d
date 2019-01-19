// Tests the templates generation code.

package ponoko2d

import (
	"bytes"
	"testing"
)

const (
	p1expect = `<?xml version="1.0" encoding="utf-8"?>
<!-- Generator: http://github.com/gmlewis/ponoko2d  -->
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
	 width="541.417px" height="541.417px" viewBox="0 0 541.417 541.417" enable-background="new 0 0 541.417 541.417"
	 xml:space="preserve">
<g id="Template_Background">
	<rect x="0" y="0" fill="#F6921E" width="541.417" height="541.417"/>
	<rect x="14.1735" y="14.1735" fill="#FFFFFF" width="513.07" height="513.07"/>
</g>
<g id="Your_Design">
<g style="fill:none;stroke:#00f;stroke-width:0.01mm">
<line x1="270.000" y1="271.418" x2="342.000" y2="271.418" />
<circle cx="270.000" cy="271.418" r="225.000" />
</g>
</g>
</svg>
`

	p2expect = `<?xml version="1.0" encoding="utf-8"?>
<!-- Generator: http://github.com/gmlewis/ponoko2d  -->
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
	 width="1116.85px" height="1116.85px" viewBox="0 0 1116.85 1116.85" enable-background="new 0 0 1116.85 1116.85"
	 xml:space="preserve">
<g id="Template_Background">
	<rect x="0" y="0" fill="#F6921E" width="1116.85" height="1116.85"/>
	<rect x="14.1735" y="14.1735" fill="#FFFFFF" width="1088.503" height="1088.503"/>
</g>
<g id="Your_Design">
<g style="fill:none;stroke:#00f;stroke-width:0.01mm">
<circle cx="557.716" cy="559.133" r="504.000" />
<line x1="557.716" y1="559.133" x2="989.716" y2="559.133" />
</g>
</g>
</svg>
`

	p3expect = `<?xml version="1.0" encoding="utf-8"?>
<!-- Generator: http://github.com/gmlewis/ponoko2d  -->
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
	 width="2267.717px" height="1116.851px" viewBox="0 0 2267.717 1116.851" enable-background="new 0 0 2267.717 1116.851"
	 xml:space="preserve">
<g id="Template_Background">
	<rect x="0" y="0" fill="#F6921E" width="2267.717" height="1116.851"/>
	<rect x="14.1735" y="14.1735" fill="#FFFFFF" width="2239.37" height="1088.5040000000001"/>
</g>
<g id="Your_Design">
<g style="fill:none;stroke:#00f;stroke-width:0.01mm">
<circle cx="557.716" cy="559.133" r="504.213" />
<circle cx="1133.858" cy="558.425" r="533.101" />
</g>
</g>
</svg>
`
)

func TestP1(t *testing.T) {
	b := new(bytes.Buffer)
	s := NewP1(b)
	s.GCut()
	s.Line(270, 271.418, 342, 271.418)
	s.Circle(270, 271.418, 225)
	s.Gend()
	s.End()
	if string(b.Bytes()) != p1expect {
		t.Errorf("NewP1 =\n%v\n, want:\n%v", string(b.Bytes()), p1expect)
	}
}

func TestP2(t *testing.T) {
	b := new(bytes.Buffer)
	s := NewP2(b)
	s.GCut()
	s.Circle(557.716, 559.133, 504)
	s.Line(557.716, 559.133, 989.716, 559.133)
	s.Gend()
	s.End()
	if string(b.Bytes()) != p2expect {
		t.Errorf("NewP2 =\n%v\n, want:\n%v", string(b.Bytes()), p2expect)
	}
}

func TestP3(t *testing.T) {
	b := new(bytes.Buffer)
	s := NewP3(b)
	s.GCut()
	s.OuterCircle(557.716, 559.133, 504)
	s.InnerCircle(1133.858, 558.425, 533.314)
	s.Gend()
	s.End()
	if string(b.Bytes()) != p3expect {
		t.Errorf("NewP3 =\n%v\n, want:\n%v", string(b.Bytes()), p3expect)
	}
}
