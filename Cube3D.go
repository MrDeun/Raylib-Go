package main

func Cube3D(size float64) Geometry {
	arr := []Vector3{
		vector3(0.5*size, 0.5*size, -0.5*size),
		vector3(0.5*size, -0.5*size, -0.5*size),
		vector3(-0.5*size, -0.5*size, -0.5*size),
		vector3(-0.5*size, 0.5*size, -0.5*size),

		vector3(0.5*size, 0.5*size, 0.5*size),
		vector3(0.5*size, -0.5*size, 0.5*size),
		vector3(-0.5*size, -0.5*size, 0.5*size),
		vector3(-0.5*size, 0.5*size, 0.5*size),
	}

	geo := Geometry{
		position: vector3zero(),
		rotation: vector3zero(),

		verticies: arr,
	}

	geo.moveZ(2.0)
	return geo
}
