package copypasta

import (
	"math"
	"sort"
	"strings"
)

// sort.Ints 性能测试 https://codeforces.ml/contest/977/submission/75301978

// 有些 OJ 不支持 sort.Slice，只能用 sort.Sort
type _pair struct{ x, y int }
type pairSlice []_pair

func (p pairSlice) Len() int { return len(p) }
func (p pairSlice) Less(i, j int) bool {
	a, b := p[i], p[j]
	return a.x < b.x || a.x == b.x && a.y < b.y
}
func (p pairSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func sortCollections() {
	sortString := func(s string) string {
		// 可以转成 []byte，也可以……
		a := strings.Split(s, "")
		sort.Strings(a)
		return strings.Join(a, "")
	}
	reverseSort := func(a []int) {
		sort.Sort(sort.Reverse(sort.IntSlice(a)))
	}

	// lowerBound-1 为 <x 的最大值的下标（-1 表示不存在），存在多个最大值时下标取最大的
	// upperBound-1 为 <=x 的最大值的下标（-1 表示不存在），存在多个最大值时下标取最大的
	lowerBound := sort.SearchInts
	upperBound := func(a []int, x int) int { return sort.Search(len(a), func(i int) bool { return a[i] > x }) }

	// NOTE: Pass n+1 if you wanna search range [0,n]
	// NOTE: 二分时特判下限！（例如 0）
	// TIPS: 如果输出的不是二分值而是一个与之相关的值，可以在 return false/true 前记录该值

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

	searchRange := func(l, r int, f func(int) bool) int {
		for l < r {
			m := (l + r) >> 1 // 注意 l+r 是否超 int
			if f(m) {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
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

	// NOTE: step 取多少合适：
	// 如果返回结果不是答案的话，注意误差对答案的影响（由于误差累加的缘故，某些题目误差对查案的影响可以达到 n=2e5 倍，见 CF578C）

	binarySearch := func(l, r float64, f func(x float64) bool) float64 {
		step := int(math.Log2((r - l) / eps)) // eps 取 1e-8 比较稳妥（一般来说是保留小数位+2）
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

	// 实数三分
	// https://codeforces.com/blog/entry/60702
	// 模板题 https://www.luogu.com.cn/problem/P3382
	// 题目推荐 https://cp-algorithms.com/num_methods/ternary_search.html#toc-tgt-4
	ternarySearch := func(l, r float64, f func(x float64) float64) float64 {
		step := int(math.Log((r-l)/eps) / math.Log(1.5)) // eps 取 1e-8 比较稳妥（一般来说是保留小数位+2）
		for i := 0; i < step; i++ {
			m1 := l + (r-l)/3
			m2 := r - (r-l)/3
			v1, v2 := f(m1), f(m2)
			if v1 < v2 {
				r = m2 // 若求最大值写成 l = m1
			} else {
				l = m1 // 若求最大值写成 r = m2
			}
		}
		return (l + r) / 2
	}

	// 整数三分
	// NOTE: 若有大量相同的离散点，该方法在某些数据下会失效（例如三分的时候把存在最小值的「洼地」 skip 了）
	// https://codeforces.ml/problemset/problem/1301/B (只是举例，不用三分也可做)
	ternarySearchInt := func(l, r int, f func(x int) int) int {
		for r-l > 4 { // 最小区间长度根据题目可以扩大点
			m1 := l + (r-l)/3
			m2 := r - (r-l)/3
			v1, v2 := f(m1), f(m2)
			if v1 < v2 {
				r = m2 // 若求最大值写成 l = m1
			} else {
				l = m1 // 若求最大值写成 r = m2
			}
		}
		min, minI := f(l), l
		for i := l + 1; i <= r; i++ {
			if v := f(i); v < min {
				min, minI = v, i
			}
		}
		return minI
	}

	// 整体二分
	// todo https://oi-wiki.org/misc/parallel-binsearch/
	// todo https://codeforces.com/blog/entry/45578

	_ = []interface{}{
		sortString, reverseSort, lowerBound, upperBound, reverse,
		searchRange, search64, binarySearch, ternarySearch, ternarySearchInt,
	}
}
