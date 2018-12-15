package main
import "fmt"
import "crypto/sha256"

func main() {
	fmt.Printf("x=%08b\tX=%08b\n", 'x', 'X')
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println("==================")
	ba := [3]int{1,2,3}
	fmt.Println("outer: ", ba)
	a(ba)
	fmt.Println("outer: ", ba)
	zero(&ba)
	fmt.Println("outer: ", ba)
}



func a( arr [3]int) {
	arr[1] = 4
	fmt.Println("inner: ", arr)
}

func zero(pa *[3]int) {
	for i, v := range pa { //pa == *pa auto derefenenced
		fmt.Println(i," ", v)
		pa[i] = 0 // pa[i] ==(*p)[i] auto dereferenced
	}
	fmt.Println("inner: ", *pa)
}