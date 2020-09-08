package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, x, y int
	var s []byte
	Fscan(in, &n, &q)
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	exp := make([]int, n+1)
	var f func(int) int
	f = func(x int) int {
		if fa[x] != x {
			ffx := f(fa[x])
			if fa[x] != ffx {
				exp[x] += exp[fa[x]]
			}
			fa[x] = ffx
		}
		return fa[x]
	}
	for ; q > 0; q-- {
		Fscan(in, &s, &x)
		if s[0] == 'j' {
			Fscan(in, &y)
			if x, y = f(x), f(y); x != y {
				exp[x] -= exp[y]
				fa[x] = y
			}
		} else if s[0] == 'a' {
			Fscan(in, &y)
			exp[f(x)] += y
		} else {
			f(x)
			e := exp[x]
			if fa[x] != x {
				e += exp[fa[x]]
			}
			Fprintln(out, e)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
