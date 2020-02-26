package copypasta

import (
	. "fmt"
	"io"
	"math"
	"math/bits"
	"sort"
)

// 从数据范围找思路：
// 1e9~1e18 √n logn 1     二分 二进制
// 1e5~1e6  nlogn nαn n   二分 RMQ 并查集
// 1e3~1e4  n^2 n√n       RMQ DP 分块
// 300~500  n^3           DP 二分图

// 注意：若不止两个数相加，要特别注意 inf 的选择
// 一个 Golang 的注意事项：for-range array 时，遍历 i 时修改 i 后面的元素的值是不影响 ai 的，只能用 a[i] 获取

func commonCollection() {
	// HELPER
	const mod int64 = 1e9 + 7 // 998244353
	const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	pow2 := [...]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144}
	pow10 := [...]int{1, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9}
	factorial := [...]int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800 /*10!*/, 39916800, 479001600}
	// 注意：对于格点来说，dir4 对应的方向是右下左上（东南西北）
	dir4 := [...][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dir4R := [...][2]int{{1, 1}, {-1, 1}, {-1, -1}, {1, -1}}
	dir8 := [...][2]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	orderP3 := [6][3]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}

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
	mins := func(vals ...int) int {
		ans := vals[0]
		for _, val := range vals[1:] {
			if val < ans {
				ans = val
			}
		}
		return ans
	}
	maxs := func(vals ...int) int {
		ans := vals[0]
		for _, val := range vals[1:] {
			if val > ans {
				ans = val
			}
		}
		return ans
	}
	sort3 := func(a ...int) (x, y, z int) { sort.Ints(a); return a[0], a[1], a[2] }
	ternaryI := func(cond bool, r1, r2 int) int {
		if cond {
			return r1
		}
		return r2
	}
	ternaryS := func(cond bool, r1, r2 string) string {
		if cond {
			return r1
		}
		return r2
	}
	toInts := func(s []byte) []int {
		ints := make([]int, len(s))
		for i, b := range s {
			ints[i] = int(b)
		}
		return ints
	}
	xor := func(b1, b2 bool) bool { return b1 && !b2 || !b1 && b2 }
	zip := func(a, b []int) {
		n := len(a)
		type pair struct{ x, y int }
		ps := make([]pair, n)
		for i := range ps {
			ps[i] = pair{a[i], b[i]}
		}
	}
	zipI := func(a []int) {
		n := len(a)
		type pair struct{ x, y int }
		ps := make([]pair, n)
		for i := range ps {
			ps[i] = pair{a[i], i}
		}
	}
	// END HELPER

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	absAll := func(a []int) {
		for i, v := range a {
			if v < 0 {
				a[i] = -v
			}
		}
	}

	// https://en.wikipedia.org/wiki/Exponentiation_by_squaring
	pow := func(x int64, n int, mod int64) int64 {
		x %= mod
		res := int64(1) % mod
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	calcFactorial := func(n int) int64 {
		ans := int64(1)
		for i := 2; i <= n; i++ {
			ans *= int64(i)
		}
		return ans
	}

	// 从低位到高位
	toAnyBase := func(x, base int) (res []int) {
		for ; x > 0; x /= base {
			res = append(res, x%base)
		}
		return
	}

	var sum2d [][]int
	initSum2D := func(mat [][]int) {
		n, m := len(mat), len(mat[0])
		sum2d = make([][]int, n+1)
		sum2d[0] = make([]int, m+1)
		for i, row := range mat {
			sum2d[i+1] = make([]int, m+1)
			for j, mij := range row {
				sum2d[i+1][j+1] = sum2d[i+1][j] + sum2d[i][j+1] - sum2d[i][j] + mij
			}
		}
	}
	// r1<=r<=r2 && c1<=c<=c2
	querySum2D := func(r1, c1, r2, c2 int) int {
		r2++
		c2++
		return sum2d[r2][c2] - sum2d[r2][c1] - sum2d[r1][c2] + sum2d[r1][c1]
	}

	//

	copyMat := func(mat [][]int) [][]int {
		n, m := len(mat), len(mat[0])
		dst := make([][]int, n)
		for i, row := range mat {
			dst[i] = make([]int, m)
			copy(dst[i], row)
		}
		return dst
	}

	hash01Mat := func(mat [][]int) int {
		hash := 0
		cnt := uint(0)
		for _, row := range mat {
			for _, mij := range row {
				hash |= mij << cnt
				cnt++
			}
		}
		return hash
	}

	reverse := func(a []byte) []byte {
		n := len(a)
		r := make([]byte, n)
		for i, v := range a {
			r[n-1-i] = v
		}
		return r
	}
	reverseSelf := func(s []byte) {
		for i, j := 0, len(s)-1; i < j; {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}

	equals := func(a, b []int) bool {
		// assert len(a) == len(b)
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}

	// a b 必须是有序的（可以为空）
	merge := func(a, b []int) []int {
		i, n := 0, len(a)
		j, m := 0, len(b)
		res := make([]int, 0, n+m)
		for {
			if i == n {
				return append(res, b[j:]...)
			}
			if j == m {
				return append(res, a[i:]...)
			}
			if a[i] < b[j] { // 改成 > 为降序
				res = append(res, a[i])
				i++
			} else {
				res = append(res, b[j])
				j++
			}
		}
	}

	// a 必须是有序的
	unique := func(a []int) (res []int) {
		n := len(a)
		if n == 0 {
			return
		}
		res = make([]int, 1, n)
		res[0] = a[0]
		for i := 1; i < n; i++ {
			if a[i] != a[i-1] {
				res = append(res, a[i])
			}
		}
		return
	}

	// 离散化 discrete([]int{100,20,50,50}, 1) => []int{3,1,2,2}
	// 相当于转换成第几小
	discrete := func(arr []int, start int) (kthArr []int) {
		n := len(arr)
		if n == 0 {
			return
		}

		type pair struct{ val, i int }
		ps := make([]pair, n)
		for i, v := range arr {
			ps[i] = pair{v, i}
		}
		sort.Slice(ps, func(i, j int) bool { return ps[i].val < ps[j].val })
		kthArr = make([]int, n)

		// 有重复
		kth := start
		kthArr[ps[0].i] = kth
		for i := 1; i < n; i++ {
			if ps[i].val != ps[i-1].val {
				kth++
			}
			kthArr[ps[i].i] = kth
		}

		// 无重复
		//for i, p := range ps {
		//	kthArr[p.i] = i + start
		//}
		return
	}

	// 哈希编号，也可以理解成另一种离散化
	// 编号从 0 开始
	indexMap := func(a []string) map[string]int {
		mp := map[string]int{}
		for _, v := range a {
			if _, ok := mp[v]; !ok {
				mp[v] = len(mp)
			}
		}
		return mp
	}

	allSame := func(a ...int) bool {
		for _, v := range a[1:] {
			if v != a[0] {
				return false
			}
		}
		return true
	}

	// a 相对于 [0,n) 的补集
	// a 必须是升序且无重复元素
	complement := func(n int, a []int) (res []int) {
		j := 0
		for i := 0; i < n; i++ {
			if j == len(a) || i < a[j] {
				res = append(res, i)
			} else {
				j++
			}
		}
		return
	}

	// 逆序数
	var mergeCount func([]int) int64
	mergeCount = func(a []int) int64 {
		n := len(a)
		if n <= 1 {
			return 0
		}
		b := make([]int, n/2)
		c := make([]int, n-n/2)
		copy(b, a[:n/2])
		copy(c, a[n/2:])
		cnt := mergeCount(b) + mergeCount(c)
		ai, bi, ci := 0, 0, 0
		for ai < n {
			// 归并排序的同时计算逆序数
			if bi < len(b) && (ci == len(c) || b[bi] <= c[ci]) {
				a[ai] = b[bi]
				bi++
			} else {
				cnt += int64(n/2 - bi)
				a[ai] = c[ci]
				ci++
			}
			ai++
		}
		return cnt
	}

	// 数组第 k 小
	nthElement := func(a []int, k int) int {
		// TODO
		return 0
	}

	// x 是否包含 y 中的所有元素，且顺序一致
	containsAll := func(x, y []int) bool {
		for len(y) < len(x) {
			if len(y) == 0 {
				return true
			}
			if x[0] == y[0] {
				y = y[1:]
			}
			x = x[1:]
		}
		return false
	}

	//

	// 算法导论 练习4.1-5
	maxSubArraySum := func(a []int) int {
		curSum, maxSum := a[0], a[0]
		for _, v := range a[1:] {
			curSum = max(curSum+v, v)
			maxSum = max(maxSum, curSum)
		}
		return maxSum
	}

	maxSubArrayAbsSum := func(a []int) int {
		//min, max, abs := math.Min, math.Max, math.Abs
		curMaxSum, maxSum := a[0], a[0]
		curMinSum, minSum := a[0], a[0]
		for _, v := range a[1:] {
			curMaxSum = max(curMaxSum+v, v)
			maxSum = max(maxSum, curMaxSum)
			curMinSum = min(curMinSum+v, v)
			minSum = min(minSum, curMinSum)
		}
		return max(abs(maxSum), abs(minSum))
	}

	// 扫描线
	// https://cses.fi/book/book.pdf 30.1
	// TODO 窗口的星星 https://www.luogu.com.cn/problem/P1502
	// 天际线问题 https://leetcode-cn.com/problems/the-skyline-problem/
	// TODO 矩形面积并 https://leetcode-cn.com/problems/rectangle-area-ii/ 《算法与实现》5.4.3
	// LC 套题 https://leetcode-cn.com/tag/line-sweep/
	sweepLine := func(in io.Reader, n int) {
		type event struct{ pos, delta int }
		events := make([]event, 0, 2*n)
		for i := 0; i < n; i++ {
			var l, r int
			Fscan(in, &l, &r)
			events = append(events, event{l, 1}, event{r, -1})
		}
		sort.Slice(events, func(i, j int) bool {
			a, b := events[i], events[j]
			return a.pos < b.pos || a.pos == b.pos && a.delta < b.delta // < 先出后进；> 先进后出
		})

		for _, e := range events {
			_ = e
		}
	}

	// 生成字符串 s 的长度至多为 k 的所有（非空）子串
	var _ss []string
	var _genSubStrs func(s, sub string)
	_genSubStrs = func(s, sub string) {
		_ss = append(_ss, sub)
		if len(sub) < 4 { // custom k
			for i := range s {
				_genSubStrs(s[i+1:], sub+string(s[i]))
			}
		}
	}
	genSubStrs := func(s string) []string {
		_ss = []string{}
		_genSubStrs(s, "")
		a := _ss[1:] // remove ""
		sort.Strings(a)
		n := len(a)
		res := make([]string, 1, n)
		res[0] = a[0]
		for i := 1; i < n; i++ {
			if a[i] != a[i-1] {
				res = append(res, a[i])
			}
		}
		return res
	}

	_ = []interface{}{
		pow2, pow10, dir4, dir4R, dir8, orderP3, factorial,
		min, mins, max, maxs, ternaryI, ternaryS, toInts, xor, zip, zipI,
		abs, absAll, pow, calcFactorial, toAnyBase, initSum2D, querySum2D,
		copyMat, hash01Mat, sort3, reverse, reverseSelf, equals, merge, unique, discrete, indexMap, allSame, complement, nthElement, containsAll,
		maxSubArraySum, maxSubArrayAbsSum, sweepLine, genSubStrs,
	}
}

// https://cp-algorithms.com/sequences/rmq.html
func rmqCollection() {
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

	// Sparse Table
	// https://oi-wiki.org/ds/sparse-table/
	// 模板中的核心函数 max 可以换成其他具有区间合并性质的函数，如 gcd 等
	// 模板题 https://www.luogu.com.cn/problem/P3865
	// 题目推荐 https://cp-algorithms.com/data_structures/sparse-table.html#toc-tgt-5
	const mx = 17 // 17 for 1e5, 20 for 1e6
	var st [][mx]int
	stInit := func(a []int) {
		n := len(a)
		st = make([][mx]int, n)
		for i := range st {
			st[i][0] = a[i]
		}
		for j := uint(1); 1<<j <= n; j++ {
			for i := 0; i+(1<<j)-1 < n; i++ {
				st[i][j] = max(st[i][j-1], st[i+(1<<(j-1))][j-1])
			}
		}
	}
	stQuery := func(l, r int) int { // [l,r) 注意 l r 是从 0 开始算的
		k := uint(bits.Len(uint(r-l)) - 1)
		return max(st[l][k], st[r-(1<<k)][k])
	}

	// ST 下标版本，查询返回的是区间最值的下标
	//type pair struct{ v, i int }
	//const mx = 17
	//var st [][mx]pair
	//stInit := func(a []int) {
	//	n := len(a)
	//	st = make([][mx]pair, n)
	//	for i := range st {
	//		st[i][0] = pair{a[i], i}
	//	}
	//	for j := uint(1); 1<<j <= n; j++ {
	//		for i := 0; i+(1<<j)-1 < n; i++ {
	//			st0, st1 := st[i][j-1], st[i+(1<<(j-1))][j-1]
	//			if st0.v <= st1.v { // 最小值，相等时下标取左侧
	//				st[i][j] = st0
	//			} else {
	//				st[i][j] = st1
	//			}
	//		}
	//	}
	//}
	//stQuery := func(l, r int) int { // [l,r) 注意 l r 是从 0 开始算的
	//	k := uint(bits.Len(uint(r-l)) - 1)
	//	a, b := st[l][k], st[r-(1<<k)][k]
	//	if a.v <= b.v { // 最小值，相等时下标取左侧
	//		return a.i
	//	}
	//	return b.i
	//}

	// 分块 Sqrt Decomposition
	// https://oi-wiki.org/ds/decompose/
	// https://oi-wiki.org/ds/block-array/
	// 题目推荐 https://cp-algorithms.com/data_structures/sqrt_decomposition.html#toc-tgt-8
	// TODO: 台湾的《根號算法》https://www.csie.ntu.edu.tw/~sprout/algo2018/ppt_pdf/root_methods.pdf
	type block struct {
		l, r           int // [l,r]
		arr, sortedArr []int
		//lazyAdd int
	}
	var blocks []*block
	sqrtInit := func(a []int) {
		n := len(a)
		blockSize := int(math.Sqrt(float64(n)))
		//blockSize := int(math.Sqrt(float64(n) * math.Log2(float64(n+1))))
		blockNum := (n-1)/blockSize + 1
		blocks = make([]*block, blockNum)
		for i, ai := range a {
			j := i / blockSize
			if i%blockSize == 0 {
				blocks[j] = &block{l: i, arr: make([]int, 0, blockSize)}
			}
			b := blocks[j]
			b.arr = append(b.arr, ai)
		}
		for _, b := range blocks {
			b.r = b.l + len(b.arr) - 1
			b.sortedArr = make([]int, len(b.arr))
			copy(b.sortedArr, b.arr)
			sort.Ints(b.sortedArr)
		}
	}
	sqrtOp := func(l, r int) { // [l,r], starts at 0
		for _, b := range blocks {
			if b.r < l {
				continue
			}
			if b.l > r {
				break
			}
			if l <= b.l && b.r <= r {
				// do op on full block
			} else {
				// do op on part block
				bl := max(b.l, l)
				br := min(b.r, r)
				for i := bl - b.l; i <= br-b.l; i++ {
					//b.arr[i]
				}
			}
		}
	}

	// 莫队算法
	// 分块，将左端点分配在一个较小的范围，并且按照右端点从小到大排序，
	// 这样对于每一块，指针移动的次数为 O(√n*√n+n) = O(n)
	// 此外，记录的是 [l,r)，这样能简化处理查询结果的代码
	// https://oi-wiki.org/misc/mo-algo/
	// 模板题 https://www.luogu.com.cn/problem/P1494
	// 题目推荐 https://cp-algorithms.com/data_structures/sqrt_decomposition.html#toc-tgt-8
	mo := func(in io.Reader, out io.Writer, n, q int, a []int) {
		type query struct {
			blockIdx  int
			l, r, idx int
		}
		qs := make([]query, q)
		blockSize := int(math.Round(math.Sqrt(float64(n))))
		for i := range qs {
			var l, r int
			Fscan(in, &l, &r)
			qs[i] = query{l / blockSize, l, r + 1, i}
		}
		sort.Slice(qs, func(i, j int) bool {
			qi, qj := qs[i], qs[j]
			if qi.blockIdx != qj.blockIdx {
				return qi.blockIdx < qj.blockIdx
			}
			// 奇偶化排序
			if qi.blockIdx&1 == 0 {
				return qi.r < qj.r
			}
			return qi.r > qj.r
		})

		// 从 1 开始算，方便 debug
		cnt := 0
		l, r := 1, 1
		update := func(idx, delta int) {
			// 有些题目在 delta 为 1 和 -1 时逻辑的顺序是严格对称的
			//v := a[idx-1]
			if delta == 1 {
				cnt++
			} else {
				cnt--
			}
		}
		getAns := func(q query) int {
			// 提醒：q.r 是加一后的，计算时需要注意
			//l := q.r - q.l
			return cnt
		}
		ans := make([]int, q)
		for _, q := range qs {
			// prepare
			// 有些题目需要维护差分值，因为 [l,r] 的差分是 s(r)-s(l-1)，此时 update 传入的应为 l-1
			for ; r < q.r; r++ {
				update(r, 1)
			}
			for ; l < q.l; l++ {
				update(l, -1)
			}
			for l > q.l {
				l--
				update(l, 1)
			}
			for r > q.r {
				r--
				update(r, -1)
			}
			ans[q.idx] = getAns(q)
		}
		for _, v := range ans {
			Fprintln(out, v)
		}
	}

	// TODO: 带修改的莫队

	// TODO: 树上莫队

	_ = []interface{}{
		stInit, stQuery,
		sqrtInit, sqrtOp,
		mo,
	}
}

func monotoneCollection() {
	// 推荐 https://cp-algorithms.com/data_structures/stack_queue_modification.html
	// TODO: CF1237D

	// 单调栈
	// 举例：返回每个元素两侧严格大于它的元素位置（不存在则为 -1 或 n）
	// 如何理解：把数组想象成一列山峰，站在山峰 a[i] 的顶上向两侧的上方看，是看不到高山背后的矮山的，只能看到一座座更高的山峰。
	// 这就启发我们引入一个顶小底大的单调栈，入栈前不断比较栈顶元素直到找到一个比当前元素大的。
	// 栈可以存元素值和下标，也可以只存下标但这样写就要判断栈是否为空
	// 技巧：事先压入一个边界元素到栈底，这样保证循环中栈一定不会为空，从而简化逻辑
	// https://oi-wiki.org/ds/monotonous-stack/
	// 模板题 https://www.luogu.com.cn/problem/P5788
	// 与 DP 结合 https://codeforces.com/problemset/problem/1313/C2
	monotoneStack := func(a []int) ([]int, []int) {
		n := len(a)
		const border int = 2e9 // -2e9
		type pair struct{ v, i int }
		posL := make([]int, n)
		stack := []pair{{border, -1}}
		for i, v := range a {
			for {
				if top := stack[len(stack)-1]; top.v > v { // 严格大于
					posL[i] = top.i
					break
				}
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, pair{v, i})
		}
		posR := make([]int, n)
		stack = []pair{{border, n}}
		for i := n - 1; i >= 0; i-- {
			v := a[i]
			for {
				if top := stack[len(stack)-1]; top.v > v { // 严格大于
					posR[i] = top.i
					break
				}
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, pair{v, i})
		}

		return posL, posR
	}

	// 单调队列
	// https://oi-wiki.org/ds/monotonous-queue/
	// 模板题（滑动窗口） https://www.luogu.com.cn/problem/P1886
	// 例题：CF1237D
	monotoneQueue := func(a []int, fixedSize int) ([]int, []int) {
		// 为简单起见，用数组+下标模拟
		// 举例：固定大小的区间最值（滑动窗口）
		n := len(a)
		mins := make([]int, n)
		deque := make([]int, n)
		s, t := 0, 0
		for i, v := range a {
			for ; s < t && a[deque[t-1]] >= v; t-- { // >= 意味着相等的元素取靠右的，若改成 > 表示相等的元素取靠左的
			}
			deque[t] = i
			t++
			if i+1 >= fixedSize {
				mins[i+1-fixedSize] = a[deque[s]]
				if deque[s] == i+1-fixedSize { // popL 的条件随题目的不同而变化
					s++
				}
			}
		}
		maxs := make([]int, n)
		deque = make([]int, n)
		s, t = 0, 0
		for i, v := range a {
			for ; s < t && a[deque[t-1]] <= v; t-- {
			}
			deque[t] = i
			t++
			if i+1 >= fixedSize {
				maxs[i+1-fixedSize] = a[deque[s]]
				if deque[s] == i+1-fixedSize {
					s++
				}
			}
		}
		return mins, maxs
	}

	_ = []interface{}{monotoneStack, monotoneQueue}
}

// NOTE: 对于搜索格子的题，可以不用创建 vis 而是通过修改格子的值为范围外的值（如零、负数、'#' 等）来做到这一点
func loopCollection() {
	dir4 := [...][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
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

	// 枚举 {0,1,...,n-1} 的全部子集
	loopSet := func(arr []int) {
		n := uint(len(arr))
		//outer:
		for sub := 0; sub < 1<<n; sub++ { // sub repr a subset which elements are in range [0,n)
			// do(sub)
			for i := uint(0); i < n; i++ {
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

	// 枚举 {0,1,...,n-1} 的大小为 k 的子集（按字典序）
	// 参考《挑战程序设计竞赛》P.156-158
	loopSubsetK := func(arr []int, k int) {
		n := uint(len(arr))
		for sub := 1<<uint(k) - 1; sub < 1<<n; {
			// do(sub)
			x := sub & -sub
			y := sub + x
			sub = sub&^y/x>>1 | y
		}
	}

	// 枚举排列 + 剪枝
	// 即有 n 个位置，枚举每个位置上可能出现的值（范围在 sets 中），且每个位置上的元素不能重复
	// 例题见 LC169D
	loopPerm := func(n int, sets []int) bool {
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

	dfsGrids := func(grid [][]int) (comps int) {
		// grid[i][j] = 0 or 1
		n, m := len(grid), len(grid[0])
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		var dfs func(i, j int) bool
		dfs = func(i, j int) bool {
			if i < 0 || i >= n || j < 0 || j >= m {
				return false
			}
			if grid[i][j] == 1 {
				return true
			}
			if vis[i][j] {
				return true
			}
			vis[i][j] = true
			res := true
			for _, dir := range dir4 {
				if !dfs(i+dir[0], j+dir[1]) {
					// 遍历完该连通分量再 return
					//if {
					//	res = false
					//}
				}
			}
			return res
		}
		for i, gi := range grid {
			for j, gij := range gi {
				if gij == 0 && !vis[i][j] {
					if dfs(i, j) {
						comps++
					}
				}
			}
		}
		return
	}

	/*
		遍历以 (centerI, centerJ) 为中心的欧几里得距离为 dis 范围内的格点
		例如 dis=2 时：
		  #
		 # #
		# @ #
		 # #
		  #
	*/
	searchDir4 := func(maxI, maxJ, centerI, centerJ, dis int) {
		for i, dir := range dir4 {
			dir2 := dir4[(i+1)%4]
			dx := dir2[0] - dir[0]
			dy := dir2[1] - dir[1]
			x := centerI + dir[0]*dis
			y := centerJ + dir[1]*dis
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
		#####
		#   #
		# @ #
		#   #
		#####
	*/
	searchDir4R := func(maxI, maxJ, centerI, centerJ, dis int) {
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

	// 从数组 a 中选择 [0,min(r,len(a))] 个元素，生成所有组合
	// vals 初始化为 make([]int, 0, min(r,len(a)))
	var rangeCombinations func(a []int, r int, vals []int, do func(vals []int))
	rangeCombinations = func(a []int, r int, vals []int, do func([]int)) {
		do(vals)
		if len(vals) < r {
			for i, v := range a {
				rangeCombinations(a[i+1:], r, append(vals, v), do)
			}
		}
	}

	// 从一个长度为 n 的数组中选择 r 个元素，生成所有组合，每个组合用下标表示
	// r must <= n
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

	// 从一个长度为 n 的数组中选择 r 个元素，生成所有排列，每个排列用下标表示
	// r must <= n
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

	// 注意这个不是按照字典序 do 的
	// Permute the values at index i to len(arr)-1.
	// See 910C for example.
	var permute func([]int, int, func([]int))
	permute = func(arr []int, i int, do func([]int)) {
		if i == len(arr) {
			do(arr)
			return
		}
		permute(arr, i+1, do)
		for j := i + 1; j < len(arr); j++ {
			arr[i], arr[j] = arr[j], arr[i]
			permute(arr, i+1, do)
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	permuteAll := func(arr []int, do func([]int)) { permute(arr, 0, do) }

	_ = []interface{}{
		loopDiagonal, loopDiagonal2, loopSet, loopSubset, loopSubsetK, loopPerm, dfsGrids, searchDir4, searchDir4R,
		rangeCombinations, combinations, permutations, permuteAll,
	}
}
