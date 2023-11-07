package main

import (
	"fmt"
)

type Shape interface {
	GetType() string
	Accept(Visitor)
}

// 1 тип
type square struct {
	side int
}

func NewSquare(side int) *square {
	return &square{
		side: side,
	}
}

func (s *square) Accept(v Visitor) {
	v.VisitForSquare(s)
}

func (s *square) GetType() string {
	return "Square"
}

// 2 тип
type circle struct {
	radius int
}

func NewCircle(radius int) *circle {
	return &circle{
		radius: radius,
	}
}

func (c *circle) Accept(v Visitor) {
	v.VisitForCircle(c)
}

func (c *circle) GetType() string {
	return "Circle"
}

// 3 тип
type Triangle struct {
	a int
	b int
	c int
}

func NewTriangle(a, b, c int) *Triangle {
	return &Triangle{
		a: a,
		b: b,
		c: c,
	}
}

func (t *Triangle) Accept(v Visitor) {
	v.VisitForTriangle(t)
}

func (t *Triangle) GetType() string {
	return "Triangle"
}

type Visitor interface {
	VisitForSquare(Shape)
	VisitForCircle(Shape)
	VisitForTriangle(Shape)
}

type areaCalculator struct {
	area int
}

func NewAreaCalculator() *areaCalculator {
	return &areaCalculator{}
}

func (a *areaCalculator) VisitForSquare(s Shape) {
	sid := s.(*square).side
	a.area = sid * sid
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) VisitForCircle(s Shape) {
	r := s.(*circle).radius
	a.area = r * r
	fmt.Println("Calculating area for circle")
}

func (a *areaCalculator) VisitForTriangle(s Shape) {
	// Вычисляем площадь для треугольника. После вычисления площади присваиваем её в
	// переменную area экземпляра
	fmt.Println("Calculating area for triangle")
}

func main() {
	square := NewSquare(2)
	circle := NewCircle(3)
	triangle := NewTriangle(1, 2, 3)

	areaCalculator := NewAreaCalculator()
	square.Accept(areaCalculator)
	fmt.Println(areaCalculator.area)

	circle.Accept(areaCalculator)
	fmt.Println(areaCalculator.area)

	triangle.Accept(areaCalculator)

	fmt.Println()
}
