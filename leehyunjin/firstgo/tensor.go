package main

type Tensor interface {
	Rank() int
	Size() []int
}

//type Vector interface { // Vector Space T에서의 벡터
//	Addition(any) any
//	ZeroVector() any
//	ScalarMultiplication(scalar float64) any // R-Vector space
//}
