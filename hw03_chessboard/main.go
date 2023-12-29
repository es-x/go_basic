package main

import (
	"fmt"
	"strings"
)

func main() {
	var s int

	fmt.Printf("Введите размер сетки: ")
	fmt.Scanf("%d", &s)

	for i := 1; i <= s; i++ {
		if i&1 != 0 {
			if s%2 == 0 {
				fmt.Println(strings.Repeat("# ", s/2))
			} else {
				fmt.Println(strings.Repeat("# ", s/2) + "#")
			}

		} else {
			if s%2 == 0 {
				fmt.Println(strings.Repeat(" #", s/2))
			} else {
				fmt.Println(strings.Repeat(" #", s/2) + " ")
			}
		}

	}
}
