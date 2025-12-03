package main

import "fmt"

type Shape interface {
	area() float32
	perimeter() float32
}

type Rectangle struct {
	length, breadth float32
}

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

func (c Circle) area() float32 {
	return 3.14 * c.radius * c.radius
}

func (s Square) area() float32 {
	return s.side * s.side
}

func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

func (r Rectangle) perimeter() float32 {
	return 2 * (r.length + r.breadth)
}

func calculate(s Shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}

func main() {

	rect := Rectangle{7, 4}

	calculate(rect)

	// sq := Square{5}
	// calculate(sq)

	// circ := Circle{3}
	// calculate(circ)

}
