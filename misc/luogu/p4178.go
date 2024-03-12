package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"sort"
)

// https://space.bilibili.com/206214
func p4178(_r io.Reader, out io.Writer) {
	var n, v, w, wt, k int
	in := bufio.NewReader(_r)
	Fscan(in, &n)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}
	Fscan(in, &k)

	markCentroid := make([]bool, len(g))
	size := make([]int, len(g)) // 注：其实只需要保存 ct 的邻居的 size，但这并不好维护
	var findCentroid func(int, int, int) (int, int, int)
	findCentroid = func(v, fa, compSize int) (minSize, ct, faCt int) {
		minSize = math.MaxInt
		maxSubSize := 0
		size[v] = 1
		for _, e := range g[v] {
			if w := e.to; w != fa && !markCentroid[w] {
				if minSizeW, ctW, faCtW := findCentroid(w, v, compSize); minSizeW < minSize {
					minSize, ct, faCt = minSizeW, ctW, faCtW
				}
				maxSubSize = max(maxSubSize, size[w])
				size[v] += size[w]
			}
		}
		maxSubSize = max(maxSubSize, compSize-size[v])
		if maxSubSize < minSize {
			minSize, ct, faCt = maxSubSize, v, fa
		}
		return
	}

	var dfs func(int, int, int) int
	dfs = func(v, fa, compSize int) (ans int) {
		_, ct, faCt := findCentroid(v, fa, compSize)
		markCentroid[ct] = true
		defer func() { markCentroid[ct] = false }()

		// 子问题：统计按 ct 分割后的子树中的点对数
		for _, e := range g[ct] {
			if w := e.to; !markCentroid[w] {
				if w != faCt {
					ans += dfs(w, ct, size[w])
				} else {
					ans += dfs(w, ct, compSize-size[ct])
				}
			}
		}

		countPairs := func(d []int) (res int) {
			sort.Ints(d)
			l, r := 0, len(d)-1
			for l < r {
				if d[l]+d[r] <= k {
					res += r - l
					l++
				} else {
					r--
				}
			}
			return
		}

		// 计算答案：统计经过 ct 的点对数
		// ds[0] = 0 是为了方便统计 ct 和另外一个点的点对数
		dis := make([]int, compSize)
		nd := 1
		for _, e := range g[ct] {
			w := e.to
			if markCentroid[w] {
				continue
			}
			subD := dis[nd:nd]
			var collectDis func(int, int, int)
			collectDis = func(v, fa, d int) {
				subD = append(subD, d)
				for _, e := range g[v] {
					if w := e.to; w != fa && !markCentroid[w] {
						collectDis(w, v, d+e.wt)
					}
				}
			}
			collectDis(w, ct, e.wt)
			nd += len(subD)
			ans -= countPairs(subD)
		}
		ans += countPairs(dis)
		return
	}
	Fprint(out, dfs(0, -1, len(g)))
}

//func main() { p4178(os.Stdin, os.Stdout) }
