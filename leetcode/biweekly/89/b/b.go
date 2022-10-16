package main

// https://space.bilibili.com/206214
const mod int = 1e9 + 7

func productQueries(n int, queries [][]int) []int {
	a := []int{}
	for n > 0 {
		lb := n & -n
		a = append(a, lb)
		n ^= lb
	}
	na := len(a)
	res := make([][]int, na)
	for i, x := range a {
		res[i] = make([]int, na)
		res[i][i] = x
		for j := i + 1; j < na; j++ {
			res[i][j] = res[i][j-1] * a[j] % mod
		}
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = res[q[0]][q[1]]
	}
	return ans
}
