package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	result := a * b

	if result%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
