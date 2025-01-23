## 方法一：记忆化搜索

`floor(coins[i] / 2)` 等价于 `coins[i] >> 1`。

右移运算是可以叠加的，即 `(x >> 1) >> 1` 等于 `x >> 2`。

我们可以在递归的过程中，额外记录从根节点递归到当前节点的过程中，一共执行了多少次右移，也就是子树中的每个节点值需要右移的次数。

故定义 $\textit{dfs}(i,j)$ 表示递归到以 $i$ 为根的子树，在上面已经执行了 $j$ 次右移的前提下，我们在这棵子树中最多可以得到多少积分。

用「选或不选」来思考，即是否执行右移：

- 不右移：答案为 $(\textit{coins}[i]\ \texttt{>>}\ j)-k$ 加上 $i$ 的每个子树 $\textit{ch}$ 的 $\textit{dfs}(\textit{ch},j)$。
- 右移：答案为 $\textit{coins}[i]\ \texttt{>>}\ (j+1)$ 加上 $i$ 的每个子树 $\textit{ch}$ 的 $\textit{dfs}(\textit{ch},j+1)$。

两种情况取最大值，得

$$
\textit{dfs}(i,j) = \max
\begin{cases}
(\textit{coins}[i]\ \texttt{>>}\ j)-k + \sum_{\textit{ch}} \textit{dfs}(\textit{ch},j)    \\
(\textit{coins}[i]\ \texttt{>>}\ (j+1)) +  \sum_{\textit{ch}} \textit{dfs}(\textit{ch},j+1)    \\
\end{cases}
$$

递归入口：$\textit{dfs}(0,0)$。其中 $i=0$ 表示根节点。一开始没有执行右移，所以 $j=0$。

### 细节

一个数最多右移多少次，就变成 $0$ 了？

设 $w$ 是 $\textit{coins}[i]$ 的二进制长度，那么 $\textit{coins}[i]$ 右移 $w$ 次后就是 $0$ 了。

在本题的数据范围下，$w\le 14$。

所以如果在递归过程中发现 $j+1 = 14$，就不执行右移，因为此时 $\textit{dfs}(\textit{ch},j+1)$ 子树中的每个节点值都要右移 $14$ 次，算出的结果一定是 $0$。既然都知道递归的结果了，那就不需要递归了。

此外，为避免错把父亲当作儿子，可以额外传入 $\textit{fa}$ 表示父节点，遍历 $i$ 的邻居时，跳过邻居节点是 $\textit{fa}$ 的情况。

关于记忆化搜索的原理，请看视频讲解 [动态规划入门：从记忆化搜索到递推](https://b23.tv/72onpYq)。

[本题视频讲解](https://www.bilibili.com/video/BV1tw411q7VZ/?t=45m18s)

```py [sol-Python3]
class Solution:
    def maximumPoints(self, edges: List[List[int]], coins: List[int], k: int) -> int:
        g = [[] for _ in coins]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, j: int, fa: int) -> int:
            res1 = (coins[i] >> j) - k
            res2 = coins[i] >> (j + 1)
            for ch in g[i]:
                if ch != fa:
                    res1 += dfs(ch, j, i)  # 不右移
                    if j < 13:  # j+1 >= 14 相当于 res2 += 0，无需递归
                        res2 += dfs(ch, j + 1, i)  # 右移
            return max(res1, res2)

        return dfs(0, 0, -1)
```

```java [sol-Java]
class Solution {
    public int maximumPoints(int[][] edges, int[] coins, int k) {
        int n = coins.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        int[][] memo = new int[n][14];
        for (int[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }
        return dfs(0, 0, -1, memo, g, coins, k);
    }

    private int dfs(int i, int j, int fa, int[][] memo, List<Integer>[] g, int[] coins, int k) {
        if (memo[i][j] != -1) { // 之前计算过
            return memo[i][j];
        }
        int res1 = (coins[i] >> j) - k;
        int res2 = coins[i] >> (j + 1);
        for (int ch : g[i]) {
            if (ch == fa) continue;
            res1 += dfs(ch, j, i, memo, g, coins, k); // 不右移
            if (j < 13) { // j+1 >= 14 相当于 res2 += 0，无需递归
                res2 += dfs(ch, j + 1, i, memo, g, coins, k); // 右移
            }
        }
        return memo[i][j] = Math.max(res1, res2); // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumPoints(vector<vector<int>>& edges, vector<int>& coins, int k) {
        int n = coins.size();
        vector<vector<int>> g(n);
        for (auto& e: edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        array<int, 14> init_val;
        ranges::fill(init_val, -1); // -1 表示没有计算过
        vector memo(n, init_val);
        auto dfs = [&](this auto&& dfs, int i, int j, int fa) {
            int& res = memo[i][j]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }
            int res1 = (coins[i] >> j) - k;
            int res2 = coins[i] >> (j + 1);
            for (int ch : g[i]) {
                if (ch == fa) continue;
                res1 += dfs(ch, j, i); // 不右移
                if (j < 13) { // j+1 >= 14 相当于 res2 += 0，无需递归
                    res2 += dfs(ch, j + 1, i); // 右移
                }
            }
            return res = max(res1, res2); // 记忆化
        };
        return dfs(0, 0, -1);
    }
};
```

```go [sol-Go]
func maximumPoints(edges [][]int, coins []int, k int) int {
    n := len(coins)
    g := make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }

    memo := make([][14]int, n)
    for i := range memo {
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    var dfs func(int, int, int) int
    dfs = func(i, j, fa int) (res int) {
        p := &memo[i][j]
        if *p != -1 {
            return *p
        }
        defer func() { *p = res }()
        res1 := coins[i]>>j - k
        res2 := coins[i] >> (j + 1)
        for _, ch := range g[i] {
            if ch != fa {
                res1 += dfs(ch, j, i) // 不右移
                if j < 13 { // j+1 >= 14 相当于 res2 += 0 无需递归
                    res2 += dfs(ch, j+1, i) // 右移
                }
            }
        }
        return max(res1, res2)
    }
    return dfs(0, 0, -1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{coins}$ 的长度，$U=\max(\textit{coins})$。
- 空间复杂度：$\mathcal{O}(n\log U)$。

## 方法二：递推

类似把记忆化搜索 1:1 翻译成递推的过程，我们也可以从下往上算。

去掉参数 $j$，改成每个节点 $i$ 返回一个长为 $14$ 的列表 $f_i$，其中 $f_i[j]$ 对应上面 $\textit{dfs}(i,j)$ 的计算结果。

递推式为

$$
f_i[j] = \max
\begin{cases}
(\textit{coins}[i]\ \texttt{>>}\ j)-k + \sum_{\textit{ch}} f_{\textit{ch}}[j]   \\
(\textit{coins}[i]\ \texttt{>>}\ (j+1)) +  \sum_{\textit{ch}} f_{\textit{ch}}[j+1]    \\
\end{cases}
$$

把 $\sum_{\textit{ch}} f_{\textit{ch}}[j]$ 累加到 $s[j]$ 中，上式为

$$
f_i[j] = \max
\begin{cases}
(\textit{coins}[i]\ \texttt{>>}\ j)-k + s[j]   \\
(\textit{coins}[i]\ \texttt{>>}\ (j+1)) +  s[j+1]    \\
\end{cases}
$$

特判 $j=13$ 的情况，上式为

$$
f_i[13] = \max((\textit{coins}[i]\ \texttt{>>}\ 13)-k + s[13],0)
$$

代码实现时，可以直接把算出的结果原地保存到 $s$ 数组中。

```py [sol-Python3]
class Solution:
    def maximumPoints(self, edges: List[List[int]], coins: List[int], k: int) -> int:
        g = [[] for _ in coins]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        def dfs(x: int, fa: int) -> List[int]:
            s = [0] * 14
            for y in g[x]:
                if y != fa:
                    fy = dfs(y, x)
                    for j, v in enumerate(fy):
                        s[j] += v
            for j in range(13):
                s[j] = max((coins[x] >> j) - k + s[j], (coins[x] >> (j + 1)) + s[j + 1])
            s[13] = max(s[13] + (coins[x] >> 13) - k, 0)
            return s

        return dfs(0, -1)[0]
```

```java [sol-Java]
class Solution {
    public int maximumPoints(int[][] edges, int[] coins, int k) {
        List<Integer>[] g = new ArrayList[coins.length];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }
        return dfs(0, -1, g, coins, k)[0];
    }

    private int[] dfs(int x, int fa, List<Integer>[] g, int[] coins, int k) {
        int[] s = new int[14];
        for (int y : g[x]) {
            if (y == fa) continue;
            int[] fy = dfs(y, x, g, coins, k);
            for (int j = 0; j < 14; j++) {
                s[j] += fy[j];
            }
        }
        for (int j = 0; j < 13; j++) {
            s[j] = Math.max((coins[x] >> j) - k + s[j], (coins[x] >> (j + 1)) + s[j + 1]);
        }
        s[13] = Math.max(s[13] + (coins[x] >> 13) - k, 0);
        return s;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumPoints(vector<vector<int>>& edges, vector<int>& coins, int k) {
        vector<vector<int>> g(coins.size());
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        auto dfs = [&](this auto&& dfs, int x, int fa) -> array<int, 14> {
            array<int, 14> s{};
            for (int y : g[x]) {
                if (y == fa) continue;
                auto fy = dfs(y, x);
                for (int j = 0; j < 14; j++) {
                    s[j] += fy[j];
                }
            }
            for (int j = 0; j < 13; j++) {
                s[j] = max((coins[x] >> j) - k + s[j], (coins[x] >> (j + 1)) + s[j + 1]);
            }
            s[13] = max(s[13] + (coins[x] >> 13) - k, 0);
            return s;
        };
        return dfs(0, -1)[0];
    }
};
```

```go [sol-Go]
func maximumPoints(edges [][]int, coins []int, k int) int {
    n := len(coins)
    g := make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }

    var dfs func(int, int) [14]int
    dfs = func(x, fa int) (s [14]int) {
        for _, y := range g[x] {
            if y != fa {
                fy := dfs(y, x)
                for j, v := range fy {
                    s[j] += v
                }
            }
        }
        for j := range 13 {
            s[j] = max(coins[x]>>j-k+s[j], coins[x]>>(j+1)+s[j+1])
        }
        s[13] = max(s[13]+coins[x]>>13-k, 0)
        return
    }
    return dfs(0, -1)[0]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{coins}$ 的长度，$U=\max(\textit{coins})$。
- 空间复杂度：$\mathcal{O}(n\log U)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
