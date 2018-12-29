package main

import "fmt"
import "strings"

func main() {
	f := square
	fmt.Printf("%T\n", f)
	fmt.Println( f(3) )
	fmt.Println("---------------")

	f = negative
	fmt.Printf("%T\n", f)
	fmt.Println( f(3) )
	fmt.Println("---------------")

	//f = product

	var ff func(int) int
	if ff != nil {
		ff(3)
	}

	fmt.Println(strings.Map(add1, "HAL-9000"))
	fmt.Println(strings.Map(add1, "VMS"))
	fmt.Println(strings.Map(add1, "123"))
}

func square( n int) int { return n * n}
func negative( n int) int { return -n}
func product( m, n int) int { return m * n }
func add1(r rune) rune { return r + 1 }