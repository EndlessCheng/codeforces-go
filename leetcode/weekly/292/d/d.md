#### 提示 1

用一个变量 $c$ 表示括号字符串的平衡度：遇到左括号就 $+1$，遇到右括号就 $-1$。那么合法字符串等价于任意时刻 $c\ge 0$ 且最后 $c=0$。

#### 提示 2

从起点到终点，往下走的次数是固定的，即 $m-1$ 次，往右走的次数也是固定的，即 $n-1$ 次，因此路径长度（字符串长度）是一个定值，即 $(m-1)+(n-1)+1 = m+n-1$。

极限情况下合法的字符串左半均为左括号，右半均为右括号，因此 $c$ 最大为 $\dfrac{m+n-1}{2}$。

#### 提示 3

把进入格子**前**的 $c$ 值当作格子的附加状态，那么一个格子至多有 $\dfrac{m+n-1}{2}+1=\dfrac{m+n+1}{2}$ 个不同的状态，整个网格图至多有 $\dfrac{mn(m+n+1)}{2}$ 个不同的状态。

#### 提示 4

在这些状态上 DFS：

- 起点为 $(0,0,0)$，表示从左上角 $(0,0)$ 出发，初始 $c=0$；
- 终点为 $(m-1,n-1,1)$，表示到右下角 $(m-1,n-1)$ 结束，且进入前 $c=1$（因为右下角必须为右括号）；
- 根据当前格子的字符计算 $c$ 值，然后往下或往右移动，继续 DFS。

代码实现时，由于找到合法路径就返回 `true` 了，不会继续执行 `dfs`，若 `dfs(x,y,c)` 最后返回的是 `false`，那后续访问同一个状态时（再次调用 `dfs(x,y,c)`），仍然会得到 `false`。因此没必要重复访问同一个状态，可以用一个 $\textit{vis}$ 数组标记，遇到访问过的状态可以直接返回 `false`。

另外有一个比较强的剪枝优化：由于字符串左括号和右括号的数目必须相同，因此字符串的长度为偶数，所以 $m+n-1$ 必须是偶数。我们可以在 DFS 之前就预先判断这一要求是否成立。

#### 复杂度分析

- 时间复杂度：$O(mn(m+n))$，每个状态至多访问一次。
- 空间复杂度：$O(mn(m+n))$，需要记录每个状态是否被访问过。

```Python [sol1-Python3]
class Solution:
    def hasValidPath(self, grid: List[List[str]]) -> bool:
        m, n = len(grid), len(grid[0])
        if (m + n) % 2 == 0 or grid[0][0] == ')' or grid[m - 1][n - 1] == '(': return False  # 剪枝

        @cache  # 效果类似 vis 数组
        def dfs(x: int, y: int, c: int) -> bool:
            if c > m - x + n - y - 1: return False  # 剪枝：即使后面都是 ')' 也不能将 c 减为 0
            if x == m - 1 and y == n - 1: return c == 1  # 终点一定是 ')'
            c += 1 if grid[x][y] == '(' else -1
            return c >= 0 and (x < m - 1 and dfs(x + 1, y, c) or y < n - 1 and dfs(x, y + 1, c))  # 往下或者往右
        return dfs(0, 0, 0)  # 起点
```

```java [sol1-Java]
class Solution {
    int m, n;
    char[][] grid;
    boolean[][][] vis;

    public boolean hasValidPath(char[][] grid) {
        m = grid.length;
        n = grid[0].length;
        if ((m + n) % 2 == 0 || grid[0][0] == ')' || grid[m - 1][n - 1] == '(') return false; // 剪枝
        this.grid = grid;
        vis = new boolean[m][n][(m + n + 1) / 2];
        return dfs(0, 0, 0);
    }

    boolean dfs(int x, int y, int c) {
        if (c > m - x + n - y - 1) return false; // 剪枝：即使后面都是 ')' 也不能将 c 减为 0
        if (x == m - 1 && y == n - 1) return c == 1; // 终点一定是 ')'
        if (vis[x][y][c]) return false; // 重复访问
        vis[x][y][c] = true;
        c += grid[x][y] == '(' ? 1 : -1;
        return c >= 0 && (x < m - 1 && dfs(x + 1, y, c) || y < n - 1 && dfs(x, y + 1, c)); // 往下或者往右
    }
}
```

```C++ [sol1-C++]
class Solution {
public:
    bool hasValidPath(vector<vector<char>> &grid) {
        int m = grid.size(), n = grid[0].size();
        if ((m + n) % 2 == 0 || grid[0][0] == ')' || grid[m - 1][n - 1] == '(') return false; // 剪枝
        bool vis[m][n][(m + n + 1) / 2]; memset(vis, 0, sizeof(vis));
        function<bool(int, int, int)> dfs = [&](int x, int y, int c) -> bool {
            if (c > m - x + n - y - 1) return false; // 剪枝：即使后面都是 ')' 也不能将 c 减为 0
            if (x == m - 1 && y == n - 1) return c == 1; // 终点一定是 ')'
            if (vis[x][y][c]) return false; // 重复访问
            vis[x][y][c] = true;
            c += grid[x][y] == '(' ? 1 : -1;
            return c >= 0 && (x < m - 1 && dfs(x + 1, y, c) || y < n - 1 && dfs(x, y + 1, c)); // 往下或者往右
        };
        return dfs(0, 0, 0);
    }
};
```

```go [sol1-Go]
func hasValidPath(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	if (m+n)%2 == 0 || grid[0][0] == ')' || grid[m-1][n-1] == '(' { // 剪枝
		return false
	}

	vis := make([][][]bool, m)
	for i := range vis {
		vis[i] = make([][]bool, n)
		for j := range vis[i] {
			vis[i][j] = make([]bool, (m+n+1)/2)
		}
	}
	var dfs func(x, y, c int) bool
	dfs = func(x, y, c int) bool {
		if c > m-x+n-y-1 { // 剪枝：即使后面都是 ')' 也不能将 c 减为 0
			return false
		}
		if x == m-1 && y == n-1 { // 终点
			return c == 1 // 终点一定是 ')'
		}
		if vis[x][y][c] { // 重复访问
			return false
		}
		vis[x][y][c] = true
		if grid[x][y] == '(' {
			c++
		} else if c--; c < 0 { // 非法括号字符串
			return false
		}
		return x < m-1 && dfs(x+1, y, c) || y < n-1 && dfs(x, y+1, c) // 往下或者往右
	}
	return dfs(0, 0, 0) // 起点
}
```

#### 评注

值得注意的是，DFS 的写法相比某些递推的写法要快 $10$ 倍以上，这是因为有很多状态是无法访问到的：比如 $(x=2,y=3,c=100)$ 这个状态就是不可达的，此时还没走几步，$c$ 不可能这么大。或者对于一些随机的网格图，$c$ 的值也会比较小。这种情况下 DFS 的优势就发挥出来了，DFS 可以十分自然地遍历到所有合法的状态，加上自带的剪枝效果，可以大大降低访问到的状态数。
