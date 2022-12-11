[视频讲解](https://www.bilibili.com/video/BV1QK41167cr/) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

# 方法一：离线询问 + 并查集

把矩阵的元素值从小到大排序，询问也从小到大排序。

用双指针遍历矩阵元素值和询问，如果矩阵元素值小于询问值，就把该格子和周围四个格子中的小于询问值的格子相连。

用并查集可以实现相连的过程，同时维护每个连通块的大小。

答案就是左上角的连通块的大小（前提是左上角小于询问值）。

[视频讲解](https://www.bilibili.com/video/BV1QK41167cr/) 中用的是对边权排序的做法，两个格子之间的边权用格子值的最大值表示。

#### 非常相似的题目

- [1697. 检查边长度限制的路径是否存在](https://leetcode.cn/problems/checking-existence-of-edge-length-limited-paths/)

```py [sol1-Python3]
class Solution:
    def maxPoints(self, grid: List[List[int]], queries: List[int]) -> List[int]:
        m, n = len(grid), len(grid[0])
        mn = m * n

        # 并查集模板
        fa = list(range(mn))
        size = [1] * mn
        def find(x: int) -> int:
            if fa[x] != x:
                fa[x] = find(fa[x])
            return fa[x]
        def merge(from_: int, to: int) -> None:
            from_ = find(from_)
            to = find(to)
            if from_ != to:
                fa[from_] = to
                size[to] += size[from_]

        # 矩阵元素从小到大排序，方便离线
        a = sorted((x, i, j) for i, row in enumerate(grid) for j, x in enumerate(row))

        ans, j = [0] * len(queries), 0
        # 查询的下标按照查询值从小到大排序，方便离线
        for i, q in sorted(enumerate(queries), key=lambda p: p[1]):
            while j < mn and a[j][0] < q:
                _, x, y = a[j]
                # 枚举周围四个格子，值小于 q 才可以合并
                for x2, y2 in (x + 1, y), (x - 1, y), (x, y + 1), (x, y - 1):
                    if 0 <= x2 < m and 0 <= y2 < n and grid[x2][y2] < q:
                        merge(x * n + y, x2 * n + y2)  # 把坐标压缩成一维的编号
                j += 1
            if grid[0][0] < q:
                ans[i] = size[find(0)]  # 左上角的连通块的大小
        return ans
```

```java [sol1-Java]
class Solution {
    private static final int[][] dirs = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
    private int[] fa, size;

    public int[] maxPoints(int[][] grid, int[] queries) {
        int m = grid.length, n = grid[0].length, mn = m * n;

        // 并查集
        fa = new int[mn];
        for (var i = 0; i < mn; i++) fa[i] = i;
        size = new int[mn];
        Arrays.fill(size, 1);

        // 矩阵元素从小到大排序，方便离线
        var a = new int[mn][3];
        for (var i = 0; i < m; ++i)
            for (var j = 0; j < n; ++j)
                a[i * n + j] = new int[]{grid[i][j], i, j};
        Arrays.sort(a, (p, q) -> p[0] - q[0]);

        // 查询的下标按照查询值从小到大排序，方便离线
        var k = queries.length;
        var id = IntStream.range(0, k).boxed().toArray(Integer[]::new);
        Arrays.sort(id, (i, j) -> queries[i] - queries[j]);

        var ans = new int[k];
        var j = 0;
        for (var i : id) {
            var q = queries[i];
            for (; j < mn && a[j][0] < q; ++j) {
                int x = a[j][1], y = a[j][2];
                // 枚举周围四个格子，值小于 q 才可以合并
                for (var d : dirs) {
                    int x2 = x + d[0], y2 = y + d[1];
                    if (0 <= x2 && x2 < m && 0 <= y2 && y2 < n && grid[x2][y2] < q)
                        merge(x * n + y, x2 * n + y2); // 把坐标压缩成一维的编号
                }
            }
            if (grid[0][0] < q)
                ans[i] = size[find(0)]; // 左上角的连通块的大小
        }
        return ans;
    }

    // 并查集模板
    private int find(int x) {
        if (fa[x] != x) fa[x] = find(fa[x]);
        return fa[x];
    }

    private void merge(int from, int to) {
        from = find(from);
        to = find(to);
        if (from != to) {
            fa[from] = to;
            size[to] += size[from];
        }
    }
}
```

```cpp [sol1-C++]
class Solution {
    const int dirs[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
public:
    vector<int> maxPoints(vector<vector<int>> &grid, vector<int> &queries) {
        int m = grid.size(), n = grid[0].size(), mn = m * n;

        // 并查集模板
        int fa[mn], size[mn];
        iota(fa, fa + mn, 0);
        fill(size, size + mn, 1);
        function<int(int)> find = [&](int x) -> int { return fa[x] == x ? x : fa[x] = find(fa[x]); };
        auto merge = [&](int from, int to) {
            from = find(from);
            to = find(to);
            if (from != to) {
                fa[from] = to;
                size[to] += size[from];
            }
        };

        // 矩阵元素从小到大排序，方便离线
        array<int, 3> a[mn];
        for (int i = 0; i < m; ++i)
            for (int j = 0; j < n; ++j)
                a[i * n + j] = {grid[i][j], i, j};
        sort(a, a + mn);

        // 查询的下标按照查询值从小到大排序，方便离线
        int k = queries.size(), id[k];
        iota(id, id + k, 0);
        sort(id, id + k, [&](int i, int j) { return queries[i] < queries[j]; });

        vector<int> ans(k);
        int j = 0;
        for (int i : id) {
            int q = queries[i];
            for (; j < mn && a[j][0] < q; ++j) {
                int x = a[j][1], y = a[j][2];
                // 枚举周围四个格子，值小于 q 才可以合并
                for (auto &d : dirs) {
                    int x2 = x + d[0], y2 = y + d[1];
                    if (0 <= x2 && x2 < m && 0 <= y2 && y2 < n && grid[x2][y2] < q)
                        merge(x * n + y, x2 * n + y2); // 把坐标压缩成一维的编号
                }
            }
            if (grid[0][0] < q)
                ans[i] = size[find(0)]; // 左上角的连通块的大小
        }
        return ans;
    }
};
```

```go [sol1-Go]
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maxPoints(grid [][]int, queries []int) []int {
	m, n := len(grid), len(grid[0])
	mn := m * n

	// 并查集模板
	fa := make([]int, mn)
	size := make([]int, mn)
	for i := range fa {
		fa[i] = i
		size[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) {
		from = find(from)
		to = find(to)
		if from != to {
			fa[from] = to
			size[to] += size[from]
		}
	}

	// 矩阵元素从小到大排序，方便离线
	type tuple struct{ x, i, j int }
	a := make([]tuple, 0, mn)
	for i, row := range grid {
		for j, x := range row {
			a = append(a, tuple{x, i, j})
		}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })

	// 查询的下标按照查询值从小到大排序，方便离线
	id := make([]int, len(queries))
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return queries[id[i]] < queries[id[j]] })

	ans := make([]int, len(queries))
	j := 0
	for _, i := range id {
		q := queries[i]
		for ; j < mn && a[j].x < q; j++ {
			x, y := a[j].i, a[j].j
			// 枚举周围四个格子，值小于 q 才可以合并
			for _, d := range dirs {
				x2, y2 := x+d.x, y+d.y
				if 0 <= x2 && x2 < m && 0 <= y2 && y2 < n && grid[x2][y2] < q {
					merge(x*n+y, x2*n+y2) // 把坐标压缩成一维的编号
				}
			}
		}
		if grid[0][0] < q {
			ans[i] = size[find(0)] // 左上角的连通块的大小
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$O(mn\log mn + k\log k)$，其中 $m$ 和 $n$ 分别为矩阵 $\textit{grid}$ 的行数和列数，$k$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$O(mn)$。

# 方法二：离线询问 + 最小堆

仍然是离线询问，还可以从左上角出发向外搜索，用最小堆，初始把左上角的元素值及其坐标入堆。对每个询问，不断循环，如果堆顶元素值小于询问值，则弹出堆顶，继续搜索。

循环结束时，出堆的元素个数就是答案。

代码实现时，可以用 $\textit{grid}$ 作为 $\textit{vis}$ 数组，从而优化空间复杂度。

```py [sol2-Python3]
class Solution:
    def maxPoints(self, grid: List[List[int]], queries: List[int]) -> List[int]:
        m, n = len(grid), len(grid[0])
        ans = [0] * len(queries)
        h = [(grid[0][0], 0, 0)]
        grid[0][0] = 0  # 充当 vis 数组的作用
        cnt = 0
        # 查询的下标按照查询值从小到大排序，方便离线
        for qi, q in sorted(enumerate(queries), key=lambda p: p[1]):
            while h and h[0][0] < q:
                cnt += 1
                _, i, j = heappop(h)
                for x, y in (i + 1, j), (i - 1, j), (i, j + 1), (i, j - 1):  # 枚举周围四个格子
                    if 0 <= x < m and 0 <= y < n and grid[x][y]:
                        heappush(h, (grid[x][y], x, y))
                        grid[x][y] = 0  # 充当 vis 数组的作用
            ans[qi] = cnt
        return ans
```

```java [sol2-Java]
class Solution {
    private static final int[][] dirs = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

    public int[] maxPoints(int[][] grid, int[] queries) {
        // 查询的下标按照查询值从小到大排序，方便离线
        var k = queries.length;
        var id = IntStream.range(0, k).boxed().toArray(Integer[]::new);
        Arrays.sort(id, (i, j) -> queries[i] - queries[j]);

        var ans = new int[k];
        var pq = new PriorityQueue<int[]>((a, b) -> a[0] - b[0]);
        pq.add(new int[]{grid[0][0], 0, 0});
        grid[0][0] = 0; // 充当 vis 数组的作用
        int m = grid.length, n = grid[0].length, cnt = 0;
        for (var i : id) {
            var q = queries[i];
            while (!pq.isEmpty() && pq.peek()[0] < q) {
                ++cnt;
                var p = pq.poll();
                for (var d : dirs) { // 枚举周围四个格子
                    int x = p[1] + d[0], y = p[2] + d[1];
                    if (0 <= x && x < m && 0 <= y && y < n && grid[x][y] > 0) {
                        pq.add(new int[]{grid[x][y], x, y});
                        grid[x][y] = 0; // 充当 vis 数组的作用
                    }
                }
            }
            ans[i] = cnt;
        }
        return ans;
    }
}
```

```cpp [sol2-C++]
class Solution {
    const int dirs[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
public:
    vector<int> maxPoints(vector<vector<int>> &grid, vector<int> &queries) {
        // 查询的下标按照查询值从小到大排序，方便离线
        int k = queries.size(), id[k];
        iota(id, id + k, 0);
        sort(id, id + k, [&](int i, int j) { return queries[i] < queries[j]; });

        vector<int> ans(k);
        priority_queue<tuple<int, int, int>, vector<tuple<int, int, int>>, greater<>> pq; // 最小堆
        pq.emplace(grid[0][0], 0, 0);
        grid[0][0] = 0; // 充当 vis 数组的作用
        int m = grid.size(), n = grid[0].size(), cnt = 0;
        for (int qi : id) {
            int q = queries[qi];
            while (!pq.empty() && get<0>(pq.top()) < q) {
                ++cnt;
                auto[_, i, j] = pq.top();
                pq.pop();
                for (auto &d : dirs) { // 枚举周围四个格子
                    int x = i + d[0], y = j + d[1];
                    if (0 <= x && x < m && 0 <= y && y < n && grid[x][y]) {
                        pq.emplace(grid[x][y], x, y);
                        grid[x][y] = 0; // 充当 vis 数组的作用
                    }
                }
            }
            ans[qi] = cnt;
        }
        return ans;
    }
};
```

```go [sol2-Go]
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maxPoints(grid [][]int, queries []int) []int {
	m, n := len(grid), len(grid[0])

	// 查询的下标按照查询值从小到大排序，方便离线
	id := make([]int, len(queries))
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return queries[id[i]] < queries[id[j]] })

	ans := make([]int, len(queries))
	h := hp{{grid[0][0], 0, 0}}
	grid[0][0] = 0 // 充当 vis 数组的作用
	cnt := 0
	for _, i := range id {
		q := queries[i]
		for len(h) > 0 && h[0].val < q {
			cnt++
			p := heap.Pop(&h).(tuple)
			for _, d := range dirs { // 枚举周围四个格子
				x, y := p.i+d.x, p.j+d.y
				if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] > 0 {
					heap.Push(&h, tuple{grid[x][y], x, y})
					grid[x][y] = 0 // 充当 vis 数组的作用
				}
			}
		}
		ans[i] = cnt
	}
	return ans
}

type tuple struct{ val, i, j int }
type hp []tuple
func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：$O(mn\log(m+n) + k\log k)$，其中 $m$ 和 $n$ 分别为矩阵 $\textit{grid}$ 的行数和列数，$k$ 为 $\textit{queries}$ 的长度。类似网格图上的 BFS，堆里面至多有 $O(m+n)$ 个元素。
- 空间复杂度：$O(m+n+k)$。
