package main
import "fmt"


func main() {
	months := [...]string{
		1: "January",
		2: "February",
		3: "March",
		4: "April",
		5: "May",
		6: "June",
		7: "July",
		8: "August",
		9: "September",
		10: "October",
		11: "November",
		12: "December",
	}
	fmt.Println(months[0])
	fmt.Println(months)
	fmt.Println(months[7])

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Printf("Q2's length=%d, capacity=%d\n", len(Q2), cap(Q2))
	fmt.Println(summer)
	fmt.Printf("summer's length=%d, capacity=%d\n", len(summer), cap(summer))

	fmt.Println("----------------------")
	fmt.Printf("befor modified: Q2[2]=%s summer[0]=%s\n", Q2[2], summer[0])
	months[6] = "六月"
	fmt.Printf("after modified: Q2[2]=%s summer[0]=%s\n", Q2[2], summer[0])

	for _, s := range summer {
		for _,  q := range Q2 {
			if s == q {
				fmt.Println(s)
			}
		}
	}

	//fmt.Println(summer[:20])
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)
	fmt.Println("======================")
	rev(months[1:13])
	fmt.Println(months)
}

func rev(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}