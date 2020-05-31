package copypasta

import (
	. "fmt"
	"sort"
)

/* 枚举
- 枚举所有 2^n 子集
- 枚举子集的所有子集
- 枚举大小为 k 的子集
- 枚举格点周围（曼哈顿距离、切比雪夫距离）
*/
func loopCollection() {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 枚举 {0,1,...,n-1} 的全部子集
	loopSet := func(arr []int) {
		n := len(arr)
		//outer:
		for sub := 0; sub < 1<<n; sub++ { // sub repr a subset which elements are in range [0,n)
			// do(sub)
			for i := 0; i < n; i++ {
				if sub>>i&1 == 1 { // choose i in sub
					_ = arr[i]
					// do(arr[i]) or continue outer
				}
			}
		}
	}

	// 枚举 subset 的全部子集
	// 作为结束条件，处理完 0 之后，会有 -1&subset == subset
	loopSubset := func(n, subset int) {
		sub := subset
		for ok := true; ok; ok = sub != subset {
			// do(sub)
			sub = (sub - 1) & subset
		}
	}

	// 枚举大小为 n 的集合的大小为 k 的子集（按字典序）
	// 参考《挑战程序设计竞赛》p.156-158
	// 比如在 n 个数中求满足某种性质的最大子集，则可以从 n 开始倒着枚举子集大小，直到找到一个符合性质的子集
	// 例题（TS1）https://codingcompetitions.withgoogle.com/codejam/round/0000000000007706/0000000000045875
	loopSubsetK := func(arr []int, k int) {
		n := len(arr)
		for sub := 1<<k - 1; sub < 1<<n; {
			// do(arr, sub) ...
			x := sub & -sub
			y := sub + x
			sub = sub&^y/x>>1 | y
		}
	}

	/*
		遍历以 (centerI, centerJ) 为中心的曼哈顿距离为 dis 范围内的格点
		例如 dis=2 时：
		  #
		 # #
		# @ #
		 # #
		  #
	*/
	type pair struct{ x, y int }
	dir4 := [...]pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右
	loopAroundManhattan := func(maxI, maxJ, centerI, centerJ, dis int) {
		for i, d := range dir4 {
			d2 := dir4[(i+1)%4]
			dx := d2.x - d.x
			dy := d2.y - d.y
			x := centerI + d.x*dis
			y := centerJ + d.y*dis
			for _i := 0; _i < dis; _i++ {
				if x >= 0 && x < maxI && y >= 0 && y < maxJ {
					// do
				}
				x += dx
				y += dy
			}
		}
	}

	/*
		遍历以 (centerI, centerJ) 为中心的切比雪夫距离为 dis 范围内的格点
		#####
		#   #
		# @ #
		#   #
		#####
	*/
	loopAroundChebyshev := func(maxI, maxJ, centerI, centerJ, dis int) {
		// 上下
		for _, x := range [...]int{centerI - dis, centerI + dis} {
			if x >= 0 && x < maxI {
				for y := max(centerJ-dis, 0); y < min(centerJ+dis, maxJ); y++ {
					// do
				}
			}
		}
		// 左右
		for _, y := range [...]int{centerJ - dis, centerJ + dis} {
			if y >= 0 && y < maxJ {
				for x := max(centerI-dis, 0); x < min(centerI+dis, maxI); x++ {
					// do
				}
			}
		}
	}

	loopDiagonal := func(mat [][]int) {
		n, m := len(mat), len(mat[0])
		for j := 0; j < m; j++ {
			for i := 0; i < n; i++ {
				if i > j {
					break
				}
				_ = mat[i][j-i]
			}
		}
		for i := 1; i < n; i++ {
			for j := m - 1; j >= 0; j-- {
				if i+m-1-j >= n {
					break
				}
				_ = mat[i+m-1-j][j]
			}
		}
	}

	loopDiagonal2 := func(n int) {
		for sum := 0; sum < 2*n-1; sum++ {
			for x := 0; x <= sum; x++ {
				y := sum - x
				if x >= n || y >= n {
					continue
				}
				Println(x, y)
			}
			Println()
		}
	}

	_ = []interface{}{
		loopSet, loopSubset, loopSubsetK,
		loopAroundManhattan, loopAroundChebyshev,
		loopDiagonal, loopDiagonal2,
	}
}

/* 排列组合 剪枝
- 指数型枚举（递归或位运算）
- 组合型枚举（递归+剪枝）
- 排列型枚举（递归+跳过以及枚举的值）
- 搜索+剪枝
*/
func searchCollection() {
	// 指数型枚举：从 1~n 中不重复地选取任意个元素（递归写法）
	chooseAny := func(n int) {
		chosen := []int{}
		var f func(p int)
		f = func(p int) {
			if p == n+1 {
				// do chosen...

				return
			}
			// 不选 p
			f(p + 1)
			// 选 p
			chosen = append(chosen, p)
			f(p + 1)
			chosen = chosen[:len(chosen)-1]
		}
		f(1)
	}

	// 组合型枚举：从 1~n 中不重复地选取至多 m 个元素 (0<=m<=n)
	chooseAtMost := func(n, m int) {
		chosen := []int{}
		var f func(p int)
		f = func(p int) {
			if len(chosen) > m || len(chosen)+n-p+1 < m {
				return
			}
			if p == n+1 {
				// do chosen...

				return
			}
			// 不选 p
			f(p + 1)
			// 选 p
			chosen = append(chosen, p)
			f(p + 1)
			chosen = chosen[:len(chosen)-1]
		}
		f(1)
	}

	// 组合型枚举（可重复）
	// 每个数至多可选 upper[i] 个，从中随机选择 m 个（m<=∑upper），求满足题设条件的概率
	// 枚举每个数选了多少个，根据乘法原理计算某个组合的个数（例如 upper=[4,3,1]，m=4，其中选2个0，2个1就有C(4,2)*C(3,2)种）
	// 总数有 C(∑upper,m) 种
	// 以 LC5427/周赛191D 为例 https://leetcode-cn.com/problems/probability-of-a-two-boxes-having-the-same-number-of-distinct-balls/
	searchCombinations := func(upper []int) float64 {
		const mx = 48
		C := [mx + 1][mx + 1]int{}
		for i := 0; i <= mx; i++ {
			C[i][0], C[i][i] = 1, 1
			for j := 1; j < i; j++ {
				C[i][j] = C[i-1][j-1] + C[i-1][j]
			}
		}

		n := len(upper)
		sum := 0
		for _, v := range upper {
			sum += v
		}
		sum /= 2

		okWays := 0
		var f func(p, s, cntL, cntR, ways int)
		f = func(p, s, cntL, cntR, ways int) {
			//if s > sum {
			//	return
			//}
			if p == n {
				// do...
				if s == sum && cntL == cntR {
					okWays += ways
				}
				return
			}
			for i := 0; i <= upper[p] && s+i <= sum; i++ {
				cl, cr := cntL, cntR
				if i > 0 {
					cl++
				}
				if i < upper[p] {
					cr++
				}
				f(p+1, s+i, cl, cr, ways*C[upper[p]][i]) // 乘法原理
			}
		}
		f(0, 0, 0, 0, 1)
		return float64(okWays) / float64(C[2*sum][sum])
	}

	// 排列型枚举（不能重复）
	// 即有 n 个位置，从左往右地枚举每个位置上可能出现的值（值必须在 a 中且不能重复）
	// 例题见 LC1307/周赛169D https://leetcode-cn.com/problems/verbal-arithmetic-puzzle/
	searchPermutations := func(a []int) bool {
		n := len(a)
		used := make([]bool, n)
		var f func(p, sum int) bool
		f = func(p, sum int) bool {
			if p == n {
				// do sum...

				return sum == 0
			}
			// 对每个位置，枚举可能出现的值，跳过已经枚举的值
			for i, v := range a {
				if used[i] {
					continue
				}
				used[i] = true
				// copy sum and do v...
				s := sum
				s += v
				if f(p+1, s) {
					//used[i] = false
					return true
				}
				used[i] = false
			}
			return false
		}
		return f(0, 0)
	}

	//

	// 生成字符串 s 的所有长度至多为 m 的非空子串（去重，按字典序返回）
	// https://codeforces.ml/problemset/problem/120/H
	genSubStrings := func(s string, m int) []string {
		ss := []string{}
		var f func(s, sub string)
		f = func(s, sub string) {
			ss = append(ss, sub)
			if len(sub) == m {
				return
			}
			for i, b := range s {
				f(s[i+1:], sub+string(b))
			}
		}
		f(s, "")
		ss = ss[1:] // 去掉空字符串
		sort.Strings(ss)
		j := 0
		for i := 1; i < len(ss); i++ {
			if ss[j] != ss[i] {
				j++
				ss[j] = ss[i]
			}
		}
		return ss[:j+1]
	}

	//

	// 从 n 个元素中选择 r 个元素，按字典序生成所有组合，每个组合用下标表示  r <= n
	// 由于实现上直接传入了 indexes，所以在 do 中不能修改 ids。若要修改则代码在传入前需要 copy 一份
	// 参考 https://docs.python.org/3/library/itertools.html#itertools.combinations
	// https://stackoverflow.com/questions/41694722/algorithm-for-itertools-combinations-in-python
	combinations := func(n, r int, do func(ids []int)) {
		ids := make([]int, r)
		for i := range ids {
			ids[i] = i
		}
		do(ids)
		for {
			i := r - 1
			for ; i >= 0; i-- {
				if ids[i] != i+n-r {
					break
				}
			}
			if i == -1 {
				return
			}
			ids[i]++
			for j := i + 1; j < r; j++ {
				ids[j] = ids[j-1] + 1
			}
			do(ids)
		}
	}

	// 从 n 个元素中选择 k 个元素，允许重复选择同一个元素，按字典序生成所有组合，每个组合用下标表示
	// 由于实现上直接传入了 indexes，所以在 do 中不能修改 ids。若要修改则代码在传入前需要 copy 一份
	// 参考 https://docs.python.org/3/library/itertools.html#itertools.combinations_with_replacement
	// https://en.wikipedia.org/wiki/Combination#Number_of_combinations_with_repetition
	// 方案数 H(n,k)=C(n+k-1,k) https://oeis.org/A059481
	// 相当于长度为 k，元素范围在 [0,n-1] 的非降序列的个数
	combinationsWithRepetition := func(n, k int, do func(ids []int)) {
		ids := make([]int, k)
		do(ids)
		for {
			i := k - 1
			for ; i >= 0; i-- {
				if ids[i] != n-1 {
					break
				}
			}
			if i == -1 {
				return
			}
			ids[i]++
			for j := i + 1; j < k; j++ {
				ids[j] = ids[i]
			}
			do(ids)
		}
	}

	// 从一个长度为 n 的数组中选择 r 个元素，按字典序生成所有排列，每个排列用下标表示  r <= n
	// 由于实现上直接传入了 indexes，所以在 do 中不能修改 ids。若要修改则代码在传入前需要 copy 一份
	// 参考 https://docs.python.org/3/library/itertools.html#itertools.permutations
	permutations := func(n, r int, do func(ids []int)) {
		ids := make([]int, n)
		for i := range ids {
			ids[i] = i
		}
		do(ids[:r])
		cycles := make([]int, r)
		for i := range cycles {
			cycles[i] = n - i
		}
		for {
			i := r - 1
			for ; i >= 0; i-- {
				cycles[i]--
				if cycles[i] == 0 {
					tmp := ids[i]
					copy(ids[i:], ids[i+1:])
					ids[n-1] = tmp
					cycles[i] = n - i
				} else {
					j := cycles[i]
					ids[i], ids[n-j] = ids[n-j], ids[i]
					do(ids[:r])
					break
				}
			}
			if i == -1 {
				return
			}
		}
	}

	// 生成全排列（不保证字典序，若要用保证字典序的，见 permutations）
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

	//

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
		chooseAny, chooseAtMost, searchCombinations, searchPermutations,
		genSubStrings,
		combinations, combinationsWithRepetition, permutations, permuteAll,
	}
}

func gridCollection() {
	type _p struct{ x, y int }
	dir4 := [...]_p{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

	type point struct{ x, y int }

	isValidPoint := func(g [][]byte, p point) bool {
		n, m := len(g), len(g[0])
		return 0 <= p.x && p.x < n && 0 <= p.y && p.y < m && g[p.x][p.y] != '#'
	}

	// DFS 格点找有多少个连通分量
	// 下列代码来自 LC1254/周赛162C https://leetcode-cn.com/problems/number-of-closed-islands/
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

	// 网格图，能否从 (s.x,s.y) 到 (t.x,t.y)，'#' 为障碍物
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
				if xx, yy := p.x+d.x, p.y+d.y; 0 <= xx && xx < n && 0 <= yy && yy < m && !vis[xx][yy] && g[xx][yy] != '#' {
					vis[xx][yy] = true
					q = append(q, point{xx, yy})
				}
			}
		}
		return false
	}

	// 网格图，从 (s.x,s.y) 到 (t.x,t.y) 的最短距离，'#' 为障碍物
	// 无法到达时返回 -1
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
				if xx, yy := p.x+d.x, p.y+d.y; 0 <= xx && xx < n && 0 <= yy && yy < m && !vis[xx][yy] && g[xx][yy] != '#' {
					vis[xx][yy] = true
					q = append(q, pair{point{xx, yy}, p.dep + 1})
				}
			}
		}
		return -1
	}

	// 从 s 出发寻找 tar，返回所有 tar 所处的坐标
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
				if xx, yy := p.x+d.x, p.y+d.y; 0 <= xx && xx < n && 0 <= yy && yy < m && !vis[xx][yy] && g[xx][yy] != '#' {
					vis[xx][yy] = true
					q = append(q, point{xx, yy})
				}
			}
		}
		return
	}

	_ = []interface{}{
		isValidPoint,
		dfsGrids,
		findOneTargetAnyWhere, countTargetAnyWhere,
		reachable, bfsDis, findAllReachableTargets,
	}
}
