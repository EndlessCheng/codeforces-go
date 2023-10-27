package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF777C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, m, q, l, r int
	Fscan(in, &n, &m)
	minL := make([]int, n)
	up := make([]int, m)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		minL[i] = i
		for j := range a[i] {
			Fscan(in, &a[i][j])
			if i > 0 && a[i][j] < a[i-1][j] {
				up[j] = i
			}
			minL[i] = min(minL[i], up[j])
		}
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &l, &r)
		if minL[r-1] <= l-1 {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

//func main() { CF777C(os.Stdin, os.Stdout) }
