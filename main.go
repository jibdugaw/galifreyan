package main

// TODO - add placeholder circles for Glyphs at each step ont he perimeter of the Colloquy

import (
	"fmt"
	"math"
	"strings"

	gglyphs "github.com/jibdugaw/galifreyan/glyphs"

	"github.com/fogleman/gg"
)

// RADIANS defines the number of RADIANS comprising a circle
// fractions of this value are used throughout
const RADIANS float64 = 2 * math.Pi

/*

func main() {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}
	dc.SavePNG("out.png")
}
*/

func main() {
	input := "gallifrey falls no more"
	words := strings.Fields(input)
	wordCount := float64(len(words))
	wordStep := RADIANS / wordCount
	// MAGERY: Gallifreyan characters start at the 6 o'clock position when
	// looking at a clock face, or 90º and run counter clockwise
	firstWordStart := .75 * RADIANS

	// our globals which wil be the canvas size and then some basic starting point stuff
	canvasSide := float64(1000)
	originX := canvasSide / 2
	originY := canvasSide / 2

	// TODO - retrieve sentences and create colloquy's for each
	// TODO - parse sentence creating glyph holder for each word
	// TODO - parse parse letters into glyph drawing profile
	sentence := gglyphs.NewColloquy(originX, originY, float64(0.8)*originX, wordStep, "", make([]*gglyphs.Glyph, len(words)))
	myStep := firstWordStart
	for i := range words {
		g := gglyphs.NewGlyph(sentence, -1.0, -1.0, words[i], myStep, 0.0)
		sentence.Glyphs = append(sentence.Glyphs, g)
		// update myStep and handle wrap arounds
		myStep -= wordStep
		if myStep < 0 {
			myStep += 2 * RADIANS
		}
	}

	dc := gg.NewContext(1000, 1000)
	// using originX, originY, radius and offsets for each glyph in the colloquy
	// using originX, originY and radius of colloquy draw each word glyph
	dc.SetLineWidth(10)
	DrawColloquy(originX, originY, dc, sentence)
	//DrawB(dc)
	//dc.DrawArc(400, 325, 300, Radians(0), Radians(90))
	//dc.SetRGB(255, 188, 55)
	//dc.Fill()
	dc.Stroke()
	dc.SavePNG("out.png")
}

// Radians converts signed degrees into signed radians
func Radians(degrees float64) float64 {
	return (degrees * float64(math.Pi) / float64(180))
}

// DrawColloquy iterates over Glyphs of a Coloquy and draws them
func DrawColloquy(x0 float64, y0 float64, dc *gg.Context, c *gglyphs.Colloquy) {
	// iterate glyphs and draw those
	for i := range c.Glyphs {
		fmt.Printf("working on Glyph: %v\n", c.Glyphs[i])
		deltaX := float64(c.Radius) * math.Cos(c.Glyphs[i].Step)
		deltaY := float64(c.Radius) * math.Sin(c.Glyphs[i].Step)

		switch {
		case c.Glyphs[i].Step <= .25*RADIANS:
			if deltaX < 0 {
				deltaX = -deltaX
			}
			if deltaY > 0 {
				deltaY = -deltaY
			}
			dc.SetRGB(255, 0, 0)
			fmt.Printf("I think I'm in quadrant I (deltaX, deltaY) = (%v,%v)\n", deltaX, deltaY)
			fmt.Printf("(x,y) = (%v, %v)\n", float64(x0)+deltaX, float64(y0)+deltaY)

			dc.DrawCircle(c.Glyphs[i].XPos, c.Glyphs[i].YPos, c.Glyphs[i].Radius)
			break
		case c.Glyphs[i].Step <= .5*RADIANS:
			if deltaX > 0 {
				deltaX = -deltaX
			}
			if deltaY > 0 {
				deltaY = -deltaY
			}
			dc.SetRGB(0, 255, 0)
			fmt.Printf("I think I'm in quadrant II (deltaX, deltaY) = (%v,%v)\n", deltaX, deltaY)
			fmt.Printf("(x,y) = (%v, %v)\n", float64(x0)+deltaX, float64(y0)+deltaY)
			dc.DrawCircle(c.Glyphs[i].XPos, c.Glyphs[i].YPos, c.Glyphs[i].Radius)
			break
		case c.Glyphs[i].Step <= .75*RADIANS:
			if deltaX > 0 {
				deltaX = -deltaX
			}
			if deltaY < 0 {
				deltaY = -deltaY
			}
			dc.SetRGB(0, 0, 255)
			fmt.Printf("I think I'm in quadrant III (deltaX, deltaY) = (%v,%v)\n", deltaX, deltaY)
			fmt.Printf("(x,y) = (%v, %v)\n", float64(x0)+deltaX, float64(y0)+deltaY)
			dc.DrawCircle(c.Glyphs[i].XPos, c.Glyphs[i].YPos, c.Glyphs[i].Radius)
			break
		case c.Glyphs[i].Step <= RADIANS:
			if deltaX < 0 {
				deltaX = -deltaX
			}
			if deltaY < 0 {
				deltaY = -deltaY
			}
			dc.SetRGB(0, 0, 0)
			fmt.Printf("I think I'm in quadrant IV (deltaX, deltaY) = (%v,%v)\n", deltaX, deltaY)
			fmt.Printf("(x,y) = (%v, %v)\n", float64(x0)+deltaX, float64(y0)+deltaY)
			dc.DrawCircle(c.Glyphs[i].XPos, c.Glyphs[i].YPos, c.Glyphs[i].Radius)
			break
		}

		//x := float64(x0) + deltaX
		//y := float64(y0) + deltaY

	}
}

// DrawB draws a Galifreyan B
func DrawB(dc *gg.Context) {
	dc.DrawArc(500, 500, 250, Radians(110), Radians(430)) //8.125)
	dc.DrawArc(250, 250, 30, Radians(0), Radians(90))     //8.125)
	//dc.
}
