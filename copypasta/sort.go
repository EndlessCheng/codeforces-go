package copypasta

func sortCollections() {
	// NOTE: Golang already has a binary search function in sort package, see 1077D for example
	type bsFunc func(int) bool
	reverse := func(f bsFunc) bsFunc {
		return func(i int) bool {
			return !f(i)
		}
	}
	// 写法1: sort.Search(n, reverse(func(i int) bool {...}))
	// 写法2:
	// sort.Search(n, func(i int) (ok bool) {
	//	 defer func() { ok = !ok }()
	//	 ...
	// })
	// 写法3（推荐）:
	// sort.Search(n, func(i int) (ok bool) {
	//	 ...
	//   return !true
	// })
	// 最后的 ans := Search(...) - 1
	// 如果 f 有副作用，需要在 Search 后调用下 f(ans)

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

	_ = []interface{}{reverse, binarySearchF, ternarySearch}
}
