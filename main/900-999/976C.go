package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf976C(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]struct{ l, r, i int }, n)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r)
		a[i].i = i
	}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.l < b.l || a.l == b.l && a.r > b.r })
	for i := 1; i < n; i++ {
		if a[i].r <= a[i-1].r {
			Fprintln(out, a[i].i+1, a[i-1].i+1)
			return
		}
	}
	Fprintln(out, -1, -1)
}

//func main() { cf976C(bufio.NewReader(os.Stdin), os.Stdout) }
