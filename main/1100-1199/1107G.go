package main

import (
	. "fmt"
	"io"
	"math/bits"
	"slices"
)

// https://github.com/EndlessCheng
type st07 [][19][2]int

func newST07(a []int) st07 {
	n := len(a)
	st := make(st07, n)
	for i, v := range a {
		st[i][0] = [2]int{v, v}
	}
	for j := 1; j < 19; j++ {
		for i := range n - 1<<j + 1 {
			st[i][j][0] = min(st[i][j-1][0], st[i+1<<(j-1)][j-1][0])
			st[i][j][1] = max(st[i][j-1][1], st[i+1<<(j-1)][j-1][1])
		}
	}
	return st
}

func (st st07) query(l, r int) [2]int {
	k := bits.Len(uint(r-l)) - 1
	p, q := st[l][k], st[r-1<<k][k]
	return [2]int{min(p[0], q[0]), max(p[1], q[1])}
}

func cf1107G(in io.Reader, out io.Writer) {
	var n, earn, c, ans int
	Fscan(in, &n, &earn)
	d := make([]int, n)
	s := make([]int, n+1)
	for i := range d {
		Fscan(in, &d[i], &c)
		s[i+1] = s[i] + earn - c
		ans = max(ans, earn-c)
	}
	t := newST07(s)

	type pair struct{ d, i int }
	ds := make([]pair, n-1)
	left := make([]int, n-1)
	st := []int{-1}
	for i := range n - 1 {
		v := d[i+1] - d[i]
		ds[i] = pair{v, i}
		for len(st) > 1 && d[st[len(st)-1]+1]-d[st[len(st)-1]] <= v {
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}
	right := make([]int, n-1)
	st = []int{n - 1}
	for i := n - 2; i >= 0; i-- {
		for len(st) > 1 && d[st[len(st)-1]+1]-d[st[len(st)-1]] <= d[i+1]-d[i] {
			st = st[:len(st)-1]
		}
		right[i] = st[len(st)-1]
		st = append(st, i)
	}

	slices.SortFunc(ds, func(a, b pair) int { return a.d - b.d })
	for _, p := range ds {
		i := p.i
		ans = max(ans, t.query(i+2, right[i]+2)[1]-t.query(left[i]+1, i+1)[0]-p.d*p.d)
	}
	Fprint(out, ans)
}

//func main() { cf1107G(bufio.NewReader(os.Stdin), os.Stdout) }
