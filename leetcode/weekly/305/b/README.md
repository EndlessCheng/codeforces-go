树上 DFS（Python/Java/C++/Go/JS/Rust）

---
 
用哈希集合（或者布尔数组）记录哪些节点是受限的。

然后建树，只有当两个节点都不是受限的（都不在哈希集合中）才连边。

从 $0$ 出发 DFS 这棵树，统计能访问到的节点个数，即为答案。

```py [sol-Python3]
class Solution:
    def reachableNodes(self, n: int, edges: List[List[int]], restricted: List[int]) -> int:
        r = set(restricted)
        g = [[] for _ in range(n)]
        for x, y in edges:
            if x not in r and y not in r:
                g[x].append(y)  # 都不受限才连边
                g[y].append(x)
        def dfs(x: int, fa: int) -> int:
            cnt = 1
            for y in g[x]:
                if y != fa:
                    cnt += dfs(y, x)
            return cnt
        return dfs(0, -1)
```

```java [sol-Java]
class Solution {
    public int reachableNodes(int n, int[][] edges, int[] restricted) {
        boolean[] isRestricted = new boolean[n];
        for (int x : restricted) {
            isRestricted[x] = true;
        }
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1];
            if (!isRestricted[x] && !isRestricted[y]) {
                g[x].add(y); // 都不受限才连边
                g[y].add(x);
            }
        }
        return dfs(0, -1, g);
    }

    private int dfs(int x, int fa, List<Integer>[] g) {
        int cnt = 1;
        for (int y : g[x]) {
            if (y != fa) {
                cnt += dfs(y, x, g);
            }
        }
        return cnt;
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<vector<int>> g;

    int dfs(int x, int fa) {
        int cnt = 1;
        for (int y : g[x]) {
            if (y != fa) {
                cnt += dfs(y, x);
            }
        }
        return cnt;
    };

public:
    int reachableNodes(int n, vector<vector<int>> &edges, vector<int> &restricted) {
        unordered_set<int> r(restricted.begin(), restricted.end());
        g.resize(n);
        for (auto &e : edges) {
            int x = e[0], y = e[1];
            if (!r.contains(x) && !r.contains(y)) {
                g[x].push_back(y); // 都不受限才连边
                g[y].push_back(x);
            }
        }
        return dfs(0, -1);
    }
};
```

```go [sol-Go]
func reachableNodes(n int, edges [][]int, restricted []int) (ans int) {
	r := make(map[int]bool, len(restricted))
	for _, x := range restricted {
		r[x] = true
	}
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		if !r[x] && !r[y] { // 都不受限才连边
			g[x] = append(g[x], y)
			g[y] = append(g[y], x)
		}
	}
	var dfs func(int, int)
	dfs = func(x, fa int) {
		ans++
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)
	return
}
```

```js [sol-JavaScript]
var reachableNodes = function(n, edges, restricted) {
    const r = new Set(restricted);
    const g = Array(n).fill(null).map(() => []);
    for (const [x, y] of edges) {
        if (!r.has(x) && !r.has(y)) {
            g[x].push(y); // 都不受限才连边
            g[y].push(x);
        }
    }
    let ans = 0;
    function dfs(x, fa) {
        ans++;
        for (const y of g[x]) {
            if (y !== fa) {
                dfs(y, x);
            }
        }
    }
    dfs(0, -1);
    return ans;
};
```

```rust [sol-Rust]
use std::collections::HashSet;

impl Solution {
    pub fn reachable_nodes(n: i32, edges: Vec<Vec<i32>>, restricted: Vec<i32>) -> i32 {
        let r = restricted.into_iter().collect::<HashSet<_>>();
        let mut g = vec![vec![]; n as usize];
        for e in &edges {
            let x = e[0];
            let y = e[1];
            if !r.contains(&x) && !r.contains(&y) {
                g[x as usize].push(y as usize); // 都不受限才连边
                g[y as usize].push(x as usize);
            }
        }
        fn dfs(x: usize, fa: usize, g: &Vec<Vec<usize>>) -> i32 {
            let mut cnt = 1;
            for &y in &g[x] {
                if y != fa {
                    cnt += dfs(y, x, g);
                }
            }
            cnt
        }
        dfs(0, 0, &g)
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
