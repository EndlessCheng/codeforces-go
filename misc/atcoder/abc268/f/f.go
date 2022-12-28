package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, c, ans int
	var s string
	Fscan(in, &n)
	a := make([]struct{ x, s int }, n)
	for i := range a {
		Fscan(in, &s)
		for _, c := range s {
			if c == 'X' {
				a[i].x++
			} else {
				ans += a[i].x * int(c-'0')
				a[i].s += int(c - '0')
			}
		}
	}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.x*b.s > a.s*b.x })
	for _, p := range a {
		ans += c * p.s
		c += p.x
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
