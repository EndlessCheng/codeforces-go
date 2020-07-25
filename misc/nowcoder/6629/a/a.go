package main

// github.com/EndlessCheng/codeforces-go
func solve(_ int, a []int) (c int) {
	vis := map[int]bool{}
	for _, v := range a {
		for ; v&1 == 0 && !vis[v]; v >>= 1 {
			vis[v] = true
			c++
		}
	}
	return
}
