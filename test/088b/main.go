package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	list := make([]int, n)
	for i := range list {
		fmt.Scan(&list[i])
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i] > list[j]
	})

	alice, bob := 0, 0

	for i := range list {
		if i%2 == 0 {
			alice += list[i]
		} else {
			bob += list[i]
		}
	}

	fmt.Println(alice - bob)
}
