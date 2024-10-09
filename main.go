package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const WINDOW_WIDTH = 1600
const WINDOW_HEIGHT = 900

const centerX = WINDOW_WIDTH / 2
const centerY = WINDOW_HEIGHT / 2

func choiceInput() rune {
	if rl.IsKeyPressed(rl.KeyOne) {
		return '1'
	} else if rl.IsKeyPressed(rl.KeyTwo) {
		return '2'
	} else if rl.IsKeyPressed(rl.KeyThree) {
		return '3'
	} else {
		return 'n'
	}
}

func main() {
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Raylib testing!")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	chosen := false
	for !chosen && !rl.WindowShouldClose() {
		choice := choiceInput()
		switch choice {
		case '1':
			RunWheel()
			chosen = true
		case '2':
			PlayCube3D()
			chosen = true
		case '3':
			PlaySphere3D()
			chosen = true
		default:
			rl.BeginDrawing()
			rl.ClearBackground(rl.Black)
			rl.DrawText("Test app - pick a program", centerX, centerY-10.0, 10.0, rl.White)
			rl.DrawText("1. Wheel Of Fortune", centerX, centerY, 10.0, rl.White)
			rl.DrawText("2. 2D Rotaing Cube", centerX, centerY+10.0, 10.0, rl.White)
			rl.DrawText("3. 2D Rotating Cylinder (WIP)", centerX, centerY+20.0, 10.0, rl.White)
			rl.EndDrawing()
		}
	}

}
