package copypasta

import "sort"

func commonCollection() {
	// 注意：若有超过两个数相加，要特别注意 inf 的选择！
	const inf int = 0x3f3f3f3f
	const inf64 int64 = 0x3f3f3f3f3f3f3f3f
	pow2 := [...]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144}
	pow10 := [...]int{1, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9}
	dirOffset4 := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	dirOffset8 := [8][2]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

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

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	const mod int64 = 1e9 + 7
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

	//

	// NOTE: arr must be sorted
	unique := func(arr []int) (newArr []int) {
		n := len(arr)
		if n == 0 {
			return
		}
		newArr = make([]int, 1, n)
		newArr[0] = arr[0]
		for i := 1; i < n; i++ {
			if arr[i] != arr[i-1] {
				newArr = append(newArr, arr[i])
			}
		}
		return
	}

	discrete := func(arr []int, start int) (newArr []int) {
		n := len(arr)
		if n == 0 {
			return
		}
		uniqueArr := make([]int, n)
		copy(uniqueArr, arr)
		sort.Ints(uniqueArr)
		uniqueArr = unique(uniqueArr)
		newArr = make([]int, n)
		for i, v := range arr {
			newArr[i] = start + sort.Search(len(uniqueArr), func(j int) bool { return uniqueArr[j] >= v })
		}
		return
	}

	ifElse := func(cond bool, r1, r2 string) string {
		if cond {
			return r1
		}
		return r2
	}

	_ = []interface{}{pow2, pow10, dirOffset4, dirOffset8, min, mins, max, maxs, abs, quickPow, discrete, ifElse}
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
