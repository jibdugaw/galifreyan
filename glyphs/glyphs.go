package glyphs

import (
	"math"
	"strings"

	"github.com/fogleman/gg"
)

var decorations = [6]int{0, 2, 3, 3, 1, 2}

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

// Circle is basically a Galifreyan Cologuy, Glyph "circle" element
// TODO - rename this monstrosity I was obviously tired
type Circle struct {
	X      float64
	Y      float64
	Radius float64
}

// Glyph represents the galifrayen transliteration of a character
type Glyph struct {
	Owner      *Colloquy // used to find relative position
	Coords     *Circle
	Angle      float64 // for non-combining letter shapes where do they go on the Glyph
	Word       string  // original word we received TODO - save on memory if we bypass this
	Characters []*GlyphChar
}

// GlyphChar represents a single character to be converted into a Galifreyan word glyph
// TODO - decide if this makes sense or not
type GlyphChar struct {
	Coords     *Circle // our position is RELATIVE to the center of the Glyph
	Base       int     // base character shape
	Decoration int     // what decorations does the basic shape receive
	Letter     string  // TODO - Gallifreyan is phonic some sounds have multiple characters
}

// Colloquy represents a string of words (aka sentence) in Galifreyan
type Colloquy struct {
	Coords   *Circle
	Step     float64
	Sentence string
	Glyphs   []*Glyph
}

// NewColloquy creates a new Galifreyan sentence
// x - x coordinate around which to draw
// y - x coordinate around which to draw
// r - radius used to draw colloquy
// sentence - the sentence to be represented
//
// IMPORTANT NOTE:
//
// Radius (r) provided is a guidlines. The actual radius will be modified based on the number
// of words (Glyphs) in the sentence (Colloquy) which may in turn be modified basd on word
// complexity such as characters (CharGlyph) and how those combine. The actual radius used to draw
// the Colloquy will be less than or equal to the radius provided.
func NewColloquy(x0 float64, y0 float64, r float64, sentence string) *Colloquy {

	// use provide radius as a seed
	ret := &Colloquy{&Circle{x0, y0, r}, 0, sentence, nil}
	//words := strings.Fields(sentence)
	//ret.Step = RADIANS / float64(len(words))
	ret.Parse()

	return ret
}

// Parse parses a Glyph and seeds the GlyphChar
func (g *Glyph) Parse() {
	midWord := true
	var letter string

	for i := 0; i < len(g.Word); i++ {
		if i == len(g.Word)-1 {
			midWord = !midWord
		}

		// defaut letter
		letter = string(g.Word[i])

		switch g.Word[i] {
		case 'c', 's', 't':
			// if next is 'h' then substitute
			if midWord {
				if g.Word[i+1] == 'h' {
					letter += string(g.Word[i+1])
					i++
				} else if g.Word[i] == 'c' {
					// rune 'c' has two special cases make it soft or hard
					// SOFT - 'c' followed by any one of [iey]
					if g.Word[i+1] == 'i' || g.Word[i+1] == 'e' || g.Word[i+1] == 'y' {
						letter = string('s')
					} else {
						letter = string('k')
					}
				}
			}

		// check for rune combinaion 'ng'
		case 'n':
			if midWord && g.Word[i+1] == 'g' {
				letter += string(g.Word[i+1])
				i++
			}

		// check for rune combinaion 'qu'
		case 'q':
			if midWord && g.Word[i+1] == 'u' {
				letter += string(g.Word[i+1])
				i++
			}

		default:
		}
		g.Characters = append(g.Characters, newGlyphChar(letter))
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
	//dc.DrawCircle(g.X, g.Y, g.Radius)
	if err := dc.LoadFontFace("/Library/Fonts/Arial.ttf", 14); err != nil {
		panic(err)
	}
	// TODO - following is useful for DEBUG mode
	//dc.DrawStringAnchored(g.Word, g.Coords.X, g.Coords.Y, .5, .5)
	dc.DrawCircle(g.Coords.X, g.Coords.Y, g.Coords.Radius)
	// draw each glyph ;-)
	//x := 0.0
	//y := 0.0
	step := Step(len(g.Characters))
	offSet := 0.0
	x := 0.0
	y := 0.0
	letterRadius := ScaleRadius(g.Coords.Radius, len(g.Characters))
	for i := range g.Characters {
		offSet = Offset(step, i)
		x = g.Coords.X + float64(g.Coords.Radius)*math.Cos(offSet)
		y = g.Coords.Y + float64(g.Coords.Radius)*math.Sin(offSet)
		dc.DrawStringAnchored(g.Characters[i].Letter, x, y, .5, .5)
		dc.DrawCircle(x, y, letterRadius)
	}
}

// Parse will parse the Colloquy sentence and seed Glyphs
// Side Effects:
//   -- sets the actual radius to be used when drawing the Colloqur
//   -- sets the step for the Colloquy based on the words in the provded sentence
func (c *Colloquy) Parse() {

	words := strings.Fields(c.Sentence)
	count := len(words)
	c.Step = Step(count)

	// initiate slice of Glyphs for the Colloquey
	c.Glyphs = make([]*Glyph, count)

	glyphRadius := ScaleRadius(c.Coords.Radius, count)

	// special case manage the Colloquy as a side effect when
	// number of Glyphs is one (1)
	if count == 1 {
		c.Coords.Radius = .8 * c.Coords.Radius
	}

	// create Glyph for each word
	for i := range words {
		g := NewGlyph(c, words[i], Offset(c.Step, i))
		// created Glyph does not have an origin YET
		circle := &Circle{0.0, 0.0, glyphRadius}
		circle.X = c.Coords.X + float64(c.Coords.Radius-glyphRadius)*math.Cos(g.Angle)
		circle.Y = c.Coords.Y + float64(c.Coords.Radius-glyphRadius)*math.Sin(g.Angle)

		g.Coords = circle
		// calculate Glyph X, Y
		//
		g.Parse()
		c.Glyphs[i] = g
	}

}

// NewGlyph returns a properly defaulted Glyph object
// NOTE: a properly defaulted Glyph does not know its origin for painting
// this may be problematic
func NewGlyph(c *Colloquy, w string, angle float64) *Glyph {
	ret := &Glyph{}
	ret.Coords = &Circle{}
	ret.Owner = c
	ret.Word = w
	ret.Angle = angle
	// calculate the radius of this glyph based on parent
	// TODO - likely to need to change radius basedd on words in the sentence
	// TODO - determine who should own this "glyph" radius
	// TODO - likely need similar thoughts around individual CharGlyphs
	/*
		ret.Coords.Radius = 10 // TODO - this is a future BUG!
		ret.Coords.X = c.Coords.X + float64(c.Coords.Radius)*math.Cos(ret.Angle)
		ret.Coords.Y = c.Coords.Y + float64(c.Coords.Radius)*math.Sin(ret.Angle)
	*/
	return ret
}

func newGlyphChar(l string) *GlyphChar {
	ret := &GlyphChar{}
	ret.Letter = l
	ret.Base = charMap[l][0]
	ret.Decoration = charMap[l][1]
	return ret

}
