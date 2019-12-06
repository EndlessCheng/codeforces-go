package copypasta

import (
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

	var sum2d [][]int
	initSum2D := func(mat [][]int) {
		sum2d = make([][]int, len(mat)+1)
		sum2d[0] = make([]int, len(mat[0])+1)
		for i, mi := range mat {
			sum2d[i+1] = make([]int, len(mi)+1)
			for j, mij := range mi {
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

	sort3 := func(a ...int) (int, int, int) {
		sort.Ints(a)
		return a[0], a[1], a[2]
	}

	reverse := func(s []byte) {
		for i, j := 0, len(s)-1; i < j; {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}
	reverseS := func(ss string) string {
		n := len(ss)
		s := make([]byte, n)
		for i := range s {
			s[i] = ss[n-1-i]
		}
		return string(s)
	}

	// NOTE: arr must be sorted
	unique := func(arr []int) (uniqueArr []int) {
		n := len(arr)
		if n == 0 {
			return
		}
		uniqueArr = make([]int, 1, n)
		uniqueArr[0] = arr[0]
		for i := 1; i < n; i++ {
			if arr[i] != arr[i-1] {
				uniqueArr = append(uniqueArr, arr[i])
			}
		}
		return
	}

	// discrete([]int{100,20,50,50}, 1) => []int{3,1,2,2}
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

	// 单调栈优化
	// TODO: CF1237D

	//

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

	_ = []interface{}{
		pow2, pow10, dirOffset4, dirOffset4R, dirOffset8, orderP3,
		min, mins, max, maxs, ifElseI, ifElseS,
		abs, quickPow, initSum2D, querySum2D,
		sort3, reverse, reverseS, unique, discrete, floatToRat, complement,
		stInit, stQuery,
	}
}

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
		# * #
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
		# * #
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

	_ = []interface{}{loopSet, dfsGrids, searchDirOffset4, searchDirOffset4R}
}

// Permute the values at index i to len(arr)-1.
// See 910C for example.
func _permute(arr []int, i int, do func([]int)) {
	if i == len(arr) {
		do(arr)
		return
	}
	_permute(arr, i+1, do)
	for j := i + 1; j < len(arr); j++ {
		arr[i], arr[j] = arr[j], arr[i]
		_permute(arr, i+1, do)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
func permute(arr []int, do func([]int)) { _permute(arr, 0, do) }

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
// n*(n-1)*(n*n-5*n+18)/24+1
