package copypasta

func commonCollection() {
	const inf int = 0x3f3f3f3f
	const inf64 int64 = 0x3f3f3f3f3f3f3f3f

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

	dirOffset4 := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	dirOffset8 := [8][2]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

	_ = []interface{}{min, max, abs, quickPow, reverse, dirOffset4, dirOffset8}
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
