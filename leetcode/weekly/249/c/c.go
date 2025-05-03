package main

// github.com/EndlessCheng/codeforces-go
func colorTheGrid(m, n int) int {
	const mod = 1_000_000_007
	pow3 := make([]int, m)
	pow3[0] = 1
	for i := 1; i < m; i++ {
		pow3[i] = pow3[i-1] * 3
	}

	valid := []int{}
next:
	for color := range pow3[m-1] * 3 {
		for i := range m - 1 {
			if color/pow3[i+1]%3 == color/pow3[i]%3 { // 相邻颜色相同
				continue next
			}
		}
		valid = append(valid, color)
	}

	nv := len(valid)
	nxt := make([][]int, nv)
	for i, color1 := range valid {
	next2:
		for j, color2 := range valid {
			for _, p3 := range pow3 {
				if color1/p3%3 == color2/p3%3 { // 相邻颜色相同
					continue next2
				}
			}
			nxt[i] = append(nxt[i], j)
		}
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, nv)
	}
	for j := range f[0] {
		f[0][j] = 1
	}
	for i := 1; i < n; i++ {
		for j := range f[i] {
			for _, k := range nxt[j] {
				f[i][j] += f[i-1][k]
			}
			f[i][j] %= mod
		}
	}

	ans := 0
	for _, fv := range f[n-1] {
		ans += fv
	}
	return ans % mod
}
