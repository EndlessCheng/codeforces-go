package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1490E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct {
		v int64
		i int
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &a[i].v)
			a[i].i = i + 1
		}
		sort.Slice(a, func(i, j int) bool { return a[i].v < a[j].v })
		s, st := int64(0), 0
		for i, p := range a {
			if s < p.v {
				st = i
			}
			s += p.v
		}
		a = a[st:]
		Fprintln(out, len(a))
		sort.Slice(a, func(i, j int) bool { return a[i].i < a[j].i })
		for _, v := range a {
			Fprint(out, v.i, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1490E(os.Stdin, os.Stdout) }
