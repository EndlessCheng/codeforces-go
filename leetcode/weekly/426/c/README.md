**核心思路**：对于第一棵树的节点 $i$，新添加的边的一个端点必然是 $i$。因为用其他节点当作端点，只会让第二棵树的节点到 $i$ 距离变得更远。

新添加的边，连到第二棵树的哪个节点上呢？

暴力枚举第二棵树的节点 $j$，用 DFS 计算距离 $j$ 不超过 $k-1$ 的节点个数 $\textit{cnt}_j$。这里 $k-1$ 是因为新添加的边也算在距离中。所有 $\textit{cnt}_j$ 取最大值，记作 $\textit{max}_2$。新添加的边就连到 $\textit{max}_2$ 对应的节点上。

同样地，暴力枚举第一棵树的节点 $i$，用 DFS 计算距离 $i$ 不超过 $k$ 的节点个数 $\textit{cnt}_i$。那么 $\textit{answer}[i] = \textit{cnt}_i + \textit{max}_2$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1tAzoY1EUN/?t=16m03s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def buildTree(self, edges: list[list[int]], k: int) -> Callable[[int, int, int], int]:
        g = [[] for _ in range(len(edges) + 1)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        def dfs(x: int, fa: int, d: int) -> int:
            if d > k:
                return 0
            cnt = 1
            for y in g[x]:
                if y != fa:
                    cnt += dfs(y, x, d + 1)
            return cnt
        return dfs

    def maxTargetNodes(self, edges1: List[List[int]], edges2: List[List[int]], k: int) -> List[int]:
        max2 = 0
        if k:
            dfs = self.buildTree(edges2, k - 1)  # 注意这里传的是 k-1
            max2 = max(dfs(i, -1, 0) for i in range(len(edges2) + 1))

        dfs = self.buildTree(edges1, k)
        return [dfs(i, -1, 0) + max2 for i in range(len(edges1) + 1)]
```

```java [sol-Java]
class Solution {
    public int[] maxTargetNodes(int[][] edges1, int[][] edges2, int k) {
        int max2 = 0;
        if (k > 0) {
            List<Integer>[] g = buildTree(edges2);
            for (int i = 0; i < edges2.length + 1; i++) {
                max2 = Math.max(max2, dfs(i, -1, 0, g, k - 1)); // 注意这里传的是 k-1
            }
        }

        List<Integer>[] g = buildTree(edges1);
        int[] ans = new int[edges1.length + 1];
        for (int i = 0; i < ans.length; i++) {
            ans[i] = dfs(i, -1, 0, g, k) + max2;
        }
        return ans;
    }

    private List<Integer>[] buildTree(int[][] edges) {
        List<Integer>[] g = new ArrayList[edges.length + 1];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }
        return g;
    }

    private int dfs(int x, int fa, int d, List<Integer>[] g, int k) {
        if (d > k) {
            return 0;
        }
        int cnt = 1;
        for (int y : g[x]) {
            if (y != fa) {
                cnt += dfs(y, x, d + 1, g, k);
            }
        }
        return cnt;
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<vector<int>> buildTree(vector<vector<int>>& edges) {
        vector<vector<int>> g(edges.size() + 1);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }
        return g;
    }

    int dfs(int x, int fa, int d, vector<vector<int>>& g, int k) {
        if (d > k) {
            return 0;
        }
        int cnt = 1;
        for (int y : g[x]) {
            if (y != fa) {
                cnt += dfs(y, x, d + 1, g, k);
            }
        }
        return cnt;
    }

public:
    vector<int> maxTargetNodes(vector<vector<int>>& edges1, vector<vector<int>>& edges2, int k) {
        int max2 = 0;
        if (k > 0) {
            auto g = buildTree(edges2);
            for (int i = 0; i < edges2.size() + 1; i++) {
                max2 = max(max2, dfs(i, -1, 0, g, k - 1)); // 注意这里传的是 k-1
            }
        }

        auto g = buildTree(edges1);
        vector<int> ans(edges1.size() + 1);
        for (int i = 0; i < ans.size(); i++) {
            ans[i] = dfs(i, -1, 0, g, k) + max2;
        }
        return ans;
    }
};
```

```go [sol-Go]
func buildTree(edges [][]int, k int) func(int, int, int) int {
	g := make([][]int, len(edges)+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int, int) int
	dfs = func(x, fa, d int) int {
		if d > k {
			return 0
		}
		cnt := 1
		for _, y := range g[x] {
			if y != fa {
				cnt += dfs(y, x, d+1)
			}
		}
		return cnt
	}
	return dfs
}

func maxTargetNodes(edges1, edges2 [][]int, k int) []int {
	max2 := 0
	if k > 0 {
		dfs := buildTree(edges2, k-1) // 注意这里传的是 k-1
		for i := range len(edges2) + 1 {
			max2 = max(max2, dfs(i, -1, 0))
		}
	}

	dfs := buildTree(edges1, k)
	ans := make([]int, len(edges1)+1)
	for i := range ans {
		ans[i] = dfs(i, -1, 0) + max2
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2+m^2)$，其中 $n$ 是 $\textit{edges}_1$ 的长度，$m$ 是 $\textit{edges}_2$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. 【本题相关】[链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
