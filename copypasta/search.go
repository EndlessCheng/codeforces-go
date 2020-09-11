package copypasta

import (
	. "fmt"
	"sort"
)

/* 状态空间
一个实际问题的各种可能情况构成的集合
由小及大：当状态空间位于边界上或某个小范围内等特殊情形，该状态空间的解往往是已知的。
		若能将此解的应用场景扩大到原问题的状态空间，并且扩展过程的每个步骤具有相似性，就可以考虑使用递推或递归求解。
		换句话说，程序在每个步骤上应该面对相同种类的问题，这些问题都是原问题的一个「子问题」，可能仅在规模或者某些限制条件上有所区别，并且能够使用「求解原问题的程序」进行求解。

https://oeis.org/A096969 Number of directed Hamiltonian paths in (n X n)-grid graph
1, 8, 40, 552, 8648, 458696, 27070560, 6046626568, 1490832682992, 1460089659025264, 1573342970540617696, 6905329711608694708440, 33304011435341069362631160, 663618176813467308855850585056, 14527222735920532980525200234503048

https://oeis.org/A236753 Number of simple (non-intersecting) directed paths in (n X n)-grid graph
1, 28, 653, 28512, 3060417, 873239772, 687430009069, 1532025110398168, 9829526954625359697, 183563561823425961932572, 10056737067604248527218979485, 1626248896102138091401810358337184

https://oeis.org/A001411 Number of n-step self-avoiding walks on square lattice
1, 4, 12, 36, 100, 284, 780, 2172, 5916, 16268, 44100, 120292, 324932, 881500, 2374444, 6416596, 17245332, 46466676, 124658732, 335116620, 897697164, 2408806028, 6444560484, 17266613812, 46146397316, 123481354908, 329712786220, 881317491628

Number of simple (non-intersecting) directed paths [of length n] in (n X n)-grid graph
1, 8, 44, 232, 972, 4008, 14932, 55104, 191068, 657848 [10], 2176716, 7157296, 22902052, 72898328, 227471396, 706797600, 2162946116

https://oeis.org/A038373 Number of n-step self-avoiding paths on quadrant grid starting at quadrant origin
1, 2, 4, 10, 24, 60, 146, 366, 912, 2302, 5800, 14722, 37368, 95304, 243168, 622518, 1594622, 4094768, 10521384, 27085436, 69768478, 179982688, 464564220, 1200563864, 3104192722, 8034256412, 20803994184, 53915334890, 139785953076, 362681515714, 941361260956, 2444866458524, 6351963691964

Number of n-step self-avoiding paths on quadrant grid starting at center
1, 2, 8, 20, 64, 172, 520, 1432, 4176, 11504, 32824, 90024, 252992, 690596, 1919328, 5217716, 14380256, 38957328, 106676600

https://oeis.org/A145157 Number of Greek-key tours on an n X n board; i.e., self-avoiding walks on n X n grid starting in top left corner
1, 2, 8, 52, 824, 22144, 1510446, 180160012, 54986690944, 29805993260994, 41433610713353366, 103271401574007978038, 660340630211753942588170, 7618229614763015717175450784, 225419381425094248494363948728158

https://oeis.org/A000532 Number of Hamiltonian paths from NW to SW corners in an n X n grid
1, 1, 2, 8, 86, 1770, 88418, 8934966, 2087813834, 1013346943033, 1111598871478668, 2568944901392936854, 13251059359839620127088, 145194816279817259193401518, 3524171261632305641165676374930
*/

/* 搜索+剪枝
任意子集（不需要剪枝的话可以直接位运算枚举）
部分子集
排列（递归+跳过已经枚举的值）
*/
func searchCollection() {
	// 任意子集：从集合 1~n 中不重复地选取任意个元素
	// 位运算写法见下面的 loopCollection
	// 模板题 https://ac.nowcoder.com/acm/contest/6913/A
	chooseAny := func(n int) {
		{
			cnt := 0
			chosen := []int{}
			var f func(int)
			f = func(p int) {
				if p == n+1 {
					// do chosen... or just cnt++
					cnt++
					return
				}

				// 剪枝：能否继续...

				// 不选 p
				f(p + 1)

				// 选 p
				// 剪枝：能否选 p（是否与 chosen 中的元素冲突等）...

				chosen = append(chosen, p)
				f(p + 1) // 如果可以重复，这里写 f(p)
				chosen = chosen[:len(chosen)-1]
			}
			f(1)
		}

		{
			cnt := 0
			used := make([]bool, n+1)
			var f func(int)
			f = func(p int) {
				if p == n+1 {
					// do used... or just cnt++
					cnt++
					return
				}

				// 剪枝：能否继续...

				// 不选 p
				f(p + 1)

				// 选 p
				// 剪枝：能否选 p（是否与 used 中的元素冲突等）...

				used[p] = true
				f(p + 1)
				used[p] = false
			}
			f(1)
		}
	}

	// 部分子集：从集合 1~n 中不重复地选取至多 m 个元素 (0<=m<=n)
	chooseAtMost := func(n, m int) {
		chosen := []int{}
		var f func(int)
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

	// 可重复组合
	// 以 LC1467/周赛191D 为例 https://leetcode-cn.com/problems/probability-of-a-two-boxes-having-the-same-number-of-distinct-balls/
	// 每个数至多可选 upper[i] 个，从中随机选择 m 个（m<=∑upper），求满足题设条件的概率
	// 枚举每个数选了多少个，根据乘法原理计算某个组合的个数（例如 upper=[4,3,1]，m=4，其中选2个0，2个1就有C(4,2)*C(3,2)种）
	// 总数有 C(∑upper,m) 种
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

	// 排列（不能重复）
	// 即有 n 个位置，从左往右地枚举每个位置上可能出现的值（值必须在 a 中且不能重复）
	// 对比上面的子集搜索，那是对每个位置枚举是否选择（两个分支），而这里每个位置有 n 个分支
	// https://www.luogu.com.cn/problem/P1118
	// LC1307/周赛169D https://leetcode-cn.com/problems/verbal-arithmetic-puzzle/
	searchPermutations := func(a []int) bool {
		n := len(a)
		used := make([]bool, n)
		var f func(p, sum int) bool
		f = func(p, sum int) bool {
			//if sum > ... { } // 剪枝
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
	// https://codeforces.com/problemset/problem/120/H
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
	// https://codeforces.com/problemset/problem/910/C
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
	// todo https://blog.csdn.net/weixin_43914593/article/details/104613920 算法竞赛专题解析（7）：搜索进阶(2)--剪枝

	// A*:
	// todo https://blog.csdn.net/weixin_43914593/article/details/104935011 算法竞赛专题解析（9）：搜索进阶(4)--A*搜索

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

/* 枚举
枚举所有 2^n 子集
枚举子集的所有子集
枚举大小为 k 的子集
枚举格点周围（曼哈顿距离、切比雪夫距离）
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
	loopSet := func(a []int) {
		n := len(a)
		f := func(sub int) (res int) {
			for i, v := range a {
				if sub>>i&1 == 1 {
					// do(v)...
					_ = v

				}
			}
			return
		}
		for sub := 0; sub < 1<<n; sub++ {
			f(sub)
		}
	}

	// 枚举 subset 的全部子集
	// 作为结束条件，处理完 0 之后，会有 -1&subset == subset
	loopSubset := func(n, subset int) {
		sub := subset
		for ok := true; ok; ok = sub != subset {
			// do(sub)...

			sub = (sub - 1) & subset
		}

		{
			// 非空子集
			for sub := subset; sub > 0; sub = (sub - 1) & subset {
				// do(sub)...

			}
		}

		{
			// 真子集
			for sub := (subset - 1) & subset; sub != subset; sub = (sub - 1) & subset {
				// do(sub)...

			}
		}

		{
			// 非空真子集
			for sub := (subset - 1) & subset; sub > 0; sub = (sub - 1) & subset {
				// do(sub)...

			}
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

func gridCollection() {
	type point struct{ x, y int }
	dir4 := []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右
	type pair struct {
		point
		dep int
	}

	// 矩形网格图，返回从起点 (s.x,s.y) 到目标 (t.x,t.y) 的最短距离。'#' 表示无法通过的格子   bfsGrid
	// 无法到达时返回 1e9
	// t 也可是别的东西，比如某个特殊符号等
	// https://ac.nowcoder.com/acm/contest/6781/B
	disST := func(g [][]byte, s point, t point) int {
		n, m := len(g), len(g[0])
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		vis[s.x][s.y] = true
		q := []pair{{s, 0}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			// g[p.x][p.y] == t
			if p.point == t {
				return p.dep
			}
			for _, d := range dir4 {
				if xx, yy := p.x+d.x, p.y+d.y; 0 <= xx && xx < n && 0 <= yy && yy < m && !vis[xx][yy] && g[xx][yy] != '#' { //
					vis[xx][yy] = true
					q = append(q, pair{point{xx, yy}, p.dep + 1})
				}
			}
		}
		return 1e9
	}

	// 从 s 出发寻找 t，返回所有 t 所处的坐标。'#' 表示无法通过的格子
	// https://leetcode-cn.com/contest/season/2020-spring/problems/xun-bao/
	findAllReachableTargets := func(g [][]byte, s point, t byte) (ps []point) {
		n, m := len(g), len(g[0])
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		vis[s.x][s.y] = true
		q := []point{s}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if g[p.x][p.y] == t {
				ps = append(ps, p)
			}
			for _, d := range dir4 {
				if xx, yy := p.x+d.x, p.y+d.y; 0 <= xx && xx < n && 0 <= yy && yy < m && !vis[xx][yy] && g[xx][yy] != '#' { //
					vis[xx][yy] = true
					q = append(q, point{xx, yy})
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
				if v == valid && !vis[i][j] {
					cnt++
					f(i, j)
				}
			}
		}
		return
	}

	// 下列代码来自 LC1254/周赛162C https://leetcode-cn.com/problems/number-of-closed-islands/
	// NOTE: 对于搜索格子的题，可以不用创建 vis 而是通过修改格子的值为范围外的值（如零、负数、'#' 等）来做到这一点
	dfsValidGrids := func(g [][]byte) (comps int) {
		n, m := len(g), len(g[0])
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		const target byte = '.'
		//var targetsPos []point
		var f func(i, j int) bool
		f = func(i, j int) bool {
			if i < 0 || i >= n || j < 0 || j >= m {
				return false
			} // 出边界的不算
			if vis[i][j] || g[i][j] != target {
				return true
			} // 已访问或到达合法边界
			vis[i][j] = true
			//targetsPos = append(targetsPos, point{i, j})
			validCC := true
			for _, d := range dir4 {
				if !f(i+d.x, j+d.y) {
					validCC = false // 遍历完该连通分量再 return，保证不重不漏
				}
			}
			return validCC
		}
		for i, row := range g {
			for j, v := range row {
				if v == target && !vis[i][j] {
					//targetsPos = []point{}
					if f(i, j) {
						comps++
						// do targetsPos...
					}
				}
			}
		}
		return
	}

	// other help functions

	isValidPoint := func(g [][]byte, p point) bool {
		n, m := len(g), len(g[0])
		return 0 <= p.x && p.x < n && 0 <= p.y && p.y < m && g[p.x][p.y] != '#'
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

	findAllTargetsAnyWhere := func(g [][]byte, tar byte) (ps []point) {
		for i, row := range g {
			for j, b := range row {
				if b == tar {
					ps = append(ps, point{i, j})
				}
			}
		}
		return
	}

	_ = []interface{}{
		disST, findAllReachableTargets,
		cntCC, dfsValidGrids,
		isValidPoint, findOneTargetAnyWhere, findAllTargetsAnyWhere,
	}
}
