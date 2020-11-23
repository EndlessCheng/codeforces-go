package copypasta

import (
	"bytes"
	"math"
	"sort"
)

/*
sort.Ints 性能测试 https://codeforces.com/contest/977/submission/75301978

https://oeis.org/A001768 Sorting numbers: number of comparisons for merge insertion sort of n elements
https://oeis.org/A001855 Sorting numbers: maximal number of comparisons for sorting n elements by binary insertion
https://oeis.org/A003071 Sorting numbers: maximal number of comparisons for sorting n elements by list merging
https://oeis.org/A036604 Sorting numbers: minimal number of comparisons needed to sort n elements

《挑战》3.1 节练习题
3258 https://www.luogu.com.cn/problem/P2855 二分最小值
3273 https://www.luogu.com.cn/problem/P2884 二分最大值
3104 https://codeforces.com/gym/101649 D http://poj.org/problem?id=3104 二分答案，判断条件是 Σmax(0,(ai-t)/k)<=t
3045 https://www.luogu.com.cn/problem/P1842 贪心，按 s+w 排序
2976 http://poj.org/problem?id=2976 0-1 分数规划
3111 https://codeforces.com/gym/101649 K http://poj.org/problem?id=3111 0-1 分数规划
3579 http://poj.org/problem?id=3579 排序后二分答案
3685 http://poj.org/problem?id=3685 斜着二分可以保证但单调性
2010 https://www.luogu.com.cn/problem/P4952 https://www.luogu.com.cn/problem/P3963
	算法一：排序后二分中位数，copy 数组两侧搞 nthElement
	算法二：排序后用两个堆来维护前缀最小 k 个元素和，以及后缀最小 k 个元素和，然后枚举中位数
3662 https://www.luogu.com.cn/problem/P1948 二分答案，判断条件是 0-1 最短路 <=k
1759 http://poj.org/problem?id=1759 递推式变形成差分，这样可以二分 B，判断最小值是否非负
3484 http://poj.org/problem?id=3484 见「防线」

隐藏的二分 https://codeforces.com/problemset/problem/1354/D

*/

// 有些 OJ 的 Go 版本过低，不支持 sort.Slice，只能用 sort.Sort
type _pair struct{ x, y int }
type pairs []_pair

func (p pairs) Len() int           { return len(p) }
func (p pairs) Less(i, j int) bool { a, b := p[i], p[j]; return a.x < b.x || a.x == b.x && a.y < b.y }
func (p pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sortCollections() {
	{
		var a []int
		// 判断是否为非降序列
		sort.IntsAreSorted(a)
		// 判断是否为严格递增序列
		sort.SliceIsSorted(a, func(i, j int) bool { return a[i] <= a[j] })
	}

	// 插入排序
	// 相关题目 LC1536 https://leetcode-cn.com/contest/weekly-contest-200/problems/minimum-swaps-to-arrange-a-binary-grid/
	insertionSort := func(a []int) {
		n := len(a)
		for i := 1; i < n; i++ {
			v := a[i]
			j := i // 也可以用二分求出循环终点从而减少比较次数
			for ; j > 0 && a[j-1] > v; j-- {
				a[j] = a[j-1]
			}
			a[j] = v
		}
	}

	// sort.SearchInts 的使用技巧
	lowerBound := sort.SearchInts
	upperBound := func(a []int, x int) int { return sort.SearchInts(a, x+1) }
	upperBound = func(a []int, x int) int { return sort.Search(len(a), func(i int) bool { return a[i] > x }) }
	// lowerBound-1 为 <x 的最大值的下标（-1 表示不存在），存在多个最大值时下标取最大的
	// upperBound-1 为 <=x 的最大值的下标（-1 表示不存在），存在多个最大值时下标取最大的

	// sort.Search 的使用技巧
	// 由于 sort.Search 处理的 f(x) 是小 false 大 true 的，在遇到小 true 大 false 的 f(x) 时
	// 若目标是找到最大的使 f(x) == true 的 x
	// 可以考虑二分 !f(x)，则二分结果是最小的使 f(x) == false 的 x，将其 -1 就得到了最大的使 f(x) == true 的 x
	// 由于要对结果 -1，sort.Search 传入的上界需要 +1
	// 好题 https://atcoder.jp/contests/abc149/tasks/abc149_e

	// 指定上下界 [l,r)
	searchRange := func(l, r int, f func(int) bool) int {
		for l < r {
			m := (l + r) >> 1 // 注意 l+r 是否超 int，必要时使用 int(uint(i+j) >> 1) 来转换
			if f(m) {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}
	// 若 l 非负，也可以这样写
	{
		var l, r int
		sort.Search(r, func(x int) bool {
			if x < l {
				return false
			}
			// do ...

			return true
		})
	}

	//

	// 字符串二分 · 其一
	// 字符串有固定长度 n，二分范围从 "aaa...a" 到 "zzz...z"
	binarySearchS1 := func(n int) []byte {
		up := 1 // int64
		for i := 0; i < n; i++ {
			up *= 26
		}
		kthString := func(k int) []byte {
			s := bytes.Repeat([]byte{'a'}, n)
			for i := n - 1; i >= 0 && k > 0; i-- {
				s[i] += byte(k % 26)
				k /= 26
			}
			return s
		}
		kth := sort.Search(up, func(k int) bool {
			s := kthString(k)
			_ = s
			// do s ...

			return true
		})
		return kthString(kth)
	}

	// 字符串二分 · 其二
	// 字符串长度不固定，最长为 n，二分范围从 "a", "b" 到 "zzz...z"
	binarySearchS2 := func(n int) []byte {
		up := 1 // int64
		for i := 0; i < n; i++ {
			up *= 26
		}
		up = (up - 1) / 25 * 26
		kthString := func(k int) []byte {
			s := []byte{}
			for k++; k > 0; k /= 26 {
				k--
				s = append(s, byte('a'+k%26))
			}
			for i, n := 0, len(s); i < n/2; i++ {
				s[i], s[n-1-i] = s[n-1-i], s[i]
			}
			return s
		}
		kth := sort.Search(up, func(k int) bool {
			s := kthString(k)
			_ = s
			// do s ...

			return true
		})
		return kthString(kth)
	}

	//

	// TIPS: 二分三分中的 step 取多少合适：
	// 如果返回结果不是答案的话，注意误差对答案的影响（由于误差累加的缘故，某些题目误差对答案的影响可以达到 n=2e5 倍，见 https://codeforces.com/problemset/problem/578/C）
	// NOTE: l 和 r 最好稍微往左右取宽点，从而保证触发相关逻辑
	// 见 https://codeforces.com/edu/course/2/lesson/6/3/practice/contest/285083/problem/D

	// 实数二分
	// 最大化平均值 https://codeforces.com/edu/course/2/lesson/6/4/practice/contest/285069/problem/A
	binarySearchF := func(l, r float64, f func(x float64) bool) float64 {
		step := int(math.Log2((r - l) / eps)) // eps 取 1e-8 比较稳妥（一般来说是保留小数位+2）
		for ; step > 0; step-- {
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
	// NOTE: 多个下凸函数的乘积仍然是下凸函数；上凸同理 ABC130F
	// https://codeforces.com/blog/entry/60702
	// 模板题 https://www.luogu.com.cn/problem/P3382
	// 题目推荐 https://cp-algorithms.com/num_methods/ternary_search.html#toc-tgt-4
	ternarySearchF := func(l, r float64, f func(x float64) float64) float64 {
		step := int(math.Log((r-l)/eps) / math.Log(1.5)) // eps 取 1e-8 比较稳妥（一般来说是保留小数位+2）
		for ; step > 0; step-- {
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
	// https://codeforces.com/problemset/problem/1301/B (只是举例，不用三分也可做)
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

	//

	// 0-1 分数规划
	// https://oi-wiki.org/misc/frac-programming/
	// https://www.luogu.com.cn/blog/yestoday/post-01-fen-shuo-gui-hua-yang-xie
	// 模板题 https://codeforces.com/edu/course/2/lesson/6/4/practice/contest/285069/problem/C http://poj.org/problem?id=2976
	//       https://codeforces.com/gym/101649 K
	//       https://www.luogu.com.cn/problem/P1570
	// 连续子段的算数平均值 https://codeforces.com/edu/course/2/lesson/6/4/practice/contest/285069/problem/A https://codeforces.com/problemset/problem/1003/C
	// 与 0-1 背包结合，即最优比率背包 https://www.luogu.com.cn/problem/P4377 https://ac.nowcoder.com/acm/contest/2271/F
	// 与生成树结合，即最优比率生成树 https://www.luogu.com.cn/problem/P4951 http://poj.org/problem?id=2728
	// 与负环判定结合，即最优比率环 https://www.luogu.com.cn/problem/P1768 https://www.luogu.com.cn/problem/P2868 https://www.luogu.com.cn/problem/P3199 http://poj.org/problem?id=3621
	//     O(nm) https://www.luogu.com.cn/blog/rqy/solution-p3199
	// 与网络流结合，即最大密度子图 https://www.luogu.com.cn/problem/UVA1389 http://poj.org/problem?id=3155
	// 与费用流结合，即最优比率流 https://www.luogu.com.cn/problem/P3705
	search01 := func(ps [][2]int, k int) float64 {
		// 必须/至少选 k 对，最大化 ∑ai/∑bi
		// 如果是算术平均值的话，bi=1
		n := len(ps)
		const eps = 1e-8
		f := func(rate float64) bool {
			a := make([]float64, n)
			for i, p := range ps {
				a[i] = float64(p[0]) - rate*float64(p[1])
			}
			sort.Float64s(a) // 由于只需要求最大的 k 个数，也可以用 nthElement
			s := 0.0
			for _, v := range a[n-k:] {
				s += v
			}
			return s < 0
		}
		l, r := 0.0, 1e5 // r=max{ai}/min{bi}   也就是根据 ∑ai/∑bi 算出下界和上界，最好松一点
		for step := int(math.Log2((r - l) / eps)); step > 0; step-- {
			mid := (l + r) / 2
			if f(mid) {
				r = mid
			} else {
				l = mid
			}
		}
		return (l + r) / 2
	}

	// CDQ 分治
	// todo https://oi-wiki.org/misc/cdq-divide/
	//      https://www.bilibili.com/video/BV1mC4y1s7ic
	//      [学习笔记]CDQ分治和整体二分 https://www.luogu.com.cn/blog/Owencodeisking/post-xue-xi-bi-ji-cdq-fen-zhi-hu-zheng-ti-er-fen
	//      https://www.luogu.com.cn/blog/ljc20020730/cdq-fen-zhi-xue-xi-bi-ji
	//      动态逆序对 https://www.luogu.com.cn/problem/P3157 https://www.luogu.com.cn/problem/UVA11990

	// 整体二分
	// todo https://oi-wiki.org/misc/parallel-binsearch/
	//      https://codeforces.com/blog/entry/45578

	// WQS 二分
	// todo https://www.cnblogs.com/CreeperLKF/p/9045491.html
	// todo https://www.luogu.com.cn/blog/juruoforever/wqs-er-fen-qian-xi
	// todo https://taodaling.github.io/blog/2020/07/31/WQS%E4%BA%8C%E5%88%86/

	_ = []interface{}{
		insertionSort,
		lowerBound, upperBound,
		searchRange,
		binarySearchS1, binarySearchS2, binarySearchF, ternarySearchF, ternarySearchInt,
		search01,
	}
}
