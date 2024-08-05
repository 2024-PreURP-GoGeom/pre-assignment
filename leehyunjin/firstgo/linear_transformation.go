package main

type LinearTransformation interface {
	Transform(Coordinate) Coordinate
}

func (m1 Matrix) Transform(m2 Matrix) Matrix {
	return m1.MatrixMultiplication(m2)
}
