package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF461C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, tp, l, r int
	Fscan(in, &n, &q)
	sum := make([]int, n+1)
	for i := range sum {
		sum[i] = i
	}
	for st, ed, rev := 0, n, false; q > 0; q-- {
		if Fscan(in, &tp, &l); tp == 1 {
			if l > ed-l {
				l, rev = ed-l, !rev
			}
			if rev {
				for i := ed - l; i < ed; i++ {
					sum[st+(ed-l)*2-i] += n - sum[st+i]
				}
			} else {
				st += l
				for i := 0; i < l; i++ {
					sum[st+i] -= sum[st-i]
				}
			}
			ed -= l
		} else {
			Fscan(in, &r)
			if rev {
				l, r = ed-r, ed-l
			}
			Fprintln(out, sum[st+r]-sum[st+l])
		}
	}
}

//func main() { CF461C(os.Stdin, os.Stdout) }
