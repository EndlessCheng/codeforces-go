如果没有相邻节点的限制，那么本题求的就是树的直径上的点的个数，见 [1245. 树的直径](https://leetcode-cn.com/problems/tree-diameter/)。

考虑用树形 DP 求直径。枚举子树 $x$ 的所有子树 $y$，维护从 $x$ 出发的最长路径 $\textit{maxLen}$，那么可以更新答案为从 $y$ 出发的最长路径加上 $\textit{maxLen}$，再加上 $1$（边 $x-y$），即**合并从 $x$ 出发的两条路径**。递归结束时返回 $\textit{maxLen}$。

对于本题的限制，我们可以在从子树 $y$ 转移过来时，仅考虑从满足 $s[y]\ne s[x]$ 的子树 $y$ 转移过来，所以对上述做法加个 `if` 判断就行了。

由于本题求的是点的个数，所以答案为最长路径的长度加一。

```Python [sol1-Python3]
class Solution:
    def longestPath(self, parent: List[int], s: str) -> int:
        n = len(parent)
        g = [[] for _ in range(n)]
        for i in range(1, n):
            g[parent[i]].append(i)

        ans = 0
        def dfs(x: int) -> int:
            nonlocal ans
            max_len = 0
            for y in g[x]:
                len = dfs(y) + 1
                if s[y] != s[x]:
                    ans = max(ans, max_len + len)
                    max_len = max(max_len, len)
            return max_len
        dfs(0)
        return ans + 1
```

```java [sol1-Java]
class Solution {
    private List<Integer>[] g;
    private char[] s;
    private int ans;

    public int longestPath(int[] parent, String s) {
        this.s = s.toCharArray();
        int n = parent.length;
        g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (int i = 1; i < n; i++) {
            g[parent[i]].add(i);
        }
        dfs(0);
        return ans + 1;
    }

    private int dfs(int x) {
        int maxLen = 0;
        for (int y : g[x]) {
            int len = dfs(y) + 1;
            if (s[y] != s[x]) {
                ans = Math.max(ans, maxLen + len);
                maxLen = Math.max(maxLen, len);
            }
        }
        return maxLen;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int longestPath(vector<int> &parent, string &s) {
        int n = parent.size();
        vector<vector<int>> g(n);
        for (int i = 1; i < n; i++) {
            g[parent[i]].push_back(i);
        }
        int ans = 0;
        auto dfs = [&](auto&& dfs, int x) -> int {
            int maxLen = 0;
            for (int y : g[x]) {
                int len = dfs(dfs, y) + 1;
                if (s[y] != s[x]) {
                    ans = max(ans, maxLen + len);
                    maxLen = max(maxLen, len);
                }
            }
            return maxLen;
        };
        dfs(dfs, 0);
        return ans + 1;
    }
};
```

```go [sol1-Go]
func longestPath(parent []int, s string) (ans int) {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		pa := parent[i]
		g[pa] = append(g[pa], i)
	}

	var dfs func(int) int
	dfs = func(x int) (maxLen int) {
		for _, y := range g[x] {
			len := dfs(y) + 1
			if s[y] != s[x] {
				ans = max(ans, maxLen+len)
				maxLen = max(maxLen, len)
			}
		}
		return
	}
	dfs(0)
	return ans + 1
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$。
- 空间复杂度：$O(n)$。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
