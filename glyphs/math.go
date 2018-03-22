package glyphs

import "math"

// RADIANS - the number of radians in a 360º cirle useful for fster maths
const RADIANS = 2 * math.Pi

// Dist calculates the distance between two circles (e.g. Glyphs, Coloquys or CharGlyphs
func Dist(c1, c2 *Circle) float64 {
	return math.Sqrt(math.Pow((c1.X-c2.X), 2) + math.Pow((c1.Y-c2.Y), 2))
}

// Angle calculates the angle at which the two circles intersect
func Angle(c1, c2 *Circle) float64 {
	d := Dist(c1, c2)
	return math.Acos((c2.Radius*c2.Radius - d*d - c1.Radius*c1.Radius) / (-2 * d * c1.Radius))

}

// FindPointFromAngle calculates point relative to center of circle at provide angle (in Radians)
func FindPointFromAngle(c *Circle, r float64, rads float64) (float64, float64) {
	return c.X + math.Cos(rads)*r, c.Y + math.Sin(rads)*r
}

// Step will calculate the radians by which each Glyph or GlyphChar must rotate
// before being drawn
func Step(count int) float64 {
	return RADIANS / float64(count)
}

// Offset - calculates the radian rotation necessary to manage Glyph or CharGlyph locations
func Offset(step float64, count int) float64 {
	ret := .25*RADIANS - (float64(count) * step)
	for ret > RADIANS {
		ret -= 2 * RADIANS
	}
	return ret
}

// ScaleRadius - scales a "default" radius based on the number of children the Colloquy/Glyph
// is to include
func ScaleRadius(radius float64, children int) float64 {
	// calculate radius for Glyphs that comprise the Colloquy
	ret := 2.5 * radius / float64(children+4)
	if children == 1 {
		ret = .8 * radius
		//c.Coords.Radius = .8 * c.Coords.Radius
	}
	return ret
}
