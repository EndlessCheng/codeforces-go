package main

// github.com/EndlessCheng/codeforces-go
func maxDistance(colors []int) int {
	n := len(colors)
	c := colors[0]
	if c != colors[n-1] {
		return n - 1
	}

	// 找最右边的颜色不等于 c 的房子
	// 题目保证至少有两栋颜色不同的房子
	r := n - 2
	for colors[r] == c {
		r--
	}

	// 找最左边的颜色不等于 c 的房子
	l := 1
	for colors[l] == c {
		l++
	}

	return max(r, n-1-l)
}
