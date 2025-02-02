package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

z_buffer := []float32{}

func PlayCube3D() {
	cube := Cube3D(512)
	step := 0.01
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		for _, t := range cube.triangles {
			temp1 := rl.Vector2{float32(t[0].x + centerX), float32(t[0].y + centerY)}
			temp2 := rl.Vector2{float32(t[1].x + centerX), float32(t[1].y + centerY)}
			temp3 := rl.Vector2{float32(t[2].x + centerX), float32(t[2].y + centerY)}
			rl.DrawTriangle(temp1, temp2, temp3, rl.Red)
			rl.DrawTriangle(temp3, temp2, temp1, rl.Red)
		}
		for _, v := range cube.verticies {
			rl.DrawCircle(int32(centerX+v.x), int32(centerY+v.y), 4.0, rl.White)
		}
		for _, l := range cube.lines {
			l.Draw()
		}

		rl.EndDrawing()
		cube.Control(step)
	}
}

func PlaySphere3D() {
	cyl := Cylinder3D(256, 4, 4)
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
