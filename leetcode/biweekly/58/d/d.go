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
		for ; i >= hl && i+hl < n && s[i-hl] == s[i+hl]; hl++ {}
		if i+hl > r {
			mid, r = i, i+hl
		}
		halfLen[i] = hl
	}

	leftTo := make([]int, n)
	rightTo := make([]int, n)
	for i, hl := range halfLen {
		pl, pr := i-hl+1, i+hl-1
		leftTo[pr] = max(leftTo[pr], hl*2-1)
		rightTo[pl] = max(rightTo[pl], hl*2-1)
	}
	for i := n - 2; i >= 0; i-- {
		leftTo[i] = max(leftTo[i], leftTo[i+1]-2)
	}
	for i := 1; i < n; i++ {
		leftTo[i] = max(leftTo[i], leftTo[i-1])
	}
	for i := 1; i < n; i++ {
		rightTo[i] = max(rightTo[i], rightTo[i-1]-2)
	}
	for i := n - 2; i >= 0; i-- {
		rightTo[i] = max(rightTo[i], rightTo[i+1])
	}

	ans := 0
	for i := 1; i < n; i++ {
		ans = max(ans, leftTo[i-1]*rightTo[i])
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
