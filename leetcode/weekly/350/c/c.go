package main

// https://space.bilibili.com/206214
func specialPerm(nums []int) (ans int) {
	const mod int = 1e9 + 7
	n := len(nums)
	m := 1 << n
	f := make([][]int, m)
	f[0] = make([]int, n)
	for j := range f[0] {
		f[0][j] = 1
	}
	for i := 1; i < m; i++ {
		f[i] = make([]int, n)
		for j, x := range nums {
			for k, y := range nums {
				if i>>k&1 > 0 && (x%y == 0 || y%x == 0) {
					f[i][j] += f[i^(1<<k)][k]
				}
			}
		}
	}
	for j := range nums {
		ans += f[(m-1)^(1<<j)][j]
	}
	return ans % mod
}
