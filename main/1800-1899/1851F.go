package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1851F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		type vi struct{ v, i int }
		a := make([]vi, n)
		for i := range a {
			Fscan(in, &a[i].v)
			a[i].i = i
		}
		sort.Slice(a, func(i, j int) bool { return a[i].v < a[j].v })

		type xa struct {
			x int
			a []vi
		}
		q := []xa{{0, a}}
		for k--; k >= 0; k-- {
			tmp := q
			q = nil
			for _, p := range tmp {
				a := p.a
				for i, n := 0, len(a); i < n; {
					st := i
					x := a[st].v >> k & 1
					for i++; i < n && a[i].v>>k&1 == x; i++ {
					}
					if i-st > 1 {
						q = append(q, xa{p.x | (x^1)<<k, a[st:i]})
					}
				}
			}
			if q == nil {
				mask := 1<<k - 1
				q = tmp
				for _, p := range q {
					a := p.a
					sort.Slice(a, func(i, j int) bool { return a[i].v&mask < a[j].v&mask })
				}
			}
		}
		p := q[0]
		Fprintln(out, p.a[0].i+1, p.a[1].i+1, p.x)
	}
}

//func main() { CF1851F(os.Stdin, os.Stdout) }
