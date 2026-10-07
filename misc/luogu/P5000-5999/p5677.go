package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
type fenwick []int

func (t fenwick) update(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] += val
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func (t fenwick) query(l, r int) int {
	return t.pre(r) - t.pre(l-1)
}

func p5677(in io.Reader, out io.Writer) {
	var n, m, l, r, ans int
	Fscan(in, &n, &m)
	a := make([]int, n)
	pos := map[int]int{}
	for i := range a {
		Fscan(in, &a[i])
		pos[a[i]] = i
	}
	sorted := slices.Clone(a)
	slices.Sort(sorted)

	type pair struct{ l, i int }
	ls := make([][]pair, n+1)
	for i := 1; i <= m; i++ {
		Fscan(in, &l, &r)
		ls[r] = append(ls[r], pair{l, i})
	}

	t := make(fenwick, n+1)
	todo := make([][]int, n+1)
	for i, v := range a {
		j := sort.SearchInts(sorted, v)

		d1 := int(1e9)
		p := n
		if j > 0 {
			d1 = v - sorted[j-1]
			p = pos[sorted[j-1]]
		}

		d2 := int(1e9)
		q := n
		if j < n-1 {
			d2 = sorted[j+1] - v
			q = pos[sorted[j+1]]
		}

		if d1 <= d2 {
			if p < i {
				t.update(p+1, 1)
			} else {
				todo[p] = append(todo[p], i)
			}
		}
		if d2 <= d1 {
			if q < i {
				t.update(q+1, 1)
			} else {
				todo[q] = append(todo[q], i)
			}
		}

		for _, j := range todo[i] {
			t.update(j+1, 1)
		}

		for _, p := range ls[i+1] {
			ans += t.query(p.l, i+1) * p.i
		}
	}

	Fprint(out, ans)
}

//func main() { p5677(bufio.NewReader(os.Stdin), os.Stdout) }
