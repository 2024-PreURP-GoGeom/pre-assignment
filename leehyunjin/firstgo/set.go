package main

import "reflect"

type Set interface {
	Contains(any) bool
}

type FiniteSet interface {
	Len(FiniteSet) int
}

type Iterable interface {
	Len() int
}

type GeneralSet struct {
	ElementType reflect.Type
}

func (N GeneralSet) Contains(x any) bool {

	if reflect.TypeOf(x) == N.ElementType {
		return true
	}
	return false
} // GeneralSet 구조체는 전부 Set 인터페이스의 조건을 만족한다.
