package main

type LinearTransformation interface {
	Multiplication(Coordinate) Coordinate
}

func (m1 Matrix) Multiplication(m2 Matrix) Matrix {
	return MatrixMultiplication(m1, m2)
}
