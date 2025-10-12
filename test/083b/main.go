package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	sum := 0
	for i := 1; i <= n; i++ {
		result := digitSum(i)

		if result >= a && result <= b {
			sum += i
		}
	}
	fmt.Println(sum)
}

func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
