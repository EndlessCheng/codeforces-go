package copypasta

import (
	"bytes"
	"math"
	"slices"
	"sort"
)

/*
https://en.algorithmica.org/hpc/data-structures/binary-search/
BFPRT https://en.wikipedia.org/wiki/Median_of_medians
sort.Ints 性能测试 https://codeforces.com/contest/977/submission/75301978
打造 Go 语言最快的排序算法 https://blog.csdn.net/ByteDanceTech/article/details/124464192

LC853 https://leetcode.cn/problems/car-fleet/
自定义排序 LC1366 https://leetcode.cn/problems/rank-teams-by-votes/ 1626

### 二分查找·题单
- [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/)
- [35. 搜索插入位置](https://leetcode.cn/problems/search-insert-position/)
- [704. 二分查找](https://leetcode.cn/problems/binary-search/)
- [2529. 正整数和负整数的最大计数](https://leetcode.cn/problems/maximum-count-of-positive-integer-and-negative-integer/) 1196 *应用
- [2300. 咒语和药水的成功对数](https://leetcode.cn/problems/successful-pairs-of-spells-and-potions/) 1477 *应用
- [278. 第一个错误的版本](https://leetcode.cn/problems/first-bad-version/)
- [162. 寻找峰值](https://leetcode.cn/problems/find-peak-element/)
- [1901. 寻找峰值 II](https://leetcode.cn/problems/find-a-peak-element-ii/)
- [153. 寻找旋转排序数组中的最小值](https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/)
- [33. 搜索旋转排序数组](https://leetcode.cn/problems/search-in-rotated-sorted-array/)
- [540. 有序数组中的单一元素](https://leetcode.cn/problems/single-element-in-a-sorted-array/)
- [702. 搜索长度未知的有序数组](https://leetcode.cn/problems/search-in-a-sorted-array-of-unknown-size/)（会员题）
- [2936. 包含相等值数字块的数量](https://leetcode.cn/problems/number-of-equal-numbers-blocks/)（会员题）
https://codeforces.com/problemset/problem/600/B 1300
https://codeforces.com/problemset/problem/165/B 1500
https://atcoder.jp/contests/abc248/tasks/abc248_d

### 二分答案·原理

为什么二分的结果一定就是我们要求的，可不可能无法由数组中的元素组合得到？
我在 https://leetcode.cn/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/solution/san-chong-suan-fa-bao-li-er-fen-da-an-du-k1vd/ 中说到：
设答案为 s，那么必然有 f(s−1)<k 且 f(s)≥k。注意这和「第 k 小」是等价的。

### 二分答案 · 题单（右边数字为难度分）

#### 二分答案
- [275. H 指数 II](https://leetcode.cn/problems/h-index-ii/) *经典题
- [1283. 使结果不超过阈值的最小除数](https://leetcode.cn/problems/find-the-smallest-divisor-given-a-threshold/) 1542
- [2187. 完成旅途的最少时间](https://leetcode.cn/problems/minimum-time-to-complete-trips/) 1641 *典型题
- [2226. 每个小孩最多能分到多少糖果](https://leetcode.cn/problems/maximum-candies-allocated-to-k-children/) 1646
- [1870. 准时到达的列车最小时速](https://leetcode.cn/problems/minimum-speed-to-arrive-on-time/) 1676
- [1011. 在 D 天内送达包裹的能力](https://leetcode.cn/problems/capacity-to-ship-packages-within-d-days/) 1725
- [875. 爱吃香蕉的珂珂](https://leetcode.cn/problems/koko-eating-bananas/) 1766 *经典题
- [1898. 可移除字符的最大数目](https://leetcode.cn/problems/maximum-number-of-removable-characters/) 1913
- [1482. 制作 m 束花所需的最少天数](https://leetcode.cn/problems/minimum-number-of-days-to-make-m-bouquets/) 1946
- [1642. 可以到达的最远建筑](https://leetcode.cn/problems/furthest-building-you-can-reach/) 1962
- [2861. 最大合金数](https://leetcode.cn/problems/maximum-number-of-alloys/) 1981
- [3007. 价值和小于等于 K 的最大数字](https://leetcode.cn/problems/maximum-number-that-sum-of-the-prices-is-less-than-or-equal-to-k/) 2258
- [2258. 逃离火灾](https://leetcode.cn/problems/escape-the-spreading-fire/) 2347
- [2137. 通过倒水操作让所有的水桶所含水量相等](https://leetcode.cn/problems/pour-water-between-buckets-to-make-water-levels-equal/)（会员题）
- [2604. 吃掉所有谷子的最短时间](https://leetcode.cn/problems/minimum-time-to-eat-all-grains/)（会员题）
- [2702. 使数字变为非正数的最小操作次数](https://leetcode.cn/problems/minimum-operations-to-make-numbers-non-positive/)（会员题）
https://codeforces.com/contest/1701/problem/C 1400
https://codeforces.com/contest/670/problem/D2 1500
https://codeforces.com/problemset/problem/1610/C 1600
https://codeforces.com/contest/1843/problem/E 1600
https://codeforces.com/problemset/problem/1118/D2 1700
DP https://codeforces.com/contest/883/problem/I 1900

#### 最小化最大值（二分最大值 mx，如果满足要求，例如所有元素最后都 <= mx 则返回 true，否则返回 false，也就是满足要求就让 right 变小，不满足要求就让 left 变大）
- [410. 分割数组的最大值](https://leetcode.cn/problems/split-array-largest-sum/)
- [2064. 分配给商店的最多商品的最小值](https://leetcode.cn/problems/minimized-maximum-of-products-distributed-to-any-store/) 1886
- [1760. 袋子里最少数目的球](https://leetcode.cn/problems/minimum-limit-of-balls-in-a-bag/) 1940
- [1631. 最小体力消耗路径](https://leetcode.cn/problems/path-with-minimum-effort/) 1948
- [2439. 最小化数组中的最大值](https://leetcode.cn/problems/minimize-maximum-of-array/) 1965
- [2560. 打家劫舍 IV](https://leetcode.cn/problems/house-robber-iv/) 2081
- [778. 水位上升的泳池中游泳](https://leetcode.cn/problems/swim-in-rising-water/) 2097 *相当于最小化路径最大值
- [2616. 最小化数对的最大差值](https://leetcode.cn/problems/minimize-the-maximum-difference-of-pairs/) 2155
- [2513. 最小化两个数组中的最大值](https://leetcode.cn/problems/minimize-the-maximum-of-two-arrays/) 2302
https://www.lanqiao.cn/problems/5129/learning/?contest_id=144

#### 最大化最小值（二分最小值 mn+1，如果满足要求，例如所有元素最后都 >= mn+1 则返回 false，否则返回 true，为什么要这样返回请看下面的【sort.Search 的使用技巧·其一】）
- [1552. 两球之间的磁力](https://leetcode.cn/problems/magnetic-force-between-two-balls/) 1920
   同一题 [2517. 礼盒的最大甜蜜度](https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/) 2021
- [2861. 最大合金数](https://leetcode.cn/problems/maximum-number-of-alloys/) 1981
- [2812. 找出最安全路径](https://leetcode.cn/problems/find-the-safest-path-in-a-grid/) 2154
- [2528. 最大化城市的最小供电站数目](https://leetcode.cn/problems/maximize-the-minimum-powered-city/) 2236
- [LCP 12. 小张刷题计划](https://leetcode.cn/problems/xiao-zhang-shua-ti-ji-hua/)
- [1231. 分享巧克力](https://leetcode.cn/problems/divide-chocolate/)（会员题）2029
http://codeforces.com/problemset/problem/460/C 1700
https://codeforces.com/problemset/problem/1550/E 2500

#### 第 K 小/大（部分题目也可以用堆解决）
- [378. 有序矩阵中第 K 小的元素](https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix/)
- [668. 乘法表中第 K 小的数](https://leetcode.cn/problems/kth-smallest-number-in-multiplication-table/)
- [373. 查找和最小的 K 对数字](https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/)
- [719. 找出第 K 小的数对距离](https://leetcode.cn/problems/find-k-th-smallest-pair-distance/)
- [1439. 有序矩阵中的第 k 个最小数组和](https://leetcode.cn/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/) 2134
- [786. 第 K 个最小的素数分数](https://leetcode.cn/problems/k-th-smallest-prime-fraction/) 2169
- [2040. 两个有序数组的第 K 小乘积](https://leetcode.cn/problems/kth-smallest-product-of-two-sorted-arrays/) 2518
- [2386. 找出数组的第 K 大和](https://leetcode.cn/problems/find-the-k-sum-of-an-array/) 2648
- [1918. 第 K 小的子数组和](https://leetcode.cn/problems/kth-smallest-subarray-sum/)（会员题）*滑动窗口
综合 https://atcoder.jp/contests/abc155/tasks/abc155_d

#### 最大化中位数
https://codeforces.com/problemset/problem/1201/C  也可以贪心做

#### 不好想到的二分（这也能二分？！）
https://codeforces.com/problemset/problem/1707/A

《挑战》3.1 节练习题
3258 https://www.luogu.com.cn/problem/P2855 二分最小值
3273 https://www.luogu.com.cn/problem/P2884 二分最大值
3104 https://codeforces.com/gym/101649 D http://poj.org/problem?id=3104 二分答案，判断条件是 ∑max(0,(ai-t)/k)<=t
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
3484 https://www.acwing.com/problem/content/122/ 二分位置

https://codeforces.com/problemset/problem/1697/D
隐藏的二分 https://atcoder.jp/contests/abc203/tasks/abc203_d
隐藏的二分 https://codeforces.com/problemset/problem/1354/D
转换的好题 https://codeforces.com/problemset/problem/1181/D

第 k 小子序列和 https://codeforces.com/gym/101234/problem/G https://leetcode.cn/problems/find-the-k-sum-of-an-array/
- 思路见我的题解 https://leetcode.cn/problems/find-the-k-sum-of-an-array/solution/zhuan-huan-dui-by-endlesscheng-8yiq/

https://oeis.org/A001768 Sorting numbers: number of comparisons for merge insertion sort of n elements
https://oeis.org/A001855 Sorting numbers: maximal number of comparisons for sorting n elements by binary insertion
https://oeis.org/A003071 Sorting numbers: maximal number of comparisons for sorting n elements by list merging
https://oeis.org/A036604 Sorting numbers: minimal number of comparisons needed to sort n elements
*/

// 把两个数组绑起来排序
// 使用方法：sort.Sort(zip{a, b})
type zip struct {
	a []int
	b []int
}

//func (p zip) Less(i, j int) bool { return p.a[i] < p.a[j] || p.a[i] == p.a[j] && p.b[i] < p.b[j] }
func (p zip) Less(i, j int) bool { return p.a[i] < p.a[j] }
func (p zip) Len() int           { return len(p.a) }
func (p zip) Swap(i, j int) {
	p.a[i], p.a[j] = p.a[j], p.a[i]
	p.b[i], p.b[j] = p.b[j], p.b[i]
}

// 记录排序过程中交换元素的下标
// r := swapRecorder{a, &[][2]int{}}
// sort.Sort(r)
// 相关题目 https://codeforces.com/problemset/problem/266/C
type swapRecorder struct {
	sort.IntSlice
	swaps *[][2]int
}

func (r swapRecorder) Swap(i, j int) {
	// 快排时可能会有 i 和 j 相等的情况
	if i == j {
		return
	}
	*r.swaps = append(*r.swaps, [2]int{i, j})
	r.IntSlice[i], r.IntSlice[j] = r.IntSlice[j], r.IntSlice[i]
}

func sortCollections() {
	{
		var a []int

		// 判断是否为非降序列
		sort.IntsAreSorted(a)

		// 判断是否为严格递增序列
		sort.SliceIsSorted(a, func(i, j int) bool { return a[i] <= a[j] })

		// 判断是否为非增序列
		sort.IsSorted(sort.Reverse(sort.IntSlice(a)))
	}

	// 在多个左闭右开区间中，查找与 [l,r) 有交集的所有区间
	// https://codeforces.com/problemset/problem/1817/A
	type interval struct{ l, r int }
	searchIntervals := func(a []interval, l, r int) {
		li := sort.Search(len(a), func(i int) bool { return a[i].r > l })
		if li < len(a) && a[li].l < r { // 至少有一个区间
			ri := sort.Search(len(a), func(i int) bool { return a[i].l >= r }) - 1
			leftL := max(l, a[li].l)
			rightR := min(r, a[ri].r)
			if li == ri { // 只有一个区间 [leftL, rightR)
				_ = rightR - leftL

			} else { // 多个区间
				// [leftL, a[li].r) + ... + [a[ri].l, rightR)
				midFull := ri - li - 1
				_ = midFull

			}
		}
	}

	// 把数组排序（元素互不相同），需要的最小交换次数
	// 做法：离散化后求置换环
	// LC2471 https://leetcode.cn/problems/minimum-number-of-operations-to-sort-a-binary-tree-by-level/
	minSwaps := func(a []int) int {
		id := make([]int, len(a))
		for i := range id {
			id[i] = i
		}
		sort.Slice(id, func(i, j int) bool { return a[id[i]] < a[id[j]] }) // 简单离散化

		ans := len(a)
		for i, v := range id {
			if v < 0 {
				continue
			}
			for id[i] >= 0 {
				nxt := id[i]
				id[i] = -1
				i = nxt
			}
			ans--
		}
		return ans
	}

	// 插入排序
	// 相关题目 LC1536 https://leetcode-cn.com/problems/minimum-swaps-to-arrange-a-binary-grid/
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
	// 等于 x 的下标范围：[lowerBound(x), upperBound(x))
	// lowerBound-1 为 <x 的最大值的下标（-1 表示不存在），存在多个最大值时下标取最大的
	// upperBound-1 为 <=x 的最大值的下标（-1 表示不存在），存在多个最大值时下标取最大的

	// sort.Search 的使用技巧·其一
	// sort.Search(n, f) 需要满足当 x 从小到大时，f(x) 先 false 后 true
	// 若 f(x) 是先 true 后 false，且目标是找到最大的使 f(x) 为 true 的 x
	// 这种情况可以考虑二分 !f(x)，则二分结果是最小的使 f(x) 为 false 的 x，将其 -1 就得到了最大的使 f(x) 为 true 的 x
	// 由于要对结果 -1，sort.Search 传入的上界需要 +1
	// 更加简单的写法是，在 f(x) 内部将 x++，这样就不需要对上界和结果调整 ±1 了
	//
	// 下面以二分求 int(sqrt(90)) 为例来说明这一技巧
	// 这相当于求最大的满足 x*x<=90 的 x
	// 于是定义 f(x) 返回 x*x<=90，注意这是一个先 true 后 false 的 f(x)
	// 我们可以改为判断 f(x+1)，即用 f(x+1) 的返回结果代替 f(x) 的返回结果
	// 同时，将 f(x) 改为先 false 后 true，即返回 x*x>90
	// 这样二分的结果就恰好停在最大的满足原 f(x) 为 true 的 x 上
	sort.Search(10, func(x int) bool {
		x++
		return x*x > 90
	})

	// 当然，这种求最大值的二分也可以用下面这种左开右闭的写法（参考 sort.Search 源码）
	search2 := func(n int, f func(int) bool) int {
		// Define f(0) == true and f(n+1) == false.
		// Invariant: f(l) == true, f(r+1) == false.
		// 这样定义的好处见下面 return 前的注释
		l, r := 0, n
		for l < r {
			mid := int(uint(l+r+1) >> 1) // mid=⌈(l+r)/2⌉，从而保证 mid 落在区间 (l,r] 内
			// l < mid ≤ r
			if f(mid) {
				l = mid // preserves f(l) == true
			} else {
				r = mid - 1 // preserves f(r+1) == false
			}
		}
		// l == r, f(r+1) == false, and f(l) (= f(r)) == true  =>  answer is l.
		return l
	}

	// 好题 https://atcoder.jp/contests/abc149/tasks/abc149_e

	// sort.Search 的使用技巧·其二
	// 若要求出一个和二分结果相关的东西
	// 可以在返回值为 true 时记录下相关数据（若有多个地方返回 true，可以用 defer 来简化）
	// 这样可以避免在二分结束后再计算一次
	// 为了保证能至少触发一次 true，某些情况下需要将二分上界 +1
	// https://codeforces.com/problemset/problem/1100/E

	// 指定上下界 [l,r)
	searchRange := func(l, r int) int {
		return l + sort.Search(r-l, func(x int) bool {
			x += l
			// ...

			return false
		})
	}

	searchRange64 := func(l, r int64, f func(int64) bool) int64 {
		for l < r {
			m := (l + r) >> 1 // l + (r-l)>>1
			if f(m) {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}

	//

	// 字符串二分 · 其一
	// 字符串有固定长度 n，二分范围从 "aaa...a" 到 "zzz...z"
	binarySearchS1 := func(n int) []byte {
		up := 1
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
	// 字符串长度不固定，最长为 L，二分范围从 "a", "b" 到 "zzz...z"
	binarySearchS2 := func(L int) []byte {
		up := 1
		for i := 0; i < L; i++ {
			up *= 26
		}
		up = (up - 1) / 25 * 26 // 求字符串的个数（等比数列之和 26 + 26*26 + ... + 26^L）
		kthString := func(k int) []byte {
			s := []byte{}
			for k++; k > 0; k /= 26 {
				k--
				s = append(s, byte('a'+k%26))
			}
			slices.Reverse(s)
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

	// 有序矩阵中的第 k 小
	// 有序矩阵：每行和每列元素均为不降序列
	// LC378 https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/
	// LC719 https://leetcode-cn.com/problems/find-k-th-smallest-pair-distance/
	kthSmallest := func(a [][]int, k int) int {
		// 注意 k 从 1 开始
		n, m := len(a), len(a[0])
		mi, mx := a[0][0], a[n-1][m-1]
		ans := sort.Search(mx-mi, func(v int) bool {
			v += mi
			cnt := 0
			for i, j := 0, m-1; i < n && j >= 0; {
				if v < a[i][j] {
					j--
				} else {
					cnt += j + 1
					i++
				}
			}
			return cnt >= k
		}) + mi
		return ans
	}

	// 区间和的第 k 小（数组元素均为非负）
	// 每个区间和可以视作一个有序上三角矩阵中的元素，在数组元素均为非负时，该矩阵从左往右和从下往上均为非降序列
	// 1508 https://leetcode-cn.com/problems/range-sum-of-sorted-subarray-sums/
	kthSmallestRangeSum := func(a []int, k int) int {
		// 1 <= k <= n*(n+1)/2
		n := len(a)
		sum := make([]int, n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + v
		}
		ans := sort.Search(sum[n], func(v int) bool {
			cnt := 0
			for l, r := 0, 1; r <= n; {
				if v < sum[r]-sum[l] {
					l++
				} else {
					cnt += r - l
					r++
				}
			}
			return cnt >= k
		})
		return ans
	}

	// 子集和的第 k 小（数组元素均为非负）
	// k 从 1 开始
	// 除了二分，另一种求法是使用最小堆
	// 初始时插入 (a[0],0)，然后执行 k-1 次操作：取出堆顶，插入 (top.v+a[top.i+1],top.i+1) 以及 (top.v+a[top.i+1]-a[top.i],top.i+1)
	// 代码见 https://codeforces.com/gym/101234/submission/116219928
	// https://codeforces.com/gym/101234/problem/G
	kthSubsetSum := func(a []int, k int) int {
		slices.Sort(a)
		// 上界不会超过 a 的前 log(k) 个元素之和
		ans := sort.Search(2e9, func(sum int) bool {
			c := 0
			var f func(p, s int)
			f = func(p, s int) {
				if c >= k || p == len(a) || s+a[p] > sum {
					return
				}
				c++
				f(p+1, s+a[p])
				f(p+1, s)
			}
			f(0, 0)
			return c >= k
		})
		return ans
	}

	//

	// NOTE: 实数二分/三分中的 step 取多少合适：
	// 如果返回结果不是答案的话，注意误差对答案的影响
	// 由于误差累加的缘故，某些题目误差对答案的影响可以达到 n=2e5 倍，见 https://codeforces.com/problemset/problem/578/C
	// 乘法带来的误差见 https://codeforces.com/problemset/problem/653/D

	// NOTE: l 和 r 最好稍微往左右取宽点，从而保证触发相关逻辑
	// 见 https://codeforces.com/edu/course/2/lesson/6/3/practice/contest/285083/problem/D

	// 实数二分
	// LC2137 https://leetcode.cn/problems/pour-water-between-buckets-to-make-water-levels-equal/
	// 最大化平均值 https://codeforces.com/edu/course/2/lesson/6/4/practice/contest/285069/problem/A
	// 0-1 分数规划见后面
	binarySearchF := func(l, r float64, f func(x float64) bool) float64 {
		// 松一点
		l--
		r++
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
	// 另一种写法是用黄金比率，效率更高
	// NOTE: 多个下凸函数的乘积仍然是下凸函数；上凸同理 ABC130F
	// https://codeforces.com/blog/entry/60702
	// 模板题 https://www.luogu.com.cn/problem/P3382
	// 模板题 https://ac.nowcoder.com/acm/contest/64272/c
	// 题目推荐 https://cp-algorithms.com/num_methods/ternary_search.html#toc-tgt-4
	ternarySearchF := func(l, r float64, f func(x float64) float64) float64 {
		// 松一点
		l--
		r++
		const eps = 1e-8 // 保留小数位+2
		step := int(math.Log((r-l)/eps) / math.Log(1.5))
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
		return (l + r) / 2 // f((l + r) / 2)
	}

	// 整数三分·写法一
	// 比较两个三分点值的大小，每次去掉 1/3 的区间
	// https://codeforces.com/blog/entry/11497
	// https://codeforces.com/blog/entry/43440
	// https://codeforces.com/blog/entry/60702
	// NOTE: 若有大量相同的离散点则可能会失效（例如三分的时候把存在最小值的「洼地」 skip 了）
	// https://codeforces.com/problemset?order=BY_SOLVED_DESC&tags=ternary+search
	// https://codeforces.com/problemset/problem/1355/E
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

	// 整数三分·写法二
	// 二分导数零点（准确说是一阶差分），即比较 f(m) 和 f(m+1), m=(l+r)/2
	// 这种写法的优点是两次运算可以将枚举范围减半，而三分点的写法两次运算仅去掉了 1/3 的范围（效率比 log(2)/log(1.5) ≈ 1.71）
	// 但是，如果存在相邻 f 值相同，且只有两个的情况：f(i-1)<f(i)=f(i+1)<f(i+2)，这种写法将会失效，而三分点的写法保证了两个三分点的间隔，可以正常运行
	ternarySearchInt2 := func(l, r int, f func(x int) int) int {
		return l + sort.Search(r-l, func(m int) bool { return f(l+m) < f(l+m+1) }) // < 求最小值   > 求最大值
	}

	// 整数三分·写法三
	// 黄金比率实现，效率更高

	//

	// 0-1 分数规划
	// 求 min{∑ai/∑bi}：在 check(k) 中判断是否有 min∑(ai-k*bi) > 0 成立，若成立说明 k 取小了，否则 k 取大了
	// 求 max{∑ai/∑bi}：在 check(k) 中判断是否有 max∑(ai-k*bi) > 0 成立，若成立说明 k 取小了，否则 k 取大了
	// https://oi-wiki.org/misc/frac-programming/
	// https://www.luogu.com.cn/blog/yestoday/post-01-fen-shuo-gui-hua-yang-xie
	// 模板题 https://codeforces.com/edu/course/2/lesson/6/4/practice/contest/285069/problem/C http://poj.org/problem?id=2976
	//       https://codeforces.com/gym/101649 K
	//       https://www.luogu.com.cn/problem/P1570
	//       https://loj.ac/p/149
	// 有长度限制的连续子段的（最大/最小）算数平均值
	//     https://codeforces.com/edu/course/2/lesson/6/4/practice/contest/285069/problem/A
	//     https://codeforces.com/problemset/problem/1003/C
	//     https://www.luogu.com.cn/problem/P1404
	//     https://www.acwing.com/problem/content/104/
	//     LC644 https://leetcode-cn.com/problems/maximum-average-subarray-ii/
	//     O(n) 做法见 04 年集训队周源论文《浅谈数形结合思想在信息学竞赛中的应用》（或紫书 p.243 例题 8-9，UVa 1451 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=447&page=show_problem&problem=4197）
	// 与 0-1 背包结合，即最优比率背包 https://www.luogu.com.cn/problem/P4377 https://ac.nowcoder.com/acm/contest/2271/F
	// 与生成树结合，即最优比率生成树 https://www.luogu.com.cn/problem/P4951 http://poj.org/problem?id=2728
	// 与负环判定结合，即最优比率环 https://www.luogu.com.cn/problem/P1768 https://www.luogu.com.cn/problem/P2868 https://www.luogu.com.cn/problem/P3199 http://poj.org/problem?id=3621
	//     O(nm) https://www.luogu.com.cn/blog/rqy/solution-p3199
	// 与网络流结合，即最大密度子图 https://www.luogu.com.cn/problem/UVA1389 http://poj.org/problem?id=3155
	// 与费用流结合，即最优比率流 https://www.luogu.com.cn/problem/P3705
	// 其他的一些题：
	//      与 DP 结合 https://codeforces.com/problemset/problem/489/E
	// 最优比率路径 https://atcoder.jp/contests/abc324/tasks/abc324_f
	//            https://codeforces.com/edu/course/2/lesson/6/4/practice/contest/285069/problem/B
	search01 := func(ps [][2]int, k int) float64 {
		// 必须/至少选 k 对，最大化 ∑ai/∑bi
		// 如果是算术平均值的话，bi=1
		n := len(ps)
		// 稳妥起见，eps 可以设的比要求的精度高两个，如果题目没有给出精度要求（例如求方案），可以将 eps 设为 1/(100∑bi)
		// 注意：若时限比较紧，可以适当调低精度
		const eps = 1e-8
		f := func(rate float64) bool {
			a := make([]float64, n)
			for i, p := range ps {
				a[i] = float64(p[0]) - rate*float64(p[1])
			}
			slices.Sort(a) // 由于只需要求最大的 k 个数，也可以用 nthElement
			s := .0
			for _, v := range a[n-k:] {
				s += v
			}
			return s < 0
		}
		l, r := -1.0, 1e5+1 // r=max{ai}/min{bi}   也就是根据 ∑ai/∑bi 算出下界和上界，最好松一点
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

	// WQS 二分
	// 见 dp.go

	// CDQ 分治
	// 见 dp.go

	// 整体二分 Parallel Binary Search
	// https://oi-wiki.org/misc/parallel-binsearch/
	// https://codeforces.com/blog/entry/45578
	// todo 整体二分解决静态区间第 k 小的优化 https://www.luogu.com.cn/blog/2-6-5-3-5/zheng-ti-er-fen-xie-jue-jing-tai-ou-jian-di-k-xiao-di-you-hua
	// 模板题 https://www.luogu.com.cn/problem/P3527
	// todo https://atcoder.jp/contests/agc002/tasks/agc002_d
	//  https://www.hackerrank.com/contests/hourrank-23/challenges/selective-additions/problem
	//  https://www.codechef.com/problems/MCO16504
	parallelBinarySearch := func(n int, qs []struct{ l, r, v int }) []int {
		// 读入询问时可以处理成左闭右开的形式

		ans := make([]int, n)
		tar := make([]int, n)
		for i := range tar {
			tar[i] = i
		}
		var f func([]int, int, int)
		f = func(tar []int, ql, qr int) {
			if len(tar) == 0 {
				return
			}
			if ql+1 == qr {
				for _, c := range tar {
					ans[c] = ql // qr
				}
				return
			}
			qm := (ql + qr) / 2
			for _, q := range qs[ql:qm] {
				_ = q
				// apply(q)

			}

			// 根据此刻查询的结果将 tar 分成左右两部分
			var left, right []int
			for _, who := range tar {
				_ = who

			}

			for _, q := range qs[ql:qm] {
				_ = q
				// rollback(q)

			}
			f(left, ql, qm)
			f(right, qm, qr)
		}
		f(tar, 0, len(qs)+1) // 这样可以将无法满足要求的 ans[i] 赋值为 len(qs)
		return ans
	}

	// 倍增
	// https://www.acwing.com/problem/content/description/111/
	binaryLifting := func(a []int, check func(a []int) bool) (ans int) {
		n := len(a)
		for r := 0; r < n; { // 注意这里是 <
			l := r
			for sz := 1; sz > 0; {
				if r+sz <= n && check(a[l:r+sz]) {
					r += sz
					sz *= 2
				} else {
					sz /= 2
				}
			}
			ans++
		}
		return
	}

	_ = []interface{}{
		searchIntervals,
		minSwaps,
		insertionSort,
		lowerBound, upperBound, search2,
		searchRange, searchRange64,
		binarySearchS1, binarySearchS2,
		kthSmallest, kthSmallestRangeSum, kthSubsetSum,
		binarySearchF, ternarySearchF, ternarySearchInt, ternarySearchInt2,
		search01,
		parallelBinarySearch,
		binaryLifting,
	}
}
