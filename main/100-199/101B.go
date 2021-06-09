package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF101B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int = 1e9 + 7

	var n, m, ans int
	Fscan(in, &n, &m)
	a := make([]struct{ l, r int }, m)
	b := make(sort.IntSlice, 0, m*2)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r)
		b = append(b, a[i].l, a[i].r)
	}
	b.Sort()
	if m == 0 || b[0] != 0 || b[m*2-1] != n {
		Fprint(out, 0)
		return
	}

	tree := make([]int, m*2+1)
	add := func(i, val int) {
		for ; i <= m*2; i += i & -i {
			tree[i] = (tree[i] + val) % mod
		}
	}
	sum := func(i int) (res int) {
		for ; i > 0; i &= i - 1 {
			res = (res + tree[i]) % mod
		}
		return
	}
	add(1, 1)
	sort.Slice(a, func(i, j int) bool { return a[i].r < a[j].r })
	for _, p := range a {
		r := b.Search(p.r)
		res := (sum(r) + mod - sum(b.Search(p.l))) % mod
		if p.r < n {
			add(r+1, res)
		} else {
			ans = (ans + res) % mod
		}
	}
	Fprint(out, ans)
}

//func main() { CF101B(os.Stdin, os.Stdout) }
