package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1635F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, l, r int
	Fscan(in, &n, &m)
	x := make([]int, n)
	w := make([]int, n)
	tree := make([]int64, n+1)
	for i := range x {
		Fscan(in, &x[i], &w[i])
		tree[i+1] = 9e18
	}
	type query struct{ r, i int }
	qs := make([][]query, n)
	ans := make([]int64, m)
	for i := range ans {
		Fscan(in, &l, &r)
		qs[l-1] = append(qs[l-1], query{r, i})
	}

	s := []int{}
	for i := n - 1; i >= 0; i-- {
		for len(s) > 0 {
			r := s[len(s)-1]
			v := int64(x[r]-x[i]) * int64(w[i]+w[r])
			for r := r + 1; r <= n; r += r & -r {
				if v < tree[r] {
					tree[r] = v
				}
			}
			if w[r] < w[i] {
				break
			}
			s = s[:len(s)-1]
		}
		s = append(s, i)
		for _, q := range qs[i] {
			minV := int64(9e18)
			for r := q.r; r > 0; r &= r - 1 {
				if tree[r] < minV {
					minV = tree[r]
				}
			}
			ans[q.i] = minV
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { CF1635F(os.Stdin, os.Stdout) }
