package copypasta

func commonCollection() {
	// 注意：若有超过两个数相加，要特别注意 inf 的选择！
	const inf int = 0x3f3f3f3f
	const inf64 int64 = 0x3f3f3f3f3f3f3f3f
	pow2 := [...]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144}

	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a >= b {
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

	ifelse := func(cond bool, r1, r2 string) string {
		if cond {
			return r1
		}
		return r2
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	const mod = int64(1e9 + 7)
	quickPow := func(x, n int64) int64 {
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	// NOTE: Golang already has a binary search function in sort package, see 1077D for example
	type bsFunc func(int) bool
	reverse := func(f bsFunc) bsFunc {
		return func(i int) bool {
			return !f(i)
		}
	}
	//sort.Search(n, reverse(func(i int) bool {...}))

	// step = 100
	binarySearchF := func(l, r float64, step int, f func(x float64) bool) float64 {
		for i := 0; i < step; i++ {
			mid := (l + r) / 2
			if f(mid) {
				r = mid
			} else {
				l = mid
			}
			// 在精度容易确定时，可以加上 if r-l < eps {break}
			// 例如保留 4 位小数时，eps 取 1e-6
		}
		return (l + r) / 2
	}

	// step = 100
	ternarySearch := func(l, r float64, step int, f func(x float64) float64) float64 {
		for i := 0; i < step; i++ {
			m1 := l + (r-l)/3
			m2 := r - (r-l)/3
			v1, v2 := f(m1), f(m2)
			if v1 < v2 {
				r = m2 // 若求最大值，则 l = m1
			} else {
				l = m1 // 若求最大值，则 r = m2
			}
		}
		return (l + r) / 2
	}

	dirOffset4 := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	dirOffset8 := [8][2]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

	_ = []interface{}{pow2, min, mins, max, maxs, ifelse, abs, quickPow, reverse, binarySearchF, ternarySearch, dirOffset4, dirOffset8}
}

// Permute the values at index i to len(arr)-1.
// See 910C for example.
func permute(arr []int, i int, do func([]int)) {
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
