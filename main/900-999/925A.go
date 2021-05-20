package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF925A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, m, cl, ce, v, q, x1, y1, x2, y2 int
	Fscan(in, &n, &m, &cl, &ce, &v)
	ls := make([]int, cl)
	for i := range ls {
		Fscan(in, &ls[i])
	}
	es := make([]int, ce)
	for i := range es {
		Fscan(in, &es[i])
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &x1, &y1, &x2, &y2)
		if x1 == x2 {
			Fprintln(out, abs(y1-y2))
			continue
		}
		ans := int(1e9)
		dx := abs(x1 - x2)
		i := sort.SearchInts(ls, y1)
		if i < cl {
			ans = min(ans, ls[i]-y1+dx+abs(y2-ls[i]))
		}
		if i > 0 {
			ans = min(ans, y1-ls[i-1]+dx+abs(y2-ls[i-1]))
		}
		t := (dx-1)/v + 1
		i = sort.SearchInts(es, y1)
		if i < ce {
			ans = min(ans, es[i]-y1+t+abs(y2-es[i]))
		}
		if i > 0 {
			ans = min(ans, y1-es[i-1]+t+abs(y2-es[i-1]))
		}
		Fprintln(out, ans)
	}
}

//func main() { CF925A(os.Stdin, os.Stdout) }
