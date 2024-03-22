## 方法一：单调栈优化 DP

### 前置知识

1. [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://b23.tv/72onpYq)
2. [单调栈【基础算法精讲 26】](https://www.bilibili.com/video/BV1VN411J7S7/)
3. [二分查找【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)

### 思路

题目要计算从 $(0,0)$ 出发，到达 $(m-1,n-1)$ 经过的最少格子数。假设我们移动到了 $(i,j)$，那么问题就变成从 $(i,j)$ 到 $(m-1,n-1)$ 经过的最少格子数，这是一个和原问题相似的子问题。

定义 $f[i][j]$ 表示从 $(i,j)$ 到 $(m-1,n-1)$ 经过的最少格子数。

- 从 $(i,j)$ 向**右**移动一次，到达 $(i,k)$，问题变成从 $(i,k)$ 到 $(m-1,n-1)$ 经过的最少格子数，即 $f[i][j] = f[i][k] + 1$。
- 从 $(i,j)$ 向**下**移动一次，到达 $(k,j)$，问题变成从 $(k,j)$ 到 $(m-1,n-1)$ 经过的最少格子数，即 $f[i][j] = f[k][j] + 1$。

设 $g=\textit{grid}[i][j]$，有

$$
f[i][j] = \min\left\{\min_{k=j+1}^{j+g} f[i][k], \min_{k=i+1}^{i+g} f[k][j]\right\} + 1
$$

倒序枚举 $i$ 和 $j$ 计算。答案为 $f[0][0]$。

这样做时间复杂度为 $\mathcal{O}(mn(m+n))$，太慢了。

### 优化

由于有「区间查询」和「单点更新」，我们可以用线段树来优化。

但其实不需要线段树，有更轻量级的做法。

对于 $\min\limits_{k=j+1}^{j+g} f[i][k]$，在倒序枚举 $j$ 时，$k$ 的左边界 $j+1$ 是在单调减小的，但右边界没有单调性。联想到滑动窗口最小值的做法，我们用一个 $f$ 值底小顶大的**单调栈**来维护 $f[i][j]$ 及其下标 $j$。由于 $j$ 是倒序枚举，单调栈中的下标是底大顶小的，从那么在单调栈上**二分查找**最大的不超过 $j+g$ 的下标 $k$，对应的 $f[i][k]$ 就是要计算的最小值。

对于 $\min\limits_{k=i+1}^{i+g} f[k][j]$ 也同理，每一列都需要维护一个单调栈。

```py [sol-Python3]
class Solution:
    def minimumVisitedCells(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        col_stacks = [[] for _ in range(n)]  # 每列的单调栈
        for i in range(m - 1, -1, -1):
            row_st = []  # 当前行的单调栈
            for j in range(n - 1, -1, -1):
                g = grid[i][j]
                col_st = col_stacks[j]
                mn = inf if i < m - 1 or j < n - 1 else 1
                if g:  # 可以向右/向下跳
                    # 在单调栈上二分查找最优转移来源
                    k = bisect_left(row_st, -(j + g), key=lambda p: p[1])
                    if k < len(row_st):
                        mn = row_st[k][0] + 1
                    k = bisect_left(col_st, -(i + g), key=lambda p: p[1])
                    if k < len(col_st):
                        mn = min(mn, col_st[k][0] + 1)
                if mn < inf:
                    # 插入单调栈
                    while row_st and mn <= row_st[-1][0]:
                        row_st.pop()
                    row_st.append((mn, -j))  # 保证下标单调递增，方便调用 bisect_left
                    while col_st and mn <= col_st[-1][0]:
                        col_st.pop()
                    col_st.append((mn, -i))  # 保证下标单调递增，方便调用 bisect_left
        return mn if mn < inf else -1  # 最后一个算出的 mn 就是 f[0][0]
```

```java [sol-Java]
class Solution {
    public int minimumVisitedCells(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        int mn = 0;
        List<int[]>[] colStacks = new ArrayList[n]; // 每列的单调栈，为了能二分用 ArrayList
        Arrays.setAll(colStacks, i -> new ArrayList<int[]>());
        List<int[]> rowSt = new ArrayList<>(); // 行单调栈
        for (int i = m - 1; i >= 0; i--) {
            rowSt.clear();
            for (int j = n - 1; j >= 0; j--) {
                int g = grid[i][j];
                List<int[]> colSt = colStacks[j];
                mn = i < m - 1 || j < n - 1 ? Integer.MAX_VALUE : 1;
                if (g > 0) { // 可以向右/向下跳
                    // 在单调栈上二分查找最优转移来源
                    int k = search(rowSt, j + g);
                    if (k < rowSt.size()) {
                        mn = rowSt.get(k)[0] + 1;
                    }
                    k = search(colSt, i + g);
                    if (k < colSt.size()) {
                        mn = Math.min(mn, colSt.get(k)[0] + 1);
                    }
                }
                if (mn < Integer.MAX_VALUE) {
                    // 插入单调栈
                    while (!rowSt.isEmpty() && mn <= rowSt.get(rowSt.size() - 1)[0]) {
                        rowSt.remove(rowSt.size() - 1);
                    }
                    rowSt.add(new int[]{mn, j});
                    while (!colSt.isEmpty() && mn <= colSt.get(colSt.size() - 1)[0]) {
                        colSt.remove(colSt.size() - 1);
                    }
                    colSt.add(new int[]{mn, i});
                }
            }
        }
        return mn < Integer.MAX_VALUE ? mn : -1; // 最后一个算出的 mn 就是 f[0][0]
    }

    // 开区间二分，见 https://www.bilibili.com/video/BV1AP41137w7/
    private int search(List<int[]> st, int target) {
        int left = -1, right = st.size(); // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            int mid = left + (right - left) / 2;
            if (st.get(mid)[1] <= target) {
                right = mid; // 范围缩小到 (left, mid)
            } else {
                left = mid; // 范围缩小到 (mid, right)
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumVisitedCells(vector<vector<int>> &grid) {
        int m = grid.size(), n = grid[0].size(), mn;
        vector<vector<pair<int, int>>> col_stacks(n); // 每列的单调栈
        vector<pair<int, int>> row_st; // 行单调栈
        for (int i = m - 1; i >= 0; i--) {
            row_st.clear();
            for (int j = n - 1; j >= 0; j--) {
                int g = grid[i][j];
                auto &col_st = col_stacks[j];
                mn = i < m - 1 || j < n - 1 ? INT_MAX : 1;
                if (g) { // 可以向右/向下跳
                    // 在单调栈上二分查找最优转移来源
                    auto it = lower_bound(row_st.begin(), row_st.end(), j + g, [](const auto &a, const int b) {
                        return a.second > b;
                    });
                    if (it < row_st.end()) mn = it->first + 1;
                    it = lower_bound(col_st.begin(), col_st.end(), i + g, [](const auto &a, const int b) {
                        return a.second > b;
                    });
                    if (it < col_st.end()) mn = min(mn, it->first + 1);
                }
                if (mn < INT_MAX) {
                    // 插入单调栈
                    while (!row_st.empty() && mn <= row_st.back().first) {
                        row_st.pop_back();
                    }
                    row_st.emplace_back(mn, j);
                    while (!col_st.empty() && mn <= col_st.back().first) {
                        col_st.pop_back();
                    }
                    col_st.emplace_back(mn, i);
                }
            }
        }
        return mn < INT_MAX ? mn : -1; // 最后一个算出的 mn 就是 f[0][0]
    }
};
```

```go [sol-Go]
func minimumVisitedCells(grid [][]int) (mn int) {
	m, n := len(grid), len(grid[0])
	type pair struct{ x, i int }
	colStacks := make([][]pair, n) // 每列的单调栈
	rowSt := []pair{}             // 行单调栈
	for i := m - 1; i >= 0; i-- {
		rowSt = rowSt[:0]
		for j := n - 1; j >= 0; j-- {
			colSt := colStacks[j]
			if i < m-1 || j < n-1 {
				mn = math.MaxInt
			}
			if g := grid[i][j]; g > 0 { // 可以向右/向下跳
				// 在单调栈上二分查找最优转移来源
				k := sort.Search(len(rowSt), func(k int) bool { return rowSt[k].i <= j+g })
				if k < len(rowSt) {
					mn = rowSt[k].x
				}
				k = sort.Search(len(colSt), func(k int) bool { return colSt[k].i <= i+g })
				if k < len(colSt) {
					mn = min(mn, colSt[k].x)
				}
			}
			if mn < math.MaxInt {
				mn++ // 加上 (i,j) 这个格子
				// 插入单调栈
				for len(rowSt) > 0 && mn <= rowSt[len(rowSt)-1].x {
					rowSt = rowSt[:len(rowSt)-1]
				}
				rowSt = append(rowSt, pair{mn, j})
				for len(colSt) > 0 && mn <= colSt[len(colSt)-1].x {
					colSt = colSt[:len(colSt)-1]
				}
				colStacks[j] = append(colSt, pair{mn, i})
			}
		}
	}
	// 最后一个算出的 mn 就是 f[0][0]
	if mn == math.MaxInt {
		return -1
	}
	return
}
```

```js [sol-JavaScript]
var minimumVisitedCells = function(grid) {
    const m = grid.length, n = grid[0].length;
    const colStacks = Array.from({length: n}).map(() => []); // 每列的单调栈
    const rowSt = []; // 行单调栈
    let mn;
    for (let i = m - 1; i >= 0; i--) {
        rowSt.length = 0;
        for (let j = n - 1; j >= 0; j--) {
            const g = grid[i][j];
            const colSt = colStacks[j];
            mn = i < m - 1 || j < n - 1 ? Infinity : 1;
            if (g) { // 可以向右/向下跳
                // 在单调栈上二分查找最优转移来源
                let k = search(rowSt, j + g);
                if (k < rowSt.length) {
                    mn = rowSt[k][0] + 1;
                }
                k = search(colSt, i + g);
                if (k < colSt.length) {
                    mn = Math.min(mn, colSt[k][0] + 1);
                }
            }
            if (mn < Infinity) {
                // 插入单调栈
                while (rowSt.length && mn <= rowSt[rowSt.length - 1][0]) {
                    rowSt.pop();
                }
                rowSt.push([mn, j]);
                while (colSt.length && mn <= colSt[colSt.length - 1][0]) {
                    colSt.pop();
                }
                colSt.push([mn, i]);
            }
        }
    }
    return mn < Infinity ? mn : -1; // 最后一个算出的 mn 就是 f[0][0]
};

// 开区间二分，见 https://www.bilibili.com/video/BV1AP41137w7/
function search(st, target) {
    let left = -1, right = st.length; // 开区间 (left, right)
    while (left + 1 < right) { // 区间不为空
        const mid = Math.floor((left + right) / 2);
        if (st[mid][1] <= target) {
            right = mid; // 范围缩小到 (left, mid)
        } else {
            left = mid; // 范围缩小到 (mid, right)
        }
    }
    return right;
}
```

```rust [sol-Rust]
impl Solution {
    pub fn minimum_visited_cells(grid: Vec<Vec<i32>>) -> i32 {
        let m = grid.len();
        let n = grid[0].len();
        let mut mn = 0;
        let mut col_stacks: Vec<Vec<(i32, usize)>> = vec![vec![]; n]; // 每列的单调栈
        let mut row_st: Vec<(i32, usize)> = Vec::new(); // 行单调栈
        for (i, row) in grid.iter().enumerate().rev() {
            row_st.clear();
            for (j, &g) in row.iter().enumerate().rev() {
                let col_st = &mut col_stacks[j];
                mn = if i < m - 1 || j < n - 1 { i32::MAX } else { 1 };
                if g > 0 { // 可以向右/向下跳
                    let g = g as usize;
                    // 在单调栈上二分查找最优转移来源
                    let k = row_st.partition_point(|(_, idx)| *idx > j + g);
                    if k < row_st.len() {
                        mn = row_st[k].0 + 1;
                    }
                    let k = col_st.partition_point(|(_, idx)| *idx > i + g);
                    if k < col_st.len() {
                        mn = mn.min(col_st[k].0 + 1);
                    }
                }
                if mn < i32::MAX {
                    // 插入单调栈
                    while !row_st.is_empty() && mn <= row_st.last().unwrap().0 {
                        row_st.pop();
                    }
                    row_st.push((mn, j));
                    while !col_st.is_empty() && mn <= col_st.last().unwrap().0 {
                        col_st.pop();
                    }
                    col_st.push((mn, i));
                }
            }
        }
        // 最后一个算出的 mn 就是 f[0][0]
        if mn < i32::MAX { mn } else { -1 }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn(\log m+\log n))$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。

## 方法二：贪心+最小堆

这个思路类似 [Dijkstra 算法](https://leetcode.cn/problems/network-delay-time/solution/liang-chong-dijkstra-xie-fa-fu-ti-dan-py-ooe8/)。

从起点 $(0,0)$ 开始，小到大枚举 $i$ 和 $j$。

假设枚举到 $(i,j)$，设从 $(0,0)$ 到 $(i,j)$ 经过了 $f$ 个格子，设 $g=\textit{grid}[i][j]$。

- 从 $(i,j)$ 出发向右，我们最远可以到达第 $g+j$ 列，把数对 $(f,g+j)$ 保存到一个数据结构 $\textit{rowH}$ 中。
- 从 $(i,j)$ 出发向下，我们最远可以到达第 $g+i$ 行，把数对 $(f,g+i)$ 保存到另一个数据结构 $\textit{colH}$ 中。

怎么算出 $f$？答案就在 $\textit{rowH}$ 和 $\textit{colH}$ 中。

- 从 $\textit{rowH}$ 中找到一个 $f$ 值最小的，且可以到达第 $j$ 列的数对。
- 从 $\textit{colH}$ 中找到一个 $f$ 值最小的，且可以到达第 $i$ 行的数对。

选择哪种数据结构实现？这样的数据结构需要支持：

- 添加元素。
- 查找最小元素。
- 如果最小元素不满足要求（无法到达行/列），将其从数据结构中移除。

**最小堆**完美符合上述要求。

具体来说，对于每一行和每一列，都有一个对应的最小堆。下面的 $\textit{rowH}$ 维护的是第 $i$ 行的数对，$\textit{colH}$ 维护的是第 $j$ 列的数对。

- 对于行最小堆 $\textit{rowH}$，如果堆顶列号（也就是之前保存的 $g+j$）小于当前列号 $j$，说明无法到达第 $j$ 列（以及更右的列），弹出堆顶，重复该过程直到堆为空或者堆顶列号大于等于 $j$。如果堆不为空，取堆顶的 $f$ 值加一，作为 $(i,j)$ 位置的 $f$。
- 对于列最小堆 $\textit{colH}$，如果堆顶行号（也就是之前保存的 $g+i$）小于当前行号 $i$，说明无法到达第 $i$ 行（以及更下的行），弹出堆顶，重复该过程直到堆为空或者堆顶行号大于等于 $i$。如果堆不为空，取堆顶的 $f$ 值加一，作为 $(i,j)$ 位置的 $f$。
- 这两种情况取最小值。
- 特别地，如果 $i=0$ 且 $j=0$，我们位于起点，$f=1$。
- 特别地，如果无法到达 $(i,j)$，则 $f=\infty$。

```py [sol-Python3]
class Solution:
    def minimumVisitedCells(self, grid: List[List[int]]) -> int:
        col_heaps = [[] for _ in grid[0]]  # 每一列的最小堆
        for i, row in enumerate(grid):
            row_h = []  # 第 i 行的最小堆
            for j, (g, col_h) in enumerate(zip(row, col_heaps)):
                while row_h and row_h[0][1] < j:  # 无法到达第 j 列
                    heappop(row_h)  # 弹出无用数据
                while col_h and col_h[0][1] < i:  # 无法到达第 i 行
                    heappop(col_h)  # 弹出无用数据
                f = inf if i or j else 1  # 起点算 1 个格子
                if row_h: f = row_h[0][0] + 1  # 从左边跳过来
                if col_h: f = min(f, col_h[0][0] + 1)  # 从上边跳过来
                if g and f < inf:
                    heappush(row_h, (f, g + j))  # 经过的格子数，向右最远能到达的列号
                    heappush(col_h, (f, g + i))  # 经过的格子数，向下最远能到达的行号
        return f if f < inf else -1  # 此时的 f 是在 (m-1, n-1) 处算出来的
```

```java [sol-Java]
class Solution {
    public int minimumVisitedCells(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        int f = 0;
        PriorityQueue<int[]>[] colHeaps = new PriorityQueue[n]; // 每一列的最小堆
        Arrays.setAll(colHeaps, i -> new PriorityQueue<int[]>((a, b) -> a[0] - b[0]));
        PriorityQueue<int[]> rowH = new PriorityQueue<>((a, b) -> a[0] - b[0]); // 行最小堆
        for (int i = 0; i < m; i++) {
            rowH.clear();
            for (int j = 0; j < n; j++) {
                while (!rowH.isEmpty() && rowH.peek()[1] < j) { // 无法到达第 j 列
                    rowH.poll(); // 弹出无用数据
                }
                PriorityQueue<int[]> colH = colHeaps[j];
                while (!colH.isEmpty() && colH.peek()[1] < i) { // 无法到达第 i 行
                    colH.poll(); // 弹出无用数据
                }

                f = i > 0 || j > 0 ? Integer.MAX_VALUE : 1; // 起点算 1 个格子
                if (!rowH.isEmpty()) {
                    f = rowH.peek()[0] + 1; // 从左边跳过来
                }
                if (!colH.isEmpty()) {
                    f = Math.min(f, colH.peek()[0] + 1); // 从上边跳过来
                }

                int g = grid[i][j];
                if (g > 0 && f < Integer.MAX_VALUE) {
                    rowH.offer(new int[]{f, g + j}); // 经过的格子数，向右最远能到达的列号
                    colH.offer(new int[]{f, g + i}); // 经过的格子数，向下最远能到达的行号
                }
            }
        }
        return f < Integer.MAX_VALUE ? f : -1; // 此时的 f 是在 (m-1, n-1) 处算出来的
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumVisitedCells(vector<vector<int>> &grid) {
        int m = grid.size(), n = grid[0].size(), f;
        using min_heap_t = priority_queue<pair<int, int>, vector<pair<int, int>>, greater<>>;
        vector<min_heap_t> col_heaps(n); // 每一列的最小堆
        for (int i = 0; i < m; i++) {
            min_heap_t row_h; // 第 i 行的最小堆
            for (int j = 0; j < n; j++) {
                while (!row_h.empty() && row_h.top().second < j) { // 无法到达第 j 列
                    row_h.pop(); // 弹出无用数据
                }
                auto &col_h = col_heaps[j];
                while (!col_h.empty() && col_h.top().second < i) { // 无法到达第 i 行
                    col_h.pop(); // 弹出无用数据
                }

                f = i || j ? INT_MAX : 1; // 起点算 1 个格子
                if (!row_h.empty()) f = row_h.top().first + 1; // 从左边跳过来
                if (!col_h.empty()) f = min(f, col_h.top().first + 1); // 从上边跳过来

                int g = grid[i][j];
                if (g && f < INT_MAX) {
                    row_h.emplace(f, g + j); // 经过的格子数，向右最远能到达的列号
                    col_h.emplace(f, g + i); // 经过的格子数，向下最远能到达的行号
                }
            }
        }
        return f < INT_MAX ? f : -1; // 此时的 f 是在 (m-1, n-1) 处算出来的
    }
};
```

```go [sol-Go]
func minimumVisitedCells(grid [][]int) int {
	colHeaps := make([]hp, len(grid[0])) // 每一列的最小堆
	rowH := hp{} // 行最小堆
	f := 1 // 起点算 1 个格子
	for i, row := range grid {
		rowH = rowH[:0]
		for j, g := range row {
			for len(rowH) > 0 && rowH[0].idx < j { // 无法到达第 j 列
				heap.Pop(&rowH) // 弹出无用数据
			}
			colH := colHeaps[j]
			for len(colH) > 0 && colH[0].idx < i { // 无法到达第 i 行
				heap.Pop(&colH) // 弹出无用数据
			}
			if i > 0 || j > 0 {
				f = math.MaxInt
			}
			if len(rowH) > 0 {
				f = rowH[0].f + 1 // 从左边跳过来
			}
			if len(colH) > 0 {
				f = min(f, colH[0].f+1) // 从上边跳过来
			}
			if g > 0 && f < math.MaxInt {
				heap.Push(&rowH, pair{f, g + j}) // 经过的格子数，向右最远能到达的列号
				heap.Push(&colH, pair{f, g + i}) // 经过的格子数，向下最远能到达的行号
			}
			colHeaps[j] = colH // 注意上面的入堆出堆操作只改了 colH，没有改 colHeaps[j]
		}
	}
	// 此时的 f 是在 (m-1, n-1) 处算出来的
	if f < math.MaxInt {
		return f
	}
	return -1
}

type pair struct{ f, idx int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].f < h[j].f }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```

```js [sol-JavaScript]
var minimumVisitedCells = function (grid) {
    const m = grid.length, n = grid[0].length;
    const colHeaps = Array.from({length: n}, () => new MinPriorityQueue({priority: (p) => p[0]})); // 每一列的最小堆
    const rowH = new MinPriorityQueue({priority: (p) => p[0]}); // 行最小堆
    let f;
    for (let i = 0; i < m; i++) {
        rowH.clear();
        for (let j = 0; j < n; j++) {
            while (!rowH.isEmpty() && rowH.front().element[1] < j) { // 无法到达第 j 列
                rowH.dequeue(); // 弹出无用数据
            }
            const colH = colHeaps[j];
            while (!colH.isEmpty() && colH.front().element[1] < i) { // 无法到达第 i 行
                colH.dequeue(); // 弹出无用数据
            }

            f = i || j ? Infinity : 1; // 起点算 1 个格子
            if (!rowH.isEmpty()) f = rowH.front().element[0] + 1; // 从左边跳过来
            if (!colH.isEmpty()) f = Math.min(f, colH.front().element[0] + 1); // 从上边跳过来

            const g = grid[i][j]
            if (g && f < Infinity) {
                rowH.enqueue([f, g + j]); // 经过的格子数，向右最远能到达的列号
                colH.enqueue([f, g + i]); // 经过的格子数，向下最远能到达的行号
            }
        }
    }
    return f < Infinity ? f : -1; // 此时的 f 是在 (m-1, n-1) 处算出来的
};
```

```rust [sol-Rust]
use std::collections::BinaryHeap;

impl Solution {
    pub fn minimum_visited_cells(grid: Vec<Vec<i32>>) -> i32 {
        let mut col_heaps: Vec<BinaryHeap<(i32, usize)>> = vec![BinaryHeap::new(); grid[0].len()]; // 每一列的最小堆
        let mut row_h: BinaryHeap<(i32, usize)> = BinaryHeap::new(); // 行最小堆
        let mut f = 0;
        for (i, row) in grid.iter().enumerate() {
            row_h.clear();
            for (j, &g) in row.iter().enumerate() {
                while !row_h.is_empty() && row_h.peek().unwrap().1 < j {
                    row_h.pop();
                }
                let col_h = &mut col_heaps[j];
                while !col_h.is_empty() && col_h.peek().unwrap().1 < i {
                    col_h.pop();
                }

                f = if i > 0 || j > 0 { i32::MAX } else { 1 }; // 起点算 1 个格子
                if let Some(&(fr, _)) = row_h.peek() {
                    f = -fr + 1; // 从左边跳过来
                }
                if let Some(&(fc, _)) = col_h.peek() {
                    f = f.min(-fc + 1); // 从上边跳过来
                }

                if g > 0 && f < i32::MAX {
                    // f 加负号变成最小堆
                    row_h.push((-f, g as usize + j)); // 经过的格子数，向右最远能到达的列号
                    col_h.push((-f, g as usize + i)); // 经过的格子数，向下最远能到达的行号
                }
            }
        }
        // 此时的 f 是在 (m-1, n-1) 处算出来的
        if f < i32::MAX { f } else { -1 }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn(\log m+\log n))$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
