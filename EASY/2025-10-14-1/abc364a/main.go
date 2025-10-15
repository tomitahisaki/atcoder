package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	list := make([]string, n)

	for i := range list {
		fmt.Scan(&list[i])
	}

	for i := 0; i < len(list)-1; i++ {
		if list[i] == "sweet" && list[i+1] == "sweet" {
			if i+2 == len(list) {
				break
			} else {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
