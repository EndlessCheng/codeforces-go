package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1621E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type group struct {
		a []int
		s int64
		i int
	}

	var T, n, m, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		te := make(sort.IntSlice, n)
		for i := range te {
			Fscan(in, &te[i])
		}
		sort.Sort(sort.Reverse(te))
		te = te[:m]
		gs := make([]group, m)
		for i := range gs {
			Fscan(in, &k)
			a := make([]int, k)
			for j := range a {
				Fscan(in, &a[j])
				gs[i].s += int64(a[j])
			}
			gs[i].a = a
			gs[i].i = i
		}
		sort.Slice(gs, func(i, j int) bool { a, b := gs[i], gs[j]; return a.s*int64(len(b.a)) > b.s*int64(len(a.a)) })

		match := make([]int, m+1)
		matchL := make([]int, m+1)
		matchR := make([]int, m+1)
		for i, v := range te {
			match[i+1] = match[i]
			if int64(v)*int64(len(gs[i].a)) >= gs[i].s {
				match[i+1]++
			}
			matchL[i+1] = matchL[i]
			if i > 0 && int64(v)*int64(len(gs[i-1].a)) >= gs[i-1].s {
				matchL[i+1]++
			}
			matchR[i+1] = matchR[i]
			if i < m-1 && int64(v)*int64(len(gs[i+1].a)) >= gs[i+1].s {
				matchR[i+1]++
			}
		}

		ans := make([][]byte, m)
		for i, g := range gs {
			res := bytes.Repeat([]byte{'0'}, len(g.a))
			for ri, v := range g.a {
				v := int64(v)
				l := int64(len(g.a))
				if v*l == g.s {
					if match[m] == m {
						res[ri] = '1'
					}
					continue
				}
				j := sort.Search(len(te), func(j int) bool { return int64(te[j])*(l-1) < g.s-v }) - 1
				if j < 0 {
					continue
				}
				if v*l < g.s { // avg ↑
					if matchL[i+1]-matchL[j+1] == i-j && match[m]-(match[i+1]-match[j]) == m-(i+1-j) {
						res[ri] = '1'
					}
				} else { // avg ↓
					if matchR[j]-matchR[i] == j-i && match[m]-(match[j+1]-match[i]) == m-(j+1-i) {
						res[ri] = '1'
					}
				}
			}
			ans[g.i] = res
		}
		for _, s := range ans {
			Fprintf(out, "%s", s)
		}
		Fprintln(out)
	}
}

//func main() { CF1621E(os.Stdin, os.Stdout) }
