package shapes

import "math"

// Shape defines a generic shape interface
type Shape interface {
	Area() float64
}

// Rectangle defines shape with width/height representing a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates metric for rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle defines shape with radius representing a circle
type Circle struct {
	Radius float64
}

// Area calculates metric for circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle defines shape with base/height representing a triangle
type Triangle struct {
	Base   float64
	Height float64
}

// Area calculates metric for triangle
func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

// Perimeter returns rectangle perimeter based on width/height
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
