## 提示 1

定义一个点的**度数**为其邻居个数。如果一个点的度数为 $1$，那么这个点叫做**叶子节点**，例如示例 2 的 $3,4,6,7$ 都是叶子节点。

如果叶子节点没有金币，我们有必要移动到叶子节点吗？没有必要。

那么可以先把这些没有金币的叶子节点去掉。如果去掉后又产生了新的没有金币的叶子节点，就继续去掉。

怎么实现？**拓扑排序**。一开始，把没有金币的叶子节点都加到队列中。然后不断循环直到队列为空。每次循环，弹出队首的节点 $x$，并删除 $x$ 及其邻居之间的边。我们并不需要实际删除边，只需要把邻居的度数减少 $1$。如果一个邻居的度数减少为 $1$ 且没有金币，就加到队列中，继续拓扑排序。

## 提示 2

看示例 2，在去掉节点 $6$ 之后，现在每个叶子节点上都有金币。

由于可以「收集距离当前节点距离为 $2$ 以内的所有金币」，我们没有必要移动到叶子节点再收集，而是移动到叶子节点的父节点的父节点，就能收集到叶子节点上的金币。

那么，去掉所有叶子，然后再去掉新产生的叶子，剩余节点就是必须要访问的节点。

## 提示 3

由于题目要求最后回到出发点，无论从哪个点出发，每条边都必须走两次。这是因为把出发点作为树根，递归遍历这棵树，那么往下「递」是一次，往上「归」又是一次，每条边都会经过两次。

所以答案就是剩余边数乘 $2$。当我们删除节点时，也可以看成是删除这个点到其父节点的边。

特别地，如果所有点都要被删除，那么当剩下两个点时，这两个点之间的边我们会删除两次，这会导致剩余边数等于 $-1$，而此时答案应该是 $0$。所以最后答案要和 $0$ 取最大值。

**代码实现时，由于我们不需要得到一个严格的拓扑序，所以简单地用栈或者数组代替队列，也是可以的。**

```py [sol1-Python3]
class Solution:
    def collectTheCoins(self, coins: List[int], edges: List[List[int]]) -> int:
        n = len(coins)
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)  # 建图
        deg = list(map(len, g))  # 每个节点的度数（邻居个数）

        left_edges = n - 1  # 剩余边数
        # 拓扑排序，去掉没有金币的子树
        q = []
        for i, (d, c) in enumerate(zip(deg, coins)):
            if d == 1 and c == 0:  # 没有金币的叶子
                q.append(i)
        while q:
            left_edges -= 1  # 删除节点到其父节点的边
            for y in g[q.pop()]:
                deg[y] -= 1
                if deg[y] == 1 and coins[y] == 0:  # 没有金币的叶子
                    q.append(y)

        # 再次拓扑排序
        for i, (d, c) in enumerate(zip(deg, coins)):
            if d == 1 and c:  # 有金币的叶子（判断 c 是避免把没有金币的叶子也算进来）
                q.append(i)
        left_edges -= len(q)  # 删除所有叶子（到其父节点的边）
        for x in q:  # 遍历所有叶子
            for y in g[x]:
                deg[y] -= 1
                if deg[y] == 1:  # y 现在是叶子了
                    left_edges -= 1  # 删除 y（到其父节点的边）
        return max(left_edges * 2, 0)
```

```java [sol1-Java]
class Solution {
    public int collectTheCoins(int[] coins, int[][] edges) {
        int n = coins.length;
        List<Integer> g[] = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        var deg = new int[n];
        for (var e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x); // 建图
            deg[x]++;
            deg[y]++; // 统计每个节点的度数（邻居个数）
        }

        int leftEdges = n - 1; // 剩余边数
        // 拓扑排序，去掉没有金币的子树
        var q = new ArrayDeque<Integer>();
        for (int i = 0; i < n; i++) {
            if (deg[i] == 1 && coins[i] == 0) { // 没有金币的叶子
                q.add(i);
            }
        }
        while (!q.isEmpty()) {
            leftEdges--; // 删除节点到其父节点的边
            for (int y : g[q.poll()]) {
                if (--deg[y] == 1 && coins[y] == 0) { // 没有金币的叶子
                    q.add(y);
                }
            }
        }

        // 再次拓扑排序
        for (int i = 0; i < n; i++) {
            if (deg[i] == 1 && coins[i] == 1) { // 有金币的叶子（判断 coins[i] 是避免把没有金币的叶子也算进来）
                q.add(i);
            }
        }
        leftEdges -= q.size(); // 删除所有叶子（到其父节点的边）
        for (int x : q) { // 遍历所有叶子
            for (int y : g[x]) {
                if (--deg[y] == 1) { // y 现在是叶子了
                    leftEdges--; // 删除 y（到其父节点的边）
                }
            }
        }
        return Math.max(leftEdges * 2, 0);
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int collectTheCoins(vector<int> &coins, vector<vector<int>> &edges) {
        int n = coins.size();
        vector<vector<int>> g(n);
        vector<int> deg(n);
        for (auto &e: edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x); // 建图
            deg[x]++;
            deg[y]++; // 统计每个节点的度数（邻居个数）
        }

        int left_edges = n - 1; // 剩余边数
        // 拓扑排序，去掉没有金币的子树
        vector<int> q;
        for (int i = 0; i < n; i++)
            if (deg[i] == 1 && coins[i] == 0) // 没有金币的叶子
                q.push_back(i);
        while (!q.empty()) {
            left_edges--; // 删除节点 x（到其父节点的边）
            int x = q.back(); q.pop_back();
            for (int y: g[x])
                if (--deg[y] == 1 && coins[y] == 0) // 没有金币的叶子
                    q.push_back(y);
        }

        // 再次拓扑排序
        for (int i = 0; i < n; i++)
            if (deg[i] == 1 && coins[i]) // 有金币的叶子（判断 coins[i] 是避免把没有金币的叶子也算进来）
                q.push_back(i);
        left_edges -= q.size(); // 删除所有叶子（到其父节点的边）
        for (int x: q) // 遍历所有叶子
            for (int y: g[x])
                if (--deg[y] == 1) // y 现在是叶子了
                    left_edges--; // 删除 y（到其父节点的边）
        return max(left_edges * 2, 0);
    }
};
```

```go [sol1-Go]
func collectTheCoins(coins []int, edges [][]int) int {
	n := len(coins)
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建图
		deg[x]++
		deg[y]++ // 统计每个节点的度数（邻居个数）
	}

	leftEdges := n - 1 // 剩余边数
	// 拓扑排序，去掉没有金币的子树
	q := []int{}
	for i, d := range deg {
		if d == 1 && coins[i] == 0 { // 没有金币的叶子
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		x := q[len(q)-1]
		q = q[:len(q)-1]
		leftEdges-- // 删除节点 x 到其父节点的边
		for _, y := range g[x] {
			deg[y]--
			if deg[y] == 1 && coins[y] == 0 { // 没有金币的叶子
				q = append(q, y)
			}
		}
	}

	// 再次拓扑排序
	for i, d := range deg {
		if d == 1 && coins[i] > 0 { // 有金币的叶子（判断 coins[i] 是避免把没有金币的叶子也算进来）
			q = append(q, i)
		}
	}
	leftEdges -= len(q) // 删除所有叶子（到其父节点的边）
	for _, x := range q { // 遍历所有叶子
		for _, y := range g[x] {
			deg[y]--
			if deg[y] == 1 { // y 现在是叶子了
				leftEdges-- // 删除 y（到其父节点的边）
			}
		}
	}
	return max(leftEdges*2, 0)
}

func max(a, b int) int { if b > a { return b }; return a }
```

```js [sol1-JavaScript]
var collectTheCoins = function (coins, edges) {
    const n = coins.length;
    const g = Array(n).fill(null).map(() => []);
    for (const [x, y] of edges) {
        g[x].push(y);
        g[y].push(x); // 建图
    }
    const deg = g.map((neighbors) => neighbors.length); // 每个节点的度数（邻居个数）

    let leftEdges = n - 1; // 剩余边数
    // 拓扑排序，去掉没有金币的子树
    const q = [];
    for (let i = 0; i < n; i++) {
        if (deg[i] === 1 && coins[i] === 0) { // 没有金币的叶子
            q.push(i);
        }
    }
    while (q.length) {
        leftEdges--; // 删除节点到其父节点的边
        for (const x of g[q.pop()]) {
            if (--deg[x] === 1 && coins[x] === 0) { // 没有金币的叶子
                q.push(x);
            }
        }
    }

    // 再次拓扑排序
    for (let i = 0; i < n; i++) {
        if (deg[i] === 1 && coins[i]) { // 有金币的叶子（判断 coins[i] 是避免把没有金币的叶子也算进来）
            q.push(i);
        }
    }
    leftEdges -= q.length; // 删除所有叶子（到其父节点的边）
    for (const x of q) { // 遍历所有叶子
        for (const y of g[x]) {
            if (--deg[y] === 1) { // y 现在是叶子了
                leftEdges--; // 删除 y（到其父节点的边）
            }
        }
    }
    return Math.max(leftEdges * 2, 0);
};
```

```rust [sol1-Rust]
impl Solution {
    pub fn collect_the_coins(coins: Vec<i32>, edges: Vec<Vec<i32>>) -> i32 {
        let n = coins.len();
        let mut g = vec![vec![]; n];
        let mut deg = vec![0; n];
        for e in &edges {
            let x = e[0] as usize;
            let y = e[1] as usize;
            g[x].push(y);
            g[y].push(x); // 建图
            deg[x] += 1;
            deg[y] += 1; // 统计每个节点的度数（邻居个数）
        }

        let mut left_edges = n as i32 - 1; // 剩余边数
        // 拓扑排序，去掉没有金币的子树
        let mut q = Vec::new();
        for i in 0..n {
            if deg[i] == 1 && coins[i] == 0 { // 没有金币的叶子
                q.push(i);
            }
        }
        while !q.is_empty() {
            left_edges -= 1; // 删除节点到其父节点的边
            for &y in &g[q.pop().unwrap()] {
                deg[y] -= 1;
                if deg[y] == 1 && coins[y] == 0 { // 没有金币的叶子
                    q.push(y);
                }
            }
        }

        // 再次拓扑排序
        for i in 0..n {
            if deg[i] == 1 && coins[i] == 1 { // 有金币的叶子（判断 coins[i] 是避免把没有金币的叶子也算进来）
                q.push(i);
            }
        }
        left_edges -= q.len() as i32; // 删除所有叶子（到其父节点的边）
        for &x in &q { // 遍历所有叶子
            for &y in &g[x] {
                deg[y] -= 1;
                if deg[y] == 1 { // y 现在是叶子了
                    left_edges -= 1; // 删除 y（到其父节点的边）
                }
            }
        }
        0.max(left_edges * 2)
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{coins}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

- [310. 最小高度树](https://leetcode.cn/problems/minimum-height-trees/)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

更多精彩题解，请看 [往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
