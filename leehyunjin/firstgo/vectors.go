package main

import (
	"reflect"
)

var Euclidean1DSpace = GeneralSet{reflect.TypeOf([1]float64{3.14})}

type Coordinate struct {
	elements []float64
}

type Tuple[T any] struct {
	elements []T
} // 아예 모든 것을 튜플로 정의해버리자.

func (c Coordinate) Len() int {
	return len(c.elements)
}

// Coordinate is a Tensor
func (c Coordinate) Size() []int {
	return []int{c.Len()}
}
func (c Coordinate) Rank() int {
	return 1
}

func (v1 Coordinate) Addition(v2 Coordinate) Coordinate {
	var v3 Coordinate
	c1 := v1.elements
	c2 := v2.elements
	if v1.Len() == v2.Len() {
		dim := v1.Len()
		// fmt.Println("Dimension of coord", dim)
		c3 := make([]float64, dim)
		for i := 0; i <= dim-1; i++ {
			c3[i] = c1[i] + c2[i]
		}
		v3 = Coordinate{elements: c3}
		return v3
	}
	return Coordinate{} //panic으로 바꿀 것.
}

func (v1 Coordinate) AdditiveInverse() Coordinate {
	var vinv Coordinate
	c1 := v1.elements
	dim := v1.Len()
	cinv := make([]float64, dim)
	for i := 0; i <= dim-1; i++ {
		cinv[i] = -c1[i]
	}
	vinv = Coordinate{elements: cinv}
	return vinv
}

func (v1 Coordinate) Zerovector() Coordinate {
	return v1.Addition(v1.AdditiveInverse())
}

func (v1 Coordinate) ScalarMultiplication(scalar float64) Coordinate {
	var vs Coordinate
	c1 := v1.elements
	dim := v1.Len()
	cs := make([]float64, dim)
	for i := 0; i <= dim-1; i++ {
		cs[i] = c1[i] * scalar
	}
	vs = Coordinate{elements: cs}
	return vs
}

func (v1 Coordinate) InnerProduct(v2 Coordinate) float64 {
	var innerprod float64 = 0
	c1 := v1.elements
	c2 := v2.elements
	dim := v1.Len()

	for i := 0; i <= dim-1; i++ {
		innerprod += c1[i] * c2[i]
	}
	return innerprod
}
