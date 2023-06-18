### 视频讲解

见[【周赛 350】](https://www.bilibili.com/video/BV1Hj411D7Tr/)第四题，欢迎点赞投币！

## 方法一：选或不选

### 前置知识：动态规划入门

详见 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)

### 思路

用「选或不选」的思路来思考：

- 如果付费刷第 $n-1$ 堵墙，那么问题变成：刷前 $n-2$ 堵墙，且付费时间和为 $\textit{time}[n-1]$，免费时间和 $0$ 的最少开销（开销指 $\textit{cost}$ 的和）。
- 如果免费刷第 $n-1$ 堵墙，那么问题变成：刷前 $n-2$ 堵墙，且付费时间和为 $0$，免费时间和为 $1$ 的最少开销。

因此，定义 $\textit{dfs}(i,j,k)$ 表示刷前 $i$ 堵墙，且付费时间和为 $j$，免费时间和为 $k$ 的最少开销。递归到终点时，如果 $j\ge k$，说明这种方案是合法的，否则不合法。

但是，这样定义的话，状态个数就太多了，需要优化。由于最后是比较的 $j$ 和 $k$ 的「相对大小」，那么不妨把 $j$ 重新定义为【付费时间和】减去【免费时间和】，这样递归到终点时，如果 $j\ge 0$，说明这种方案是合法的，否则不合法。这样一来，状态个数就大大减少了。

分类讨论：

- 如果付费刷第 $i$ 堵墙：$\textit{dfs}(i,j) = \textit{dfs}(i-1,j+\textit{time}[i]) +\textit{cost}[i]$。
- 如果免费刷第 $i$ 堵墙：$\textit{dfs}(i,j) = \textit{dfs}(i-1,j-1)$。

两种情况取最小值：

$$
\textit{dfs}(i,j) = \min(\textit{dfs}(i-1,j+\textit{time}[i]) +\textit{cost}[i], \textit{dfs}(i-1,j-1))
$$

递归边界：如果 $j>i$，那么剩余的墙都可以免费刷，即 $\textit{dfs}(i,j) = 0$，否则 $\textit{dfs}(-1,j) = \infty$。

递归入口：$\textit{dfs}(n-1,0)$。

```py [sol-Python3]
class Solution:
    def paintWalls(self, cost: List[int], time: List[int]) -> int:
        @cache  # 记忆化搜索
        def dfs(i: int, j: int) -> int:
            if j > i: return 0  # 剩余的墙都可以免费刷
            if i < 0: return inf
            return min(dfs(i - 1, j + time[i]) + cost[i], dfs(i - 1, j - 1))
        return dfs(len(cost) - 1, 0)
```

```go [sol-Go]
func paintWalls(cost, time []int) int {
	n := len(cost)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n*2+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if j-n > i { // 注意 j 加上了偏移量 n
			return 0 // 剩余的墙都可以免费刷
		}
		if i < 0 {
			return math.MaxInt / 2 // 防止加法溢出
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		*p = min(dfs(i-1, j+time[i])+cost[i], dfs(i-1, j-1))
		return *p
	}
	return dfs(n-1, n) // 加上偏移量 n，防止负数
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{cost}$ 的长度。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n^2)$，单个状态的计算时间为 $\mathcal{O}(1)$，因此时间复杂度为 $\mathcal{O}(n^2)$。
- 空间复杂度：$\mathcal{O}(n^2)$。

## 方法二：转换成 0-1 背包

### 前置知识：0-1 背包

详见[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)

### 思路

看着方法一的状态转移方程，你可能会觉得：怎么感觉很像 0-1 背包呢？没错，这题就是 0-1 背包！

根据题意，付费刷墙个数 $+$ 免费刷墙个数 $=n$。

同时，付费刷墙时间之和必须 $\ge$ 免费刷墙个数。

结合这两个式子，得到：付费刷墙时间之和 $\ge n\ -$ 付费刷墙个数。

移项，得到：「付费刷墙时间$+1$」之和 $\ge n$。（把个数拆分成 $1+1+1+\cdots$。）

把 $\textit{time}[i]+1$ 看成物品体积，$\textit{cost}[i]$ 看成物品价值，问题变成：

从 $n$ 个物品中选择体积和**至少**为 $n$ 的物品，价值和最小是多少？

这是 0-1 背包的一种「至少装满」的变形。我们可以定义 $\textit{dfs}(i,j)$ 表示考虑前 $i$ 个物品，**剩余**还需要凑出 $j$ 的体积，此时的最小价值和。

依然是用「选或不选」来思考，可以得到类似的状态转移方程：

$$
\textit{dfs}(i,j) = \min(\textit{dfs}(i-1,j-\textit{time}[i]-1)+\textit{cost}[i], \textit{dfs}(i-1,j))
$$

递归边界：如果 $j\le 0$，那么不需要再选任何物品了，返回 $0$；如果 $i<0$，返回无穷大。

递归入口：$\textit{dfs}(n-1,n)$，表示体积和至少为 $n$，这正是我们要计算的。

```py [sol-Python3]
class Solution:
    def paintWalls(self, cost: List[int], time: List[int]) -> int:
        @cache  # 记忆化搜索
        def dfs(i: int, j: int) -> int:  # j 表示剩余需要的体积
            if j <= 0: return 0  # 没有约束：后面什么也不用选了
            if i < 0: return inf  # 此时 j>0，但没有物品可选，不合法
            return min(dfs(i - 1, j - time[i] - 1) + cost[i], dfs(i - 1, j))
        n = len(cost)
        return dfs(n - 1, n)
```

```go [sol-Go]
func paintWalls(cost, time []int) int {
	n := len(cost)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int { // j 表示剩余需要的体积
		if j <= 0 { // 没有约束：后面什么也不用选了
			return 0
		}
		if i < 0 { // 此时 j>0，但没有物品可选，不合法
			return math.MaxInt / 2 // 防止加法溢出
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		*p = min(dfs(i-1, j-time[i]-1)+cost[i], dfs(i-1, j))
		return *p
	}
	return dfs(n-1, n)
}

func min(a, b int) int { if b < a { return b }; return a }
```

然后改成递推，方法在[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)中讲过了。

```py [sol-Python3]
class Solution:
    def paintWalls(self, cost: List[int], time: List[int]) -> int:
        n = len(cost)
        f = [0] + [inf] * n  # f[i][0] 表示 j<=0 的状态
        for c, t in zip(cost, time):
            for j in range(n, 0, -1):
                f[j] = min(f[j], f[max(j - t - 1, 0)] + c)
        return f[n]
```

```java [sol-Java]
class Solution {
    public int paintWalls(int[] cost, int[] time) {
        int n = cost.length;
        var f = new int[n + 1];
        Arrays.fill(f, Integer.MAX_VALUE / 2); // 防止加法溢出
        f[0] = 0;
        for (int i = 0; i < n; i++) {
            int c = cost[i], t = time[i] + 1;  // 注意这里加一了
            for (int j = n; j > 0; j--)
                f[j] = Math.min(f[j], f[Math.max(j - t, 0)] + c);
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int paintWalls(vector<int> &cost, vector<int> &time) {
        int n = cost.size(), f[n + 1];
        memset(f, 0x3f, sizeof(f));
        f[0] = 0;
        for (int i = 0; i < n; i++) {
            int c = cost[i], t = time[i] + 1; // 注意这里加一了
            for (int j = n; j; j--)
                f[j] = min(f[j], f[max(j - t, 0)] + c);
        }
        return f[n];
    }
};
```

```go [sol-Go]
func paintWalls(cost, time []int) int {
	n := len(cost)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2 // 防止加法溢出
	}
	for i, c := range cost {
		t := time[i] + 1 // 注意这里加一了
		for j := n; j > 0; j-- {
			f[j] = min(f[j], f[max(j-t, 0)]+c)
		}
	}
	return f[n]
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{cost}$ 的长度。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n^2)$，单个状态的计算时间为 $\mathcal{O}(1)$，因此时间复杂度为 $\mathcal{O}(n^2)$。
- 空间复杂度：$\mathcal{O}(n)$。

### 思考题

如果同一面墙可以反复刷要怎么做？

把 $\textit{cost}$ 去掉，如果有 $x$ 位付费油漆匠和 $y$ 位免费油漆匠，它们可以同时工作，刷完所有墙的最小耗时是多少？
