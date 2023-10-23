package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF95E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, m, v, w int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	vis := make([]bool, n)
	var dfs func(int) int
	dfs = func(v int) int {
		vis[v] = true
		sz := 1
		for _, w := range g[v] {
			if !vis[w] {
				sz += dfs(w)
			}
		}
		return sz
	}
	cnt := map[int]int{}
	for i, b := range vis {
		if !b {
			cnt[dfs(i)]++
		}
	}

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = 1e9
	}
	for w, num := range cnt {
		for rem := 0; rem < w; rem++ {
			type pair struct{ minF, j int }
			q := []pair{}
			for j := 0; j*w+rem <= n; j++ {
				t := f[j*w+rem] - j
				for len(q) > 0 && q[len(q)-1].minF >= t {
					q = q[:len(q)-1]
				}
				q = append(q, pair{t, j})
				f[j*w+rem] = q[0].minF + j
				if j-q[0].j == num {
					q = q[1:]
				}
			}
		}
	}
	ans := int(1e9)
o:
	for i := 1; i <= n; i++ {
		for x := i; x > 0; x /= 10 {
			if x%10 != 4 && x%10 != 7 {
				continue o
			}
		}
		ans = min(ans, f[i])
	}
	if ans == 1e9 {
		Fprint(out, -1)
	} else {
		Fprint(out, ans-1)
	}
}

//func main() { CF95E(os.Stdin, os.Stdout) }
