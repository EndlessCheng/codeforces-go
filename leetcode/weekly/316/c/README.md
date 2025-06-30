[视频讲解](https://www.bilibili.com/video/BV1ne4y1e7nu) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

## 方法一：枚举 + 考察变化量

将 $\textit{nums}$ 和 $\textit{cost}$ 绑在一起排序。

首先计算使所有元素都等于 $\textit{nums}[0]$ 的总开销 $\textit{total}$，以及所有 $\textit{cost}$ 的和 $\textit{sumCost}$。

然后考虑要使所有元素都等于 $\textit{nums}[1]$，$\textit{total}$ 的**变化量**是多少：

- 有 $\textit{cost}[0]$ 这么多的开销要增加 $\textit{nums}[1]-\textit{nums}[0]$；
- 有 $\textit{sumCost}-\textit{cost}[0]$ 这么多的开销要减少 $\textit{nums}[1]-\textit{nums}[0]$。

因此 $\textit{total}$ 减少了

$$
(\textit{sumCost} - 2 \cdot \textit{cost}[0]) \cdot (\textit{nums}[1]-\textit{nums}[0])
$$

按照这个公式模拟后续 $\textit{nums}[i]$，取所有 $\textit{total}$ 最小值为答案。

```py [sol-Python3]
class Solution:
    def minCost(self, nums: List[int], cost: List[int]) -> int:
        a = sorted(zip(nums, cost))
        ans = total = sum((x - a[0][0]) * c for x, c in a)
        sum_cost = sum(cost)
        for (x0, c), (x1, _) in pairwise(a):
            sum_cost -= c * 2
            total -= sum_cost * (x1 - x0)
            ans = min(ans, total)
        return ans
```

```go [sol-Go]
func minCost(nums, cost []int) int64 {
	type pair struct{ x, c int }
	a := make([]pair, len(nums))
	for i, x := range nums {
		a[i] = pair{x, cost[i]}
	}
	slices.SortFunc(a, func(p, q pair) int { return p.x - q.x })

	var total, sumCost int64
	for _, p := range a {
		total += int64(p.c) * int64(p.x-a[0].x)
		sumCost += int64(p.c)
	}
	ans := total
	for i := 1; i < len(a); i++ {
		sumCost -= int64(a[i-1].c * 2)
		total -= sumCost * int64(a[i].x-a[i-1].x)
		ans = min(ans, total)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

## 方法二：中位数贪心

把 $\textit{cost}[i]$ 理解成 $\textit{nums}[i]$ 的出现次数。

根据 [中位数贪心及其证明](https://zhuanlan.zhihu.com/p/1922938031687595039)，把所有数变成中位数是最优的。

代码实现时，仍然按照方法一那样排序，然后不断累加 $\textit{cost}[i]$，首次累加到 $\ge\dfrac{\textit{sumCost}}{2}$ 时就找到了中位数。

由于 $\textit{sumCost}$ 可能是奇数，所以要上取整，即首次累加到 $\ge\left\lceil\dfrac{\textit{sumCost}}{2}\right\rceil$ 时就找到了中位数。

```py [sol-Python3]
class Solution:
    def minCost(self, nums: List[int], cost: List[int]) -> int:
        a = sorted(zip(nums, cost))
        s, mid = 0, (sum(cost) + 1) // 2
        for x, c in a:
            s += c
            if s >= mid:
                return sum(abs(y - x) * c for y, c in a)  # 把所有数都变成 x
```

```go [sol-Go]
func minCost(nums, cost []int) (ans int64) {
	type pair struct{ x, c int }
	a := make([]pair, len(nums))
	sumCost := int64(0)
	for i, c := range cost {
		a[i] = pair{nums[i], c}
		sumCost += int64(c)
	}
	slices.SortFunc(a, func(p, q pair) int { return p.x - q.x })

	s, mid := int64(0), (sumCost+1)/2
	for _, p := range a {
		s += int64(p.c)
		if s >= mid {
			// 把所有数都变成 p.x
			for _, q := range a {
				ans += int64(abs(q.x-p.x)) * int64(q.c)
			}
			break
		}
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
