直接从左到右暴搜回文路径？写记忆化？不幸的是，记忆化需要保存后续路径的完整顺序信息（因为要判断是否回文）。考虑完全图，本质在枚举所有排列，有 $\mathcal{O}(n!)$ 条不同的路径，这太多了。

**核心思路**：回文串问题，可以**枚举回文中心**，从中心向左右两边扩展（中心扩展法）。

设访问过的节点集合为 $S$。对于已经访问过的点，其访问顺序是不重要的，我们只需要知道访问过哪些点。比如回文路径的左半边扩展到节点 $3$ 时，怎么来的不重要，无论是 $1\to 2 \to 3$ 还是 $2\to 1 \to 3$，都等同于回文路径的左半边已经包含 $1,2,3$ 这三个节点了，且左半边当前在节点 $3$（重叠子问题）。这样就无需暴力枚举节点访问顺序的排列了，只需维护节点无序集合的信息，可以用状压 DP 解决。

我们需要知道三个信息：

- 路径的左右端点 $x$ 和 $y$。
- 已经访问过的节点（包括 $x$ 和 $y$）集合 $S$。

定义 $\textit{dfs}(x,y,S)$ 表示在路径的左右端点为 $x$ 和 $y$，访问过的节点集合为 $S$ 的情况下，从 $x$ 和 $y$ 向两侧扩展，最多还能访问多少个节点（不算 $x$ 和 $y$）。

枚举 $x$ 的邻居 $v$，枚举 $y$ 的邻居 $w$，如果 $v$ 和 $w$ 都没访问过，且 $v\ne w$ 且 $\textit{label}[v] = \textit{label}[w]$，那么可以扩展，问题变成在路径的左右端点为 $v$ 和 $w$，访问过的节点集合为 $S \cup \{v,w\}$ 的情况下，从 $v$ 和 $w$ 向两侧扩展，最多还能访问多少个节点，即 $\textit{dfs}(v,w, S \cup \{v,w\})$，用该返回值加二（加上 $v$ 和 $w$），更新 $\textit{dfs}(x,y,S)$ 的返回值的最大值，即

$$
\textit{dfs}(x,y,S) = \max_{v,w} \textit{dfs}(v,w, S \cup \{v,w\}) + 2
$$

其中 $v$ 是 $x$ 的邻居，$w$ 是 $y$ 的邻居，$v\notin S$，$w\notin S$，$v\ne w$，$\textit{label}[v] = \textit{label}[w]$。

**递归边界**：无需判断。

**递归入口**：

- 奇回文串：$\textit{dfs}(x,x,\{x\})+1$。
- 偶回文串：$\textit{dfs}(x,y,\{x,y\})+2$。其中 $x$ 和 $y$ 是邻居且 $\textit{label}[x] = \textit{label}[y]$。

代码实现时，用二进制表示集合，用位运算实现集合操作，具体请看 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

**优化 1**：我们计算的是从 $x$ 和 $y$ 出发继续扩展的节点个数，根据对称性，$\textit{dfs}(x,y,S)$ 计算出的结果和 $\textit{dfs}(y,x,S)$ 计算出的结果是一样的，没必要算两次。所以递归时，可以人为规定递归参数必须满足 $x\le y$，从而减少状态个数和计算量。如果 $x>y$ 则交换。

**优化 2**：递归结束后，如果 $\textit{ans} = n$，可以直接返回 $n$。

**优化 3**：特判完全图的情况，此时路径可以是任意节点的排列，问题等价于重排 $\textit{label}$ 中的字母可以得到的最长回文串。比如 $3$ 个 $\texttt{a}$ 和 $5$ 个 $\texttt{b}$，可以选 $2$ 个 $\texttt{a}$ 一左一右，$4$ 个 $\texttt{b}$ 左右各放 $2$ 个，多出的字母只能选一个放正中间。比如回文串为 $\texttt{abbabba}$。（谢谢 [@观铃 🔔](/u/kamio_misuzu) 补充）

具体请看 [视频讲解](https://www.bilibili.com/video/BV1xSuFzHEa1/?t=28m51s)，欢迎点赞关注~

```py [sol-Python3]
# 手写 max 更快
max = lambda a, b: b if b > a else a

class Solution:
    def maxLen(self, n: int, edges: List[List[int]], label: str) -> int:
        if len(edges) == n * (n - 1) // 2:  # 完全图
            ans = odd = 0
            for c in Counter(label).values():
                ans += c - c % 2
                odd |= c % 2
            return ans + odd

        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        # 计算从 x 和 y 向两侧扩展，最多还能访问多少个节点（不算 x 和 y）
        @cache
        def dfs(x: int, y: int, vis: int) -> int:
            res = 0
            for v in g[x]:
                if vis >> v & 1:
                    continue
                for w in g[y]:
                    if vis >> w & 1 == 0 and v != w and label[w] == label[v]:
                        tv, tw = v, w  # 注意不能直接交换 v 和 w，否则下个循环的 v 就不是原来的 v 了
                        if tv > tw:  # 保证 tv < tw，减少状态个数和计算量
                            tv, tw = tw, tv
                        res = max(res, dfs(tv, tw, vis | 1 << v | 1 << w) + 2)
            return res

        ans = 0
        for x, to in enumerate(g):
            # 奇回文串，x 作为回文中心
            ans = max(ans, dfs(x, x, 1 << x) + 1)
            if ans == n:
                return n
            # 偶回文串，x 和 x 的邻居 y 作为回文中心
            for y in to:
                # 保证递归参数 x < y，减少状态个数和计算量
                if x < y and label[x] == label[y]:
                    ans = max(ans, dfs(x, y, 1 << x | 1 << y) + 2)
                    if ans == n:
                        return n
        return ans
```

```java [sol-Java]
class Solution {
    public int maxLen(int n, int[][] edges, String label) {
        char[] s = label.toCharArray();
        if (edges.length == n * (n - 1) / 2) { // 完全图
            int[] cnt = new int[26];
            for (char ch : s) {
                cnt[ch - 'a']++;
            }
            int ans = 0, odd = 0;
            for (int c : cnt) {
                ans += c - c % 2;
                odd |= c % 2;
            }
            return ans + odd;
        }

        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        int[][][] memo = new int[n][n][1 << n];
        for (int[][] mat : memo) {
            for (int[] row : mat) {
                Arrays.fill(row, -1);
            }
        }

        int ans = 0;
        for (int x = 0; x < n; x++) {
            // 奇回文串，x 作为回文中心
            ans = Math.max(ans, dfs(x, x, 1 << x, g, s, memo) + 1);
            if (ans == n) {
                return n;
            }
            // 偶回文串，x 和 x 的邻居 y 作为回文中心
            for (int y : g[x]) {
                // 保证 x < y，减少状态个数和计算量
                if (x < y && s[x] == s[y]) {
                    ans = Math.max(ans, dfs(x, y, 1 << x | 1 << y, g, s, memo) + 2);
                    if (ans == n) {
                        return n;
                    }
                }
            }
        }
        return ans;
    }

    // 计算从 x 和 y 向两侧扩展，最多还能访问多少个节点（不算 x 和 y）
    private int dfs(int x, int y, int vis, List<Integer>[] g, char[] label, int[][][] memo) {
        if (memo[x][y][vis] >= 0) { // 之前计算过
            return memo[x][y][vis];
        }
        int res = 0;
        for (int v : g[x]) {
            if ((vis >> v & 1) > 0) { // v 在路径中
                continue;
            }
            for (int w : g[y]) {
                if ((vis >> w & 1) == 0 && w != v && label[w] == label[v]) {
                    // 保证 v < w，减少状态个数和计算量
                    int r = dfs(Math.min(v, w), Math.max(v, w), vis | 1 << v | 1 << w, g, label, memo);
                    res = Math.max(res, r + 2);
                }
            }
        }
        return memo[x][y][vis] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxLen(int n, vector<vector<int>>& edges, string label) {
        if (edges.size() == n * (n - 1) / 2) { // 完全图
            int cnt[26]{};
            for (char ch : label) {
                cnt[ch - 'a']++;
            }
            int ans = 0, odd = 0;
            for (int c : cnt) {
                ans += c - c % 2;
                odd |= c % 2;
            }
            return ans + odd;
        }

        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        vector memo(n, vector(n, vector<int>(1 << n, -1)));
        // 计算从 x 和 y 向两侧扩展，最多还能访问多少个节点（不算 x 和 y）
        auto dfs = [&](this auto&& dfs, int x, int y, int vis) -> int {
            int& res = memo[x][y][vis]; // 注意这里是引用
            if (res >= 0) { // 之前计算过
                return res;
            }
            res = 0;
            for (int v : g[x]) {
                if (vis >> v & 1) { // v 在路径中
                    continue;
                }
                for (int w : g[y]) {
                    if ((vis >> w & 1) == 0 && w != v && label[w] == label[v]) {
                        // 保证 v < w，减少状态个数和计算量
                        int r = dfs(min(v, w), max(v, w), vis | 1 << v | 1 << w);
                        res = max(res, r + 2);
                    }
                }
            }
            return res;
        };

        int ans = 0;
        for (int x = 0; x < n; x++) {
            // 奇回文串，x 作为回文中心
            ans = max(ans, dfs(x, x, 1 << x) + 1);
            if (ans == n) {
                return n;
            }
            // 偶回文串，x 和 x 的邻居 y 作为回文中心
            for (int y : g[x]) {
                // 保证 x < y，减少状态个数和计算量
                if (x < y && label[x] == label[y]) {
                    ans = max(ans, dfs(x, y, 1 << x | 1 << y) + 2);
                    if (ans == n) {
                        return n;
                    }
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxLen(n int, edges [][]int, label string) (ans int) {
	if len(edges) == n*(n-1)/2 { // 完全图
		cnt := [26]int{}
		for _, ch := range label {
			cnt[ch-'a']++
		}
		odd := 0
		for _, c := range cnt {
			ans += c - c%2
			odd |= c % 2
		}
		return ans + odd
	}

	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	memo := make([][][]int, n)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, 1<<n)
			for p := range memo[i][j] {
				memo[i][j][p] = -1
			}
		}
	}

	// 计算从 x 和 y 向两侧扩展，最多还能访问多少个节点（不算 x 和 y）
	var dfs func(int, int, int) int
	dfs = func(x, y, vis int) (res int) {
		p := &memo[x][y][vis]
		if *p >= 0 { // 之前计算过
			return *p
		}
		for _, v := range g[x] {
			if vis>>v&1 > 0 { // v 在路径中
				continue
			}
			for _, w := range g[y] {
				if vis>>w&1 == 0 && w != v && label[w] == label[v] {
					// 保证 v < w，减少状态个数和计算量
					r := dfs(min(v, w), max(v, w), vis|1<<v|1<<w)
					res = max(res, r+2)
				}
			}
		}
		*p = res // 记忆化
		return
	}

	for x, to := range g {
		// 奇回文串，x 作为回文中心
		ans = max(ans, dfs(x, x, 1<<x)+1)
		if ans == n {
			return
		}
		// 偶回文串，x 和 x 的邻居 y 作为回文中心
		for _, y := range to {
			// 保证 x < y，减少状态个数和计算量
			if x < y && label[x] == label[y] {
				ans = max(ans, dfs(x, y, 1<<x|1<<y)+2)
				if ans == n {
					return
				}
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^4 2^n)$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n^2 2^n)$，最坏情况下（完全图）单个状态的计算时间为 $\mathcal{O}(n^2)$，所以总的时间复杂度为 $\mathcal{O}(n^4 2^n)$。
- 空间复杂度：$\mathcal{O}(n^2 2^n)$。保存多少状态，就需要多少空间。

## 专题训练

见下面动态规划题单的「**§9.2 排列型 ② 相邻相关**」。

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
