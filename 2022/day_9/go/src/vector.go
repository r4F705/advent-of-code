package main

import "math"

type Vector2 struct {
	X int
	Y int
}

func (v *Vector2) Distance(other Vector2) float64 {
	return math.Sqrt(math.Pow(float64(v.X-other.X), 2) + math.Pow(float64(v.Y-other.Y), 2))
}

func (v *Vector2) Difference(other Vector2) Vector2 {
	return Vector2{v.X - other.X, v.Y - other.Y}
}
