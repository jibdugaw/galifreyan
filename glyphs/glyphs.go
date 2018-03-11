package glyphs

import (
	"fmt"
	"math"
	"strings"

	"github.com/fogleman/gg"
)

// RADIANS defines the number of RADIANS comprising a circle
// fractions of this value are used throughout
const RADIANS float64 = 2 * math.Pi

// charMap maps "characters" to GlyphChar base and decoration
var charMap = map[string][]int{
	"b":  {0, 0},
	"ch": {0, 1},
	"d":  {0, 2},
	"f":  {0, 3},
	"g":  {0, 4},
	"h":  {0, 5},

	"j": {1, 0},
	"k": {1, 1},
	"l": {1, 2},
	"m": {1, 3},
	"n": {1, 4},
	"p": {1, 5},

	"t":  {2, 0},
	"sh": {2, 1},
	"r":  {2, 2},
	"s":  {2, 3},
	"v":  {2, 4},
	"w":  {2, 5},

	"th": {3, 0},
	"y":  {3, 1},
	"Z":  {3, 2},
	"ng": {3, 3},
	"q":  {3, 4},
	"x":  {3, 5},

	"a": {4, 0},
	"e": {4, 1},
	"i": {4, 2},
	"o": {4, 3},
	"u": {4, 4},
}

/*
//assigns the subtype
var map = {
    "b": 1, "ch": 2, "d": 3, "f": 4, "g": 5, "h": 6,
    "j": 1, "k": 2, "l": 3, "m": 4, "n": 5, "p": 6,
    "t": 1, "sh": 2, "r": 3, "s": 4, "v": 5, "w": 6,
    "th": 1, "y": 2, "z": 3, "ng": 4, "qu": 5, "x": 6,
    "a": 1, "e": 2, "i": 3, "o": 4, "u": 5
};
*/

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
	Letter     string // TODO - Gallifreyan is phonic some sounds have multiple characters
	Base       int
	Decoration int
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

// Parse parses a Glyph and seeds the GlyphChar
func (g *Glyph) Parse() {
	midWord := true
	var letter string
	for i := 0; i < len(g.Word); i++ {
		if i+1 >= len(g.Word) {
			midWord = !midWord
		}
		switch g.Word[i] {
		case 'c':
		case 's':
		case 't':
			// if next is 'h' then substitute
			if midWord && g.Word[i+1] == 'h' {
				letter = string(g.Word[i])
				letter += string(g.Word[i+1])
				i++
				break
			}
			// rune 'c' has two special cases make it soft or hard
			// SOFT - 'c' followed by any one of [iey]
			if midWord && g.Word[i] == 'c' {

				if g.Word[i] == 'i' || g.Word[i] == 'e' || g.Word[i] == 'y' {
					letter = string('s')
					break
				} else {
					letter = string('k')
					break
				}
			}
			letter = string(g.Word[i])
			break
			// check for rune combinaion 'ng'
		case 'n':
			if midWord && g.Word[i+1] == 'g' {
				letter = string(g.Word[i])
				letter += string(g.Word[i+1])
				i++
				break
			}
			// check for rune combinaion 'qu'
		case 'q':
			if midWord && g.Word[i+1] == 'u' {
				letter = string(g.Word[i])
				letter += string(g.Word[i+1])
				i++
				break
			}
		default:
			letter = string(g.Word[i])
			break
		}
	}
	g.Characters = append(g.Characters, newGlyphChar(letter))
}

// Draw iterates over Glyphs of a Coloquy and draws them to the provided canvas
func (c *Colloquy) Draw(dc *gg.Context) {

	for i := range c.Glyphs {
		c.Glyphs[i].Draw(dc)
	}

}

// Draw draws an individual Glyph
func (g *Glyph) Draw(dc *gg.Context) {
	fmt.Printf("TODO - implement individual GlyphChar drawing: %v\n", g.Word)
	fmt.Printf("Glyph: (%v) to be anchored at (%v, %v)\n", g.Word, g.X, g.Y)
	//dc.DrawCircle(g.X, g.Y, g.Radius)
	if err := dc.LoadFontFace("/Library/Fonts/Arial.ttf", 14); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored(g.Word, g.X, g.Y, .5, .5)
	dc.DrawCircle(g.X, g.Y, 20)
}

// Parse will parse the Colloquy sentence and seed Glyphs
func (c *Colloquy) Parse() {

	words := strings.Fields(c.Sentence)
	firstWordStart := .25 * RADIANS // TODO - verify
	myStep := firstWordStart
	c.Glyphs = make([]*Glyph, len(words))
	for i := range words {
		g := NewGlyph(c, words[i], myStep)
		g.Parse()
		c.Glyphs[i] = g
		// update myStep and handle wrap arounds
		myStep -= c.Step
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

func newGlyphChar(l string) *GlyphChar {
	ret := &GlyphChar{}
	ret.Letter = l
	ret.Base = charMap[l][0]
	ret.Decoration = charMap[l][1]
	return ret

}
