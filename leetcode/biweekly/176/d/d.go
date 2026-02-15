package main

import (
	"math/bits"
	"strconv"
	"strings"
)

// https://space.bilibili.com/206214
// 模板来自我的题单 https://leetcode.cn/circle/discuss/mOr1u6/
type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1) // 使用下标 1 到 n
}

// a[i] ^= val
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f fenwick) update(i int, val int) {
	for ; i < len(f); i += i & -i {
		f[i] ^= val
	}
}

// 计算前缀异或和 a[1] ^ ... ^ a[i]
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res ^= f[i]
	}
	return
}

func palindromePath(n int, edges [][]int, s string, queries []string) (ans []bool) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	mx := bits.Len(uint(n))
	pa := make([][16]int, n)
	dep := make([]int, n)
	timeIn := make([]int, n) // DFS 时间戳
	timeOut := make([]int, n)
	clock := 0
	pathXorFromRoot := make([]int, n) // 从根开始的路径中的字母奇偶性的集合
	pathXorFromRoot[0] = 1 << (s[0] - 'a')

	var dfs func(int, int)
	dfs = func(x, p int) {
		pa[x][0] = p
		clock++
		timeIn[x] = clock
		for _, y := range g[x] {
			if y != p {
				dep[y] = dep[x] + 1
				pathXorFromRoot[y] = pathXorFromRoot[x] ^ 1<<(s[y]-'a')
				dfs(y, x)
			}
		}
		timeOut[x] = clock
	}
	dfs(0, -1)

	for i := range mx - 1 {
		for x := range pa {
			p := pa[x][i]
			if p != -1 {
				pa[x][i+1] = pa[p][i]
			} else {
				pa[x][i+1] = -1
			}
		}
	}

	uptoDep := func(x, d int) int {
		for k := uint32(dep[x] - d); k > 0; k &= k - 1 {
			x = pa[x][bits.TrailingZeros32(k)]
		}
		return x
	}

	// 返回 x 和 y 的最近公共祖先
	getLCA := func(x, y int) int {
		if dep[x] > dep[y] {
			x, y = y, x
		}
		y = uptoDep(y, dep[x]) // 使 y 和 x 在同一深度
		if y == x {
			return x
		}
		for i := mx - 1; i >= 0; i-- {
			px, py := pa[x][i], pa[y][i]
			if px != py {
				x, y = px, py // 同时往上跳 2^i 步
			}
		}
		return pa[x][0]
	}

	// 上面全是模板，下面开始本题逻辑

	t := []byte(s)
	f := newFenwickTree(n) // 注意树状数组是异或运算
	for _, q := range queries {
		if q[0] == 'u' {
			x, _ := strconv.Atoi(q[7 : len(q)-2])
			c := q[len(q)-1]
			val := 1<<(t[x]-'a') ^ 1<<(c-'a') // 擦除旧的，换上新的
			t[x] = c
			// 子树 x 全部异或 val，转换成对区间 [timeIn[x], timeOut[x]] 的差分更新
			f.update(timeIn[x], val)
			f.update(timeOut[x]+1, val)
		} else {
			q = q[6:]
			i := strings.IndexByte(q, ' ')
			x, _ := strconv.Atoi(q[:i])
			y, _ := strconv.Atoi(q[i+1:])
			lca := getLCA(x, y)
			// x 和 y 的 LCA 被抵消了，把 LCA 添加回来
			res := pathXorFromRoot[x] ^ pathXorFromRoot[y] ^ f.pre(timeIn[x]) ^ f.pre(timeIn[y]) ^ 1<<(t[lca]-'a')
			ans = append(ans, res&(res-1) == 0) // 至多一个字母的出现次数是奇数
		}
	}
	return
}
