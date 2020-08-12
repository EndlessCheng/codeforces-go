package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1074D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	fa, dis := map[int]int{}, map[int]int{}
	var find func(int) int
	find = func(x int) int {
		if fx, ok := fa[x]; ok && fx != x {
			ffx := find(fx)
			dis[x] ^= dis[fx]
			fa[x] = ffx
			return ffx
		}
		return x
	}
	merge := func(from, to, d int) {
		fFrom, fTo := find(from), find(to)
		if fFrom != fTo {
			dis[fFrom] = d ^ dis[from] ^ dis[to]
			fa[fFrom] = fTo
		}
	}
	same := func(x, y int) bool { return find(x) == find(y) }

	var q, t, l, r, x int
	last := 0
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &t, &l, &r)
		l ^= last
		r ^= last
		if l > r {
			l, r = r, l
		}
		r++
		if t == 1 {
			Fscan(in, &x)
			merge(l, r, x^last)
		} else {
			if same(l, r) {
				last = dis[l] ^ dis[r]
				Fprintln(out, last)
			} else {
				last = 1
				Fprintln(out, -1)
			}
		}
	}
}

//func main() { CF1074D(os.Stdin, os.Stdout) }
