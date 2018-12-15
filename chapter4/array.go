package main
import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1,2,3}
	var r [3]int = [3]int{1,2}
	fmt.Println(q)
	fmt.Println(r[2])

	qq := [...]int{1,2,3}
	fmt.Printf("%T\n", qq)

	//qqq := [3]int{1,2,3}
	//qqq = [...]int{1,2,3,4}

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "¥"}

	fmt.Println(symbol[RMB])

	a1 := [2]int{1, 2}
	b1 := [...]int{1, 2}
	c1 := [2]int{1, 3}
	fmt.Println(a1 == b1, a1 == c1, b1 == c1) // "true false false"

	//d1 := [3]int{1, 2}
	//fmt.Println(a1 == d1) //


}