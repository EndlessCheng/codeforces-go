如果你从未做过状态压缩 DP（状压 DP），推荐先阅读 [教你一步步思考状压 DP：从记忆化搜索到递推](https://leetcode.cn/problems/beautiful-arrangement/solution/jiao-ni-yi-bu-bu-si-kao-zhuang-ya-dpcong-c6kd/)。

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

## 写法一：记忆化搜索

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

## 写法二：1:1 翻译成递推

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

**注**：本题也可以转化成最短路模型，用 Dijkstra 算法解决。

## 专题训练

1. 动态规划题单的「**§9.2 排列型状压 DP ② 相邻相关**」和「**§9.3 旅行商问题（TSP）**」。
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
