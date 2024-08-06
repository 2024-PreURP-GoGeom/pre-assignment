package main

import "reflect"

var Real = GeneralSet{reflect.TypeOf(1.00)}

var Coord1D = GeneralSet{reflect.TypeOf(Coordinate{elements: []float64{3.14}})}

var Integer = GeneralSet{reflect.TypeOf(1)}
