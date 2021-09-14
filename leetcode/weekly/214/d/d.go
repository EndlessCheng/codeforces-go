package main

// github.com/EndlessCheng/codeforces-go
type fenwick struct {
	tree []int
}

func newFenwickTree(n int) fenwick {
	return fenwick{make([]int, n+1)}
}
func (f fenwick) add(i int) {
	for ; i < len(f.tree); i += i & -i {
		f.tree[i]++
	}
}
func (f fenwick) sum(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f.tree[i]
	}
	return
}
func (f fenwick) query(l, r int) (res int) {
	return f.sum(r) - f.sum(l-1)
}

func createSortedArray(a []int) (ans int) {
	const mod int = 1e9 + 7
	f := newFenwickTree(1e5)
	for _, v := range a {
		ans += min(f.sum(v-1), f.query(v+1, 1e5))
		f.add(v)
	}
	ans %= mod
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
