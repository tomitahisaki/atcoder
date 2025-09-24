package main

import "fmt"

func main() {
	var arr [8]int
	fmt.Scan(&arr[0], &arr[1], &arr[2], &arr[3], &arr[4], &arr[5], &arr[6], &arr[7])

	for i := range arr[:7] {
		if arr[i] < 100 || 675 < arr[i] || arr[i]%25 != 0 {
			fmt.Println("No")
			return
		}

		if arr[i] > arr[i+1] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
