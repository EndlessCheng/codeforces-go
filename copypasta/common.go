package copypasta

import (
	"math/bits"
	"sort"
)

func commonCollection() {
	const mod int64 = 1e9 + 7
	// 注意：若有超过两个数相加，要特别注意 inf 的选择！
	const inf int = 0x3f3f3f3f
	const inf64 int64 = 0x3f3f3f3f3f3f3f3f
	pow2 := [...]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144}
	pow10 := [...]int{1, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9}
	dirOffset4 := [...][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	dirOffset4R := [...][2]int{{1, 1}, {-1, 1}, {-1, -1}, {1, -1}}
	dirOffset8 := [...][2]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	orders := [6][3]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}

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

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	quickPow := func(x, n, mod int64) int64 {
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	//

	reverse := func(s []byte) {
		for i, j := 0, len(s)-1; i < j; {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
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

	discrete := func(arr []int, start int) (disArr []int) {
		n := len(arr)
		if n == 0 {
			return
		}
		type pair struct {
			val int
			idx int
		}
		pairs := make([]pair, n)
		for i, val := range arr {
			pairs[i] = pair{val, i}
		}
		sort.Slice(pairs, func(i, j int) bool { return pairs[i].val < pairs[j].val })
		disArr = make([]int, n)
		disVal := start
		disArr[pairs[0].idx] = disVal
		for i := 1; i < n; i++ {
			if pairs[i].val != pairs[i-1].val {
				disVal++
			}
			disArr[pairs[i].idx] = disVal
		}
		return
	}

	ifElse := func(cond bool, r1, r2 string) string {
		if cond {
			return r1
		}
		return r2
	}

	var d [][20]int
	rmqInit := func(a []int) {
		n := len(a)
		d = make([][20]int, n)
		for i := range d {
			d[i][0] = a[i]
		}
		for j := uint(1); 1<<j <= n; j++ {
			for i := 0; i+(1<<j)-1 < n; i++ {
				d[i][j] = max(d[i][j-1], d[i+(1<<(j-1))][j-1])
			}
		}
	}
	rmq := func(l, r int) int { // [l,r] 注意 l r 是从 0 开始算的
		k := uint(bits.Len(uint(r-l+1)) - 1)
		return max(d[l][k], d[r-(1<<k)+1][k])
	}

	var s string
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-'a']++
		//cnt[c-'A']++
		//cnt[c-'0']++
	}

	_ = []interface{}{
		pow2, pow10, dirOffset4, dirOffset4R, dirOffset8, orders, searchDirOffset4, searchDirOffset4R,
		min, mins, max, maxs, abs, quickPow,
		reverse, unique, discrete, ifElse,
		rmqInit, rmq,
	}
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
