package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1926G(in io.Reader, out io.Writer) {
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		pa := make([]int, n+1)
		for i := 2; i <= n; i++ {
			Fscan(in, &pa[i])
		}
		Fscan(in, &s)
		f := make([][2]int, n+1)
		for i := n; i > 0; i-- {
			g := f[i]
			if s[i-1] != 'C' {
				g[s[i-1]&1] = 1e9
			}
			f[pa[i]][0] += min(g[0], g[1]+1)
			f[pa[i]][1] += min(g[1], g[0]+1)
		}
		Fprintln(out, min(f[0][0], f[0][1]))
	}
}

//func main() { cf1926G(bufio.NewReader(os.Stdin), os.Stdout) }
