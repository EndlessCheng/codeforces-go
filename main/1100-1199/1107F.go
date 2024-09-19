package main

import (
	. "fmt"
	"io"
	"math"
	"slices"
)

func cf1107F(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	type offer struct{ a, b, k int }
	a := make([]offer, n)
	for i := range a {
		Fscan(in, &a[i].a, &a[i].b, &a[i].k)
	}
	slices.SortFunc(a, func(p, q offer) int { return q.b - p.b })

	f := make([]int, n+2)
	f[0] = math.MinInt / 2
	for _, t := range a {
		for j := n; j >= 0; j-- {
			f[j+1] = max(f[j+1]+max(t.a-t.b*t.k, 0), f[j]+t.a-t.b*(j-1))
		}
	}
	Fprint(out, slices.Max(f))
}

//func main() { cf1107F(bufio.NewReader(os.Stdin), os.Stdout) }
