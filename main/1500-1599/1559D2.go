package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type uf59 []int

func (u uf59) find(x int) int {
	if u[x] != x {
		u[x] = u.find(u[x])
	}
	return u[x]
}

func (u uf59) merge(from, to int) bool {
	x, y := u.find(from), u.find(to)
	if x == y {
		return false
	}
	u[x] = y
	return true
}

func (u uf59) same(x, y int) bool {
	return u.find(x) == u.find(y)
}

func cf1559D2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m1, m2 int
	Fscan(in, &n, &m1, &m2)
	init := func(m int) uf59 {
		u := make(uf59, n+1)
		for i := range u {
			u[i] = i
		}
		for range m {
			var v, w int
			Fscan(in, &v, &w)
			u.merge(v, w)
		}
		return u
	}
	u1 := init(m1)
	u2 := init(m2)

	ans := [][2]int{}
	for i := 1; i <= n; i++ {
		if !u1.same(i, 1) && !u2.same(i, 1) {
			ans = append(ans, [2]int{i, 1})
			u1.merge(i, 1)
			u2.merge(i, 1)
		}
	}

	var a, b []int
	for i := 1; i <= n; i++ {
		if u1.merge(i, 1) {
			a = append(a, i)
		}
		if u2.merge(i, 1) {
			b = append(b, i)
		}
	}
	for i := range min(len(a), len(b)) {
		ans = append(ans, [2]int{a[i], b[i]})
	}

	Fprintln(out, len(ans))
	for _, p := range ans {
		Fprintln(out, p[0], p[1])
	}
}

//func main() { cf1559D2(bufio.NewReader(os.Stdin), os.Stdout) }
