package main

// github.com/EndlessCheng/codeforces-go
func maxProduct(s string) int64 {
	n := len(s)
	halfLen := make([]int, n)
	for i, mid, r := 0, 0, 0; i < n; i++ {
		hl := 1
		if i < r {
			hl = min(halfLen[mid*2-i], r-i)
		}
		for ; i >= hl && i+hl < n && s[i-hl] == s[i+hl]; hl++ {
		}
		if i+hl > r {
			mid, r = i, i+hl
		}
		halfLen[i] = hl
	}

	startPL := make([]int, n)
	endPL := make([]int, n)
	for i, hl := range halfLen {
		left, right := i-hl+1, i+hl-1
		startPL[left] = max(startPL[left], hl*2-1)
		endPL[right] = max(endPL[right], hl*2-1)
	}
	for i := 1; i < n; i++ {
		startPL[i] = max(startPL[i], startPL[i-1]-2)
	}
	for i := n - 2; i >= 0; i-- {
		startPL[i] = max(startPL[i], startPL[i+1])
	}
	for i := n - 2; i >= 0; i-- {
		endPL[i] = max(endPL[i], endPL[i+1]-2)
	}
	for i := 1; i < n; i++ {
		endPL[i] = max(endPL[i], endPL[i-1])
	}

	ans := 0
	for i := 1; i < n; i++ {
		ans = max(ans, endPL[i-1]*startPL[i])
	}
	return int64(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
