package main

// https://space.bilibili.com/206214
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	nxt := make([]int, n-1)
	for i := range nxt {
		nxt[i] = i + 1
	}

	ans := make([]int, len(queries))
	cnt := n - 1
	for qi, q := range queries {
		for i, r := q[0], q[1]; nxt[i] < r; i, nxt[i] = nxt[i], r {
			cnt--
		}
		ans[qi] = cnt
	}
	return ans
}

func shortestDistanceAfterQueries2(n int, queries [][]int) []int {
	fa := make([]int, n-1)
	for i := range fa {
		fa[i] = i
	}
	// 非递归并查集
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}

	ans := make([]int, len(queries))
	cnt := n - 1 // 并查集连通块个数
	for qi, q := range queries {
		l, r := q[0], q[1]-1
		fr := find(r)
		for i := find(l); i < r; i = find(i + 1) {
			fa[i] = fr
			cnt--
		}
		ans[qi] = cnt
	}
	return ans
}
