package main

import "fmt"

type square struct{
	sideLength float64
}
type triangle struct{
	height float64
}

type shape interface{
	getArea() float64
}

func main() {
	mySquare := square{
		sideLength: 5.0,
	}

	myTriangle := triangle{
		height: 3.5,
	}

	printArea(mySquare)
	printArea(myTriangle)

}

func printArea(s shape) {
	area := s.getArea()
	fmt.Println(area)
}

func (s square) getArea() float64 {
	area := s.sideLength * s.sideLength
	return area
}

func (t triangle) getArea() float64 {
	base := 10.0
	area := (base * t.height) / 2
	return area
}