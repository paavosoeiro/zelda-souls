package main

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func (v Vec2) Sum(vec Vec2) Vec2 {
	return Vec2{
		X: v.X + vec.X,
		Y: v.Y + vec.Y,
	}
}

func (v Vec2) magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2) Normalize() Vec2 {
	magnitude := v.magnitude()
	if magnitude == 0 {
		return Vec2{0, 0}
	}
	return Vec2{v.X / magnitude, v.Y / magnitude}
}

func (v Vec2) Scale(factor float64) Vec2 {
	return Vec2{v.X * factor, v.Y * factor}
}
