下午两点直播讲题，记得关注哦~（见个人主页）

---

### 提示 1

按元素值从小往大计算。

### 提示 2

定义 $f[i][j]$ 表示到达 $\textit{mat}[i][j]$ 时所访问的单元格的最大数量。那么答案就是所有 $f[i][j]$ 的最大值。

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

代码实现时 $f[i][j]$ 可以省略，因为只需要每行每列的 $f$ 值的最大值。

对于相同元素，在全部计算出最大值后，再更新到 $\textit{rowMax}$ 和 $\textit{colMax}$ 中。

```py [sol-Python3]
class Solution:
    def maxIncreasingCells(self, mat: List[List[int]]) -> int:
        g = defaultdict(list)
        for i, row in enumerate(mat):
            for j, x in enumerate(row):
                g[x].append((i, j))  # 相同元素放在同一组，统计位置

        ans = 0
        row_max = [0] * len(mat)
        col_max = [0] * len(mat[0])
        for _, pos in sorted(g.items(), key=lambda p: p[0]):
            # 先把最大值算出来，再更新 row_max 和 col_max
            mx = [max(row_max[i], col_max[j]) + 1 for i, j in pos]
            ans = max(ans, max(mx))
            for (i, j), f in zip(pos, mx):
                row_max[i] = max(row_max[i], f)  # 更新第 i 行的最大 f 值
                col_max[j] = max(col_max[j], f)  # 更新第 j 列的最大 f 值
        return ans
```

```java [sol-Java]
class Solution {
    public int maxIncreasingCells(int[][] mat) {
        var g = new TreeMap<Integer, List<int[]>>();
        int m = mat.length, n = mat[0].length;
        for (int i = 0; i < m; i++)
            for (int j = 0; j < n; j++)
                // 相同元素放在同一组，统计位置
                g.computeIfAbsent(mat[i][j], k -> new ArrayList<>()).add(new int[]{i, j});

        int ans = 0;
        int[] rowMax = new int[m], colMax = new int[n];
        for (var pos : g.values()) {
            var mx = new int[pos.size()];  // 先把最大值算出来，再更新 rowMax 和 colMax
            for (int i = 0; i < pos.size(); i++) {
                mx[i] = Math.max(rowMax[pos.get(i)[0]], colMax[pos.get(i)[1]]) + 1;
                ans = Math.max(ans, mx[i]);
            }
            for (int k = 0; k < pos.size(); k++) {
                int i = pos.get(k)[0], j = pos.get(k)[1];
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
    int maxIncreasingCells(vector<vector<int>> &mat) {
        map<int, vector<pair<int, int>>> g;
        int m = mat.size(), n = mat[0].size();
        for (int i = 0; i < m; i++)
            for (int j = 0; j < n; j++)
                g[mat[i][j]].emplace_back(i, j); // 相同元素放在同一组，统计位置

        int ans = 0;
        vector<int> row_max(m), col_max(n);
        for (auto &[_, pos]: g) {
            vector<int> mx; // 先把最大值算出来，再更新 row_max 和 col_max
            for (auto &[i, j]: pos) {
                mx.push_back(max(row_max[i], col_max[j]) + 1);
                ans = max(ans, mx.back());
            }
            for (int k = 0; k < pos.size(); k++) {
                auto &[i, j] = pos[k];
                row_max[i] = max(row_max[i], mx[k]); // 更新第 i 行的最大 f 值
                col_max[j] = max(col_max[j], mx[k]); // 更新第 j 列的最大 f 值
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxIncreasingCells(mat [][]int) (ans int) {
	type pair struct{ x, y int }
	g := map[int][]pair{}
	for i, row := range mat {
		for j, x := range row {
			g[x] = append(g[x], pair{i, j}) // 相同元素放在同一组，统计位置
		}
	}
	a := make([]int, 0, len(g))
	for k := range g {
		a = append(a, k)
	}
	sort.Ints(a) // 从小到大

	rowMax := make([]int, len(mat))
	colMax := make([]int, len(mat[0]))
	for _, x := range a {
		pos := g[x]
		mx := make([]int, len(pos))
		for i, p := range pos {
			mx[i] = max(rowMax[p.x], colMax[p.y]) + 1 // 先把最大值算出来，再更新 rowMax 和 colMax
			ans = max(ans, mx[i])
		}
		for i, p := range pos {
			rowMax[p.x] = max(rowMax[p.x], mx[i]) // 更新第 p.x 行的最大 f 值
			colMax[p.y] = max(colMax[p.y], mx[i]) // 更新第 p.y 列的最大 f 值
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn\log (mn))$，其中 $m$ 和 $n$ 分别为 $\textit{mat}$ 的行数和列数。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(mn)$。
