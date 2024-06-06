### 提示 1

按元素值从小往大计算。

### 提示 2

定义 $f[i][j]$ 表示到达 $\textit{mat}[i][j]$ 时，访问过的单元格的最大数量（包含 $\textit{mat}[i][j]$）。那么答案就是所有 $f[i][j]$ 的最大值。

如何计算 $f[i][j]$？从哪转移过来？

请注意，我们**不需要知道具体从哪个单元格转移过来，只需要知道转移来源的 $f$ 的最大值是多少**。

### 提示 3

按照元素值从小往大计算，那么第 $i$ 行的比 $\textit{mat}[i][j]$ 小的 $f$ 值都算出来了，大于等于 $\textit{mat}[i][j]$ 的尚未计算，视作 $0$。

所以对于第 $i$ 行，相当于取这一行的 $f$ 值的最大值，作为转移来源的值。我们用一个长为 $m$ 的数组 $\textit{rowMax}$ 维护每一行的 $f$ 值的最大值。

对于每一列，也同理，用一个长为 $n$ 的数组 $\textit{colMax}$ 维护。

所以有

$$
f[i][j] = \max(\textit{rowMax}[i], \textit{colMax}[j]) + 1
$$

这里加一是把 $\textit{mat}[i][j]$ 算上。

### 细节

代码实现时 $f[i][j]$ 可以省略，因为只需要知道每行每列的 $f$ 值的最大值。

对于相同元素，先计算出 $(i,j)$ 处的最大值，再更新到 $\textit{rowMax}$ 和 $\textit{colMax}$ 中。

```py [sol-Python3]
class Solution:
    def maxIncreasingCells(self, mat: List[List[int]]) -> int:
        g = defaultdict(list)
        for i, row in enumerate(mat):
            for j, x in enumerate(row):
                g[x].append((i, j))  # 相同元素放在同一组，统计位置

        row_max = [0] * len(mat)
        col_max = [0] * len(mat[0])
        for _, pos in sorted(g.items(), key=lambda p: p[0]):
            # 先把最大值算出来，再更新 row_max 和 col_max
            mx = [max(row_max[i], col_max[j]) + 1 for i, j in pos]
            for (i, j), f in zip(pos, mx):
                row_max[i] = max(row_max[i], f)  # 更新第 i 行的最大 f 值
                col_max[j] = max(col_max[j], f)  # 更新第 j 列的最大 f 值
        return max(row_max)
```

```java [sol-Java]
class Solution {
    public int maxIncreasingCells(int[][] mat) {
        int m = mat.length;
        int n = mat[0].length;
        TreeMap<Integer, List<int[]>> g = new TreeMap<>();
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                // 相同元素放在同一组，统计位置
                g.computeIfAbsent(mat[i][j], k -> new ArrayList<>()).add(new int[]{i, j});
            }
        }

        int ans = 0;
        int[] rowMax = new int[m];
        int[] colMax = new int[n];
        for (List<int[]> pos : g.values()) {
            int[] mx = new int[pos.size()]; // 先把最大值算出来，再更新 rowMax 和 colMax
            for (int k = 0; k < pos.size(); k++) {
                int[] p = pos.get(k);
                int i = p[0];
                int j = p[1];
                mx[k] = Math.max(rowMax[i], colMax[j]) + 1;
                ans = Math.max(ans, mx[k]);
            }
            for (int k = 0; k < pos.size(); k++) {
                int[] p = pos.get(k);
                int i = p[0];
                int j = p[1];
                rowMax[i] = Math.max(rowMax[i], mx[k]); // 更新第 i 行的最大 f 值
                colMax[j] = Math.max(colMax[j], mx[k]); // 更新第 j 列的最大 f 值
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxIncreasingCells(vector<vector<int>>& mat) {
        int m = mat.size(), n = mat[0].size();
        map<int, vector<pair<int, int>>> g;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                g[mat[i][j]].emplace_back(i, j); // 相同元素放在同一组，统计位置
            }
        }

        vector<int> row_max(m), col_max(n);
        for (auto& [_, pos] : g) {
            vector<int> mx; // 先把最大值算出来，再更新 row_max 和 col_max
            for (auto& [i, j] : pos) {
                mx.push_back(max(row_max[i], col_max[j]) + 1);
            }
            for (int k = 0; k < pos.size(); k++) {
                auto& [i, j] = pos[k];
                row_max[i] = max(row_max[i], mx[k]); // 更新第 i 行的最大 f 值
                col_max[j] = max(col_max[j], mx[k]); // 更新第 j 列的最大 f 值
            }
        }
        return ranges::max(row_max);
    }
};
```

```go [sol-Go]
func maxIncreasingCells(mat [][]int) int {
    type pair struct{ x, y int }
    g := map[int][]pair{}
    for i, row := range mat {
        for j, x := range row {
            g[x] = append(g[x], pair{i, j}) // 相同元素放在同一组，统计位置
        }
    }
    keys := make([]int, 0, len(g))
    for k := range g {
        keys = append(keys, k)
    }
    slices.Sort(keys)

    rowMax := make([]int, len(mat))
    colMax := make([]int, len(mat[0]))
    for _, x := range keys {
        pos := g[x]
        mx := make([]int, len(pos))
        for i, p := range pos {
            mx[i] = max(rowMax[p.x], colMax[p.y]) + 1 // 先把最大值算出来，再更新 rowMax 和 colMax
        }
        for i, p := range pos {
            rowMax[p.x] = max(rowMax[p.x], mx[i]) // 更新第 p.x 行的最大 f 值
            colMax[p.y] = max(colMax[p.y], mx[i]) // 更新第 p.y 列的最大 f 值
        }
    }
    return slices.Max(rowMax)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn\log (mn))$，其中 $m$ 和 $n$ 分别为 $\textit{mat}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。

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
