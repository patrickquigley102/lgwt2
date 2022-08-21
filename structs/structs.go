// Package structs for learnings
package structs

import "math"

const two = 2

// Shape is cool.
type Shape interface {
	Area()
}

// Rectangle is cool.
type Rectangle struct {
	Width, Height float64
}

// Perimeter of a rectangle.
func (r Rectangle) Perimeter() float64 {
	return two * (r.Width + r.Height)
}

// Area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle is cool.
type Circle struct {
	Radius float64
}

// Area of circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
