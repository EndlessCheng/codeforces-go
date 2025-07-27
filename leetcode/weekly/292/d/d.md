## 提示 1

用一个变量 $c$ 表示括号字符串的平衡度：遇到左括号就 $+1$，遇到右括号就 $-1$。那么合法字符串等价于任意时刻 $c\ge 0$ 且最后 $c=0$。

## 提示 2

从起点到终点，往下走的次数是固定的，即 $m-1$ 次，往右走的次数也是固定的，即 $n-1$ 次，因此路径长度（字符串长度）是一个定值，即 $(m-1)+(n-1)+1 = m+n-1$。

极限情况下合法的字符串左半均为左括号，右半均为右括号，因此 $c$ 最大为 $\dfrac{m+n-1}{2}$。

## 提示 3

把进入格子**前**的 $c$ 值当作格子的附加状态，定义状态 $(x,y,c)$ 表示进入格子 $(x,y)$ 且进入**前**平衡度为 $c$。

由于一个格子至多有 $\dfrac{m+n-1}{2}+1=\dfrac{m+n+1}{2}$ 个不同的状态，整个网格图至多有 $\dfrac{mn(m+n+1)}{2}$ 个不同的状态。

## 提示 4

在这些状态上 DFS：

- 起点为 $(0,0,0)$，表示从左上角 $(0,0)$ 出发，初始 $c=0$；
- 终点为 $(m-1,n-1,1)$，表示到右下角 $(m-1,n-1)$ 结束，且进入前 $c=1$（因为右下角必须为右括号）；
- 根据当前格子的字符计算 $c$ 值，然后往下或往右移动，继续 DFS。

代码实现时，由于找到合法路径就返回 `true` 了，不会继续执行 `dfs`，若 `dfs(x,y,c)` 最后返回的是 `false`，那后续访问同一个状态时（再次调用 `dfs(x,y,c)`），仍然会得到 `false`。因此没必要重复访问同一个状态，可以用一个 $\textit{vis}$ 数组标记，遇到访问过的状态可以直接返回 `false`。

另外有一个比较强的优化：由于字符串左括号和右括号的数目必须相同，因此字符串的长度为偶数，所以 $m+n-1$ 必须是偶数（提示 2）。我们可以在 DFS 之前就预先判断这一要求是否成立。

```py [sol-Python3]
class Solution:
    def hasValidPath(self, grid: List[List[str]]) -> bool:
        m, n = len(grid), len(grid[0])
        if (m + n) % 2 == 0 or grid[0][0] == ')' or grid[m - 1][n - 1] == '(':
            return False

        @cache  # 效果类似 vis 数组
        def dfs(x: int, y: int, c: int) -> bool:
            if c > m - x + n - y - 1:  # 剪枝：即使后面都是 ')' 也不能将 c 减为 0
                return False
            if x == m - 1 and y == n - 1:  # 终点
                return c == 1  # 终点一定是 ')'
            c += 1 if grid[x][y] == '(' else -1
            return c >= 0 and (x < m - 1 and dfs(x + 1, y, c) or y < n - 1 and dfs(x, y + 1, c))  # 往下或者往右

        return dfs(0, 0, 0)  # 起点
```

```java [sol-Java]
class Solution {
    private int m, n;
    private char[][] grid;
    private boolean[][][] vis;

    public boolean hasValidPath(char[][] grid) {
        m = grid.length;
        n = grid[0].length;
        if ((m + n) % 2 == 0 || grid[0][0] == ')' || grid[m - 1][n - 1] == '(') {
            return false;
        }

        this.grid = grid;
        vis = new boolean[m][n][(m + n + 1) / 2];
        return dfs(0, 0, 0); // 起点
    }

    private boolean dfs(int x, int y, int c) {
        if (c > m - x + n - y - 1) { // 剪枝：即使后面都是 ')' 也不能将 c 减为 0
            return false;
        }
        if (x == m - 1 && y == n - 1) { // 终点
            return c == 1; // 终点一定是 ')'
        }
        if (vis[x][y][c]) {
            return false;
        }
        vis[x][y][c] = true;
        c += grid[x][y] == '(' ? 1 : -1;
        return c >= 0 && (x < m - 1 && dfs(x + 1, y, c) || y < n - 1 && dfs(x, y + 1, c)); // 往下或者往右
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool hasValidPath(vector<vector<char>>& grid) {
        int m = grid.size(), n = grid[0].size();
        if ((m + n) % 2 == 0 || grid[0][0] == ')' || grid[m - 1][n - 1] == '(') {
            return false;
        }

        bool vis[m][n][(m + n + 1) / 2];
        memset(vis, 0, sizeof(vis));

        auto dfs = [&](this auto&& dfs, int x, int y, int c) -> bool {
            if (c > m - x + n - y - 1) { // 剪枝：即使后面都是 ')' 也不能将 c 减为 0
                return false;
            }
            if (x == m - 1 && y == n - 1) { // 终点
                return c == 1; // 终点一定是 ')'
            }
            if (vis[x][y][c]) {
                return false;
            }
            vis[x][y][c] = true;
            c += grid[x][y] == '(' ? 1 : -1;
            return c >= 0 && (x < m - 1 && dfs(x + 1, y, c) || y < n - 1 && dfs(x, y + 1, c)); // 往下或者往右
        };

        return dfs(0, 0, 0); // 起点
    }
};
```

```go [sol-Go]
func hasValidPath(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	if (m+n)%2 == 0 || grid[0][0] == ')' || grid[m-1][n-1] == '(' {
		return false
	}

	vis := make([][][]bool, m)
	for i := range vis {
		vis[i] = make([][]bool, n)
		for j := range vis[i] {
			vis[i][j] = make([]bool, (m+n+1)/2)
		}
	}
	var dfs func(int, int, int) bool
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

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn(m+n))$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。每个状态至多访问一次。
- 空间复杂度：$\mathcal{O}(mn(m+n))$。

## 注

值得注意的是，DFS 的写法相比某些递推的写法要快 $10$ 倍以上，这是因为有很多状态是无法访问到的：比如 $(x=2,y=3,c=100)$ 这个状态就是不可达的，此时还没走几步，$c$ 不可能这么大。或者对于一些随机的网格图，$c$ 的值也会比较小。这种情况下 DFS 的优势就发挥出来了，DFS 可以十分自然地遍历到所有合法的状态，加上自带的剪枝效果，可以大大降低访问到的状态数。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
