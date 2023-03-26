### 本题视频讲解

见[【周赛 338】](https://www.bilibili.com/video/BV11o4y1p7Ci/?t=25m04s) 25:04。

### 前置知识：拓扑排序

见[【周赛 308】](https://www.bilibili.com/video/BV1mG411V7fj/?t=21m55s) 21:55。

### 提示 1

去掉不包含金币的子树，访问其中任何一个点都毫无意义。

做法：从没有金币的叶子出发，跑拓扑排序。

注意，去掉这些子树后，某些原来不是叶子的节点会变成叶子。

### 提示 2

只需要考虑有金币的叶子，因为不在叶子上的金币**顺路**就能收集到。

### 提示 3

从有金币的叶子出发，再次跑拓扑排序。在拓扑排序的同时，标记每个点入队的时间 $\textit{time}$。

**注意是入队的时间，不是访问到这个节点的时间。**

- 叶子入队的时间为 $0$；
- 去掉这些叶子后，又产生了**新的叶子**，这些叶子入队的时间为 $1$；
- 去掉这些叶子后，又产生了**新的叶子**，这些叶子入队的时间为 $2$；
- ……

示例 2 如下图，数字表示节点入队的时间：

![t4.png](https://pic.leetcode.cn/1679802238-QZehnH-t4.png)

那么只要走到 $\textit{time}[x]=2$ 的节点 $x$，就能收集到在叶子上的金币。

遍历所有边 $x-y$，如果满足 $\textit{time}[x]\ge 2$ 且 $\textit{time}[y]\ge 2$（上图蓝色边），那么这条边需要恰好经过 $2$ 次（因为需要回到出发点），答案加 $2$；如果不满足，则无需经过。

> 注：从任意被蓝色边连接的点出发，算出来的答案都是一样的。

```py [sol1-Python3]
class Solution:
    def collectTheCoins(self, coins: List[int], edges: List[List[int]]) -> int:
        n = len(coins)
        g = [[] for _ in range(n)]
        deg = [0] * n
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)  # 建图
            deg[x] += 1
            deg[y] += 1

        # 用拓扑排序「剪枝」：去掉没有金币的子树
        q = deque()
        for i, (d, c) in enumerate(zip(deg, coins)):
            if d == 1 and c == 0:  # 无金币叶子
                q.append(i)
        while q:
            for y in g[q.popleft()]:
                deg[y] -= 1
                if deg[y] == 1 and coins[y] == 0:
                    q.append(y)

        # 再次拓扑排序
        for i, (d, c) in enumerate(zip(deg, coins)):
            if d == 1 and c:  # 有金币叶子
                q.append(i)
        if len(q) <= 1:  # 至多一个有金币的叶子，直接收集
            return 0
        time = [0] * n
        while q:
            x = q.popleft()
            for y in g[x]:
                deg[y] -= 1
                if deg[y] == 1:
                    time[y] = time[x] + 1  # 记录入队时间
                    q.append(y)

        # 统计答案
        return sum(time[x] >= 2 and time[y] >= 2 for x, y in edges) * 2
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
            ++deg[x];
            ++deg[y];
        }

        // 用拓扑排序「剪枝」：去掉没有金币的子树
        var q = new ArrayDeque<Integer>();
        for (int i = 0; i < n; ++i)
            if (deg[i] == 1 && coins[i] == 0) // 无金币叶子
                q.add(i);
        while (!q.isEmpty()) {
            int x = q.peek();
            q.pop();
            for (int y : g[x])
                if (--deg[y] == 1 && coins[y] == 0)
                    q.add(y);
        }

        // 再次拓扑排序
        for (int i = 0; i < n; ++i)
            if (deg[i] == 1 && coins[i] == 1) // 有金币叶子
                q.add(i);
        if (q.size() <= 1) return 0; // 至多一个有金币的叶子，直接收集
        var time = new int[n];
        while (!q.isEmpty()) {
            int x = q.peek();
            q.pop();
            for (int y : g[x])
                if (--deg[y] == 1) {
                    time[y] = time[x] + 1; // 记录入队时间
                    q.add(y);
                }
        }

        // 统计答案
        int ans = 0;
        for (var e : edges)
            if (time[e[0]] >= 2 && time[e[1]] >= 2)
                ans += 2;
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int collectTheCoins(vector<int> &coins, vector<vector<int>> &edges) {
        int n = coins.size();
        vector<vector<int>> g(n);
        int deg[n]; memset(deg, 0, sizeof(deg));
        for (auto &e: edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x); // 建图
            ++deg[x];
            ++deg[y];
        }

        // 用拓扑排序「剪枝」：去掉没有金币的子树
        queue<int> q;
        for (int i = 0; i < n; ++i)
            if (deg[i] == 1 && coins[i] == 0) // 无金币叶子
                q.push(i);
        while (!q.empty()) {
            int x = q.front();
            q.pop();
            for (int y: g[x])
                if (--deg[y] == 1 && coins[y] == 0)
                    q.push(y);
        }

        // 再次拓扑排序
        for (int i = 0; i < n; ++i)
            if (deg[i] == 1 && coins[i]) // 有金币叶子
                q.push(i);
        if (q.size() <= 1) return 0; // 至多一个有金币的叶子，直接收集
        int time[n]; memset(time, 0, sizeof(time));
        while (!q.empty()) {
            int x = q.front();
            q.pop();
            for (int y: g[x])
                if (--deg[y] == 1) {
                    time[y] = time[x] + 1; // 记录入队时间
                    q.push(y);
                }
        }

        // 统计答案
        int ans = 0;
        for (auto &e: edges)
            if (time[e[0]] >= 2 && time[e[1]] >= 2)
                ans += 2;
        return ans;
    }
};
```

```go [sol1-Go]
func collectTheCoins(coins []int, edges [][]int) (ans int) {
	n := len(coins)
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建图
		deg[x]++
		deg[y]++
	}

	// 用拓扑排序「剪枝」：去掉没有金币的子树
	q := make([]int, 0, n)
	for i, d := range deg {
		if d == 1 && coins[i] == 0 { // 无金币叶子
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range g[x] {
			deg[y]--
			if deg[y] == 1 && coins[y] == 0 {
				q = append(q, y)
			}
		}
	}

	// 再次拓扑排序
	for i, d := range deg {
		if d == 1 && coins[i] == 1 { // 有金币叶子
			q = append(q, i)
		}
	}
	if len(q) <= 1 { // 至多一个有金币的叶子，直接收集
		return
	}
	time := make([]int, n)
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range g[x] {
			deg[y]--
			if deg[y] == 1 {
				time[y] = time[x] + 1 // 记录入队时间
				q = append(q, y)
			}
		}
	}

	// 统计答案
	for _, e := range edges {
		if time[e[0]] >= 2 && time[e[1]] >= 2 {
			ans += 2
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{coins}$ 的长度。
- 空间复杂度：$O(n)$。

### 思考题

如果把题目中的 $2$ 换成 $0,1,2,3,\cdots, n-1$，你能把这些情况对应的答案全部算出来吗？要求总的时间复杂度仍然是 $O(n)$。

做法见本题视频讲解。

### 相似题目

- [2050. 并行课程 III](https://leetcode.cn/problems/parallel-courses-iii/)
- [1857. 有向图中最大颜色值](https://leetcode.cn/problems/largest-color-value-in-a-directed-graph/)
