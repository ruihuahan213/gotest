package main
import (
	"fmt"
)

func main() {
	values := []int{1,2,3,4}
	fmt.Println(sum(values...))
	fmt.Printf("%T\n", sum)

	fmt.Println(values)
}

func sum( vals ...int) int {
	var total int 
	for _, val := range vals {
		total += val
	}

	return total
}