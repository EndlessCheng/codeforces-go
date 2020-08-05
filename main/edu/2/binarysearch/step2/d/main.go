package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type man struct{ t, n, rest, do int }
	var m, n int
	Fscan(in, &m, &n)
	a := make([]man, n)
	for i := range a {
		Fscan(in, &a[i].t, &a[i].n, &a[i].rest)
	}
	f := func(t int) bool {
		cnt, leftM := 0, m
		for i, p := range a {
			tt := p.t*p.n + p.rest
			c := t / tt * p.n
			if left := t % tt; left < p.t*p.n {
				c += left / p.t
			} else {
				c += p.n
			}
			if c > leftM {
				c = leftM
			}
			leftM -= c
			a[i].do = c
			cnt += c
		}
		return cnt >= m
	}
	t := sort.Search(m*200, f)
	f(t)
	Fprintln(out, t)
	for _, p := range a {
		Fprint(out, p.do, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
