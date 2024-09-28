package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sphere3D struct {
	rotation Vector3
	position Vector3

	verticies []Vector3
}

const fullCircle = math.Pi * 2

func controlSphere(sphere *Sphere3D, step float64) {
	if rl.IsKeyDown(rl.KeyDown) {
		sphere.roll(step)
	}
	if rl.IsKeyDown(rl.KeyUp) {
		sphere.roll(-step)
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		sphere.pitch(-step)
	}
	if rl.IsKeyDown(rl.KeyRight) {
		sphere.pitch(step)
	}
}

func sphere3D(size float64, vertical_points int, horizontal_point int) Sphere3D {
	var array = []Vector3{}
	for i := vertical_points; i > -vertical_points; i-- {
		for j := horizontal_point; j > -horizontal_point; j-- {
			angle := fullCircle / float64(j)
			veritce := vector3(
				math.Cos(angle)*float64(size)*float64(j),
				float64(i)*float64(size),
				math.Sin(angle)*float64(size)*float64(j))
			array = append(array, veritce)
		}
	}

	return Sphere3D{rotation: vector3zero(), position: vector3zero(), verticies: array}
}

func (sphere *Sphere3D) roll(step float64) {
	mat := rollMatrix(step)
	for i := range sphere.verticies {
		x := sphere.verticies[i].x
		y := sphere.verticies[i].y
		z := sphere.verticies[i].z

		sphere.verticies[i].x = (mat[0][0] * x) + (mat[0][1] * y) + (mat[0][2] * z)
		sphere.verticies[i].y = (mat[1][0] * x) + (mat[1][1] * y) + (mat[1][2] * z)
		sphere.verticies[i].z = (mat[2][0] * x) + (mat[2][1] * y) + (mat[2][2] * z)
	}
}
func (sphere *Sphere3D) yaw(step float64) {
	mat := yawMatrix(step)
	for i := range sphere.verticies {
		x := sphere.verticies[i].x
		y := sphere.verticies[i].y
		z := sphere.verticies[i].z

		sphere.verticies[i].x = (mat[0][0] * x) + (mat[0][1] * y) + (mat[0][2] * z)
		sphere.verticies[i].y = (mat[1][0] * x) + (mat[1][1] * y) + (mat[1][2] * z)
		sphere.verticies[i].z = (mat[2][0] * x) + (mat[2][1] * y) + (mat[2][2] * z)
	}
}
func (sphere *Sphere3D) pitch(step float64) {
	mat := pitchMatrix(step)
	for i := range sphere.verticies {
		x := sphere.verticies[i].x
		y := sphere.verticies[i].y
		z := sphere.verticies[i].z

		sphere.verticies[i].x = (mat[0][0] * x) + (mat[0][1] * y) + (mat[0][2] * z)
		sphere.verticies[i].y = (mat[1][0] * x) + (mat[1][1] * y) + (mat[1][2] * z)
		sphere.verticies[i].z = (mat[2][0] * x) + (mat[2][1] * y) + (mat[2][2] * z)
	}
}

func RunSphere() {
	var sphere = sphere3D(15.0, 128, 128)
	const step = 0.01
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		for _, v := range sphere.verticies {
			fmt.Printf("STAT: Point at ( %f, %f )\n", v.x, v.y)
			rl.DrawCircle(int32(v.x+centerX), int32(v.y+centerY), 1.0, rl.White)
		}
		rl.EndDrawing()
		//sphere.roll(step)
		controlSphere(&sphere, step)
	}

}
