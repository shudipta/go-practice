package main

import (
	"fmt";
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (c *Circle) perimeter() float64 {
	return 2.0 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2.0 * (l + w)
}

type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var peri float64
	for _, s := range shapes {
		peri += s.perimeter()
	}
	return peri
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func (m *MultiShape) perimeter() float64 {
	var peri float64
	for _, s := range m.shapes {
		peri += s.perimeter()
	}
	return peri
}

func main() {
	r := Rectangle{0, 0, 10, 10}
	c := Circle{y: 2, x: 1, r: 3}

	var x int
	var y string
	fmt.Println(x, y, c)

	m := MultiShape{ []Shape{&r, &c}}
	//fmt.Println(r.area())
	//fmt.Println(c.area())
	fmt.Println(r.area(), c.area())
	fmt.Println(totalArea(&m))
	fmt.Println(r.perimeter(), c.perimeter())
	fmt.Println(totalPerimeter(&m))
	fmt.Println(c)
}