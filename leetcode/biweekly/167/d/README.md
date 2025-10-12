## 方法一：二分答案 + 判断二分图

### 转化

假设划分因子 $\ge \textit{low}$。

$\textit{low}$ 越小，要求就越**宽松**，越能找到一个合法划分。

$\textit{low}$ 越大，要求就越**苛刻**，越不能找到一个合法划分。

据此，可以**二分猜答案**。关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

现在问题变成一个判定性问题：

- 给定 $\textit{low}$，能否把点集分成两个非空集合，使得每个集合内部的任意点对的曼哈顿距离都 $\ge \textit{low}$？

### 思路

对于曼哈顿距离 $< \textit{low}$ 的点对（非法点对），连一条边，我们可以得到一个无向图。

我们需要把这些点分成两个集合，每个集合内部不能有边（否则违背要求）。换句话说，每条边的两个端点，一定来自不同集合。

这和 [785. 判断二分图](https://leetcode.cn/problems/is-graph-bipartite/) 是完全一样的，做法见[【图解】交替染色法](https://leetcode.cn/problems/is-graph-bipartite/solutions/3803670/tu-jie-jiao-ti-ran-se-fa-pythonjavaccgoj-ov27/)。

### 细节

下面代码采用开区间二分，这仅仅是二分的一种写法，使用闭区间或者半闭半开区间都是可以的，喜欢哪种写法就用哪种。

- 开区间左端点初始值：$0$。可以随意划分，一定满足要求。
- 开区间右端点初始值：最大曼哈顿距离加一。此时一定无法满足要求。

> 对于开区间写法，简单来说 `check(mid) == true` 时更新的是谁，最后就返回谁。相比其他二分写法，开区间写法不需要思考加一减一等细节，更简单。推荐使用开区间写二分。

注意特判 $n=2$ 的情况，此时划分因子恒为 $0$。

### 答疑

**问**：为什么二分结束后，答案 $\textit{ans}$ 一定来自 $\textit{points}$ 的某两个点的曼哈顿距离？

**答**：反证法。假设 $\textit{ans}$ 不来自 $\textit{points}$ 的某两个点的曼哈顿距离，这意味着最小曼哈顿距离 $> \textit{ans}$，也就是 $\ge \textit{ans}+1$。换句话说，$\text{check}(\textit{ans}+1)=\texttt{true}$。但根据循环不变量，二分结束后 $\text{check}(\textit{ans}+1)=\texttt{false}$，矛盾。故原命题成立。

[本题视频讲解](https://www.bilibili.com/video/BV16E4uzLEdK/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxPartitionFactor(self, points: List[List[int]]) -> int:
        if len(points) == 2:
            return 0

        # 原理见 785. 判断二分图
        def check(low: int) -> bool:
            colors = [0] * len(points)

            def dfs(x: int, c: int) -> bool:
                colors[x] = c
                x1, y1 = points[x]
                for y, (x2, y2) in enumerate(points):
                    if y == x or abs(x1 - x2) + abs(y1 - y2) >= low:  # 符合要求
                        continue
                    if colors[y] == c or colors[y] == 0 and not dfs(y, -c):
                        return False  # 不是二分图
                return True

            # 可能有多个连通块
            for i, c in enumerate(colors):
                if c == 0 and not dfs(i, 1):
                    return False
            return True

        max_dis = max(abs(x1 - x2) + abs(y1 - y2)
                      for (x1, y1), (x2, y2) in combinations(points, 2))

        left, right = 0, max_dis + 1
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid):
                left = mid
            else:
                right = mid
        return left
```

```py [sol-Python3 库函数]
class Solution:
    def maxPartitionFactor(self, points: List[List[int]]) -> int:
        if len(points) == 2:
            return 0

        # 原理见 785. 判断二分图
        def check(low: int) -> bool:
            # 二分最小的不满足要求的 low+1，就可以得到最大的满足要求的 low
            low += 1
            colors = [0] * len(points)

            def dfs(x: int, c: int) -> bool:
                colors[x] = c
                x1, y1 = points[x]
                for y, (x2, y2) in enumerate(points):
                    if y == x or abs(x1 - x2) + abs(y1 - y2) >= low:  # 符合要求
                        continue
                    if colors[y] == c or colors[y] == 0 and not dfs(y, -c):
                        return False  # 不是二分图
                return True

            # 可能有多个连通块
            for i, c in enumerate(colors):
                if c == 0 and not dfs(i, 1):
                    return True
            return False

        max_dis = max(abs(x1 - x2) + abs(y1 - y2)
                      for (x1, y1), (x2, y2) in combinations(points, 2))
        return bisect_left(range(max_dis), True, key=check)
```

```java [sol-Java]
class Solution {
    public int maxPartitionFactor(int[][] points) {
        int n = points.length;
        if (n == 2) {
            return 0;
        }

        int maxDis = 0;
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                int d = Math.abs(points[i][0] - points[j][0]) + Math.abs(points[i][1] - points[j][1]);
                maxDis = Math.max(maxDis, d);
            }
        }

        int left = 0, right = maxDis + 1;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            if (check(mid, points)) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return left;
    }

    // 原理见 785. 判断二分图
    private boolean check(int low, int[][] points) {
        int n = points.length;
        int[] colors = new int[n];
        for (int i = 0; i < n; i++) {
            if (colors[i] == 0 && !dfs(i, 1, low, points, colors)) {
                return false;
            }
        }
        return true;
    }

    private boolean dfs(int x, int c, int low, int[][] points, int[] colors) {
        colors[x] = c;
        int x1 = points[x][0], y1 = points[x][1];
        for (int y = 0; y < points.length; y++) {
            if (y == x || Math.abs(x1 - points[y][0]) + Math.abs(y1 - points[y][1]) >= low) { // 符合要求
                continue;
            }
            if (colors[y] == c || colors[y] == 0 && !dfs(y, -c, low, points, colors)) {
                return false; // 不是二分图
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxPartitionFactor(vector<vector<int>>& points) {
        int n = points.size();
        if (n == 2) {
            return 0;
        }

        // 原理见 785. 判断二分图
        auto check = [&](int low) -> bool {
            vector<int8_t> colors(n);

            auto dfs = [&](this auto&& dfs, int x, int8_t c) -> bool {
                colors[x] = c;
                auto& p = points[x];
                for (int y = 0; y < n; y++) {
                    auto& q = points[y];
                    if (y == x || abs(p[0] - q[0]) + abs(p[1] - q[1]) >= low) { // 符合要求
                        continue;
                    }
                    if (colors[y] == c || colors[y] == 0 && !dfs(y, -c)) {
                        return false; // 不是二分图
                    }
                }
                return true;
            };

            // 可能有多个连通块
            for (int i = 0; i < n; i++) {
                if (colors[i] == 0 && !dfs(i, 1)) {
                    return false;
                }
            }
            return true;
        };

        int max_dis = 0;
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                max_dis = max(max_dis, abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1]));
            }
        }

        int left = 0, right = max_dis + 1;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            (check(mid) ? left : right) = mid;
        }
        return left;
    }
};
```

```go [sol-Go]
// 原理见 785. 判断二分图
func isBipartite(points [][]int, low int) bool {
	colors := make([]int8, len(points))

	var dfs func(int, int8) bool
	dfs = func(x int, c int8) bool {
		colors[x] = c
		p := points[x]
		for y, q := range points {
			if y == x || abs(p[0]-q[0])+abs(p[1]-q[1]) >= low { // 符合要求
				continue
			}
			if colors[y] == c || colors[y] == 0 && !dfs(y, -c) {
				return false // 不是二分图
			}
		}
		return true
	}

	// 可能有多个连通块
	for i, c := range colors {
		if c == 0 && !dfs(i, 1) {
			return false
		}
	}
	return true
}

func maxPartitionFactor(points [][]int) int {
	n := len(points)
	if n == 2 {
		return 0
	}

	// 不想算的话可以写 maxDis := int(4e8)
	maxDis := 0
	for i, p := range points {
		for _, q := range points[:i] {
			maxDis = max(maxDis, abs(p[0]-q[0])+abs(p[1]-q[1]))
		}
	}

	return sort.Search(maxDis, func(low int) bool {
		// 二分最小的不满足要求的 low+1，就可以得到最大的满足要求的 low
		return !isBipartite(points, low+1)
	})
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2\log U)$，其中 $n$ 是 $\textit{points}$ 的长度，$U\le 4\times 10^8$ 是曼哈顿距离的最大值。
- 空间复杂度：$\mathcal{O}(n)$。

**注**：也可以把所有曼哈顿距离存到一个列表中，排序，在列表中二分，从而做到 $\mathcal{O}(n^2\log n)$ 的时间。

## 方法二：排序 + 并查集

像 Kruskal 算法那样，按照曼哈顿距离从小到大处理点对。

这些点对是互斥的，不能在同一个集合中。

这可以用**带权并查集**解决，见 [399. 除法求值](https://leetcode.cn/problems/evaluate-division/)。

本题边权为 $1$，距离算子为模 $2$ 意义下的加法，即异或运算。

```py [sol-Python3]
class UnionFind:
    def __init__(self, n: int):
        self.fa = list(range(n))
        self.dis = [0] * n  # dis[x] 表示 x 到其代表元的距离

    # 返回 x 所在集合的代表元
    # 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
    def find(self, x: int) -> int:
        fa = self.fa
        if fa[x] != x:
            rt = self.find(fa[x])
            self.dis[x] ^= self.dis[fa[x]]  # 更新 x 到其代表元的距离
            fa[x] = rt
        return fa[x]

    # 合并两个互斥的点
    # 如果已经合并，返回是否与已知条件矛盾
    def merge(self, from_: int, to: int) -> bool:
        x, y = self.find(from_), self.find(to)
        dis = self.dis
        if x == y:  # from 和 to 在同一个集合，不合并
            return dis[from_] != dis[to]  # 必须在不同集合
        #    2 ------ 4
        #   /        /
        #  1 ------ 3
        # 如果知道 1->2 的距离和 3->4 的距离，现在合并 1 和 3，并传入 1->3 的距离（本题等于 1）
        # 由于 1->3->4 和 1->2->4 的距离相等
        # 所以 2->4 的距离为 (1->3) + (3->4) - (1->2)
        dis[x] = 1 ^ dis[to] ^ dis[from_]
        self.fa[x] = y
        return True


class Solution:
    def maxPartitionFactor(self, points: List[List[int]]) -> int:
        manhattan_tuples = [(abs(x1 - x2) + abs(y1 - y2), i, j)
                            for i, (x1, y1) in enumerate(points) for j, (x2, y2) in enumerate(points[:i])]
        manhattan_tuples.sort(key=lambda t: t[0])

        uf = UnionFind(len(points))
        for dis, x, y in manhattan_tuples:
            if not uf.merge(x, y):
                return dis  # x 和 y 必须在同一个集合，dis 就是这一划分的最小划分因子
        return 0
```

```java [sol-Java]
class UnionFind {
    private final int[] fa;
    private final int[] dis;

    public UnionFind(int n) {
        fa = new int[n];
        for (int i = 0; i < n; i++) {
            fa[i] = i;
        }
        dis = new int[n];
    }

    // 返回 x 所在集合的代表元
    // 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
    public int find(int x) {
        if (fa[x] != x) {
            int rt = find(fa[x]);
            dis[x] ^= dis[fa[x]]; // 更新 x 到其代表元的距离
            fa[x] = rt;
        }
        return fa[x];
    }

    // 合并两个互斥的点
    // 如果已经合并，返回是否与已知条件矛盾
    public boolean merge(int from, int to) {
        int x = find(from);
        int y = find(to);
        if (x == y) { // from 和 to 在同一个集合，不合并
            return dis[from] != dis[to]; // 必须在不同集合
        }
        //    2 ------ 4
        //   /        /
        //  1 ------ 3
        // 如果知道 1->2 的距离和 3->4 的距离，现在合并 1 和 3，并传入 1->3 的距离（本题等于 1）
        // 由于 1->3->4 和 1->2->4 的距离相等
        // 所以 2->4 的距离为 (1->3) + (3->4) - (1->2)
        dis[x] = 1 ^ dis[to] ^ dis[from];
        fa[x] = y;
        return true;
    }
}

class Solution {
    public int maxPartitionFactor(int[][] points) {
        int n = points.length;
        int[][] manhattanTuples = new int[n * (n - 1) / 2][]; // [dis, x, y]
        int idx = 0;
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                int dis = Math.abs(points[i][0] - points[j][0]) + Math.abs(points[i][1] - points[j][1]);
                manhattanTuples[idx++] = new int[]{dis, i, j};
            }
        }
        Arrays.sort(manhattanTuples, (a, b) -> a[0] - b[0]);

        UnionFind uf = new UnionFind(n);
        for (int[] t : manhattanTuples) {
            if (!uf.merge(t[1], t[2])) {
                return t[0]; // x=t[1] 和 y=t[2] 必须在同一个集合，dis=t[0] 就是这一划分的最小划分因子
            }
        }
        return 0;
    }
}
```

```cpp [sol-C++]
class UnionFind {
    vector<int> fa;
    vector<int8_t> dis; // dis[x] 表示 x 到其代表元的距离

public:
    UnionFind(int n) : fa(n), dis(n) {
        ranges::iota(fa, 0);
    }

    // 返回 x 所在集合的代表元
    // 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
    int find(int x) {
        if (fa[x] != x) {
            int rt = find(fa[x]);
            dis[x] ^= dis[fa[x]]; // 更新 x 到其代表元的距离
            fa[x] = rt;
        }
        return fa[x];
    }

    // 合并两个互斥的点
    // 如果已经合并，返回是否与已知条件矛盾
    bool merge(int from, int to) {
        int x = find(from), y = find(to);
        if (x == y) { // from 和 to 在同一个集合，不合并
            return dis[from] != dis[to]; // 必须在不同集合
        }
        //    2 ------ 4
        //   /        /
        //  1 ------ 3
        // 如果知道 1->2 的距离和 3->4 的距离，现在合并 1 和 3，并传入 1->3 的距离（本题等于 1）
        // 由于 1->3->4 和 1->2->4 的距离相等
        // 所以 2->4 的距离为 (1->3) + (3->4) - (1->2)
        dis[x] = 1 ^ dis[to] ^ dis[from];
        fa[x] = y;
        return true;
    }
};

class Solution {
public:
    int maxPartitionFactor(vector<vector<int>>& points) {
        int n = points.size();
        vector<tuple<int, int, int>> manhattan_tuples;
        manhattan_tuples.reserve(n * (n - 1) / 2); // 预分配空间
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                int dis = abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1]);
                manhattan_tuples.emplace_back(dis, i, j);
            }
        }
        ranges::sort(manhattan_tuples, {}, [](const auto& t) { return get<0>(t); });

        UnionFind uf(n);
        for (auto& [dis, x, y] : manhattan_tuples) {
            if (!uf.merge(x, y)) {
                return dis; // x 和 y 必须在同一个集合，dis 就是这一划分的最小划分因子
            }
        }
        return 0;
    }
};
```

```go [sol-Go]
type unionFind struct {
	fa  []int
	dis []int8 // dis[x] 表示 x 到其代表元的距离
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	dis := make([]int8, n)
	for i := range fa {
		fa[i] = i
	}
	return unionFind{fa, dis}
}

// 返回 x 所在集合的代表元
// 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
func (u unionFind) find(x int) int {
	if u.fa[x] != x {
		rt := u.find(u.fa[x])
		u.dis[x] ^= u.dis[u.fa[x]] // 更新 x 到其代表元的距离
		u.fa[x] = rt
	}
	return u.fa[x]
}

// 合并两个互斥的点
// 如果已经合并，返回是否与已知条件矛盾
func (u *unionFind) merge(from, to int) bool {
	x, y := u.find(from), u.find(to)
	if x == y { // from 和 to 在同一个集合，不合并
		return u.dis[from] != u.dis[to] // 必须在不同集合
	}
	//    2 ------ 4
	//   /        /
	//  1 ------ 3
	// 如果知道 1->2 的距离和 3->4 的距离，现在合并 1 和 3，并传入 1->3 的距离（本题等于 1）
	// 由于 1->3->4 和 1->2->4 的距离相等
	// 所以 2->4 的距离为 (1->3) + (3->4) - (1->2)
	u.dis[x] = 1 ^ u.dis[to] ^ u.dis[from]
	u.fa[x] = y
	return true
}

func maxPartitionFactor(points [][]int) int {
	n := len(points)
	type tuple struct{ dis, x, y int }
	manhattanTuples := make([]tuple, 0, n*(n-1)/2) // 预分配空间
	for i, p := range points {
		for j, q := range points[:i] {
			manhattanTuples = append(manhattanTuples, tuple{abs(p[0]-q[0]) + abs(p[1]-q[1]), i, j})
		}
	}
	slices.SortFunc(manhattanTuples, func(a, b tuple) int { return a.dis - b.dis })

	uf := newUnionFind(n)
	for _, t := range manhattanTuples {
		if !uf.merge(t.x, t.y) {
			return t.dis // t.x 和 t.y 必须在同一个集合，t.dis 就是这一划分的最小划分因子
		}
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2\log n)$，其中 $n$ 是 $\textit{points}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

## 专题训练

见下面图论题单的「**七、二分图染色**」。

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
