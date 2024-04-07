Discussion:

- If $s=t$, no movement is needed, the answer is $0$.
- If $s$ and $t$ are not in the same connected component, the answer is $-1$.
- If $s$ and $t$ are in the same connected component. Since the property of AND is that the more numbers ANDed, the smaller the result. With the condition of being able to repeat edges, the optimal solution is to traverse all edges within the connected component where $s$ lies.
So we need to know which connected component $s$ and $t$ belong to, and what the AND of edge weights is within the connected component.

This can be implemented using either DFS or Union Find.

## DFS

```py [sol-Python3]
class Solution:
    def minimumCost(self, n: int, edges: List[List[int]], query: List[List[int]]) -> List[int]:
        g = [[] for _ in range(n)]
        for x, y, w in edges:
            g[x].append((y, w))
            g[y].append((x, w))

        def dfs(x: int) -> int:
            and_ = -1
            ids[x] = len(cc_and)
            for y, w in g[x]:
                and_ &= w
                if ids[y] < 0:
                    and_ &= dfs(y)
            return and_

        ids = [-1] * n
        cc_and = []
        for i in range(n):
            if ids[i] < 0:
                cc_and.append(dfs(i))

        return [0 if s == t else -1 if ids[s] != ids[t] else cc_and[ids[s]]
                for s, t in query]
```

```java [sol-Java]
class Solution {
    public int[] minimumCost(int n, int[][] edges, int[][] query) {
        List<int[]>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1], w = e[2];
            g[x].add(new int[]{y, w});
            g[y].add(new int[]{x, w});
        }

        int[] ids = new int[n];
        Arrays.fill(ids, -1);
        List<Integer> ccAnd = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            if (ids[i] < 0) {
                ccAnd.add(dfs(i, ccAnd.size(), g, ids));
            }
        }

        int[] ans = new int[query.length];
        for (int i = 0; i < query.length; i++) {
            int s = query[i][0], t = query[i][1];
            ans[i] = s == t ? 0 : ids[s] != ids[t] ? -1 : ccAnd.get(ids[s]);
        }
        return ans;
    }

    private int dfs(int x, int curId, List<int[]>[] g, int[] ids) {
        ids[x] = curId;
        int and = -1;
        for (int[] e : g[x]) {
            and &= e[1];
            if (ids[e[0]] < 0) {
                and &= dfs(e[0], curId, g, ids);
            }
        }
        return and;
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<vector<pair<int, int>>> g;
    vector<int> cc_and, ids;

    int dfs(int x) {
        ids[x] = cc_and.size(); 
        int and_ = -1;
        for (auto &[y, w]: g[x]) {
            and_ &= w;
            if (ids[y] < 0) { 
                and_ &= dfs(y);
            }
        }
        return and_;
    }

public:
    vector<int> minimumCost(int n, vector<vector<int>> &edges, vector<vector<int>> &query) {
        g.resize(n);
        for (auto &e: edges) {
            int x = e[0], y = e[1], w = e[2];
            g[x].emplace_back(y, w);
            g[y].emplace_back(x, w);
        }

        ids.resize(n, -1); 
        for (int i = 0; i < n; i++) {
            if (ids[i] < 0) { 
                cc_and.push_back(dfs(i)); 
            }
        }

        vector<int> ans;
        ans.reserve(query.size());
        for (auto &q: query) {
            int s = q[0], t = q[1];
            ans.push_back(s == t ? 0 : ids[s] != ids[t] ? -1 : cc_and[ids[s]]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumCost(n int, edges, query [][]int) []int {
	type edge struct{ to, w int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, w})
		g[y] = append(g[y], edge{x, w})
	}

	ids := make([]int, n) 
	for i := range ids {
		ids[i] = -1
	}
	ccAnd := []int{} 
	var dfs func(int) int
	dfs = func(x int) int {
		ids[x] = len(ccAnd) 
		and := -1
		for _, e := range g[x] {
			and &= e.w
			if ids[e.to] < 0 {
				and &= dfs(e.to)
			}
		}
		return and
	}
	for i, id := range ids {
		if id < 0 { // 没有访问过
			ccAnd = append(ccAnd, dfs(i))
		}
	}

	ans := make([]int, len(query))
	for i, q := range query {
		s, t := q[0], q[1]
		if s == t {
			continue
		}
		if ids[s] != ids[t] {
			ans[i] = -1
		} else {
			ans[i] = ccAnd[ids[s]]
		}
	}
	return ans
}
```

#### Complexity Analysis

- **Time complexity**: $\mathcal{O}(n+m+q)$, where $m$ is the length of $\textit{edges}$, and $q$ is the length of $\textit{query}$.
- **Space complexity**: $\mathcal{O}(n+m)$. The return value is not included.

## Union Find

```py [sol-Python3]
class Solution:
    def minimumCost(self, n: int, edges: List[List[int]], query: List[List[int]]) -> List[int]:
        fa = list(range(n))
        and_ = [-1] * n

        def find(x: int) -> int:
            if fa[x] != x:
                fa[x] = find(fa[x])
            return fa[x]

        for x, y, w in edges:
            x = find(x)
            y = find(y)
            and_[y] &= w
            if x != y:
                and_[y] &= and_[x]
                fa[x] = y

        return [0 if s == t else -1 if find(s) != find(t) else and_[find(s)]
                for s, t in query]
```

```java [sol-Java]
class Solution {
    public int[] minimumCost(int n, int[][] edges, int[][] query) {
        int[] fa = new int[n];
        for (int i = 0; i < n; i++) {
            fa[i] = i;
        }
        int[] and_ = new int[n];
        Arrays.fill(and_, -1);

        for (int[] e : edges) {
            int x = find(e[0], fa);
            int y = find(e[1], fa);
            and_[y] &= e[2];
            if (x != y) {
                and_[y] &= and_[x];
                fa[x] = y;
            }
        }

        int[] ans = new int[query.length];
        for (int i = 0; i < query.length; i++) {
            int s = query[i][0], t = query[i][1];
            ans[i] = s == t ? 0 : find(s, fa) != find(t, fa) ? -1 : and_[find(s, fa)];
        }
        return ans;
    }

    private int find(int x, int[] fa) {
        if (fa[x] != x) {
            fa[x] = find(fa[x], fa);
        }
        return fa[x];
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<int> fa, and_;

    int find(int x) {
        if (fa[x] != x) {
            fa[x] = find(fa[x]);
        }
        return fa[x];
    };

public:
    vector<int> minimumCost(int n, vector<vector<int>> &edges, vector<vector<int>> &query) {
        fa.resize(n);
        iota(fa.begin(), fa.end(), 0);
        and_.resize(n, -1);
        for (auto &e: edges) {
            int x = find(e[0]);
            int y = find(e[1]);
            and_[y] &= e[2];
            if (x != y) {
                and_[y] &= and_[x];
                fa[x] = y;
            }
        }

        vector<int> ans;
        ans.reserve(query.size()); 
        for (auto &q: query) {
            int s = q[0], t = q[1];
            ans.push_back(s == t ? 0 : find(s) != find(t) ? -1 : and_[find(s)]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumCost(n int, edges, query [][]int) []int {
	fa := make([]int, n)
	and := make([]int, n)
	for i := range fa {
		fa[i] = i
		and[i] = -1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	for _, e := range edges {
		x, y := find(e[0]), find(e[1])
		and[y] &= e[2]
		if x != y {
			and[y] &= and[x]
			fa[x] = y
		}
	}

	ans := make([]int, len(query))
	for i, q := range query {
		s, t := q[0], q[1]
		if s == t {
			continue
		}
		if find(s) != find(t) {
			ans[i] = -1
		} else {
			ans[i] = and[find(s)]
		}
	}
	return ans
}
```

#### Complexity Analysis

- **Time complexity**: $\mathcal{O}((n+m+q)\log n)$, where $m$ is the length of $\textit{edges}$, and $q$ is the length of $\textit{query}$.
- **Space complexity**: $\mathcal{O}(n+m)$. The return value is not included.
