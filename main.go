package main

// TODO - add placeholder circles for Glyphs at each step ont he perimeter of the Colloquy

import (
	gglyphs "github.com/jibdugaw/galifreyan/glyphs"

	"github.com/fogleman/gg"
)

func main() {
	//input := "gallifrey falls no more"
	input := "Doctor it is so much bigger on the inside"
	//input := "i am the doctor"

	// our globals which will be the canvas size and then some basic starting point stuff
	canvasSide := float64(1000)
	originX := canvasSide / 2
	originY := canvasSide / 2

	// TODO - retrieve sentences and create colloquy's for each
	sentence := gglyphs.NewColloquy(originX, originY, float64(0.8)*originX, input)

	dc := gg.NewContext(1000, 1000)
	// using originX, originY, radius and offsets for each glyph in the colloquy
	// using originX, originY and radius of colloquy draw each word glyph
	dc.SetLineWidth(10)
	sentence.Draw(dc)
	//DrawB(dc)
	//dc.DrawArc(400, 325, 300, Radians(0), Radians(90))
	//dc.SetRGB(255, 188, 55)
	//dc.Fill()
	dc.Stroke()
	dc.SavePNG("out.png")
}
