package main

import raylib "github.com/gen2brain/raylib-go/raylib"
import (
	"math"
	"sort"
)

const (
	DRAWING = 0
	FOURIER = 1
	WIDTH = 1200
	HEIGHT = 800
)

var time float32 = 0
var userState = -1
var userDrawing []raylib.Vector2

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
	userDrawing = nil
	userState = DRAWING
	fourierPath = nil
}

func onMouseUp() {
	time = 0
	userState = FOURIER
	fourierTransforms = dft(userDrawing)
	sort.Slice(fourierTransforms, func(i, j int) bool {
		return fourierTransforms[i].amplitude > fourierTransforms[j].amplitude
	})
}

func update(dt float32) {
	if (raylib.IsMouseButtonPressed(raylib.MouseLeftButton)) {
		onMouseDown()
	}

	if (raylib.IsMouseButtonReleased(raylib.MouseLeftButton)) {
		onMouseUp()
	}

	switch userState {
		case DRAWING:
			userDrawing = append(userDrawing, raylib.GetMousePosition())
		
		case FOURIER:
			if (time > 2 * math.Pi) {
				fourierPath = nil
				time = 0
			}

			time += (2 * math.Pi) / float32(len(fourierTransforms))
	}
}

func drawEpicycles(x float64, y float64) raylib.Vector2 {

	for i := 0; i < len(fourierTransforms); i++ {
		var prevX = x + WIDTH / 2
		var prevY = y + HEIGHT / 2

		var k = fourierTransforms[i].frequency
		var r = fourierTransforms[i].amplitude
		var offset = fourierTransforms[i].phase

		x += r * math.Cos(float64(k) * float64(time) + offset)
		y += r * math.Sin(float64(k) * float64(time) + offset)

		raylib.DrawCircleLines(int32(math.Floor(prevX)), int32(math.Floor(prevY)), float32(r), raylib.LightGray)
	}

	return raylib.Vector2 { X: float32(x + WIDTH / 2), Y: float32(y + HEIGHT / 2) }
}

func draw() {
	raylib.BeginDrawing()
	defer raylib.EndDrawing()

	raylib.ClearBackground(raylib.Black)

	if (userState == DRAWING) {
		raylib.DrawLineStrip(userDrawing, int32(len(userDrawing)), raylib.RayWhite)
	}

	if (userState == FOURIER) {		
		var point = drawEpicycles(0, 0)
		fourierPath = append(fourierPath, point)
		raylib.DrawCircleV(point, 5, raylib.Green)
		raylib.DrawLineStrip(fourierPath, int32(len(fourierPath)), raylib.Green)
	}
}