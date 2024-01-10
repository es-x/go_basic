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
		switch {
		case i&1 != 0 && s%2 == 0:
			fmt.Println(strings.Repeat("# ", s/2))
		case i&1 != 0 && s%2 != 0:
			fmt.Println(strings.Repeat("# ", s/2) + "#")
		case i&1 != 1 && s%2 == 0:
			fmt.Println(strings.Repeat(" #", s/2))
		case i&1 != 1 && s%2 != 0:
			fmt.Println(strings.Repeat(" #", s/2) + " ")
		}
	}
}
