package main

import (
	"math"
	"reflect"
)

type Function struct {
	domain      GeneralSet
	codomain    GeneralSet
	maprelation func(x any) any
}

func (F Function) Computation(x any) any {
	ok := F.domain.Contains(x)
	if !ok {
		return nil
	}
	return F.maprelation(x)
}

var FunctionSpace = GeneralSet{reflect.TypeOf(exfunc1)}

func Diffrel(F Function) func(x any) any {
	return func(x any) any {
		if IsSafe(x, Real) {
			point := x.(float64)
			slope := ((F.Computation(point + epsilon)).(float64) - (F.Computation(point)).(float64)) / epsilon
			return slope
		}
		return nil
	}
}

func Differential(F Function) Function {
	var Derivative Function
	Derivative.domain = F.domain
	Derivative.codomain = F.codomain
	Derivative.maprelation = Diffrel(F)
	return Derivative
}

func DefiniteIntegral(F Function, a, b float64) float64 {
	if IsSafe(a, Real) && IsSafe(b, Real) {
		n := bigNum
		h := (b - a) / n
		result := 0.0
		for i := 0; i <= int(n); i++ {
			x := a + float64(i)*h
			fx := F.Computation(x).(float64)
			if i == 0 || i == int(n) {
				result += fx / 2
			} else {
				result += fx
			}
		}
		return result * h
	}
	return math.NaN()
}

func IsSafe(input any, domain GeneralSet) bool {
	ok := domain.Contains(input)
	if !ok {
		return false
	}
	return true
}

func Exmapping(input any) any {
	return math.Cosh(input.(float64))
}

var exfunc1 Function = Function{
	domain:      Real,
	codomain:    Real,
	maprelation: Exmapping}
