package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1650C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]struct{ x, w, i int }, m)
		for i := range a {
			Fscan(in, &a[i].x, &a[i].w)
			a[i].i = i
		}
		sort.Slice(a, func(i, j int) bool { return a[i].w < a[j].w })
		a = a[:n*2]
		sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })
		s := 0
		for _, p := range a {
			s += p.w
		}
		Fprintln(out, s)
		for i, p := range a[:n] {
			Fprintln(out, p.i+1, a[n*2-1-i].i+1)
		}
	}
}

//func main() { CF1650C(os.Stdin, os.Stdout) }
