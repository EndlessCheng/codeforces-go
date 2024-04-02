## 前置知识：动态规划入门

请看视频：[动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)

## 思路

为方便后面翻译成递推，这里改成从 $n-1$ 倒着跳到 $0$。

由于只能向左跳，每次跳跃之后都会把问题规模缩小，那么可以定义 $\textit{dfs}(i)$ 表示从 $i$ 跳到 $0$ 的最大跳跃次数。

用「枚举选哪个」来思考：

枚举 $j$，如果 $-\textit{target}\le \textit{nums}[i]-\textit{nums}[j] \le \textit{target}$，那么有

$$
\textit{dfs}(i) = \textit{dfs}(j) + 1
$$

取所有情况的最大值，即为 $\textit{dfs}(i)$。

如果没有这样的 $j$，那么 $\textit{dfs}(i) = -\infty$。

递归边界：$\textit{dfs}(0)=0$。

递归入口：$\textit{dfs}(n-1)$。也就是答案。如果答案是负数就返回 $-1$。

附：[视频讲解](https://www.bilibili.com/video/BV1XW4y1f7Wv/) 第二题。

```py [sol-Python3]
class Solution:
    def maximumJumps(self, nums: List[int], target: int) -> int:
        @cache
        def dfs(i: int):
            if i == 0:
                return 0
            res = -inf
            for j in range(i):
                if abs(nums[i] - nums[j]) <= target:
                    res = max(res, dfs(j) + 1)
            return res
        ans = dfs(len(nums) - 1)
        return -1 if ans < 0 else ans
```

```go [sol-Go]
func maximumJumps(nums []int, target int) int {
	n := len(nums)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1 // -1 表示没有计算过
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i == 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 { // 之前算过了
			return *p
		}
		res := math.MinInt
		for j, x := range nums[:i] {
			if -target <= nums[i]-x && nums[i]-x <= target {
				res = max(res, dfs(j)+1)
			}
		}
		*p = res // 记忆化
		return res
	}
	ans := dfs(n - 1)
	if ans < 0 {
		return -1
	}
	return ans
}
```

然后按照视频中讲的，1:1 翻译成递推。

```py [sol-Python3]
class Solution:
    def maximumJumps(self, nums: List[int], target: int) -> int:
        n = len(nums)
        f = [-inf] * n
        f[0] = 0
        for i in range(1, n):
            for j in range(i):
                if abs(nums[i] - nums[j]) <= target:
                    f[i] = max(f[i], f[j] + 1)
        return -1 if f[-1] < 0 else f[-1]
```

```go [sol-Go]
func maximumJumps(nums []int, target int) int {
	n := len(nums)
	f := make([]int, n)
	for i := 1; i < n; i++ {
		f[i] = math.MinInt
		for j, x := range nums[:i] {
			if -target <= nums[i]-x && nums[i]-x <= target {
				f[i] = max(f[i], f[j]+1)
			}
		}
	}
	if f[n-1] < 0 {
		return -1
	}
	return f[n-1]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(n)$，所以动态规划的时间复杂度为 $\mathcal{O}(n^2)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

如果改成求「最小跳跃次数」呢？

那样 BFS 也是可以做的。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
