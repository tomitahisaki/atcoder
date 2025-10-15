package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, p, q, r, s int
	fmt.Scan(&n, &p, &q, &r, &s)

	list := make([]int, n)
	for i := range list {
		fmt.Scan(&list[i])
	}

	// adjust index to 0-based
	p--
	q--
	r--
	s--

	result := swap(list, p, q, r, s)
	strList := make([]string, len(result))
	for i := range result {
		strList[i] = fmt.Sprintf("%d", result[i])
	}

	spacedString := strings.Join(strList, " ")

	fmt.Println(spacedString)
}

func swap(list []int, p, q, r, s int) []int {
	left := list[:p]
	part1 := list[p : q+1]
	middle := list[q+1 : r]
	part2 := list[r : s+1]
	right := list[s+1:]

	result := append([]int{}, left...)
	result = append(result, part2...)
	result = append(result, middle...)
	result = append(result, part1...)
	result = append(result, right...)

	return result
}
