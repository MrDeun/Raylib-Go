package main

import "math"

type Sphere3D struct {
	rotation Vector3
	position Vector3

	verticies []Vector3
}

const fullCircle = math.Pi * 2

func sphere3D(size int, vertical_points int, horizontal_point int) Sphere3D {
	var array = []Vector3{}
	fltSize := float64(size)
	for i := 0; i < vertical_points; i++ {
		for j := 0; j < horizontal_point; j++ {
			angle := fullCircle / float64(j)
			veritce := vector3(
				math.Cos(angle)*fltSize*float64(i),
				float64(i)*fltSize,
				math.Sin(angle)*fltSize*float64(i))
			array = append(array, veritce)
		}
	}

	return Sphere3D{rotation: vector3zero(), position: vector3zero(), verticies: array}
}
