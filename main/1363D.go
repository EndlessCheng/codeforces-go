package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

type (
	input1363 struct {
		n, k int
		a    [][]int
	}
	guess1363 struct{ ans []int }
	qIn1363   struct{ q []int }
	qOut1363  struct{ max int }
)

// github.com/EndlessCheng/codeforces-go
func CF1363D(in input1363, Q func(qIn1363) qOut1363) (gs guess1363) {
	n, k, a := in.n, in.k, in.a
	gs.ans = make([]int, k)
	q := []int{}
	for i := 1; i <= n; i++ {
		q = append(q, i)
	}
	allMax := Q(qIn1363{q}).max
	l, r := 0, k
	for r-l > 1 {
		m := (l + r) >> 1
		q = []int{}
		for _, ids := range a[l:m] {
			q = append(q, ids...)
		}
		if Q(qIn1363{q}).max == allMax {
			for i := m; i < r; i++ {
				gs.ans[i] = allMax
			}
			r = m
		} else {
			for i := l; i < m; i++ {
				gs.ans[i] = allMax
			}
			l = m
		}
	}
	sort.Ints(a[l])
	q = []int{}
	for i, j := 1, 0; i <= n; i++ {
		if j == len(a[l]) || i != a[l][j] {
			q = append(q, i)
		} else {
			j++
		}
	}
	gs.ans[l] = Q(qIn1363{q}).max
	return
}

func ioq1363D() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	Q := func(qi qIn1363) (resp qOut1363) {
		Fprint(out, "? ", len(qi.q))
		for _, v := range qi.q {
			Fprint(out, " ", v)
		}
		Fprintln(out)
		out.Flush()
		Fscan(in, &resp.max)
		return
	}
	var t, c int
	for Fscan(in, &t); t > 0; t-- {
		d := input1363{}
		Fscan(in, &d.n, &d.k)
		a := make([][]int, d.k)
		for i := range a {
			Fscan(in, &c)
			a[i] = make([]int, c)
			for j := range a[i] {
				Fscan(in, &a[i][j])
			}
		}
		d.a = a
		gs := CF1363D(d, Q)
		Fprint(out, "!")
		for _, v := range gs.ans {
			Fprint(out, " ", v)
		}
		Fprintln(out)
		out.Flush()
		var s []byte
		Fscan(in, &s)
	}
}

//func main() { ioq1363D() }
