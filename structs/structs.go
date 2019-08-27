package structs

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns the perimeter of a shape
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Area foobar
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle struct
type Circle struct {
	Radius float64
}

const PI = 3.1416

// Perimeter on circle
func (c Circle) Perimeter() float64 {
	return 2.0 * PI * c.Radius
}

// Area on circle
func (c Circle) Area() float64 {
	return PI * c.Radius * c.Radius
}

// Triangle struct
type Triangle struct {
	base   float64
	height float64
}

// Area of a Triangle
func (triangle Triangle) Area() float64 {
	return .5 * triangle.height * triangle.base
}

// Shape interface
type Shape interface {
	Area() float64
}
