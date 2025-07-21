package main

// https://space.bilibili.com/206214
func countTrapezoids(points [][]int) (ans int) {
	const mod = 1_000_000_007
	cnt := map[int]int{}
	for _, p := range points {
		cnt[p[1]]++ // 统计每一行（水平线）有多少个点
	}

	s := 0
	for _, c := range cnt {
		k := c * (c - 1) / 2
		ans += s * k
		s += k
	}
	return ans % mod
}
