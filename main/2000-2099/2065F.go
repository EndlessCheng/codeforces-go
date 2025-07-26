package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2065F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		ans := bytes.Repeat([]byte{'0'}, n+1)
		g := make([][]int, n)
		for range n - 1 {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
			if a[v] == a[w] {
				ans[a[v]] = '1'
			}
		}

		vis := make([]int, n+1)
		for i, to := range g {
			for _, j := range to {
				v := a[j]
				if vis[v] == i+1 {
					ans[v] = '1'
				} else {
					vis[v] = i + 1
				}
			}
		}
		Fprintf(out, "%s\n", ans[1:])
	}
}

//func main() { cf2065F(bufio.NewReader(os.Stdin), os.Stdout) }
