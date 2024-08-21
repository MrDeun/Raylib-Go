package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cube3D struct {
	position Vector3
	rotation Vector3
	vertices [8]Vector3
}

func cube3D(size float64) Cube3D {
	position := vector3zero()
	rotation := vector3zero()

	var vertices = [8]Vector3{}
	vertices[0] = vector3(0.5*size, 0.5*size, 0.5*size)
	vertices[1] = vector3(-0.5*size, 0.5*size, 0.5*size)
	vertices[2] = vector3(0.5*size, -0.5*size, 0.5*size)
	vertices[3] = vector3(-0.5*size, -0.5*size, 0.5*size)
	vertices[4] = vector3(0.5*size, 0.5*size, -0.5*size)
	vertices[5] = vector3(-0.5*size, 0.5*size, -0.5*size)
	vertices[6] = vector3(0.5*size, -0.5*size, -0.5*size)
	vertices[7] = vector3(-0.5*size, -0.5*size, -0.5*size)

	for i := range vertices {
		vertices[i].z += 2.0
	}

	return Cube3D{position: position, rotation: rotation, vertices: vertices}
}

func cube3Dyaw(cube *Cube3D, step float64) {
	cube.rotation = vector3addscalar(cube.rotation, step)
	yaw := yawMatrix(step)
	for i := range cube.vertices {
		x := cube.vertices[i].x
		y := cube.vertices[i].y
		z := cube.vertices[i].z

		cube.vertices[i].x = (yaw[0][0] * x) + (yaw[0][1] * y) + (yaw[0][2] * z)
		cube.vertices[i].y = (yaw[1][0] * x) + (yaw[1][1] * y) + (yaw[1][2] * z)
		cube.vertices[i].z = (yaw[2][0] * x) + (yaw[2][1] * y) + (yaw[2][2] * z)
	}
}
func cube3Dpitch(cube *Cube3D, step float64) {
	cube.rotation = vector3addscalar(cube.rotation, step)
	pitch := pitchMatrix(step)
	for i := range cube.vertices {
		x := cube.vertices[i].x
		y := cube.vertices[i].y
		z := cube.vertices[i].z

		cube.vertices[i].x = (pitch[0][0] * x) + (pitch[0][1] * y) + (pitch[0][2] * z)
		cube.vertices[i].y = (pitch[1][0] * x) + (pitch[1][1] * y) + (pitch[1][2] * z)
		cube.vertices[i].z = (pitch[2][0] * x) + (pitch[2][1] * y) + (pitch[2][2] * z)
	}
}

func cube3Droll(cube *Cube3D, step float64) {
	cube.rotation = vector3addscalar(cube.rotation, step)
	roll := rollMatrix(step)
	for i := range cube.vertices {
		x := cube.vertices[i].x
		y := cube.vertices[i].y
		z := cube.vertices[i].z

		cube.vertices[i].x = (roll[0][0] * x) + (roll[0][1] * y) + (roll[0][2] * z)
		cube.vertices[i].y = (roll[1][0] * x) + (roll[1][1] * y) + (roll[1][2] * z)
		cube.vertices[i].z = (roll[2][0] * x) + (roll[2][1] * y) + (roll[2][2] * z)
	}
}

func cubeInput(step float64, direction Vector2) (float64, Vector2) {
	if rl.IsKeyDown(rl.KeyRight) {
		step += 0.001
	} else if rl.IsKeyDown(rl.KeyLeft) {
		step -= 0.001
	} else if rl.IsKeyPressed(rl.KeyR) {
		step = 0.001
		direction = vector2(1.0, 1.0)
	} else if rl.IsKeyPressed(rl.KeySpace) {
		step *= -1.0
	} else if rl.IsKeyPressed(rl.KeyW) {
		direction.y += 0.1
	} else if rl.IsKeyPressed(rl.KeyS) {
		direction.y -= 0.1
	} else if rl.IsKeyPressed(rl.KeyA) {
		direction.x += 0.1
	} else if rl.IsKeyPressed(rl.KeyD) {
		direction.x -= 0.1
	}

	return step, direction
}

func RunCube() {
	var cube = cube3D(300.0)
	var step = 0.001
	direction := vector2(1.0, 1.0)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		step, direction = cubeInput(step, direction)
		rl.ClearBackground(rl.Black)
		for _, v := range cube.vertices {
			rl.DrawCircle(int32(v.x+centerX), int32(v.y+centerY), float32(v.z)*0.1+25.0, rl.Red)
		}
		cube3Dpitch(&cube, step*direction.x)
		cube3Dyaw(&cube, step*direction.y)

		speedText := fmt.Sprintf("Speed: %.04f", step)
		vectorText := fmt.Sprintf("Directional Vector: (%f,%f)", direction.x, direction.y)
		rl.DrawText(vectorText, 10.0, 20*1, 20, rl.Gray)
		rl.DrawText(speedText, 10.0, 20*2, 20, rl.Gray)
		rl.DrawText("LeftArrow - Decrease Speed", 10.0, 20*3, 20, rl.Gray)
		rl.DrawText("RightArrow - Increase Speed", 10.0, 20*4, 20, rl.Gray)
		rl.DrawText("R - reset speed", 10.0, 20*5, 20, rl.Gray)
		rl.DrawText("Spacebar - reverse direction", 10.0, 20*6, 20, rl.Gray)

		rl.EndDrawing()
	}
}
