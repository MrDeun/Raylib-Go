package main

import "math"

func Cylinder3D(size float64, vert_points, horz_points int32) Geometry {
	verticies := []Vector3{}
	for i := -vert_points; i < (vert_points); i++ {
		for j := -horz_points; j < (horz_points); j++ {
			angle := (float64(j) / float64(horz_points))
			v3 := vector3(
				size*math.Cos(angle),
				(float64(i)*size)/float64(vert_points),
				size*math.Sin(angle),
			)
			verticies = append(verticies, v3)
		}
	}
	return Geometry{position: vector3zero(), rotation: vector3zero(), verticies: verticies, lines: nil}
}
