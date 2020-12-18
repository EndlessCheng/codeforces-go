package main

// github.com/EndlessCheng/codeforces-go
func wwork(n int, a []int) (ans int) {
	tree := make([]int, n+1)
	add := func(i int) {
		for ; i <= n; i += i & -i {
			tree[i]++
		}
	}
	cnt := func(i int) (c int) {
		for ; i > 0; i &= i - 1 {
			c += tree[i]
		}
		return
	}
	pos := make([]int, n+1)
	for i, v := range a {
		pos[v] = i + 1
	}
	for i := n; i > 0; i-- {
		if p := pos[i]; p-cnt(p) != i {
			ans++
			add(p)
		}
	}
	return
}
