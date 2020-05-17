package copypasta

import "sort"

func searchCollection() {
	// 生成全排列（不保证字典序）
	// 会修改原数组
	// Permute the values at index i to len(arr)-1.
	// https://codeforces.ml/problemset/problem/910/C
	var _permute func([]int, int, func())
	_permute = func(a []int, i int, do func()) {
		if i == len(a) {
			do()
			return
		}
		_permute(a, i+1, do)
		for j := i + 1; j < len(a); j++ {
			a[i], a[j] = a[j], a[i]
			_permute(a, i+1, do)
			a[i], a[j] = a[j], a[i]
		}
	}
	permuteAll := func(a []int, do func()) { _permute(a, 0, do) }

	type _p struct{ x, y int }
	dir4 := [...]_p{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

	type point struct{ x, y int }

	valid := func(g [][]byte, p point) bool {
		n, m := len(g), len(g[0])
		return p.x >= 0 && p.x < n && p.y >= 0 && p.y < m && g[p.x][p.y] != '#'
	}

	// DFS 格点找有多少个连通分量
	// 下列代码来自 LC162C https://leetcode-cn.com/problems/number-of-closed-islands/
	// NOTE: 对于搜索格子的题，可以不用创建 vis 而是通过修改格子的值为范围外的值（如零、负数、'#' 等）来做到这一点
	dfsGrids := func(g [][]byte) (comps int) {
		n, m := len(g), len(g[0])
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		const target byte = '.'
		var targetsPos [][2]int
		var f func(i, j int) bool
		f = func(i, j int) bool {
			if i < 0 || i >= n || j < 0 || j >= m {
				return false
			} // 出边界的不算
			if vis[i][j] || g[i][j] != target {
				return true
			}
			vis[i][j] = true
			targetsPos = append(targetsPos, [2]int{i, j})
			validComp := true
			// 遍历完该连通分量再 return，保证不重不漏
			for _, d := range dir4 {
				if !f(i+d.x, j+d.y) {
					validComp = false
				}
			}
			return validComp
		}
		for i, gi := range g {
			for j, gij := range gi {
				if gij == target && !vis[i][j] {
					targetsPos = [][2]int{}
					if f(i, j) {
						comps++
						// do targetsPos...
					}
				}
			}
		}
		return
	}

	findOneTargetAnyWhere := func(g [][]byte, tar byte) point {
		for i, row := range g {
			for j, b := range row {
				if b == tar {
					return point{i, j}
				}
			}
		}
		return point{-1, -1}
	}

	countTargetAnyWhere := func(g [][]byte, tar byte) (cnt int) {
		for _, row := range g {
			for _, b := range row {
				if b == tar {
					cnt++
				}
			}
		}
		return
	}

	type pair struct {
		point
		dep int
	}

	// 网格图从 (s.x,s.y) 到 (t.x,t.y) 的最短距离，'#' 为障碍物
	// 无法到达时返回 -1
	reachable := func(g [][]byte, s, t point) bool {
		n, m := len(g), len(g[0])
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		vis[s.x][s.y] = true
		for q := []point{s}; len(q) > 0; {
			p := q[0]
			q = q[1:]
			if p == t {
				return true
			}
			for _, d := range dir4 {
				if xx, yy := p.x+d.x, p.y+d.y; xx >= 0 && xx < n && yy >= 0 && yy < m && !vis[xx][yy] && g[xx][yy] != '#' {
					vis[xx][yy] = true
					q = append(q, point{xx, yy})
				}
			}
		}
		return false
	}
	bfsDis := func(g [][]byte, s, t point) int {
		n, m := len(g), len(g[0])
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		vis[s.x][s.y] = true
		for q := []pair{{s, 0}}; len(q) > 0; {
			p := q[0]
			q = q[1:]
			if p.point == t {
				return p.dep
			}
			for _, d := range dir4 {
				if xx, yy := p.x+d.x, p.y+d.y; xx >= 0 && xx < n && yy >= 0 && yy < m && !vis[xx][yy] && g[xx][yy] != '#' {
					vis[xx][yy] = true
					q = append(q, pair{point{xx, yy}, p.dep + 1})
				}
			}
		}
		return -1
	}
	findAllReachableTargets := func(g [][]byte, s point, tar byte) (ps []point) {
		n, m := len(g), len(g[0])
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		vis[s.x][s.y] = true
		for q := []point{s}; len(q) > 0; {
			p := q[0]
			q = q[1:]
			if g[p.x][p.y] == tar {
				ps = append(ps, p)
			}
			for _, d := range dir4 {
				if xx, yy := p.x+d.x, p.y+d.y; xx >= 0 && xx < n && yy >= 0 && yy < m && !vis[xx][yy] && g[xx][yy] != '#' {
					vis[xx][yy] = true
					q = append(q, point{xx, yy})
				}
			}
		}
		return
	}

	// 生成字符串 s 的所有长度至多为 r 的非空子串
	// https://codeforces.ml/problemset/problem/120/H
	genSubStrings := func(s string, r int) []string {
		a := []string{}
		var f func(s, sub string)
		f = func(s, sub string) {
			a = append(a, sub)
			if len(sub) < r {
				for i, b := range s {
					f(s[i+1:], sub+string(b))
				}
			}
		}
		f(s, "")
		a = a[1:] // 去掉空字符串
		sort.Strings(a)
		j := 0
		for i := 1; i < len(a); i++ {
			if a[j] != a[i] {
				j++
				a[j] = a[i]
			}
		}
		return a[:j+1]
	}

	// 从左往右枚举排列（可以剪枝）
	// 即有 n 个位置，从左往右地枚举每个位置上可能出现的值（值必须在 sets 中），且每个位置上的元素不能重复
	// 例题见 LC169D https://leetcode-cn.com/problems/verbal-arithmetic-puzzle/
	dfsPermutations := func(n int, sets []int) bool {
		used := make([]bool, len(sets))
		//used := [10]bool{}
		var f func(cur, x, y int) bool
		f = func(pos, x, y int) bool {
			if pos == n {
				return true // custom
			}
			// 对每个位置，枚举可能出现的值，跳过已经枚举的值
			for i, v := range sets {
				_ = v
				// custom pruning
				//if  {
				//	continue
				//}
				if used[i] {
					continue
				}
				used[i] = true
				// custom calc x y
				if f(pos+1, x, y) {
					return true
				}
				used[i] = false
			}
			return false
		}
		return f(0, 0, 0)
	}

	// 从一个长度为 n 的数组中选择 r 个元素，按字典序生成所有组合，每个组合用下标表示  r <= n
	// 由于实现上直接传入了 indexes，所以在 do 中不能修改 indexes。若要修改则代码在传入前需要 copy 一份
	// 参考 https://docs.python.org/3/library/itertools.html#itertools.combinations
	// https://stackoverflow.com/questions/41694722/algorithm-for-itertools-combinations-in-python
	combinations := func(n, r int, do func(indexes []int)) {
		indexes := make([]int, r)
		for i := range indexes {
			indexes[i] = i
		}
		do(indexes)
		for {
			i := r - 1
			for ; i >= 0; i-- {
				if indexes[i] != i+n-r {
					break
				}
			}
			if i == -1 {
				return
			}
			indexes[i]++
			for j := i + 1; j < r; j++ {
				indexes[j] = indexes[j-1] + 1
			}
			do(indexes)
		}
	}

	// 从一个长度为 n 的数组中选择 r 个元素，按字典序生成所有排列，每个排列用下标表示  r <= n
	// 由于实现上直接传入了 indexes，所以在 do 中不能修改 indexes。若要修改则代码在传入前需要 copy 一份
	// 参考 https://docs.python.org/3/library/itertools.html#itertools.permutations
	permutations := func(n, r int, do func(indexes []int)) {
		indexes := make([]int, n)
		for i := range indexes {
			indexes[i] = i
		}
		do(indexes[:r])
		cycles := make([]int, r)
		for i := range cycles {
			cycles[i] = n - i
		}
		for {
			i := r - 1
			for ; i >= 0; i-- {
				cycles[i]--
				if cycles[i] == 0 {
					tmp := indexes[i]
					copy(indexes[i:], indexes[i+1:])
					indexes[n-1] = tmp
					cycles[i] = n - i
				} else {
					j := cycles[i]
					indexes[i], indexes[n-j] = indexes[n-j], indexes[i]
					do(indexes[:r])
					break
				}
			}
			if i == -1 {
				return
			}
		}
	}

	// 剪枝:
	// todo https://blog.csdn.net/weixin_43914593/article/details/104613920

	// A*:
	// todo https://blog.csdn.net/weixin_43914593/article/details/104935011

	// 舞蹈链
	// TODO: https://oi-wiki.org/search/dlx/
	//       https://leverimmy.blog.luogu.org/dlx-xiang-xi-jiang-jie
	//       https://www.luogu.com.cn/blog/Parabola/qian-tan-shen-xian-suan-fa-dlx

	// 对抗搜索与 Alpha-Beta 剪枝
	// https://www.luogu.com.cn/blog/pks-LOVING/zhun-bei-tou-ri-bao-di-fou-qi-yan-di-blog

	_ = []interface{}{
		valid, dfsGrids, findOneTargetAnyWhere, countTargetAnyWhere, reachable, bfsDis, findAllReachableTargets,
		genSubStrings, dfsPermutations, combinations, permutations, permuteAll,
	}
}
