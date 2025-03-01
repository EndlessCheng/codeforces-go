package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const inf int = 1e9
	var n, q, v, w int
	var s string
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	type pair struct{ dis, s int }
	f := make([][]pair, n)
	for i, v := range a {
		Fscan(in, &s)
		f[i] = make([]pair, n)
		for j, b := range s {
			if b == 'Y' {
				f[i][j] = pair{1, v}
			} else {
				f[i][j].dis = inf
			}
		}
	}

	for k := range f {
		for i := range f {
			for j := range f {
				d := f[i][k].dis + f[k][j].dis
				s := f[i][k].s + f[k][j].s
				if d < f[i][j].dis || d == f[i][j].dis && s > f[i][j].s {
					f[i][j] = pair{d, s}
				}
			}
		}
	}

	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &v, &w)
		p := f[v-1][w-1]
		if p.dis == inf {
			Fprintln(out, "Impossible")
		} else {
			Fprintln(out, p.dis, p.s+a[w-1])
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
