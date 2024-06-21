package main

// https://space.bilibili.com/206214
func goodSubsetofBinaryMatrix(grid [][]int) []int {
	n := len(grid[0])
	f := make([]int, 1<<n)
	for i := range f {
		f[i] = -1
	}
	for i, row := range grid {
		mask := 0
		for j, x := range row {
			mask |= x << j
		}
		if mask == 0 {
			return []int{i}
		}
		f[mask] = i
	}

	u := 1<<n - 1
	for s := 1; s < u; s++ {
		for b := 0; b < n; b++ {
			if s>>b&1 == 0 {
				continue
			}
			f[s] = max(f[s], f[s^1<<b])
			i := f[s]
			if i < 0 {
				continue
			}
			j := f[u^s]
			if j >= 0 {
				return []int{min(i, j), max(i, j)}
			}
		}
	}
	return nil
}
