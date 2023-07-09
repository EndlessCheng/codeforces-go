下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，欢迎关注！

## 前置知识：动态规划入门

详见 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)

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

```py [sol-Python3]
class Solution:
    def maximumJumps(self, nums: List[int], target: int) -> int:
        @cache
        def dfs(i: int):
            if i == 0:
                return 0
            res = -inf
            for j in range(i):
                if -target <= nums[i] - nums[j] <= target:
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

func max(a, b int) int { if b > a { return b }; return a }
```

然后按照 [视频](https://www.bilibili.com/video/BV1Xj411K7oF/) 中讲的，1:1 翻译成递推。

```py [sol-Python3]
class Solution:
    def maximumJumps(self, nums: List[int], target: int) -> int:
        n = len(nums)
        f = [-inf] * n
        f[0] = 0
        for i in range(1, n):
            for j in range(i):
                if -target <= nums[i] - nums[j] <= target:
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

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(n)$，所以动态规划的时间复杂度为 $\mathcal{O}(n^2)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

如果改成求「最小跳跃次数」呢？

那样 BFS 也是可以做的。
