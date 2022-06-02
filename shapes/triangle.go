package shapes

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t *Triangle) CalcPerimeter() float64 {
	return 3 * t.Side
}

func (t *Triangle) CalcArea() float64 {
	return t.Side * math.Sqrt(3) / 4
}