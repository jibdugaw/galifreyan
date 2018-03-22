package main

// TODO - add placeholder circles for Glyphs at each step ont he perimeter of the Colloquy

import (
	"strings"

	gglyphs "github.com/jibdugaw/galifreyan/glyphs"

	"math"

	"github.com/fogleman/gg"
)

func main() {

	/*
		if true {
			const S = 1024
			dc := gg.NewContext(S, S)
			dc.SetRGB(1, 1, 1)
			dc.Clear()
			dc.SetRGB(0, 0, 0)
			if err := dc.LoadFontFace("/Library/Fonts/Arial.ttf", 14); err != nil {
				panic(err)
			}
			dc.DrawStringAnchored("Hello, world!", S/2, S/2, 0.5, 0.5)
			dc.DrawStringAnchored("bob, world!", S/2, S/2, 0.5, 0.5)
			dc.DrawStringAnchored("COLLISION", S/2, S/2, 0.5, 0.5)
			dc.SavePNG("test.png")
		}
	*/
	// TODO - get input dynamically
	//input := "gallifrey falls no more"
	input := "Doctor it is so much bigger on the inside"
	//input := "is"
	//input := "so"
	//input := "much"
	//input := "i am the doctor"

	// Gallifreyan has no case sensitivity
	input = strings.ToLower(input)

	// our globals which will be the canvas size and then some basic starting point stuff
	canvasSide := float64(1000)
	originX := canvasSide / 2
	originY := canvasSide / 2
	topRadius := .9 * math.Min(originX, originY) // control for non-square canvases

	// TODO - retrieve sentences and create colloquy's for each
	sentence := gglyphs.NewColloquy(originX, originY, topRadius, input)

	dc := gg.NewContext(int(canvasSide), int(canvasSide))
	// the following three lines are "weird" but they seem to be required to get
	// text to paint properly
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	// using originX, originY, radius and offsets for each glyph in the colloquy
	// using originX, originY and radius of colloquy draw each word glyph

	//dc.SetLineWidth(10)
	sentence.Draw(dc)
	//DrawB(dc)
	//dc.DrawArc(400, 325, 300, Radians(0), Radians(90))
	//dc.SetRGB(255, 188, 55)
	//dc.Fill()
	dc.Stroke()
	dc.SavePNG("out.png")
}
