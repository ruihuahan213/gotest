package main
import "fmt"
import "unicode/utf8"

func main() {
	s := "国"
	fmt.Println(len(s))
	fmt.Printf("%0x %0x %0x\n", s[0], s[1], s[2])
	fmt.Printf("%c %c\n", s[0], s[len(s) - 1])
	fmt.Printf("%s\n", s)	

	s = "left foot"
	t := s
	s += ", right foot"
	fmt.Println(s)
	fmt.Println(t)

	//s[0] = 'L'
	fmt.Println("\xe5\x9b\xbd\n")

	fmt.Println(`then return \n not returnd
but returned here`)


	ss := "世界"
	for _, v := range []byte(ss) {
		fmt.Printf("%x ", v)		
	}
	fmt.Println()

	for _, v := range ss {
		fmt.Printf("%06x ", v)		
	}
	fmt.Println()

	fmt.Println("\xe4\xb8\x96\xe7\x95\x8c\n")
	fmt.Println("\u4e16\u754c\n")
	fmt.Println("\U00004e16\U0000754c\n")

	ss = "hello, 世界"
	fmt.Println(len(ss));
	fmt.Println(utf8.RuneCountInString(ss))

	for i := 0; i < len(ss); {
		r, size := utf8.DecodeRuneInString(ss[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	n := 0
	for _, _ = range ss {
		n++
	}
	
	fmt.Println(n)

	SSS := "我就是世界"
	fmt.Printf("the length: %d\n", len(SSS))
	fmt.Printf("% x\n", SSS)
	fmt.Println(SSS)
	rrr := []rune(SSS)
	fmt.Printf("the length of rune: %d\n", len(rrr))
	fmt.Printf("%x\n", rrr)
	fmt.Println(string(rrr))
	fmt.Printf("% x\n", string(rrr))

	fmt.Println(string(65))
	fmt.Println(string(0x4eac))
	fmt.Printf("% x\n", string(0x4eac))

	fmt.Println(string(1234567))
}