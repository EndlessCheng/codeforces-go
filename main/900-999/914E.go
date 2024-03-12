package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
func cf914E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	var s string
	Fscan(in, &n)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	Fscan(in, &s)

	deleted := make([]bool, n)
	size := make([]int, n)
	var findCentroid func(int, int, int) (int, int, int)
	findCentroid = func(v, fa, compSize int) (minSize, ct, faCt int) {
		minSize = math.MaxInt
		maxSubSize := 0
		size[v] = 1
		for _, w := range g[v] {
			if w != fa && !deleted[w] {
				minSizeW, ctW, faCtW := findCentroid(w, v, compSize)
				if minSizeW < minSize {
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

	ans := make([]int, n)
	cnt := [1 << 20]int{}

	// 更新连通块（子树）信息
	var updateCC func(int, int, int, int)
	updateCC = func(v, fa, delta, mask int) {
		mask ^= 1 << (s[v] - 'a')
		cnt[mask] += delta
		for _, w := range g[v] {
			if w != fa && !deleted[w] {
				updateCC(w, v, delta, mask)
			}
		}
	}

	// 计算「经过 v 向上，在重心拐弯，到其它子树」的路径信息
	var calc func(int, int, int) int
	calc = func(v, fa, mask int) int {
		mask ^= 1 << (s[v] - 'a')
		// 单独计算：从 v 出发的路径信息
		res := cnt[mask]
		for i := 1; i < len(cnt); i <<= 1 {
			res += cnt[mask^i]
		}
		// 把 v 下面的也加上，这样最终算出的是经过 v 的路径信息
		for _, w := range g[v] {
			if w != fa && !deleted[w] {
				res += calc(w, v, mask)
			}
		}
		ans[v] += res
		return res
	}

	var dfs func(int, int, int)
	dfs = func(v, fa, compSize int) {
		_, ct, faCt := findCentroid(v, fa, compSize)

		updateCC(ct, -1, 1, 0)
		// 单独计算：从 ct 出发的路径信息
		res := cnt[0]
		for i := 1; i < len(cnt); i <<= 1 {
			res += cnt[i]
		}
		// 再加上经过 ct 的路径信息
		for _, w := range g[ct] {
			if deleted[w] {
				continue
			}
			// 排除 w 子树后再计算
			updateCC(w, ct, -1, 1<<(s[ct]-'a'))
			res += calc(w, ct, 0)
			updateCC(w, ct, 1, 1<<(s[ct]-'a'))
		}
		ans[ct] += res / 2 // v->w 和 w->v 算了两次
		updateCC(ct, -1, -1, 0)

		deleted[ct] = true
		for _, w := range g[ct] {
			if !deleted[w] {
				if w != faCt {
					dfs(w, ct, size[w])
				} else {
					dfs(w, ct, compSize-size[ct])
				}
			}
		}
	}
	dfs(0, -1, n)
	for _, v := range ans {
		Fprint(out, v+1, " ")
	}
}

//func main() { cf914E(os.Stdin, os.Stdout) }
