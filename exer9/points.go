package exer9

import (
	"fmt"
	"math"
)

// TODO: The Point struct, NewPoint function, .String and .Norm methods

type Point struct {
   x float64
   y float64
}

// String representation
func NewPoint (valX float64, valY float64) Point{
	return Point{
		x : valX, 
		y : valY,
	}
}

func (point Point) String() string {
	return fmt.Sprintf("(%v, %v)", point.x, point.y)
}

func (point Point) Norm() float64{
	return math.Sqrt(point.x*point.x + point.y*point.y)
}

func (point *Point) Scale(f float64) {
	point.x = point.x * f
	point.y = point.y * f
}

func (point *Point) Rotate(angle float64) {
	valueX := point.x
	point.x = valueX * math.Cos(angle) - point.y * math.Sin(angle)
	point.y = valueX * math.Sin(angle) + point.y * math.Cos(angle)
}
