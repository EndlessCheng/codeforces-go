## 方法一：状压 DP

如果你从未做过状压 DP，推荐先阅读 [教你一步步思考状压 DP：从记忆化搜索到递推](https://leetcode.cn/problems/beautiful-arrangement/solution/jiao-ni-yi-bu-bu-si-kao-zhuang-ya-dpcong-c6kd/)。

本题是**相邻相关排列型状压 DP**。标准套路是定义 $\textit{dfs}(\textit{mask},i)$ 表示处理完请求集合 $\textit{mask}$，且电梯停在 $\textit{floor}_i$（最后处理的是请求 $i$），所需的最短时间。

枚举上一次处理的请求是 $j$，问题变成处理完请求集合 $\textit{mask} \setminus \{i\}$，且电梯停在 $\textit{floor}_j$，所需的最短时间，即 $\textit{dfs}(\textit{mask} \setminus \{i\}, j)$。那么完成请求 $i$ 的时间为

$$
\max(\textit{dfs}(\textit{mask} \setminus \{i\}, j) + |\textit{floor}_i - \textit{floor}_j|, \textit{arrival}_i)
$$

> 注：完成请求 $i$ 的时间不能早于 $\textit{arrival}_i$。

用上式更新 $\textit{dfs}(\textit{mask} \setminus \{i\}, j)$ 的最小值，即

$$
\textit{dfs}(\textit{mask} \setminus \{i\}, j) = \min_{j\in \textit{mask} \setminus \{i\}} \max(\textit{dfs}(\textit{mask} \setminus \{i\}, j) + |\textit{floor}_i - \textit{floor}_j|, \textit{arrival}_i)
$$

**递归边界**：$\textit{dfs}(\{i\}, i) = \max(|\textit{floor}_i - \textit{start}|, \textit{arrival}_i)$。此时只有一个请求，即我们处理的第一个请求。需要从 $\textit{start}$ 移动到 $\textit{floor}_i$。

**递归入口**：$\min\limits_{i=0}^{m-1}\textit{dfs}(U,i)$，其中 $m$ 是 $\textit{requests}$ 的长度，全集 $U=\{0,1,2,\ldots,m-1\}$。

代码实现时，需要把集合语言翻译成位运算语言，见 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

[本题视频讲解](https://www.bilibili.com/video/BV1gRbD6zECR/?t=10m29s)，欢迎点赞关注~

### 记忆化搜索

```py [sol-Python3]
class Solution:
    def elevatorRequests(self, n: int, start: int, requests: list[list[int]]) -> int:
        # 返回处理完请求集合 mask，且电梯停在 requests[i][1]，所需的最短时间
        @cache
        def dfs(mask: int, i: int) -> int:
            mask ^= 1 << i  # 这里去掉了 i
            t, x = requests[i]
            if mask == 0:  # i 是第一个被处理的请求
                return max(abs(x - start), t)
            # 处理完请求 j 的时间 + 从 j 到 i 的时间
            res = min(dfs(mask, j) + abs(x - y) for j, (_, y) in enumerate(requests) if mask >> j & 1)
            return max(res, t)  # 处理完请求 i 的时间不能早于 t

        m = len(requests)
        return min(dfs((1 << m) - 1, i) for i in range(m))  # 枚举最后处理的请求
```

```java [sol-Java]
class Solution {
    public long elevatorRequests(int n, int start, int[][] requests) {
        int m = requests.length;
        long[][] memo = new long[1 << m][m];
        for (long[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }

        long ans = Long.MAX_VALUE;
        for (int i = 0; i < m; i++) { // 枚举最后处理的请求
            ans = Math.min(ans, dfs((1 << m) - 1, i, start, requests, memo));
        }
        return ans;
    }

    // 返回处理完请求集合 mask，且电梯停在 requests[i][1]，所需的最短时间
    private long dfs(int mask, int i, int start, int[][] requests, long[][] memo) {
        mask ^= 1 << i; // 这里去掉了 i
        int[] req = requests[i];
        int t = req[0];
        int x = req[1];

        if (mask == 0) {
            // i 是第一个被处理的请求
            return Math.max(Math.abs(x - start), t);
        }

        if (memo[mask][i] != -1) { // 之前计算过
            return memo[mask][i];
        }

        long res = Long.MAX_VALUE;
        for (int j = 0; j < requests.length; j++) {
            if ((mask >> j & 1) > 0) {
                // 处理完请求 j 的时间 + 从 j 到 i 的时间
                res = Math.min(res, dfs(mask, j, start, requests, memo) + Math.abs(x - requests[j][1]));
            }
        }
        // 处理完请求 i 的时间不能早于 t
        res = Math.max(res, t);

        memo[mask][i] = res; // 记忆化
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long elevatorRequests(int n, int start, vector<vector<int>>& requests) {
        int m = requests.size();
        vector memo(1 << m, vector<long long>(m, LLONG_MAX));

        // 返回处理完请求集合 mask，且电梯停在 requests[i][1]，所需的最短时间
        auto dfs = [&](this auto&& dfs, int mask, int i) -> long long {
            mask ^= 1 << i; // 这里去掉了 i
            auto& req = requests[i];
            int t = req[0], x = req[1];

            if (mask == 0) {
                // i 是第一个被处理的请求
                return max(abs(x - start), t);
            }

            auto& res = memo[mask][i];
            if (res != LLONG_MAX) { // 之前计算过
                return res;
            }

            for (int j = 0; j < m; j++) {
                if (mask >> j & 1) {
                    // 处理完请求 j 的时间 + 从 j 到 i 的时间
                    res = min(res, dfs(mask, j) + abs(x - requests[j][1]));
                }
            }
            // 处理完请求 i 的时间不能早于 t
            res = max(res, 1LL * t);
            return res;
        };

        long long ans = LLONG_MAX;
        for (int i = 0; i < m; i++) { // 枚举最后处理的请求
            ans = min(ans, dfs((1 << m) - 1, i));
        }
        return ans;
    }
};
```

```go [sol-Go]
func elevatorRequests(n int, start int, requests [][]int) int64 {
	m := len(requests)
	memo := make([][]int, 1<<m)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}

	// 返回处理完请求集合 mask，且电梯停在 requests[i][1]，所需的最短时间
	var dfs func(int, int) int
	dfs = func(mask, i int) int {
		mask ^= 1 << i // 这里去掉了 i
		req := requests[i]
		t, x := req[0], req[1]

		if mask == 0 {
			// i 是第一个被处理的请求
			return max(abs(x-start), t)
		}

		p := &memo[mask][i]
		if *p != -1 { // 之前计算过
			return *p
		}

		res := math.MaxInt
		for j, r := range requests {
			if mask>>j&1 > 0 {
				// 处理完请求 j 的时间 + 从 j 到 i 的时间
				res = min(res, dfs(mask, j)+abs(x-r[1]))
			}
		}
		// 处理完请求 i 的时间不能早于 t
		res = max(res, t)

		*p = res // 记忆化
		return res
	}

	ans := math.MaxInt
	for i := range m { // 枚举最后处理的请求
		ans = min(ans, dfs(1<<m-1, i))
	}
	return int64(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

### 1:1 翻译成递推

```py [sol-Python3]
class Solution:
    def elevatorRequests(self, n: int, start: int, requests: list[list[int]]) -> int:
        m = len(requests)
        f = [[0] * m for _ in range(1 << m)]

        for i, (t, x) in enumerate(requests):
            f[1 << i][i] = max(abs(x - start), t)

        for mask in range(1, 1 << m):
            if mask & (mask - 1) == 0:  # mask 只有一个元素
                continue
            for i, (t, x) in enumerate(requests):
                if mask >> i & 1 == 0:
                    continue
                msk = mask ^ (1 << i)
                res = min(f[msk][j] + abs(x - y) for j, (_, y) in enumerate(requests) if msk >> j & 1)
                f[mask][i] = max(res, t)

        return min(f[-1])
```

```java [sol-Java]
class Solution {
    public long elevatorRequests(int n, int start, int[][] requests) {
        int m = requests.length;
        long[][] f = new long[1 << m][m];
        for (int i = 0; i < m; i++) {
            int[] req = requests[i];
            f[1 << i][i] = Math.max(Math.abs(req[1] - start), req[0]);
        }

        for (int mask = 1; mask < 1 << m; mask++) {
            if ((mask & (mask - 1)) == 0) { // mask 只有一个元素
                continue;
            }
            for (int i = 0; i < m; i++) {
                if ((mask >> i & 1) == 0) {
                    continue;
                }
                long res = Long.MAX_VALUE;
                int msk = mask ^ (1 << i);
                int x = requests[i][1];
                for (int j = 0; j < m; j++) {
                    if ((msk >> j & 1) > 0) {
                        res = Math.min(res, f[msk][j] + Math.abs(x - requests[j][1]));
                    }
                }
                f[mask][i] = Math.max(res, requests[i][0]);
            }
        }

        long ans = Long.MAX_VALUE;
        for (long x : f[(1 << m) - 1]) {
            ans = Math.min(ans, x);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long elevatorRequests(int n, int start, vector<vector<int>>& requests) {
        int m = requests.size();
        vector f(1 << m, vector<long long>(m));
        for (int i = 0; i < m; i++) {
            auto& req = requests[i];
            f[1 << i][i] = max(abs(req[1] - start), req[0]);
        }

        for (int mask = 1; mask < (1 << m); mask++) {
            if ((mask & (mask - 1)) == 0) { // mask 只有一个元素
                continue;
            }
            for (int i = 0; i < m; i++) {
                if ((mask >> i & 1) == 0) {
                    continue;
                }
                auto& req = requests[i];
                int t = req[0], x = req[1];
                long long res = LLONG_MAX;
                int msk = mask ^ (1 << i);
                for (int j = 0; j < m; j++) {
                    if (msk >> j & 1) {
                        res = min(res, f[msk][j] + abs(x - requests[j][1]));
                    }
                }
                f[mask][i] = max(res, 1LL * t);
            }
        }

        return ranges::min(f.back());
    }
};
```

```go [sol-Go]
func elevatorRequests(n int, start int, requests [][]int) int64 {
	m := len(requests)
	f := make([][]int, 1<<m)
	for i := range f {
		f[i] = make([]int, m)
	}

	for i, req := range requests {
		f[1<<i][i] = max(abs(req[1]-start), req[0])
	}

	for mask := 1; mask < 1<<m; mask++ {
		if mask&(mask-1) == 0 { // mask 只有一个元素
			continue
		}
		for i, req := range requests {
			if mask>>i&1 == 0 {
				continue
			}
			res := math.MaxInt
			msk := mask ^ 1<<i
			x := req[1]
			for j, r := range requests {
				if msk>>j&1 > 0 {
					res = min(res, f[msk][j]+abs(x-r[1]))
				}
			}
			f[mask][i] = max(res, req[0])
		}
	}

	return int64(slices.Min(f[1<<m-1]))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m^22^m)$，其中 $m$ 是 $\textit{requests}$ 的长度。
- 空间复杂度：$\mathcal{O}(m2^m)$。

> **注**：也可以转化成最短路模型，用 Dijkstra 算法解决。

## 方法二：区间 DP

考察电梯移动路径的逆过程：

- 电梯从某个楼层 $\textit{floor}_i$ 出发，依次处理所有请求，最后回到 $\textit{start}$。

例如原问题的答案为 $10$，某个请求的发起时间为 $\textit{arrival}_i=3$。从逆过程的角度看，我们从时刻 $10$ 开始倒流，必须在 $\ge 3$ 时刻完成该请求。

如果觉得逆向时间不好理解，也可以理解成从时刻 $0$ 开始，必须在 $\le 10-3=7$ 时刻完成该请求。换句话说，请求的「到达时间」变成了「**截止时间**」。

设 $a,b,c$ 是三个楼层，且 $a < b < c$。在从楼层 $a$ 移动到楼层 $c$ 的过程中，一定会经过楼层 $b$：

- 如果到达楼层 $b$ 时，已经超过 $b$ 的截止时间，那么移动失败，这不是一个合法的移动方案。
- 如果到达楼层 $b$ 时，尚未超过 $b$ 的截止时间，那么可以**顺带完成**楼层 $b$ 的请求。

由此可以看出，本题的逆过程和 [4023. 电梯请求 II](https://leetcode.cn/problems/elevator-requests-ii/) 是一样的。设逆过程中电梯到达过的最小楼层为 $p$，最大楼层为 $q$，那么电梯也一定到达过 $[p,q]$ 中的楼层，我们也顺带处理了相应的请求。所以我们只需考虑 $p$ 下面的最近请求楼层，以及 $q$ 上面的最近请求楼层，无需枚举所有请求楼层。并且，电梯移动后，要么位于最小楼层，要么位于最大楼层。

这样就证明了：把请求按楼层排序后，逆过程的已完成区间是连续的，即**正向过程的未完成区间是连续的**。

类似 4023 题，先插入 $-1$ 和 $n$ 两个哨兵楼层（不含哨兵的下标范围是 $[1,m-2]$），然后写一个区间 DP，定义：

- $\textit{dfs}(i,j,\texttt{false})$ 表示完成请求 $[1,i] \cup [j+1,m-2]$ 所需的最短时间，且最后一个完成的请求是 $i$。
- $\textit{dfs}(i,j,\texttt{true})$ 表示完成请求 $[1,i-1] \cup [j,m-2]$ 所需的最短时间，且最后一个完成的请求是 $j$。

设当前请求的发起时间为 $t$，楼层为 $x$。枚举上一个完成的请求是 $i-1$ 还是 $j+1$：

- 如果是 $i-1$，那么问题变成完成请求 $[1,i-1] \cup [j+1,m-2]$ 所需的最短时间，且最后一个完成的请求是 $i-1$，即 $\textit{dfs}(i-1,j,\texttt{false})$，加上从 $\textit{floor}_{i-1}$ 到当前楼层 $x$ 的距离，再与 $t$ 取最大值。
- 如果是 $j+1$，那么问题变成完成请求 $[1,i-1] \cup [j+1,m-2]$ 所需的最短时间，且最后一个完成的请求是 $j+1$，即 $\textit{dfs}(i,j+1,\texttt{true})$，加上从 $\textit{floor}_{j+1}$ 到当前楼层 $x$ 的距离，再与 $t$ 取最大值。

两种情况取最小值，即 $\textit{dfs}(i,j,\textit{isRight})$。

**递归边界**：

- 如果 $i=0$ 或者 $j=m-1$，出界，不合法，返回 $\infty$。
- 如果 $i=1$ 且 $j=m-2$，则当前请求是第一个请求，返回 $\max(|x - \textit{start}|, t)$。

**递归入口**：枚举最后一个完成的请求是 $i$，用 $\textit{dfs}(i,i,\texttt{false})$ 更新答案的最小值。这里写 $\texttt{false}$ 还是 $\texttt{true}$ 都可以，是一样的。

### 记忆化搜索

```py [sol-Python3]
class Solution:
    def elevatorRequests(self, n: int, start: int, requests: list[int]) -> int:
        requests += [[0, -1], [0, n]]  # 插入两个哨兵
        requests.sort(key=lambda r: r[1])  # 按楼层排序
        m = len(requests)  # 不含哨兵的下标范围是 [1, m-2]

        # dfs(i, j, False) 返回完成请求 [1,i] ∪ [j+1,m-2] 所需的最短时间，此时电梯在 floor[i]（最后一个完成的请求是 i）
        # dfs(i, j, True)  返回完成请求 [1,i-1] ∪ [j,m-2] 所需的最短时间，此时电梯在 floor[j]（最后一个完成的请求是 j）
        @cache  # 缓存装饰器，避免重复计算 dfs（一行代码实现记忆化）
        def dfs(i: int, j: int, is_right: bool) -> int:
            if i == 0 or j == m - 1:  # 出界
                return inf
            t, x = requests[j if is_right else i]  # 当前请求
            if i == 1 and j == m - 2:  # 当前请求是第一个请求
                return max(abs(x - start), t)  # 从 start 到当前楼层
            return min(max(dfs(i - 1, j, False) + x - requests[i - 1][1], t),  # 从 floor[i-1] 到当前楼层
                       max(dfs(i, j + 1, True) + requests[j + 1][1] - x, t))   # 从 floor[j+1] 到当前楼层

        # 枚举最后一个完成的请求
        return min(dfs(i, i, False) for i in range(1, m - 1))
```

```java [sol-Java]
class Solution {
    public long elevatorRequests(int n, int start, int[][] requests) {
        int m = requests.length + 2; // 不含哨兵的下标范围是 [1, m-2]
        int[][] a = Arrays.copyOf(requests, m);
        a[m - 2] = new int[]{0, -1};
        a[m - 1] = new int[]{0, n}; // 插入两个哨兵
        Arrays.sort(a, (p, q) -> p[1] - q[1]); // 按楼层排序

        long[][][] memo = new long[m][m][2];
        for (long[][] mat : memo) {
            for (long[] row : mat) {
                Arrays.fill(row, -1); // -1 表示该状态没有计算过
            }
        }

        // 枚举最后一个完成的请求
        long ans = Long.MAX_VALUE;
        for (int i = 1; i < m - 1; i++) {
            ans = Math.min(ans, dfs(i, i, 0, start, a, memo)); // 这里 0 和 1 是一样的
        }
        return ans;
    }

    // dfs(i, j, 0) 返回完成请求 [1,i] ∪ [j+1,m-2] 所需的最短时间，此时电梯在 floor[i]（最后一个完成的请求是 i）
    // dfs(i, j, 1) 返回完成请求 [1,i-1] ∪ [j,m-2] 所需的最短时间，此时电梯在 floor[j]（最后一个完成的请求是 j）
    private long dfs(int i, int j, int isRight, int start, int[][] requests, long[][][] memo) {
        int m = requests.length;
        if (i == 0 || j == m - 1) { // 出界
            return Long.MAX_VALUE / 2;
        }

        long res = memo[i][j][isRight];
        if (res != -1) {
            return res;
        }

        int[] req = requests[isRight > 0 ? j : i];
        int t = req[0];
        int x = req[1];
        if (i == 1 && j == m - 2) { // 当前请求是第一个请求
            res = Math.max(Math.abs(x - start), t); // 从 start 到当前楼层
        } else {
            res = Math.min(Math.max(dfs(i - 1, j, 0, start, requests, memo) + x - requests[i - 1][1], t),  // 从 floor[i-1] 到当前楼层
                           Math.max(dfs(i, j + 1, 1, start, requests, memo) + requests[j + 1][1] - x, t)); // 从 floor[j+1] 到当前楼层
        }

        memo[i][j][isRight] = res; // 记忆化
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long elevatorRequests(int n, int start, vector<vector<int>>& requests) {
        requests.insert(requests.end(), {{0, -1}, {0, n}}); // 插入两个哨兵
        ranges::sort(requests, {}, [](auto& r) { return r[1]; }); // 按楼层排序
        int m = requests.size(); // 不含哨兵的下标范围是 [1, m-2]
        vector memo(m, vector(m, array<long long, 2>{-1, -1})); // -1 表示该状态没有计算过

        // dfs(i, j, false) 返回完成请求 [1,i] ∪ [j+1,m-2] 所需的最短时间，此时电梯在 floor[i]（最后一个完成的请求是 i）
        // dfs(i, j, true)  返回完成请求 [1,i-1] ∪ [j,m-2] 所需的最短时间，此时电梯在 floor[j]（最后一个完成的请求是 j）
        auto dfs = [&](this auto&& dfs, int i, int j, bool is_right) -> long long {
            if (i == 0 || j == m - 1) { // 出界
                return LLONG_MAX / 2;
            }
            long long& res = memo[i][j][is_right]; // 注意这里是引用
            if (res != -1) {
                return res;
            }

            auto& req = requests[is_right ? j : i];
            int t = req[0], x = req[1];
            if (i == 1 && j == m - 2) { // 当前请求是第一个请求
                return res = max(abs(x - start), t); // 从 start 到当前楼层
            }
            return res = min(max(dfs(i - 1, j, false) + x - requests[i - 1][1], 1LL * t), // 从 floor[i-1] 到当前楼层
                             max(dfs(i, j + 1, true) + requests[j + 1][1] - x, 1LL * t)); // 从 floor[j+1] 到当前楼层
        };

        // 枚举最后一个完成的请求
        long long ans = LLONG_MAX;
        for (int i = 1; i < m - 1; i++) {
            ans = min(ans, dfs(i, i, false)); // 这里 false 和 true 是一样的
        }
        return ans;
    }
};
```

```go [sol-Go]
func elevatorRequests(n int, start int, requests [][]int) int64 {
	requests = append(requests, []int{0, -1}, []int{0, n}) // 插入两个哨兵
	slices.SortFunc(requests, func(a, b []int) int { return a[1] - b[1] }) // 按楼层排序
	m := len(requests) // 不含哨兵的下标范围是 [1, m-2]

	memo := make([][][2]int, m)
	for i := range memo {
		memo[i] = make([][2]int, m)
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1} // -1 表示该状态没有计算过
		}
	}

	// dfs(i, j, false) 返回完成请求 [1,i] ∪ [j+1,m-2] 所需的最短时间，此时电梯在 floor[i]（最后一个完成的请求是 i）
	// dfs(i, j, true)  返回完成请求 [1,i-1] ∪ [j,m-2] 所需的最短时间，此时电梯在 floor[j]（最后一个完成的请求是 j）
	var dfs func(int, int, uint8) int
	dfs = func(i, j int, isRight uint8) int {
		if i == 0 || j == m-1 { // 出界
			return math.MaxInt / 2
		}

		p := &memo[i][j][isRight]
		if *p != -1 {
			return *p
		}

		k := i
		if isRight > 0 {
			k = j
		}
		t, x := requests[k][0], requests[k][1]
		if i == 1 && j == m-2 { // 当前请求是第一个请求
			*p = max(abs(x-start), t) // 从 start 到当前楼层
		} else {
			*p = min(max(dfs(i-1, j, 0)+x-requests[i-1][1], t), // 从 floor[i-1] 到当前楼层
					max(dfs(i, j+1, 1)+requests[j+1][1]-x, t)) // 从 floor[j+1] 到当前楼层
		}
		return *p
	}

	ans := math.MaxInt
	// 枚举最后一个完成的请求
	for i := 1; i < m-1; i++ {
		ans = min(ans, dfs(i, i, 0))
	}
	return int64(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

### 1:1 翻译成递推

```py [sol-Python3]
class Solution:
    def elevatorRequests(self, n: int, start: int, requests: list[list[int]]) -> int:
        requests += [[0, -1], [0, n]]  # 插入两个哨兵
        requests.sort(key=lambda r: r[1])  # 按楼层排序
        m = len(requests)  # 不含哨兵的下标范围是 [1, m-2]
        f = [[[inf, inf] for _ in range(m)] for _ in range(m)]

        for i in range(1, m - 1):
            t, x = requests[i]
            for j in range(m - 2, i - 1, -1):
                t2, y = requests[j]
                if i == 1 and j == m - 2:  # 当前请求是第一个请求
                    # 从 start 到当前楼层
                    f[i][j][0] = max(abs(x - start), t)
                    f[i][j][1] = max(abs(y - start), t2)
                    continue
                f[i][j][0] = min(max(f[i - 1][j][0] + x - requests[i - 1][1], t),  # 从 floor[i-1] 到当前楼层
                                 max(f[i][j + 1][1] + requests[j + 1][1] - x, t))  # 从 floor[j+1] 到当前楼层
                f[i][j][1] = min(max(f[i - 1][j][0] + y - requests[i - 1][1], t2),  # 从 floor[i-1] 到当前楼层
                                 max(f[i][j + 1][1] + requests[j + 1][1] - y, t2))  # 从 floor[j+1] 到当前楼层

        # 枚举最后一个完成的请求
        return min(f[i][i][0] for i in range(1, m - 1))
```

```java [sol-Java]
class Solution {
    public long elevatorRequests(int n, int start, int[][] requests) {
        int m = requests.length + 2; // 不含哨兵的下标范围是 [1, m-2]
        int[][] a = Arrays.copyOf(requests, m);
        a[m - 2] = new int[]{0, -1};
        a[m - 1] = new int[]{0, n}; // 插入两个哨兵
        Arrays.sort(a, (p, q) -> p[1] - q[1]); // 按楼层排序

        long[][][] f = new long[m][m][2];
        for (long[] row : f[0]) {
            Arrays.fill(row, Long.MAX_VALUE / 2);
        }

        for (int i = 1; i < m - 1; i++) {
            f[i][m - 1][0] = f[i][m - 1][1] = Long.MAX_VALUE / 2;
            int t = a[i][0];
            int x = a[i][1];
            for (int j = m - 2; j >= i; j--) {
                int t2 = a[j][0];
                int y = a[j][1];
                if (i == 1 && j == m - 2) { // 当前请求是第一个请求
                    // 从 start 到当前楼层
                    f[i][j][0] = Math.max(Math.abs(x - start), t);
                    f[i][j][1] = Math.max(Math.abs(y - start), t2);
                    continue;
                }
                f[i][j][0] = Math.min(Math.max(f[i - 1][j][0] + x - a[i - 1][1], t),  // 从 floor[i-1] 到当前楼层
                                      Math.max(f[i][j + 1][1] + a[j + 1][1] - x, t)); // 从 floor[j+1] 到当前楼层
                f[i][j][1] = Math.min(Math.max(f[i - 1][j][0] + y - a[i - 1][1], t2),  // 从 floor[i-1] 到当前楼层
                                      Math.max(f[i][j + 1][1] + a[j + 1][1] - y, t2)); // 从 floor[j+1] 到当前楼层
            }
        }

        // 枚举最后一个完成的请求
        long ans = Long.MAX_VALUE;
        for (int i = 1; i < m - 1; i++) {
            ans = Math.min(ans, f[i][i][0]);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long elevatorRequests(int n, int start, vector<vector<int>>& requests) {
        requests.insert(requests.end(), {{0, -1}, {0, n}}); // 插入两个哨兵
        ranges::sort(requests, {}, [](auto& r) { return r[1]; }); // 按楼层排序
        int m = requests.size(); // 不含哨兵的下标范围是 [1, m-2]
        vector f(m, vector(m, array<long long, 2>{LLONG_MAX / 2, LLONG_MAX / 2}));

        for (int i = 1; i < m - 1; i++) {
            int t = requests[i][0];
            int x = requests[i][1];
            for (int j = m - 2; j >= i; j--) {
                int t2 = requests[j][0];
                int y = requests[j][1];
                if (i == 1 && j == m - 2) { // 当前请求是第一个请求
                    // 从 start 到当前楼层
                    f[i][j][0] = max(abs(x - start), t);
                    f[i][j][1] = max(abs(y - start), t2);
                    continue;
                }
                f[i][j][0] = min(max(f[i - 1][j][0] + x - requests[i - 1][1], 1LL * t),  // 从 floor[i-1] 到当前楼层
                                 max(f[i][j + 1][1] + requests[j + 1][1] - x, 1LL * t)); // 从 floor[j+1] 到当前楼层
                f[i][j][1] = min(max(f[i - 1][j][0] + y - requests[i - 1][1], 1LL * t2),  // 从 floor[i-1] 到当前楼层
                                 max(f[i][j + 1][1] + requests[j + 1][1] - y, 1LL * t2)); // 从 floor[j+1] 到当前楼层
            }
        }

        // 枚举最后一个完成的请求
        long long ans = LLONG_MAX;
        for (int i = 1; i < m - 1; i++) {
            ans = min(ans, f[i][i][0]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func elevatorRequests(n int, start int, requests [][]int) int64 {
	requests = append(requests, []int{0, -1}, []int{0, n}) // 插入两个哨兵
	slices.SortFunc(requests, func(a, b []int) int { return a[1] - b[1] }) // 按楼层排序
	m := len(requests)// 不含哨兵的下标范围是 [1, m-2]

	f := make([][][2]int, m)
	for i := range f {
		f[i] = make([][2]int, m)
	}
	for j := range f[0] {
		f[0][j] = [2]int{math.MaxInt / 2, math.MaxInt / 2}
	}

	for i := 1; i < m-1; i++ {
		f[i][m-1] = [2]int{math.MaxInt / 2, math.MaxInt / 2}
		t, x := requests[i][0], requests[i][1]
		for j := m - 2; j >= i; j-- {
			t2, y := requests[j][0], requests[j][1]
			if i == 1 && j == m-2 { // 当前请求是第一个请求
				// 从 start 到当前楼层
				f[i][j][0] = max(abs(x-start), t)
				f[i][j][1] = max(abs(y-start), t2)
				continue
			}
			f[i][j][0] = min(max(f[i-1][j][0]+x-requests[i-1][1], t), // 从 floor[i-1] 到当前楼层
							max(f[i][j+1][1]+requests[j+1][1]-x, t)) // 从 floor[j+1] 到当前楼层
			f[i][j][1] = min(max(f[i-1][j][0]+y-requests[i-1][1], t2), // 从 floor[i-1] 到当前楼层
							max(f[i][j+1][1]+requests[j+1][1]-y, t2)) // 从 floor[j+1] 到当前楼层
		}
	}

	// 枚举最后一个完成的请求
	ans := math.MaxInt
	for i := 1; i < m-1; i++ {
		ans = min(ans, f[i][i][0])
	}
	return int64(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

### 空间优化

```py [sol-Python3]
class Solution:
    def elevatorRequests(self, n: int, start: int, requests: list[list[int]]) -> int:
        requests += [[0, -1], [0, n]]  # 插入两个哨兵
        requests.sort(key=lambda r: r[1])  # 按楼层排序
        m = len(requests)  # 不含哨兵的下标范围是 [1, m-2]
        f = [[inf, inf] for _ in range(m)]

        for i in range(1, m - 1):
            t, x = requests[i]
            for j in range(m - 2, i - 1, -1):
                t2, y = requests[j]
                if i == 1 and j == m - 2:  # 当前请求是第一个请求
                    # 从 start 到当前楼层
                    f[j][0] = max(abs(x - start), t)
                    f[j][1] = max(abs(y - start), t2)
                    continue
                f[j][1] = min(max(f[j][0] + y - requests[i - 1][1], t2),  # 从 floor[i-1] 到当前楼层
                              max(f[j + 1][1] + requests[j + 1][1] - y, t2))  # 从 floor[j+1] 到当前楼层
                f[j][0] = min(max(f[j][0] + x - requests[i - 1][1], t),  # 从 floor[i-1] 到当前楼层
                              max(f[j + 1][1] + requests[j + 1][1] - x, t))  # 从 floor[j+1] 到当前楼层

        # 枚举最后一个完成的请求
        return min(f[i][0] for i in range(1, m - 1))
```

```java [sol-Java]
class Solution {
    public long elevatorRequests(int n, int start, int[][] requests) {
        int m = requests.length + 2; // 不含哨兵的下标范围是 [1, m-2]
        int[][] a = Arrays.copyOf(requests, m);
        a[m - 2] = new int[]{0, -1};
        a[m - 1] = new int[]{0, n}; // 插入两个哨兵
        Arrays.sort(a, (p, q) -> p[1] - q[1]); // 按楼层排序

        long[][] f = new long[m][2];
        for (long[] row : f) {
            Arrays.fill(row, Long.MAX_VALUE / 2);
        }

        for (int i = 1; i < m - 1; i++) {
            int t = a[i][0];
            int x = a[i][1];
            for (int j = m - 2; j >= i; j--) {
                int t2 = a[j][0];
                int y = a[j][1];
                if (i == 1 && j == m - 2) { // 当前请求是第一个请求
                    // 从 start 到当前楼层
                    f[j][0] = Math.max(Math.abs(x - start), t);
                    f[j][1] = Math.max(Math.abs(y - start), t2);
                    continue;
                }
                f[j][1] = Math.min(Math.max(f[j][0] + y - a[i - 1][1], t2), // 从 floor[i-1] 到当前楼层
                                   Math.max(f[j + 1][1] + a[j + 1][1] - y, t2)); // 从 floor[j+1] 到当前楼层
                f[j][0] = Math.min(Math.max(f[j][0] + x - a[i - 1][1], t), // 从 floor[i-1] 到当前楼层
                                   Math.max(f[j + 1][1] + a[j + 1][1] - x, t)); // 从 floor[j+1] 到当前楼层
            }
        }

        // 枚举最后一个完成的请求
        long ans = Long.MAX_VALUE;
        for (int i = 1; i < m - 1; i++) {
            ans = Math.min(ans, f[i][0]);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long elevatorRequests(int n, int start, vector<vector<int>>& requests) {
        requests.insert(requests.end(), {{0, -1}, {0, n}}); // 插入两个哨兵
        ranges::sort(requests, {}, [](auto& r) { return r[1]; }); // 按楼层排序
        int m = requests.size(); // 不含哨兵的下标范围是 [1, m-2]
        vector f(m, array<long long, 2>{LLONG_MAX / 2, LLONG_MAX / 2});

        for (int i = 1; i < m - 1; i++) {
            int t = requests[i][0];
            int x = requests[i][1];
            for (int j = m - 2; j >= i; j--) {
                int t2 = requests[j][0];
                int y = requests[j][1];
                if (i == 1 && j == m - 2) { // 当前请求是第一个请求
                    // 从 start 到当前楼层
                    f[j][0] = max(abs(x - start), t);
                    f[j][1] = max(abs(y - start), t2);
                    continue;
                }
                f[j][1] = min(max(f[j][0] + y - requests[i - 1][1], 1LL * t2), // 从 floor[i-1] 到当前楼层
                              max(f[j + 1][1] + requests[j + 1][1] - y, 1LL * t2)); // 从 floor[j+1] 到当前楼层
                f[j][0] = min(max(f[j][0] + x - requests[i - 1][1], 1LL * t), // 从 floor[i-1] 到当前楼层
                              max(f[j + 1][1] + requests[j + 1][1] - x, 1LL * t)); // 从 floor[j+1] 到当前楼层
            }
        }

        // 枚举最后一个完成的请求
        long long ans = LLONG_MAX;
        for (int i = 1; i < m - 1; i++) {
            ans = min(ans, f[i][0]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func elevatorRequests(n int, start int, requests [][]int) int64 {
	requests = append(requests, []int{0, -1}, []int{0, n}) // 插入两个哨兵
	slices.SortFunc(requests, func(a, b []int) int { return a[1] - b[1] }) // 按楼层排序
	m := len(requests) // 不含哨兵的下标范围是 [1, m-2]

	f := make([][2]int, m)
	for j := range f {
		f[j] = [2]int{math.MaxInt / 2, math.MaxInt / 2}
	}

	for i := 1; i < m-1; i++ {
		t, x := requests[i][0], requests[i][1]
		for j := m - 2; j >= i; j-- {
			t2, y := requests[j][0], requests[j][1]
			if i == 1 && j == m-2 { // 当前请求是第一个请求
				// 从 start 到当前楼层
				f[j][0] = max(abs(x-start), t)
				f[j][1] = max(abs(y-start), t2)
				continue
			}
			f[j][1] = min(max(f[j][0]+y-requests[i-1][1], t2), // 从 floor[i-1] 到当前楼层
						max(f[j+1][1]+requests[j+1][1]-y, t2)) // 从 floor[j+1] 到当前楼层
			f[j][0] = min(max(f[j][0]+x-requests[i-1][1], t), // 从 floor[i-1] 到当前楼层
						max(f[j+1][1]+requests[j+1][1]-x, t)) // 从 floor[j+1] 到当前楼层
		}
	}

	// 枚举最后一个完成的请求
	ans := math.MaxInt
	for i := 1; i < m-1; i++ {
		ans = min(ans, f[i][0])
	}
	return int64(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m^2)$，其中 $m$ 是 $\textit{requests}$ 的长度。
- 空间复杂度：$\mathcal{O}(m)$。

## 专题训练

1. 动态规划题单的「**§9.2 排列型状压 DP ② 相邻相关**」「**§9.3 旅行商问题（TSP）**」和「**八、区间 DP**」。
2. 图论题单的「**§3.1 单源最短路：Dijkstra 算法**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
