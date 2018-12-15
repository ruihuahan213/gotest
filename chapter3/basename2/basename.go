package basename2

import "strings"
import "fmt"

func BaseName(s string) string {

	fmt.Println("Output by basename version 2")
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >=0 {
		s  = s[:dot]
	} 

	return s
}