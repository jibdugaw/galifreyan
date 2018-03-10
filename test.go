package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/fogleman/gg"
)

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

type glyph struct {
	step    float64
	letters []byte
	radius  float64
}

type colloquy struct {
	step   float64
	words  []glyph
	radius float64
}

func main() {
	input := "gallifrey falls no more"
	// our globals which wil be the canvas size and then some basic starting point stuff
	canvasSide := float64(1000)
	originX := canvasSide / 2
	originY := canvasSide / 2
	// MAGERY: Gallifreyan characters start at the 6 o'clock position when
	// looking at a clock face, or 90º and run counter clockwise
	firstWordStart := 270

	// TODO - retrieve sentences and create colloquy's for each
	// TODO - parse sentence creating glyph holder for each word
	// TODO - parse parse letters into glyph drawing profile
	// sentenceCount := 1
	wordCount := 5
	wordStep := float64(360 / wordCount) // MAGERY: evenly space glyphs in sentence circle
	fmt.Printf("Assigned words are %v with step of %v\n", wordCount, wordStep)
	fmt.Printf("originX: %v \t originY: %v\n", originX, originY)
	fmt.Printf("firstWordStart: %v \n", firstWordStart)
	fmt.Printf("I think the step is: %v\n", wordStep)

	//	sentence := colloquy{wordStep, make([]glyph, wordCount), float64(320)}
	words := strings.Fields(input)
	sentence := NewColloquy(originX, originY, float64(0.8)*originX, wordStep, "", make([]*Glyph, len(words)))
	myStep := float64(firstWordStart)
	for i := range words {
		fmt.Printf("This glyp  offset goes to degrees %v\n", myStep)
		fmt.Printf("i is: %v\n", i)
		g := NewGlyph(sentence, -1.0, -1.0, words[i], 0.0, 0.0)
		sentence.glyphs[i] = g
		// update myStep and handle wrap arounds
		myStep += wordStep
		if myStep >= 360 {
			myStep -= 360
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

// DrawColloquy iterates over glyphs of a Coloquy and draws them
func DrawColloquy(x0 float64, y0 float64, dc *gg.Context, c *Colloquy) {
	// iterate glyphs and draw those
	for i := range c.glyphs {
		fmt.Printf("My step is %v\n", c.glyphs[i].step)
		deltaX := float64(c.radius) * math.Cos(Radians(c.glyphs[i].step))
		deltaY := float64(c.radius) * math.Sin(Radians(c.glyphs[i].step))
		switch {
		case c.glyphs[i].step <= 90:
			if deltaX < 0 {
				deltaX = -deltaX
			}
			if deltaY > 0 {
				deltaY = -deltaY
			}
			dc.SetRGB(255, 0, 0)
			fmt.Printf("I think I'm in quadrant I (deltaX, deltaY) = (%v,%v)\n", deltaX, deltaY)
			fmt.Printf("(x,y) = (%v, %v)\n", float64(x0)+deltaX, float64(y0)+deltaY)

			break
		case c.glyphs[i].step <= 180:
			if deltaX > 0 {
				deltaX = -deltaX
			}
			if deltaY > 0 {
				deltaY = -deltaY
			}
			dc.SetRGB(0, 255, 0)
			fmt.Printf("I think I'm in quadrant II (deltaX, deltaY) = (%v,%v)\n", deltaX, deltaY)
			fmt.Printf("(x,y) = (%v, %v)\n", float64(x0)+deltaX, float64(y0)+deltaY)
			break
		case c.glyphs[i].step <= 270:
			if deltaX > 0 {
				deltaX = -deltaX
			}
			if deltaY < 0 {
				deltaY = -deltaY
			}
			dc.SetRGB(0, 0, 255)
			fmt.Printf("I think I'm in quadrant III (deltaX, deltaY) = (%v,%v)\n", deltaX, deltaY)
			fmt.Printf("(x,y) = (%v, %v)\n", float64(x0)+deltaX, float64(y0)+deltaY)
			break
		case c.glyphs[i].step <= 360:
			if deltaX < 0 {
				deltaX = -deltaX
			}
			if deltaY < 0 {
				deltaY = -deltaY
			}
			dc.SetRGB(0, 0, 0)
			fmt.Printf("I think I'm in quadrant IV (deltaX, deltaY) = (%v,%v)\n", deltaX, deltaY)
			fmt.Printf("(x,y) = (%v, %v)\n", float64(x0)+deltaX, float64(y0)+deltaY)
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
