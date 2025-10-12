package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	var result int
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "1" {
			result++
		}
	}

	fmt.Println(result)
}
