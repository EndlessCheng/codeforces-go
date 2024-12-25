package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func maxConsecutive(bottom, top int, special []int) int {
	slices.Sort(special)
	n := len(special)
	ans := max(special[0]-bottom, top-special[n-1])
	for i := 1; i < n; i++ {
		ans = max(ans, special[i]-special[i-1]-1)
	}
	return ans
}

func maxConsecutive2(bottom, top int, special []int) (ans int) {
	special = append(special, bottom-1, top+1)
	slices.Sort(special)
	for i := 1; i < len(special); i++ {
		ans = max(ans, special[i]-special[i-1]-1)
	}
	return
}
