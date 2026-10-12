// Command mark writes the go-scikit mark, svg/color/go-scikit.svg:
//
//	cd mark && go run . > ../svg/color/go-scikit.svg
//
// go-scikit is the umbrella over seven organisations, and the mark is their
// stack drawn as a building: a base, the columns standing on it, and the
// beam they carry.
//
//   - the base is the arrays: go-ndarray and go-fft, one plate cut in two;
//   - the columns are the libraries that compute on those arrays:
//     go-numerics, go-statistics, go-dataframe and go-plot, one column each;
//   - the beam is go-notebook, the one place where all of them meet.
//
// Nothing is placed by hand. The layers below list the member organisations,
// and the geometry follows from them: a plate is cut into as many cells as its
// layer has members, the columns are spaced evenly across the width, and the
// vertical spacing leaves the same clear gap between the strokes of every two
// layers. Add a member to a layer and the mark redraws itself.
//
// The rest is the family grammar every go-* mark shares: a 256-unit rounded
// square (rx 56) filled with a diagonal two-stop gradient, a white line glyph
// at stroke-width 13 with round caps and joins. logos/brandgen renders the
// format matrix from the file this writes; the kit workflow runs this again
// and fails if a byte of svg/color differs.
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	light = "#8FA3C8" // the gradient's light stop
	dark  = "#2F3E63" // its dark stop

	stroke = 13.0 // the family's line width

	left, right = 48.0, 208.0 // the glyph's horizontal extent (stroke centres)
	plate       = 30.0        // height of a plate, between stroke centres
	gap         = 12.0        // clear space between the strokes of two layers
	radius      = 8.0         // corner radius of a plate
)

// layer is one level of the stack, bottom first.
type layer struct {
	members []string
	columns bool // drawn as columns rather than as a plate
}

var stack = []layer{
	{members: []string{"go-ndarray", "go-fft"}},
	{members: []string{"go-numerics", "go-statistics", "go-dataframe", "go-plot"}, columns: true},
	{members: []string{"go-notebook"}},
}

// num prints a coordinate with at most one decimal and no trailing zero.
func num(v float64) string {
	s := fmt.Sprintf("%.1f", math.Round(v*10)/10)
	s = strings.TrimSuffix(s, ".0")
	if s == "-0" {
		s = "0"
	}
	return s
}

func main() {
	// A column is half as tall again as a plate. The layers, strokes
	// included, are centred on the square.
	column := 1.5 * plate
	total := 0.0
	for _, l := range stack {
		if l.columns {
			total += column
		} else {
			total += plate
		}
	}
	total += float64(len(stack)-1) * (gap + stroke)

	var paths []string
	// Bottom layer first: y grows downwards, so start at the bottom edge.
	y := 128 + total/2
	for _, l := range stack {
		n := float64(len(l.members))
		if l.columns {
			// A line's round caps reach stroke/2 past its ends, exactly as a
			// plate's stroke reaches stroke/2 past its edge, so the same
			// spacing leaves the same clear gap.
			top, bottom := y-column, y
			for i := range l.members {
				x := left + (right-left)*(float64(i)+0.5)/n
				paths = append(paths, fmt.Sprintf(`<path d="M%s %sV%s"/>`, num(x), num(top), num(bottom)))
			}
			y -= column + gap + stroke
			continue
		}
		top := y - plate
		paths = append(paths, fmt.Sprintf(`<rect x="%s" y="%s" width="%s" height="%s" rx="%s"/>`,
			num(left), num(top), num(right-left), num(plate), num(radius)))
		for i := 1; i < len(l.members); i++ {
			x := left + (right-left)*float64(i)/n
			paths = append(paths, fmt.Sprintf(`<path d="M%s %sV%s"/>`, num(x), num(top), num(y)))
		}
		y -= plate + gap + stroke
	}

	fmt.Fprintf(os.Stdout, `<svg viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg" role="img" aria-label="go-scikit">
  <defs><linearGradient id="skgrad" x1="0" y1="0" x2="1" y2="1"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient></defs>
  <rect width="256" height="256" rx="56" fill="url(#skgrad)"/>
  <g fill="none" stroke="#FFFFFF" stroke-width="%s" stroke-linecap="round" stroke-linejoin="round">
    %s
  </g>
</svg>
`, light, dark, num(stroke), strings.Join(paths, "\n    "))
}
