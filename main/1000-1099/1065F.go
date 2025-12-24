package main

import (
	. "fmt"
	"io"
)

// 把 DFS 写在外面可以避免 MLE

// https://github.com/EndlessCheng
func cf1065F(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		var p int
		Fscan(in, &p)
		g[p-1] = append(g[p-1], w)
	}

	// base = 能回到 v 时，可以访问的叶子数
	// ex = 无法回到 v 时，相比 base 额外访问的叶子数
	var dfs func(int, int) (int, int, int)
	dfs = func(v, dep int) (minUpDep, base, ex int) {
		if g[v] == nil {
			return max(dep-k, 0), 1, 0
		}
		minUpDep = n
		for _, w := range g[v] {
			up, b, e := dfs(w, dep+1)
			if up > dep { // 进入 w 无法回到 v
				ex = max(ex, b+e) // ex 是 w 里的全部
			} else { // 进入 w 可以回到 v
				minUpDep = min(minUpDep, up)
				// 分开统计 b 和 e
				base += b
				ex = max(ex, e)
			}
		}
		return
	}
	_, base, ex := dfs(0, 0)
	Fprint(out, base+ex)
}

//func main() { cf1065F(bufio.NewReader(os.Stdin), os.Stdout) }
