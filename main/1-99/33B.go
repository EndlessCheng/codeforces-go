package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf33B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var s, t, v, w []byte
	var m, wt, ans int
	Fscan(in, &s, &t)
	if len(s) != len(t) {
		Fprint(out, -1)
		return
	}
	
	g := [26][26]int{}
	for i := range g {
		for j := range g[i] {
			if j != i {
				g[i][j] = 1e9
			}
		}
	}
	for Fscan(in, &m); m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		v, w := v[0]-'a', w[0]-'a'
		g[v][w] = min(g[v][w], wt)
	}
	for k := range g {
		for i := range g {
			for j := range g {
				g[i][j] = min(g[i][j], g[i][k]+g[k][j])
			}
		}
	}
	
	for i, v := range s {
		v -= 'a'
		w := t[i] - 'a'
		minJ := 0
		for j := 0; j < 26; j++ {
			if g[v][j]+g[w][j] < g[v][minJ]+g[w][minJ] {
				minJ = j
			}
		}
		ans += g[v][minJ] + g[w][minJ]
		s[i] = 'a' + byte(minJ)
	}
	if ans >= 1e9 {
		Fprint(out, -1)
	} else {
		Fprintf(out, "%d\n%s", ans, s)
	}
}

//func main() { cf33B(os.Stdin, os.Stdout) }
