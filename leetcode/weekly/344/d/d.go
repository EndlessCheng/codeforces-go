package main

// https://space.bilibili.com/206214
func minIncrements(n int, cost []int) (ans int) {
	for i := n / 2; i > 0; i-- {
		left, right := cost[i*2-1], cost[i*2]
		if left > right { // 保证 left <= right
			left, right = right, left
		}
		ans += right - left // 大的减小的
		cost[i-1] += right // 把路径和加给父节点
	}
	return
}

func minIncrements2(n int, cost []int) (ans int) {
	mx := make([]int, n*2+2)
	for i := n; i > 0; i-- {
		mx[i] = cost[i-1] + max(mx[i*2], mx[i*2+1])
	}

	var f func(int, int)
	f = func(i, s int) {
		if i > n {
			return
		}
		d := mx[1] - mx[i] - s
		ans += d

		s += cost[i-1] + d
		f(i*2, s)
		f(i*2+1, s)
	}
	f(1, 0)

	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
