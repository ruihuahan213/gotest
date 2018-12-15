package main

import "fmt"
import "time"
import "math"

const noDelay time.Duration = 0
const timeout = 5 * time.Minute

const (
	a  = 1
	b = 2
	c 
	d
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wenesday
	Thursday
	Friday
	Saturday
)

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcase
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

const (
	deadbeef = 0xdeadbeef // untyped int with value 3735928559
	a1 = uint32(deadbeef) // uint32 with value 3735928559
	b1 = float32(deadbeef) // float32 with value 3735928576 (rounded up)
	c1 = float64(deadbeef) // float64 with value 3735928559 (exact)
	//d1 = int32(deadbeef) // compile error: constant overflows int32
	//e1 = float64(1e309) // compile error: constant overflows float64
	//f1 = uint(-1)
	// compile error: constant underflows uint
	)



func main() {
	fmt.Println(math.Pi)

	fmt.Printf("%08b\n", FlagUp)
	fmt.Printf("%08b\n", FlagBroadcase)
	fmt.Printf("%08b\n", FlagLoopback)
	fmt.Printf("%08b\n", FlagPointToPoint)
	fmt.Printf("%08b\n", FlagMulticast)

	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)

	fmt.Println( a, b, c, d )

	fmt.Println(Tuesday)

	var f float64 = 212
	fmt.Println( (f - 32) * 5 / 9 )
	fmt.Println( 5 / 9 * (f - 32))
	fmt.Println( 5.0 / 9.0 * (f - 32))

	fmt.Printf("%T\n", 0) // "int"
	fmt.Printf("%T\n", 0.0) // "float64"
	fmt.Printf("%T\n", 0i) // "complex128"
	fmt.Printf("%T\n", '\000') // "int32" (rune)
}