package copypasta

import (
	. "fmt"
	"io"
	"math"
	"math/bits"
	"math/rand"
	"sort"
)

// General ideas https://codeforces.com/blog/entry/48417
// 从特殊到一般：尝试修改条件或缩小题目的数据范围，先研究某个特殊情况下的思路，然后再逐渐扩大数据范围来思考怎么改进算法

// 异类双变量：固定某变量统计另一变量的 [0,n)
//     EXTRA: 值域上的双变量，见 https://codeforces.com/contest/486/problem/D
// 同类双变量①：固定 i 统计 [0,n)
// 同类双变量②：固定 i 统计 [0,i-1]
// 套路：预处理数据（按照某种顺序排序/优先队列/BST/...），或者边遍历边维护，
//      然后固定变量 i，用均摊 O(1)~O(logn) 的复杂度统计范围内的另一变量 j
// 这样可以将复杂度从 O(n^2) 降低到 O(n) 或 O(nlogn)

/* 横看成岭侧成峰
考虑每个点产生的贡献 https://codeforces.com/problemset/problem/1009/E
考虑每条边产生的负贡献 https://atcoder.jp/contests/abc173/tasks/abc173_f
和式的另一视角。若每一项的值都在一个范围，不妨考虑另一个问题：值为 x 的项有多少个？https://atcoder.jp/contests/abc162/tasks/abc162_e
对所有排列考察所有子区间的性质，可以转换成对所有子区间考察所有排列。将子区间内部的排列和区间外部的排列进行区分，内部的性质单独研究，外部的当作 (n-(r-l))! 个排列 https://codeforces.com/problemset/problem/1284/C
转换为距离的众数 https://codeforces.com/problemset/problem/1365/C
转换为差分数组的变化 https://codeforces.com/problemset/problem/1110/E
不解释，自己感受 https://leetcode-cn.com/contest/biweekly-contest-31/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/
*/

// NOTE: 正难则反。 all => any, any => all https://codeforces.com/problemset/problem/621/C
// NOTE: 子区间和为 0 => 出现了两个同样的前缀和。这种题目建议下标从 1 开始，见 https://codeforces.com/problemset/problem/1333/C

// 尺取法套题 https://blog.csdn.net/weixin_43914593/article/details/104090474 算法竞赛专题解析（2）：尺取法（双指针）

// 栈+懒删除 https://codeforces.com/problemset/problem/1000/F

// NOTE: 若不止两个数相加，要特别注意 inf 的选择

// Golang 注意事项：
// 		和 slice 不同，for range array 时，遍历 i 时修改 i 后面的元素的值是不影响 ai 的，只能用 for+a[i] 或 forr a[:] 获取（因为 for range array 会整个拷贝一份）
// 		for-switch 内的 break 跳出的是该 switch，不是其外部的 for 循环
// 		对于存在海量小对象的情况（如 trie, treap 等），使用 debug.SetGCPercent(-1) 来禁用 GC，不去扫描大量对象，能明显减少耗时；
//		对于可以回收的情况（如 append 在超过 cap 时），使用 debug.SetGCPercent(-1) 虽然会减少些许耗时，但若有大量内存没被回收，会有 MLE 的风险；
//		其他情况下使用 debug.SetGCPercent(-1) 对耗时和内存使用无明显影响
//		对于多组数据的情况，禁用 GC 若 MLE，可在每组数据的开头或者末尾调用 debug.FreeOSMemory() 手动 GC
//		参考 https://zhuanlan.zhihu.com/p/77943973 https://draveness.me/golang/docs/part3-runtime/ch07-memory/golang-garbage-collector/
func commonCollection() {
	const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	// TIPS: dir4[i] 和 dir4[i^1] 互为相反方向
	type pair struct{ x, y int }
	dir4 := [...]pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右
	dir4C := [...]pair{ // 西东南北
		'W': {-1, 0},
		'E': {1, 0},
		'S': {0, -1},
		'N': {0, 1},
	}
	dir4c := [...]pair{ // 左右下上
		'L': {-1, 0},
		'R': {1, 0},
		'D': {0, -1},
		'U': {0, 1},
	}
	dir4R := [...]pair{{1, 1}, {-1, 1}, {-1, -1}, {1, -1}}
	dir8 := [...]pair{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	orderP3 := [6][3]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}
	pow10 := func(x int) int64 { return int64(math.Pow10(x)) } // 不需要 round

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
	mins := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if v < res {
				res = v
			}
		}
		return res
	}
	maxs := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if v > res {
				res = v
			}
		}
		return res
	}
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

	isDigit := func(b byte) bool { return '0' <= b && b <= '9' }
	isLower := func(b byte) bool { return 'a' <= b && b <= 'z' }
	isUpper := func(b byte) bool { return 'A' <= b && b <= 'Z' }
	isAlpha := func(b byte) bool { return 'A' <= b && b <= 'Z' || 'a' <= b && b <= 'z' }

	sort3 := func(a ...int) (x, y, z int) { sort.Ints(a); return a[0], a[1], a[2] }
	minString := func(a, b string) string {
		if len(a) != len(b) {
			if len(a) < len(b) {
				return a
			}
			return b
		}
		if a < b {
			return a
		}
		return b
	}
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
	getCol := func(mat [][]int, j int) (col []int) {
		for _, row := range mat {
			col = append(col, row[j])
		}
		return
	}
	copyMat := func(mat [][]int) [][]int {
		n, m := len(mat), len(mat[0])
		dst := make([][]int, n)
		for i, row := range mat {
			dst[i] = make([]int, m)
			copy(dst[i], row)
		}
		return dst
	}
	toInts := func(s []byte) []int {
		ints := make([]int, len(s))
		for i, b := range s {
			ints[i] = int(b)
		}
		return ints
	}

	// 适用于 a*b 超过 64 位范围的情况
	mul := func(a, b, mod int64) (res int64) {
		for ; b > 0; b >>= 1 {
			if b&1 == 1 {
				res = (res + a) % mod
			}
			a = (a + a) % mod
		}
		return
	}

	// https://en.wikipedia.org/wiki/Exponentiation_by_squaring
	pow := func(x, n, mod int64) int64 {
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
	digits := func(x int) (res []int) {
		for ; x > 0; x /= 10 {
			res = append(res, x%10)
		}
		return
	}

	// 带权(等差数列)前缀和
	{
		var n int // read
		a := make([]int64, n)
		sum := make([]int64, n+1)
		iSum := make([]int64, n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + v
			iSum[i+1] = iSum[i] + int64(i+1)*v
		}
		query := func(l, r int) int64 { return iSum[r] - iSum[l] - int64(l)*(sum[r]-sum[l]) } // [l,r)

		_ = query
	}

	// 二维前缀和
	var sum2d [][]int
	initSum2D := func(mat [][]int) {
		n, m := len(mat), len(mat[0])
		sum2d = make([][]int, n+1)
		sum2d[0] = make([]int, m+1)
		for i, row := range mat {
			sum2d[i+1] = make([]int, m+1)
			for j, v := range row {
				sum2d[i+1][j+1] = sum2d[i+1][j] + sum2d[i][j+1] - sum2d[i][j] + v
			}
		}
	}
	// r1<=r<=r2 && c1<=c<=c2
	querySum2D := func(r1, c1, r2, c2 int) int {
		r2++
		c2++
		return sum2d[r2][c2] - sum2d[r2][c1] - sum2d[r1][c2] + sum2d[r1][c1]
	}

	reverse := func(a []byte) []byte {
		n := len(a)
		b := make([]byte, n)
		for i, v := range a {
			b[n-1-i] = v
		}
		return b
	}
	reverseInPlace := func(a []byte) {
		for i, j := 0, len(a)-1; i < j; i++ {
			a[i], a[j] = a[j], a[i]
			j--
		}
	}

	equal := func(a, b []int) bool {
		// assert len(a) == len(b)
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}

	// 启发式合并：map 版
	mergeMap := func(a, b map[int]int) map[int]int {
		if len(a) < len(b) {
			a, b = b, a
		}
		for k, v := range b {
			a[k] += v
		}
		return a
	}

	// 合并有序数组，保留重复元素
	// a b 必须是有序的（可以为空）
	// 若不保留重复元素，则相当于求 a 和 b 的对称差（见下面 splitDifferenceAndIntersection 函数）
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

	// 求差集 A-B, B-A 和交集 A∩B
	// EXTRA: 求并集 union: A∪B = A-B+A∩B = merge(differenceA, intersection) 或 merge(differenceB, intersection)
	// EXTRA: 求对称差 symmetric_difference: A▲B = A-B ∪ B-A = merge(differenceA, differenceB)
	// a b 必须是有序的（可以为空）
	// 与图论结合 https://codeforces.com/problemset/problem/243/B
	splitDifferenceAndIntersection := func(a, b []int) (differenceA, differenceB, intersection []int) {
		i, n := 0, len(a)
		j, m := 0, len(b)
		for {
			if i == n {
				differenceB = append(differenceB, b[j:]...)
				return
			}
			if j == m {
				differenceA = append(differenceA, a[i:]...)
				return
			}
			x, y := a[i], b[j]
			if x < y { // 改成 > 为降序
				differenceA = append(differenceA, x)
				i++
			} else if x > y { // 改成 < 为降序
				differenceB = append(differenceB, y)
				j++
			} else {
				intersection = append(intersection, x)
				i++
				j++
			}
		}
	}

	// a 是否为 b 的子集（相当于 differenceA 为空）
	// a b 需要是有序的
	isSubset := func(a, b []int) bool {
		i, n := 0, len(a)
		j, m := 0, len(b)
		for {
			if i == n {
				return true
			}
			if j == m {
				return false
			}
			x, y := a[i], b[j]
			if x < y { // 改成 > 为降序
				return false
			} else if x > y { // 改成 < 为降序
				j++
			} else {
				i++
				j++
			}
		}
	}

	// EXTRA: a 是否为 b 的子序列
	// https://codeforces.com/problemset/problem/778/A
	isSubSequence := func(a, b []int) bool {
		i, n := 0, len(a)
		j, m := 0, len(b)
		for {
			if i == n {
				return true
			}
			if j == m {
				return false
			}
			if a[i] == b[j] {
				i++
				j++
			} else {
				j++
			}
		}
	}

	// 是否为不相交集合（相当于 intersection 为空）
	// a b 需要是有序的
	isDisjoint := func(a, b []int) bool {
		i, n := 0, len(a)
		j, m := 0, len(b)
		for {
			if i == n || j == m {
				return true
			}
			x, y := a[i], b[j]
			if x < y { // 改成 > 为降序
				i++
			} else if x > y { // 改成 < 为降序
				j++
			} else {
				return false
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
		//n = len(res)
		return
	}

	uniqueInPlace := func(a []int) []int {
		n := len(a)
		if n == 0 {
			return nil
		}
		j := 0
		for i := 1; i < n; i++ {
			if a[j] != a[i] {
				j++
				a[j] = a[i]
			}
		}
		//n = j + 1
		return a[:j+1]
	}

	// 离散化 discrete([]int{100,20,50,50}, 1) => []int{3,1,2,2}
	// 相当于转换成第几小
	// 若允许修改原数组，可以先将其排序去重后，再调用 discrete，注意去重后 n 需要重新赋值
	discrete := func(a []int, startIndex int) (kth []int) {
		type pair struct{ v, i int }
		n := len(a)
		if n == 0 {
			return
		}
		ps := make([]pair, n)
		for i, v := range a {
			ps[i] = pair{v, i}
		}
		sort.Slice(ps, func(i, j int) bool { return ps[i].v < ps[j].v }) // or SliceStable
		kth = make([]int, n)

		// a 有重复元素
		k := startIndex
		kth[ps[0].i] = k
		for i := 1; i < n; i++ {
			if ps[i].v != ps[i-1].v {
				k++
			}
			kth[ps[i].i] = k
		}

		// a 无重复元素
		for i, p := range ps {
			kth[p.i] = i + startIndex
		}

		return
	}

	// 离散化 discreteMap([]int{100,20,50,50}, 1) => map[int]int{100:3, 20:1, 50:2}
	// 若允许修改原数组，可以先将其排序去重后，再调用 discreteMap，注意去重后 n 需要重新赋值
	discreteMap := func(a []int, startIndex int) (kth map[int]int) {
		// assert len(a) > 0
		n := len(a)
		b := make([]int, n)
		copy(b, a)
		sort.Ints(b)

		// 有重复元素
		k := startIndex
		kth = map[int]int{b[0]: k}
		for i := 1; i < n; i++ {
			if b[i] != b[i-1] {
				k++
				kth[b[i]] = k
			}
		}

		// 无重复元素
		kth = make(map[int]int, n)
		for i, v := range b {
			kth[v] = i + startIndex
		}

		return
	}

	// 哈希编号，也可以理解成另一种离散化（无序）
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

	// 数组第 k 小 (Quick Select)       kthElement nthElement
	// 0 <= k < len(a)
	// 调用会改变数组中元素顺序
	// 代码实现参考算法第四版 p.221
	// 算法的平均比较次数为 ~2n+2kln(n/k)+2(n-k)ln(n/(n-k))
	// https://en.wikipedia.org/wiki/Quickselect
	// https://www.geeksforgeeks.org/quickselect-algorithm/
	// 模板题 LC215 https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
	// 模板题 https://codeforces.com/contest/977/problem/C
	quickSelect := func(a []int, k int) int {
		//k = len(a) - 1 - k // 求第 k 大
		rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
		for l, r := 0, len(a)-1; l < r; {
			v := a[l] // 切分元素
			i, j := l, r+1
			for {
				for i++; i < r && a[i] < v; i++ { // less(i, l)
				}
				for j--; j > l && a[j] > v; j-- { // less(l, j)
				}
				if i >= j {
					break
				}
				a[i], a[j] = a[j], a[i]
			}
			a[l], a[j] = a[j], v
			if j == k {
				break
			} else if j < k {
				l = j + 1
			} else {
				r = j - 1
			}
		}
		return a[k] //  a[:k+1]  a[k:]
	}

	contains := func(a []int, x int) bool {
		for _, v := range a {
			if v == x {
				return true
			}
		}
		return false
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

	// 扫描线
	// 某些题目需要配合线段树
	// https://cses.fi/book/book.pdf 30.1
	// TODO 窗口的星星 https://www.luogu.com.cn/problem/P1502
	// 天际线问题 LC218 https://leetcode-cn.com/problems/the-skyline-problem/
	// TODO 矩形面积并 LC850 https://leetcode-cn.com/problems/rectangle-area-ii/ 《算法与实现》5.4.3
	// 经典题 https://codeforces.com/problemset/problem/1000/C
	// LC 套题 https://leetcode-cn.com/tag/line-sweep/
	// todo CF652D
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

	// 扫描线：一维格点刷漆，返回被刷到的格点数
	countCoveredPoints := func(in io.Reader, m int) int {
		type pair struct{ p, d int }
		es := make([]pair, 0, 2*m)
		for i := 0; i < m; i++ {
			var l, r int
			Fscan(in, &l, &r)
			es = append(es, pair{l, 1}, pair{r, -1})
		}
		// assert len(es) > 0
		sort.Slice(es, func(i, j int) bool { return es[i].p < es[j].p })
		ans := es[len(es)-1].p - es[0].p + 1
		// 减去没被刷到的格点
		eventCnt, st := 0, es[0].p
		for _, e := range es {
			if eventCnt == 0 {
				if d := e.p - st - 1; d > 0 {
					ans -= d
				}
			}
			eventCnt += e.d
			if eventCnt == 0 {
				st = e.p
			}
		}
		return ans
	}

	// 二维离散化
	// 代码来源 https://atcoder.jp/contests/abc168/tasks/abc168_f
	discrete2D := func(n, m int) (ans int) {
		type line struct{ a, b, c int }
		lr := make([]line, n)
		du := make([]line, m)
		// read ...

		xs := []int{-2e9, 0, 2e9}
		ys := []int{-2e9, 0, 2e9}
		for _, l := range lr {
			a, b, c := l.a, l.b, l.c
			xs = append(xs, a, b)
			ys = append(ys, c)
		}
		for _, l := range du {
			a, b, c := l.a, l.b, l.c
			xs = append(xs, a)
			ys = append(ys, b, c)
		}
		sort.Ints(xs)
		xs = unique(xs)
		xi := discreteMap(xs, 0)
		sort.Ints(ys)
		ys = unique(ys)
		yi := discrete(ys, 0)

		lx, ly := len(xi), len(yi)
		glr := make([][]int, lx)
		gdu := make([][]int, lx)
		vis := make([][]bool, lx)
		for i := range glr {
			glr[i] = make([]int, ly)
			gdu[i] = make([]int, ly)
			vis[i] = make([]bool, ly)
		}
		for _, p := range lr {
			glr[xi[p.a]][yi[p.c]]++
			glr[xi[p.b]][yi[p.c]]--
		}
		for _, p := range du {
			gdu[xi[p.a]][yi[p.b]]++
			gdu[xi[p.a]][yi[p.c]]--
		}
		for i := 1; i < lx-1; i++ {
			for j := 1; j < ly-1; j++ {
				glr[i][j] += glr[i-1][j]
				gdu[i][j] += gdu[i][j-1]
			}
		}

		type pair struct{ x, y int }
		q := []pair{{xi[0], yi[0]}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			x, y := p.x, p.y
			if x == 0 || x == lx-1 || y == 0 || y == ly-1 {
				return -1
			} // 无穷大
			if !vis[x][y] {
				vis[x][y] = true
				ans += (xs[x+1] - xs[x]) * (ys[y+1] - ys[y])
				if glr[x][y] == 0 {
					q = append(q, pair{x, y - 1})
				}
				if glr[x][y+1] == 0 {
					q = append(q, pair{x, y + 1})
				}
				if gdu[x][y] == 0 {
					q = append(q, pair{x - 1, y})
				}
				if gdu[x+1][y] == 0 {
					q = append(q, pair{x + 1, y})
				}
			}
		}
		return
	}

	_ = []interface{}{
		pow10, dir4, dir4C, dir4c, dir4R, dir8, orderP3,
		min, mins, max, maxs, abs, absAll,
		isDigit, isLower, isUpper, isAlpha,
		ternaryI, ternaryS, toInts, xor, zip, zipI, getCol, minString,
		pow, mul, toAnyBase, digits, initSum2D, querySum2D, mergeMap,
		copyMat, sort3, reverse, reverseInPlace, equal,
		merge, splitDifferenceAndIntersection, isSubset, isSubSequence, isDisjoint,
		unique, uniqueInPlace, discrete, discreteMap, indexMap, allSame, complement, quickSelect, contains, containsAll,
		sweepLine, countCoveredPoints,
		discrete2D,
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
	// st[i][j] 对应的区间是 [i, i+1<<j)
	// https://oi-wiki.org/ds/sparse-table/
	// https://codeforces.com/blog/entry/66643
	// 模板中的核心函数 max 可以换成其他具有区间合并性质的函数（允许区间重叠），如 gcd 等
	// 模板题 https://www.luogu.com.cn/problem/P3865
	// 题目推荐 https://cp-algorithms.com/data_structures/sparse-table.html#toc-tgt-5
	const mx = 17 // 131072, 262144, 524288, 1048576
	var st [][mx]int
	stInit := func(a []int) {
		n := len(a)
		st = make([][mx]int, n)
		for i, v := range a {
			st[i][0] = v
		}
		for j := 1; 1<<j <= n; j++ {
			for i := 0; i+1<<j <= n; i++ {
				st[i][j] = max(st[i][j-1], st[i+1<<(j-1)][j-1])
			}
		}
	}
	// [l,r) 注意 l r 是从 0 开始算的
	stQuery := func(l, r int) int { k := bits.Len(uint(r-l)) - 1; return max(st[l][k], st[r-1<<k][k]) }

	// Sparse Table 下标版本，查询返回的是区间最值的下标
	{
		type pair struct{ v, i int }
		const mx = 17
		var st [][mx]pair
		stInit := func(a []int) {
			n := len(a)
			st = make([][mx]pair, n)
			for i, v := range a {
				st[i][0] = pair{v, i}
			}
			for j := 1; 1<<j <= n; j++ {
				for i := 0; i+1<<j <= n; i++ {
					if a, b := st[i][j-1], st[i+1<<(j-1)][j-1]; a.v <= b.v { // 最小值，相等时下标取左侧
						st[i][j] = a
					} else {
						st[i][j] = b
					}
				}
			}
		}
		stQuery := func(l, r int) int { // [l,r) 注意 l r 是从 0 开始算的
			k := bits.Len(uint(r-l)) - 1
			a, b := st[l][k], st[r-1<<k][k]
			if a.v <= b.v { // 最小值，相等时下标取左侧
				return a.i
			}
			return b.i
		}
		_, _ = stInit, stQuery
	}

	// 分块 Sqrt Decomposition
	// https://oi-wiki.org/ds/decompose/
	// https://oi-wiki.org/ds/block-array/
	// TODO: 台湾的《根號算法》https://www.csie.ntu.edu.tw/~sprout/algo2018/ppt_pdf/root_methods.pdf
	// 题目推荐 https://cp-algorithms.com/data_structures/sqrt_decomposition.html#toc-tgt-8
	// 好题 https://codeforces.com/problemset/problem/91/E
	// todo 动态逆序对 https://www.luogu.com.cn/problem/P3157 https://www.luogu.com.cn/problem/UVA11990
	type block struct {
		l, r           int // [l,r]
		origin, sorted []int
		//lazyAdd int
	}
	var blocks []block
	sqrtInit := func(a []int) {
		n := len(a)
		blockSize := int(math.Sqrt(float64(n)))
		//blockSize := int(math.Sqrt(float64(n) * math.Log2(float64(n+1))))
		blockNum := (n-1)/blockSize + 1
		blocks = make([]block, blockNum)
		for i, v := range a {
			j := i / blockSize
			if i%blockSize == 0 {
				blocks[j] = block{l: i, origin: make([]int, 0, blockSize)}
			}
			blocks[j].origin = append(blocks[j].origin, v)
		}
		for i := range blocks {
			b := &blocks[i]
			b.r = b.l + len(b.origin) - 1
			b.sorted = make([]int, len(b.origin))
			copy(b.sorted, b.origin)
			sort.Ints(b.sorted)
		}
	}
	sqrtOp := func(l, r int, v int) { // [l,r], starts at 0
		for i := range blocks {
			b := &blocks[i]
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
				for j := bl - b.l; j <= br-b.l; j++ {
					// do b.origin[j]...
				}
			}
		}
	}

	_ = []interface{}{
		stInit, stQuery,
		sqrtInit, sqrtOp,
	}
}

/* 平方根算法：组合两种算法从而降低复杂度 O(n^2) -> O(n√n)
参考 Competitive Programmer’s Handbook Ch.27

有 n 个对象，每个对象有一个「关于其他对象的统计量」ci（一个数、一个集合的元素个数，等等）
为方便起见，假设 ∑ci 的数量级和 n 一样，下面用 n 表示 ∑ci
当 ci > √n 时，这样的对象不超过 √n 个，暴力枚举这些对象之间的关系（或者，该对象与其他所有对象的关系），时间复杂度为 O(n) 或 O(n√n)。此乃算法一
当 ci ≤ √n 时，这样的对象有 O(n) 个，由于统计量 ci 很小，暴力枚举当前对象的统计量，时间复杂度为 O(n√n)。此乃算法二
这样，以 √n 为界，我们将所有对象划分成了两组，并用两个不同的算法处理
这两种算法是看待同一个问题的两种不同方式，通过恰当地组合这两个算法，复杂度由 O(n^2) 降至 O(n√n)
注意：**枚举时要做到不重不漏**

另一种题型是注意到 n 的整数分拆中，不同数字的个数至多有 O(√n) 种

好题 LCP16 https://leetcode-cn.com/problems/you-le-yuan-de-you-lan-ji-hua/
*/

// 莫队算法：对询问分块
// 分块，每一块的大小为 √n，这样可以将左端点分配在一个较小的范围，并且按照右端点从小到大排序，
// 从而对于每一块，指针移动的次数为 O(√n*√n+n) = O(n)，从而整体复杂度为 O(n√n) （注：这里假设询问次数等同于 n）
// 此外，记录的是 [l,r)，这样能简化处理查询结果的代码
// https://oi-wiki.org/misc/mo-algo/
// 模板题 https://www.luogu.com.cn/problem/P1494
// https://www.luogu.com.cn/problem/P4462
// 题目推荐 https://cp-algorithms.com/data_structures/sqrt_decomposition.html#toc-tgt-8
func moAlgorithm() {
	// 若 block 改变会对询问有影响，可以先放入不同的 block 然后再排序
	mo := func(in io.Reader, a []int, q int) []int {
		n := len(a)
		type query struct{ bid, l, r, qid int }
		qs := make([]query, q)
		blockSize := int(math.Round(math.Sqrt(float64(n))))
		for i := range qs {
			var l, r int
			Fscan(in, &l, &r) // 从 1 开始
			qs[i] = query{l / blockSize, l, r + 1, i}
		}
		sort.Slice(qs, func(i, j int) bool {
			qi, qj := qs[i], qs[j]
			if qi.bid != qj.bid {
				return qi.bid < qj.bid
			}
			// 奇偶化排序
			if qi.bid&1 == 0 {
				return qi.r < qj.r
			}
			return qi.r > qj.r
		})

		cnt := 0
		l, r := 1, 1 // 区间从 1 开始，方便 debug
		update := func(idx, delta int) {
			// NOTE: 有些题目在 delta 为 1 和 -1 时逻辑的顺序是严格对称的
			// v := a[idx-1]
			// ...
			if delta > 0 {
				cnt++
			} else {
				cnt--
			}
		}
		getAns := func(q query) int {
			// 提醒：q.r 是加一后的，计算时需要注意
			// sz := q.r - q.l
			// ...
			return cnt
		}
		ans := make([]int, q)
		for _, q := range qs {
			// prepare
			// NOTE: 有些题目需要维护差分值，由于 [l,r] 的差分是 s(r)-s(l-1)，此时 update 传入的应为 l-1
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
			ans[q.qid] = getAns(q)
		}
		return ans
	}

	// TODO: 带修改的莫队
	// https://www.luogu.com.cn/blog/deco/qian-tan-ji-chu-gen-hao-suan-fa-fen-kuai

	// TODO: 树上莫队
	// https://blog.csdn.net/weixin_43914593/article/details/108485396

	_ = mo
}

func monotoneCollection() {
	// 推荐 https://cp-algorithms.com/data_structures/stack_queue_modification.html

	// 单调栈
	// 举例：返回每个元素两侧严格大于它的元素位置（不存在则为 -1 或 n）
	// 如何理解：把数组想象成一列山峰，站在 a[i] 的山顶仰望两侧的山峰，是看不到高山背后的矮山的，只能看到一座座更高的山峰
	//          这就启发我们引入一个底大顶小的单调栈，入栈时不断比较栈顶元素直到找到一个比当前元素大的
	// 技巧：事先压入一个边界元素到栈底，这样保证循环时栈一定不会为空，从而简化逻辑
	// https://oi-wiki.org/ds/monotonous-stack/
	// 模板题 https://www.luogu.com.cn/problem/P5788
	//       https://www.luogu.com.cn/problem/P2866 http://poj.org/problem?id=3250
	//       https://leetcode-cn.com/problems/next-greater-element-i/ LC496/周赛18BA
	//       https://leetcode-cn.com/problems/next-greater-element-ii/ LC503/周赛18BB
	// 柱状图中最大的矩形 LC84 https://leetcode-cn.com/problems/largest-rectangle-in-histogram/ http://poj.org/problem?id=2559 http://poj.org/problem?id=2082
	// 后缀数组+不同矩形对应方案数之和 https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/D
	// 与 DP 结合
	//     https://codeforces.com/problemset/problem/1313/C2
	//     https://codeforces.com/problemset/problem/1407/D
	// 全 1 子矩阵个数 O(n^2) LC周赛196C https://leetcode-cn.com/contest/weekly-contest-196/problems/count-submatrices-with-all-ones/ 原题为 http://poj.org/problem?id=3494
	monotoneStack := func(a []int) ([]int, []int) {
		const border int = 2e9 // 求两侧小的话用 -1
		type pair struct{ v, i int }

		// 求左侧严格大于
		n := len(a)
		posL := make([]int, n)
		for i := range posL {
			posL[i] = -1
		}
		stack := []pair{{border, -1}}
		for i, v := range a {
			for {
				if top := stack[len(stack)-1]; top.v > v { //
					posL[i] = top.i
					break
				}
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, pair{v, i})
		}

		// 求右侧严格大于
		posR := make([]int, n)
		for i := range posR {
			posR[i] = n
		}
		stack = []pair{{border, n}}
		for i := n - 1; i >= 0; i-- {
			v := a[i]
			for {
				if top := stack[len(stack)-1]; top.v > v { //
					posR[i] = top.i
					break
				}
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, pair{v, i})
		}

		return posL, posR
	}

	/* 单调队列
	需要不断维护队列的单调性，即保证队列(指向的)数组元素从大到小或从小到大
		为简单起见，这里用数组+双下标模拟双端队列
		为保证有足够空间，队列初始大小应和指向的数组长度相同
		队列存储的是数组元素的下标
		l == r 表示队列为空
		l < r 表示队列不为空，队列指向数组元素为 a[idQ[l]], a[idQ[l]+1], ..., a[idQ[r-1]]
			注意：某些情况下这不等同于考察的区间就是 [idQ[l], idQ[r-1]]，但至少包含这一区间

	一般的写法是：   [pop]-push-query
		1. 初始化单调队列 idQ（初始大小为指向的数组长度），队首队尾下标 l r 指向 0
		2. 循环枚举右端点 i
			1. 循环枚举队头 idQ[l]，若区间长度 = i-idQ[l]+1 > 上界 M 则弹出队头，直至队列为空或不超出上界
				注：这里只是举了一个区间长度上界作为约束的例子，更复杂的约束见后面的代码
				注：若无约束这一步可忽略
			2. 准备插入右端点 i，为保证插入后的队列单调性，需要检查并弹出若干队尾元素
			3. 插入右端点 i
			4. 此时当前区间满足约束，可以查询区间最值等信息，此时队头就是右端点为 i 时的最优选择
			注意：若查询区间不包含右端点 i，或者说查询的区间右端点是 i-1，则上述步骤需要稍作改动，
				顺序是 1423    [pop]-query-push
				①的+1去掉（因为右端点是 i-1）
				④需要先检查队列是否为空再查询

	有些题目枚举左端点更为方便，细节见下面的 cf1237d

	https://oi-wiki.org/ds/monotonous-queue/

	todo http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=1070
	*/

	// 模板题 - 固定区间大小的区间最值（滑动窗口）   pop-push-query
	// https://www.luogu.com.cn/problem/P1886 http://poj.org/problem?id=2823
	// https://codeforces.com/problemset/status/940/problem/E
	fixedSizeMinMax := func(a []int, fixedSize int) (mins, maxs []int) {
		n := len(a)

		idQ := make([]int, n)
		l, r := 0, 0
		for i, v := range a {
			if l < r && i-idQ[l]+1 > fixedSize {
				l++
			}
			for ; l < r && a[idQ[r-1]] >= v; r-- { // >= 意味着相等的元素取靠右的，若改成 > 表示相等的元素取靠左的
			}
			idQ[r] = i
			r++
			if i+1 >= fixedSize {
				mins = append(mins, a[idQ[l]])
			}
		}

		l, r = 0, 0
		for i, v := range a {
			if l < r && i-idQ[l]+1 > fixedSize {
				l++
			}
			for ; l < r && a[idQ[r-1]] <= v; r-- { // <= 表示首大尾小
			}
			idQ[r] = i
			r++
			if i+1 >= fixedSize {
				maxs = append(maxs, a[idQ[l]])
			}
		}

		return
	}

	// 查询区间的右端点为 i-1    [pop]-query-push
	// 代码来自 LC1499/周赛195D https://leetcode-cn.com/problems/max-value-of-equation/
	findMaxValueOfEquation := func(points [][]int, k int) (ans int) {
		max := func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}
		f := func(p []int) int { return p[1] - p[0] }

		ans = -1e18
		idQ := make([]int, len(points))
		l, r := 0, 0
		for i, p := range points {
			for ; l < r && p[0]-points[idQ[l]][0] > k; l++ {
			}
			if l < r {
				ans = max(ans, p[0]+p[1]+f(points[idQ[l]]))
			}
			for ; l < r && f(points[idQ[r-1]]) <= f(p); r-- {
			}
			idQ[r] = i
			r++
		}
		return
	}

	// 子数组和至少为 k 的最短子数组长度    push-[query-pop]
	// 由于求的是子数组和，可以转化为前缀和之差，若枚举区间右端点 i，则查询的是 [x,i] 的最小值
	// x 为右端点为 i 时的符合和至少为 k 的子数组的左端点
	// LC862 https://leetcode-cn.com/problems/shortest-subarray-with-sum-at-least-k/
	shortestSubarray := func(a []int, k int) (ans int) {
		min := func(a, b int) int {
			if a < b {
				return a
			}
			return b
		}
		n := len(a)

		const inf int = 1e9
		ans = inf
		sum := make([]int, n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + v
		}
		idQ := make([]int, n+1)
		l, r := 0, 0
		for i, s := range sum {
			for ; l < r && sum[idQ[r-1]] >= s; r-- {
			}
			idQ[r] = i
			r++
			for ; l < r && s-sum[idQ[l]] >= k; l++ { // 不断查询+弹出队首直到队列为空或不满足要求
				ans = min(ans, i-idQ[l])
			}
		}
		if ans == inf {
			ans = -1
		}
		return
	}

	// 有区间上界的最大子数组和     pop-push-query
	// https://ac.nowcoder.com/acm/contest/1006/D
	upperSizeMaxSum := func(a []int, upperSize int) (ans int) {
		max := func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}
		n := len(a)
		ans = -1e18

		sum := make([]int, n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + v
		}
		idQ := make([]int, n+1)
		l, r := 0, 0
		for i, s := range sum {
			if l < r && i-idQ[l] > upperSize {
				l++
			}
			for ; l < r && sum[idQ[r-1]] >= s; r-- {
			}
			idQ[r] = i
			r++
			ans = max(ans, s-sum[idQ[l]])
		}
		return
	}

	// 枚举区间左端点更为方便的情况    [query-push]-pop
	// 下面的代码来自 https://codeforces.com/problemset/problem/1237/D
	cf1237d := func(a []int, n int) (ans []int) {
		a = append(append(a, a...), a...)
		idQ := make([]int, 3*n) // 队首为区间最值
		l, r := 0, 0
		for i, j := 0, 0; i < n; i++ { // 枚举区间左端点 i
			// 不断扩大区间右端点 j 直至不满足题目要求
			for ; j < 3*n && (l == r || 2*a[j] >= a[idQ[l]]); j++ {
				for ; l < r && a[idQ[r-1]] <= a[j]; r-- {
				}
				idQ[r] = j
				r++
			}
			maxLen := j - i
			if maxLen > 2*n {
				maxLen = -1
			}
			ans = append(ans, maxLen)
			// 若 i 不在下一个考察区间内，则弹出队首
			if l < r && idQ[l] == i {
				l++
			}
		}
		return
	}

	// 统计区间个数：区间最大值 >= 2*区间最小值
	// todo https://ac.nowcoder.com/acm/contest/6778/C

	_ = []interface{}{
		monotoneStack,
		fixedSizeMinMax, findMaxValueOfEquation, shortestSubarray, upperSizeMaxSum, cf1237d,
	}
}
