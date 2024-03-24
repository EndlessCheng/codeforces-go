package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf601A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	Fscan(in, &n, &m)
	g := make([][]bool, n)
	for i := range g {
		g[i] = make([]bool, n)
	}
	for ; m > 0; m-- {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v][w] = true
		g[w][v] = true
	}
	tar := !g[0][n-1]
	vis := make([]bool, n)
	vis[0] = true
	q := []int{0}
	for ans := 1; len(q) > 0; ans++ {
		tmp := q
		q = nil
		for _, v := range tmp {
			for w, b := range g[v] {
				if b == tar && !vis[w] {
					if w == n-1 {
						Fprint(out, ans)
						return
					}
					vis[w] = true
					q = append(q, w)
				}
			}
		}
	}
	Fprint(out, -1)
}

//func main() { cf601A(os.Stdin, os.Stdout) }
