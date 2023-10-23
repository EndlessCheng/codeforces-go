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
	maxJ := 0
	for w, num := range cnt {
		maxJ += w * num
		for k1 := 1; num > 0; k1 <<= 1 {
			k := min(k1, num)
			for j := maxJ; j >= k*w; j-- {
				f[j] = min(f[j], f[j-k*w]+k)
			}
			num -= k
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
