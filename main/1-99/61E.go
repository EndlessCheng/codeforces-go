package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF61E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	ps := make([]struct{ v, i int }, n)
	for i := range ps {
		Fscan(in, &ps[i].v)
		ps[i].i = i
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].v < ps[j].v })
	a := make([]int, n)
	for i, p := range ps {
		a[p.i] = i + 1
	}

	tree := make([]int, n+1)
	add := func(i int) {
		for ; i <= n; i += i & -i {
			tree[i]++
		}
	}
	sum := func(i int) (res int) {
		for ; i > 0; i &= i - 1 {
			res += tree[i]
		}
		return
	}
	query := func(l, r int) int { return sum(r) - sum(l-1) }
	tree2 := make([]int64, n+1)
	add2 := func(i, val int) {
		for ; i <= n; i += i & -i {
			tree2[i] += int64(val)
		}
	}
	sum2 := func(i int) (res int64) {
		for ; i > 0; i &= i - 1 {
			res += tree2[i]
		}
		return
	}
	query2 := func(l, r int) int64 { return sum2(r) - sum2(l-1) }

	ans := int64(0)
	for _, v := range a {
		ans += query2(v+1, n)
		add2(v, query(v+1, n))
		add(v)
	}
	Fprint(out, ans)
}

//func main() { CF61E(os.Stdin, os.Stdout) }
