package main

import "math"

type Matrix3 [3][3]float64

func yawMatrix(degree float64) Matrix3 {
	var tempTable = [3][3]float64{}
	tempTable[0][0] = math.Cos(degree)
	tempTable[0][1] = -math.Sin(degree)
	tempTable[0][2] = 0.0

	tempTable[1][0] = math.Sin(degree)
	tempTable[1][1] = math.Cos(degree)
	tempTable[1][2] = 0.0

	tempTable[2][0] = 0.0
	tempTable[2][1] = 0.0
	tempTable[2][2] = 1.0
	return tempTable
}
func pitchMatrix(degree float64) Matrix3 {
	var tempTable = [3][3]float64{}
	tempTable[0][0] = math.Cos(degree)
	tempTable[0][1] = 0.0
	tempTable[0][2] = math.Sin(degree)

	tempTable[1][0] = 0.0
	tempTable[1][1] = 1.0
	tempTable[1][2] = 0.0

	tempTable[2][0] = -math.Sin(degree)
	tempTable[2][1] = 0
	tempTable[2][2] = math.Cos(degree)
	return tempTable
}
func rollMatrix(degree float64) Matrix3 {
	var tempTable = [3][3]float64{}
	tempTable[0][0] = 1.0
	tempTable[0][1] = 0.0
	tempTable[0][2] = 0.0

	tempTable[1][0] = 0.0
	tempTable[1][1] = math.Cos(degree)
	tempTable[1][2] = -math.Sin(degree)

	tempTable[2][0] = 0
	tempTable[2][1] = math.Sin(degree)
	tempTable[2][2] = math.Cos(degree)
	return tempTable
}
