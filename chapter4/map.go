package main
import "fmt"

func main() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	/* ages := map[string]int{
		"alice": 31,
		"charlie": 34,
	} */
	
	fmt.Println(ages)
	fmt.Println(ages["alice"])

	delete(ages, "alice")
	fmt.Println(ages)

}