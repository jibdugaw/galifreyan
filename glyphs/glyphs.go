package glyphs

import (
	"fmt"
	"math"
	"strings"

	"github.com/fogleman/gg"
)

const (
	open = iota
	inside
	arc
	on
)

// RADIANS defines the number of RADIANS comprising a circle
// fractions of this value are used throughout
const RADIANS float64 = 2 * math.Pi

// Glyph represents the galifrayen transliteration of a character
type Glyph struct {
	Owner      *Colloquy // used to find relative position
	X          float64   // needed to tell composition Glyphs where to go
	Y          float64   // needed to tell composition Glyphs where to go
	Radius     float64   //
	Angle      float64   // for non-combining glyphs where do they go on the Glyph
	Word       string    // original word we received TODO - save on memory if we bypass this
	Characters []*GlyphChar
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
// x - x coordinate around which to draw
// y - x coordinate around which to draw
// r - radius used to draw colloquy
// sentence - the sentence to be represented
func NewColloquy(x0 float64, y0 float64, r float64, sentence string) *Colloquy {

	ret := &Colloquy{x0, y0, r, 0, sentence, nil}
	words := strings.Fields(sentence)
	ret.Step = RADIANS / float64(len(words))

	ret.Parse()

	return ret
}

// parse parses a Glyph and seeds the GlyphChar
func (g *Glyph) parse() {
	midWord := true
	for i := 0; i < len(g.Word); i++ {
		if i+1 >= len(g.Word) {
			midWord = !midWord
		}
		gc := &GlyphChar{"  ", -1, -1}
		switch g.Word[i] {
		case 'c':
		case 's':
		case 't':
			// if next is 'h' then substitute
			if midWord && g.Word[i+1] == 'h' {
				gc.Letter = string(g.Word[i])
				gc.Letter += string(g.Word[i+1])
				i++
				break
			}
			// rune 'c' has two special cases make it soft or hard
			// SOFT - 'c' followed by any one of [iey]
			if midWord && g.Word[i] == 'c' {

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
			if midWord && g.Word[i+1] == 'g' {
				gc.Letter = string(g.Word[i])
				gc.Letter += string(g.Word[i+1])
				i++
				break
			}
			// check for rune combinaion 'qu'
		case 'q':
			if midWord && g.Word[i+1] == 'u' {
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

// Draw iterates over Glyphs of a Coloquy and draws them to the provided canvas
func (c *Colloquy) Draw(dc *gg.Context) {

	for i := range c.Glyphs {
		c.Glyphs[i].Draw(dc)
	}

}

// Draw draws an individual Glyph
func (g *Glyph) Draw(dc *gg.Context) {
	fmt.Println("TODO - implement individual GlyphChar drawing")
	dc.DrawCircle(g.X, g.Y, g.Radius)
}

// Parse will parse the Colloquy sentence and seed Glyphs
func (c *Colloquy) Parse() {

	words := strings.Fields(c.Sentence)
	firstWordStart := .75 * RADIANS // TODO - verify
	myStep := firstWordStart
	c.Glyphs = make([]*Glyph, len(words))
	for i := range words {
		g := NewGlyph(c, words[i], myStep)
		g.Parse()
		c.Glyphs[i] = g
		// update myStep and handle wrap arounds
		myStep += c.Step
		if myStep > RADIANS {
			myStep -= 2 * RADIANS
		}
	}

}

// NewGlyph returns a properly defaulted Glyph object
//func NewGlyph(c *Colloquy, x float64, y float64, w string, step float64, radius float64) *Glyph {
func NewGlyph(c *Colloquy, w string, angle float64) *Glyph {
	ret := &Glyph{}
	ret.Owner = c
	ret.Word = w
	ret.Angle = angle
	ret.Radius = 10 // TODO - this is a future BUG!
	ret.X = c.X + float64(c.Radius)*math.Cos(ret.Angle)
	ret.Y = c.Y + float64(c.Radius)*math.Sin(ret.Angle)
	return ret
}

// Parse will parse the word assciated with the glyph seeding GlyphChars
func (*Glyph) Parse() {
	// TODO
}
