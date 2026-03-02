package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf601D(in io.Reader, out io.Writer) {
	var n int
	var s string
	Fscan(in, &n)
	c := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &c[i])
	}
	Fscan(in, &s)
	s = " " + s
	g := make([][]int, n+1)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ch := make([][26]int, n+1)
	f := make([]int, n+1)
	var merge func(int, int) int
	merge = func(x, y int) int {
		if x == 0 || y == 0 {
			if x != 0 {
				return x
			}
			return y
		}
		f[x] = 1
		for i := range 26 {
			ch[x][i] = merge(ch[x][i], ch[y][i])
			f[x] += f[ch[x][i]]
		}
		return x
	}

	ans, cnt := 0, 0
	var dfs func(int, int)
	dfs = func(v, fa int) {
		for _, w := range g[v] {
			if w != fa {
				dfs(w, v)
				ch[v][s[w]-'a'] = merge(ch[v][s[w]-'a'], w)
			}
		}
		f[v] = 1
		for i := range 26 {
			f[v] += f[ch[v][i]]
		}
		res := f[v] + c[v]
		if res > ans {
			ans, cnt = res, 1
		} else if res == ans {
			cnt++
		}
	}
	dfs(1, 0)
	Fprintln(out, ans)
	Fprint(out, cnt)
}

//func main() { cf601D(bufio.NewReader(os.Stdin), os.Stdout) }
