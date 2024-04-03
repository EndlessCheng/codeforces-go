package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"sort"
)

// https://space.bilibili.com/206214
func cf371E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, s, ss, mnR int
	Fscan(in, &n)
	a := make([]struct{ v, i int }, n)
	for i := range a {
		Fscan(in, &a[i].v)
		a[i].i = i + 1
	}
	Fscan(in, &k)

	sort.Slice(a, func(i, j int) bool { return a[i].v < a[j].v })
	mn := math.MaxInt
	for r, p := range a {
		ss += p.v*min(r-1, k-1) - s
		s += p.v
		if l := r + 1 - k; l >= 0 {
			if ss < mn {
				mn = ss
				mnR = r + 1
			}
			s -= a[l].v
			ss -= s - a[l].v*(k-1)
		}
	}
	for _, p := range a[mnR-k : mnR] {
		Fprint(out, p.i, " ")
	}
}

//func main() { cf371E(os.Stdin, os.Stdout) }
