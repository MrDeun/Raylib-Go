package main

import rl "github.com/gen2brain/raylib-go/raylib"

func PlayCube3D() {
	cube := Cube3D(256)
	step := 0.01
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		for _, v := range cube.verticies {
			rl.DrawCircle(int32(centerX+v.x), int32(centerY+v.y), 4.0, rl.White)
		}
		rl.EndDrawing()
		cube.Control(step)
	}
}

func PlaySphere3D() {
	cyl := Cylinder3D(64, 16, 16)
	step := 0.01

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		for _, v := range cyl.verticies {
			rl.DrawCircle(int32(centerX+v.x), int32(centerY+v.y), 4.0, rl.White)
		}
		rl.EndDrawing()
		cyl.Control(step)
	}
}
