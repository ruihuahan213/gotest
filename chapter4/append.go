package main
import "fmt"

func main() {
	/*
	arr := [6]int{1,2,3,4,5, 6}
	a := arr[:5]
	b := a[:6]
	fmt.Printf("the 6 element of the slice that have 5 elements totally: %d\n", b[5])

	fmt.Println(a)

	

	fmt.Println(appendInt(a, 7))

	fmt.Println(arr)

	*/

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1

	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	

	}

	z[len(x)] = y
	
	return z
}