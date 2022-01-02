本文将介绍处理基环树问题的一些通用技巧。

从 $i$ 向 $\textit{favorite}[i]$ 连边，我们可以得到一张有向图。由于每个大小为 $k$ 的连通块都有 $k$ 个点和 $k$ 条边，所以每个连通块必定有且仅有一个环，且由于每个点的出度均为 $1$，这样的有向图又叫做内向**基环树 (pseudotree)**，由基环树组成的森林叫**基环树森林 (pseudoforest)**。

每一个内向基环树（连通块）都由一个**基环**和其余指向基环的**树枝**组成，例如示例 $3$ 可以得到如下内向基环树，其基环由节点 $0$、$1$、$3$ 和 $4$ 组成，节点 $2$ 为其树枝：

![1.png](https://pic.leetcode-cn.com/1641096462-IsWZUX-1.png)


特别地，我们得到的基环可能只包含两个节点，例如示例 $1$ 可以得到如下内向基环树，其基环只包含节点 $1$ 和 $2$，而节点 $0$ 和 $3$ 组成其树枝：

![2.png](https://pic.leetcode-cn.com/1641096467-KCwxMo-2.png)


对于本题来说，这两类基环树在组成圆桌时会有明显区别，下文会说明这一点。

先来看看基环大小大于 $2$ 的情况，显然基环上的节点组成了一个环，因而可以组成一个圆桌；而树枝上的点，若插入圆桌上 $v\rightarrow w$ 这两人中间，会导致节点 $v$ 无法和其喜欢的员工坐在一起，因此树枝上的点是无法插入圆桌的；此外，树枝上的点也不能单独组成圆桌，因为这样会存在一个出度为 $0$ 的节点，其无法和其喜欢的员工坐在一起。对于其余内向基环树（连通块）上的节点，和树枝同理，也无法插入该基环组成的圆桌。

因此，对于基环大小大于 $2$ 的情况，圆桌的最大员工数目即为最大的基环大小，记作 $\textit{maxRingSize}$。

下面来分析基环大小等于 $2$ 的情况。

以如下基环树为例，$0$ 和 $1$ 组成基环，其余节点组成树枝：

![3.png](https://pic.leetcode-cn.com/1641096473-JtGBgY-3.png)


我们可以先让 $0$ 和 $1$ 坐在圆桌旁（假设 $0$ 坐在 $1$ 左侧），那么 $0$ 这一侧的树枝只能坐在 $0$ 的左侧，而 $1$ 这一侧的树枝只能坐在 $1$ 的右侧。

$2$ 可以紧靠着坐在 $0$ 的左侧，而 $3$ 和 $4$ 只能选一个坐在 $2$ 的左侧（如果 $4$ 紧靠着坐在 $2$ 的左侧，那么 $3$ 是无法紧靠着坐在 $4$ 的左侧的，反之亦然）。

这意味着从 $0$ 出发倒着找树枝上的点（即沿着反图上的边），每个点只能在其反图上选择其中一个子节点，因此 $0$ 这一侧的节点必须组成一条链，那么我们可以找最长的那条链，即上图加粗的节点。

对于 $1$ 这一侧也同理，将这两条最长链拼起来即为该基环树能组成的圆桌的最大员工数。

对于多个基环大小等于 $2$ 的基环树，每个基环树所对应的链，都可以拼在其余链的末尾，因此我们可以将这些链全部拼成一个圆桌，其大小记作 $\textit{sumListSize}$。

答案即为 $\max(\textit{maxRingSize},\textit{sumListSize})$。

---

下面介绍基环树问题的通用写法。

我们可以通过一次拓扑排序「剪掉」所有树枝，这样就可以将基环和树枝分开，从而简化后续处理流程：

- 如果要遍历基环，可以从入度不为 $0$ 的节点出发，在图上搜索；
- 如果要遍历树枝，可以以基环与树枝的连接处为起点，顺着反图来搜索树枝，从而将问题转化成一个树形问题。

对于本题，我们可以遍历所有基环，并按基环大小分类计算：

- 对于大小大于 $2$ 的基环，我们取基环大小的最大值；
- 对于大小等于 $2$ 的基环，我们可以从基环上的点出发，在反图上找到最大的树枝节点深度。
 
```go [sol1-Go]
func maximumInvitations(favorite []int) int {
	n := len(favorite)
	g := make([][]int, n)
	rg := make([][]int, n) // 图 g 的反图
	deg := make([]int, n)  // 图 g 上每个节点的入度
	for v, w := range favorite {
		g[v] = append(g[v], w)
		rg[w] = append(rg[w], v)
		deg[w]++
	}

	// 拓扑排序，剪掉图 g 上的所有树枝
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, w := range g[v] {
			if deg[w]--; deg[w] == 0 {
				q = append(q, w)
			}
		}
	}

	// 寻找图 g 上的基环
	ring := []int{}
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(v int) {
		vis[v] = true
		ring = append(ring, v)
		for _, w := range g[v] {
			if !vis[w] {
				dfs(w)
			}
		}
	}

	// 通过反图 rg 寻找树枝上最深的链
	maxDepth := 0
	var rdfs func(int, int, int)
	rdfs = func(v, fa, depth int) {
		maxDepth = max(maxDepth, depth)
		for _, w := range rg[v] {
			if w != fa {
				rdfs(w, v, depth+1)
			}
		}
	}

	maxRingSize, sumListSize := 0, 0
	for i, b := range vis {
		if !b && deg[i] > 0 { // 遍历基环上的点（拓扑排序后入度不为 0）
			ring = []int{}
			dfs(i)
			if len(ring) == 2 { // 基环大小为 2
				v, w := ring[0], ring[1]
				maxDepth = 0
				rdfs(v, w, 1)
				sumListSize += maxDepth // 累加 v 这一侧的最长链的长度
				maxDepth = 0
				rdfs(w, v, 1)
				sumListSize += maxDepth // 累加 w 这一侧的最长链的长度
			} else {
				maxRingSize = max(maxRingSize, len(ring)) // 取所有基环的最大值
			}
		}
	}
	return max(maxRingSize, sumListSize)
}

func max(a, b int) int { if b > a { return b }; return a }
```

```C++ [sol1-C++]
class Solution {
public:
    int maximumInvitations(vector<int> &favorite) {
        int n = favorite.size();
        vector<vector<int>> g(n), rg(n); // rg 为图 g 的反图
        vector<int> deg(n); // 图 g 上每个节点的入度
        for (int v = 0; v < n; ++v) {
            int w = favorite[v];
            g[v].emplace_back(w);
            rg[w].emplace_back(v);
            ++deg[w];
        }

        // 拓扑排序，剪掉图 g 上的所有树枝
        queue<int> q;
        for (int i = 0; i < n; ++i) {
            if (deg[i] == 0) {
                q.emplace(i);
            }
        }
        while (!q.empty()) {
            int v = q.front();
            q.pop();
            for (int w : g[v]) {
                if (--deg[w] == 0) {
                    q.emplace(w);
                }
            }
        }

        // 寻找图 g 上的基环
        vector<int> ring;
        vector<int> vis(n);
        function<void(int)> dfs = [&](int v) {
            vis[v] = true;
            ring.emplace_back(v);
            for (int w: g[v]) {
                if (!vis[w]) {
                    dfs(w);
                }
            }
        };

        // 通过反图 rg 寻找树枝上最深的链
        int max_depth = 0;
        function<void(int, int, int)> rdfs = [&](int v, int fa, int depth) {
            max_depth = max(max_depth, depth);
            for (int w: rg[v]) {
                if (w != fa) {
                    rdfs(w, v, depth + 1);
                }
            }
        };

        int max_ring_size = 0, sum_list_size = 0;
        for (int i = 0; i < n; ++i) {
            if (!vis[i] && deg[i]) { // 遍历基环上的点（拓扑排序后入度不为 0）
                ring.resize(0);
                dfs(i);
                int sz = ring.size();
                if (sz == 2) { // 基环大小为 2
                    int v = ring[0], w = ring[1];
                    max_depth = 0;
                    rdfs(v, w, 1);
                    sum_list_size += max_depth; // 累加 v 这一侧的最长链的长度
                    max_depth = 0;
                    rdfs(w, v, 1);
                    sum_list_size += max_depth; // 累加 w 这一侧的最长链的长度
                } else {
                    max_ring_size = max(max_ring_size, sz); // 取所有基环的最大值
                }
            }
        }
        return max(max_ring_size, sum_list_size);
    }
};
```

```python [sol1-Python]
class Solution:
    def maximumInvitations(self, favorite: List[int]) -> int:
        n = len(favorite)
        g = [[] for _ in range(n)]
        rg = [[] for _ in range(n)]  # 图 g 的反图
        deg = [0] * n  # 图 g 上每个节点的入度
        for v, w in enumerate(favorite):
            g[v].append(w)
            rg[w].append(v)
            deg[w] += 1

        # 拓扑排序，剪掉图 g 上的所有树枝
        q = deque(i for i, d in enumerate(deg) if d == 0)
        while q:
            v = q.popleft()
            for w in g[v]:
                deg[w] -= 1
                if deg[w] == 0:
                    q.append(w)

        # 寻找图 g 上的基环
        ring = []
        vis = [False] * n
        def dfs(v: int):
            vis[v] = True
            ring.append(v)
            for w in g[v]:
                if not vis[w]:
                    dfs(w)

        # 通过反图 rg 寻找树枝上最深的链
        max_depth = 0
        def rdfs(v: int, fa: int, depth: int):
            nonlocal max_depth
            max_depth = max(max_depth, depth)
            for w in rg[v]:
                if w != fa:
                    rdfs(w, v, depth + 1)

        max_ring_size, sum_list_size = 0, 0
        for i, b in enumerate(vis):
            if not b and deg[i]:  # 遍历基环上的点（拓扑排序后入度不为 0）
                ring = []
                dfs(i)
                if len(ring) == 2:  # 基环大小为 2
                    v, w = ring
                    max_depth = 0
                    rdfs(v, w, 1)
                    sum_list_size += max_depth  # 累加 v 这一侧的最长链的长度
                    max_depth = 0
                    rdfs(w, v, 1)
                    sum_list_size += max_depth  # 累加 w 这一侧的最长链的长度
                else:
                    max_ring_size = max(max_ring_size, len(ring))  # 取所有基环的最大值
                    
        return max(max_ring_size, sum_list_size)
```
