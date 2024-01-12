package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf500C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, ans int
	Fscan(in, &n, &m)
	w := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &w[i])
	}
	b := make([]int, m+1)
	last := make([]int, n+1)
	vis := make([]int, n+1)
	for i := 1; i <= m; i++ {
		Fscan(in, &b[i])
		for _, v := range b[last[b[i]]+1 : i] {
			if vis[v] != i {
				vis[v] = i
				ans += w[v]
			}
		}
		last[b[i]] = i
	}
	Fprint(out, ans)
}

//func main() { cf500C(os.Stdin, os.Stdout) }
