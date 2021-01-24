package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func above(x int) (res int) {
	s := 0
	for i := 1; s+i <= x; i++ {
		s += i
		res += i - 1
	}
	if s < x {
		res += x - s - 1
	}
	return
}

func minimumBoxes(n int) int {
	return sort.Search(2e6, func(x int) bool {
		sum := x
		for {
			y := above(x)
			if y == 0 {
				break
			}
			sum += y
			x = y
		}
		return sum >= n
	})
}
