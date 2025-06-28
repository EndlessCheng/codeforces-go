package copypasta

import (
	"math"
	"sort"
)

/*
网格/矩阵上的搜索
NOTE: 对于 n*m 的网格图，BFS 最多只占用 O(min(n,m)) 的空间，而 DFS 最多会占用 O(nm) 的空间

网格图 DFS
- [417. 太平洋大西洋水流问题](https://leetcode.cn/problems/pacific-atlantic-water-flow/)
   - https://codeforces.com/problemset/problem/1651/D 1900
- [827. 最大人工岛](https://leetcode.cn/problems/making-a-large-island/) 1934
   - https://codeforces.com/contest/616/problem/C 1600
   - 可以改一排或一列 https://codeforces.com/problemset/problem/1985/H1
   - 可以改一排和一列 https://codeforces.com/problemset/problem/1985/H2
https://codeforces.com/problemset/problem/1948/C 1300
https://codeforces.com/problemset/problem/723/D 1600
https://codeforces.com/problemset/problem/598/D 1700
https://codeforces.com/problemset/problem/1365/D 1700

网格图 BFS
https://codeforces.com/problemset/problem/35/C 1500
https://codeforces.com/problemset/problem/329/B 1500
https://codeforces.com/problemset/problem/2041/D 1700
https://codeforces.com/problemset/problem/1955/H 2300
https://codeforces.com/problemset/problem/1301/F 2600 BFS 进阶玩法
- 同色入队 - 往四周走 - 同色跳跃 - 往四周走 - 同色跳跃 - ...
- 记录访问过的颜色
https://atcoder.jp/contests/abc317/tasks/abc317_e
另见 graph.go 中的 0-1 BFS

综合
- [2258. 逃离火灾](https://leetcode.cn/problems/escape-the-spreading-fire/) 2347
   - https://www.luogu.com.cn/problem/UVA11624
易错题 https://codeforces.com/problemset/problem/540/C 2000

其它
- [54. 螺旋矩阵](https://leetcode.cn/problems/spiral-matrix/)
- [59. 螺旋矩阵 II](https://leetcode.cn/problems/spiral-matrix-ii/)

*/
func _() {
	type pair struct{ x, y int }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

	// 矩形网格图，返回从起点 (s.x,s.y) 到其余所有可达点的最短距离。'#' 表示无法通过的格子   bfsGridAll 单源最短距离
	// https://codeforces.com/contest/1520/problem/G
	// LC2146 https://leetcode.cn/problems/k-highest-ranked-items-within-a-price-range/
	bfsAll := func(g [][]byte, sx, sy int) [][]int {
		n, m := len(g), len(g[0])
		dis := make([][]int, n)
		for i := range dis {
			dis[i] = make([]int, m)
			for j := range dis[i] {
				dis[i][j] = -1
			}
		}
		dis[sx][sy] = 0
		q := []pair{{sx, sy}}
		for step := 1; len(q) > 0; step++ {
			tmp := q
			q = nil
			for _, p := range tmp {
				for _, d := range dir4 {
					if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < n && 0 <= y && y < m && g[x][y] != '#' && dis[x][y] < 0 { //
						dis[x][y] = step
						q = append(q, pair{x, y})
					}
				}
			}
		}
		return dis
	}

	// 返回 (sx,sy) 到其他格子的最短距离
	// 0-1 BFS
	// https://leetcode.cn/problems/grid-teleportation-traversal/
	bfs01 := func(a [][]int, sx, sy int) [][]int {
		n, m := len(a), len(a[0])
		dirs := []pair{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

		dis := make([][]int, n)
		for i := range dis {
			dis[i] = make([]int, m)
			for j := range dis[i] {
				dis[i][j] = math.MaxInt
			}
		}
		dis[sx][sy] = 0

		// 或者写 q := [2][]pair{{{sx, sy}}}
		q0 := []pair{{sx, sy}}
		q1 := []pair{}

		for len(q0) > 0 || len(q1) > 0 {
			var p pair
			if len(q0) > 0 {
				p, q0 = q0[len(q0)-1], q0[:len(q0)-1]
			} else {
				p, q1 = q1[0], q1[1:]
			}
			d := dis[p.x][p.y]
			//if p.x == tx && p.y == ty { return d }

			for _, dir := range dirs {
				x, y := p.x+dir.x, p.y+dir.y
				if 0 <= x && x < n && 0 <= y && y < m && a[x][y] != '#' {
					wt := a[x][y]
					newD := d + wt
					if newD >= dis[x][y] {
						continue
					}
					dis[x][y] = newD
					if wt == 0 {
						q0 = append(q0, pair{x, y})
					} else {
						q1 = append(q1, pair{x, y})
					}
				}
			}
		}

		return dis
	}

	// 矩形网格图，返回从起点 (s.x,s.y) 到目标 (t.x,t.y) 的最短距离。'#' 表示无法通过的格子   bfsGridDep 最短距离
	// 无法到达时返回 inf
	// t 也可是别的东西，比如某个特殊符号等
	// https://ac.nowcoder.com/acm/contest/6781/B
	// https://atcoder.jp/contests/abc184/tasks/abc184_e
	bfsST := func(g [][]byte, sx, sy, tx, ty int) int {
		n, m := len(g), len(g[0])
		const inf int = 1e9 // 1e18

		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		vis[sx][sy] = true
		q := []pair{{sx, sy}}
		for step := 0; len(q) > 0; step++ {
			tmp := q
			q = nil
			for _, p := range tmp {
				// g[p.x][p.y] == 'T'
				if p.x == tx && p.y == ty {
					return step
				}
				for _, d := range dir4 {
					if xx, yy := p.x+d.x, p.y+d.y; 0 <= xx && xx < n && 0 <= yy && yy < m && !vis[xx][yy] && g[xx][yy] != '#' { //
						//if p.x == tx && p.y == ty {
						//	return step
						//}
						vis[xx][yy] = true
						q = append(q, pair{xx, yy})
					}
				}
			}
		}
		return inf
	}

	// 从 s 出发寻找 t，返回所有 t 所处的坐标。'#' 表示无法通过的格子   bfsGrid 可达
	// https://leetcode.cn/contest/season/2020-spring/problems/xun-bao/
	bfsFindAllReachableTargets := func(g [][]byte, s pair, t byte) (ps []pair) {
		n, m := len(g), len(g[0])
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		vis[s.x][s.y] = true
		q := []pair{s}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			x, y := p.x, p.y
			if g[x][y] == t { // x == n-1 && y == m-1
				ps = append(ps, p)
			}
			for _, d := range dir4 {
				if xx, yy := x+d.x, y+d.y; 0 <= xx && xx < n && 0 <= yy && yy < m && !vis[xx][yy] && g[xx][yy] != '#' { //
					vis[xx][yy] = true
					q = append(q, pair{xx, yy})
				}
			}
		}
		return
	}

	// DFS 格点找有多少个连通分量   dfsGrid
	cntCC := func(g [][]byte) (cnt int) {
		n, m := len(g), len(g[0])
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		const valid byte = '.'
		var f func(int, int)
		f = func(x, y int) {
			if x < 0 || x >= n || y < 0 || y >= m || vis[x][y] || g[x][y] != valid {
				return
			}
			vis[x][y] = true
			for _, d := range dir4 {
				f(x+d.x, y+d.y)
			}
		}
		for i, row := range g {
			for j, v := range row {
				if v != valid && !vis[i][j] {
					continue
				}
				cnt++
				f(i, j)
			}
		}
		return
	}

	// 下列代码来自 LC1254 https://leetcode.cn/problems/number-of-closed-islands/
	// NOTE: 对于搜索格子的题，可以不用创建 vis 而是通过修改格子的值为范围外的值（如零、负数、'#' 等）来做到这一点  dfsGrid
	dfsValidGrids := func(g [][]byte) (comps [][]pair) {
		n, m := len(g), len(g[0])
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		const validCell byte = '.'
		var comp []pair
		var f func(int, int) bool
		f = func(i, j int) bool {
			if i < 0 || i >= n || j < 0 || j >= m {
				return false
			}
			if vis[i][j] || g[i][j] != validCell {
				return true
			}
			vis[i][j] = true
			comp = append(comp, pair{i, j})
			validCC := true
			for _, d := range dir4 {
				x, y := i+d.x, j+d.y
				if !f(x, y) {
					validCC = false // 遍历完该连通分量再 return，保证不重不漏
				}
			}
			return validCC
		}
		for i, row := range g {
			for j, v := range row {
				if v != validCell && !vis[i][j] {
					continue
				}
				comp = []pair{}
				if f(i, j) {
					comps = append(comps, comp)
					// do comp ...
				}
			}
		}
		return
	}

	// 周赛 212D https://leetcode.cn/problems/rank-transform-of-a-matrix/
	findSameValueCC := func(mat [][]int) {
		type pair struct{ x, y int }
		type vPos struct {
			v   int
			pos []pair
		}
		allPos := map[int][]pair{}
		for i, row := range mat {
			for j, v := range row {
				allPos[v] = append(allPos[v], pair{i, j})
			}
		}
		vps := []vPos{}
		for v, pos := range allPos {
			np := len(pos)
			g := make([][]int, np)
			for i := 1; i < np; i++ {
				if pos[i].x == pos[i-1].x {
					g[i] = append(g[i], i-1)
					g[i-1] = append(g[i-1], i)
				}
			}
			pid := map[pair]int{}
			col := map[int][]int{} // 按列分组的横坐标
			for i, p := range pos {
				pid[p] = i
				col[p.y] = append(col[p.y], p.x)
			}
			for j, xs := range col {
				for k := 1; k < len(xs); k++ {
					i := pid[pair{xs[k-1], j}]
					i2 := pid[pair{xs[k], j}]
					g[i] = append(g[i], i2)
					g[i2] = append(g[i2], i)
				}
			}
			// 寻找值相同且同行列的所有位置
			var cc []pair
			vis := make([]bool, np)
			var f func(int)
			f = func(v int) {
				vis[v] = true
				cc = append(cc, pos[v])
				for _, w := range g[v] {
					if !vis[w] {
						f(w)
					}
				}
				return
			}
			for i, b := range vis {
				if !b {
					cc = nil
					f(i)
					vps = append(vps, vPos{v, cc})
				}
			}
		}
		sort.Slice(vps, func(i, j int) bool { return vps[i].v < vps[j].v })
		//for _, vp := range vps {
		//	v, pos := vp.v, vp.pos
		//
		//}
	}

	// other help functions

	isValidPoint := func(g [][]byte, p pair) bool {
		n, m := len(g), len(g[0])
		return 0 <= p.x && p.x < n && 0 <= p.y && p.y < m && g[p.x][p.y] != '#'
	}

	findOneTargetAnyWhere := func(g [][]byte, tar byte) pair {
		for i, row := range g {
			for j, b := range row {
				if b == tar {
					return pair{i, j}
				}
			}
		}
		return pair{-1, -1}
	}

	findAllTargetsAnyWhere := func(g [][]byte, tar byte) (ps []pair) {
		for i, row := range g {
			for j, b := range row {
				if b == tar {
					ps = append(ps, pair{i, j})
				}
			}
		}
		return
	}

	_ = []interface{}{
		bfsAll, bfs01, bfsST, bfsFindAllReachableTargets,
		cntCC, dfsValidGrids,
		findSameValueCC,
		isValidPoint, findOneTargetAnyWhere, findAllTargetsAnyWhere,
	}
}
