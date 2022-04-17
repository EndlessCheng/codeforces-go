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
				ans = max(ans, len+maxLen)
				maxLen = max(maxLen, len)
			}
		}
		return
	}
	dfs(0)
	return ans + 1
}

func max(a, b int) int { if b > a { return b }; return a }
```

```C++ [sol1-C++]
class Solution {
public:
    int longestPath(vector<int> &parent, string &s) {
        int n = parent.size();
        vector<vector<int>> g(n);
        for (int i = 1; i < n; ++i)
            g[parent[i]].push_back(i);

        int ans = 0;
        function<int(int)> dfs = [&](int x) -> int {
            int maxLen = 0;
            for (int y : g[x]) {
                int len = dfs(y) + 1;
                if (s[y] != s[x]) {
                    ans = max(ans, maxLen + len);
                    maxLen = max(maxLen, len);
                }
            }
            return maxLen;
        };
        dfs(0);
        return ans + 1;
    }
};
```

```java [sol1-Java]
class Solution {
    List<Integer>[] g;
    String s;
    int ans;

    public int longestPath(int[] parent, String s) {
        this.s = s;
        var n = parent.length;
        g = new ArrayList[n];
        for (var i = 0; i < n; i++) g[i] = new ArrayList<>();
        for (var i = 1; i < n; i++) g[parent[i]].add(i);

        dfs(0);
        return ans + 1;
    }

    int dfs(int x) {
        var maxLen = 0;
        for (var y : g[x]) {
            int len = dfs(y) + 1;
            if (s.charAt(y) != s.charAt(x)) {
                ans = Math.max(ans, maxLen + len);
                maxLen = Math.max(maxLen, len);
            }
        }
        return maxLen;
    }
}
```
