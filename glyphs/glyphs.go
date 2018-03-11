package glyphs

import (
	"strings"
)

const (
	open = iota
	inside
	arc
	on
)

// Glyph represents the galifrayen transliteration of a character
type Glyph struct {
	Owner      *Colloquy // used to find relative position
	XPos       float64   // TODO - Renderer knows positioning
	YPos       float64   // TODO - Renderer knows positioning
	Word       string    // original word we received TODO - save on memory if we bypass this
	Characters []*GlyphChar
	Step       float64 // TODO - maybe Renderer
	Radius     float64 // TODO - maybe renderer
}

// GlyphChar represents a single character to be converted into a Galifreyan word glyph
// TODO - decide if this makes sense or not
type GlyphChar struct {
	Letter  string // TODO - Gallifreyan is phonic some sounds have multiple characters
	Type    int
	SubType int
}

// Colloquy represents a string of words (aka sentence) in Galifreyan
type Colloquy struct {
	X        float64
	Y        float64
	Radius   float64
	Step     float64
	Sentence string
	Glyphs   []*Glyph
}

// NewColloquy creates a new Galifreyan sentence
func NewColloquy(x0 float64, y0 float64, r float64, s float64, sent string, gs []*Glyph) *Colloquy {
	ret := &Colloquy{x0, y0, r, s, sent, gs}
	words := strings.Fields(sent)
	glyphs := make([]*Glyph, len(words))
	for word := range words {
		g := NewGlyph(ret, float64(-1), float64(-1), words[word], float64(word), float64(-1))
		glyphs[word] = g
	}
	ret.Glyphs = glyphs
	return ret
}

// NewGlyph returns a properly defaulted Glyph object
func NewGlyph(c *Colloquy, x float64, y float64, w string, step float64, radius float64) *Glyph {
	// split sentence by white space
	g := &Glyph{c, x, y, w, nil, 0, 0}
	g.parse()
	return g
}

func (g *Glyph) parse() {
	for i := 0; i < len(g.Word); i++ {
		gc := &GlyphChar{"  ", -1, -1}
		switch g.Word[i] {
		case 'c':
		case 's':
		case 't':
			// if next is 'h' then substitute
			if g.Word[i+1] == 'h' {
				gc.Letter = string(g.Word[i])
				gc.Letter += string(g.Word[i+1])
				i++
				break
			}
			// rune 'c' has two special cases make it soft or hard
			// SOFT - 'c' followed by any one of [iey]
			if g.Word[i] == 'c' {

				if g.Word[i] == 'i' || g.Word[i] == 'e' || g.Word[i] == 'y' {
					gc.Letter = string('s')
					break
				} else {
					gc.Letter = string('k')
					break
				}
			}
			gc.Letter = string(g.Word[i])
			break
			// check for rune combinaion 'ng'
		case 'n':
			if g.Word[i+1] == 'g' {
				gc.Letter = string(g.Word[i])
				gc.Letter += string(g.Word[i+1])
				i++
				break
			}
			// check for rune combinaion 'qu'
		case 'q':
			if g.Word[i+1] == 'u' {
				gc.Letter = string(g.Word[i])
				gc.Letter += string(g.Word[i+1])
				i++
				break
			}
		default:
			gc.Letter = string(g.Word[i])
			break
		}
	}
}
