package main

import rl "github.com/gen2brain/raylib-go/raylib"
import "fmt"
import "math"

const pointCount = 60

var buttonBuffer = 0

func wheelInput(speed *float64) {
	if rl.IsKeyPressed(rl.KeySpace) {
		(*speed) *= -1.0
	} else if rl.IsKeyDown(rl.KeyRight) {
		if buttonBuffer == 10 {
			(*speed) += 0.1
			buttonBuffer = 0
		} else {
			buttonBuffer++
		}
	} else if rl.IsKeyDown(rl.KeyLeft) {
		if buttonBuffer == 10 {
			(*speed) -= 0.1
			buttonBuffer = 0
		} else {
			buttonBuffer++
		}
	} else if rl.IsKeyPressed(rl.KeyR) {
		*speed = 1.0
	}
}

func RunWheel() {
	var arrayVec []Vector2

	for i := 0; i < pointCount; i++ {
		degree := (2 * (math.Pi)) / pointCount * float64(i)
		arrayVec = append(arrayVec, vector2com(degree))
	}

	const radius = 300.0
	speed := 1.0

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		wheelInput(&speed)
		rl.ClearBackground(rl.Black)
		for i, v := range arrayVec {
			rl.DrawCircle(int32(v.x*radius)+centerX, int32(v.y*radius)+centerY, 10.0, rl.Red)
			arrayVec[i] = moveVector2Degree(arrayVec[i], 0.01*speed)
		}
		speedText := fmt.Sprintf("Speed: %.2f", speed)
		rl.DrawText(speedText, 10.0, 20*1, 20, rl.Gray)
		rl.DrawText("LeftArrow - Decrease Speed", 10.0, 20*2, 20, rl.Gray)
		rl.DrawText("RightArrow - Increase Speed", 10.0, 20*3, 20, rl.Gray)
		rl.DrawText("R - reset speed", 10.0, 20*4, 20, rl.Gray)
		rl.DrawText("Spacebar - reverse direction", 10.0, 20*5, 20, rl.Gray)

		rl.EndDrawing()
	}
}
