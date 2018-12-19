package main
import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	a := []string{"aaa","","ccc","ddd", "","fff"}
	fmt.Println(a)
    b := nonempty(a) // usually we use a = nonempty(a)
	fmt.Println(b)
	fmt.Println(a)

	a[0] = "modified"
	fmt.Println(a)
	fmt.Println(b)





}