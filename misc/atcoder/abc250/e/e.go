package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, q, x, y int
	Fscan(in, &n)
	a := make([]int, n+1)
	id := map[int]int{}
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		if id[v] == 0 {
			id[v] = len(id) + 1
		}
		a[i] = max(a[i-1], id[v])
	}

	b := make([]int, n+1)
	set := map[int]struct{}{}
	for i, mx := 1, 0; i <= n; i++ {
		Fscan(in, &v)
		v = id[v]
		if v == 0 {
			mx = 1e9
		}
		mx = max(mx, v)
		set[v] = struct{}{}
		if mx == len(set) {
			b[i] = mx
		}
	}

	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &x, &y)
		if a[x] == b[y] {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
