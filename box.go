package golang_united_school_homework

import (
	"errors"
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return errors.New("It goes out of the shapesCapacity range")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if ok, e := b.checkIndex(i); ok == false && e != nil {
		return nil, e
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if ok, e := b.checkIndex(i); ok == false && e != nil {
		return nil, e
	}
	shape := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if ok, e := b.checkIndex(i); ok == false && e != nil {
		return nil, e
	}
	b.shapes[i] = shape
	return shape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	sum = 0

	for _, v := range b.shapes {
		sum += v.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	sum = 0
	for _, v := range b.shapes {
		sum += v.CalcArea()
	}
	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	exist := false
	for i, v := range b.shapes {
		if isCircle(v) == true {
			b.ExtractByIndex(i)
			exist = true
		}
	}

	if exist == true {
		return nil
	} else {
		return errors.New("Circles are not exist in the list")
	}
}

func (b *box) checkIndex(index int) (bool, error) {
	if index >= b.shapesCapacity {
		return false, errors.New("Index went out of the range")
	}

	if index >= len(b.shapes) {
		return false, errors.New("Shape by index doesn't exist")
	}

	return true, nil
}

func isCircle(t interface{}) bool {
	switch t.(type) {
	case Circle:
		return true
	default:
		return false
	}
}

func main() {
	c := Circle{Radius: 10}
	fmt.Println(c)
}
