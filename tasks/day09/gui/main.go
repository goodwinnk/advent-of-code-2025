package main

import (
	"AdventOfCode2025/internal/util"
	"AdventOfCode2025/tasks/day09"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenW = 1000
	screenH = 1000
)

type game struct {
	points []day09.Point
}

func (g *game) Update() error {
	// No updates needed for a static line.
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	// Fill background
	screen.Fill(color.RGBA{R: 20, G: 24, B: 28, A: 255})

	low := 248
	lowLeft := 217
	lowMax := 29
	highMax := 470

	prev := g.points[len(g.points)-1]
	for _, p := range g.points[1:] {
		x1 := float32(prev.X) / 100
		y1 := float32(prev.Y) / 100
		x2 := float32(p.X) / 100
		y2 := float32(p.Y) / 100

		vector.StrokeLine(screen, x1, y1, x2, y2, 1, color.White, false)

		prev = p
	}

	for _, p := range g.points {
		x2 := float32(p.X) / 100
		y2 := float32(p.Y) / 100

		clr := color.RGBA{R: 0x99, G: 0x00, B: 0x00, A: 0xff}
		if p.Id == low {
			clr = color.RGBA{R: 0x00, G: 0xdd, B: 0x00, A: 0xff}
		}

		if p.X >= g.points[low].X {
			clr = color.RGBA{R: 0xdd, G: 0xdd, B: 0xff, A: 0xff}
		}

		if p.X < g.points[lowLeft].X {
			clr = color.RGBA{R: 0xdd, G: 0xdd, B: 0xff, A: 0xff}
		}

		if p.Id == lowLeft || p.Id == low || p.Id == lowMax {
			clr = color.RGBA{R: 0x00, G: 0xdd, B: 0xff, A: 0xff}
		}

		vector.DrawFilledCircle(screen, x2, y2, 2, clr, false)
	}

	lowLine := float32(g.points[lowMax].Y / 100)
	vector.StrokeLine(screen, 0, lowLine, 1000, lowLine, 1, color.White, false)

	highLine := float32(g.points[highMax].Y / 100)
	vector.StrokeLine(screen, 0, highLine, 1000, highLine, 1, color.White, false)

	lowDx := g.points[low].X - g.points[lowLeft].X
	lowDy := g.points[lowLeft].Y - g.points[low].Y
	vector.DrawFilledRect(screen,
		float32(g.points[lowLeft].X)/100, float32(g.points[low].Y)/100,
		float32(lowDx)/100, float32(lowDy)/100,
		color.RGBA{R: 0x22, G: 0x00, B: 0x22, A: 0x22}, false)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenW, screenH
}

func main() {
	input := util.Input(9)
	points, err := day09.Parse(input)
	if err != nil {
		panic(err)
	}

	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Day 09 — Show Lines")
	if err := ebiten.RunGame(&game{
		points: points,
	}); err != nil {
		panic(err)
	}
}
