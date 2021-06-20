package copypasta

/* 单调栈 Monotone Stack
举例：返回每个元素两侧严格大于它的元素位置（不存在则为 -1 或 n）
如何理解：把数组想象成一列山峰，站在 a[i] 的山顶仰望两侧的山峰，是看不到高山背后的矮山的，只能看到一座座更高的山峰
         这就启发我们引入一个底大顶小的单调栈，入栈时不断比较栈顶元素直到找到一个比当前元素大的
技巧：事先压入一个边界元素到栈底，这样保证循环时栈一定不会为空，从而简化逻辑
一些转换：
    若区间 [l,r] 的最大值等于 a[r]，则 l 必须 > posL[r]
    若区间 [l,r] 的最大值等于 a[l]，则 r 必须 < posR[l]
    这一结论可以用于思考一些双变量的题目
https://oi-wiki.org/ds/monotonous-stack/
https://cp-algorithms.com/data_structures/stack_queue_modification.html

模板题 https://www.luogu.com.cn/problem/P5788
      https://www.luogu.com.cn/problem/P2866 http://poj.org/problem?id=3250
      https://leetcode-cn.com/problems/next-greater-element-i/ LC496/周赛18BA
      https://leetcode-cn.com/problems/next-greater-element-ii/ LC503/周赛18BB
      NEERC05，UVa 1619 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=825&page=show_problem&problem=4494
柱状图中最大的矩形 LC84 https://leetcode-cn.com/problems/largest-rectangle-in-histogram/ http://poj.org/problem?id=2559 http://poj.org/problem?id=2082
最大全 1 矩形 LC85（实现见下面的 maximalRectangleArea）https://leetcode-cn.com/problems/maximal-rectangle/
接雨水 LC42 https://leetcode-cn.com/problems/trapping-rain-water/
后缀数组+不同矩形对应方案数之和 https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/D
与 DP 结合
    https://codeforces.com/problemset/problem/1313/C2
    https://codeforces.com/problemset/problem/1407/D
全 1 子矩阵个数 O(n^2) LC周赛196C https://leetcode-cn.com/contest/weekly-contest-196/problems/count-submatrices-with-all-ones/ 原题为 http://poj.org/problem?id=3494
已知部分 posR 还原全部 posR；已知 posR 还原 a https://codeforces.com/problemset/problem/1158/C
*/
func monotoneStack(a []int) ([]int, []int) {
	const border int = -2e9 // 求两侧大的话用 2e9
	type pair struct{ v, i int }

	// 求左侧严格小于
	n := len(a)
	posL := make([]int, n)
	stack := []pair{{border, -1}}
	for i, v := range a {
		for {
			if top := stack[len(stack)-1]; top.v < v { //
				posL[i] = top.i
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{v, i})
	}

	// 求右侧严格小于
	posR := make([]int, n)
	stack = []pair{{border, n}}
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		for {
			if top := stack[len(stack)-1]; top.v < v { //
				posR[i] = top.i
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{v, i})
	}

	// EXTRA
	mx := 0
	for i, v := range a {
		l, r := posL[i]+1, posR[i] // [l,r)
		if v *= r - l; v > mx {
			mx = v
		}
	}

	return posL, posR
}

// 注：若输入的是一个 1~n 的排列，有更简单的写法（求两侧大于位置）
// 为简单起见，求出的下标从 1 开始（不存在时表示为 0 或 n+1）
// https://codeforces.com/contest/1156/problem/E
func permPosLR(a []int) ([]int, []int) {
	n := len(a)
	posV := make([]int, n+1)
	posL := make([]int, n+2)
	posR := make([]int, n+1)
	for i := 1; i <= n; i++ {
		posV[a[i-1]] = i
		posL[i], posR[i] = i-1, i+1
	}
	// 正序遍历求出的是两侧大于位置
	// 倒序遍历求出的是两侧小于位置
	for v := 1; v <= n; v++ {
		p := posV[v]
		posR[posL[p]] = posR[p]
		posL[posR[p]] = posL[p]
	}
	return posL, posR
}

// 最大全 1 矩形
// LC85 https://leetcode-cn.com/problems/maximal-rectangle/
func maximalRectangleArea(mat [][]int) (ans int) {
	const target = 1
	n, m := len(mat), len(mat[0])
	heights := make([][]int, n) // heights[i][j] 表示从 (i,j) 往上看的高度，a[i][j] = 0 时为 0
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

	type pair struct{ h, i int }
	for _, hs := range heights {
		posL := make([]int, m)
		stack := []pair{{-1, -1}}
		for j, h := range hs {
			for {
				if top := stack[len(stack)-1]; top.h < h {
					posL[j] = top.i
					break
				}
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, pair{h, j})
		}
		posR := make([]int, m)
		stack = []pair{{-1, m}}
		for j := m - 1; j >= 0; j-- {
			h := hs[j]
			for {
				if top := stack[len(stack)-1]; top.h < h {
					posR[j] = top.i
					break
				}
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, pair{h, j})
		}
		for j, h := range hs {
			if s := (posR[j] - posL[j] - 1) * h; s > ans {
				ans = s
			}
		}
	}
	return
}
