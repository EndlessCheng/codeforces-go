package main

import (
	. "fmt"
	"io"
)

func cf633F(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	var dfs func(int, int) (int, int, int)
	dfs = func(v, fa int) (int, int, int) {
		val := a[v]
		maxChain := val     // 最大链
		maxPathV := val     // 含 v 最大路
		maxPathW := 0       // 不含 v 最大路
		maxChainPath := val // 最大链路
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			chainW, pathW, chainPathW := dfs(w, v)

			// max(含 v 最大路, 不含 v 最大路) + w 路
			// 最大链 + w 链路
			// 最大链路 + w 链
			ans = max(ans, max(maxPathV, maxPathW)+pathW, maxChain+chainPathW, maxChainPath+chainW)

			// 注意一定是链往上，不能是含 v 最大路 + w 链
			// w 链路 + a[v]
			// 不含 v 最大路 + w 链 + a[v]
			// 最大链 + w 路
			maxChainPath = max(maxChainPath, chainPathW+val, maxPathW+chainW+val, maxChain+pathW)

			maxPathV = max(maxPathV, maxChain+chainW)
			maxPathW = max(maxPathW, pathW)
			maxChain = max(maxChain, chainW+val)
		}
		return maxChain, max(maxPathV, maxPathW), maxChainPath
	}
	dfs(0, -1)
	Fprint(out, ans)
}

//func main() { cf633F(bufio.NewReader(os.Stdin), os.Stdout) }
