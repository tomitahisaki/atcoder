package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	list := make([]int, n)
	for i := range list {
		fmt.Scan(&list[i])
	}

	count := 0
	for {
		for i := range list {
			if list[i]%2 != 0 {
				fmt.Println(count)
				return
			}
			list[i] /= 2
		}
		count++
	}
}
