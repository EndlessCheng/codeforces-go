package main

import (
	"cmp"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1601D(in io.Reader, out io.Writer) {
	var n, d, ans int
	Fscan(in, &n, &d)
	type pair struct{ s, a int }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].s, &a[i].a)
	}
	slices.SortFunc(a, func(a, b pair) int { return cmp.Or(max(a.s, a.a)-max(b.s, b.a), a.s-b.s) })
	for _, p := range a {
		if p.s >= d {
			d = max(d, p.a)
			ans++
		}
	}
	Fprint(out, ans)
}

//func main() { cf1601D(bufio.NewReader(os.Stdin), os.Stdout) }
