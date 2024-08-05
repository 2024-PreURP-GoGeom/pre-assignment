package main

import (
	"fmt"
)

const epsilon float64 = 0.00000000001
const bigNum float64 = 10000

func main() {
	var x = 5.00
	var y = 5
	c1 := []float64{1.00, 5.00, 7.56}
	c2 := []float64{2.00, 5.00, 8.45}
	v1 := Coordinate{elements: c1}
	v2 := Coordinate{elements: c2}
	var m1 Matrix = Matrix{elements: []Coordinate{v1, v2}}
	var m2 Matrix = Matrix{elements: []Coordinate{v2, v1}}

	fmt.Println(Real.Contains(x))
	fmt.Println(Real.Contains(y))
	fmt.Println(Integer.Contains(y))
	fmt.Println(exfunc1.maprelation(5.00))
	fmt.Println(FunctionSpace.Contains(exfunc1))
	fmt.Println(exfunc1.Computation(5))
	fmt.Println(Differential(exfunc1).Computation(5.00))
	fmt.Println(DefiniteIntegral(exfunc1, 1, 5))
	fmt.Println(v1.Addition(v2).elements)
	fmt.Println(v1.AdditiveInverse().elements)
	fmt.Println(m1.elements)
	fmt.Println(m2.elements)
	fmt.Println(m1.Addition(m2))
	fmt.Println(m2.AdditiveInverse())
	fmt.Println(m1.Size())
	fmt.Println(m1.Transpose())
	fmt.Println(m1.MatrixMultiplication(m2.Transpose()))
}
