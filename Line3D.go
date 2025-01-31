package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Line3D struct {
	begin *Vector3
	end   *Vector3
}

func createLine3D(_begin, _end *Vector3) Line3D {
	new_line := Line3D{}
	new_line.begin = _begin
	new_line.end = _end
	return new_line
}

func (_line *Line3D) findEquation() (float64, float64) {
	A := (_line.end.y - _line.begin.y) / (_line.end.x - _line.begin.x)
	B := _line.end.y - A*_line.end.x
	return A, B
}

func (_line *Line3D) Draw() {
	rl.DrawLine(
		int32(_line.begin.x+centerX),
		int32(_line.begin.y+centerY),
		int32(_line.end.x+centerX),
		int32(_line.end.y+centerY),
		rl.White,
	)
}

func (_line *Line3D) doesPointBelong(x, y, equation_A, equation_B float64) bool {
	if x >= _line.begin.x && _line.end.x >= x {
		if y >= _line.begin.y && _line.end.y >= y {
			if y == equation_A*x+equation_B {
				return true
			}
		}
	}

	return false
}
