package main
import (
	"math"
	"fmt"
)


func main(){
	fmt.Println( hypot(3,4))

	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", zero)
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

// Like parameters, results may be named. In that case, each name declares a local variable
// initialized to the zero value for its type.

func add(x int, y int) int { return x +y }
func sub( x, y int) ( z int ) { z = x + y; return z }
func first(x int, _ int) int { return x}
func zero( int, int) int { return 0}

// Go has no concept of default parameter values, nor any way to secify arguments
// by name, so names of parameters and results don't matter to the caller except
// documentation. 

// Parameters are local variables within the body of the function, with its 
// initial values set to the arguments supplied by the caller.

// Arguments are passed by value, so the function receives a copy of each 
// arguments;
// However, if the argument contains some kind of reference, like a pointer,
// slice, map, function, or channel, then the caller may affected by any 
// modifications the function makes to variables indirectly refered to by
// the argument. (notice: struct is not included in the reference types)

