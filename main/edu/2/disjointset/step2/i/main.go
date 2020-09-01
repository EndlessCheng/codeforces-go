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

	var n, q, op, v, w int
	Fscan(in, &n, &q)
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	dis := make([]int8, n)
	var f func(int) int
	f = func(x int) int {
		if fa[x] != x {
			ffx := f(fa[x])
			dis[x] ^= dis[fa[x]]
			fa[x] = ffx
		}
		return fa[x]
	}
	merge := func(from, to int) {
		if ff, ft := f(from), f(to); ff != ft {
			dis[ff] = 1 ^ dis[to] ^ dis[from]
			fa[ff] = ft
		}
	}
	for s := 0; q > 0; q-- {
		Fscan(in, &op, &v, &w)
		v = (v + s) % n
		w = (w + s) % n
		if op == 0 {
			merge(v, w)
		} else {
			f(v)
			f(w)
			if dis[v] == dis[w] {
				s++
				Fprintln(out, "YES")
			} else {
				Fprintln(out, "NO")
			}
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
