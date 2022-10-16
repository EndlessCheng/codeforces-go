package main

// https://space.bilibili.com/206214
func componentValue(nums []int, edges [][]int) int {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var target int
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		sum := nums[x] // 价值
		for _, y := range g[x] {
			if y != fa {
				res := dfs(y, x)
				if res < 0 {
					return -1
				}
				sum += res
			}
		}
		if sum > target {
			return -1
		}
		if sum == target {
			return 0
		}
		return sum
	}

	total, mx := 0, 0
	for _, x := range nums {
		total += x
		mx = max(mx, x)
	}
	for i := min(n, total/mx); ; i-- {
		if total%i == 0 {
			target = total / i
			if dfs(0, -1) == 0 {
				return i - 1
			}
		}
	}
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
