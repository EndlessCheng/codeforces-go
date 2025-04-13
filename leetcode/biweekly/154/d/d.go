package main

// https://space.bilibili.com/206214
type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1)
}

// 把下标 i 的元素增加 val
func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

// [1,i] 的元素和
func (f fenwick) pre(i int) (s int) {
	for ; i > 0; i &= i - 1 {
		s += f[i]
	}
	return
}

func treeQueries(n int, edges [][]int, queries [][]int) (ans []int) {
	g := make([][]int, n+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	in := make([]int, n+1)
	out := make([]int, n+1)
	clock := 0
	var dfs func(int, int)
	dfs = func(x, fa int) {
		clock++
		in[x] = clock // 进来的时间
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x)
			}
		}
		out[x] = clock // 离开的时间
	}
	dfs(1, 0)

	// 对于一条边 x-y（y 是 x 的儿子），把边权保存在 weight[y] 中
	weight := make([]int, n+1)
	diff := newFenwickTree(n)
	update := func(x, y, w int) {
		// 保证 y 是 x 的儿子
		if in[x] > in[y] {
			y = x
		}
		d := w - weight[y] // 边权的增量
		weight[y] = w
		// 把子树 y 中的最短路长度都增加 d（用差分树状数组维护）
		diff.update(in[y], d)
		diff.update(out[y]+1, -d)
	}

	for _, e := range edges {
		update(e[0], e[1], e[2])
	}
	for _, q := range queries {
		if q[0] == 1 {
			update(q[1], q[2], q[3])
		} else {
			ans = append(ans, diff.pre(in[q[1]]))
		}
	}
	return
}
