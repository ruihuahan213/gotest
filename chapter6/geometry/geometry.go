package geometry

import "math"
import "image/color"

type Point struct{ X, Y float64 }
type Path []Point
type ColoredPoint struct {
	Point
	Color color.RGBA
}


func Distance(p, q Point) float64 {
	return math.Hypot( q.X - p.X, q.Y - p.Y)
}

func (p Point) Distance( q  Point) float64 {
	return math.Hypot( q.X - p.X, q.Y - p.Y )
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}

	return sum
}

func PathDistance(path Path) float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i].Distance(path[i-1])
		}
	}

	return sum
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) ScaleBy2(factor float64 ) {

	p.X *= factor
	p.Y *= factor
}

