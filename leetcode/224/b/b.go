package main

// github.com/EndlessCheng/codeforces-go
func tupleSameProduct(a []int) (ans int) {
	cnt := map[int]int{}
	for i, v := range a {
		for _, w := range a[i+1:] {
			cnt[v*w]++
		}
	}
	for _, c := range cnt {
		ans += 4 * c * (c - 1)
	}
	return
}
