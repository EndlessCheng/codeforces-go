下午 2 点在 B 站直播讲周赛的题目，包括本题**拓扑排序的原理**，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

#### 提示 1

数字之间的约束只发生在行与行、列于列，而行与列之间没有任何约束。

因此我们可以分别处理行与列中数字的相对顺序，如何求出这个相对顺序呢？

#### 提示 2

拓扑排序。

#### 提示 3

对于 $\textit{rowConditions}$，我们可以从 $\textit{above}_i$ 向 $\textit{below}_i$ 连一条有向边，得到一张有向图。在这张图上跑拓扑排序，得到的拓扑序就是行与行中数字的相对顺序，这样我们就知道了每一行要填哪个数字。如果得到的拓扑序长度不足 $k$，说明图中有环，无法构造，答案不存在。

对 $\textit{colConditions}$ 也执行上述过程，得到每一列要填哪个数字，进而得到每个数字要填到哪一列中，这样我们就知道每一行的数字要填到哪一列了。

#### 复杂度分析

- 时间复杂度：$O(k+m)$，其中 $m$ 为 $\textit{rowConditions}$ 的长度。
- 空间复杂度：$O(k+m)$。忽略返回值占用的空间复杂度。

```py [sol1-Python3]
class Solution:
    def buildMatrix(self, k: int, rowConditions: List[List[int]], colConditions: List[List[int]]) -> List[List[int]]:
        def topo_sort(edges: List[List[int]]) -> List[int]:
            g = [[] for _ in range(k)]
            in_deg = [0] * k
            for x, y in edges:
                g[x - 1].append(y - 1)  # 顶点编号从 0 开始，方便计算
                in_deg[y - 1] += 1
            order = []
            q = deque(i for i, d in enumerate(in_deg) if d == 0)
            while q:
                x = q.popleft()
                order.append(x)
                for y in g[x]:
                    in_deg[y] -= 1
                    if in_deg[y] == 0:
                        q.append(y)
            return order if len(order) == k else None

        if (row := topo_sort(rowConditions)) is None or (col := topo_sort(colConditions)) is None:
            return []
        pos = {x: i for i, x in enumerate(col)}
        ans = [[0] * k for _ in range(k)]
        for i, x in enumerate(row):
            ans[i][pos[x]] = x + 1
        return ans
```

```java [sol1-Java]
class Solution {
    int[] topoSort(int k, int[][] edges) {
        List<Integer>[] g = new ArrayList[k];
        Arrays.setAll(g, e -> new ArrayList<>());
        var inDeg = new int[k];
        for (var e : edges) {
            int x = e[0] - 1, y = e[1] - 1; // 顶点编号从 0 开始，方便计算
            g[x].add(y);
            ++inDeg[y];
        }

        var order = new ArrayList<Integer>();
        var q = new ArrayDeque<Integer>();
        for (var i = 0; i < k; ++i)
            if (inDeg[i] == 0) q.push(i);
        while (!q.isEmpty()) {
            var x = q.pop();
            order.add(x);
            for (var y : g[x])
                if (--inDeg[y] == 0) q.push(y);
        }
        return order.stream().mapToInt(x -> x).toArray();
    }

    public int[][] buildMatrix(int k, int[][] rowConditions, int[][] colConditions) {
        int[] row = topoSort(k, rowConditions), col = topoSort(k, colConditions);
        if (row.length < k || col.length < k) return new int[][]{};
        var pos = new int[k];
        for (var i = 0; i < k; ++i)
            pos[col[i]] = i;
        var ans = new int[k][k];
        for (var i = 0; i < k; ++i)
            ans[i][pos[row[i]]] = row[i] + 1;
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
    vector<int> topo_sort(int k, vector<vector<int>> &edges) {
        vector<vector<int>> g(k);
        vector<int> in_deg(k);
        for (auto &e : edges) {
            int x = e[0] - 1, y = e[1] - 1; // 顶点编号从 0 开始，方便计算
            g[x].push_back(y);
            ++in_deg[y];
        }

        vector<int> order;
        queue<int> q;
        for (int i = 0; i < k; ++i)
            if (in_deg[i] == 0)
                q.push(i);
        while (!q.empty()) {
            int x = q.front();
            q.pop();
            order.push_back(x);
            for (int y : g[x])
                if (--in_deg[y] == 0)
                    q.push(y);
        }
        return order;
    }

public:
    vector<vector<int>> buildMatrix(int k, vector<vector<int>> &rowConditions, vector<vector<int>> &colConditions) {
        auto row = topo_sort(k, rowConditions), col = topo_sort(k, colConditions);
        if (row.size() < k || col.size() < k) return {};
        vector<int> pos(k);
        for (int i = 0; i < k; ++i)
            pos[col[i]] = i;
        vector<vector<int>> ans(k, vector<int>(k));
        for (int i = 0; i < k; ++i)
            ans[i][pos[row[i]]] = row[i] + 1;
        return ans;
    }
};
```

```go [sol1-Go]
func topoSort(k int, edges [][]int) []int {
	g := make([][]int, k)
	inDeg := make([]int, k)
	for _, e := range edges {
		x, y := e[0]-1, e[1]-1 // 顶点编号从 0 开始，方便计算
		g[x] = append(g[x], y)
		inDeg[y]++
	}
	q := make([]int, 0, k)
	orders := q // 复用队列作为拓扑序
	for i, d := range inDeg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range g[x] {
			if inDeg[y]--; inDeg[y] == 0 {
				q = append(q, y)
			}
		}
	}
	if cap(q) > 0 {
		return nil
	}
	return orders[:k]
}

func buildMatrix(k int, rowConditions, colConditions [][]int) [][]int {
	row := topoSort(k, rowConditions)
	col := topoSort(k, colConditions)
	if row == nil || col == nil {
		return nil
	}
	pos := make([]int, k)
	for i, v := range col {
		pos[v] = i
	}
	ans := make([][]int, k)
	for i, x := range row {
		ans[i] = make([]int, k)
		ans[i][pos[x]] = x + 1
	}
	return ans
}
```

#### 思考题

如果问题变成一个三维的立方格，再添加一个 $z$ 轴上的数字约束，要怎么做？
