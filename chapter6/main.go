package main

import (
	"fmt"
	"gotest/chapter6/geometry"
	"image/color"
	"gotest/chapter6/intset"
)

type P *int

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}
	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))

	perim := geometry.Path{
		{1,1},
		{5,1},
		{5,4},
		{1,1},
	}

	fmt.Println(perim.Distance())
	fmt.Println(geometry.PathDistance(perim))

	fmt.Println("original p: ", p.X, p.Y)
	p.ScaleBy2(2) // the receiver passed to the method is a value type
	fmt.Println("p after scaled by value-receiver method", p.X, p.Y)
	
	// compiler performed an implicte reference, such as &p 
	p.ScaleBy(2) // the receiver passed to the method is a pointer type
	fmt.Println("p after scaled by pointer-receiver method", p.X, p.Y)
	(&p).ScaleBy(2) // the formal invocation method
	fmt.Println("p after scaled by pointer-receiver method", p.X, p.Y)
	fmt.Printf("the type of ScaleBy method: %T\n", p.ScaleBy)

	fmt.Println("--------Color Point-----------------")
	var cp geometry.ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{255,0,0,255}
	blue := color.RGBA{0,0,255,255}
	var pp = geometry.ColoredPoint{ geometry.Point{1,1}, red}
	var qq = geometry.ColoredPoint{ geometry.Point{5, 4}, blue}
	fmt.Println(pp.Distance(qq.Point))

	pp.ScaleBy(2)
	qq.ScaleBy(2)
	fmt.Println(pp.Distance(qq.Point))

	fmt.Println("----------------method values and expression------------")
	ppp := geometry.Point{1, 2}
	qqq := geometry.Point{4, 6}

	distanceFromP := ppp.Distance // get from a instance of the struct
	fmt.Println(distanceFromP(qqq))

	var origin geometry.Point
	fmt.Println(distanceFromP(origin))

	scaleP := ppp.ScaleBy
	scaleP(2)
	fmt.Println(ppp)
	scaleP(3)
	fmt.Println(ppp)	
	scaleP(10)
	fmt.Println(ppp)

	fmt.Println("----------------method values and expression------------")

	pppp := geometry.Point{1, 2}
	qqqq := geometry.Point{4, 6}

	distance := geometry.Point.Distance // get from the struct type
	fmt.Println(distance(pppp, qqqq))
	fmt.Printf("%T\n", distance)

	scale := (*geometry.Point).ScaleBy
	scale(&pppp, 2)
	fmt.Println(pppp)
	fmt.Printf("%T\n", scale);

	fmt.Println("---------------- intset------------")

	var x, y intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	fmt.Println("the string from &x ",&x)
	fmt.Println(x)
}

// pinter type can not have method
//func (p P) f() {
	
//}