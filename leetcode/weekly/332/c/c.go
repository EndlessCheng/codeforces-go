package main

// https://space.bilibili.com/206214
func substringXorQueries(s string, queries [][]int) [][]int {
	type pair struct{ l, r int }
	m := map[int]pair{}
	for l := range s {
		for r, x := l, 0; r < l+30 && r < len(s); r++ {
			x = x<<1 | int(s[r]&1)
			if p, ok := m[x]; !ok || r-l < p.r-p.l {
				m[x] = pair{l, r}
			}
		}
	}

	ans := make([][]int, len(queries))
	notFound := []int{-1, -1} // 避免重复创建
	for i, q := range queries {
		p, ok := m[q[0]^q[1]]
		if !ok {
			ans[i] = notFound
		} else {
			ans[i] = []int{p.l, p.r}
		}
	}
	return ans
}
