下午两点【bilibili@灵茶山艾府】直播讲题，记得关注哦~

---

### 提示 1

如果让你把所有**从起点到终点**的路径上的格子都做个标记，这些标记的「轮廓」会是什么样的？

如果可以使矩阵不连通，「轮廓」应该有什么特点？

### 提示 2

如果可以使矩阵不连通，那么你翻转的那个格子必然会使「轮廓」也不连通（断开）。

如何找到「轮廓」？

### 提示 3

从 $(1,0)$ 出发，优先向下走，其次向右走，得到下轮廓。

从 $(0,1)$ 出发，优先向右走，其次向下走，得到上轮廓。

如果两个轮廓有交集（除了终点），那么翻转交集中的任意一个格子，都可以使矩阵不连通。

代码实现时，可以直接把下轮廓的格子值修改成 $0$（除了终点），如果再从 $(0,1)$ 出发，无法到达终点，则说明可以使矩阵不连通。也就是说，得到下轮廓后，无需求上轮廓，只需要看 $(0,1)$ 和终点之间是否有通路即可。

```py [sol1-Python3]
class Solution:
    def isPossibleToCutPath(self, g: List[List[int]]) -> bool:
        m, n = len(g), len(g[0])

        def dfs(x: int, y: int) -> bool:  # 返回能否到达终点
            if x == m - 1 and y == n - 1: return True
            g[x][y] = 0  # 直接修改
            return x < m - 1 and g[x + 1][y] and dfs(x + 1, y) or \
                   y < n - 1 and g[x][y + 1] and dfs(x, y + 1)

        # 提前特判一些可以直接得到答案的情况
        return m * n > 2 and (m == 1 or n == 1 or g[1][0] == 0 or g[0][1] == 0 or g[-2][-1] == 0 or g[-1][-2] == 0 or
                              not dfs(1, 0) or not dfs(0, 1))
```

```java [sol1-Java]
class Solution {
    private int[][] g;
    private int m, n;

    public boolean isPossibleToCutPath(int[][] grid) {
        g = grid; m = g.length; n = g[0].length;
        // 提前特判一些可以直接得到答案的情况
        return m * n > 2 && (m == 1 || n == 1 || g[1][0] == 0 || g[0][1] == 0 || g[m - 2][n - 1] == 0 || g[m - 1][n - 2] == 0 ||
                             !dfs(1, 0) || !dfs(0, 1));
    }

    private boolean dfs(int x, int y) { // 返回能否到达终点
        if (x == m - 1 && y == n - 1) return true;
        g[x][y] = 0; // 直接修改
        return x < m - 1 && g[x + 1][y] > 0 && dfs(x + 1, y) ||
               y < n - 1 && g[x][y + 1] > 0 && dfs(x, y + 1);
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    bool isPossibleToCutPath(vector<vector<int>> &g) {
        int m = g.size(), n = g[0].size();

        function<bool(int, int)> dfs = [&](int x, int y) -> bool { // 返回能否到达终点
            if (x == m - 1 && y == n - 1) return true;
            g[x][y] = 0; // 直接修改
            return x < m - 1 && g[x + 1][y] && dfs(x + 1, y) ||
                   y < n - 1 && g[x][y + 1] && dfs(x, y + 1);
        };

        // 提前特判一些可以直接得到答案的情况
        return m * n > 2 && (m == 1 || n == 1 || g[1][0] == 0 || g[0][1] == 0 || g[m - 2][n - 1] == 0 || g[m - 1][n - 2] == 0 ||
                             !dfs(1, 0) || !dfs(0, 1));
    }
};
```

```go [sol1-Go]
func isPossibleToCutPath(g [][]int) bool {
    m, n := len(g), len(g[0])

    var dfs func(int, int) bool
    dfs = func(x, y int) bool { // 返回能否到达终点
        if x == m-1 && y == n-1 {
            return true
        }
        g[x][y] = 0 // 直接修改，同时保证每个点至多访问一次
        return x < m-1 && g[x+1][y] > 0 && dfs(x+1, y) ||
               y < n-1 && g[x][y+1] > 0 && dfs(x, y+1)
    }

    // 提前特判一些可以直接得到答案的情况
    return m*n > 2 && (m == 1 || n == 1 || g[1][0] == 0 || g[0][1] == 0 || g[m-2][n-1] == 0 || g[m-1][n-2] == 0 ||
                       !dfs(1, 0) || !dfs(0, 1))
}
```

### 复杂度分析

- 时间复杂度：$O(mn)$，其中 $m$ 为 $\textit{grid}$ 的长度，$n$ 为 $\textit{grid}[i]$ 的长度。
- 空间复杂度：$O(m+n)$。递归需要 $O(m+n)$ 的栈空间。

### 思考题

如果题目还允许向上和向左，要如何做呢？

欢迎在评论区发表你的做法。

---

如果你觉得自己的思维能力有些薄弱，可以做做 [从周赛中学算法 - 2022 年周赛题目总结（下篇）](https://leetcode.cn/circle/discuss/WR1MJP/) 中的「思维题」这节，所有题目我都写了题解。
