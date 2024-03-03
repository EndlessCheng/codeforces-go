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

        def dfs(x: int, fa: int) -> (int, int):
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
        long f0 = 0, f1 = Long.MIN_VALUE; // f[x][0] 和 f[x][1]
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

#### 第一个观察

由于一个数异或两次（偶数次）$k$ 后保持不变，所以对于一条从 $x$ 到 $y$ 的简单路径，我们把路径上的所有边操作后，路径上除了 $x$ 和 $y$ 的其它节点都恰好操作两次，所以只有 $\textit{nums}[x]$ 和 $\textit{nums}[y]$ 都异或了 $k$，其余元素不变。

所以题目中的操作可以作用在**任意两个数**上。我们不需要建树，$\textit{edges}$ 是多余的。

#### 第二个观察

注意到，无论操作多少次，总是有**偶数个**元素异或了 $k$，其余元素不变，理由如下：

- 如果我们操作的两个数之前都没有异或过，那么操作后，异或 $k$ 的元素增加了 $2$。
- 如果我们操作的两个数之前都异或过，那么操作后，异或 $k$ 的元素减少了 $2$。
- 如果我们操作的两个数之前一个异或过，另一个没有异或过，那么操作后，异或 $k$ 的元素加一减一，不变。

#### 思路

结合这两个观察，问题变成：

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
        long f0 = 0, f1 = Long.MIN_VALUE;
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

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
