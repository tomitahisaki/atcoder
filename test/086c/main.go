package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	travelList := make([][3]int, n+1)

	for i := range travelList {
		if i == 0 {
			travelList[i] = [3]int{0, 0, 0}
			continue
		}
		fmt.Scan(&travelList[i][0], &travelList[i][1], &travelList[i][2])
	}

	available := "Yes"

	for i := range travelList {
		if i == 0 {
			continue
		}

		dist := abs(travelList[i][1]-travelList[i-1][1]) + abs(travelList[i][2]-travelList[i-1][2])
		dt := travelList[i][0] - travelList[i-1][0]

		if dist > dt || (dt-dist)%2 != 0 {
			available = "No"
			break
		}
	}
	fmt.Println(available)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
