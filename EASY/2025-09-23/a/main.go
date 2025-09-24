package main

import (
	"fmt"
)

func main() {
	var ab, ac, bc string

	if _, err := fmt.Scan(&ab, &ac, &bc); err != nil {
		return
	}

	a := 0
	b := 0
	c := 0

	if ab == ">" {
		a++
	} else {
		b++
	}

	if ac == ">" {
		a++
	} else {
		c++
	}

	if bc == ">" {
		b++
	} else {
		c++
	}

	if a == 1 {
		fmt.Println("A")
	} else if b == 1 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
