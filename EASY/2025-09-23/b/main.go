package main

import "fmt"

func main() {
	var arr [8]int
	fmt.Scan(&arr[0], &arr[1], &arr[2], &arr[3], &arr[4], &arr[5], &arr[6], &arr[7])

	result := true

	for i := range arr[:7] {
		if arr[i] >= arr[i+1] {
			result = false
		}

		if arr[i] < 100 || 675 < arr[i] {
			result = false
		}

		if arr[i]%25 != 0 {
			result = false
		}
	}

	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
