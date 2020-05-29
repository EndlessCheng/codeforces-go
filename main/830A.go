package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF830A(_r io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	type pair struct{ l, r int }
	in := bufio.NewReader(_r)
	var n, m, o int
	Fscan(in, &n, &m, &o)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	k := make([]int, m)
	for i := range k {
		Fscan(in, &k[i])
	}
	sort.Ints(k)
	ans := sort.Search(2e9, func(t int) bool {
		ps := []pair{}
		for _, p := range k {
			if d := t - abs(p-o); d >= 0 {
				ps = append(ps, pair{p - d, p + d})
			}
		}
		sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.r < b.r || a.r == b.r && a.l < b.l })
		i := -1
		for _, p := range a {
			for i++; i < len(ps) && ps[i].r < p; i++ {
			}
			if i == len(ps) || p < ps[i].l {
				return false
			}
		}
		return true
	})
	Fprint(out, ans)
}

//func main() { CF830A(os.Stdin, os.Stdout) }
