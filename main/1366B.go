package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1366B(_r io.Reader, out io.Writer) {
	type pair struct{ l, r int }
	in := bufio.NewReader(_r)
	var t, n, x, m, l, r int
	for Fscan(in, &t); t > 0; t-- {
		var a, es []pair
		for Fscan(in, &n, &x, &m); m > 0; m-- {
			Fscan(in, &l, &r)
			if l <= x && x <= r {
				a = append(a, pair{l, r})
				es = append(es, pair{l, 1}, pair{r, -1})
				continue
			}
			for _, p := range a {
				if p.l <= r && p.r >= l {
					a = append(a, pair{l, r})
					es = append(es, pair{l, 1}, pair{r, -1})
					break
				}
			}
		}
		if len(es) == 0 {
			Fprintln(out, 1)
			continue
		}
		sort.Slice(es, func(i, j int) bool { return es[i].l < es[j].l })
		ans := es[len(es)-1].l - es[0].l + 1
		c, st := 0, es[0].l
		for _, e := range es {
			if c == 0 {
				if d := e.l - st - 1; d > 0 {
					ans -= d
				}
			}
			c += e.r
			if c == 0 {
				st = e.l
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1366B(os.Stdin, os.Stdout) }
