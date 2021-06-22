package main

import (
	"log"
	"math"
)

func main() {
	triangle := TriangularObject{
		height: 1,
		base:   2,
		length: 3,
	}
	rectangle := RectangularObject{
		height: 1,
		length: 2,
		base:   4,
	}
	cylinder := CylinderObject{
		height: 4,
		radius: 5,
	}

	// x is a configuration parameter that can be set via a json config file. Depend
	// on its value or a combination a different feature object would be chosen
	var x interface{} = "cyc"
	var object Object

	switch x.(type) {
	case int:
		object = triangle
	case float64:
		object = rectangle
	default:
		object = cylinder
	}

	volumecalculator := VolumneCalculator{object}
	volume := volumecalculator.Volume()
	log.Printf("volume of %T = %f \n", object, volume)

}

type Object interface {
	Perimeter() float64
	Height() float64
}
type VolumneCalculator struct {
	object Object
}

func (calculator VolumneCalculator) Volume() float64 {
	return calculator.object.Perimeter() * calculator.object.Height() //bridge that links feature to interface
}

//implementation of an Object: Triangle
type TriangularObject struct {
	height float64
	base   float64
	length float64
}

func (t TriangularObject) Perimeter() float64 {
	return t.height * t.base / 2
}

func (t TriangularObject) Height() float64 {
	return t.length
}

//implementation of an Object: Rectangle
type RectangularObject struct {
	height float64
	length float64
	base   float64
}

func (r RectangularObject) Perimeter() float64 {
	return r.height * r.base * 2
}

func (r RectangularObject) Height() float64 {
	return r.length
}

// implementation of an Object: Cylinder
type CylinderObject struct {
	height float64
	radius float64
}

func (c CylinderObject) Perimeter() float64 {
	return math.Pi*2*c.radius*c.height + 2*math.Pi*c.radius*c.height
}

func (c CylinderObject) Height() float64 {
	return c.height
}
