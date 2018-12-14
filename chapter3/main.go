package main

import (
	"os"
	"fmt"
	"gotest/chapter3/basename1"
)

func main() {

	ss := os.Args[1]
	sss := basename1.BaseName(ss)
	fmt.Println(sss)
}

 