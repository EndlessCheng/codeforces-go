package copypasta

/* 单调栈 Monotone Stack
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

模板题
https://www.luogu.com.cn/problem/P5788
https://www.luogu.com.cn/problem/P2866 http://poj.org/problem?id=3250
https://leetcode-cn.com/problems/next-greater-element-i/ LC496/周赛18BA
https://leetcode-cn.com/problems/next-greater-element-ii/ LC503/周赛18BB
NEERC05，UVa 1619 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=825&page=show_problem&problem=4494

计算贡献
LC907 https://leetcode.cn/problems/sum-of-subarray-minimums/
LC1856 https://leetcode.cn/problems/maximum-subarray-min-product/
LC2104 https://leetcode.cn/problems/sum-of-subarray-ranges/
LC2281 https://leetcode.com/problems/sum-of-total-strength-of-wizards/

与 DP 结合
https://codeforces.com/problemset/problem/1313/C2
https://codeforces.com/problemset/problem/1407/D
结合线段树，或者巧妙地在单调栈中去维护最值 https://codeforces.com/problemset/problem/1483/C
单调队列优化 LC375 猜数字大小 II https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/

其他
LC42 接雨水 https://leetcode-cn.com/problems/trapping-rain-water/
     评注：接雨水有三种不同的解法（DP、单调栈和双指针），其中双指针是 DP 的空间优化写法
          本质上是两种计算策略：计算每个下标处的接水量（纵向累加），计算一段高度对应的接水宽度（横向累加）
LC84 柱状图中最大的矩形 https://leetcode-cn.com/problems/largest-rectangle-in-histogram/ http://poj.org/problem?id=2559 http://poj.org/problem?id=2082
LC85 最大全 1 矩形（实现见下面的 maximalRectangleArea）https://leetcode-cn.com/problems/maximal-rectangle/ 原题为 http://poj.org/problem?id=3494
LC1504/周赛196C 全 1 矩形个数（实现见下面的 numSubmat）https://leetcode-cn.com/problems/count-submatrices-with-all-ones/
后缀数组+不同矩形对应方案数之和 https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/D
与 bitOpTrickCnt 结合（见 bits.go）https://codeforces.com/problemset/problem/875/D
已知部分 right 还原全部 right；已知 right 还原 a https://codeforces.com/problemset/problem/1158/C
*/
func monotoneStack(a []int) ([]int, []int) {
	const mod int = 1e9 + 7

	// 考察局部最小
	// 如果有相同元素，需要把某一侧循环内的符号改成小于等于

	// 求左侧严格小于 a[i] 的最近位置 left[i]，这样 a[i] 就是区间 [left[i]+1,i] 内最小的元素（之一）
	// 如果改成小于等于，那么 a[i] 就是区间 [left[i]+1,i] 内独一无二的最小元素
	n := len(a)
	left := make([]int, n)
	const border int = -2e9 // 求两侧大的话用 2e9
	type pair struct{ v, i int }
	stack := []pair{{border, -1}}
	for i, v := range a {
		for {
			if top := stack[len(stack)-1]; top.v < v { //
				left[i] = top.i
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{v, i})
	}

	// 求右侧严格小于 a[i] 的最近位置 right[i]，这样 a[i] 就是区间 [i,right[i]-1] 内最小的元素（之一）
	// 如果改成小于等于，那么 a[i] 就是区间 [i,right[i]-1] 内独一无二的最小元素
	right := make([]int, n)
	stack = []pair{{border, n}}
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		for {
			if top := stack[len(stack)-1]; top.v < v { //
				right[i] = top.i
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{v, i})
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
		st := []int{}
		for i, v := range a {
			for len(st) > 0 && a[st[len(st)-1]] >= v { // 这里是右侧小于等于
				right[st[len(st)-1]] = i
				st = st[:len(st)-1]
			}
			if len(st) > 0 {
				left[i] = st[len(st)-1]
			} else {
				left[i] = -1
			}
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

// 注：若输入的是一个 1~n 的排列，有更简单的写法（求两侧大于位置）
// 为简单起见，求出的下标从 1 开始（不存在时表示为 0 或 n+1）
// https://codeforces.com/contest/1156/problem/E
func permLR(a []int) ([]int, []int) {
	n := len(a)
	idx := make([]int, n+1)
	left := make([]int, n+2)
	right := make([]int, n+1)
	for i := 1; i <= n; i++ {
		idx[a[i-1]] = i
		left[i], right[i] = i-1, i+1
	}
	// 正序遍历求出的是两侧大于位置
	// 倒序遍历求出的是两侧小于位置
	for v := 1; v <= n; v++ {
		i := idx[v]
		right[left[i]] = right[i]
		left[right[i]] = left[i]
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
// LC1504/周赛196C https://leetcode-cn.com/problems/count-submatrices-with-all-ones/
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
