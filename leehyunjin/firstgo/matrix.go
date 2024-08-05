package main

import "fmt"

type Matrix struct {
	elements []Coordinate
}

func (m Matrix) Size() [2]int {
	if len(m.elements) == 0 {
		return [2]int{0, 0}
	}
	return [2]int{
		len(m.elements),
		len(m.elements[0].elements),
	}
}

func (m1 Matrix) Addition(m2 Matrix) Matrix {
	var m3 Matrix
	c1 := m1.elements
	c2 := m2.elements
	dim := len(m1.elements)
	c3 := make([]Coordinate, dim)
	if m1.Size() == m2.Size() {
		for i := 0; i <= dim-1; i++ {
			c3[i] = c1[i].Addition(c2[i])
		}
		m3 = Matrix{elements: c3}
		return m3
	}
	return Matrix{} //panic으로 바꿀 것.
}

func (m1 Matrix) AdditiveInverse() Matrix {
	var minv Matrix
	c1 := m1.elements
	dim := len(m1.elements)
	cinv := make([]Coordinate, dim)
	for i := 0; i <= dim-1; i++ {
		cinv[i] = c1[i].AdditiveInverse()
	}
	minv = Matrix{elements: cinv}
	return minv
}

func (m1 Matrix) ZeroVector() Matrix {
	return m1.Addition(m1.AdditiveInverse())
}

func (m1 Matrix) ScalarMultiplication(scalar float64) Matrix {
	var ms Matrix
	c1 := m1.elements
	dim := len(m1.elements)
	cs := make([]Coordinate, dim)
	for i := 0; i <= dim-1; i++ {
		cs[i] = c1[i].ScalarMultiplication(scalar)
	}
	ms = Matrix{elements: cs}
	return ms
}

func (m1 Matrix) Transpose() Matrix {
	var transpose Matrix
	newdim := m1.Size()[1]
	cs := make([]Coordinate, newdim)
	for i := 0; i <= newdim-1; i++ {
		cs[i] = Coordinate{elements: make([]float64, m1.Size()[0])}
		for j := 0; j <= m1.Size()[0]-1; j++ {
			cs[i].elements[j] = m1.elements[j].elements[i]
		}
	}
	transpose = Matrix{elements: cs}
	return transpose
}

func (m1 Matrix) MatrixMultiplication(m2 Matrix) Matrix {
	// 행렬곱 가능한지 체크
	fmt.Println(m1.Size(), m2.Size())
	if m1.Size()[1] != m2.Size()[0] {
		panic("matrices size incompatible!!!")
	}

	c1 := m1.elements
	c2 := m2.Transpose().elements
	newSize := [2]int{m1.Size()[0], m2.Size()[1]}

	cm := make([]Coordinate, newSize[0])

	for i := 0; i < newSize[0]; i++ {
		cm[i] = Coordinate{elements: make([]float64, newSize[1])}
		for j := 0; j < newSize[1]; j++ {
			cm[i].elements[j] = c1[i].InnerProduct(c2[j])
		}
	}
	return Matrix{elements: cm}
}
