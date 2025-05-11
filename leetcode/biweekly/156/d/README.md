树形 DP：从 O(nk) 到 O(n)（Python/Java/C++/Go）

---

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

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def subtreeInversionSum(self, edges: List[List[int]], nums: List[int], k: int) -> int:
        n = len(nums)
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        memo = {}

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

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(nk)$。

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
