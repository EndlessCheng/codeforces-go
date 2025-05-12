[视频讲解](https://www.bilibili.com/video/BV1AU411F7Fp/) 第四题。

## 方法一：树形 DP

**前置知识**：[树形 DP【基础算法精讲 24】](https://www.bilibili.com/video/BV1vu4y1f7dn/)

用「选或不选」思考。

对于以 $x$ 为根的子树，考虑 $x$ 和它的儿子 $y$ 之间的边是否操作。

- 定义 $f[x][0]$ 表示 $x$ 操作偶数次时，子树 $x$ 的除去 $x$ 的最大价值和。
- 定义 $f[x][1]$ 表示 $x$ 操作奇数次时，子树 $x$ 的除去 $x$ 的最大价值和。

初始化 $f[x][0]=0,\ f[x][1] = -\infty$。遍历并递归计算 $x$ 的所有儿子，设当前遍历到的儿子为 $y$，

- 情况一，不操作 $x$ 和 $y$：
    - 设 $r_0 = \max(f[y][0] + \textit{nums}[y], f[y][1] + (\textit{nums}[y] \oplus k))$。这是不操作 $x$ 和 $y$ 时，子树 $y$ 的最大价值和。
    - $f[x][0] = f[x][0] + r_0$。
    - $f[x][1] = f[x][1] + r_0$。
- 情况二，操作 $x$ 和 $y$：
    - 设 $r_1 = \max(f[y][1] + \textit{nums}[y], f[y][0] + (\textit{nums}[y] \oplus k))$。这是操作 $x$ 和 $y$ 时，子树 $y$ 的最大价值和。
    - $f[x][0] = f[x][1] + r_1$。注意操作后，$x$ 的操作次数的奇偶性变化了。
    - $f[x][1] = f[x][0] + r_1$。

两种情况取最大值，有

$$
\begin{aligned}
&f[x][0] = \max(f[x][0] + r_0, f[x][1] + r_1)\\
&f[x][1] = \max(f[x][1] + r_0, f[x][0] + r_1) 
\end{aligned}
$$

注意这两个转移是同时发生的。

最后答案为根节点对应的 $r_0$。

```py [sol-Python3]
class Solution:
    def maximumValueSum(self, nums: List[int], k: int, edges: List[List[int]]) -> int:
        g = [[] for _ in nums]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        def dfs(x: int, fa: int) -> Tuple[int, int]:
            f0, f1 = 0, -inf  # f[x][0] 和 f[x][1]
            for y in g[x]:
                if y != fa:
                    r0, r1 = dfs(y, x)
                    f0, f1 = max(f0 + r0, f1 + r1), max(f1 + r0, f0 + r1)
            return max(f0 + nums[x], f1 + (nums[x] ^ k)), max(f1 + nums[x], f0 + (nums[x] ^ k))
        return dfs(0, -1)[0]
```

```java [sol-Java]
class Solution {
    public long maximumValueSum(int[] nums, int k, int[][] edges) {
        int n = nums.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }
        return dfs(0, -1, g, nums, k)[0];
    }

    private long[] dfs(int x, int fa, List<Integer>[] g, int[] nums, int k) {
        long f0 = 0;
        long f1 = Long.MIN_VALUE; // f[x][0] 和 f[x][1]
        for (int y : g[x]) {
            if (y != fa) {
                long[] r = dfs(y, x, g, nums, k);
                long t = Math.max(f1 + r[0], f0 + r[1]);
                f0 = Math.max(f0 + r[0], f1 + r[1]);
                f1 = t;
            }
        }
        return new long[]{Math.max(f0 + nums[x], f1 + (nums[x] ^ k)),
                          Math.max(f1 + nums[x], f0 + (nums[x] ^ k))};
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumValueSum(vector<int> &nums, int k, vector<vector<int>> &edges) {
        int n = nums.size();
        vector<vector<int>> g(n);
        for (auto &e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        function<pair<long long, long long>(int, int)> dfs = [&](int x, int fa) -> pair<long long, long long> {
            long long f0 = 0, f1 = LLONG_MIN; // f[x][0] 和 f[x][1]
            for (auto &y : g[x]) {
                if (y != fa) {
                    auto [r0, r1] = dfs(y, x);
                    long long t = max(f1 + r0, f0 + r1);
                    f0 = max(f0 + r0, f1 + r1);
                    f1 = t;
                }
            }
            return {max(f0 + nums[x], f1 + (nums[x] ^ k)), max(f1 + nums[x], f0 + (nums[x] ^ k))};
        };
        return dfs(0, -1).first;
    }
};
```

```go [sol-Go]
func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		f0, f1 := 0, math.MinInt // f[x][0] 和 f[x][1]
		for _, y := range g[x] {
			if y != fa {
				r0, r1 := dfs(y, x)
				f0, f1 = max(f0+r0, f1+r1), max(f1+r0, f0+r1)
			}
		}
		return max(f0+nums[x], f1+(nums[x]^k)), max(f1+nums[x], f0+(nums[x]^k))
	}
	ans, _ := dfs(0, -1)
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：结论 + 状态机 DP

### 性质一

由于一个数异或两次 $k$ 后保持不变，所以对于一条从 $x$ 到 $y$ 的简单路径，我们把路径上的所有边操作后，路径上除了 $x$ 和 $y$ 的其它节点都恰好操作两次，所以只有 $\textit{nums}[x]$ 和 $\textit{nums}[y]$ 都异或了 $k$，其余元素不变。

所以题目中的操作可以作用在**任意两个数**上。我们不需要建树，$\textit{edges}$ 是多余的！

### 性质二

注意到，无论操作多少次，总是有**偶数个**元素异或了 $k$。理由如下：

- 如果我们操作的两个数之前都没有异或过，那么操作后，异或 $k$ 的元素增加了 $2$。
- 如果我们操作的两个数之前都异或过，那么操作后，异或 $k$ 的元素减少了 $2$。
- 如果我们操作的两个数之前一个异或过，另一个没有异或过，那么操作后，异或 $k$ 的元素加一减一，不变。

### 思路

结合这两个性质，问题变成：

- 选择 $\textit{nums}$ 中的**偶数**个元素，把这些数都异或 $k$，数组的最大元素和是多少？

这可以用状态机 DP 解决。

- 定义 $f[i][0]$ 表示选择 $\textit{nums}$ 的前 $i$ 数中的偶数个元素异或 $k$，得到的最大元素和。
- 定义 $f[i][1]$ 表示选择 $\textit{nums}$ 的前 $i$ 数中的奇数个元素异或 $k$，得到的最大元素和。

设 $x=\textit{nums}[i]$。

- 情况一，不操作 $x$：
  - $f[i+1][0] = f[i][0] + x$。
  - $f[i+1][1] = f[i][1] + x$。
- 情况二，操作 $x$：
  - $f[i+1][0] = f[i][1] + (x\oplus k)$。
  - $f[i+1][1] = f[i][0] + (x\oplus k)$。

两种情况取最大值，有

$$
\begin{aligned}
&f[i+1][0] = \max(f[i][0] + x, f[i][1] + (x\oplus k))\\
&f[i+1][1] = \max(f[i][1] + x, f[i][0] + (x\oplus k))
\end{aligned}
$$

初始值 $f[0][0] = 0,\ f[0][1] = -\infty$。

答案为 $f[n][0]$。

代码实现时，$f$ 数组的第一个维度可以压缩掉。

```py [sol-Python3]
class Solution:
    def maximumValueSum(self, nums: List[int], k: int, _: List[List[int]]) -> int:
        f0, f1 = 0, -inf
        for x in nums:
            f0, f1 = max(f0 + x, f1 + (x ^ k)), max(f1 + x, f0 + (x ^ k))
        return f0
```

```java [sol-Java]
class Solution {
    public long maximumValueSum(int[] nums, int k, int[][] edges) {
        long f0 = 0;
        long f1 = Long.MIN_VALUE;
        for (int x : nums) {
            long t = Math.max(f1 + x, f0 + (x ^ k));
            f0 = Math.max(f0 + x, f1 + (x ^ k));
            f1 = t;
        }
        return f0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumValueSum(vector<int> &nums, int k, vector<vector<int>> &edges) {
        long long f0 = 0, f1 = LLONG_MIN;
        for (int x : nums) {
            long long t = max(f1 + x, f0 + (x ^ k));
            f0 = max(f0 + x, f1 + (x ^ k));
            f1 = t;
        }
        return f0;
    }
};
```

```go [sol-Go]
func maximumValueSum(nums []int, k int, _ [][]int) int64 {
	f0, f1 := 0, math.MinInt
	for _, x := range nums {
		f0, f1 = max(f0+x, f1+(x^k)), max(f1+x, f0+(x^k))
	}
	return int64(f0)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
