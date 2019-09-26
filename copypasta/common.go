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

	_ = []interface{}{min, max, abs, quickPow}
}

// Permute the values at index i to len(arr)-1.
func permutations(arr []int, i int, f func([]int)) {
	if i == len(arr) {
		f(arr)
		return
	}
	permutations(arr, i+1, f)
	for j := i + 1; j < len(arr); j++ {
		arr[i], arr[j] = arr[j], arr[i]
		permutations(arr, i+1, f)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
