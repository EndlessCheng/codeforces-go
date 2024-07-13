package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://blog.csdn.net/cherrt/article/details/109487315
type fenwick67 struct{ t []int }

func newFenwickTree67(n int) fenwick67 {
	return fenwick67{make([]int, n+1)}
}
func (f fenwick67) add(i, val int) {
	for ; i < len(f.t); i += i & -i {
		f.t[i] += val
	}
}
func (f fenwick67) sum(i int) (s int) {
	for ; i > 0; i &= i - 1 {
		s += f.t[i]
	}
	return
}

func CF1167F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 1_000_000_007
	var n int
	Fscan(in, &n)
	a := make([]struct{ v, i int }, n)
	for i := range a {
		Fscan(in, &a[i].v)
		a[i].i = i + 1
	}
	sort.Slice(a, func(i, j int) bool { return a[i].v < a[j].v })

	ans := 0
	f0, f1 := newFenwickTree67(n), newFenwickTree67(n)
	for _, p := range a {
		i := p.i
		ans = (ans + (f0.sum(i-1)*(n+1-i)+(f1.sum(n)-f1.sum(i)+(n+1-i))*i)%mod*p.v) % mod
		f0.add(i, i)
		f1.add(i, n+1-i)
	}
	Fprint(out, ans)
}

//func main() { CF1167F(os.Stdin, os.Stdout) }
