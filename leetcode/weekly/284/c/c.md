### 提示 1

考虑哪些数能在 $k$ 次操作后成为栈顶，答案是这些数的最大值。

### 提示 2

从特殊位置 $\textit{nums}[0]$ 入手，思考如何让 $\textit{nums}[0]$ 成为栈顶。（尝试枚举 $k=0,1,2,3,\cdots$）

### 提示 3

推广到其余 $\textit{nums}[i]$。

### 思路

设 $\textit{nums}$ 的长度为 $n$，根据题意：

- 如果 $n=1$，那么我们只能在栈不为空时删除栈顶，栈为空时将 $\textit{nums}[0]$ 入栈。因此 $k$ 为奇数时，$k$ 次操作后栈为空，返回 $-1$；$k$ 为偶数时则返回 $\textit{nums}[0]$。
- 如果 $k=0$，无法执行任何操作，直接返回 $\textit{nums}[0]$。

其余情况，按数组元素的下标 $i$ 分类讨论：

- 如果 $i=0$，我们可以不断地删除-添加 $\textit{nums}[0]$，如果 $k$ 为偶数，那么最后栈顶为 $\textit{nums}[0]$；如果 $k$ 为奇数（这里要求 $k>1$），我们可以在倒数第二步删除 $\textit{nums}[1]$，最后一步将 $\textit{nums}[0]$ 入栈，从而保证 $\textit{nums}[0]$ 可以为栈顶。 
- 如果 $0<i<k-1$，我们仍然可以仿造上述流程操作。
- 如果 $i=k-1$，最后一步操作只能删除 $\textit{nums}[i]$，所以无法将 $\textit{nums}[i]$ 置于栈顶。
- 如果 $i=k$，那么可以删除前 $k$ 个元素，将 $\textit{nums}[i]$ 置于栈顶。
- 如果 $i>k$，$\textit{nums}[i]$ 前面的元素无法删除，所以无法将 $\textit{nums}[i]$ 置于栈顶。

综上所述，我们可以让 $i<k-1$ 或 $i=k$ 的数组元素作为 $k$ 次操作后的栈顶。这些元素的最大值即为答案。

```Python [sol-Python3]
class Solution:
    def maximumTop(self, nums: List[int], k: int) -> int:
        if k % 2 and len(nums) == 1:
            return -1
        return max(num for i, num in enumerate(nums)
                       if i < k - 1 or i == k)
```

```go [sol-Go]
func maximumTop(a []int, k int) int {
	n := len(a)
	if n == 1 || k == 0 {
		if k%2 == 1 { return -1 }
		return a[0]
	}
	// 删除 a[k-1] 以及 a[k+1:]，下面直接取 a 的最大值
	if k < n {
		a = append(a[:k-1], a[k]) 
	} else if k == n {
		a = a[:n-1]
	}
	return slices.Max(a)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
