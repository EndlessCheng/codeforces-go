package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, ans int
	Fscan(in, &n, &n, &m)
	type pair struct{ v, l int }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].v, &a[i].l)
	}
	b := make([]pair, m)
	for i := range b {
		Fscan(in, &b[i].v, &b[i].l)
	}

	i, j := 0, 0
	for i < n && j < m {
		if a[i].v == b[j].v {
			ans += min(a[i].l, b[j].l)
		}
		if a[i].l > b[j].l {
			a[i].l -= b[j].l
			j++
		} else {
			b[j].l -= a[i].l
			i++
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if b < a { return b }; return a }
