package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1038E(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	f := [4]int{}
	s := [4]int{}
	g := [4][4]bool{}
	for i := range g {
		g[i][i] = true
	}
	minWt := int(1e9)
	for range n {
		var v, wt, w int
		Fscan(in, &v, &wt, &w)
		v--
		w--
		if v != w {
			minWt = min(minWt, wt)
		}
		f[v] ^= 1
		f[w] ^= 1
		s[v] += wt
		g[v][w] = true
		g[w][v] = true
	}

	for k := range 4 {
		for i := range 4 {
			for j := range 4 {
				if g[i][k] && g[k][j] {
					g[i][j] = true
				}
			}
		}
	}

	ok := true
	for i := range 4 {
		if !g[0][i] || f[i] == 0 {
			ok = false
			break
		}
	}

	if ok {
		sum := 0
		for i := 0; i < 4; i++ {
			sum += s[i]
		}
		Fprint(out, sum-minWt)
		return
	}

	ans := 0
	for i := range 4 {
		sum := 0
		for j := range 4 {
			if g[i][j] {
				sum += s[j]
			}
		}
		ans = max(ans, sum)
	}
	Fprint(out, ans)
}

//func main() { cf1038E(bufio.NewReader(os.Stdin), os.Stdout) }
