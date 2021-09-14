package main

// github.com/EndlessCheng/codeforces-go
func minOperations(boxes string) (ans []int) {
	ans = make([]int, len(boxes))
	for i := range boxes {
		for j, b := range boxes {
			if b == '1' {
				ans[i] += abs(i - j)
			}
		}
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
