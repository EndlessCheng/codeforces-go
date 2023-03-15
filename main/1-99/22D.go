package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF22D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, l, r int
	Fscan(in, &n)
	a := make([]struct{ l, r int }, n)
	for i := range a {
		Fscan(in, &l, &r)
		if l > r {
			l, r = r, l
		}
		a[i].l, a[i].r = l, r
	}
	sort.Slice(a, func(i, j int) bool { return a[i].r < a[j].r })
	ans := []int{-1e9}
	for _, p := range a {
		if p.l > ans[len(ans)-1] {
			ans = append(ans, p.r)
		}
	}
	Fprintln(out, len(ans)-1)
	for _, v := range ans[1:] {
		Fprint(out, v, " ")
	}
}

//func main() { CF22D(os.Stdin, os.Stdout) }
