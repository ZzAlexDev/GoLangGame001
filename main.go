package main

import (
	"image/color"
	"math"
	"time"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 600
	screenHeight = 300

	ballRadius = 15
)

type Game struct{}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

var (
	ballPositionX  = float64(screenWidth) / 2
	ballPositionY  = float64(screenHeight) / 2
	ballMovementX  = float64(0.0000006)
	ballMovementY  = float64(0.0000004)
	prevUpdateTime = time.Now()
)

func (g *Game) Update() error {
	timeDelta := float64(time.Since(prevUpdateTime))
	prevUpdateTime = time.Now()

	ballPositionX += ballMovementX * timeDelta
	ballPositionY += ballMovementY * timeDelta

	const minX = ballRadius
	const minY = ballRadius
	const maxX = screenWidth - ballRadius
	const maxY = screenHeight - ballRadius

	if ballPositionX >= maxX || ballPositionX <= minX {
		if ballPositionX > maxX {
			ballPositionX = maxX
		} else if ballPositionX < minX {
			ballPositionX = minX
		}

		ballMovementX *= -1
	}

	if ballPositionY >= maxY || ballPositionY <= minY {
		if ballPositionY > maxY {
			ballPositionY = maxY
		} else if ballPositionY < minY {
			ballPositionY = minY
		}

		ballMovementY *= -1
	}

	return nil
}

func (g *Game) drawCircle(screen *ebiten.Image, x, y, radius int, clr color.Color, fill bool) {
	radius64 := float64(radius)
	minAngle := math.Acos(1 - 1/radius64)

	for angle := float64(0); angle <= 360; angle += minAngle {
		xDelta := radius64 * math.Cos(angle)
		yDelta := radius64 * math.Sin(angle)

		x1 := int(math.Round(float64(x) + xDelta))
		y1 := int(math.Round(float64(y) + yDelta))

		if fill {
			if y1 < y {
				for y2 := y1; y2 <= y; y2++ {
					screen.Set(x1, y2, clr)
				}
			} else {
				for y2 := y1; y2 > y; y2-- {
					screen.Set(x1, y2, clr)
				}
			}
		}

		screen.Set(x1, y1, clr)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	purpleClr := color.RGBA{255, 0, 255, 255}

	x := int(math.Round(ballPositionX))
	y := int(math.Round(ballPositionY))

	g.drawCircle(screen, x, y, ballRadius, purpleClr, true)
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("The Moving Ball")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
