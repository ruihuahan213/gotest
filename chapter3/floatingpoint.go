package main

import "fmt"
import "math"

func main() {

	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1) // "true"!

	f=100e2

	fmt.Println(f)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d eA = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 0+Inf InfNaN"
}