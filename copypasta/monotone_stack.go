package copypasta

/* 单调栈 Monotone Stack

【图解单调栈】两种方法，两张图秒懂
https://leetcode.cn/problems/next-greater-node-in-linked-list/solution/tu-jie-dan-diao-zhan-liang-chong-fang-fa-v9ab/

举例：返回每个元素两侧严格大于它的元素位置（不存在则为 -1 或 n）
如何理解：把数组想象成一列山峰，站在 a[i] 的山顶仰望两侧的山峰，是看不到高山背后的矮山的，只能看到一座座更高的山峰
         这就启发我们引入一个底大顶小的单调栈，入栈时不断比较栈顶元素直到找到一个比当前元素大的
技巧：事先压入一个边界元素到栈底，这样保证循环时栈一定不会为空，从而简化逻辑
一些转换：
    若区间 [l,r] 的最大值等于 a[r]，则 l 必须 > left[r]
    若区间 [l,r] 的最大值等于 a[l]，则 r 必须 < right[l]
    这一结论可以用于思考一些双变量的题目
https://oi-wiki.org/ds/monotonous-stack/
https://cp-algorithms.com/data_structures/stack_queue_modification.html

#### 单调栈

- [496. 下一个更大元素 I](https://leetcode.cn/problems/next-greater-element-i/)（单调栈模板题）
- [503. 下一个更大元素 II](https://leetcode.cn/problems/next-greater-element-ii/)
- [2454. 下一个更大元素 IV](https://leetcode.cn/problems/next-greater-element-iv/)
- [456. 132 模式](https://leetcode.cn/problems/132-pattern/)
- [739. 每日温度](https://leetcode.cn/problems/daily-temperatures/)
- [901. 股票价格跨度](https://leetcode.cn/problems/online-stock-span/)
- [1019. 链表中的下一个更大节点](https://leetcode.cn/problems/next-greater-node-in-linked-list/)
- [1124. 表现良好的最长时间段](https://leetcode.cn/problems/longest-well-performing-interval/)
- [1475. 商品折扣后的最终价格](https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/)
- [2289. 使数组按非递减顺序排列](https://leetcode.cn/problems/steps-to-make-array-non-decreasing/)

#### 矩形系列

- [84. 柱状图中最大的矩形](https://leetcode.cn/problems/largest-rectangle-in-histogram/)
- [85. 最大矩形](https://leetcode.cn/problems/maximal-rectangle/)
- [1504. 统计全 1 子矩形](https://leetcode.cn/problems/count-submatrices-with-all-ones/)

#### 字典序最小

- [316. 去除重复字母](https://leetcode.cn/problems/remove-duplicate-letters/)
- [316 扩展：重复个数不超过 limit](https://leetcode.cn/contest/tianchi2022/problems/ev2bru/)
- [402. 移掉 K 位数字](https://leetcode.cn/problems/remove-k-digits/)
- [321. 拼接最大数](https://leetcode.cn/problems/create-maximum-number/)

#### 贡献法

- [907. 子数组的最小值之和](https://leetcode.cn/problems/sum-of-subarray-minimums/)
- [1856. 子数组最小乘积的最大值](https://leetcode.cn/problems/maximum-subarray-min-product/)
- [2104. 子数组范围和](https://leetcode.cn/problems/sum-of-subarray-ranges/)
- [2281. 巫师的总力量和](https://leetcode.cn/problems/sum-of-total-strength-of-wizards/)

模板题
https://www.luogu.com.cn/problem/P5788
https://www.luogu.com.cn/problem/P2866 http://poj.org/problem?id=3250

NEERC05，UVa 1619 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=825&page=show_problem&problem=4494
转换 https://codeforces.com/problemset/problem/280/B
转换 LC2289 https://leetcode.cn/problems/steps-to-make-array-non-decreasing/
max >= sum https://codeforces.com/problemset/problem/1691/D
LC1124 https://leetcode.cn/problems/longest-well-performing-interval/
你从单调栈学到了什么思想？LC1944 https://leetcode.cn/problems/number-of-visible-people-in-a-queue/
下下个最大元素 LC2454 https://leetcode.cn/problems/next-greater-element-iv/
- 应用 https://atcoder.jp/contests/abc140/tasks/abc140_e
max(最小值*子数组和) LC1856 https://leetcode.cn/problems/maximum-subarray-min-product/

字典序最小
LC316 https://leetcode.cn/problems/remove-duplicate-letters/
- 扩展：重复个数不超过 limit https://leetcode.cn/contest/tianchi2022/problems/ev2bru/
LC402 https://leetcode.cn/problems/remove-k-digits/
LC321 https://leetcode.cn/problems/create-maximum-number/

计算贡献（所有子数组的……的和）
最小值 LC907 https://leetcode.cn/problems/sum-of-subarray-minimums/
最大值-最小值 LC2104 https://leetcode.cn/problems/sum-of-subarray-ranges/
最小值*和 LC2281 https://leetcode.cn/problems/sum-of-total-strength-of-wizards/
第二大 https://atcoder.jp/contests/abc140/tasks/abc140_e

与 DP 结合
https://codeforces.com/problemset/problem/5/E
https://codeforces.com/problemset/problem/1313/C2
https://codeforces.com/problemset/problem/1407/D
结合线段树，或者巧妙地在单调栈中去维护最值 https://codeforces.com/problemset/problem/1483/C
按照最大值分类讨论 LC1335 https://leetcode.cn/problems/minimum-difficulty-of-a-job-schedule/
LC2355 https://leetcode.cn/problems/maximum-number-of-books-you-can-take/

其他
LC42 接雨水 https://leetcode-cn.com/problems/trapping-rain-water/
     评注：接雨水有三种不同的解法（DP、单调栈和双指针），其中双指针是 DP 的空间优化写法，讲解见 https://www.bilibili.com/video/BV1Qg411q7ia/
          本质上是两种计算策略：计算每个下标处的接水量（纵向累加），计算一段高度对应的接水宽度（横向累加）
LC84 柱状图中最大的矩形 https://leetcode-cn.com/problems/largest-rectangle-in-histogram/ http://poj.org/problem?id=2559 http://poj.org/problem?id=2082
LC85 最大全 1 矩形（实现见下面的 maximalRectangleArea）https://leetcode-cn.com/problems/maximal-rectangle/ 原题为 http://poj.org/problem?id=3494
LC1504 全 1 矩形个数（实现见下面的 numSubmat）https://leetcode-cn.com/problems/count-submatrices-with-all-ones/
LC768 https://leetcode.cn/problems/max-chunks-to-make-sorted-ii/
后缀数组+不同矩形对应方案数之和 https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/D
与 bitOpTrickCnt 结合（见 bits.go）https://codeforces.com/problemset/problem/875/D
已知部分 right 还原全部 right；已知 right 还原 a https://codeforces.com/problemset/problem/1158/C
*/
func monotoneStack(a []int) ([]int, []int) {
	const mod int = 1e9 + 7

	// 考察局部最小
	// 如果有相同元素，需要把某一侧循环内的符号改成小于等于

	// 求左侧严格小于 a[i] 的最近位置 left[i]，这样 a[i] 就是区间 [left[i]+1,i] 内最小的元素（之一）
	// 如果改成求左侧小于等于，那么 a[i] 就是区间 [left[i]+1,i] 内独一无二的最小元素
	// 不存在时 left[i] = -1
	// 虽然写了个二重循环，但站在每个元素的视角看，这个元素在二重循环中最多入栈出栈各一次，因此整个二重循环的时间复杂度为 O(n)
	n := len(a)
	left := make([]int, n)
	st := []int{-1} // 栈底哨兵
	for i, v := range a {
		for len(st) > 1 && a[st[len(st)-1]] >= v { // 不断弹出 >= v 的，循环结束后栈顶就是 < v 的（下面 left[i] = st[len(st)-1]）
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}

	// 求右侧严格小于 a[i] 的最近位置 right[i]，这样 a[i] 就是区间 [i,right[i]-1] 内最小的元素（之一）
	// 如果改成求右侧小于等于，那么 a[i] 就是区间 [i,right[i]-1] 内独一无二的最小元素
	// 不存在时 right[i] = n
	right := make([]int, n)
	st = []int{n}
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		for len(st) > 1 && a[st[len(st)-1]] >= v {
			st = st[:len(st)-1]
		}
		right[i] = st[len(st)-1]
		st = append(st, i)
	}

	sum := make([]int, n+1) // int64
	for i, v := range a {
		sum[i+1] = (sum[i] + v) % mod
	}

	// EXTRA：计算贡献（注意取模时避免出现负数）
	ans := 0
	for i, v := range a {
		l, r := left[i]+1, right[i] // [l,r) 左闭右开
		// ...

		tot := (sum[r] + mod - sum[l]) % mod
		ans = (ans + v*tot) % mod
	}

	{
		// TIPS: 如果有一侧定义成小于等于，还可以一次遍历求出 left 和 right
		left := make([]int, n)
		right := make([]int, n)
		for i := range right {
			right[i] = n
		}
		st := []int{-1}
		for i, v := range a {
			for len(st) > 1 && a[st[len(st)-1]] >= v { // 这里是 right 小于等于
				right[st[len(st)-1]] = i
				st = st[:len(st)-1]
			}
			left[i] = st[len(st)-1]
			st = append(st, i)
		}
	}

	// EXTRA: 求所有长为 i 的子区间的最小值的最大值
	// https://codeforces.com/problemset/problem/547/B LC1950 https://leetcode-cn.com/problems/maximum-of-minimum-values-in-all-subarrays/
	{
		ans := make([]int, n+1)
		for i := range ans {
			ans[i] = -2e9
		}
		for i, v := range a {
			sz := right[i] - left[i] - 1
			if v > ans[sz] {
				ans[sz] = v
			}
		}
		for i := n - 1; i > 0; i-- {
			if ans[i+1] > ans[i] {
				ans[i] = ans[i+1]
			}
		}
		// ans[1:]
	}

	return left, right
}

// 注：若输入的是一个 1~n 的排列，求两侧大于/小于位置有更简单的写法
// 用双向链表思考（代码实现时用的数组）：
// - 把 perm 转换成双向链表，按元素值**从小到大**遍历 perm[i]，那么 perm[i] 左右两侧的就是大于 perm[i] 的元素
// - 算完 perm[i] 后把 perm[i] 从链表中删掉
// 为避免判断下标越界，传入的 perm 虽然下标是从 0 开始的，但视作从 1 开始（不存在时表示为 0 或 n+1）
// https://codeforces.com/contest/1156/problem/E
// https://atcoder.jp/contests/abc140/tasks/abc140_e
func permLR(perm []int) ([]int, []int) {
	n := len(perm)
	pos := make([]int, n+1)
	left := make([]int, n+2)
	right := make([]int, n+1)
	for i := 1; i <= n; i++ {
		pos[perm[i-1]] = i
		left[i], right[i] = i-1, i+1
	}
	right[0] = 1
	left[n+1] = n // 哨兵（本题不需要这两行，但是某些题目需要，比如 https://codeforces.com/problemset/problem/1154/E）
	del := func(i int) {
		l, r := left[i], right[i]
		right[l] = r
		left[r] = l
	}

	// 正序遍历求出的是两侧大于位置
	// 倒序遍历求出的是两侧小于位置
	for v := 1; v <= n; v++ {
		i := pos[v]
		l, r := left[i], right[i]
		// do ...
		_, _ = l, r

		del(i) // 从链表中删除 v
	}
	return left, right
}

// 最大全 1 矩形
// LC85 https://leetcode-cn.com/problems/maximal-rectangle/
func maximalRectangleArea(mat [][]int) (ans int) {
	const target = 1
	n, m := len(mat), len(mat[0])
	heights := make([][]int, n) // heights[i][j] 表示从 (i,j) 往上看的高度（连续 1 的长度），mat[i][j] = 0 时为 0
	for i, row := range mat {
		heights[i] = make([]int, m)
		for j, v := range row {
			if v == target {
				if i == 0 {
					heights[i][j] = 1
				} else {
					heights[i][j] = heights[i-1][j] + 1
				}
			}
		}
	}

	// 然后枚举每一行，就变成 LC84 这题了
	type pair struct{ h, i int }
	for _, hs := range heights {
		left := make([]int, m)
		stack := []pair{{-1, -1}}
		for j, h := range hs {
			for {
				if top := stack[len(stack)-1]; top.h < h {
					left[j] = top.i
					break
				}
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, pair{h, j})
		}
		right := make([]int, m)
		stack = []pair{{-1, m}}
		for j := m - 1; j >= 0; j-- {
			h := hs[j]
			for {
				if top := stack[len(stack)-1]; top.h < h {
					right[j] = top.i
					break
				}
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, pair{h, j})
		}
		for j, h := range hs {
			if area := (right[j] - left[j] - 1) * h; area > ans {
				ans = area
			}
		}
	}
	return
}

// 全 1 矩形个数
// LC1504 https://leetcode-cn.com/problems/count-submatrices-with-all-ones/
// 参考 https://leetcode.com/problems/count-submatrices-with-all-ones/discuss/720265/Java-Detailed-Explanation-From-O(MNM)-to-O(MN)-by-using-Stack
func numSubmat(mat [][]int) (ans int) {
	m := len(mat[0])
	heights := make([]int, m)
	for _, row := range mat {
		sum := make([]int, m)
		type pair struct{ h, j int }
		stack := []pair{{-1, -1}}
		for j, v := range row {
			if v == 0 {
				heights[j] = 0
			} else {
				heights[j]++
			}
			h := heights[j]
			for {
				if top := stack[len(stack)-1]; top.h < h {
					if pre := top.j; pre < 0 {
						sum[j] = (j + 1) * h
					} else {
						sum[j] = sum[pre] + (j-pre)*h
					}
					ans += sum[j]
					break
				}
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, pair{h, j})
		}
	}
	return
}

// 字典序最小的无重复字符的子序列
// LC316 https://leetcode.cn/problems/remove-duplicate-letters/
// EXTRA: 重复个数不超过 limit https://leetcode.cn/contest/tianchi2022/problems/ev2bru/
func removeDuplicateLetters(s string) string {
	left := ['z' + 1]int{}
	for _, c := range s {
		left[c]++
	}
	inSt := ['z' + 1]bool{}
	st := []rune{}
	for _, c := range s {
		left[c]--
		if inSt[c] {
			continue
		}
		for len(st) > 0 && c < st[len(st)-1] && left[st[len(st)-1]] > 0 {
			top := st[len(st)-1]
			st = st[:len(st)-1]
			inSt[top] = false // top > c，且 top 后面还有，那么可以重新加进来
		}
		st = append(st, c)
		inSt[c] = true
	}
	return string(st)
}

// 求 a 的最长的子数组，其元素和大于 lowerSum
// 返回任意一个符合要求的子数组的左右端点（闭区间）
// 如果不存在，返回 [-1,-1]
// 讲解：https://leetcode.cn/problems/longest-well-performing-interval/solution/liang-chong-zuo-fa-liang-zhang-tu-miao-d-hysl/
// LC962 https://leetcode.cn/problems/maximum-width-ramp/
// LC1124 https://leetcode.cn/problems/longest-well-performing-interval/
// 有点相关 http://codeforces.com/problemset/problem/1788/E
func longestSubarrayWithLowerSum(a []int, lowerSum int) (int, int) {
	n := len(a)
	sum := make([]int, n+1)
	st := []int{0}
	for j, v := range a {
		j++
		sum[j] = sum[j-1] + v
		if sum[j] < sum[st[len(st)-1]] {
			st = append(st, j)
		}
	}

	l, r := -1, 0
	for i := n; i > 0; i-- {
		for len(st) > 0 && sum[i]-sum[st[len(st)-1]] > lowerSum {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			if l < 0 || i-j < r-l {
				l, r = j, i
			}
		}
	}
	r-- // 闭区间

	return l, r
}
