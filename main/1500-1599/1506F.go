package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1506F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ r, c int }

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]pair, n, n+1)
		for i := range a {
			Fscan(in, &a[i].r)
		}
		for i := range a {
			Fscan(in, &a[i].c)
		}
		a = append(a, pair{1, 1})
		sort.Slice(a, func(i, j int) bool { return a[i].r < a[j].r })
		s := 0
		for i, p := range a[:n] {
			if d := a[i+1].r - a[i+1].c - (p.r-p.c)&^1; d > 1 {
				s += d / 2
			} else if d == 0 {
				s += a[i+1].r - p.r
			}
		}
		Fprintln(out, s)
	}
}

//func main() { CF1506F(os.Stdin, os.Stdout) }
