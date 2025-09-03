## 方法一：暴力计算

计算 $\textit{topLeft}[i][j]$：从 $(i-1,j-1)$ 开始，往左上方向 ↖ 遍历，直到到达矩阵边界。在这个过程中，把遍历到的数加到哈希集合中，最终 $\textit{topLeft}[i][j]$ 就是哈希集合的大小。

计算 $\textit{bottomRight}[i][j]$：从 $(i+1,j+1)$ 开始，往右下方向 ↘ 遍历，直到到达矩阵边界。在这个过程中，把遍历到的数加到哈希集合中，最终 $\textit{bottomRight}[i][j]$ 就是哈希集合的大小。

```py [sol-Python3]
class Solution:
    def differenceOfDistinctValues(self, grid: List[List[int]]) -> List[List[int]]:
        m, n = len(grid), len(grid[0])
        ans = [[0] * n for _ in range(m)]
        st = set()
        for i in range(m):
            for j in range(n):
                # 计算 top_left[i][j]
                st.clear()
                x, y = i - 1, j - 1
                while x >= 0 and y >= 0:
                    st.add(grid[x][y])
                    x -= 1
                    y -= 1
                top_left = len(st)

                # 计算 bottom_right[i][j]
                st.clear()
                x, y = i + 1, j + 1
                while x < m and y < n:
                    st.add(grid[x][y])
                    x += 1
                    y += 1
                bottom_right = len(st)

                ans[i][j] = abs(top_left - bottom_right)
        return ans
```

```java [sol-Java]
class Solution {
    public int[][] differenceOfDistinctValues(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        int[][] ans = new int[m][n];
        Set<Integer> st = new HashSet<>();
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                // 计算 topLeft[i][j]
                st.clear();
                for (int x = i - 1, y = j - 1; x >= 0 && y >= 0; x--, y--) {
                    st.add(grid[x][y]);
                }
                int topLeft = st.size();

                // 计算 bottomRight[i][j]
                st.clear();
                for (int x = i + 1, y = j + 1; x < m && y < n; x++, y++) {
                    st.add(grid[x][y]);
                }
                int bottomRight = st.size();

                ans[i][j] = Math.abs(topLeft - bottomRight);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> differenceOfDistinctValues(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        vector ans(m, vector<int>(n));
        unordered_set<int> st;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                // 计算 top_left[i][j]
                st.clear();
                for (int x = i - 1, y = j - 1; x >= 0 && y >= 0; x--, y--) {
                    st.insert(grid[x][y]);
                }
                int top_left = st.size();

                // 计算 bottom_right[i][j]
                st.clear();
                for (int x = i + 1, y = j + 1; x < m && y < n; x++, y++) {
                    st.insert(grid[x][y]);
                }
                int bottom_right = st.size();

                ans[i][j] = abs(top_left - bottom_right);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func differenceOfDistinctValues(grid [][]int) [][]int {
    m, n := len(grid), len(grid[0])
    ans := make([][]int, m)
    set := map[int]struct{}{}
    for i := range m {
        ans[i] = make([]int, n)
        for j := range n {
            // 计算 topLeft[i][j]
            clear(set)
            for x, y := i-1, j-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
                set[grid[x][y]] = struct{}{}
            }
            topLeft := len(set)

            // 计算 bottomRight[i][j]
            clear(set)
            for x, y := i+1, j+1; x < m && y < n; x, y = x+1, y+1 {
                set[grid[x][y]] = struct{}{}
            }
            bottomRight := len(set)

            ans[i][j] = abs(topLeft - bottomRight)
        }
    }
    return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn\cdot \min(m,n))$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。计算一个 $\textit{topLeft}[i][j]$ 和 $\textit{bottomRight}[i][j]$ 需要 $\mathcal{O}(\min(m,n))$ 的时间。
- 空间复杂度：$\mathcal{O}(\min(m,n))$。返回值不计入。

## 方法二：前后缀分解

对于同一条对角线，方法一会多次遍历。比如计算 $\textit{ans}[0][0]$ 的时候我们会遍历主对角线，计算 $\textit{ans}[1][1]$ 的时候我们又会遍历主对角线。怎么减少遍历次数？

考察某一条对角线 $d$，把它视作一个一维数组。对于 $d[i]$ 来说：

- $\textit{topLeft}$ 是在 $d[i]$ 左边的不同元素个数。我们可以从左到右遍历 $d$，同时把元素加到一个哈希集合中。遍历到 $d[i]$ 时，哈希集合的大小就是 $\textit{topLeft}$。
- $\textit{bottomRight}$ 是在 $d[i]$ 右边的不同元素个数。我们可以从右到左遍历 $d$，同时把元素加到一个哈希集合中。遍历到 $d[i]$ 时，哈希集合的大小就是 $\textit{bottomRight}$。

如何一条一条地枚举对角线？见[【模板】遍历对角线](https://leetcode.cn/problems/sort-matrix-by-diagonals/solutions/3068709/mo-ban-mei-ju-dui-jiao-xian-pythonjavacg-pjxp/)。

代码实现时，可以直接把 $\textit{topLeft}$ 和 $\textit{bottomRight}$ 保存到 $\textit{ans}[i][j]$ 中。

```py [sol-Python3]
class Solution:
    def differenceOfDistinctValues(self, grid: List[List[int]]) -> List[List[int]]:
        m, n = len(grid), len(grid[0])
        ans = [[0] * n for _ in range(m)]
        st = set()

        # 第一排在右上，最后一排在左下
        # 每排从左上到右下
        # 令 k=i-j+n，那么右上角 k=1，左下角 k=m+n-1
        for k in range(1, m + n):
            # 核心：计算 j 的最小值和最大值
            min_j = max(n - k, 0) # i=0 的时候，j=n-k，但不能是负数
            max_j = min(m + n - 1 - k, n - 1)  # i=m-1 的时候，j=m+n-1-k，但不能超过 n-1

            st.clear()
            for j in range(min_j, max_j + 1):
                i = k + j - n
                ans[i][j] = len(st)  # top_left[i][j] == len(st)
                st.add(grid[i][j])

            st.clear()
            for j in range(max_j, min_j - 1, -1):
                i = k + j - n
                ans[i][j] = abs(ans[i][j] - len(st))  # bottom_right[i][j] == len(st)
                st.add(grid[i][j])
        return ans
```

```java [sol-Java]
class Solution {
    public int[][] differenceOfDistinctValues(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        int[][] ans = new int[m][n];
        Set<Integer> set = new HashSet<>();

        // 第一排在右上，最后一排在左下
        // 每排从左上到右下
        // 令 k=i-j+n，那么右上角 k=1，左下角 k=m+n-1
        for (int k = 1; k < m + n; k++) {
            // 核心：计算 j 的最小值和最大值
            int minJ = Math.max(n - k, 0); // i=0 的时候，j=n-k，但不能是负数
            int maxJ = Math.min(m + n - 1 - k, n - 1); // i=m-1 的时候，j=m+n-1-k，但不能超过 n-1

            set.clear();
            for (int j = minJ; j <= maxJ; j++) {
                int i = k + j - n;
                ans[i][j] = set.size(); // topLeft[i][j] == set.size()
                set.add(grid[i][j]);
            }

            set.clear();
            for (int j = maxJ; j >= minJ; j--) {
                int i = k + j - n;
                ans[i][j] = Math.abs(ans[i][j] - set.size()); // bottomRight[i][j] == set.size()
                set.add(grid[i][j]);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> differenceOfDistinctValues(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        vector ans(m, vector<int>(n));
        unordered_set<int> st;

        // 第一排在右上，最后一排在左下
        // 每排从左上到右下
        // 令 k=i-j+n，那么右上角 k=1，左下角 k=m+n-1
        for (int k = 1; k < m + n; k++) {
            // 核心：计算 j 的最小值和最大值
            int min_j = max(n - k, 0); // i=0 的时候，j=n-k，但不能是负数
            int max_j = min(m + n - 1 - k, n - 1); // i=m-1 的时候，j=m+n-1-k，但不能超过 n-1

            st.clear();
            for (int j = min_j; j <= max_j; j++) {
                int i = k + j - n;
                ans[i][j] = st.size(); // top_left[i][j] == st.size()
                st.insert(grid[i][j]);
            }

            st.clear();
            for (int j = max_j; j >= min_j; j--) {
                int i = k + j - n;
                ans[i][j] = abs(ans[i][j] - (int) st.size()); // bottom_right[i][j] == st.size()
                st.insert(grid[i][j]);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func differenceOfDistinctValues(grid [][]int) [][]int {
    m, n := len(grid), len(grid[0])
    ans := make([][]int, m)
    for i := range ans {
        ans[i] = make([]int, n)
    }
    set := map[int]struct{}{}

    // 第一排在右上，最后一排在左下
    // 每排从左上到右下
    // 令 k=i-j+n，那么右上角 k=1，左下角 k=m+n-1
    for k := 1; k < m+n; k++ {
        // 核心：计算 j 的最小值和最大值
        minJ := max(n-k, 0)       // i=0 的时候，j=n-k，但不能是负数
        maxJ := min(m+n-1-k, n-1) // i=m-1 的时候，j=m+n-1-k，但不能超过 n-1

        clear(set)
        for j := minJ; j <= maxJ; j++ {
            i := k + j - n
            ans[i][j] = len(set) // topLeft[i][j] == len(set)
            set[grid[i][j]] = struct{}{}
        }

        clear(set)
        for j := maxJ; j >= minJ; j-- {
            i := k + j - n
            ans[i][j] = abs(ans[i][j] - len(set)) // bottomRight[i][j] == len(set)
            set[grid[i][j]] = struct{}{}
        }
    }
    return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。每个单元格访问两次。
- 空间复杂度：$\mathcal{O}(\min(m,n))$。返回值不计入。

## 方法三：位运算优化

本题值域范围为 $[1,50]$，我们可以用 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/) 中的技巧，用二进制数（$64$ 位整数）记录哪些数字出现过。遍历到 $x=\textit{grid}[i][j]$ 时，方法二要加到集合中，方法三改成把二进制数的从低到高第 $x$ 位置为 $1$。集合大小就是二进制数中的 $1$ 的个数。

```py [sol-Python3]
class Solution:
    def differenceOfDistinctValues(self, grid: List[List[int]]) -> List[List[int]]:
        m, n = len(grid), len(grid[0])
        ans = [[0] * n for _ in range(m)]

        for k in range(1, m + n):
            min_j = max(n - k, 0)
            max_j = min(m + n - 1 - k, n - 1)

            st = 0
            for j in range(min_j, max_j + 1):
                i = k + j - n
                ans[i][j] = st.bit_count()  # 计算 st 中 1 的个数
                st |= 1 << grid[i][j]  # 把 grid[i][j] 加到 st 中

            st = 0
            for j in range(max_j, min_j - 1, -1):
                i = k + j - n
                ans[i][j] = abs(ans[i][j] - st.bit_count())
                st |= 1 << grid[i][j]
        return ans
```

```java [sol-Java]
class Solution {
    public int[][] differenceOfDistinctValues(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        int[][] ans = new int[m][n];

        for (int k = 1; k < m + n; k++) {
            int minJ = Math.max(n - k, 0);
            int maxJ = Math.min(m + n - 1 - k, n - 1);

            long set = 0;
            for (int j = minJ; j <= maxJ; j++) {
                int i = k + j - n;
                ans[i][j] = Long.bitCount(set); // 计算 set 中 1 的个数
                set |= 1L << grid[i][j]; // 把 grid[i][j] 加到 set 中
            }

            set = 0;
            for (int j = maxJ; j >= minJ; j--) {
                int i = k + j - n;
                ans[i][j] = Math.abs(ans[i][j] - Long.bitCount(set));
                set |= 1L << grid[i][j];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> differenceOfDistinctValues(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        vector ans(m, vector<int>(n));

        for (int k = 1; k < m + n; k++) {
            int min_j = max(n - k, 0);
            int max_j = min(m + n - 1 - k, n - 1);

            uint64_t st = 0;
            for (int j = min_j; j <= max_j; j++) {
                int i = k + j - n;
                ans[i][j] = popcount(st); // st 的大小
                st |= 1ULL << grid[i][j]; // 把 grid[i][j] 加到 st 中
            }

            st = 0;
            for (int j = max_j; j >= min_j; j--) {
                int i = k + j - n;
                ans[i][j] = abs(ans[i][j] - popcount(st));
                st |= 1ULL << grid[i][j];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func differenceOfDistinctValues(grid [][]int) [][]int {
    m, n := len(grid), len(grid[0])
    ans := make([][]int, m)
    for i := range ans {
        ans[i] = make([]int, n)
    }

    for k := 1; k < m+n; k++ {
        minJ := max(n-k, 0)
        maxJ := min(m+n-1-k, n-1)

        set := uint(0)
        for j := minJ; j <= maxJ; j++ {
            i := k + j - n
            ans[i][j] = bits.OnesCount(set) // set 的大小
            set |= 1 << grid[i][j] // 把 grid[i][j] 加到 set 中
        }

        set = 0
        for j := maxJ; j >= minJ; j-- {
            i := k + j - n
            ans[i][j] = abs(ans[i][j] - bits.OnesCount(set))
            set |= 1 << grid[i][j]
        }
    }
    return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。每个单元格访问两次。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

更多相似题目，见下面动态规划题单中的「**专题：前后缀分解**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
