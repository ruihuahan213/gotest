package main

import "fmt"

func main() {

	var u uint8 = 255
	fmt.Println( u, u + 1, u * u)

	var i int8 = 127
	fmt.Println(i, i + 1, i * i )

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	
	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}
	fmt.Printf("%08b\n", x&y) // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y) // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y) // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
		fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}

	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i]) // "bronze", "silver", "gold"
	}	


	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii) // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 D 'D'"
	fmt.Printf("%d %[1]q\n", newline) // "10 '\n'"
}