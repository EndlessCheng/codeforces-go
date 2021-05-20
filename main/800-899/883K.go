package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF883K(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, g int
	Fscan(in, &n)
	a := make([]struct{ s, mi, mx int }, n)
	for i := range a {
		Fscan(in, &a[i].s, &g)
		a[i].mi = a[i].s
		a[i].mx = a[i].s + g
	}

	for i := 1; i < n; i++ {
		a[i].mx = min(a[i].mx, a[i-1].mx+1)
		a[i].mi = max(a[i].mi, a[i-1].mi-1)
	}
	for i := n - 2; i >= 0; i-- {
		a[i].mx = min(a[i].mx, a[i+1].mx+1)
		a[i].mi = max(a[i].mi, a[i+1].mi-1)
	}

	ans := int64(0)
	for _, p := range a {
		if p.mx < p.mi {
			Fprint(out, -1)
			return
		}
		ans += int64(p.mx - p.s)
	}
	Fprintln(out, ans)
	for _, p := range a {
		Fprint(out, p.mx, " ")
	}
}

//func main() { CF883K(os.Stdin, os.Stdout) }
