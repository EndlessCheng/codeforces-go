package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF711B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, x, y int
	Fscan(in, &n)
	if n == 1 {
		Fprint(out, 1)
		return
	}

	a := make([][]int64, n)
	for i := range a {
		a[i] = make([]int64, n)
		for j := range a[i] {
			Fscan(in, &a[i][j])
			if a[i][j] == 0 {
				x, y = i, j
			}
		}
	}

	sum := int64(0)
	ri := 0
	if x == 0 {
		ri = 1
	}
	for j, v := range a[ri] {
		sum += v
		if j != y {
			a[x][y] -= a[x][j]
		}
	}
	a[x][y] += sum
	if a[x][y] < 1 {
		Fprint(out, -1)
		return
	}
	for _, r := range a {
		s := int64(0)
		for _, v := range r {
			s += v
		}
		if s != sum {
			Fprint(out, -1)
			return
		}
	}
	for j := range a[0] {
		s := int64(0)
		for _, r := range a {
			s += r[j]
		}
		if s != sum {
			Fprint(out, -1)
			return
		}
	}
	var s, s2 int64
	for i, r := range a {
		s += r[i]
		s2 += r[n-1-i]
	}
	if s != sum || s2 != sum {
		Fprint(out, -1)
	} else {
		Fprint(out, a[x][y])
	}
}

//func main() { CF711B(os.Stdin, os.Stdout) }
