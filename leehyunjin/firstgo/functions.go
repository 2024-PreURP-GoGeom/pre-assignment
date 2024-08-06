package main

import (
	"math"
	"reflect"
)

type Function struct {
	domain      GeneralSet
	codomain    GeneralSet
	maprelation func(x Coordinate) Matrix
}

func (F Function) Computation(v Coordinate) Matrix {
	// ok := F.domain.Contains(v)
	ok := true
	if !ok {
		return Matrix{}
	}
	return F.maprelation(v)
}

var FunctionSpace = GeneralSet{reflect.TypeOf(exfunc1)}

func Diffrel(F Function, dom GeneralSet) func(x Coordinate) Matrix {
	return func(x Coordinate) Matrix {
		// IsSafe 이후 수정
		if true {
			point := x
			domainDim := len(point.elements)
			totalderiv := make([]Coordinate, domainDim)

			for i := range point.elements {
				// fmt.Println(i, "dimension!!")
				infinitesimal := make([]float64, domainDim)
				infinitesimal[i] = epsilon
				epsilonCoord := Coordinate{elements: infinitesimal}

				slope := (F.Computation(point.Addition(epsilonCoord)).Addition(F.Computation(point).AdditiveInverse())).ScalarMultiplication(bigNum)
				totalderiv[i] = Coordinate{slope.elements[0].elements}
			}

			return Matrix{elements: totalderiv}
		}
		return Matrix{}
	}
}

func Differential(F Function) Function {
	var Derivative Function
	Derivative.domain = F.domain
	Derivative.codomain = F.codomain
	Derivative.maprelation = Diffrel(F, F.domain)
	return Derivative
}

var Diff Function = Function{
	domain:      FunctionSpace,
	codomain:    FunctionSpace,
	maprelation: Differential,
}

func IsSafe(input any, domain GeneralSet) bool {
	ok := domain.Contains(input)
	if !ok {
		return false
	}
	return true
}

func Exmapping(input Coordinate) Matrix {
	return Matrix{elements: []Coordinate{{elements: []float64{math.Cos(input.elements[0])}}}}
}

func Exmapping2(input Coordinate) Matrix {
	x := input.elements[0]
	y := input.elements[1]
	return Matrix{elements: []Coordinate{{elements: []float64{x*x + y*y}}}}
}

var exfunc1 Function = Function{
	domain:      Real,
	codomain:    Real,
	maprelation: Exmapping}

var exfunc2 Function = Function{
	domain:      GeneralSet{},
	codomain:    GeneralSet{},
	maprelation: Exmapping2,
}
