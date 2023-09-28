package main

import raylib "github.com/gen2brain/raylib-go/raylib"
import (
	"fmt"
	"math"
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
}

func onMouseUp() {
	time = 0
	userState = FOURIER
	fourierTransforms = dft(userDrawing)
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
				time = 0
			}

			time += (2 * math.Pi) / float32(len(fourierTransforms))
	}
}

// func calculatePhase() int {
// 	if (userState == FOURIER) {
// 		return int(math.Floor(float64(fourierTime / userTime) * float64(len(userDrawing)))) % len(userDrawing)
// 	}

// 	return len(userDrawing)
// }

func drawEpicycles(x int32, y int32) {

	
	
	// Draw arbitrary circles
	var max = 5
	for i := 0; i < max; i++ {
		var n = 2 * i + 1
		var r float64 = 100 * ( 4 / (float64(n) * math.Pi))
		raylib.DrawCircleLines(x + WIDTH / 2, y + HEIGHT / 2, float32(r), raylib.RayWhite)

		x += int32(math.Floor(r * math.Cos(float64(n) * float64(time))))
		y += int32(math.Floor(r * math.Sin(float64(n) * float64(time))))
	}

	// Draw circles in fourierTransforms
	// for i := 0; i < len(fourierTransforms	); i++ {
	// 	var ft = fourierTransforms[i]
	// 	var n = 2 * i + 1
	// 	var r float64 = 100 * ( 4 / (float64(n) * math.Pi))
	// 	raylib.DrawCircleLines(x + WIDTH / 2, y + HEIGHT / 2, float32(r), raylib.RayWhite)

	// 	x += int32(math.Floor(r * math.Cos(float64(n) * float64(time) + ft.phase)))
	// 	y += int32(math.Floor(r * math.Sin(float64(n) * float64(time) + ft.phase)))
	// }
}

func draw() {
	raylib.BeginDrawing()
	defer raylib.EndDrawing()

	raylib.ClearBackground(raylib.Black)

	if (userState == DRAWING) {
		raylib.DrawLineStrip(userDrawing, int32(len(userDrawing)), raylib.RayWhite)
	}

	if (userState == FOURIER && len(fourierTransforms) > 0) {
		
		raylib.DrawText(fmt.Sprint("drawing: ", len(userDrawing)), 100, 400, 48, raylib.RayWhite)
		raylib.DrawText(fmt.Sprint("transforms: ", len(fourierTransforms)), 100, 200, 48, raylib.RayWhite)

		drawEpicycles(0, 0)
	}
}