## 思路一

设矩形的左下角为 $(x_1,y_1)$，右上角为 $(x_2,y_2)$。

**核心思路**：计算横坐标 $\le x_1$ 且纵坐标在 $[y_1,y_2]$ 中的点的个数，记作 $\textit{pre}$。然后计算横坐标 $\le x_2$ 且纵坐标在 $[y_1,y_2]$ 中的点的个数，记作 $\textit{cur}$。如果 $\textit{pre}+2 = \textit{cur}$，说明新增的 $2$ 个点是矩形的右下角和右上角，进而说明矩形区域（含边界）只有 $4$ 个点，是合法矩形，用矩形面积 $(x_2-x_1)\cdot(y_2-y_1)$ 更新答案的最大值。

把所有点按照横坐标从小到大排序，横坐标相同的点，按照纵坐标从小到大排序。得到排序后的数组 $\textit{points}$。

遍历 $\textit{points}$，把纵坐标（离散化后）加到**树状数组**中。注意树状数组只保存纵坐标。

1. 对于每个矩形，我们首先会遇到矩形的左下角和左上角（这两个点在 $\textit{points}$ 中一定相邻）。在树状数组中查询纵坐标在 $[y_1,y_2]$ 中的点的个数 $C_1$。用哈希表（数组）记录，key 是 $y_2$，value 是左下角的坐标以及 $C_1$。
2. 然后继续遍历，遇到矩形的右下角和右上角（这两个点在 $\textit{points}$ 中也一定相邻）。在树状数组中查询纵坐标在 $[y_1,y_2]$ 中的点的个数 $C_2$。如果哈希表（数组）中记录的左下角纵坐标等于 $y_1$，且 $C_1 + 2 = C_2$，那么这是一个合法的矩形。

> 注意 $\textit{points}$ 中的两个相邻的点，既可以是某个矩形的右下角和右上角，又可以是另一个矩形的左下角和左上角。

```py [sol-Python3]
# 树状数组模板
class Fenwick:
    def __init__(self, n: int):
        self.tree = [0] * (n + 1)

    def add(self, i: int) -> None:
        while i < len(self.tree):
            self.tree[i] += 1
            i += i & -i

    # [1,i] 中的元素和
    def pre(self, i: int) -> int:
        res = 0
        while i > 0:
            res += self.tree[i]
            i &= i - 1
        return res

    # [l,r] 中的元素和
    def query(self, l: int, r: int) -> int:
        return self.pre(r) - self.pre(l - 1)

class Solution:
    def maxRectangleArea(self, xCoord: List[int], yCoord: List[int]) -> int:
        points = sorted(zip(xCoord, yCoord))
        ys = sorted(set(yCoord))  # 离散化用

        ans = -1
        tree = Fenwick(len(ys))
        tree.add(bisect_left(ys, points[0][1]) + 1)  # 离散化
        pre = {}
        for (x1, y1), (x2, y2) in pairwise(points):
            y = bisect_left(ys, y2) + 1  # 离散化
            tree.add(y)
            if x1 != x2:  # 两点不在同一列
                continue
            cur = tree.query(bisect_left(ys, y1) + 1, y)
            if y2 in pre and pre[y2][1] == y1 and pre[y2][2] + 2 == cur:
                ans = max(ans, (x2 - pre[y2][0]) * (y2 - y1))
            pre[y2] = (x1, y1, cur)
        return ans
```

```java [sol-Java]
class Fenwick {
    private final int[] tree;

    Fenwick(int n) {
        tree = new int[n];
    }

    void add(int i) {
        while (i < tree.length) {
            tree[i]++;
            i += i & -i;
        }
    }

    // [1,i] 中的元素和
    int pre(int i) {
        int res = 0;
        while (i > 0) {
            res += tree[i];
            i -= i & -i;
        }
        return res;
    }

    // [l,r] 中的元素和
    int query(int l, int r) {
        return pre(r) - pre(l - 1);
    }
}

class Solution {
    public long maxRectangleArea(int[] xCoord, int[] ys) {
        int n = xCoord.length;
        int[][] points = new int[n][2];
        for (int i = 0; i < n; i++) {
            points[i][0] = xCoord[i];
            points[i][1] = ys[i];
        }
        Arrays.sort(points, (a, b) -> a[0] != b[0] ? a[0] - b[0] : a[1] - b[1]);
        Arrays.sort(ys); // 离散化用

        long ans = -1;
        Fenwick tree = new Fenwick(n + 1);
        tree.add(Arrays.binarySearch(ys, points[0][1]) + 1); // 离散化
        int[][] pre = new int[n][3];
        for (int i = 1; i < n; i++) {
            int x1 = points[i - 1][0];
            int y1 = points[i - 1][1];
            int x2 = points[i][0];
            int y2 = points[i][1];
            int y = Arrays.binarySearch(ys, y2); // 离散化
            tree.add(y + 1);
            if (x1 != x2) { // 两点不在同一列
                continue;
            }
            int cur = tree.query(Arrays.binarySearch(ys, y1) + 1, y + 1);
            int[] p = pre[y];
            if (p[2] > 0 && p[2] + 2 == cur && p[1] == y1) {
                ans = Math.max(ans, (long) (x2 - p[0]) * (y2 - y1));
            }
            p[0] = x1;
            p[1] = y1;
            p[2] = cur;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
// 树状数组模板
class Fenwick {
    vector<int> tree;

public:
    Fenwick(int n) : tree(n, 0) {}

    void add(int i) {
        while (i < tree.size()) {
            tree[i]++;
            i += i & -i;
        }
    }

    // [1,i] 中的元素和
    int pre(int i) {
        int res = 0;
        while (i > 0) {
            res += tree[i];
            i -= i & -i;
        }
        return res;
    }

    // [l,r] 中的元素和
    int query(int l, int r) {
        return pre(r) - pre(l - 1);
    }
};

class Solution {
public:
    long long maxRectangleArea(vector<int>& xCoord, vector<int>& ys) {
        vector<pair<int, int>> points;
        for (int i = 0; i < xCoord.size(); i++) {
            points.emplace_back(xCoord[i], ys[i]);
        }
        ranges::sort(points);

        // 离散化用
        ranges::sort(ys);
        ys.erase(unique(ys.begin(), ys.end()), ys.end()); // 去重

        long long ans = -1;
        Fenwick tree(ys.size() + 1);
        tree.add(ranges::lower_bound(ys, points[0].second) - ys.begin() + 1); // 离散化
        vector<tuple<int, int, int>> pre(ys.size(), {-1, -1, -1});
        for (int i = 1; i < points.size(); i++) {
            auto& [x1, y1] = points[i - 1];
            auto& [x2, y2] = points[i];
            int y = ranges::lower_bound(ys, y2) - ys.begin(); // 离散化
            tree.add(y + 1);
            if (x1 != x2) { // 两点不在同一列
                continue;
            }
            int cur = tree.query(ranges::lower_bound(ys, y1) - ys.begin() + 1, y + 1);
            auto& [pre_x, pre_y, p] = pre[y];
            if (pre_y == y1 && p + 2 == cur) {
                ans = max(ans, (long long) (x2 - pre_x) * (y2 - y1));
            }
            pre[y] = {x1, y1, cur};
        }
        return ans;
    }
};
```

```go [sol-Go]
// 树状数组模板
type fenwick []int

func (f fenwick) add(i int) {
	for ; i < len(f); i += i & -i {
		f[i]++
	}
}

// [1,i] 中的元素和
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

// [l,r] 中的元素和
func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

func maxRectangleArea(xCoord, ys []int) int64 {
	type pair struct{ x, y int }
	points := make([]pair, len(xCoord))
	for i := range xCoord {
		points[i] = pair{xCoord[i], ys[i]}
	}
	slices.SortFunc(points, func(a, b pair) int { return cmp.Or(a.x-b.x, a.y-b.y) })

	// 离散化用
	slices.Sort(ys)
	ys = slices.Compact(ys)

	ans := -1
	tree := make(fenwick, len(ys)+1)
	tree.add(sort.SearchInts(ys, points[0].y) + 1) // 离散化
	type tuple struct{ x, y, c int }
	pre := make([]tuple, len(ys))
	for i := 1; i < len(points); i++ {
		x1, y1 := points[i-1].x, points[i-1].y
		x2, y2 := points[i].x, points[i].y
		y := sort.SearchInts(ys, y2) // 离散化
		tree.add(y + 1)
		if x1 != x2 { // 两点不在同一列
			continue
		}
		cur := tree.query(sort.SearchInts(ys, y1)+1, y+1)
		p := pre[y]
		if p.c > 0 && p.c+2 == cur && p.y == y1 {
			ans = max(ans, (x2-p.x)*(y2-y1))
		}
		pre[y] = tuple{x1, y1, cur}
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{xCoord}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 思路二

我们来分别解决两个更一般的问题：

1. 找到所有矩形，这些矩形的边界（除去顶点）不含其他点。
2. 计算每个矩形区域中的点的个数。本题要求点的个数恰好等于 $4$（四个顶点）。

由于矩形边界不能包含其他点，那么枚举矩形的右上角 $(x_2,y_2)$，它正左边的点就是左上角，正下方的点就是右下角。

左上角的横坐标 $x_1$ 和右下角的纵坐标 $y_1$ 就组成了矩形的左下角 $(x_1,y_1)$。

> 需要判断：矩形右下角的左边的点的横坐标必须是 $x_1$，矩形左上角的下边的点的纵坐标必须是 $y_1$。

把点按照横坐标分组，即可快速知道每个点的正下方的点；按照纵坐标分组，即可快速知道每个点的正左边的点。

现在，问题变成一个（静态的）**二维数点问题**：

- 给你一些询问，对于每个询问，你需要计算横坐标在 $[x_1,x_2]$ 中且纵坐标在 $[y_1,y_2]$ 中的点的个数。如果恰好有 $4$ 个点，用矩形面积 $(x_2-x_1)\cdot(y_2-y_1)$ 更新答案的最大值。

这可以**离线**解决，核心想法是：

- 计算横坐标 $\le x_2$ 且纵坐标在 $[y_1,y_2]$ 中的点的个数。
- 计算横坐标 $\le x_1-1$ 且纵坐标在 $[y_1,y_2]$ 中的点的个数。
- 二者相减，即为矩形区域中的点的个数。

具体来说：

1. 把询问按照 $x_1-1$ 和 $x_2$ 的值分组。
2. 对于 $x_1-1$，记录如下信息：
    - 询问的编号。
    - 计算点的个数时，这部分要减去，系数是 $-1$。
    - 纵坐标 $y_1$ 和 $y_2$。
3. 对于 $x_2$，记录如下信息：
    - 询问的编号。
    - 计算点的个数时，这部分要加上，系数是 $1$。
    - 纵坐标 $y_1$ 和 $y_2$。

然后用 [树状数组](https://leetcode.cn/problems/range-sum-query-mutable/solution/dai-ni-fa-ming-shu-zhuang-shu-zu-fu-shu-lyfll/) 按照 $x$ 从小到大的顺序记录、计算点的个数（区间中的 $y$ 的个数）。

由于数据范围很大，需要离散化。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1YeqHYSEhK/?t=24m52s)，欢迎点赞关注~

```py [sol-Python3]
# 树状数组模板
class Fenwick:
    def __init__(self, n: int):
        self.tree = [0] * (n + 1)

    def add(self, i: int) -> None:
        while i < len(self.tree):
            self.tree[i] += 1
            i += i & -i

    # [1,i] 中的元素和
    def pre(self, i: int) -> int:
        res = 0
        while i > 0:
            res += self.tree[i]
            i &= i - 1
        return res

    # [l,r] 中的元素和
    def query(self, l: int, r: int) -> int:
        return self.pre(r) - self.pre(l - 1)

class Solution:
    def maxRectangleArea(self, xCoord: List[int], yCoord: List[int]) -> int:
        x_map = defaultdict(list)  # 同一列的所有点的纵坐标
        y_map = defaultdict(list)  # 同一行的所有点的横坐标
        for x, y in zip(xCoord, yCoord):
            x_map[x].append(y)
            y_map[y].append(x)

        # 预处理每个点的正下方的点
        below = {}
        for x, ys in x_map.items():
            ys.sort()
            for y1, y2 in pairwise(ys):
                below[(x, y2)] = y1

        # 预处理每个点的正左边的点
        left = {}
        for y, xs in y_map.items():
            xs.sort()
            for x1, x2 in pairwise(xs):
                left[(x2, y)] = x1

        # 离散化用
        xs = sorted(x_map)
        ys = sorted(y_map)

        # 收集询问：矩形区域（包括边界）的点的个数
        queries = []
        # 枚举 (x2,y2) 作为矩形的右上角
        for x2, list_y in x_map.items():
            for y1, y2 in pairwise(list_y):
                # 计算矩形左下角 (x1,y1)
                x1 = left.get((x2, y2), None)
                # 矩形右下角的左边的点的横坐标必须是 x1
                # 矩形左上角的下边的点的纵坐标必须是 y1
                if x1 is not None and left.get((x2, y1), None) == x1 and below.get((x1, y2), None) == y1:
                    queries.append((
                        bisect_left(xs, x1),  # 离散化
                        bisect_left(xs, x2),
                        bisect_left(ys, y1),
                        bisect_left(ys, y2),
                        (x2 - x1) * (y2 - y1),
                    ))

        # 离线询问
        grouped_queries = [[] for _ in range(len(xs))]
        for i, (x1, x2, y1, y2, _) in enumerate(queries):
            if x1 > 0:
                grouped_queries[x1 - 1].append((i, -1, y1, y2))
            grouped_queries[x2].append((i, 1, y1, y2))

        # 回答询问
        res = [0] * len(queries)
        tree = Fenwick(len(ys))
        for x, qs in zip(xs, grouped_queries):
            # 把横坐标为 x 的所有点都加到树状数组中
            for y in x_map[x]:
                tree.add(bisect_left(ys, y) + 1)  # 离散化
            for qid, sign, y1, y2 in qs:
                # 查询横坐标 <= x（已满足）且纵坐标在 [y1,y2] 中的点的个数
                res[qid] += sign * tree.query(y1 + 1, y2 + 1)

        ans = -1
        for cnt, q in zip(res, queries):
            if cnt == 4:
                ans = max(ans, q[4])  # q[4] 保存着矩形面积
        return ans
```

```java [sol-Java]
class Fenwick {
    private final int[] tree;

    Fenwick(int n) {
        tree = new int[n];
    }

    void add(int i) {
        while (i < tree.length) {
            tree[i]++;
            i += i & -i;
        }
    }

    // [1,i] 中的元素和
    int pre(int i) {
        int res = 0;
        while (i > 0) {
            res += tree[i];
            i -= i & -i;
        }
        return res;
    }

    // [l,r] 中的元素和
    int query(int l, int r) {
        return pre(r) - pre(l - 1);
    }
}

class Solution {
    private record Query(int x1, int x2, int y1, int y2, long area) {
    }

    private record Data(int qid, int sign, int y1, int y2) {
    }

    public long maxRectangleArea(int[] xCoord, int[] yCoord) {
        Map<Integer, List<Integer>> xMap = new HashMap<>(); // 同一列的所有点的纵坐标
        Map<Integer, List<Integer>> yMap = new HashMap<>(); // 同一行的所有点的横坐标
        int n = xCoord.length;
        for (int i = 0; i < n; i++) {
            xMap.computeIfAbsent(xCoord[i], k -> new ArrayList<>()).add(yCoord[i]);
            yMap.computeIfAbsent(yCoord[i], k -> new ArrayList<>()).add(xCoord[i]);
        }

        // 预处理每个点的正下方的点
        Map<Long, Integer> below = new HashMap<>();
        for (var e : xMap.entrySet()) {
            int x = e.getKey();
            List<Integer> ys = e.getValue();
            ys.sort(null);
            for (int i = 1; i < ys.size(); i++) {
                // 通过 x<<32|y 的方式，把 (x,y) 压缩成一个 long
                below.put((long) x << 32 | ys.get(i), ys.get(i - 1));
            }
        }

        // 预处理每个点的正左边的点
        Map<Long, Integer> left = new HashMap<>();
        for (var e : yMap.entrySet()) {
            int y = e.getKey();
            List<Integer> xs = e.getValue();
            xs.sort(null);
            for (int i = 1; i < xs.size(); i++) {
                left.put((long) xs.get(i) << 32 | y, xs.get(i - 1));
            }
        }

        // 离散化用
        List<Integer> xs = new ArrayList<>(xMap.keySet());
        List<Integer> ys = new ArrayList<>(yMap.keySet());
        xs.sort(null);
        ys.sort(null);

        // 收集询问：矩形区域（包括边界）的点的个数
        List<Query> queries = new ArrayList<>();
        // 枚举 (x2,y2) 作为矩形的右上角
        for (var e : xMap.entrySet()) {
            int x2 = e.getKey();
            List<Integer> listY = e.getValue();
            for (int i = 1; i < listY.size(); i++) {
                // 计算矩形左下角 (x1,y1)
                int y2 = listY.get(i);
                int x1 = left.getOrDefault((long) x2 << 32 | y2, -1);
                if (x1 < 0) {
                    continue;
                }
                int y1 = listY.get(i - 1); // (x2,y2) 下面的点（矩形右下角）的纵坐标
                // 矩形右下角的左边的点的横坐标必须是 x1
                if (left.getOrDefault((long) x2 << 32 | y1, -1) != x1) {
                    continue;
                }
                // 矩形左上角的下边的点的纵坐标必须是 y1
                if (below.getOrDefault((long) x1 << 32 | y2, -1) != y1) {
                    continue;
                }
                queries.add(new Query(
                     Collections.binarySearch(xs, x1), // 离散化
                     Collections.binarySearch(xs, x2),
                     Collections.binarySearch(ys, y1),
                     Collections.binarySearch(ys, y2),
                     (long) (x2 - x1) * (y2 - y1)
                ));
            }
        }

        // 离线询问
        List<Data>[] qs = new ArrayList[xs.size()];
        Arrays.setAll(qs, i -> new ArrayList<>());
        for (int i = 0; i < queries.size(); i++) {
            Query q = queries.get(i);
            if (q.x1 > 0) {
                qs[q.x1 - 1].add(new Data(i, -1, q.y1, q.y2));
            }
            qs[q.x2].add(new Data(i, 1, q.y1, q.y2));
        }

        // 回答询问
        int[] res = new int[queries.size()];
        Fenwick tree = new Fenwick(ys.size() + 1);
        for (int i = 0; i < xs.size(); i++) {
            // 把横坐标为 xs[i] 的所有点都加到树状数组中
            for (int y : xMap.get(xs.get(i))) {
                tree.add(Collections.binarySearch(ys, y) + 1); // 离散化
            }
            for (Data q : qs[i]) {
                // 查询横坐标 <= x（已满足）且纵坐标在 [y1,y2] 中的点的个数
                res[q.qid] += q.sign * tree.query(q.y1 + 1, q.y2 + 1);
            }
        }

        long ans = -1;
        for (int i = 0; i < res.length; i++) {
            if (res[i] == 4) { // 矩形区域（包括边界）恰好有 4 个点
                ans = Math.max(ans, queries.get(i).area);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
// 树状数组模板
class Fenwick {
    vector<int> tree;

public:
    Fenwick(int n) : tree(n, 0) {}

    void add(int i) {
        while (i < tree.size()) {
            tree[i]++;
            i += i & -i;
        }
    }

    // [1,i] 中的元素和
    int pre(int i) {
        int res = 0;
        while (i > 0) {
            res += tree[i];
            i -= i & -i;
        }
        return res;
    }

    // [l,r] 中的元素和
    int query(int l, int r) {
        return pre(r) - pre(l - 1);
    }
};

class Solution {
public:
    long long maxRectangleArea(vector<int>& xCoord, vector<int>& yCoord) {
        unordered_map<int, vector<int>> x_map; // 同一列的所有点的纵坐标
        unordered_map<int, vector<int>> y_map; // 同一行的所有点的横坐标
        int n = xCoord.size();
        for (int i = 0; i < n; i++) {
            x_map[xCoord[i]].push_back(yCoord[i]);
            y_map[yCoord[i]].push_back(xCoord[i]);
        }

        // 预处理每个点的正下方的点
        unordered_map<long long, int> below;
        for (auto& [x, ys] : x_map) {
            ranges::sort(ys);
            for (int i = 1; i < ys.size(); i++) {
                // 通过 x<<32|y 的方式，把 (x,y) 压缩成一个 long long
                below[(long long) x << 32 | ys[i]] = ys[i - 1];
            }
        }

        // 预处理每个点的正左边的点
        unordered_map<long long, int> left;
        for (auto& [y, xs] : y_map) {
            ranges::sort(xs);
            for (int i = 1; i < xs.size(); i++) {
                left[(long long) xs[i] << 32 | y] = xs[i - 1];
            }
        }

        // 离散化用
        vector<int> xs;
        for (auto& [x, _] : x_map) {
            xs.push_back(x);
        }
        ranges::sort(xs);

        vector<int> ys;
        for (auto& [y, _] : y_map) {
            ys.push_back(y);
        }
        ranges::sort(ys);

        // 收集询问：矩形区域（包括边界）的点的个数
        struct Query {
            int x1, x2, y1, y2;
            long long area;
        };
        vector<Query> queries;
        // 枚举 (x2,y2) 作为矩形的右上角
        for (auto& [x2, list_y] : x_map) {
            for (int i = 1; i < list_y.size(); i++) {
                // 计算矩形左下角 (x1,y1)
                int y2 = list_y[i];
                auto it = left.find((long long) x2 << 32 | y2);
                if (it == left.end()) {
                    continue;
                }
                int x1 = it->second;
                int y1 = list_y[i - 1]; // (x2,y2) 下面的点（矩形右下角）的纵坐标
                // 矩形右下角的左边的点的横坐标必须是 x1
                it = left.find((long long) x2 << 32 | y1);
                if (it == left.end() || it->second != x1) {
                    continue;
                }
                // 矩形左上角的下边的点的纵坐标必须是 y1
                it = below.find((long long) x1 << 32 | y2);
                if (it == left.end() || it->second != y1) {
                    continue;
                }
                queries.emplace_back(
                     ranges::lower_bound(xs, x1) - xs.begin(), // 离散化
                     ranges::lower_bound(xs, x2) - xs.begin(),
                     ranges::lower_bound(ys, y1) - ys.begin(),
                     ranges::lower_bound(ys, y2) - ys.begin(),
                     (long long) (x2 - x1) * (y2 - y1)
                );
            }
        }

        // 离线询问
        struct Data {
            int qid, sign, y1, y2;
        };
        vector<vector<Data>> qs(xs.size());
        for (int i = 0; i < queries.size(); i++) {
            auto& [x1, x2, y1, y2, _] = queries[i];
            if (x1 > 0) {
                qs[x1 - 1].emplace_back(i, -1, y1, y2);
            }
            qs[x2].emplace_back(i, 1, y1, y2);
        }

        // 回答询问
        vector<int> res(queries.size());
        Fenwick tree(ys.size() + 1);
        for (int i = 0; i < xs.size(); i++) {
            // 把横坐标为 xs[i] 的所有点都加到树状数组中
            for (int y : x_map[xs[i]]) {
                tree.add(ranges::lower_bound(ys, y) - ys.begin() + 1); // 离散化
            }
            for (auto& [qid, sign, y1, y2] : qs[i]) {
                // 查询横坐标 <= x（已满足）且纵坐标在 [y1,y2] 中的点的个数
                res[qid] += sign * tree.query(y1 + 1, y2 + 1);
            }
        }

        long long ans = -1;
        for (int i = 0; i < res.size(); i++) {
            if (res[i] == 4) { // 矩形区域（包括边界）恰好有 4 个点
                ans = max(ans, queries[i].area);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
// 树状数组模板
type fenwick []int

func (f fenwick) add(i int) {
	for ; i < len(f); i += i & -i {
		f[i]++
	}
}

// [1,i] 中的元素和
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

// [l,r] 中的元素和
func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

func maxRectangleArea(xCoord, yCoord []int) int64 {
	xMap := map[int][]int{} // 同一列的所有点的纵坐标
	yMap := map[int][]int{} // 同一行的所有点的横坐标
	for i, x := range xCoord {
		y := yCoord[i]
		xMap[x] = append(xMap[x], y)
		yMap[y] = append(yMap[y], x)
	}

	// 预处理每个点的正下方的点
	type pair struct{ x, y int }
	below := map[pair]int{}
	for x, ys := range xMap {
		slices.Sort(ys)
		for i := 1; i < len(ys); i++ {
			below[pair{x, ys[i]}] = ys[i-1]
		}
	}

	// 预处理每个点的正左边的点
	left := map[pair]int{}
	for y, xs := range yMap {
		slices.Sort(xs)
		for i := 1; i < len(xs); i++ {
			left[pair{xs[i], y}] = xs[i-1]
		}
	}

	// 离散化用
	xs := slices.Sorted(maps.Keys(xMap))
	ys := slices.Sorted(maps.Keys(yMap))

	// 收集询问：矩形区域（包括边界）的点的个数
	type query struct{ x1, x2, y1, y2, area int }
	queries := []query{}
	// 枚举 (x2,y2) 作为矩形的右上角
	for x2, listY := range xMap {
		for i := 1; i < len(listY); i++ {
			// 计算矩形左下角 (x1,y1)
			y2 := listY[i]
			x1, ok := left[pair{x2, y2}]
			if !ok {
				continue
			}
			y1 := listY[i-1] // (x2,y2) 下面的点（矩形右下角）的纵坐标
			// 矩形右下角的左边的点的横坐标必须是 x1
			if x, ok := left[pair{x2, y1}]; !ok || x != x1 {
				continue
			}
			// 矩形左上角的下边的点的纵坐标必须是 y1
			if y, ok := below[pair{x1, y2}]; !ok || y != y1 {
				continue
			}
			queries = append(queries, query{
				sort.SearchInts(xs, x1), // 离散化
				sort.SearchInts(xs, x2),
				sort.SearchInts(ys, y1),
				sort.SearchInts(ys, y2),
				(x2 - x1) * (y2 - y1),
			})
		}
	}

	// 离线询问
	type data struct{ qid, sign, y1, y2 int }
	qs := make([][]data, len(xs))
	for i, q := range queries {
		if q.x1 > 0 {
			qs[q.x1-1] = append(qs[q.x1-1], data{i, -1, q.y1, q.y2})
		}
		qs[q.x2] = append(qs[q.x2], data{i, 1, q.y1, q.y2})
	}

	// 回答询问
	res := make([]int, len(queries))
	tree := make(fenwick, len(ys)+1)
	for i, x := range xs {
		// 把横坐标为 x 的所有点都加到树状数组中
		for _, y := range xMap[x] {
			tree.add(sort.SearchInts(ys, y) + 1) // 离散化
		}
		for _, q := range qs[i] {
			// 查询横坐标 <= x（已满足）且纵坐标在 [y1,y2] 中的点的个数
			res[q.qid] += q.sign * tree.query(q.y1+1, q.y2+1)
		}
	}

	ans := -1
	for i, cnt := range res {
		if cnt == 4 { // 矩形区域（包括边界）恰好有 4 个点
			ans = max(ans, queries[i].area)
		}
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{xCoord}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

- [CF1899G. Unusual Entertainment](https://codeforces.com/contest/1899/problem/G)

欢迎评论补充其他相似题目。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. 【本题相关】[常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
