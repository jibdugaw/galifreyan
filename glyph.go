package main // TODO - put this is in appropriate namespace

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
	owner      *Colloquy // used to find relative position
	xPos       float64   // TODO - Renderer knows positioning
	yPos       float64   // TODO - Renderer knows positioning
	word       string    // original word we received TODO - save on memory if we bypass this
	characters []*glyphChar
	step       float64 // TODO - maybe Renderer
	radius     float64 // TODO - maybe renderer
}

// TODO - decide if this makes sense or not
type glyphChar struct {
	letter string // TODO - Gallifreyan is phonic some sounds have multiple characters
	gType  int
	gSub   int
}

// Colloquy represents a string of words (aka sentence) in Galifreyan
type Colloquy struct {
	x        float64
	y        float64
	radius   float64
	step     float64
	sentence string
	glyphs   []*Glyph
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
	ret.glyphs = glyphs
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
	for i := 0; i < len(g.word); i++ {
		gc := &glyphChar{"  ", -1, -1}
		switch g.word[i] {
		case 'c':
		case 's':
		case 't':
			// if next is 'h' then substitute
			if g.word[i+1] == 'h' {
				gc.letter = string(g.word[i])
				gc.letter += string(g.word[i+1])
				i++
				break
			}
			// rune 'c' has two special cases make it soft or hard
			// SOFT - 'c' followed by any one of [iey]
			if g.word[i] == 'c' {

				if g.word[i] == 'i' || g.word[i] == 'e' || g.word[i] == 'y' {
					gc.letter = string('s')
					break
				} else {
					gc.letter = string('k')
					break
				}
			}
			gc.letter = string(g.word[i])
			break
			// check for rune combinaion 'ng'
		case 'n':
			if g.word[i+1] == 'g' {
				gc.letter = string(g.word[i])
				gc.letter += string(g.word[i+1])
				i++
				break
			}
			// check for rune combinaion 'qu'
		case 'q':
			if g.word[i+1] == 'u' {
				gc.letter = string(g.word[i])
				gc.letter += string(g.word[i+1])
				i++
				break
			}
		default:
			gc.letter = string(g.word[i])
			break
		}
	}
}
