package main

import "math"

type Vector2 struct {
	x      float64
	y      float64
	degree float64
}

func vector2(x float64, y float64, degree ...float64) Vector2 {
	if len(degree) > 0 {
		return Vector2{x, y, degree[0]}
	} else {
		return Vector2{x, y, 0.0}
	}
}

func moveVector2Degree(vec Vector2, step float64) Vector2 {
	return vector2com(vec.degree + step)
}

func vector2com(degress float64) Vector2 {
	x := math.Cos(degress)
	y := math.Sin(degress)
	return vector2(x, y, degress)
}

func normalize(vec Vector2) Vector2 {
	length := math.Sqrt(((vec.x * vec.x) + (vec.y * vec.y)))
	x := vec.x / length
	y := vec.y / length
	return Vector2{x, y, vec.degree}
}
