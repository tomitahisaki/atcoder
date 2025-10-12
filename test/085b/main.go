package main

import (
	"fmt"
	"slices"
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

	uniquedList := slices.Compact(list)

	fmt.Println(len(uniquedList))
}
