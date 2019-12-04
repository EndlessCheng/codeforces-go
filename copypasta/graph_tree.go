package copypasta

import "math/bits"

// https://oi-wiki.org/graph/lca/#rmq

// https://www.csie.ntu.edu.tw/~hsinmu/courses/_media/dsa_13spring/horowitz_306_311_biconnected.pdf
// low(u) is the lowest dfn that we can reach from u using a path of descendants followed by at most one back edge

// namespace
type tree struct{}

func (*tree) lca(n, root int) {
	g := make([][]int, n)
	for i := 0; i < n-1; i++ {
		var v, w int
		//v, w := read()-1, read()-1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	dfn := make([]int, 0, 2*n-1) // 欧拉序列
	pos := make([]int, n)        // pos[v] 表示 v 在 vs 中第一次出现的位置编号
	depths := make([]int, 0, 2*n-1)
	var dfs func(v, fa, d int)
	dfs = func(v, fa, d int) {
		pos[v] = len(dfn)
		dfn = append(dfn, v)
		depths = append(depths, d)
		for _, w := range g[v] {
			if w != fa {
				dfs(w, v, d+1)
				dfn = append(dfn, v)
				depths = append(depths, d)
			}
		}
	}
	dfs(root, -1, 0)

	type pair struct{ v, i int }
	var st [][20]pair
	stInit := func(a []int) {
		n := len(a)
		st = make([][20]pair, n)
		for i := range st {
			st[i][0] = pair{a[i], i}
		}
		for j := uint(1); 1<<j <= n; j++ {
			for i := 0; i+(1<<j)-1 < n; i++ {
				st0, st1 := st[i][j-1], st[i+(1<<(j-1))][j-1]
				if st0.v < st1.v {
					st[i][j] = st0
				} else {
					st[i][j] = st1
				}
			}
		}
	}
	stQuery := func(l, r int) int { // [l,r] 注意 l r 是从 0 开始算的
		k := uint(bits.Len(uint(r-l+1)) - 1)
		st0, st1 := st[l][k], st[r-(1<<k)+1][k]
		if st0.v < st1.v {
			return st0.i
		}
		return st1.i
	}
	// 注意下标的换算，输出的话要 +1
	calcLCA := func(v, w int) int {
		pv, pw := pos[v], pos[w]
		if pv > pw {
			pv, pw = pw, pv
		}
		return dfn[stQuery(pv, pw)]
	}

	stInit(depths)
	// ...

	_ = calcLCA
}
