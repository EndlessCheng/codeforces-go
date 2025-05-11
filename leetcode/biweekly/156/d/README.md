## 方法一：多维 DP

反转操作的距离约束相当于：一旦执行了一次反转操作，那么会有 $k$ 秒的冷却期（CD），在冷却中不能执行反转操作。每往下走一步，CD 减一。

定义 $\textit{dfs}(x, \textit{cd}, \textit{parity})$ 表示当前递归到节点 $x$，剩余冷却时间为 $\textit{cd}$，$x$ 的祖先节点执行的反转操作次数的奇偶性是 $\textit{parity}$ 时，$x$ 子树的最大点权和。

设 $y$ 是 $x$ 的儿子。用选或不选（是否反转）思考：

- 不反转：往下递归到 $\textit{dfs}(y,\max(\textit{cd}-1,0), \textit{parity})$。
- 反转（前提是 $\textit{cd}=0$）：往下递归到 $\textit{dfs}(y,k-1, \textit{parity}\oplus 1)$。其中 $\oplus$ 表示异或运算。

两种情况，分别累加 $\textit{dfs}$ 的返回值，再加上 $\textit{nums}[x]$（不反转/反转）后的值，分别得到点权和 $s_0$ 和 $s_1$。

$\textit{dfs}(x, \textit{cd}, \textit{parity})$ 的返回值就是 $\max(s_0,s_1)$。

递归入口：$\textit{dfs}(0,0,0)$，即答案。

由于会重复访问状态，需要写记忆化搜索。

代码实现时，额外传入一个变量 $\textit{fa}$ 表示 $x$ 父节点，避免我们从 $x$ 递归到 $x$ 的父节点。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1m7EuzqEqr/?t=24m)，欢迎点赞关注~

### 写法一：记忆化搜索

```py [sol-Python3]
class Solution:
    def subtreeInversionSum(self, edges: List[List[int]], nums: List[int], k: int) -> int:
        n = len(nums)
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        memo = {}  # 手写 @cache

        # 这里为了计算方便，把 parity 改成 mul = 1 或者 -1
        def dfs(x: int, fa: int, cd: int, mul: int) -> int:
            t = (x, cd, mul)
            if t in memo:
                return memo[t]

            # 不反转
            res = nums[x] * mul
            for y in g[x]:
                if y != fa:
                    res += dfs(y, x, cd - 1 if cd else 0, mul)

            # 反转
            if cd == 0:
                mul *= -1
                s = nums[x] * mul
                for y in g[x]:
                    if y != fa:
                        s += dfs(y, x, k - 1, mul)
                if s > res:
                    res = s

            memo[t] = res
            return res

        return dfs(0, -1, 0, 1)
```

```java [sol-Java]
class Solution {
    public long subtreeInversionSum(int[][] edges, int[] nums, int k) {
        int n = nums.length;
        List<Integer>[] g = new List[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        long[][][] memo = new long[n][k][2];
        for (long[][] mat : memo) {
            for (long[] row : mat) {
                Arrays.fill(row, Long.MIN_VALUE);
            }
        }
        return dfs(0, -1, 0, 0, g, nums, k, memo);
    }

    private long dfs(int x, int fa, int cd, int parity, List<Integer>[] g, int[] nums, int k, long[][][] memo) {
        if (memo[x][cd][parity] != Long.MIN_VALUE) {
            return memo[x][cd][parity];
        }

        // 不反转
        long res = parity > 0 ? -nums[x] : nums[x];
        for (int y : g[x]) {
            if (y != fa) {
                res += dfs(y, x, Math.max(cd - 1, 0), parity, g, nums, k, memo);
            }
        }

        // 反转
        if (cd == 0) {
            long s = parity > 0 ? nums[x] : -nums[x];
            for (int y : g[x]) {
                if (y != fa) {
                    s += dfs(y, x, k - 1, parity ^ 1, g, nums, k, memo); // 重置 CD
                }
            }
            res = Math.max(res, s);
        }

        return memo[x][cd][parity] = res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long subtreeInversionSum(vector<vector<int>>& edges, vector<int>& nums, int k) {
        int n = nums.size();
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        vector memo(n, vector<array<long long, 2>>(k, {LLONG_MIN, LLONG_MIN}));
        auto dfs = [&](this auto&& dfs, int x, int fa, int cd, bool parity) -> long long {
            auto& res = memo[x][cd][parity]; // 注意这里是引用
            if (res != LLONG_MIN) {
                return res;
            }

            // 不反转
            res = parity ? -nums[x] : nums[x];
            for (int y : g[x]) {
                if (y != fa) {
                    res += dfs(y, x, max(cd - 1, 0), parity);
                }
            }

            // 反转
            if (cd == 0) {
                long long s = parity ? nums[x] : -nums[x];
                for (int y : g[x]) {
                    if (y != fa) {
                        s += dfs(y, x, k - 1, !parity); // 重置 CD
                    }
                }
                res = max(res, s);
            }

            return res;
        };

        return dfs(0, -1, 0, 0);
    }
};
```

```go [sol-Go]
func subtreeInversionSum(edges [][]int, nums []int, k int) int64 {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	memo := make([][][2]int, n)
	for i := range memo {
		memo[i] = make([][2]int, k)
		for j := range memo[i] {
			for p := range memo[i][j] {
				memo[i][j][p] = math.MinInt
			}
		}
	}
	var dfs func(int, int, int, int) int
	dfs = func(x, fa, cd, parity int) int {
		p := &memo[x][cd][parity]
		if *p != math.MinInt {
			return *p
		}

		// 不反转
		res := nums[x] * (1 - parity*2)
		for _, y := range g[x] {
			if y != fa {
				res += dfs(y, x, max(cd-1, 0), parity)
			}
		}

		// 反转
		if cd == 0 {
			s := nums[x] * (parity*2 - 1)
			for _, y := range g[x] {
				if y != fa {
					s += dfs(y, x, k-1, parity^1) // 重置 CD
				}
			}
			res = max(res, s)
		}

		*p = res
		return res
	}

	return int64(dfs(0, -1, 0, 0))
}
```

### 写法二：递推

```py [sol-Python3]
max = lambda a, b: b if b > a else a  # 手写 max 效率更高

class Solution:
    def subtreeInversionSum(self, edges: List[List[int]], nums: List[int], k: int) -> int:
        n = len(nums)
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        def dfs(x: int, fa: int) -> List[List[int]]:
            v = nums[x]
            res = [[v, -v] for _ in range(k)]
            s0, s1 = -v, v
            for y in g[x]:
                if y == fa:
                    continue
                fy = dfs(y, x)
                # 不反转
                for cd in range(k):
                    res[cd][0] += fy[max(cd - 1, 0)][0]
                    res[cd][1] += fy[max(cd - 1, 0)][1]
                # 反转
                s0 += fy[k - 1][1]
                s1 += fy[k - 1][0]
            # 反转
            res[0][0] = max(res[0][0], s0)
            res[0][1] = max(res[0][1], s1)
            return res

        return dfs(0, -1)[0][0]
```

```java [sol-Java]
class Solution {
    public long subtreeInversionSum(int[][] edges, int[] nums, int k) {
        int n = nums.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }
        return dfs(0, -1, g, nums, k)[0][0];
    }

    private long[][] dfs(int x, int fa, List<Integer>[] g, int[] nums, int k) {
        int v = nums[x];
        long[][] res = new long[k][2];
        for (int i = 0; i < k; i++) {
            res[i][0] = v;
            res[i][1] = -v;
        }
        long s0 = -v;
        long s1 = v;
        for (int y : g[x]) {
            if (y == fa) {
                continue;
            }
            long[][] fy = dfs(y, x, g, nums, k);
            // 不反转
            for (int cd = 0; cd < k; cd++) {
                res[cd][0] += fy[Math.max(cd - 1, 0)][0];
                res[cd][1] += fy[Math.max(cd - 1, 0)][1];
            }
            // 反转
            s0 += fy[k - 1][1];
            s1 += fy[k - 1][0];
        }
        // 反转
        res[0][0] = Math.max(res[0][0], s0);
        res[0][1] = Math.max(res[0][1], s1);
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long subtreeInversionSum(vector<vector<int>>& edges, vector<int>& nums, int k) {
        int n = nums.size();
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        auto dfs = [&](this auto&& dfs, int x, int fa) -> vector<array<long long, 2>> {
            int v = nums[x];
            vector<array<long long, 2>> res(k, {v, -v});
            long long s0 = -v, s1 = v;
            for (int y : g[x]) {
                if (y == fa) {
                    continue;
                }
                auto fy = dfs(y, x);
                // 不反转
                for (int cd = 0; cd < k; cd++) {
                    res[cd][0] += fy[max(cd - 1, 0)][0];
                    res[cd][1] += fy[max(cd - 1, 0)][1];
                }
                // 反转
                s0 += fy[k - 1][1];
                s1 += fy[k - 1][0];
            }
            // 反转
            res[0][0] = max(res[0][0], s0);
            res[0][1] = max(res[0][1], s1);
            return res;
        };

        return dfs(0, -1)[0][0];
    }
};
```

```go [sol-Go]
func subtreeInversionSum(edges [][]int, nums []int, k int) int64 {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int) [][2]int
	dfs = func(x, fa int) [][2]int {
		v := nums[x]
		res := make([][2]int, k)
		for cd := range res {
			res[cd] = [2]int{v, -v}
		}
		s0, s1 := -v, v
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			fy := dfs(y, x)
			// 不反转
			for cd := range res {
				res[cd][0] += fy[max(cd-1, 0)][0]
				res[cd][1] += fy[max(cd-1, 0)][1]
			}
			// 反转
			s0 += fy[k-1][1]
			s1 += fy[k-1][0]
		}
		// 反转
		res[0][0] = max(res[0][0], s0)
		res[0][1] = max(res[0][1], s1)
		return res
	}

	return int64(dfs(0, -1)[0][0])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(nk)$。

## 方法二：树上刷表法

设这棵树的所有点权和为 $S$，我们来算算，在 $S$ 的基础上，用反转操作，可以让 $S$ 增加多少。我们算的是这个增量的最大值。

$\textit{dfs}(x)$ 返回三个数：

- 子树 $x$ 的点权和 $s$。
- 在 $x$ 上面反转了偶数次的情况下的最大增量 $\textit{res}_0$。
- 在 $x$ 上面反转了奇数次的情况下的最大增量 $\textit{res}_1$。

设 $x$ 的儿子为 $y$。

**不反转** $x$，问题变成儿子 $y$ 的最大增量：

- 对于 $y$ 来说，上面反转次数的奇偶性不变。
- 累加 $\textit{dfs}(y)$ 的 $\textit{res}_0$，得到 $\textit{notInv}_0$。
- 累加 $\textit{dfs}(y)$ 的 $\textit{res}_1$，得到 $\textit{notInv}_1$。

**反转** $x$，问题变成 $x$ 的 $k$ 级后代的最大增量。我们分两部分计算。

第一部分：

- 如果 $x$ 上面反转了偶数次，我们会立刻得到 $-2s$ 的增量。
- 如果 $x$ 上面反转了奇数次，我们会立刻得到 $2s$ 的增量。

比如 $0\text{-}1\text{-}2$ 这棵树（这是一条链），点权为 $\textit{nums}=[-1,-1,1]$，$k=2$。反转根节点 $0$，我们会得到 $-2s = 2$ 的增量。此时点权为 $1,1,-1$，对于末端的叶子来说，反转它还可以再得到 $2s=2$ 的增量（注意这里的 $s$ 是基于原始 $\textit{nums}[2]=1$ 计算的）。所以在 $S=-1-1+1=-1$ 的基础上，（通过执行两次反转操作）一共可以得到 $2+2=4$ 的增量，最终答案为 $S+4=3$。

第二部分：

- 如果 $x$ 上面反转了偶数次，反转 $x$ 后，对于 $x$ 的 $k$ 级后代来说，上面反转了奇数次，所以累加的是 $x$ 的所有 $k$ 级后代的 $\textit{res}_1$。
- 如果 $x$ 上面反转了偶数次，反转 $x$ 后，对于 $x$ 的 $k$ 级后代来说，上面反转了偶数次，所以累加的是 $x$ 的所有 $k$ 级后代的 $\textit{res}_0$。

第一部分与第二部分相加（按照 $x$ 上面反转次数的奇偶性分开计算），结果分别记作 $\textit{inv}_0$ 和 $\textit{inv}_1$。

**注**：这里有两类转移来源。不反转 $x$，从 $x$ 的儿子 $y$ 转移过来；反转 $x$，从 $x$ 的 $k$ 级后代转移过来。读者可以联系 [309. 买卖股票的最佳时机含冷冻期](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/) 理解这个做法。注意反转 $x$ 是不能从儿子 $y$ 转移过来的，因为 $y$ 的返回值在计算下文的 $\max$ 时，可能取的是反转 $y$ 的情况，这违背了 $k$ 的约束。

计算 $\textit{dfs}(x)$ 的返回值：

$$
\begin{aligned}
\textit{res}_0 &= \max(\textit{notInv}_0, \textit{Inv}_0)     \\
\textit{res}_1 &= \max(\textit{notInv}_1, \textit{Inv}_1)     \\
\end{aligned}
$$

最后，如何累加 $x$ 的所有 $k$ 级后代 $z$ 的 $\textit{res}_i$？站在 $x$ 的视角计算，很不方便；但如果站在 $z$ 的视角，去更新 $x$，就很好算了。

```py [sol-Python3]
max = lambda a, b: b if b > a else a  # 手写 max 效率更高

class Solution:
    def subtreeInversionSum(self, edges: List[List[int]], nums: List[int], k: int) -> int:
        n = len(nums)
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        f = []

        def dfs(x: int, fa: int) -> Tuple[int, int, int]:
            f.append([0, 0])  # 用于刷表

            s = nums[x]  # 子树和
            not_inv0 = not_inv1 = 0  # 不反转 x 时的额外增量（0 表示上面反转了偶数次，1 表示上面反转了奇数次）
            for y in g[x]:
                if y == fa:
                    continue
                sy, y0, y1 = dfs(y, x)
                s += sy
                # 不反转 x，反转次数的奇偶性不变
                not_inv0 += y0
                not_inv1 += y1

            sub_res0, sub_res1 = f.pop()  # 被刷表后的结果

            # 反转 x
            # x 上面反转了偶数次，反转 x 会带来 -2 倍子树和的增量，且对于 x 的 k 级后代来说，上面反转了奇数次（所以是 sub_res1）
            inv0 = sub_res1 - s * 2
            # x 上面反转了奇数次，反转 x 会带来 2 倍子树和的增量，且对于 x 的 k 级后代来说，上面反转了偶数次（所以是 sub_res0）
            inv1 = sub_res0 + s * 2

            res0 = max(not_inv0, inv0)
            res1 = max(not_inv1, inv1)

            # 刷表法：更新 x 的 k 级祖先的状态
            if len(f) >= k:
                f[-k][0] += res0
                f[-k][1] += res1

            return s, res0, res1

        s, res0, _ = dfs(0, -1)
        return s + res0  # 对于根节点来说，上面一定反转了偶数次（0 次）
```

```java [sol-Java]
class Solution {
    public long subtreeInversionSum(int[][] edges, int[] nums, int k) {
        int n = nums.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        List<long[]> f = new ArrayList<>();
        long[] res = dfs(0, -1, g, nums, k, f);
        return res[0] + res[1]; // 对于根节点来说，上面一定反转了偶数次（0 次）
    }

    private long[] dfs(int x, int fa, List<Integer>[] g, int[] nums, int k, List<long[]> f) {
        f.add(new long[]{0, 0}); // 用于刷表

        long s = nums[x];  // 子树和
        long notInv0 = 0;
        long notInv1 = 0;  // 不反转 x 时的额外增量（0 表示上面反转了偶数次，1 表示上面反转了奇数次）
        for (int y : g[x]) {
            if (y == fa) {
                continue;
            }
            long[] resY = dfs(y, x, g, nums, k, f);
            s += resY[0];
            // 不反转 x，反转次数的奇偶性不变
            notInv0 += resY[1];
            notInv1 += resY[2];
        }

        long[] subRes = f.removeLast(); // 被刷表后的结果

        // 反转 x
        // x 上面反转了偶数次，反转 x 会带来 -2 倍子树和的增量，且对于 x 的 k 级后代来说，上面反转了奇数次（所以是 subRes1）
        long inv0 = subRes[1] - s * 2;
        // x 上面反转了奇数次，反转 x 会带来 2 倍子树和的增量，且对于 x 的 k 级后代来说，上面反转了偶数次（所以是 subRes0）
        long inv1 = subRes[0] + s * 2;

        long res0 = Math.max(notInv0, inv0);
        long res1 = Math.max(notInv1, inv1);

        // 刷表法：更新 x 的 k 级祖先的状态
        if (f.size() >= k) {
            long[] ancestor = f.get(f.size() - k);
            ancestor[0] += res0;
            ancestor[1] += res1;
        }

        return new long[]{s, res0, res1};
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long subtreeInversionSum(vector<vector<int>>& edges, vector<int>& nums, int k) {
        int n = nums.size();
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        vector<pair<long long, long long>> f;

        auto dfs = [&](this auto&& dfs, int x, int fa) -> tuple<long long, long long, long long> {
            f.emplace_back(0, 0); // 用于刷表

            long long s = nums[x]; // 子树和
            long long not_inv0 = 0, not_inv1 = 0; // 不反转 x 时的额外增量（0 表示上面反转了偶数次，1 表示上面反转了奇数次）
            for (int y : g[x]) {
                if (y == fa) {
                    continue;
                }
                auto [sy, y0, y1] = dfs(y, x);
                s += sy;
                // 不反转 x，反转次数的奇偶性不变
                not_inv0 += y0;
                not_inv1 += y1;
            }

            auto [sub_res0, sub_res1] = f.back(); // 被刷表后的结果
            f.pop_back();

            // 反转 x
            // x 上面反转了偶数次，反转 x 会带来 -2 倍子树和的增量，且对于 x 的 k 级后代来说，上面反转了奇数次（所以是 sub_res1）
            long long inv0 = sub_res1 - s * 2;
            // x 上面反转了奇数次，反转 x 会带来 2 倍子树和的增量，且对于 x 的 k 级后代来说，上面反转了偶数次（所以是 sub_res0）
            long long inv1 = sub_res0 + s * 2;

            long long res0 = max(not_inv0, inv0);
            long long res1 = max(not_inv1, inv1);

            // 刷表法：更新 x 的 k 级祖先的状态
            if (f.size() >= k) {
                f[f.size() - k].first += res0;
                f[f.size() - k].second += res1;
            }

            return {s, res0, res1};
        };

        auto [s, res0, _] = dfs(0, -1);
        return s + res0;  // 对于根节点来说，上面一定反转了偶数次（0 次）
    }
};
```

```go [sol-Go]
func subtreeInversionSum(edges [][]int, nums []int, k int) int64 {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	f := [][2]int{}
	var dfs func(int, int) (int, int, int)
	dfs = func(x, fa int) (int, int, int) {
		f = append(f, [2]int{}) // 用于刷表

		s := nums[x] // 子树和
		notInv0, notInv1 := 0, 0 // 不反转 x 时的额外增量（0 表示上面反转了偶数次，1 表示上面反转了奇数次）
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			sy, y0, y1 := dfs(y, x)
			s += sy
			// 不反转 x，反转次数的奇偶性不变
			notInv0 += y0
			notInv1 += y1
		}

		subRes := f[len(f)-1] // 被刷表后的结果
		f = f[:len(f)-1]

		// 反转 x
		// x 上面反转了偶数次，反转 x 会带来 -2 倍子树和的增量，且对于 x 的 k 级后代来说，上面反转了奇数次（所以是 subRes1）
		inv0 := subRes[1] - s*2
		// x 上面反转了奇数次，反转 x 会带来 2 倍子树和的增量，且对于 x 的 k 级后代来说，上面反转了偶数次（所以是 subRes0）
		inv1 := subRes[0] + s*2

		res0 := max(notInv0, inv0)
		res1 := max(notInv1, inv1)

		// 刷表法：更新 x 的 k 级祖先的状态
		if len(f) >= k {
			f[len(f)-k][0] += res0
			f[len(f)-k][1] += res1
		}

		return s, res0, res1
	}

	s, res0, _ := dfs(0, -1)
	return int64(s + res0) // 对于根节点来说，上面一定反转了偶数次（0 次）
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面动态规划题单的「**十二、树形 DP**」。

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
