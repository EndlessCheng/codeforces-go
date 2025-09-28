package main

// https://space.bilibili.com/206214
type pair struct{ x, y int }
var dirs = []pair{'L': {-1, 0}, 'R': {1, 0}, 'D': {0, -1}, 'U': {0, 1}}

func distinctPoints(s string, k int) int {
	p := pair{}
	set := map[pair]struct{}{p: {}} // 第一个窗口
	for i := k; i < len(s); i++ {
		in, out := s[i], s[i-k]
		p.x += dirs[in].x - dirs[out].x
		p.y += dirs[in].y - dirs[out].y
		set[p] = struct{}{}
	}
	return len(set)
}
