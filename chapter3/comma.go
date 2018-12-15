package main

import (
	"fmt"
	"os"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Println(" Wrong number of arguments")
		os.Exit(1)
	}

	fmt.Println(comma(os.Args[1]))

}b

func comma(s string) string {
	n := len(s)
	if n <=3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}