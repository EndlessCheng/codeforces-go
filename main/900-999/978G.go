package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF978G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type exam struct{ i, l, r, c int }

	var n, m, l, r, c int
	Fscan(in, &n, &m)
	ans := make([]int, n)
	es := make([]exam, m)
	for i := range es {
		Fscan(in, &l, &r, &c)
		es[i] = exam{i + 1, l - 1, r - 1, c}
		ans[r-1] = m + 1
	}
	sort.Slice(es, func(i, j int) bool { return es[i].r < es[j].r })
	for _, e := range es {
		for i := e.l; e.c > 0 && i < e.r; i++ {
			if ans[i] == 0 {
				ans[i] = e.i
				e.c--
			}
		}
		if e.c > 0 {
			Fprint(out, -1)
			return
		}
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF978G(os.Stdin, os.Stdout) }
