**前置知识**：[【图解】一张图秒懂换根 DP！](https://leetcode.cn/problems/sum-of-distances-in-tree/solution/tu-jie-yi-zhang-tu-miao-dong-huan-gen-dp-6bgb/)

## 第一次 DFS

计算以 $0$ 为根时，每棵子树 $x$ 的最大得分（一定包含节点 $x$），记作 $\textit{subScore}[x]$。注意是子树 $x$，不是子图 $x$，前者不包含 $x$ 的父节点。

1. 对于 $x$ 的儿子 $y$，递归计算子树 $y$ 的最大得分（一定包含节点 $y$）。
2. 如果子树 $y$ 的最大得分是负数，那么不选子树 $y$ 的得分，否则选子树 $y$ 的最大得分，即累加 $\max(\textit{subScore}[y],0)$。
3. 最后把节点 $x$ 的贡献算到 $\textit{subScore}[x]$ 中：如果 $\textit{good}[x] = 1$ 则加一，否则减一。

## 第二次 DFS

假设我们算出了**子图** $x$ 的得分 $\textit{scoreX}$，现在要对 $x$ 的儿子 $y$，计算子图 $y$ 的得分。

从「以 $x$ 为根」换到「以 $y$ 为根」，那么 $x$ 去掉子树 $y$ 后的剩余部分，就变成挂在 $y$ 下面的一棵子树了。

子图 $y$ 的得分由两部分组成：

1. $\textit{subScore}[y]$：这里面的节点只在子树 $y$ 中。
2. 来自 $y$ 的父节点 $x$ 的最大得分：从 $\textit{scoreX}$ 中减去子树 $y$ 的贡献 $\max(\textit{subScore}[y],0)$，即为来自 $x$ 的最大得分。

[本题视频讲解](https://www.bilibili.com/video/BV1sv2fB4Evi/)，欢迎点赞关注~

```py [sol-Python3]
# 手写 max 更快
max = lambda a, b: b if b > a else a

class Solution:
    def maxSubgraphScore(self, n: int, edges: List[List[int]], good: List[int]) -> List[int]:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        # sub_score[x] 表示（以 0 为根时）子树 x 的最大得分（一定包含节点 x）
        sub_score = [0] * n

        # 计算并返回 sub_score[x]
        def dfs(x: int, fa: int) -> int:
            for y in g[x]:
                if y != fa:
                    # 如果子树 y 的得分是负数，不选子树 y，否则选子树 y
                    sub_score[x] += max(dfs(y, x), 0)
            sub_score[x] += 1 if good[x] else -1  # sub_score[x] 一定包含 x
            return sub_score[x]

        dfs(0, -1)

        ans = [0] * n

        # 计算子图 x 的最大得分 score_x，其中 faScore 表示来自父节点 fa 的最大得分（一定包含节点 fa）
        def reroot(x: int, fa: int, fa_score: int) -> None:
            ans[x] = score_x = sub_score[x] + max(fa_score, 0)
            for y in g[x]:
                if y != fa:
                    # score_x - max(sub_score[y], 0) 是不含子树 y 的最大得分
                    reroot(y, x, score_x - max(sub_score[y], 0))

        reroot(0, -1, 0)
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

        // subScore[x] 表示（以 0 为根时）子树 x 的最大得分（一定包含节点 x）
        int[] subScore = new int[n];
        dfs(0, -1, g, good, subScore);

        int[] ans = new int[n];
        reroot(0, -1, 0, g, subScore, ans);
        return ans;
    }

    // 计算并返回 subScore[x]
    private int dfs(int x, int fa, List<Integer>[] g, int[] good, int[] subScore) {
        for (int y : g[x]) {
            if (y != fa) {
                // 如果子树 y 的得分是负数，不选子树 y，否则选子树 y
                subScore[x] += Math.max(dfs(y, x, g, good, subScore), 0);
            }
        }
        subScore[x] += good[x] == 1 ? 1 : -1; // subScore[x] 一定包含 x
        return subScore[x];
    }

    // 计算子图 x 的最大得分 scoreX，其中 faScore 表示来自父节点 fa 的最大得分（一定包含节点 fa）
    private void reroot(int x, int fa, int faScore, List<Integer>[] g, int[] subScore, int[] ans) {
        int scoreX = subScore[x] + Math.max(faScore, 0);
        ans[x] = scoreX;
        for (int y : g[x]) {
            if (y != fa) {
                // scoreX-max(subScore[y],0) 是不含子树 y 的最大得分
                reroot(y, x, scoreX - Math.max(subScore[y], 0), g, subScore, ans);
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

        // sub_score[x] 表示（以 0 为根时）子树 x 的最大得分（一定包含节点 x）
        vector<int> sub_score(n);
        // 计算并返回 sub_score[x]
        auto dfs = [&](this auto&& dfs, int x, int fa) -> int {
            for (int y : g[x]) {
                if (y != fa) {
                    // 如果子树 y 的得分是负数，不选子树 y，否则选子树 y
                    sub_score[x] += max(dfs(y, x), 0);
                }
            }
            sub_score[x] += good[x] ? 1 : -1; // sub_score[x] 一定包含 x
            return sub_score[x];
        };
        dfs(0, -1);

        vector<int> ans(n);
        // 计算子图 x 的最大得分 score_x，其中 faScore 表示来自父节点 fa 的最大得分（一定包含节点 fa）
        auto reroot = [&](this auto&& reroot, int x, int fa, int fa_score) -> void {
            int score_x = sub_score[x] + max(fa_score, 0);
            ans[x] = score_x;
            for (int y : g[x]) {
                if (y != fa) {
                    // score_x - max(sub_score[y], 0) 是不含子树 y 的最大得分
                    reroot(y, x, score_x - max(sub_score[y], 0));
                }
            }
        };
        reroot(0, -1, 0);
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

	// subScore[x] 表示（以 0 为根时）子树 x 的最大得分（一定包含节点 x）
	subScore := make([]int, n)
	// 计算并返回 subScore[x]
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		for _, y := range g[x] {
			if y != fa {
				// 如果子树 y 的得分是负数，不选子树 y，否则选子树 y
				subScore[x] += max(dfs(y, x), 0)
			}
		}
		subScore[x] += good[x]*2 - 1 // subScore[x] 一定包含 x
		return subScore[x]
	}
	dfs(0, -1)

	ans := make([]int, n)
	// 计算子图 x 的最大得分 scoreX，其中 faScore 表示来自父节点 fa 的最大得分（一定包含节点 fa）
	var reroot func(int, int, int)
	reroot = func(x, fa, faScore int) {
		scoreX := subScore[x] + max(faScore, 0)
		ans[x] = scoreX
		for _, y := range g[x] {
			if y != fa {
				// scoreX-max(subScore[y],0) 是不含子树 y 的最大得分
				reroot(y, x, scoreX-max(subScore[y], 0))
			}
		}
	}
	reroot(0, -1, 0)
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
