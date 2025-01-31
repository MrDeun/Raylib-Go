package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Geometry struct {
	position Vector3
	rotation Vector3

	verticies []Vector3
	lines     []Line3D
	triangles [][3]*Vector3
}

func (geo *Geometry) Control(step float64) {
	if rl.IsKeyDown(rl.KeyUp) {
		geo.rotateX(step)
	}
	if rl.IsKeyDown(rl.KeyDown) {
		geo.rotateX(-step)
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		geo.rotateY(-step)
	}
	if rl.IsKeyDown(rl.KeyRight) {
		geo.rotateY(step)
	}

	if rl.IsKeyDown(rl.KeyA) {
		geo.moveX(-step * 100)
	}
	if rl.IsKeyDown(rl.KeyD) {
		geo.moveX(step * 100)
	}
	if rl.IsKeyDown(rl.KeyW) {
		geo.moveY(-step * 100)
	}
	if rl.IsKeyDown(rl.KeyS) {
		geo.moveY(step * 100)
	}

}

func (geo *Geometry) moveX(step float64) {
	for i := range geo.verticies {
		geo.verticies[i].x += step
	}
}

func (geo *Geometry) moveY(step float64) {
	for i := range geo.verticies {
		geo.verticies[i].y += step
	}
}

func (geo *Geometry) moveZ(step float64) {
	for i := range geo.verticies {
		geo.verticies[i].z += step
	}
}
func (geo *Geometry) rotateX(step float64) {
	mat := rollMatrix(step)
	for i := range geo.verticies {
		x := geo.verticies[i].x
		y := geo.verticies[i].y
		z := geo.verticies[i].z

		geo.verticies[i].x = (mat[0][0] * x) + (mat[0][1] * y) + (mat[0][2] * z)
		geo.verticies[i].y = (mat[1][0] * x) + (mat[1][1] * y) + (mat[1][2] * z)
		geo.verticies[i].z = (mat[2][0] * x) + (mat[2][1] * y) + (mat[2][2] * z)
	}
}

func (geo *Geometry) rotateY(step float64) {
	mat := pitchMatrix(step)
	for i := range geo.verticies {
		x := geo.verticies[i].x
		y := geo.verticies[i].y
		z := geo.verticies[i].z

		geo.verticies[i].x = (mat[0][0] * x) + (mat[0][1] * y) + (mat[0][2] * z)
		geo.verticies[i].y = (mat[1][0] * x) + (mat[1][1] * y) + (mat[1][2] * z)
		geo.verticies[i].z = (mat[2][0] * x) + (mat[2][1] * y) + (mat[2][2] * z)
	}
}

func (geo *Geometry) rotateZ(step float64) {
	mat := yawMatrix(step)
	for i := range geo.verticies {
		x := geo.verticies[i].x
		y := geo.verticies[i].y
		z := geo.verticies[i].z

		geo.verticies[i].x = (mat[0][0] * x) + (mat[0][1] * y) + (mat[0][2] * z)
		geo.verticies[i].y = (mat[1][0] * x) + (mat[1][1] * y) + (mat[1][2] * z)
		geo.verticies[i].z = (mat[2][0] * x) + (mat[2][1] * y) + (mat[2][2] * z)
	}
}
