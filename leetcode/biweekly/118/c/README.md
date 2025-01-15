## 一、寻找子问题

我们需要解决的问题是：「获得第 $1$ 个及其后面的水果所需要的最少金币数」。

第 $1$ 个水果一定要买，然后呢？

第 $2$ 个水果可以购买，也可以免费获得：

- 如果购买，那么需要解决的问题为：「在购买第 $2$ 个水果的前提下，获得第 $2$ 个及其后面的水果所需要的最少金币数」。
- 如果免费获得，那么根据题意，第 $3$ 个水果必须购买，需要解决的问题为：「在购买第 $3$ 个水果的前提下，获得第 $3$ 个及其后面的水果所需要的最少金币数」。

无论哪种情况都会把原问题变成一个**和原问题相似的、规模更小的子问题**，所以可以用**递归**解决。

> DP 做题技巧：**见微知著**，想清楚第一步（或者最后一步）怎么做，就想清楚了递归过程中的每一步要怎么做。

## 二、状态定义与状态转移方程

从上面的讨论可以知道，只需要一个 $i$ 就能表达子问题，即定义 $\textit{dfs}(i)$ 表示在购买第 $i$ 个水果的前提下，获得第 $i$ 个及其后面的水果所需要的最少金币数。注意 $i$ 从 $1$ 开始。

买第 $i$ 个水果，那么从 $i+1$ 到 $2i$ 的水果都是免费的。枚举下一个购买的水果 $j$，问题变成：在购买第 $j$ 个水果的前提下，获得第 $j$ 个及其后面的水果所需要的最少金币数，即 $\textit{dfs}(j)$。

$j$ 的范围是 $[i+1,2i+1]$。其中 $2i+1$ 表示免费获得从 $i+1$ 到 $2i$ 的所有水果，那么第 $2i+1$ 个水果不能免费，一定要买。

这些 $\textit{dfs}(j)$ 取最小值，再加上购买第 $i$ 个水果的花费 $\textit{prices}[i]$，得

$$
\textit{dfs}(i) = \textit{prices}[i] + \min_{j=i+1}^{2i+1} \textit{dfs}(j)
$$

**递归边界**：注意到当 $2i\ge n$，即 $i\ge \left\lceil\dfrac{n}{2}\right\rceil = \left\lfloor\dfrac{n+1}{2}\right\rfloor$ 时，后面的水果都可以免费获得了，所以递归边界为

$$
\textit{dfs}(i)=\textit{prices}[i]
$$

其中 $i\ge \left\lfloor\dfrac{n+1}{2}\right\rfloor$。

**递归入口**：$\textit{dfs}(1)$，这是原问题，也是答案。

由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

原理见 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含如何把记忆化搜索 1:1 翻译成递推的技巧。

[本题视频讲解](https://www.bilibili.com/video/BV1Rw411P72r/?t=10m16s)，欢迎点赞关注~

> 注：由于传入的 $\textit{prices}$ 数组的下标是从 $0$ 开始的，代码实现时下标要减一。

```py [sol-Python3]
class Solution:
    def minimumCoins(self, prices: List[int]) -> int:
        n = len(prices)
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int) -> int:
            if i * 2 >= n:
                return prices[i - 1]  # i 从 1 开始
            return min(dfs(j) for j in range(i + 1, i * 2 + 2)) + prices[i - 1]
        return dfs(1)
```

```java [sol-Java]
class Solution {
    public int minimumCoins(int[] prices) {
        int n = prices.length;
        int[] memo = new int[(n + 1) / 2];
        return dfs(1, prices, memo);
    }

    private int dfs(int i, int[] prices, int[] memo) {
        if (i * 2 >= prices.length) {
            return prices[i - 1]; // i 从 1 开始
        }
        if (memo[i] != 0) { // 之前算过
            return memo[i];
        }
        int res = Integer.MAX_VALUE;
        for (int j = i + 1; j <= i * 2 + 1; j++) {
            res = Math.min(res, dfs(j, prices, memo));
        }
        return memo[i] = res + prices[i - 1]; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumCoins(vector<int>& prices) {
        int n = prices.size();
        vector<int> memo((n + 1) / 2);
        auto dfs = [&](this auto&& dfs, int i) -> int {
            if (i * 2 >= n) {
                return prices[i - 1]; // i 从 1 开始
            }
            int& res = memo[i]; // 注意这里是引用
            if (res) { // 之前算过
                return res;
            }
            res = INT_MAX;
            for (int j = i + 1; j <= i * 2 + 1; j++) {
                res = min(res, dfs(j));
            }
            res += prices[i - 1];
            return res;
        };
        return dfs(1);
    }
};
```

```go [sol-Go]
func minimumCoins(prices []int) int {
	n := len(prices)
	memo := make([]int, (n+1)/2)
	var dfs func(int) int
	dfs = func(i int) (res int) {
		if i*2 >= n {
			return prices[i-1] // i 从 1 开始
		}
		p := &memo[i]
		if *p != 0 { // 之前算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		res = math.MaxInt
		for j := i + 1; j <= i*2+1; j++ {
			res = min(res, dfs(j))
		}
		return res + prices[i-1]
	}
	return dfs(1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{prices}$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(n)$，所以动态规划的时间复杂度为 $\mathcal{O}(n^2)$。
- 空间复杂度：$\mathcal{O}(n)$。有多少个状态，$\textit{memo}$ 数组的大小就是多少。

## 三、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i]$ 的定义和 $\textit{dfs}(i)$ 的定义是一样的，都表示在购买第 $i$ 个水果的前提下，获得第 $i$ 个及其后面的水果所需要的最少金币数。注意 $i$ 从 $1$ 开始。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i] = \textit{prices}[i] + \min_{j=i+1}^{2i+1} f[j]
$$

> 注：由于从比 $i$ 更大的 $j$ 转移过来，所以必须倒着计算 $f$。

初始值：当 $i\ge \left\lfloor\dfrac{n+1}{2}\right\rfloor$ 时，$f[i]=\textit{prices}[i]$，翻译自递归边界 $\textit{dfs}(i)=\textit{prices}[i]$。

答案：$f[1]$，翻译自递归入口 $\textit{dfs}(1)$。

> 注：由于传入的 $\textit{prices}$ 数组的下标是从 $0$ 开始的，代码实现时下标要减一。

代码实现时，可以直接把 $\textit{prices}$ 当作 $f$ 数组。

```py [sol-Python3]
class Solution:
    def minimumCoins(self, f: List[int]) -> int:
        n = len(f)
        for i in range((n + 1) // 2 - 1, 0, -1):
            f[i - 1] += min(f[i: i * 2 + 1])
        return f[0]
```

```java [sol-Java]
class Solution {
    public int minimumCoins(int[] f) {
        int n = f.length;
        for (int i = (n + 1) / 2 - 1; i > 0; i--) {
            int mn = Integer.MAX_VALUE;
            for (int j = i; j <= i * 2; j++) {
                mn = Math.min(mn, f[j]);
            }
            f[i - 1] += mn;
        }
        return f[0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumCoins(vector<int>& f) {
        int n = f.size();
        for (int i = (n + 1) / 2 - 1; i > 0; i--) {
            f[i - 1] += *min_element(f.begin() + i, f.begin() + i * 2 + 1);
        }
        return f[0];
    }
};
```

```go [sol-Go]
func minimumCoins(f []int) int {
	n := len(f)
	for i := (n+1)/2 - 1; i > 0; i-- {
		f[i-1] += slices.Min(f[i : i*2+1])
	}
	return f[0]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{prices}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。Python 忽略切片开销。

## 四、单调队列优化

由于随着 $i$ 的变小，$j$ 的范围 $[i+1,2i+1]$ 的左右边界也在变小，所以 $[i+1,2i+1]$ 是一个向左的**滑动窗口**。

计算 $\min\limits_{j=i+1}^{2i+1} f[j]$ 的过程本质上是在计算**滑动窗口最小值**，原理见 [单调队列【基础算法精讲 27】](https://www.bilibili.com/video/BV1bM411X72E/)。

下面代码中的队首在左边，队尾在右边。

> 注：也可以把 $f$ 存到 $\textit{prices}$ 中，双端队列中只保存下标，但是这样可读性比较差，我在双端队列中保存了下标和对应的 $f[i]$。

```py [sol-Python3]
class Solution:
    def minimumCoins(self, prices: List[int]) -> int:
        n = len(prices)
        q = deque([(n + 1, 0)])  # 哨兵
        for i in range(n, 0, -1):
            while q[-1][0] > i * 2 + 1:  # 右边离开窗口
                q.pop()
            f = prices[i - 1] + q[-1][1]
            while f <= q[0][1]:
                q.popleft()
            q.appendleft((i, f))  # 左边进入窗口
        return q[0][1]
```

```java [sol-Java]
class Solution {
    public int minimumCoins(int[] prices) {
        int n = prices.length;
        Deque<int[]> q = new ArrayDeque<>();
        q.addLast(new int[]{n + 1, 0}); // 哨兵
        for (int i = n; i > 0; i--) {
            while (q.peekLast()[0] > i * 2 + 1) { // 右边离开窗口
                q.pollLast();
            }
            int f = prices[i - 1] + q.peekLast()[1];
            while (f <= q.peekFirst()[1]) {
                q.pollFirst();
            }
            q.addFirst(new int[]{i, f}); // 左边进入窗口
        }
        return q.peekFirst()[1];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumCoins(vector<int>& prices) {
        int n = prices.size();
        deque<pair<int, int>> q;
        q.emplace_front(n + 1, 0); // 哨兵
        for (int i = n; i > 0; i--) {
            while (q.back().first > i * 2 + 1) { // 右边离开窗口
                q.pop_back();
            }
            int f = prices[i - 1] + q.back().second;
            while (f <= q.front().second) {
                q.pop_front();
            }
            q.emplace_front(i, f); // 左边进入窗口
        }
        return q.front().second;
    }
};
```

```go [sol-Go]
func minimumCoins(prices []int) int {
	n := len(prices)
	type pair struct{ i, f int }
	q := []pair{{n + 1, 0}} // 哨兵
	for i := n; i > 0; i-- {
		for q[0].i > i*2+1 { // 右边离开窗口
			q = q[1:]
		}
		f := prices[i-1] + q[0].f
		for f <= q[len(q)-1].f {
			q = q[:len(q)-1]
		}
		q = append(q, pair{i, f}) // 左边进入窗口
	}
	return q[len(q)-1].f
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{prices}$ 的长度。每个下标只会入队出队各至多一次。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面动态规划题单中的「**§11.3 单调队列优化 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
