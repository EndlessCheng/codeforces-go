package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF659B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct {
		v int
		s string
	}

	var n, m, id, v int
	var s string
	Fscan(in, &n, &m)
	a := make([][]pair, m)
	for ; n > 0; n-- {
		Fscan(in, &s, &id, &v)
		a[id-1] = append(a[id-1], pair{v, s})
	}
	for _, ps := range a {
		sort.Slice(ps, func(i, j int) bool { return ps[i].v > ps[j].v })
		if len(ps) > 2 && ps[1].v == ps[2].v {
			Fprintln(out, "?")
		} else {
			Fprintln(out, ps[0].s, ps[1].s)
		}
	}
}

//func main() { CF659B(os.Stdin, os.Stdout) }
