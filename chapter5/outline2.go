package main
import ( 
	"fmt"
	"os"
	"net/http"
	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks3: %v\n", err)
			os.Exit(1)
		}
/*
		for _, link := range links {
			fmt.Println(link)
		}
*/
	}
}

func findLinks(url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()

	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	//return visit(nil, doc), nil

	// where is the slice []strig created? why can we passed a nil?

	forEachNode(doc, startElement, endElement)

	return nil
}

func visit( links []string, n *html.Node) []string {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
		
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c:= n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int
func startElement( n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<<%s>>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement( n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s<</%s>>\n", depth*2, "", n.Data)
	}
}


