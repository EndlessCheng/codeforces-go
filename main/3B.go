package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF3B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, sz, i, s int
	Fscan(in, &n, &sz)
	type pair struct{ v, w, i int }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].w, &a[i].v)
		a[i].i = i
	}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.v*b.w > b.v*a.w })
	for ; i < n && a[i].w <= sz; i++ {
		sz -= a[i].w
		s += a[i].v
	}
	if i < n && sz > 0 {
		j := i - 1
		for ; j >= 0 && a[j].w == 2; j-- {
		}
		for k := i + 1; k < n; k++ {
			if a[k].w == 1 {
				if j >= 0 && a[i].v-a[j].v > a[k].v {
					break
				}
				s += a[k].v
				a[i].i = a[k].i
				i++
				goto print
			}
		}
		if j >= 0 && a[i].v > a[j].v {
			s += a[i].v - a[j].v
			a[j].i = a[i].i
		}
	}
print:
	Fprintln(out, s)
	for _, p := range a[:i] {
		Fprint(out, p.i+1, " ")
	}
}

//func main() { CF3B(os.Stdin, os.Stdout) }
