package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

// code for problem https://codeforces.com/contest/1363/problem/D

type (
	input struct {
		n, k int
		a    [][]int
	}
	guess struct{ ans []int }
	qIn   struct{ q []int }
	qOut  struct{ max int }
)

// github.com/EndlessCheng/codeforces-go
func run(in input, Q func(qIn) qOut) (gs guess) {
	n, k, a := in.n, in.k, in.a
	gs.ans = make([]int, k)
	q := []int{}
	for i := 1; i <= n; i++ {
		q = append(q, i)
	}
	allMax := Q(qIn{q}).max
	l, r := 0, k
	for r-l > 1 {
		m := (l + r) >> 1
		q = []int{}
		for _, ids := range a[l:m] {
			q = append(q, ids...)
		}
		if Q(qIn{q}).max == allMax {
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
	gs.ans[l] = Q(qIn{q}).max
	return
}

func ioq() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	Q := func(qi qIn) (resp qOut) {
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
		d := input{}
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
		gs := run(d, Q)
		Fprint(out, "!")
		for _, v := range gs.ans {
			Fprint(out, " ", v)
		}
		Fprintln(out)
		out.Flush()
		var s string
		Fscan(in, &s)
	}
}

func main() { ioq() }
