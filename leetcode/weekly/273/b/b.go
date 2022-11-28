package main

// github.com/EndlessCheng/codeforces-go
var dirs = [][2]int{'L': {0, -1}, 'R': {0, 1}, 'U': {-1, 0}, 'D': {1, 0}}

// TODO
func executeInstructions(n int, startPos []int, s string) []int {
	ns := len(s)
	ans := make([]int, ns)
	for i := range ans {
		ans[i] = ns - i
	}
	x0, y0 := startPos[0], startPos[1]
	cnt := make(map[int][]int, n*2)
	add := func(x, v int) { cnt[x] = append(cnt[x], v) }
	add(x0<<1, 0)
	add(y0<<1|1, 0)
	x, y := 0, 0
	for i, ch := range s {
		x += dirs[ch][0]
		y += dirs[ch][1]
		add((x0-x)<<1, i+1)
		add((y0-y)<<1|1, i+1)
		for _, key := range []int{(n - x) << 1, (n-y)<<1 | 1, (-x - 1) << 1, (-y-1)<<1 | 1} {
			for _, j := range cnt[key] {
				if ans[j] > i-j {
					ans[j] = i - j
				}
			}
			cnt[key] = nil
		}
	}
	return ans
}
