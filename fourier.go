package main

import (
	"math"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	WIDTH  = 1200
	HEIGHT = 800
)

const (
	DRAWING = 0
	FOURIER = 1
)

var time float32 = 0
var state = -1

var userPath []raylib.Vector2

var fourierTransforms []DiscreteFourierTransform
var fourierPath []raylib.Vector2

func main() {
	raylib.InitWindow(WIDTH, HEIGHT, "Epicycles")
	defer raylib.CloseWindow()

	raylib.SetTargetFPS(60)

	for !raylib.WindowShouldClose() {
		dt := raylib.GetFrameTime()
		update(dt)
		draw()
	}
}

func onMouseDown() {
	state = DRAWING
	userPath = nil
	fourierPath = nil
}

func onMouseUp() {
	time = 0
	state = FOURIER
	fourierTransforms = dft(userPath)
}

func update(dt float32) {
	if raylib.IsMouseButtonPressed(raylib.MouseLeftButton) {
		onMouseDown()
	}

	if raylib.IsMouseButtonReleased(raylib.MouseLeftButton) {
		onMouseUp()
	}
}

func drawEpicycles(x float64, y float64) raylib.Vector2 {
	for i := 0; i < len(fourierTransforms); i++ {
		var prevX = x + WIDTH/2
		var prevY = y + HEIGHT/2

		var freq = fourierTransforms[i].frequency
		var radius = fourierTransforms[i].amplitude
		var phi = fourierTransforms[i].phase

		x += radius * math.Cos(freq*float64(time)+phi)
		y += radius * math.Sin(freq*float64(time)+phi)

		raylib.DrawCircleLines(int32(math.Floor(prevX)), int32(math.Floor(prevY)), float32(radius), raylib.LightGray)
	}

	return raylib.Vector2{X: float32(x + WIDTH/2), Y: float32(y + HEIGHT/2)}
}

func draw() {
	raylib.BeginDrawing()
	defer raylib.EndDrawing()

	raylib.ClearBackground(raylib.Black)

	switch state {
	case DRAWING:
		userPath = append(userPath, raylib.GetMousePosition())
		raylib.DrawLineStrip(userPath, int32(len(userPath)), raylib.RayWhite)

	case FOURIER:
		if time > 2*math.Pi {
			fourierPath = nil
			time = 0
		}

		time += (2 * math.Pi) / float32(len(fourierTransforms))

		fourierPath = append(fourierPath, drawEpicycles(0, 0))
		raylib.DrawLineStrip(fourierPath, int32(len(fourierPath)), raylib.Green)
	}
}
