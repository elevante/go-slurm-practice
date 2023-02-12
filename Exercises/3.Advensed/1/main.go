package main

import (
	"fmt"
	"math"
)

const pi float64 = 3.14

type Rectangle struct {
	a float64
	b float64
}

type Circle struct {
	r float64
}

func newCircle(a float64) *Circle {
	return &Circle{
		r: a,
	}
}

func newRectangle(a, b float64) *Rectangle {
	return &Rectangle{
		a: a,
		b: b,
	}
}

func (r *Rectangle) Area() float64 {
	return r.a * r.b
}

func (c *Circle) Area() float64 {
	return pi * math.Pow(c.r, 2)
}

func (r *Rectangle) Type() string {
	return "rectangle"
}

func (c *Circle) Type() string {
	return "circle"
}

func main() {
	rectangleNew := newRectangle(2.0, 4.0)
	circleNew := newCircle(2.0)

	fmt.Println("Square", rectangleNew.Type(), rectangleNew.Area())
	fmt.Println("Square", circleNew.Type(), circleNew.Area())

}
