package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf758E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	type edge struct{ v, w, wt, p int }
	es := make([]edge, n-1)
	type nb struct{ to, i int }
	g := make([][]nb, n+1)
	for i := range es {
		var v, w, wt, p int
		Fscan(in, &v, &w, &wt, &p)
		es[i] = edge{v, w, wt, p}
		g[v] = append(g[v], nb{w, i})
	}

	// 第一次 DFS：预处理每棵子树的 minSum（子树最小重量和）以及 extraDec（见下面注释）
	a := make([]struct{ minSum, extraDec int }, n+1)
	var dfs func(int) (int, int)
	dfs = func(v int) (minSum, maxSum int) {
		for _, e := range g[v] {
			w := e.to
			// 递归计算子树 w 的最小重量和 mn、最大重量和 mx
			mn, mx := dfs(w)
			p := es[e.i].p
			// v-w 边的强度不能小于子树 w 的最小重量和
			if mn < 0 || p < mn {
				return -1, 0
			}
			wt := es[e.i].wt
			// v-w 边的强度可以减少到子树 w 的最小重量和
			minSum += max(wt-(p-mn), 1) + mn
			// 子树 w 的最大重量和不能超过 v-w 边的强度
			maxSum += wt + min(mx, p)
			// 如果 v-w 这条边的强度 p < mx，那么 w 子树的最大重量和 mx 要额外减少 mx-p（减少到 p）
			// 说【额外】是因为 mx 已经是 w 子树内部减重之后的最大重量和了
			// 如果 p < mx，那么 mx 内部还要再减少 mx-p
			a[w].extraDec = max(mx-p, 0)
		}
		a[v].minSum = minSum
		// 最后返回子树 v 的最小重量和、最大重量和
		return
	}
	minSum, _ := dfs(1)
	if minSum < 0 {
		Fprint(out, -1)
		return
	}

	// 第二次 DFS：减重
	// 核心思想：优先减重最下面的边
	// 如果不这样做，先减重上面的，那么由于上面的边强度变小，下面的边也得跟着减重，不如先减重下面的边优
	dec := 0
	var modify func(int)
	modify = func(v int) {
		for _, ew := range g[v] {
			w := ew.to
			// 递归之前，只累加需要减重的量，在递归之后处理减重，这样就可以保证下面的边先减重
			dec += a[w].extraDec
			// 处理 w 子树内部的减重
			modify(w)
			// 处理 v-w 这条边的减重
			e := &es[ew.i]
			// v-w 这条边，重量可以减到 1，强度可以减到子树 w 的最小重量和
			d := min(e.wt-1, e.p-a[w].minSum, dec)
			e.wt -= d
			e.p -= d
			dec -= d
		}
	}
	modify(1)

	Fprintln(out, n)
	for _, e := range es {
		Fprintln(out, e.v, e.w, e.wt, e.p)
	}
}

//func main() { cf758E(bufio.NewReader(os.Stdin), os.Stdout) }
