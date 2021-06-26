package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF301D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, l, r int
	Fscan(in, &n, &q)
	pos := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &r)
		pos[r] = i
	}
	targetPos := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j += i {
			l, r := pos[i], pos[j]
			if l > r {
				l, r = r, l
			}
			targetPos[l] = append(targetPos[l], r)
		}
	}
	type query struct{ r, i int }
	qs := make([][]query, n+1)
	for i := 0; i < q; i++ {
		Fscan(in, &l, &r)
		qs[l] = append(qs[l], query{r, i})
	}
	ans := make([]int, q)
	tree := make([]int, n+1)
	for i := n; i > 0; i-- {
		for _, r := range targetPos[i] {
			for ; r <= n; r += r & -r {
				tree[r]++
			}
		}
		for _, q := range qs[i] {
			for r := q.r; r > 0; r &= r - 1 {
				ans[q.i] += tree[r]
			}
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { CF301D(os.Stdin, os.Stdout) }
