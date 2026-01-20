package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1338D(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	g := make([][]int, n+1)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	var dfs func(int, int) (int, int)
	dfs = func(v, fa int) (int, int) {
		maxL, foot := 0, 1
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			subL, subFoot := dfs(w, v)
			ans = max(ans, maxL+max(subL, subFoot), foot+subL)
			maxL = max(maxL, max(subL, subFoot)+len(g[v])-2)
			foot = max(foot, subL+1)
		}
		ans = max(ans, foot)
		return maxL, foot
	}
	dfs(1, 0)
	Fprint(out, ans)
}

//func main() { cf1338D(bufio.NewReader(os.Stdin), os.Stdout) }
