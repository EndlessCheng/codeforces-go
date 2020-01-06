package copypasta

import (
	"math"
	"math/bits"
	"sort"
	"strconv"
	"strings"
)

// 注意：若有超过两个数相加，要特别注意 inf 的选择

func commonCollection() {
	// HELPER
	const mod int64 = 1e9 + 7 // 998244353
	const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	pow2 := [...]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144}
	pow10 := [...]int{1, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9}
	dirOffset4 := [...][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	dirOffset4R := [...][2]int{{1, 1}, {-1, 1}, {-1, -1}, {1, -1}}
	dirOffset8 := [...][2]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
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
	ifElseI := func(cond bool, r1, r2 int) int {
		if cond {
			return r1
		}
		return r2
	}
	ifElseS := func(cond bool, r1, r2 string) string {
		if cond {
			return r1
		}
		return r2
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

	quickPow := func(x int64, n int, mod int64) int64 {
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

	sort3 := func(a ...int) (int, int, int) {
		sort.Ints(a)
		return a[0], a[1], a[2]
	}

	reverseArr := func(s []byte) {
		for i, j := 0, len(s)-1; i < j; {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}
	reverseStr := func(s string) []byte {
		n := len(s)
		res := make([]byte, n)
		for i := range s {
			res[i] = s[n-1-i]
		}
		return res
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

	same := func(a ...int) bool {
		for _, v := range a[1:] {
			if v != a[0] {
				return false
			}
		}
		return true
	}

	// floatStr must contain a .
	// all decimal part must have same length
	// floatToInt("3.000100", 1e6) => 3000100
	// "3.0001" is not allowed
	floatToInt := func(floatStr string, shift10 int) int {
		splits := strings.SplitN(floatStr, ".", 2)
		i, _ := strconv.Atoi(splits[0])
		d, _ := strconv.Atoi(splits[1])
		return i*shift10 + d
	}

	// floatToRat("1.2", 1e1) => (6, 5)
	floatToRat := func(floatStr string, shift10 int) (m, n int) {
		m = floatToInt(floatStr, shift10)
		n = shift10
		var gcd int // calcGCD(m, n)
		m /= gcd
		n /= gcd
		return
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

	// TODO: 单调栈/单调队列
	// 推荐 https://cp-algorithms.com/data_structures/stack_queue_modification.html
	// https://oi-wiki.org/ds/monotonous-stack/
	// https://oi-wiki.org/ds/monotonous-queue/
	// TODO: CF1237D

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

	_ = []interface{}{
		pow2, pow10, dirOffset4, dirOffset4R, dirOffset8, orderP3,
		min, mins, max, maxs, ifElseI, ifElseS,
		abs, absAll, quickPow, toAnyBase, initSum2D, querySum2D,
		copyMat, hash01Mat, sort3, reverseArr, reverseStr, merge, unique, discrete, same,
		floatToRat, complement, containsAll, maxSubArraySum, maxSubArrayAbsSum,
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
	stQuery := func(l, r int) int { // [l,r] 注意 l r 是从 0 开始算的
		k := uint(bits.Len(uint(r-l+1)) - 1)
		return max(st[l][k], st[r-(1<<k)+1][k])
	}

	// 下标版本，查询返回的是区间最值的下标
	//type pair struct{ v, i int }
	//var st [][20]pair
	//stInit := func(a []int) {
	//	n := len(a)
	//	st = make([][20]pair, n)
	//	for i := range st {
	//		st[i][0] = pair{a[i], i}
	//	}
	//	for j := uint(1); 1<<j <= n; j++ {
	//		for i := 0; i+(1<<j)-1 < n; i++ {
	//			st0, st1 := st[i][j-1], st[i+(1<<(j-1))][j-1]
	//			if st0.v < st1.v { // 最小值
	//				st[i][j] = st0
	//			} else {
	//				st[i][j] = st1
	//			}
	//		}
	//	}
	//}
	//stQuery := func(l, r int) int { // [l,r] 注意 l r 是从 0 开始算的
	//	k := uint(bits.Len(uint(r-l+1)) - 1)
	//	st0, st1 := st[l][k], st[r-(1<<k)+1][k]
	//	if st0.v < st1.v { // 最小值
	//		return st0.i
	//	}
	//	return st1.i
	//}

	// Sqrt Decomposition
	// https://oi-wiki.org/ds/decompose/
	// https://oi-wiki.org/ds/block-array/
	// 题目推荐 https://cp-algorithms.com/data_structures/sqrt_decomposition.html#toc-tgt-8
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
	// 题目推荐 https://cp-algorithms.com/data_structures/sqrt_decomposition.html#toc-tgt-8
	mo := func(n, q int, a []int) {
		type query struct {
			blockIdx  int
			l, r, idx int
		}
		qs := make([]query, q)
		blockSize := int(math.Round(math.Sqrt(float64(n))))
		for i := range qs {
			var l, r int
			//Fscan(in, &l, &r)
			qs[i] = query{l / blockSize, l, r + 1, i}
		}
		sort.Slice(qs, func(i, j int) bool {
			qi, qj := qs[i], qs[j]
			return qi.blockIdx < qj.blockIdx || qi.blockIdx == qj.blockIdx && qi.r < qj.r
		})

		// 从 1 开始算，方便 debug
		l, r := 1, 1
		update := func(idx, delta int) {
			// 有些题目在 delta 为 1 和 -1 时逻辑的顺序是对称性的
			//v := a[idx]
			if delta == 1 {

			} else {

			}
		}
		ans := make([]int, q)
		getAns := func(q query) int {
			// 提醒：q.r 是加一后的，计算时需要注意

			return 0
		}
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
			_ = v
			//Fprintln(out, v)
		}
	}

	_ = []interface{}{
		stInit, stQuery,
		sqrtInit, sqrtOp,
		mo,
	}
}

//（含组合排列）
func loopCollection() {
	dirOffset4 := [...][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
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

	loopSet := func(arr []int) (ans int) {
		n := uint(len(arr))
		//outer:
		for i := 0; i < 1<<n; i++ {
			sum := 0
			for j := uint(0); j < n; j++ {
				if i>>j&1 == 1 { // choose j in range [0,n)
					// sum+=do(arr[j]) or continue outer
				}
			}
			ans = max(ans, sum)
		}
		return
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
			for _, dir := range dirOffset4 {
				if !dfs(i+dir[0], j+dir[1]) {
					// 遍历完该连通分量再 return
					res = false
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
	searchDirOffset4 := func(maxI, maxJ, centerI, centerJ, dis int) {
		for i, dir := range dirOffset4 {
			dir2 := dirOffset4[(i+1)%4]
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
	searchDirOffset4R := func(maxI, maxJ, centerI, centerJ, dis int) {
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
		loopSet, dfsGrids, searchDirOffset4, searchDirOffset4R,
		combinations, permutations, permuteAll,
	}
}

//func grayCode(length int) []int {
//	if length == 1 {
//		return []int{0, 1}
//	}
//	part0 := grayCode(length - 1)
//	part1 := make([]int, len(part0))
//	for i, v := range part0 {
//		part1[len(part0)-i-1] = v
//	}
//	for i, v := range part1 {
//		part1[i] = v | 1<<uint(length-1)
//	}
//	return append(part0, part1...)
//}
func grayCode(length int) []int {
	ans := make([]int, 1<<uint(length))
	for i := range ans {
		ans[i] = i ^ i>>1
	}
	return ans
}

// https://oeis.org/A001227
func consecutiveNumbersSum(n int) (ans int) {
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i&1 == 1 {
				ans++
			}
			if i*i < n && n/i&1 == 1 {
				ans++
			}
		}
	}
	return
}

// https://oeis.org/A000127
// Maximal number of regions obtained by joining n points around a circle by straight lines.
// Also number of regions in 4-space formed by n-1 hyperplanes.
//
//     n*(n-1)*(n*n-5*n+18)/24+1

// https://leetcode-cn.com/contest/weekly-contest-139/problems/adding-two-negabinary-numbers/
func addNegabinary(a1, a2 []int) []int {
	if len(a1) < len(a2) {
		a1, a2 = a2, a1
	}
	for i, j := len(a1)-1, len(a2)-1; j >= 0; {
		a1[i] += a2[j]
		i--
		j--
	}
	ans := append(make([]int, 2), a1...)
	for i := len(ans) - 1; i >= 0; i-- {
		if ans[i] >= 2 {
			ans[i] -= 2
			if ans[i-1] >= 1 {
				ans[i-1]--
			} else {
				ans[i-1]++
				ans[i-2]++
			}
		}
	}
	for i, v := range ans {
		if v != 0 {
			return ans[i:]
		}
	}
	return []int{0}
}

// https://leetcode.com/problems/convert-to-base-2/
func toNegabinary(n int) (res string) {
	if n == 0 {
		return "0"
	}
	for ; n != 0; n = -(n >> 1) {
		res = string('0'+n&1) + res
	}
	return
}
