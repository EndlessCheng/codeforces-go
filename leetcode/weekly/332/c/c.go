package main

import "strings"

// https://space.bilibili.com/206214
func substringXorQueries(s string, queries [][]int) [][]int {
	type pair struct{ l, r int }
	m := map[int]pair{}
	if i := strings.IndexByte(s, '0'); i >= 0 {
		m[0] = pair{i, i}
	}
	for l, c := range s {
		if c == '0' {
			continue
		}
		for r, x := l, 0; r < l+30 && r < len(s); r++ {
			x = x<<1 | int(s[r]&1)
			if _, ok := m[x]; !ok {
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
