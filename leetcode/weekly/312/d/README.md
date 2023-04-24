### 视频讲解

见[【周赛 312】](https://www.bilibili.com/video/BV1ve411K7P5/)第四题。

### 提示 1

当涉及到最大/最小的约束时，往往要按照节点值的顺序思考。

应该从大到小，还是从小到大呢？

### 提示 2

一种比较暴力的思路是，先考虑节点值最大的点，从这些点中任选两个点，作为路径的开始节点和结束节点，这样的路径都是满足题目要求的。

设节点值最大的点有 $x$ 个，那么会形成 $C(x,2) = \dfrac{x(x-1)}{2}$ 条长度大于 $1$ 的好路径。（单个节点的好路径单独统计。）

然后从树中删去这些点，再递归考虑剩下的各个连通块内的好路径个数。

但这样每个节点会被多次统计，最坏情况下的时间复杂度为 $O(n^2)$。

### 提示 3

倒着思考这一过程，把删除改为合并，你想到了哪个数据结构？

### 提示 4

并查集。

### 提示 5

按节点值从小到大考虑。

用并查集合并时，**总是从节点值小的点往节点值大的点合并**，这样可以保证连通块的代表元的节点值是最大的。

对于节点 $x$ 及其邻居 $y$，如果 $y$ 所处的连通分量的最大节点值不超过 $\textit{vals}[x]$，那么可以把 $y$ 所处的连通块合并到 $x$ 所处的连通块中。

如果此时这两个连通块的最大节点值相同，那么可以根据乘法原理，把这两个连通块内的等于最大节点值的节点个数相乘，加到答案中。

```py [sol1-Python3]
class Solution:
    def numberOfGoodPaths(self, vals: List[int], edges: List[List[int]]) -> int:
        n = len(vals)
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)  # 建图

        # 并查集模板
        fa = list(range(n))
        # size[x] 表示节点值等于 vals[x] 的节点个数，
        # 如果按照节点值从小到大合并，size[x] 也是连通块内的等于最大节点值的节点个数
        size = [1] * n
        def find(x: int) -> int:
            if fa[x] != x:
                fa[x] = find(fa[x])
            return fa[x]

        ans = n  # 单个节点的好路径
        for vx, x in sorted(zip(vals, range(n))):
            fx = find(x)
            for y in g[x]:
                y = find(y)
                if y == fx or vals[y] > vx:
                    continue  # 只考虑最大节点值不超过 vx 的连通块
                if vals[y] == vx:  # 可以构成好路径
                    ans += size[fx] * size[y]  # 乘法原理
                    size[fx] += size[y]  # 统计连通块内节点值等于 vx 的节点个数
                fa[y] = fx  # 把小的节点值合并到大的节点值上
        return ans
```

```java [sol1-Java]
class Solution {
    public int numberOfGoodPaths(int[] vals, int[][] edges) {
        int n = vals.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (var e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x); // 建图
        }

        fa = new int[n];
        var id = new Integer[n];
        for (int i = 0; i < n; i++)
            fa[i] = id[i] = i;
        Arrays.sort(id, (i, j) -> vals[i] - vals[j]);
        // size[x] 表示节点值等于 vals[x] 的节点个数，
        // 如果按照节点值从小到大合并，size[x] 也是连通块内的等于最大节点值的节点个数
        var size = new int[n];
        Arrays.fill(size, 1);

        int ans = n; // 单个节点的好路径
        for (var x : id) {
            int vx = vals[x], fx = find(x);
            for (var y : g[x]) {
                y = find(y);
                if (y == fx || vals[y] > vx)
                    continue; // 只考虑最大节点值不超过 vx 的连通块
                if (vals[y] == vx) { // 可以构成好路径
                    ans += size[fx] * size[y]; // 乘法原理
                    size[fx] += size[y]; // 统计连通块内节点值等于 vx 的节点个数
                }
                fa[y] = fx; // 把小的节点值合并到大的节点值上
            }
        }
        return ans;
    }

    private int[] fa;

    private int find(int x) {
        if (fa[x] != x) fa[x] = find(fa[x]);
        return fa[x];
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int numberOfGoodPaths(vector<int> &vals, vector<vector<int>> &edges) {
        int n = vals.size();
        vector<vector<int>> g(n);
        for (auto &e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x); // 建图
        }

        // 并查集模板
        // size[x] 表示节点值等于 vals[x] 的节点个数，
        // 如果按照节点值从小到大合并，size[x] 也是连通块内的等于最大节点值的节点个数
        int id[n], fa[n], size[n]; // id 后面排序用
        iota(id, id + n, 0);
        iota(fa, fa + n, 0);
        fill(size, size + n, 1);
        function<int(int)> find = [&](int x) -> int { return fa[x] == x ? x : fa[x] = find(fa[x]); };

        int ans = n; // 单个节点的好路径
        sort(id, id + n, [&](int i, int j) { return vals[i] < vals[j]; });
        for (int x : id) {
            int vx = vals[x], fx = find(x);
            for (int y : g[x]) {
                y = find(y);
                if (y == fx || vals[y] > vx)
                    continue; // 只考虑最大节点值不超过 vx 的连通块
                if (vals[y] == vx) { // 可以构成好路径
                    ans += size[fx] * size[y]; // 乘法原理
                    size[fx] += size[y]; // 统计连通块内节点值等于 vx 的节点个数
                }
                fa[y] = fx; // 把小的节点值合并到大的节点值上
            }
        }
        return ans;
    }
};
```

```go [sol1-Go]
func numberOfGoodPaths(vals []int, edges [][]int) int {
	n := len(vals)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建图
	}

	// 并查集模板
	fa := make([]int, n)
	// size[x] 表示节点值等于 vals[x] 的节点个数，
	// 如果按照节点值从小到大合并，size[x] 也是连通块内的等于最大节点值的节点个数
	size := make([]int, n) 
	id := make([]int, n) // 后面排序用
	for i := range fa {
		fa[i] = i
		size[i] = 1
		id[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	ans := n // 单个节点的好路径
	sort.Slice(id, func(i, j int) bool { return vals[id[i]] < vals[id[j]] })
	for _, x := range id {
		vx := vals[x]
		fx := find(x)
		for _, y := range g[x] {
			y = find(y)
			if y == fx || vals[y] > vx {
				continue // 只考虑最大节点值不超过 vx 的连通块
			}
			if vals[y] == vx { // 可以构成好路径
				ans += size[fx] * size[y] // 乘法原理
				size[fx] += size[y] // 统计连通块内节点值等于 vx 的节点个数
			}
			fa[y] = fx // 把小的节点值合并到大的节点值上
		}
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{vals}$ 的长度。
- 空间复杂度：$O(n)$。

### 相似题目

- [Codeforces 915F. Imbalance Value of a Tree](https://codeforces.com/problemset/problem/915/F)
