package main

import "fmt"

type pointer struct {
	x, y int
}


func main() {

	a := pointer{ x: 5, y: 6}

	fmt.Printf("a before: x=%d y=%d\n", a.x, a.y)
	b := scale(a, 2)

	fmt.Printf("a after: x=%d y=%d\n", a.x, a.y)
	fmt.Printf("the returned b : x=%d y=%d\n", b.x, b.y)

	c := pointer{1,2}
	d := pointer{1,2}

	fmt.Println(c==d)

	type Point struct {
		X, Y int
	}
	
	type Circle struct {
		//Center Point 
		Point
		Radius int
	}

	type Wheel struct {
		Circle 
		spokes int
	}

	var w Wheel
	//w.Circle.Center.X = 8
	w.X = 8	
	//w.Circle.Center.Y = 8
	w.Y = 8
	//w.Circle.Radius = 5
	w.Radius = 5
	w.spokes = 20

	w = Wheel{Circle{Point{8,8}, 5}, 20}

	fmt.Println(w)

	w = Wheel{
		Circle: Circle {
			Point: Point {
				X: 8,
				Y: 8,
			},
			Radius: 5,
		},
		spokes: 20,
	}

	fmt.Printf("%#v\n", w)



}


func scale(p pointer, factor int) pointer {
	p.x  *= factor
	p.y  *= factor
	fmt.Printf("inter: x=%d y=%d\n", p.x, p.y)

	return p
}