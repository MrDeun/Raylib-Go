package main

func Cube3D(size float64) Geometry {
	vertices_arr := []Vector3{
		vector3(0.5*size, 0.5*size, -0.5*size),
		vector3(0.5*size, -0.5*size, -0.5*size),
		vector3(-0.5*size, -0.5*size, -0.5*size),
		vector3(-0.5*size, 0.5*size, -0.5*size),

		vector3(0.5*size, 0.5*size, 0.5*size),
		vector3(0.5*size, -0.5*size, 0.5*size),
		vector3(-0.5*size, -0.5*size, 0.5*size),
		vector3(-0.5*size, 0.5*size, 0.5*size),
	}

	lines_arr := []Line3D{
		createLine3D(&vertices_arr[0], &vertices_arr[1]),
		createLine3D(&vertices_arr[1], &vertices_arr[2]),
		createLine3D(&vertices_arr[2], &vertices_arr[3]),
		createLine3D(&vertices_arr[3], &vertices_arr[0]),

		createLine3D(&vertices_arr[4], &vertices_arr[5]),
		createLine3D(&vertices_arr[5], &vertices_arr[6]),
		createLine3D(&vertices_arr[6], &vertices_arr[7]),
		createLine3D(&vertices_arr[7], &vertices_arr[4]),

		createLine3D(&vertices_arr[0], &vertices_arr[4]),
		createLine3D(&vertices_arr[1], &vertices_arr[5]),
		createLine3D(&vertices_arr[3], &vertices_arr[7]),
		createLine3D(&vertices_arr[2], &vertices_arr[6]),
	}

	triangles_arr := [][3]*Vector3{
		{&vertices_arr[0], &vertices_arr[1], &vertices_arr[2]},
		{&vertices_arr[0], &vertices_arr[3], &vertices_arr[2]},

		{&vertices_arr[4], &vertices_arr[5], &vertices_arr[6]},
		{&vertices_arr[4], &vertices_arr[7], &vertices_arr[6]},
	}

	geo := Geometry{
		position:  vector3zero(),
		rotation:  vector3zero(),
		lines:     lines_arr,
		verticies: vertices_arr,
		triangles: triangles_arr,
	}

	geo.moveZ(2.0)
	return geo
}
