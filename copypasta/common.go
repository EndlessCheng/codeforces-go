package copypasta

import (
	. "fmt"
	"io"
	"math"
	"math/bits"
	"math/rand"
	"reflect"
	"sort"
	"unsafe"
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

// 利用前缀和实现巧妙的构造 https://www.luogu.com.cn/blog/duyi/qian-zhui-he
// 邻项修改->前缀和->单项修改 https://codeforces.com/problemset/problem/1254/B2 https://ac.nowcoder.com/acm/contest/7612/C

/* 横看成岭侧成峰

有相当多的题目考察思维角度的转换
一个技巧是，思考：从答案出发，倒推怎样才能得到这个答案
逆向思维 https://leetcode-cn.com/contest/biweekly-contest-9/problems/minimum-time-to-build-blocks/

考虑每个点产生的贡献 https://codeforces.com/problemset/problem/1009/E
考虑每条边产生的负贡献 https://atcoder.jp/contests/abc173/tasks/abc173_f
和式的另一视角。若每一项的值都在一个范围，不妨考虑另一个问题：值为 x 的项有多少个？https://atcoder.jp/contests/abc162/tasks/abc162_e
对所有排列考察所有子区间的性质，可以转换成对所有子区间考察所有排列。将子区间内部的排列和区间外部的排列进行区分，内部的性质单独研究，外部的当作 (n-(r-l))! 个排列 https://codeforces.com/problemset/problem/1284/C
转换为距离的众数 https://codeforces.com/problemset/problem/1365/C
转换为差分数组的变化 https://codeforces.com/problemset/problem/1110/E
不解释，自己感受 https://leetcode-cn.com/contest/biweekly-contest-31/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/
从最大值入手 https://codeforces.com/problemset/problem/1381/B
等效性 https://leetcode-cn.com/contest/biweekly-contest-8/problems/maximum-number-of-ones/

奇偶性 https://codeforces.com/problemset/problem/763/B
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
	pow10 := func(x int) int64 { return int64(math.Pow10(x)) } // 不需要 round

	// TIPS: dir4[i] 和 dir4[i^1] 互为相反方向
	type pair struct{ x, y int }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右
	dir4C := []pair{ // 西东南北
		'W': {-1, 0},
		'E': {1, 0},
		'S': {0, -1},
		'N': {0, 1},
	}
	dir4c := []pair{ // 左右下上
		'L': {-1, 0},
		'R': {1, 0},
		'D': {0, -1},
		'U': {0, 1},
	}
	dir4R := []pair{{1, 1}, {-1, 1}, {-1, -1}, {1, -1}}
	dir8 := []pair{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	perm3 := [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}
	perm4 := [][]int{
		{0, 1, 2, 3}, {0, 1, 3, 2}, {0, 2, 1, 3}, {0, 2, 3, 1}, {0, 3, 1, 2}, {0, 3, 2, 1},
		{1, 0, 2, 3}, {1, 0, 3, 2}, {1, 2, 0, 3}, {1, 2, 3, 0}, {1, 3, 0, 2}, {1, 3, 2, 0},
		{2, 0, 1, 3}, {2, 0, 3, 1}, {2, 1, 0, 3}, {2, 1, 3, 0}, {2, 3, 0, 1}, {2, 3, 1, 0},
		{3, 0, 1, 2}, {3, 0, 2, 1}, {3, 1, 0, 2}, {3, 1, 2, 0}, {3, 2, 0, 1}, {3, 2, 1, 0},
	}

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
	ceil := func(a, b int) int {
		// assert a >= 0 && b > 0
		if a == 0 {
			return 0
		}
		return (a-1)/b + 1
	}
	// 另一种写法，无需考虑 a 为 0 的情况
	ceil = func(a, b int) int {
		return (a + b - 1) / b
	}
	bin := func(v int) []byte {
		const maxLen = 30 // 62 for int64
		s := make([]byte, maxLen+1)
		for i := range s {
			s[i] = byte(v >> (maxLen - i) & 1)
		}
		return s
	}

	// 超过 cap(a) 的数据是未知的
	sliceToArray := func(a []int) [10]int {
		return *(*[10]int)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&a)).Data))
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

	// 顺时针旋转矩阵 90°
	rotate := func(a [][]int) [][]int {
		n, m := len(a), len(a[0])
		b := make([][]int, m)
		for i := range b {
			b[i] = make([]int, n)
		}
		for i, r := range a {
			for j, v := range r {
				b[j][n-1-i] = v
			}
		}
		return b
	}
	// 转置
	transpose := func(a [][]int) [][]int {
		n, m := len(a), len(a[0])
		b := make([][]int, m)
		for i := range b {
			b[i] = make([]int, n)
			for j, r := range a {
				b[i][j] = r[i]
			}
		}
		return b
	}

	// 适用于 mod 超过 int32 范围的情况
	mul := func(a, b, mod int64) (res int64) {
		for ; b > 0; b >>= 1 {
			if b&1 == 1 {
				res = (res + a) % mod
			}
			a = (a + a) % mod
		}
		return
	}

	// 另一种写法，随机数据下比上面的龟速乘快 10 倍左右
	// 这里就假设 a b 均为非负了
	mul = func(a, b, mod int64) int64 {
		hi, lo := bits.Mul64(uint64(a), uint64(b))
		h, l := int64(hi%uint64(mod)), int64(lo%uint64(mod))
		p32 := int64(1) << 32 % mod
		return (p32*p32%mod*h + l) % mod
	}

	// 还有一种用浮点数的写法，此略

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

	// 合并有序数组，保留重复元素
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

	// 返回 a 的各个子集的元素和（不保证返回结果有序）
	subSum := func(a []int) []int {
		sum := make([]int, 1<<len(a)) // int64
		for i, v := range a {
			for j := 0; j < 1<<i; j++ {
				sum[1<<i|j] = sum[j] + v
				// NOTE: 若要直接在此写循环遍历 sum，注意别漏了 sum[0] = 0 的情况
			}
		}
		return sum
	}

	// 返回 a 的各个子集的元素和，且保证返回结果有序
	// 若已求出前 i-1 个数的有序子集和 b，那么前 i 个数的有序子集和可以由 b 和 {b 的每个数加上 a[i]} 归并得到
	// 复杂度为 1+2+4+...+2^n = O(2^n)
	// 参考 https://leetcode-cn.com/problems/closest-subsequence-sum/solution/o2n2de-zuo-fa-by-heltion-0yn7/
	subSumSorted := func(a []int) []int {
		sum := []int{0}
		for _, v := range a {
			b := make([]int, len(sum))
			for i, w := range sum {
				b[i] = w + v
			}
			sum = merge(sum, b)
		}
		return sum
	}

	// 分组前缀和（具体见 query 上的注释）
	// 周赛 216C https://leetcode-cn.com/contest/weekly-contest-216/problems/ways-to-make-a-fair-array/
	groupPrefixSum := func(a []int, k int) {
		// 补 0 简化后续逻辑
		n := len(a)
		for len(a)%k > 0 {
			a = append(a, 0)
		}
		sum := make([]int, len(a)+k) // int64
		for i, v := range a {
			sum[i+k] = sum[i] + v
		}
		pre := func(x, m int) int {
			if x%k <= m {
				return sum[x/k*k+m]
			}
			return sum[(x+k-1)/k*k+m]
		}
		// 求下标在 [l,r) 范围内且下标同余于 m 的元素和 (0<=m<k)
		query := func(l, r, m int) int {
			return pre(r, m) - pre(l, m)
		}
		a = a[:n] // 如果要枚举等，可能需要复原

		_ = query
	}

	// 环形区间和 [l,r) 0<=l<r
	circularRangeSum := func(a []int) {
		n := len(a)
		sum := make([]int64, n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + int64(v)
		}
		pre := func(p int) int64 {
			return sum[n]*int64(p/n) + sum[p%n]
		}
		query := func(l, r int) int64 {
			return pre(r) - pre(l)
		}

		_ = query
	}

	// 带权(等差数列)前缀和
	{
		var n int // read
		a := make([]int64, n)
		// read a ...

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
	initSum2D := func(a [][]int) {
		n, m := len(a), len(a[0])
		sum2d = make([][]int, n+1)
		sum2d[0] = make([]int, m+1)
		for i, row := range a {
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

	// 利用每个数产生的贡献计算 Σ|ai-aj|, i!=j
	// 相关题目 https://codeforces.com/contest/1311/problem/F
	contributionSum := func(a []int) (sum int64) {
		n := len(a)
		sort.Ints(a)
		for i, v := range a {
			sum += int64(v) * int64(2*i+1-n)
		}
		return
	}

	// 二维差分
	// todo https://blog.csdn.net/weixin_43914593/article/details/113782108
	//      https://www.luogu.com.cn/problem/P3397

	reverse := func(a []byte) []byte {
		n := len(a)
		b := make([]byte, n)
		for i, v := range a {
			b[n-1-i] = v
		}
		return b
	}
	reverseInPlace := func(a []byte) {
		for i, n := 0, len(a); i < n/2; i++ {
			a[i], a[n-1-i] = a[n-1-i], a[i]
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

	// 求交集简洁写法
	intersection := func(a, b []int) []int {
		mp := map[int]bool{}
		for _, v := range a {
			mp[v] = true
		}
		mp2 := map[int]bool{}
		for _, v := range b {
			if mp[v] {
				mp2[v] = true
			}
		}
		mp = mp2

		keys := make([]int, 0, len(mp))
		for k := range mp {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		return keys
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

	// 去重
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

	// 直接在 a 上去重
	uniqueInPlace := func(a []int) []int {
		n := len(a)
		if n == 0 {
			return nil
		}
		k := 0
		for _, v := range a[1:] {
			if a[k] != v {
				k++
				a[k] = v
			}
		}
		//n = k + 1
		return a[:k+1]
	}

	// 离散化，不保留原始数据（保留原始数据的版本见下面的 discreteMap）
	// discrete([]int{100,20,50,50}, 1) => []int{3,1,2,2}
	// https://leetcode-cn.com/contest/biweekly-contest-18/problems/rank-transform-of-an-array/
	discrete := func(a []int, startIndex int) (kth []int) {
		if len(a) == 0 {
			return
		}

		type pair struct{ v, i int }
		ps := make([]pair, len(a))
		for i, v := range a {
			ps[i] = pair{v, i}
		}
		sort.Slice(ps, func(i, j int) bool { return ps[i].v < ps[j].v }) // or SliceStable
		kth = make([]int, len(a))

		// a 有重复元素
		k := startIndex
		kth[ps[0].i] = k
		for i := 1; i < len(ps); i++ {
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

	// 简化版，不要求值连续 [10,30,20,20] => [0,3,1,1]
	discrete2 := func(a []int, startIndex int) []int {
		b := append([]int(nil), a...)
		sort.Ints(b)
		for i, v := range a {
			a[i] = sort.SearchInts(b, v) + startIndex
		}
		return a
	}

	// 保留原始数据的离散化
	// 返回一个名次 map
	// discreteMap([]int{100,20,20,50}, 1) => map[int]int{20:1, 50:2, 100:3}
	// 例题：LC327 https://leetcode-cn.com/problems/count-of-range-sum/
	discreteMap := func(a []int, startIndex int) (kth map[int]int) {
		// assert len(a) > 0
		sorted := append([]int(nil), a...)
		sort.Ints(sorted)

		// 有重复元素
		k := startIndex
		kth = map[int]int{sorted[0]: k}
		for i := 1; i < len(sorted); i++ {
			if sorted[i] != sorted[i-1] {
				k++
				kth[sorted[i]] = k
			}
		}

		// 无重复元素
		kth = make(map[int]int, len(sorted))
		for i, v := range sorted {
			kth[v] = i + startIndex
		}

		// EXTRA: 第 k 小元素在原数组中的下标 kthPos
		pos := make(map[int][]int, k-startIndex)
		for i, v := range a {
			pos[v] = append(pos[v], i)
		}
		kthPos := make([][]int, k+1)
		for v, k := range kth {
			kthPos[k] = pos[v]
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
	//       LC973 https://leetcode-cn.com/problems/k-closest-points-to-origin/submissions/
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
	// 线段相交统计（栈） https://codeforces.com/contest/1278/problem/D
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
			return a.pos < b.pos || a.pos == b.pos && a.delta < b.delta // 先出后进。改成 a.delta > b.delta 为先进后出
		})

		for _, e := range events {
			if e.delta > 0 {

			} else {

			}
		}
	}

	// 扫描线另一种写法，把 delta 压缩进 pos
	// 这样可以避免写一个复杂的 sort.Slice
	sweepLine2 := func(in io.Reader, n int) {
		events := make([]int, 0, 2*n)
		for i := 0; i < n; i++ {
			var l, r int
			Fscan(in, &l, &r)
			events = append(events, l<<1|1, r<<1) // 先出后进
			//events = append(events, l<<1, r<<1|1) // 先进后出
		}
		sort.Ints(events)

		for _, e := range events {
			pos, delta := e>>1, e&1
			_ = pos
			if delta > 0 { // 根据上面的写法来定义何为出何为进

			} else {

			}
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
		pow10, dir4, dir4C, dir4c, dir4R, dir8, perm3, perm4,
		min, mins, max, maxs, abs, ceil, bin,
		sliceToArray,
		isDigit, isLower, isUpper, isAlpha,
		ternaryI, ternaryS, zip, zipI, rotate, transpose, minString,
		pow, mul, toAnyBase, digits,
		subSum, subSumSorted, groupPrefixSum, circularRangeSum, initSum2D, querySum2D,
		contributionSum,
		sort3, reverse, reverseInPlace, equal,
		merge, splitDifferenceAndIntersection, intersection, isSubset, isSubSequence, isDisjoint,
		unique, uniqueInPlace, discrete, discrete2, discreteMap, indexMap, allSame, complement, quickSelect, contains, containsAll,
		sweepLine, sweepLine2, countCoveredPoints,
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

	// 稀疏表 Sparse Table
	// st[i][j] 对应的区间是 [i, i+1<<j)
	// https://oi-wiki.org/ds/sparse-table/
	// https://codeforces.com/blog/entry/66643
	// 模板中的核心函数 core 可以换成其他具有区间合并性质的函数（允许区间重叠），如 gcd 等
	// 模板题 https://www.luogu.com.cn/problem/P3865
	// 变长/种类 https://www.jisuanke.com/contest/11346/challenges
	// 题目推荐 https://cp-algorithms.com/data_structures/sparse-table.html#toc-tgt-5
	const mx = 17 // 17: n<131072, 18: n<262144, 19: n<524288, 20: n<1048576     mx = bits.Len(uint(n))
	core := max
	var st [][mx]int
	stInit := func(a []int) {
		n := len(a)
		st = make([][mx]int, n)
		for i, v := range a {
			st[i][0] = v
		}
		for j := 1; 1<<j <= n; j++ {
			for i := 0; i+1<<j <= n; i++ {
				st[i][j] = core(st[i][j-1], st[i+1<<(j-1)][j-1])
			}
		}
	}
	// [l,r) 注意 l r 是从 0 开始算的
	stQuery := func(l, r int) int { k := bits.Len(uint(r-l)) - 1; return core(st[l][k], st[r-1<<k][k]) }

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
	// 浅谈基础根号算法——分块 https://www.luogu.com.cn/blog/deco/qian-tan-ji-chu-gen-hao-suan-fa-fen-kuai
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
			b.sorted = append([]int(nil), b.origin...)
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
// todo 各类莫队综述
//    https://www.cnblogs.com/WAMonster/p/10118934.html
//    https://ouuan.github.io/post/%E8%8E%AB%E9%98%9F%E5%B8%A6%E4%BF%AE%E8%8E%AB%E9%98%9F%E6%A0%91%E4%B8%8A%E8%8E%AB%E9%98%9F%E8%AF%A6%E8%A7%A3/
//    https://blog.csdn.net/weixin_43914593/article/details/108485396
// todo【推荐】文章及题单 https://www.luogu.com.cn/training/2914
// https://cp-algorithms.com/data_structures/sqrt_decomposition.html#toc-tgt-8
func moAlgorithm() {
	// 普通莫队（没有修改操作）
	// 本质是通过巧妙地改变询问的顺序，使区间左右端点移动的次数由 O(nm) 降为了 O(n√m)
	// 在块大小取 n/√m 时可达到最优复杂度 O(n√m)，见 https://www.luogu.com.cn/blog/codesonic/mosalgorithm
	// https://oi-wiki.org/misc/mo-algo/
	// 模板题 https://www.luogu.com.cn/problem/P1494
	// todo https://www.luogu.com.cn/problem/P2709
	// todo https://www.luogu.com.cn/problem/P4462
	// 区间 mex https://blog.csdn.net/includelhc/article/details/79593496
	//     反向构造题 https://www.luogu.com.cn/problem/P6852
	// todo https://codeforces.com/contest/86/problem/D
	//      https://codeforces.com/contest/220/problem/B
	//      https://codeforces.com/contest/617/problem/E
	//      https://codeforces.com/contest/877/problem/F
	//      https://www.codechef.com/problems/QCHEF
	mo := func(in io.Reader, a []int, q int) []int {
		n := len(a)
		blockSize := int(math.Ceil(float64(n) / math.Sqrt(float64(q))))
		type query struct{ lb, l, r, qid int }
		qs := make([]query, q)
		for i := range qs {
			var l, r int
			Fscan(in, &l, &r) // 从 1 开始，[l,r)
			qs[i] = query{l / blockSize, l, r + 1, i}
		}
		sort.Slice(qs, func(i, j int) bool {
			a, b := qs[i], qs[j]
			if a.lb != b.lb {
				return a.lb < b.lb
			}
			// 奇偶化排序
			if a.lb&1 == 0 {
				return a.r < b.r
			}
			return a.r > b.r
		})

		cnt := 0
		l, r := 1, 1 // 区间从 1 开始，方便 debug
		move := func(idx, delta int) {
			// NOTE: 有些题目在 delta 为 1 和 -1 时逻辑的顺序是严格对称的
			// v := a[idx-1]
			// ...
			// cnt += delta
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
			for ; r < q.r; r++ {
				move(r, 1)
			}
			for ; l < q.l; l++ {
				move(l, -1)
			}
			for l > q.l {
				l--
				move(l, 1)
			}
			for r > q.r {
				r--
				move(r, -1)
			}
			ans[q.qid] = getAns(q)
		}
		return ans
	}

	// 带修莫队（支持单点修改）
	// https://oi-wiki.org/misc/modifiable-mo-algo/
	// https://codeforces.com/blog/entry/72690
	// 模板题 数颜色 https://www.luogu.com.cn/problem/P1903
	// https://codeforces.com/problemset/problem/940/F
	// https://codeforces.com/problemset/problem/1476/G
	// todo https://www.codechef.com/FEB17/problems/DISTNUM3
	// todo 二逼平衡树（树套树）https://www.luogu.com.cn/problem/P3380
	moWithUpdate := func(in io.Reader) []int {
		var n, q int
		Fscan(in, &n, &q)
		a := make([]int, n+1) // 从 1 开始，方便 debug
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}
		blockSize := int(math.Round(math.Pow(float64(n), 2.0/3)))
		type query struct{ lb, rb, l, r, t, qid int }
		type modify struct{ pos, val int }
		qs := []query{}
		ms := []modify{}
		for ; q > 0; q-- {
			var op string
			if Fscan(in, &op); op[0] == 'Q' {
				var l, r int
				Fscan(in, &l, &r)
				// 改成左闭右开
				qs = append(qs, query{l / blockSize, (r + 1) / blockSize, l, r + 1, len(ms), len(qs)})
			} else {
				var pos, val int
				Fscan(in, &pos, &val)
				ms = append(ms, modify{pos, val})
			}
		}
		sort.Slice(qs, func(i, j int) bool {
			a, b := qs[i], qs[j]
			if a.lb != b.lb {
				return a.lb < b.lb
			}
			if a.rb != b.rb {
				if a.lb&1 == 0 {
					return a.rb < b.rb
				}
				return a.rb > b.rb
			}
			if a.rb&1 == 0 {
				return a.t < b.t
			}
			return a.t > b.t
		})

		const mx int = 1e6 // TODO
		cnt, cc := [mx + 1]int{}, 0
		l, r, now := 1, 1, 0
		add := func(val int) {
			if cnt[val] == 0 {
				cc++
			}
			cnt[val]++
		}
		del := func(val int) {
			cnt[val]--
			if cnt[val] == 0 {
				cc--
			}
		}
		// 注：由于函数套函数不会内联，直接写到主流程的 for now 循环中会快不少
		timeSlip := func(l, r int) {
			m := ms[now]
			p, v := m.pos, m.val
			if l <= p && p < r {
				del(a[p])
				add(v)
			}
			a[p], ms[now].val = v, a[p]
		}
		getAns := func(q query) int {
			// 提醒：q.r 是加一后的，计算时需要注意
			// sz := q.r - q.l
			// ...
			return cc
		}
		ans := make([]int, len(qs))
		for _, q := range qs {
			for ; r < q.r; r++ {
				add(a[r])
			}
			for ; l < q.l; l++ {
				del(a[l])
			}
			for l > q.l {
				l--
				add(a[l])
			}
			for r > q.r {
				r--
				del(a[r])
			}
			for ; now < q.t; now++ {
				timeSlip(q.l, q.r)
			}
			for now > q.t {
				now--
				timeSlip(q.l, q.r)
			}
			ans[q.qid] = getAns(q)
		}
		return ans
	}

	// 回滚莫队
	// 复杂度同普通莫队
	// https://oi-wiki.org/misc/rollback-mo-algo/
	// 浅谈回滚莫队 https://www.luogu.com.cn/blog/bfqaq/qian-tan-hui-gun-mu-dui
	// todo 回滚莫队及其简单运用 https://www.cnblogs.com/Parsnip/p/10969989.html
	// 模板题 历史研究 https://www.luogu.com.cn/problem/AT1219 https://atcoder.jp/contests/joisc2014/tasks/joisc2014_c
	// todo https://www.luogu.com.cn/problem/P5906
	// todo https://www.luogu.com.cn/problem/P5386
	// todo https://www.luogu.com.cn/problem/P6072
	rollbackMo := func(in io.Reader) []int {
		var n, q int
		Fscan(in, &n, &q)
		a := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}
		ans := make([]int, q)
		B := int(math.Ceil(float64(n) / math.Sqrt(float64(q))))
		type query struct{ lb, l, r, qid int }
		qs := []query{}
		cnt := make([]int, n+1)
		for i := 0; i < q; i++ {
			var l, r int
			Fscan(in, &l, &r)
			r++
			if r-l > B {
				qs = append(qs, query{l / B, l, r, i})
				continue
			}
			// 小区间暴力计算
			res := 0
			for _, v := range a[l:r] {
				cnt[v]++
				// ...
			}
			ans[i] = res
			// 重置数据 ...
			for _, v := range a[l:r] {
				cnt[v] = 0
			}
		}
		sort.Slice(qs, func(i, j int) bool {
			a, b := qs[i], qs[j]
			if a.lb != b.lb {
				return a.lb < b.lb
			}
			return a.r < b.r
		})

		var l, r int
		var res int
		add := func(i int) {
			v := a[i]
			cnt[v]++
			// ...
		}
		getAns := func(q query) int {
			// ...
			return res
		}
		for i, q := range qs {
			l0 := (q.lb + 1) * B
			if i == 0 || q.lb > qs[i-1].lb {
				l, r = l0, l0
				// 重置数据 ...
				res = 0
				cnt = make([]int, n+1)
			}
			for ; r < q.r; r++ {
				add(r)
			}
			tmp := res
			for l > q.l {
				l--
				add(l)
			}
			ans[q.qid] = getAns(q)
			res = tmp
			for ; l < l0; l++ {
				// 回滚 ...
				cnt[a[l]]--
			}
		}
		return ans
	}

	// 树上莫队
	// 通过 DFS 序转化成序列上的查询
	// NOTE: 对于带修莫队，去掉 timeSlip 中的参数，且 if l <= p && p < r 替换成 if vis[p] https://www.luogu.com.cn/record/46714923
	// https://oi-wiki.org/misc/mo-algo-on-tree/
	// 有关树分块的内容见 graph_tree.go 中的 limitSizeDecomposition
	// 模板题 糖果公园 https://www.luogu.com.cn/problem/P4074
	//       https://www.acwing.com/problem/content/2536/ https://www.luogu.com.cn/problem/SP10707
	moOnTree := func(n, root, q int, g [][]int, vals []int) []int {
		vs := make([]int, 0, 2*n)
		tin := make([]int, n)
		tout := make([]int, n)
		var initTime func(v, fa int)
		initTime = func(v, fa int) {
			tin[v] = len(vs)
			vs = append(vs, v)
			for _, w := range g[v] {
				if w != fa {
					initTime(w, v)
				}
			}
			tout[v] = len(vs)
			vs = append(vs, v)
		}
		initTime(root, -1)

		// initTime 的逻辑可以并到求 pa dep 的 DFS 中
		var _lca func(v, w int) int // 见 tree.lcaBinarySearch

		blockSize := int(math.Ceil(float64(2*n) / math.Sqrt(float64(q)))) // int(math.Round(math.Pow(float64(2*n), 2.0/3)))
		type query struct{ lb, l, r, lca, qid int }
		qs := make([]query, q)
		for i := range qs {
			var v, w int
			//Fscan(in, &v, &w)
			v--
			w--
			if tin[v] > tin[w] {
				v, w = w, v
			}
			if lca := _lca(v, w); lca != v {
				qs[i] = query{tout[v] / blockSize, tout[v], tin[w] + 1, lca, i}
			} else {
				qs[i] = query{tin[v] / blockSize, tin[v], tin[w] + 1, -1, i}
			}
		}
		sort.Slice(qs, func(i, j int) bool {
			a, b := qs[i], qs[j]
			if a.lb != b.lb {
				return a.lb < b.lb
			}
			if a.lb&1 == 0 {
				return a.r < b.r
			}
			return a.r > b.r
		})

		var k int // vals 不同元素个数
		cnt := make([]int, k+1)
		cc := 0
		l, r := 0, 0
		vis := make([]bool, n)
		move := func(v int) {
			x := vals[v]
			if vis[v] = !vis[v]; vis[v] {
				if cnt[x] == 0 {
					cc++
				}
				cnt[x]++
			} else {
				cnt[x]--
				if cnt[x] == 0 {
					cc--
				}
			}
		}
		getAns := func(q query) int {
			return cc
		}
		ans := make([]int, q)
		for _, q := range qs {
			for ; r < q.r; r++ {
				move(vs[r])
			}
			for ; l < q.l; l++ {
				move(vs[l])
			}
			for l > q.l {
				l--
				move(vs[l])
			}
			for r > q.r {
				r--
				move(vs[r])
			}
			if q.lca >= 0 {
				move(q.lca)
			}
			ans[q.qid] = getAns(q)
			if q.lca >= 0 {
				move(q.lca)
			}
		}
		return ans
	}

	// 二次离线莫队
	// todo https://www.luogu.com.cn/problem/P4887

	_ = []interface{}{mo, moWithUpdate, rollbackMo, moOnTree}
}

// 单调队列（相关代码见 monotoneCollection 中的「单调队列」部分）
type mqData struct {
	val int
	del int // 懒删除标记
}
type monotoneQueue struct {
	data []mqData // 初始化时可以 make([]mqData, 0, n) 来减少扩容的开销
	size int      // 单调队列对应的区间的长度
}

func (mq monotoneQueue) less(a, b mqData) bool { return a.val >= b.val } // >= 维护区间最大值；<= 维护区间最小值
func (mq *monotoneQueue) push(v int) {
	mq.size++
	d := mqData{v, 1}
	for len(mq.data) > 0 && mq.less(d, mq.data[len(mq.data)-1]) {
		d.del += mq.data[len(mq.data)-1].del
		mq.data = mq.data[:len(mq.data)-1]
	}
	mq.data = append(mq.data, d)
}
func (mq *monotoneQueue) pop() {
	mq.size--
	if mq.data[0].del > 1 {
		mq.data[0].del--
	} else {
		mq.data = mq.data[1:]
	}
}
func (mq monotoneQueue) top() int { return mq.data[0].val } // 调用前需要判断 size > 0

func monotoneCollection() {
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

	// 推荐 https://cp-algorithms.com/data_structures/stack_queue_modification.html

	// 单调栈
	// 举例：返回每个元素两侧严格大于它的元素位置（不存在则为 -1 或 n）
	// 如何理解：把数组想象成一列山峰，站在 a[i] 的山顶仰望两侧的山峰，是看不到高山背后的矮山的，只能看到一座座更高的山峰
	//          这就启发我们引入一个底大顶小的单调栈，入栈时不断比较栈顶元素直到找到一个比当前元素大的
	// 技巧：事先压入一个边界元素到栈底，这样保证循环时栈一定不会为空，从而简化逻辑
	// 一些转换：
	//     若区间 [l,r] 的最大值等于 a[r]，则 l 必须 > posL[r]
	//     若区间 [l,r] 的最大值等于 a[l]，则 r 必须 < posR[l]
	//     这一结论可以用于思考一些双变量的题目
	// https://oi-wiki.org/ds/monotonous-stack/
	// 模板题 https://www.luogu.com.cn/problem/P5788
	//       https://www.luogu.com.cn/problem/P2866 http://poj.org/problem?id=3250
	//       https://leetcode-cn.com/problems/next-greater-element-i/ LC496/周赛18BA
	//       https://leetcode-cn.com/problems/next-greater-element-ii/ LC503/周赛18BB
	// 柱状图中最大的矩形 LC84 https://leetcode-cn.com/problems/largest-rectangle-in-histogram/ http://poj.org/problem?id=2559 http://poj.org/problem?id=2082
	// 最大全 1 矩形 LC85（实现见下面的 maximalRectangleArea）https://leetcode-cn.com/problems/maximal-rectangle/
	// 接雨水 LC42 https://leetcode-cn.com/problems/trapping-rain-water/
	// 后缀数组+不同矩形对应方案数之和 https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/D
	// 与 DP 结合
	//     https://codeforces.com/problemset/problem/1313/C2
	//     https://codeforces.com/problemset/problem/1407/D
	// 全 1 子矩阵个数 O(n^2) LC周赛196C https://leetcode-cn.com/contest/weekly-contest-196/problems/count-submatrices-with-all-ones/ 原题为 http://poj.org/problem?id=3494
	// * 已知部分 posR 还原全部 posR；已知 posR 还原 a https://codeforces.com/problemset/problem/1158/C
	monotoneStack := func(a []int) ([]int, []int) {
		const border int = 2e9 // 求两侧小的话用 -1
		type pair struct{ v, i int }

		// 求左侧严格大于
		n := len(a)
		posL := make([]int, n)
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

	// 最大全 1 矩形 LC85 https://leetcode-cn.com/problems/maximal-rectangle/
	maximalRectangleArea := func(a [][]int) int {
		const target = 1
		n, m, ans := len(a), len(a[0]), 0
		heights := make([][]int, n) // heights(i,j) 表示从 (i,j) 往上看的高度，a(i,j) = 0 时为 0
		for i, row := range a {
			heights[i] = make([]int, m)
			for j, v := range row {
				if v == target {
					if i == 0 {
						heights[i][j] = 1
					} else {
						heights[i][j] = heights[i-1][j] + 1
					}
				}
			}
		}
		type pair struct{ h, i int }
		for _, hs := range heights {
			posL := make([]int, m)
			stack := []pair{{-1, -1}}
			for j, h := range hs {
				for {
					if top := stack[len(stack)-1]; top.h < h {
						posL[j] = top.i
						break
					}
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, pair{h, j})
			}
			posR := make([]int, m)
			stack = []pair{{-1, m}}
			for j := m - 1; j >= 0; j-- {
				h := hs[j]
				for {
					if top := stack[len(stack)-1]; top.h < h {
						posR[j] = top.i
						break
					}
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, pair{h, j})
			}
			for j, h := range hs {
				ans = max(ans, (posR[j]-posL[j]-1)*h)
			}
		}
		return ans
	}

	/* 单调队列
	需要不断维护队列的单调性，时刻保证队列元素从大到小或从小到大
	https://oi-wiki.org/ds/monotonous-queue/
	https://oi-wiki.org/dp/opt/monotonous-queue-stack/
	https://blog.csdn.net/weixin_43914593/article/details/105791217 算法竞赛专题解析（13）：DP优化(3)--单调队列优化
	todo https://www.luogu.com.cn/problem/P2627

	NOTE: 某些题目需要特殊判断数组长度为 1 的情况

	todo http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=1070
	*/

	// 模板题 - 固定区间大小的区间最值
	// https://www.luogu.com.cn/problem/P1886 http://poj.org/problem?id=2823
	// https://codeforces.com/problemset/problem/940/E
	fixedSizeMax := func(a []int, fixedSize int) []int {
		n := len(a)
		q := monotoneQueue{} // 最大/最小由 less 来控制
		ans := make([]int, 0, n-fixedSize+1)
		for i, v := range a {
			q.push(v)
			if q.size > fixedSize {
				q.pop()
			}
			// 插入新元素并保证单调队列大小后，获取区间最值
			if i+1 >= fixedSize {
				ans = append(ans, q.top())
			}
		}
		return ans
	}

	// 模板题 - 最大子序和
	// https://www.acwing.com/problem/content/137/ https://ac.nowcoder.com/acm/contest/1006/D
	maxSubSumWithLimitSize := func(a []int, sizeLimit int) int {
		n := len(a)
		sum := make([]int, n+1) // int64
		for i, v := range a {
			sum[i+1] = sum[i] + v
		}
		ans := int(-1e9)     // -1e18
		q := monotoneQueue{} // 维护区间最小值
		q.push(sum[0])
		for r := 1; r <= n; r++ {
			if q.size > sizeLimit {
				q.pop()
			}
			ans = max(ans, sum[r]-q.top())
			q.push(sum[r])
		}
		return ans
	}

	// 子数组和至少为 k 的最短非空子数组长度
	// 这题的关键在于，当右端点向右（枚举）时，左端点是绝对不会向左的（因为向左肯定会比当前求出的最短长度要长）
	// 想明白这一点就可以愉快地使用单调队列了
	// LC862 https://leetcode-cn.com/problems/shortest-subarray-with-sum-at-least-k/
	shortestSubSumAtLeastK := func(a []int, k int) int {
		n := len(a)
		ans := n + 1
		sum := make([]int, n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + v
		}
		q := monotoneQueue{} // 维护区间最小值
		q.push(sum[0])
		for r := 1; r <= n; r++ {
			for q.size > 0 && sum[r]-q.top() >= k {
				ans = min(ans, q.size)
				q.pop()
			}
			q.push(sum[r])
		}
		if ans > n {
			return -1
		}
		return ans
	}

	// 枚举区间左端点更为方便的情况 · 其一
	// 统计区间个数：区间最大值 >= 2*区间最小值
	// https://ac.nowcoder.com/acm/contest/6778/C
	//
	// 思路：转变成求「区间最大值 < 2*区间最小值」的区间个数
	// 随着左端点向右，右端点必然不会向左
	countSubArrayByMinMax := func(a []int) int {
		n := len(a)
		ans := n * (n + 1) / 2 // int64
		mx := monotoneQueue{}  // 维护区间最大值
		mi := monotoneQueue{}  // 维护区间最小值（需要新定义一个有不同 less 的 monotoneQueue）
		for i, j := 0, 0; i < n; i++ {
			// 确保符合条件再插入
			for ; j < n && (mx.size == 0 || mi.size == 0 || max(mx.top(), a[j]) < 2*min(mi.top(), a[j])); j++ {
				mx.push(a[j])
				mi.push(a[j])
			}
			sz := j - i
			ans -= sz
			// 若单调队列指向的区间的左端点为 i，则对应元素在下一次循环时将不再使用。故弹出之
			if mx.size == sz {
				mx.pop()
			}
			if mi.size == sz {
				mi.pop()
			}
		}
		return ans
	}

	// 枚举区间左端点更为方便的情况 · 其二
	// https://codeforces.com/problemset/problem/1237/D
	// 注意这题和 countSubArrayByMinMax 的不同之处：不满足要求的最小值一定要在最大值的右侧
	// 也可以枚举右端点，见 https://www.luogu.com.cn/blog/qianshang/solution-cf1237d
	balancedPlaylist := func(a []int, n int) (ans []int) {
		a = append(append(a, a...), a...)
		q := monotoneQueue{} // 维护区间最大值
		for i, j := 0, 0; i < n; i++ {
			// 不断扩大区间右端点 j 直至不满足题目要求
			for ; j < 3*n && (q.size == 0 || q.top() <= 2*a[j]); j++ {
				q.push(a[j])
			}
			sz := j - i
			if sz > 2*n {
				sz = -1
			}
			ans = append(ans, sz)
			if q.size == sz {
				q.pop()
			}
		}
		return
	}

	_ = []interface{}{
		monotoneStack, maximalRectangleArea,
		fixedSizeMax, maxSubSumWithLimitSize, shortestSubSumAtLeastK, countSubArrayByMinMax, balancedPlaylist,
	}
}
