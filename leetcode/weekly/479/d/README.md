**前置知识**：[【图解】一张图秒懂换根 DP！](https://leetcode.cn/problems/sum-of-distances-in-tree/solution/tu-jie-yi-zhang-tu-miao-dong-huan-gen-dp-6bgb/)

## 第一次 DFS

计算以 $0$ 为根时，每棵子树 $x$ 的最大得分（一定包含节点 $x$），记作 $\textit{subScore}[x]$。注意是子树 $x$，不是子图 $x$，前者不包含 $x$ 的父节点。

1. 把节点 $x$ 算到 $\textit{subScore}[x]$ 中：如果 $\textit{good}[x] = 1$ 则初始化 $\textit{subScore}[x]=1$，否则初始化 $\textit{subScore}[x]=-1$。
2. 对于 $x$ 的儿子 $y$，递归计算子树 $y$ 的最大得分（一定包含节点 $y$）。
3. 如果子树 $y$ 的最大得分大于 $0$，那么选子树 $y$，否则不选。也就是把 $\max(\textit{subScore}[y],0)$ 加到 $\textit{subScore}[x]$ 中。

## 第二次 DFS

假设我们算出了**子图** $x$ 的最大得分 $\textit{ans}[x]$，现在要对 $x$ 的儿子 $y$，计算子图 $y$ 的最大得分。

从「以 $x$ 为根」换到「以 $y$ 为根」：

1. 从 $x$ 中去掉子树 $y$，剩余部分记作 $F$。
2. 换根后，$F$ 变成挂在 $y$ 下面的一棵子树。

子图 $y$ 的得分由两部分组成：

1. 原来的 $\textit{subScore}[y]$。
2. $\textit{ans}[x] - \max(\textit{subScore}[y], 0)$。这是子树 $F$ 的最大得分。如果大于 $0$ 则选子树 $F$，否则不选。 

合并两部分，得到从 $x$ 换根到 $y$ 的转移方程

$$
\textit{ans}[y] = \textit{subScore}[y] + \max(\textit{ans}[x] - \max(\textit{subScore}[y], 0), 0)
$$

[本题视频讲解](https://www.bilibili.com/video/BV1sv2fB4Evi/)，欢迎点赞关注~

## 写法一

```py [sol-Python3]
# 手写 max 更快
max = lambda a, b: b if b > a else a

class Solution:
    def maxSubgraphScore(self, n: int, edges: List[List[int]], good: List[int]) -> List[int]:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        # sub_score[x] 表示（以 0 为根时）包含 x 的子树 x 的最大得分（注意是子树不是子图）
        sub_score = [0] * n
        # 计算 sub_score[x]
        def dfs(x: int, fa: int) -> None:
            sub_score[x] = 1 if good[x] else -1  # sub_score[x] 一定包含 x
            for y in g[x]:
                if y != fa:
                    dfs(y, x)
                    # 如果子树 y 的最大得分 > 0，选子树 y，否则不选
                    sub_score[x] += max(sub_score[y], 0)
        dfs(0, -1)

        ans = [0] * n
        ans[0] = sub_score[0]
        # 对于 x 的儿子 y，计算包含 y 的子图最大得分
        def reroot(x: int, fa: int) -> None:
            for y in g[x]:
                if y != fa:
                    # 从 ans[x] 中去掉子树 y。换根后，这部分内容变成 y 的一棵子树（记作 F）
                    score_f = ans[x] - max(sub_score[y], 0)
                    # 如果子树 F 的最大得分 > 0，选子树 F，否则不选
                    ans[y] = sub_score[y] + max(score_f, 0)
                    reroot(y, x)
        reroot(0, -1)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] maxSubgraphScore(int n, int[][] edges, int[] good) {
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        // subScore[x] 表示（以 0 为根时）包含 x 的子树 x 的最大得分（注意是子树不是子图）
        int[] subScore = new int[n];
        dfs(0, -1, g, good, subScore);

        int[] ans = new int[n];
        ans[0] = subScore[0];
        reroot(0, -1, g, subScore, ans);
        return ans;
    }

    // 计算 subScore[x]
    private void dfs(int x, int fa, List<Integer>[] g, int[] good, int[] subScore) {
        subScore[x] = good[x] == 0 ? -1 : 1; // subScore[x] 一定包含 x
        for (int y : g[x]) {
            if (y != fa) {
                dfs(y, x, g, good, subScore);
                // 如果子树 y 的最大得分 > 0，选子树 y，否则不选
                subScore[x] += Math.max(subScore[y], 0);
            }
        }
    }

    // 对于 x 的儿子 y，计算包含 y 的子图最大得分
    private void reroot(int x, int fa, List<Integer>[] g, int[] subScore, int[] ans) {
        for (int y : g[x]) {
            if (y != fa) {
                // 从 ans[x] 中去掉子树 y。换根后，这部分内容变成 y 的一棵子树（记作 F）
                int scoreF = ans[x] - Math.max(subScore[y], 0);
                // 如果子树 F 的最大得分 > 0，选子树 F，否则不选
                ans[y] = subScore[y] + Math.max(scoreF, 0);
                reroot(y, x, g, subScore, ans);
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> maxSubgraphScore(int n, vector<vector<int>>& edges, vector<int>& good) {
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        // sub_score[x] 表示（以 0 为根时）包含 x 的子树 x 的最大得分（注意是子树不是子图）
        vector<int> sub_score(n);
        // 计算 sub_score[x]
        auto dfs = [&](this auto&& dfs, int x, int fa) -> void {
            sub_score[x] = good[x] ? 1 : -1; // subScore[x] 一定包含 x
            for (int y : g[x]) {
                if (y != fa) {
                    dfs(y, x);
                    // 如果子树 y 的最大得分 > 0，选子树 y，否则不选
                    sub_score[x] += max(sub_score[y], 0);
                }
            }
        };
        dfs(0, -1);

        vector<int> ans(n);
        ans[0] = sub_score[0];
        // 对于 x 的儿子 y，计算包含 y 的子图最大得分
        auto reroot = [&](this auto&& reroot, int x, int fa) -> void {
            for (int y : g[x]) {
                if (y != fa) {
                    // 从 ans[x] 中去掉子树 y。换根后，这部分内容变成 y 的一棵子树（记作 F）
                    int score_f = ans[x] - max(sub_score[y], 0);
                    // 如果子树 F 的最大得分 > 0，选子树 F，否则不选
                    ans[y] = sub_score[y] + max(score_f, 0);
                    reroot(y, x);
                }
            }
        };
        reroot(0, -1);
        return ans;
    }
};
```

```go [sol-Go]
func maxSubgraphScore(n int, edges [][]int, good []int) []int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// subScore[x] 表示（以 0 为根时）包含 x 的子树 x 的最大得分（注意是子树不是子图）
	subScore := make([]int, n)
	// 计算 subScore[x]
	var dfs func(int, int)
	dfs = func(x, fa int) {
		subScore[x] = good[x]*2 - 1 // subScore[x] 一定包含 x
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x)
				// 如果子树 y 的最大得分 > 0，选子树 y，否则不选
				subScore[x] += max(subScore[y], 0)
			}
		}
	}
	dfs(0, -1)

	ans := make([]int, n)
	ans[0] = subScore[0]
	// 对于 x 的儿子 y，计算包含 y 的子图最大得分
	var reroot func(int, int)
	reroot = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				// 从 ans[x] 中去掉子树 y。换根后，这部分内容变成 y 的一棵子树（记作 F）
				scoreF := ans[x] - max(subScore[y], 0)
				// 如果子树 F 的最大得分 > 0，选子树 F，否则不选
				ans[y] = subScore[y] + max(scoreF, 0)
				reroot(y, x)
			}
		}
	}
	reroot(0, -1)
	return ans
}
```

## 写法二

观察上面的代码，$\textit{subScore}[y]$ 使用完后就不再用到了，所以可以把 $\textit{ans}[y]$ 记在 $\textit{subScore}[y]$ 中，无需创建 $\textit{ans}$ 数组。

进一步地，$\textit{subScore}$ 数组也无需创建，直接把数据记在 $\textit{good}$ 数组中。

```py [sol-Python3]
# 手写 max 更快
max = lambda a, b: b if b > a else a

class Solution:
    def maxSubgraphScore(self, n: int, edges: List[List[int]], ans: List[int]) -> List[int]:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        def dfs(x: int, fa: int) -> None:
            ans[x] = 1 if ans[x] else -1
            for y in g[x]:
                if y != fa:
                    dfs(y, x)
                    # 如果子树 y 的最大得分 > 0，选子树 y，否则不选
                    ans[x] += max(ans[y], 0)
        dfs(0, -1)

        # 对于 x 的儿子 y，计算包含 y 的子图最大得分
        def reroot(x: int, fa: int) -> None:
            for y in g[x]:
                if y != fa:
                    # 从 ans[x] 中去掉子树 y。换根后，这部分内容变成 y 的一棵子树（记作 F）
                    score_f = ans[x] - max(ans[y], 0)
                    # 如果子树 F 的最大得分 > 0，选子树 F，否则不选
                    ans[y] += max(score_f, 0)
                    reroot(y, x)
        reroot(0, -1)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] maxSubgraphScore(int n, int[][] edges, int[] good) {
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        dfs(0, -1, g, good);
        reroot(0, -1, g, good);
        return good;
    }

    private void dfs(int x, int fa, List<Integer>[] g, int[] ans) {
        ans[x] = ans[x] == 0 ? -1 : 1;
        for (int y : g[x]) {
            if (y != fa) {
                dfs(y, x, g, ans);
                // 如果子树 y 的最大得分 > 0，选子树 y，否则不选
                ans[x] += Math.max(ans[y], 0);
            }
        }
    }

    // 对于 x 的儿子 y，计算包含 y 的子图最大得分
    private void reroot(int x, int fa, List<Integer>[] g, int[] ans) {
        for (int y : g[x]) {
            if (y != fa) {
                // 从 ans[x] 中去掉子树 y。换根后，这部分内容变成 y 的一棵子树（记作 F）
                int scoreF = ans[x] - Math.max(ans[y], 0);
                // 如果子树 F 的最大得分 > 0，选子树 F，否则不选
                ans[y] += Math.max(scoreF, 0);
                reroot(y, x, g, ans);
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> maxSubgraphScore(int n, vector<vector<int>>& edges, vector<int>& ans) {
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        auto dfs = [&](this auto&& dfs, int x, int fa) -> void {
            ans[x] = ans[x] ? 1 : -1;
            for (int y : g[x]) {
                if (y != fa) {
                    dfs(y, x);
                    // 如果子树 y 的最大得分 > 0，选子树 y，否则不选
                    ans[x] += max(ans[y], 0);
                }
            }
        };
        dfs(0, -1);

        // 对于 x 的儿子 y，计算包含 y 的子图最大得分
        auto reroot = [&](this auto&& reroot, int x, int fa) -> void {
            for (int y : g[x]) {
                if (y != fa) {
                    // 从 ans[x] 中去掉子树 y。换根后，这部分内容变成 y 的一棵子树（记作 F）
                    int score_f = ans[x] - max(ans[y], 0);
                    // 如果子树 F 的最大得分 > 0，选子树 F，否则不选
                    ans[y] += max(score_f, 0);
                    reroot(y, x);
                }
            }
        };
        reroot(0, -1);
        return ans;
    }
};
```

```go [sol-Go]
func maxSubgraphScore(n int, edges [][]int, ans []int) []int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int)
	dfs = func(x, fa int) {
		ans[x] = ans[x]*2 - 1
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x)
				// 如果子树 y 的最大得分 > 0，选子树 y，否则不选
				ans[x] += max(ans[y], 0)
			}
		}
	}
	dfs(0, -1)

	// 对于 x 的儿子 y，计算包含 y 的子图最大得分
	var reroot func(int, int)
	reroot = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				// 从 ans[x] 中去掉子树 y。换根后，这部分内容变成 y 的一棵子树（记作 F）
				scoreF := ans[x] - max(ans[y], 0)
				// 如果子树 F 的最大得分 > 0，选子树 F，否则不选
				ans[y] += max(scoreF, 0)
				reroot(y, x)
			}
		}
	}
	reroot(0, -1)
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面动态规划题单的「**§12.4 换根 DP**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
