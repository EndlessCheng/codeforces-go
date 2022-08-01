package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF175C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, t int
	var pre, ans, f int64
	Fscan(in, &n)
	a := make([]struct{ k, c int64 }, n)
	for i := range a {
		Fscan(in, &a[i].k, &a[i].c)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].c < a[j].c })
	Fscan(in, &t)
	p := make([]int64, t, t+1)
	for i := range p {
		Fscan(in, &p[i])
		pre, p[i] = p[i], p[i]-pre
	}
	p = append(p, 1e18)

	for _, kc := range a {
		k, c := kc.k, kc.c
		for ; k > p[f]; f++ {
			ans += p[f] * c * (f + 1)
			k -= p[f]
		}
		ans += k * c * (f + 1)
		p[f] -= k
	}
	Fprint(out, ans)
}

//func main() { CF175C(os.Stdin, os.Stdout) }
