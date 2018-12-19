package main
import "fmt"

func main() {
	a := []string{"aaa", "",  "ccc"}
	a = nonempty2(a)
	fmt.Println(a)
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out,s)
		}
	}

	return out
}