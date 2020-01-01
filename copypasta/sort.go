package copypasta

import (
	"math"
	"sort"
)

func sortCollections() {
	_ = sort.Search
	// NOTE: Pass n+1 if you wanna search range [0,n]
	// NOTE: 二分时特判下限！（例如 0）

	type bsFunc func(int) bool
	reverse := func(f bsFunc) bsFunc {
		return func(x int) bool {
			return !f(x)
		}
	}
	// 写法1: sort.Search(n, reverse(func(x int) bool {...}))
	// 写法2:
	// sort.Search(n, func(x int) (ok bool) {
	//	 defer func() { ok = !ok }()
	//	 ...
	// })
	// 写法3（推荐）:
	// sort.Search(n, func(x int) (ok bool) {
	//	 ...
	//   return !true
	// })
	// 最后的 ans := Search(...) - 1
	// 如果 f 有副作用，需要在 Search 后调用下 f(ans)
	// 也可以理解成 return false 就是抬高下限，反之减小上限

	// 经常遇到需要从 1 开始二分的情况……
	searchRange := func(l, r int, f func(int) bool) int {
		// Define f(l-1) == false and f(r) == true.
		i, j := l, r
		for i < j {
			h := (i + j) >> 1
			if f(h) {
				j = h
			} else {
				i = h + 1
			}
		}
		return i
	}
	// ……当然，这种情况也可以这样写
	//sort.Search(r, func(x int) bool {
	//	if x < l {
	//		return false
	//	}
	//	...
	//})

	search64 := func(n int64, f func(int64) bool) int64 {
		// Define f(-1) == false and f(n) == true.
		i, j := int64(0), n
		for i < j {
			h := (i + j) >> 1
			if f(h) {
				j = h
			} else {
				i = h + 1
			}
		}
		return i
	}

	binarySearch := func(l, r float64, f func(x float64) bool) float64 {
		step := int(math.Log2((r - l) / eps)) // eps 取 1e-8 比较稳妥
		for i := 0; i < step; i++ {
			mid := (l + r) / 2
			if f(mid) {
				r = mid // 减小 x
			} else {
				l = mid // 增大 x
			}
		}
		return (l + r) / 2
	}

	// 题目推荐 https://cp-algorithms.com/num_methods/ternary_search.html#toc-tgt-4
	ternarySearch := func(l, r float64, f func(x float64) float64) float64 {
		step := int(math.Log((r-l)/eps) / math.Log(1.5)) // eps 取 1e-8 比较稳妥
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

	// TODO: 整体二分 https://oi-wiki.org/misc/parallel-binsearch/

	// TODO: https://oi-wiki.org/search/dlx/

	_ = []interface{}{reverse, searchRange, search64, binarySearch, ternarySearch}
}
