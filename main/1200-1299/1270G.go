package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1270G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			g[i] = i - v
		}
		x := 1
		for ; g[x] > 0; x = -g[x] {
			g[x] = -g[x]
		}
		ans := []any{x}
		st := x
		for x = -g[x]; x != st; x = -g[x] {
			ans = append(ans, x)
		}
		Fprintln(out, len(ans))
		Fprintln(out, ans...)
	}
}

//func main() { cf1270G(os.Stdin, os.Stdout) }
