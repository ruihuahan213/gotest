package main
import (
	"fmt"
	"os"
	"log"

	"gotest/chapter5/links"
)

func main() {
	fmt.Println(os.Args[1:])
	breadthFirst(crawl, os.Args[1:])
}

func breadthFirst( f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
	//fmt.Println(worklist)
}



func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	//fmt.Println("in the crawl, the list: ", list)
	return list
}