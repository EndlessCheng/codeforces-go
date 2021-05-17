package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF141C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct {
		s string
		v int
	}

	var n int
	Fscan(in, &n)
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].s, &a[i].v)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].v < a[j].v })
	ans := []pair{}
	for i := 0; i < n; {
		c := a[i].v
		if len(ans) < c {
			Fprint(out, -1)
			return
		}
		b := append([]pair(nil), ans[c:]...)
		ans = ans[:c]
		for ; i < n && a[i].v == c; i++ {
			ans = append(ans, pair{a[i].s, n - c})
		}
		ans = append(ans, b...)
	}
	for _, p := range ans {
		Fprintln(out, p.s, p.v)
	}
}

//func main() { CF141C(os.Stdin, os.Stdout) }
