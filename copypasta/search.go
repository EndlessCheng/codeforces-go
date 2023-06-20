package copypasta

import (
	"math/bits"
	"sort"
)

/* 状态空间
一个实际问题的各种可能情况构成的集合
由小及大：当状态空间位于边界上或某个小范围内等特殊情形，该状态空间的解往往是已知的。
若能将此解的应用场景扩大到原问题的状态空间，并且扩展过程的每个步骤具有相似性，就可以考虑使用递推或递归求解。
换句话说，程序在每个步骤上应该面对相同种类的问题，这些问题都是原问题的一个「子问题」，可能仅在规模或者某些限制条件上有所区别，并且能够使用「求解原问题的程序」进行求解。

Self-Avoiding Walk https://mathworld.wolfram.com/Self-AvoidingWalk.html

COUNTING SELF-AVOIDING WALKS https://arxiv.org/pdf/1304.7216.pdf

https://oeis.org/A096969 Number of directed Hamiltonian paths in (n X n)-grid graph
1, 8, 40, 552, 8648, 458696, 27070560, 6046626568, 1490832682992, 1460089659025264, 1573342970540617696, 6905329711608694708440, 33304011435341069362631160, 663618176813467308855850585056, 14527222735920532980525200234503048

https://oeis.org/A236753 Number of simple (non-intersecting) directed paths in (n X n)-grid graph
1, 28, 653, 28512, 3060417, 873239772, 687430009069, 1532025110398168, 9829526954625359697, 183563561823425961932572, 10056737067604248527218979485, 1626248896102138091401810358337184

https://oeis.org/A001411 Number of n-step self-avoiding walks on square lattice
1, 4, 12, 36, 100, 284, 780, 2172, 5916, 16268, 44100, 120292, 324932, 881500, 2374444, 6416596, 17245332, 46466676, 124658732, 335116620, 897697164, 2408806028, 6444560484, 17266613812, 46146397316, 123481354908, 329712786220, 881317491628

https://oeis.org/A046170 Number of self-avoiding walks on a 2-D lattice of length n which start at the origin, take first step in the {+1,0} direction and whose vertices are always nonnegative in x and y
1, 2, 5, 12, 30, 73, 183, 456, 1151, 2900, 7361, 18684, 47652, 121584, 311259, 797311, 2047384, 5260692, 13542718, 34884239, 89991344, 232282110, 600281932, 1552096361, 4017128206, 10401997092, 26957667445, 69892976538

https://oeis.org/A007764 Number of non-intersecting (or self-avoiding) rook paths joining opposite corners of an n X n grid
1, 2, 12, 184, 8512, 1262816, 575780564, 789360053252, 3266598486981642, 41044208702632496804, 1568758030464750013214100, 182413291514248049241470885236, 64528039343270018963357185158482118, 69450664761521361664274701548907358996488

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

https://oeis.org/A000129 Pell numbers: a(0) = 0, a(1) = 1; for n > 1, a(n) = 2*a(n-1) + a(n-2)
https://en.wikipedia.org/wiki/Pell_number
Number of lattice paths from (0,0) to the line x=n-1 consisting of U=(1,1), D=(1,-1) and H=(2,0) steps (i.e., left factors of Grand Schroeder paths)
for example, a(3)=5, counting the paths H, UD, UU, DU and DD
https://oeis.org/A048739 A000129 的前缀和

https://oeis.org/A001333 Number of n-step non-selfintersecting paths starting at (0,0) with steps of types (1,0), (-1,0) or (0,1)
https://codeforces.com/problemset/problem/954/F

*/

/* 搜索+剪枝
任意子集（不需要剪枝的话可以直接位运算枚举）
部分子集
排列（递归+跳过已经枚举的值）
https://leetcode.cn/tag/backtracking/problemset/
https://www.luogu.com.cn/problem/P1379
https://codeforces.com/problemset/problem/429/C
*/
func searchCollection() {
	// 指数型，即 n 层循环
	// https://codeforces.com/contest/459/problem/C
	loopAny := func(n, low, up int) { // or lows ups []int
		vals := make([]int, n)
		var f func(int)
		f = func(p int) {
			if p == n {
				// do vals...

				return
			}
			for vals[p] = low; vals[p] <= up; vals[p]++ {
				f(p + 1)
			}
		}
		f(0)
	}

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
	// 以 LC1467 为例 https://leetcode-cn.com/problems/probability-of-a-two-boxes-having-the-same-number-of-distinct-balls/
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
	// LC1307 https://leetcode-cn.com/problems/verbal-arithmetic-puzzle/
	searchPermutations := func(a []int) bool {
		n := len(a)
		used := 0
		var f func(p, sum int) bool
		f = func(p, sum int) bool {
			//if sum > ... { } // 剪枝
			if p == n {
				// do sum...

				return sum == 0
			}
			// 对每个位置，枚举可能出现的值，跳过已经枚举的值
			for i, v := range a {
				if used>>i&1 > 0 {
					continue
				}
				used |= 1 << i
				// copy sum and do v...
				s := sum
				s += v
				if f(p+1, s) {
					//used[i] = false
					return true
				}
				used ^= 1 << i
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

	// 每个位置独立，枚举 [0,limits[i]] 范围内的数
	iterWithLimits := func(limits []int, do func(upp []int) bool) {
		n := len(limits)
		upp := make([]int, n)
		var f func(p int) bool
		f = func(p int) bool {
			if p == n {
				return do(upp)
			}
			for upp[p] = 0; upp[p] <= limits[p]; upp[p]++ {
				if f(p + 1) {
					return true
				}
			}
			return false
		}
		f(0)
	}

	// 每个位置独立，枚举 [0,limits[i]] 范围内的数，且和为 sum
	iterWithLimitsAndSum := func(sum int, limits []int, do func(a []int) bool) {
		n := len(limits)
		a := make([]int, n)
		var f func(int, int) bool
		f = func(p, s int) bool {
			if s > sum {
				return false
			}
			if p == n {
				if s < sum {
					return false
				}
				return do(a)
			}
			for a[p] = 0; a[p] <= limits[p]; a[p]++ {
				if f(p+1, s+a[p]) {
					return true
				}
			}
			return false
		}
		f(0, 0)
	}

	//

	// 从 n 个元素中选择 r 个元素，按字典序生成所有组合，每个组合用下标表示  r <= n
	// 由于实现上直接传入了 indexes，所以在 do 中不能修改 ids。若要修改则代码在传入前需要 copy 一份
	// 参考 https://docs.python.org/3/library/itertools.html#itertools.combinations
	// https://stackoverflow.com/questions/41694722/algorithm-for-itertools-combinations-in-python
	combinations := func(n, r int, do func(ids []int) (Break bool)) {
		ids := make([]int, r)
		for i := range ids {
			ids[i] = i
		}
		if do(ids) {
			return
		}
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
			if do(ids) {
				return
			}
		}
	}

	// 从 n 个元素中选择 k 个元素，允许重复选择同一个元素，按字典序生成所有组合，每个组合用下标表示
	// 由于实现上直接传入了 indexes，所以在 do 中不能修改 ids。若要修改则代码在传入前需要 copy 一份
	// 参考 https://docs.python.org/3/library/itertools.html#itertools.combinations_with_replacement
	// https://en.wikipedia.org/wiki/Combination#Number_of_combinations_with_repetition
	// 方案数 H(n,k)=C(n+k-1,k) https://oeis.org/A059481
	// 相当于长度为 k，元素范围在 [0,n-1] 的非降序列的个数
	combinationsWithRepetition := func(n, k int, do func(ids []int) (Break bool)) {
		ids := make([]int, k)
		if do(ids) {
			return
		}
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
			if do(ids) {
				return
			}
		}
	}

	// 从一个长度为 n 的数组中选择 r 个元素，按字典序生成所有排列，每个排列用下标表示  r <= n
	// 由于实现上直接传入了 indexes，所以在 do 中不能修改 ids。若要修改则代码在传入前需要 copy 一份
	// 参考 https://docs.python.org/3/library/itertools.html#itertools.permutations
	permutations := func(n, r int, do func(ids []int) (Break bool)) {
		ids := make([]int, n)
		for i := range ids {
			ids[i] = i
		}
		if do(ids[:r]) {
			return
		}
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
					if do(ids[:r]) {
						return
					}
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

	reverse := func(a []int) {
		for i, n := 0, len(a); i < n/2; i++ {
			a[i], a[n-1-i] = a[n-1-i], a[i]
		}
	}
	// 调用完之后
	// 返回 true：a 修改为其下一个排列（即比 a 大且字典序最小的排列）
	// 返回 false：a 修改为其字典序最小的排列（即 a 排序后的结果）
	nextPermutation := func(a []int) bool {
		n := len(a)
		i := n - 2
		for i >= 0 && a[i] >= a[i+1] {
			i--
		}
		defer reverse(a[i+1:])
		if i < 0 {
			return false
		}
		j := n - 1
		for j >= 0 && a[i] >= a[j] {
			j--
		}
		a[i], a[j] = a[j], a[i]
		return true
	}

	// 康托展开 Cantor Expansion
	// 返回所给排列 perm（元素在 [1,n]）的字典序名次（可以从 0 或从 1 开始，具体看代码末尾）
	// 核心思想：对于第 i 个位置，若该位置的数是未出现在其左侧的数中第 k 大的，那么有 (k−1)×(N−i)! 种方案在该位置上比这个排列小
	// 结合康托展开和逆康托展开，可以求出一个排列的下 k 个排列
	// https://zh.wikipedia.org/wiki/%E5%BA%B7%E6%89%98%E5%B1%95%E5%BC%80
	// https://oi-wiki.org/math/cantor/
	// https://www.luogu.com.cn/problem/P5367
	// 有重复元素 LC1830 https://leetcode-cn.com/problems/minimum-number-of-operations-to-make-string-sorted/
	// https://codeforces.com/problemset/problem/1443/E
	rankPermutation := func(perm []int) int64 {
		const mod int64 = 1e9 + 7
		n := len(perm)
		F := make([]int64, n)
		F[0] = 1
		for i := 1; i < n; i++ {
			F[i] = F[i-1] * int64(i) % mod
		}

		tree := make([]int, n+1)
		add := func(i, val int) {
			for ; i <= n; i += i & -i {
				tree[i] += val
			}
		}
		sum := func(i int) (res int) {
			for ; i > 0; i &= i - 1 {
				res += tree[i]
			}
			return
		}
		for i := 1; i <= n; i++ {
			add(i, 1)
		}

		ans := int64(0)
		for i, v := range perm {
			ans += int64(sum(v-1)) * F[n-1-i] % mod
			add(v, -1)
		}
		ans++ // 从 1 开始的排名
		ans %= mod
		return ans
	}

	// 逆康托展开 Inverse Cantor Expansion
	// 返回字典序第 k 小的排列，元素范围为 [1,n]
	// LC60 https://leetcode-cn.com/problems/permutation-sequence/
	// https://codeforces.com/problemset/problem/1443/E
	kthPermutation := func(n, k int) []int {
		F := make([]int, n)
		F[0] = 1
		for i := 1; i < n; i++ {
			F[i] = F[i-1] * i
		}

		k-- // 如果输入是从 1 开始的，改成从 0 开始
		perm := make([]int, n)
		valid := make([]int, n+1)
		for i := 1; i <= n; i++ {
			valid[i] = 1
		}
		for i := 1; i <= n; i++ {
			order := k/F[n-i] + 1
			for j := 1; j <= n; j++ { // 从 1 开始的排列    TODO 用线段树优化
				order -= valid[j]
				if order == 0 {
					perm = append(perm, j)
					valid[j] = 0
					break
				}
			}
			k %= F[n-i]
		}
		return perm
	}

	// 迭代加深搜索
	// 限制 DFS 深度（不断提高搜索深度）
	// http://poj.org/problem?id=2248

	// 折半枚举/双向搜索 Meet in the middle
	// https://codeforces.com/problemset/problem/1006/F https://atcoder.jp/contests/abc271/tasks/abc271_f https://leetcode.com/discuss/interview-question/2324457/Google-Online-Assessment-Question
	// LC805 https://leetcode.cn/problems/split-array-with-same-average/
	// LC2035 https://leetcode.cn/problems/partition-array-into-two-arrays-to-minimize-sum-difference/
	// O(3^(n/2)) 放A组/放B组/不选 https://www.luogu.com.cn/problem/P3067 https://www.luogu.com.cn/record/88785388
	// https://www.luogu.com.cn/problem/P5194
	// https://www.luogu.com.cn/problem/P4799
	// https://codeforces.com/problemset/problem/327/E
	// https://atcoder.jp/contests/abc184/tasks/abc184_f

	// 折半枚举 - 超大背包问题
	// https://atcoder.jp/contests/abc184/tasks/abc184_f
	bigKnapsack := func(a []int, size int) (ans int) {
		n := len(a)
		if n == 1 {
			if a[0] > size {
				return
			}
			return a[0]
		}

		sumW, ws, end := 0, []int{}, n/2
		var f func(int)
		f = func(p int) {
			if p == end {
				if sumW <= size {
					ws = append(ws, sumW)
				}
				return
			}
			f(p + 1)
			sumW += a[p]
			f(p + 1)
			sumW -= a[p]
		}
		f(0)
		l := ws
		sort.Ints(l)
		// l 去重略

		ws, end = nil, n
		f(n / 2)
		for _, w := range ws {
			// <= size-w 的第一个数（因为 l[0]==0 所以 p 一定非负）
			p := sort.SearchInts(l, size-w+1) - 1
			if l[p]+w > ans {
				ans = l[p] + w
			}
		}
		return
	}

	type pair struct{ w, v int }
	bigKnapsack2 := func(items []pair, size int) (ans int) {
		n := len(items)
		if n == 1 {
			if items[0].w > size {
				return
			}
			return items[0].v
		}

		sumW, sumV, ps, end := 0, 0, []pair{}, n/2
		var f func(int)
		f = func(p int) {
			if p == end {
				ps = append(ps, pair{sumW, sumV})
				return
			}
			f(p + 1)
			it := items[p]
			sumW += it.w
			sumV += it.v
			f(p + 1)
			sumV -= it.v
			sumW -= it.w
		}
		f(0)

		// 去重，确保重量越大，价值严格越大
		l := ps
		nl := 1
		for i := 1; i < len(l); i++ {
			if l[nl-1].v < l[i].v {
				l[nl] = l[i]
				nl++
			}
		}
		l = l[:nl]

		ps, end = nil, n
		f(n / 2)
		for _, p := range ps {
			// <= size-p.w 的第一个数（因为 l[0].w==0 所以 i 一定非负）
			i := sort.Search(len(l), func(i int) bool { return l[i].w+p.w > size }) - 1
			if l[i].v+p.v > ans {
				ans = l[i].v + p.v
			}
		}
		return
	}

	// 剪枝:
	// todo https://blog.csdn.net/weixin_43914593/article/details/104613920 算法竞赛专题解析（7）：搜索进阶(2)--剪枝

	// A*:
	// todo https://blog.csdn.net/weixin_43914593/article/details/104935011 算法竞赛专题解析（9）：搜索进阶(4)--A*搜索
	//  https://www.redblobgames.com/pathfinding/a-star/introduction.html

	// 舞蹈链 Dancing Links 精确覆盖问题
	// https://en.wikipedia.org/wiki/Dancing_Links
	// TODO: https://oi-wiki.org/search/dlx/
	//       https://leverimmy.blog.luogu.org/dlx-xiang-xi-jiang-jie
	//       https://www.luogu.com.cn/blog/Parabola/qian-tan-shen-xian-suan-fa-dlx
	//       https://www.cnblogs.com/grenet/p/3145800.html
	//       https://www.cnblogs.com/grenet/p/3163550.html
	//   https://lsr2002.blog.luogu.org/wu-dao-lian
	// 模板题+讲解
	//       todo http://hihocoder.com/contest/hiho101/problem/1
	//       http://hihocoder.com/contest/hiho102/problem/1
	//       https://www.luogu.com.cn/problem/P4929

	// 对抗搜索与 Alpha-Beta 剪枝
	// https://www.luogu.com.cn/blog/pks-LOVING/zhun-bei-tou-ri-bao-di-fou-qi-yan-di-blog

	_ = []interface{}{
		loopAny, chooseAny, chooseAtMost, searchCombinations, searchPermutations,
		genSubStrings,
		iterWithLimits, iterWithLimitsAndSum,
		combinations, combinationsWithRepetition,
		permutations, permuteAll, nextPermutation, rankPermutation, kthPermutation,
		bigKnapsack, bigKnapsack2,
	}
}

/* 枚举
枚举所有 2^n 子集
枚举子集的所有子集
枚举大小为 k 的子集
枚举格点周围（曼哈顿距离、切比雪夫距离）
*/
func _(min, max func(int, int) int) {
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

	// 枚举 set 的全部子集
	// 作为结束条件，处理完 0 之后，会有 -1&set == set
	//
	// 你可能会好奇，为什么 sub = (sub - 1) & set 这样写一定可以「跳到」下一个子集呢？会不会漏呢？
	// 因为二进制的减法的特点是，每次会把 lowbit 那个 1 改成 0，lowbit 右边的 0 全部改成 1
	// 由于下一个子集必然比 sub 小，减法的这种特点可以保证 sub-1 之后的二进制数必然包含下一个子集
	loopSubset := func(n, set int) {
		// 所有子集
		for sub, ok := set, true; ok; ok = sub != set {
			// do(sub)...

			sub = (sub - 1) & set
		}

		// 所有子集（写法二）
		for sub := set; ; sub = (sub - 1) & set {
			// do(sub)...

			if sub == 0 {
				break
			}
		}

		// 非空子集
		for sub := set; sub > 0; sub = (sub - 1) & set {
			// do(sub)...

		}

		// 真子集
		for sub := (set - 1) & set; sub != set; sub = (sub - 1) & set {
			// do(sub)...

		}

		// 非空真子集
		for sub := (set - 1) & set; sub > 0; sub = (sub - 1) & set {
			// do(sub)...

		}

		{
			// EXTRA: 求多个集合（状压）的所有非空子集组成的集合
			// https://ac.nowcoder.com/acm/contest/7607/B
			has := [1e6 + 1]bool{0: true}
			var f func(uint)
			f = func(v uint) {
				if has[v] {
					return
				}
				has[v] = true
				for w := v; w > 0; w &= w - 1 {
					f(v ^ w&-w)
				}
			}
			//for _, v := range a {
			//	f(v)
			//}
		}
	}

	// 枚举 set 的全部超集（父集）ss
	loopSuperset := func(n, set int) {
		for ss := set; ss < 1<<n; ss = (ss + 1) | set {
			// do(ss)...

		}
	}

	// Gosper's Hack：枚举大小为 n 的集合的大小为 k 的子集（按字典序）
	// 我的视频讲解 https://www.bilibili.com/video/BV1na41137jv?t=15m43s
	// https://en.wikipedia.org/wiki/Combinatorial_number_system#Applications
	// 比如在 n 个数中求满足某种性质的最大子集，则可以从 n 开始倒着枚举子集大小，直到找到一个符合性质的子集
	// 例题（TS1）GCJ 2018 R2 Costume Change https://codingcompetitions.withgoogle.com/codejam/round/0000000000007706/0000000000045875
	loopSubsetK := func(a []int, k int) {
		n := len(a)
		for sub := 1<<k - 1; sub < 1<<n; {
			// do(a, sub) ...
			lb := sub & -sub
			x := sub + lb
			// 下式等价于 sub = (sub^x)/lb>>2 | x
			// 把除法改成右移 bits.TrailingZeros 可以快好几倍
			sub = (sub^x)>>bits.TrailingZeros(uint(lb))>>2 | x
		}
	}

	// 枚举各个 1 位的另一种方法
	// 每次统计尾 0 的个数，然后移除最右侧的 1
	// benchmark 了一下，效率比一个个位上去检查是否为 1 要快
	{
		var mask uint
		for ; mask > 0; mask &= mask - 1 {
			p := bits.TrailingZeros(mask)
			_ = p
		}
	}

	//

	// 获取螺旋遍历的所有坐标         螺旋矩阵 Spiral Matrix
	// LC54 https://leetcode.cn/problems/spiral-matrix/
	// LC59 https://leetcode.cn/problems/spiral-matrix-ii/
	// LC885 https://leetcode.cn/problems/spiral-matrix-iii/
	// LC2326 https://leetcode.cn/problems/spiral-matrix-iv/
	// https://ac.nowcoder.com/acm/contest/6489/C
	type pair struct{ x, y int }
	loopSpiralMatrix := func(n, m int) []pair { // n 行 m 列
		dir4 := []struct{ x, y int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上
		mat := make([][]int, n)
		for i := range mat {
			mat[i] = make([]int, m)
			for j := range mat[i] {
				mat[i][j] = -1
			}
		}
		pos := make([]pair, n*m)
		i, j, di := 0, 0, 0
		for id := 0; id < n*m; id++ {
			pos[id] = pair{i, j}
			mat[i][j] = id
			d := dir4[di]
			if x, y := i+d.x, j+d.y; x < 0 || x >= n || y < 0 || y >= m || mat[x][y] != -1 {
				di = (di + 1) % 4
				d = dir4[di]
			}
			i += d.x
			j += d.y
		}
		return pos
	}

	// 顺时针遍历矩阵从外向内的第 d 圈（保证不自交）
	// LC1914 https://leetcode-cn.com/problems/cyclically-rotating-a-grid/
	loopAround := func(a [][]int, d int) []int {
		n, m := len(a), len(a[0])
		b := make([]int, 0, (n+m-d*4-2)*2)
		for j := d; j < m-d; j++ { // →
			b = append(b, a[d][j])
		}
		for i := d + 1; i < n-d; i++ { // ↓
			b = append(b, a[i][m-1-d])
		}
		for j := m - d - 2; j >= d; j-- { // ←
			b = append(b, a[n-1-d][j])
		}
		for i := n - d - 2; i > d; i-- { // ↑
			b = append(b, a[i][d])
		}
		return b
	}

	// 获取之字遍历的所有坐标
	loopZigZag := func(n, m int) []pair { // n 行 m 列
		pos := make([]pair, 0, n*m)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				pos = append(pos, pair{i, j})
			}
			i++
			if i == n {
				break
			}
			for j := m - 1; j >= 0; j-- {
				pos = append(pos, pair{i, j})
			}
		}
		return pos
	}

	/*
		遍历以 (ox, oy) 为中心的曼哈顿距离为 dis 范围内的格点
		例如 dis=2 时：
		  #
		 # #
		# @ #
		 # #
		  #
	*/
	dir4r := []struct{ x, y int }{{-1, 1}, {-1, -1}, {1, -1}, {1, 1}} // 逆时针
	loopAroundManhattan := func(n, m, ox, oy, dis int, f func(x, y int)) {
		if dis == 0 {
			f(ox, oy)
			return
		}
		x, y := ox+dis, oy // 从最右顶点出发，逆时针移动
		for _, d := range dir4r {
			for k := 0; k < dis; k++ {
				if 0 <= x && x < n && 0 <= y && y < m {
					f(x, y)
				}
				x += d.x
				y += d.y
			}
		}
	}

	// 曼哈顿圈序遍历
	// LC1030 https://leetcode-cn.com/problems/matrix-cells-in-distance-order/
	loopAllManhattan := func(n, m, ox, oy int, f func(x, y int)) {
		f(ox, oy)
		maxDist := max(ox, n-1-ox) + max(oy, m-1-oy)
		for dis := 1; dis <= maxDist; dis++ {
			x, y := ox+dis, oy // 从最右顶点出发，逆时针移动
			for _, d := range dir4r {
				for k := 0; k < dis; k++ {
					if 0 <= x && x < n && 0 <= y && y < m {
						f(x, y)
					}
					x += d.x
					y += d.y
				}
			}
		}
	}

	/*
		遍历以 (ox, oy) 为中心的切比雪夫距离为 dis 范围内的格点
		#####
		#   #
		# @ #
		#   #
		#####
	*/
	loopAroundChebyshev := func(n, m, ox, oy, dis int) {
		// 上下
		for _, x := range []int{ox - dis, ox + dis} {
			if 0 <= x && x < n {
				for y := max(oy-dis, 0); y <= min(oy+dis, m-1); y++ {
					// do ...
				}
			}
		}
		// 左右（注意四角已经被上面的循环枚举到了）
		for _, y := range []int{oy - dis, oy + dis} {
			if 0 <= y && y < m {
				for x := max(ox-dis, 0) + 1; x <= min(ox+dis, n-1)-1; x++ {
					// do ...
				}
			}
		}
	}

	// 第一排在右上，最后一排在左下
	// 每排从左上到右下
	// LC2711 https://leetcode.cn/problems/difference-of-number-of-distinct-values-on-diagonals/
	loopDiagonal := func(n, m int) {
		for s := 1; s < n+m; s++ {
			minJ := max(0, m-s)
			maxJ := min(m-1, n+m-1-s)
			for j := minJ; j <= maxJ; j++ {
				i := s + j - m
				_ = i

			}
		}
	}

	// 第一排在左上，最后一排在右下
	// 每排从左下到右上
	// LC498 https://leetcode.cn/problems/diagonal-traverse/
	loopAntiDiagonal := func(n, m int) {
		for s := 0; s < n+m-1; s++ {
			minJ := max(0, s-n+1)
			maxJ := min(m-1, s)
			for j := minJ; j <= maxJ; j++ {
				i := s - j
				_ = i

			}
		}
	}

	// 以主对角线为第一列（行），然后向右（下）平移遍历
	// 例如
	// 0 3 6 9
	// 10 1 4 7
	// 8 11 2 5
	// https://codeforces.com/problemset/problem/1276/C
	circleLoopDiagonal := func(n, m int) {
		if n <= m {
			// 向右平移
			for rc := 0; rc < n*m; rc++ {
				_c, _r := rc/n, rc%n
				i, j := _r, (_c+_r)%m
				_, _ = i, j

			}
		} else {
			// 向下平移
			for rc := 0; rc < n*m; rc++ {
				_r, _c := rc/m, rc%m
				i, j := (_r+_c)%n, _c
				_, _ = i, j

			}
		}
	}

	// 保证边界在范围内且 x0 <= x1 且 y0 <= y1
	loopBorder := func(x0, y0, x1, y1 int) {
		if y0 == y1 {
			for i := x0; i <= x1; i++ {
				// do(i, y0) ...

			}
			return
		}
		for i := x0; i <= x1; i++ {
			for j := y0; j <= y1; {
				// do(i, j) ...

				if i == x0 || i == x1 {
					j++
				} else {
					j += y1 - y0
				}
			}
		}
	}

	_ = []interface{}{
		loopSet, loopSubset, loopSuperset, loopSubsetK,
		loopSpiralMatrix, loopAround, loopZigZag,
		loopAroundManhattan, loopAllManhattan, loopAroundChebyshev,
		loopDiagonal, loopAntiDiagonal, circleLoopDiagonal,
		loopBorder,
	}
}

/*
## 题单

#### 网格图 DFS

- [200. 岛屿数量](https://leetcode.cn/problems/number-of-islands/)
- [1254. 统计封闭岛屿的数目](https://leetcode.cn/problems/number-of-closed-islands/)
- [1020. 飞地的数量](https://leetcode.cn/problems/number-of-enclaves/)
- [695. 岛屿的最大面积](https://leetcode.cn/problems/max-area-of-island/)
- [面试题 16.19. 水域大小](https://leetcode.cn/problems/pond-sizes-lcci/)
- [463. 岛屿的周长](https://leetcode.cn/problems/island-perimeter/)
- [130. 被围绕的区域](https://leetcode.cn/problems/surrounded-regions/)
- [417. 太平洋大西洋水流问题](https://leetcode.cn/problems/pacific-atlantic-water-flow/)
- [529. 扫雷游戏](https://leetcode.cn/problems/minesweeper/)
- [827. 最大人工岛](https://leetcode.cn/problems/making-a-large-island/)
- [1034. 边界着色](https://leetcode.cn/problems/coloring-a-border/)

#### 网格图 BFS

- [542. 01 矩阵](https://leetcode.cn/problems/01-matrix/)
- [934. 最短的桥](https://leetcode.cn/problems/shortest-bridge/)
- [994. 腐烂的橘子](https://leetcode.cn/problems/rotting-oranges/)
- [1162. 地图分析](https://leetcode.cn/problems/as-far-from-land-as-possible/)
- [2146. 价格范围内最高排名的 K 样物品](https://leetcode.cn/problems/k-highest-ranked-items-within-a-price-range/)
- [2258. 逃离火灾](https://leetcode.cn/problems/escape-the-spreading-fire/)
- [2577. 在网格图中访问一个格子的最少时间](https://leetcode.cn/problems/minimum-time-to-visit-a-cell-in-a-grid/)

#### 综合应用

- [778. 水位上升的泳池中游泳](https://leetcode.cn/problems/swim-in-rising-water/)
- [1631. 最小体力消耗路径](https://leetcode.cn/problems/path-with-minimum-effort/)
- [1263. 推箱子](https://leetcode.cn/problems/minimum-moves-to-move-a-box-to-their-target-location/)
- [LCP 75. 传送卷轴](https://leetcode.cn/problems/rdmXM7/)
*/

// 网格/矩阵上的搜索
// NOTE: 对于 n*m 的网格图，BFS 最多只占用 O(min(n,m)) 的空间，而 DFS 最多会占用 O(nm) 的空间
// 易错题 https://codeforces.com/problemset/problem/540/C
// 思维转换 LCP31 https://leetcode-cn.com/problems/Db3wC1/
// LC778 https://leetcode.cn/problems/swim-in-rising-water/
// LC1631 https://leetcode.cn/problems/path-with-minimum-effort/
// BFS:
// LC542 https://leetcode.cn/problems/01-matrix/
// LC994 https://leetcode.cn/problems/rotting-oranges/
// LC1162 https://leetcode.cn/problems/as-far-from-land-as-possible/
// LC2146 https://leetcode.cn/problems/k-highest-ranked-items-within-a-price-range/
// LC2258 https://leetcode.cn/problems/escape-the-spreading-fire/
// LC2577 https://leetcode.cn/problems/minimum-time-to-visit-a-cell-in-a-grid/
// LCP13 https://leetcode.cn/problems/xun-bao/
// LCP75 https://leetcode.cn/problems/rdmXM7/
func gridCollection() {
	type pair struct{ x, y int }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

	// 获取网格图的 labelS, labelT 的坐标
	getST := func(g [][]byte, labelS, labelT byte) (pair, pair) {
		var s, t pair
		for i, row := range g {
			for j, b := range row {
				p := pair{i, j}
				if b == labelS {
					s = p
				} else if b == labelT {
					t = p
				}
			}
		}
		return s, t
	}

	// 矩形网格图，返回从起点 (s.x,s.y) 到其余所有可达点的最短距离。'#' 表示无法通过的格子   bfsGridAll 单源最短距离
	// https://codeforces.com/contest/1520/problem/G
	// LC2146 https://leetcode-cn.com/problems/k-highest-ranked-items-within-a-price-range/
	disAll := func(g [][]byte, sx, sy int) [][]int {
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

	// 矩形网格图，返回从起点 (s.x,s.y) 到目标 (t.x,t.y) 的最短距离。'#' 表示无法通过的格子   bfsGridDep 最短距离
	// 无法到达时返回 inf
	// t 也可是别的东西，比如某个特殊符号等
	// https://ac.nowcoder.com/acm/contest/6781/B
	// https://atcoder.jp/contests/abc184/tasks/abc184_e
	disST := func(g [][]byte, sx, sy, tx, ty int) int {
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
	// https://leetcode-cn.com/contest/season/2020-spring/problems/xun-bao/
	findAllReachableTargets := func(g [][]byte, s pair, t byte) (ps []pair) {
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
				if v == valid && !vis[i][j] {
					cnt++
					f(i, j)
				}
			}
		}
		return
	}

	// 下列代码来自 LC1254 https://leetcode-cn.com/problems/number-of-closed-islands/
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
				if v == validCell && !vis[i][j] {
					comp = []pair{}
					if f(i, j) {
						comps = append(comps, comp)
						// do comp ...
					}
				}
			}
		}
		return
	}

	// 周赛 212D https://leetcode-cn.com/problems/rank-transform-of-a-matrix/
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
		getST,
		disAll, disST, findAllReachableTargets,
		cntCC, dfsValidGrids,
		findSameValueCC,
		isValidPoint, findOneTargetAnyWhere, findAllTargetsAnyWhere,
	}
}
