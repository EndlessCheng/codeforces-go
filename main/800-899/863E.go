package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF863E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]struct{ l, r, i int }, n)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r)
		a[i].i = i
	}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.l < b.l || a.l == b.l && a.r > b.r })
	for i := 1; i < n; i++ {
		if a[i-1].r >= a[i].r {
			Fprint(out, a[i].i+1)
			return
		}
	}
	for i := 1; i < n-1; i++ {
		if a[i-1].r+1 >= a[i+1].l {
			Fprint(out, a[i].i+1)
			return
		}
	}
	Fprint(out, -1)
}

//func main() { CF863E(os.Stdin, os.Stdout) }
