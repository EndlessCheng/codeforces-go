package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1098B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	ans := make([][]byte, n)
	t := make([][]byte, n)
	a := make([][]byte, n)
	for i := range a {
		Fscan(in, &a[i])
		ans[i] = make([]byte, m)
		t[i] = make([]byte, m)
	}
	mi := int(1e9)
	for _, p := range []string{"AGCT", "ACGT", "ATGC", "GCAT", "GTAC", "CTAG"} {
		d := 0
		for i, r := range a {
			x, y := 0, 0
			for j, b := range r {
				if b != p[i&1<<1|j&1] {
					x++
				}
				if b != p[i&1<<1|j&1^1] {
					y++
				}
			}
			if x < y {
				d += x
				for j := range r {
					t[i][j] = p[i&1<<1|j&1]
				}
			} else {
				d += y
				for j := range r {
					t[i][j] = p[i&1<<1|j&1^1]
				}
			}
		}
		if d < mi {
			mi = d
			for i, r := range t {
				copy(ans[i], r)
			}
		}
		d = 0
		for j := range a[0] {
			x, y := 0, 0
			for i, r := range a {
				if r[j] != p[j&1<<1|i&1] {
					x++
				}
				if r[j] != p[j&1<<1|i&1^1] {
					y++
				}
			}
			if x < y {
				d += x
				for i := range t {
					t[i][j] = p[j&1<<1|i&1]
				}
			} else {
				d += y
				for i := range t {
					t[i][j] = p[j&1<<1|i&1^1]
				}
			}
		}
		if d < mi {
			mi = d
			for i, r := range t {
				copy(ans[i], r)
			}
		}
	}
	for _, r := range ans {
		Fprintf(out, "%s\n", r)
	}
}

//func main() { CF1098B(os.Stdin, os.Stdout) }
