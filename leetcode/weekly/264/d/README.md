定义 $f[i]$ 表示完成第 $i$ 门课程需要花费的最少月份数。根据题意，只有当 $i$ 的所有先修课程都完成时，才可以开始学习第 $i$ 门课程，并且可以立即开始。

因此 

$$
f[i]=\textit{time}[i] + \max_{j} f[j]
$$

其中 $j$ 是 $i$ 的先修课程。

由于题目保证图是一个有向无环图，所以一定存在拓扑序。我们可以在计算拓扑序的同时，计算状态转移。

具体来说，设当前节点为 $x$，我们可以在计算出 $f[x]$ 后，更新 $y$ 的所有先修课程耗时的最大值，这里 $x$ 是 $y$ 的先修课程。

答案就是所有 $f[i]$ 的最大值。

```py [sol-Python3]
class Solution:
    def minimumTime(self, n: int, relations: List[List[int]], time: List[int]) -> int:
        g = [[] for _ in range(n)]
        deg = [0] * n  # deg[i] 表示 i 的先修课的个数
        for x, y in relations:
            g[x - 1].append(y - 1)  # 建图
            deg[y - 1] += 1

        q = deque(i for i, d in enumerate(deg) if d == 0)  # 没有先修课
        f = [0] * n
        while q:
            x = q.popleft()  # x 出队，意味着 x 的所有先修课都上完了
            f[x] += time[x]  # 加上当前课程的时间，就得到了最终的 f[x]
            for y in g[x]:  # 遍历 x 的邻居 y
                f[y] = max(f[y], f[x])  # 更新 f[y] 的所有先修课程耗时的最大值
                deg[y] -= 1
                if deg[y] == 0:  # y 的先修课已上完
                    q.append(y)
        return max(f)
```

```java [sol-Java]
class Solution {
    public int minimumTime(int n, int[][] relations, int[] time) {
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        var deg = new int[n];
        for (var r : relations) {
            int x = r[0] - 1, y = r[1] - 1;
            g[x].add(y);
            deg[y]++;
        }

        var q = new ArrayDeque<Integer>();
        for (int i = 0; i < n; i++)
            if (deg[i] == 0) // 没有先修课
                q.add(i);
        var f = new int[n];
        int ans = 0;
        while (!q.isEmpty()) {
            int x = q.poll(); // x 出队，意味着 x 的所有先修课都上完了
            f[x] += time[x]; // 加上当前课程的时间，就得到了最终的 f[x]
            ans = Math.max(ans, f[x]);
            for (int y : g[x]) { // 遍历 x 的邻居 y
                f[y] = Math.max(f[y], f[x]); // 更新 f[y] 的所有先修课程耗时的最大值
                if (--deg[y] == 0) // y 的先修课已上完
                    q.add(y);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumTime(int n, vector<vector<int>> &relations, vector<int> &time) {
        vector<vector<int>> g(n);
        vector<int> deg(n);
        for (auto &r: relations) {
            int x = r[0] - 1, y = r[1] - 1;
            g[x].push_back(y); // 建图
            deg[y]++; // 可以理解为 y 的先修课的个数
        }

        queue<int> q;
        for (int i = 0; i < n; i++)
            if (deg[i] == 0) // 没有先修课
                q.push(i);
        vector<int> f(n);
        while (!q.empty()) {
            int x = q.front();
            q.pop();
            // x 出队，意味着 x 的所有先修课都上完了
            f[x] += time[x]; // 加上当前课程的时间，就得到了最终的 f[x]
            for (int y: g[x]) { // 遍历 x 的邻居 y
                f[y] = max(f[y], f[x]); // 更新 f[y] 的所有先修课程耗时的最大值
                if (--deg[y] == 0) // y 的先修课已上完
                    q.push(y);
            }
        }
        return *max_element(f.begin(), f.end());
    }
};
```

```go [sol-Go]
func minimumTime(n int, relations [][]int, time []int) (ans int) {
	g := make([][]int, n)
	deg := make([]int, n)
	for _, r := range relations {
		x, y := r[0]-1, r[1]-1
		g[x] = append(g[x], y) // 建图
		deg[y]++ // 可以理解为 y 的先修课的个数
	}

	q := []int{}
	for i, d := range deg {
		if d == 0 { // 没有先修课
			q = append(q, i)
		}
	}
	f := make([]int, n)
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		// x 出队，意味着 x 的所有先修课都上完了
		f[x] += time[x] // 加上当前课程的时间，就得到了最终的 f[x]
		ans = max(ans, f[x])
		for _, y := range g[x] { // 遍历 x 的邻居 y
			f[y] = max(f[y], f[x]) // 更新 f[y] 的所有先修课程耗时的最大值
			if deg[y]--; deg[y] == 0 { // y 的先修课已上完
				q = append(q, y)
			}
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

```js [sol-JavaScript]
var minimumTime = function (n, relations, time) {
    let g = Array(n).fill(null).map(() => []);
    let deg = Array(n).fill(0); // deg[i] 表示 i 的先修课的个数
    for (const [x, y] of relations) {
        g[x - 1].push(y - 1);
        deg[y - 1]++;
    }

    let q = new Queue();
    for (let i = 0; i < n; i++)
        if (deg[i] === 0) // 没有先修课
            q.enqueue(i);
    let f = Array(n).fill(0);
    while (!q.isEmpty()) {
        const x = q.dequeue(); // x 出队，意味着 x 的所有先修课都上完了
        f[x] += time[x]; // 加上当前课程的时间，就得到了最终的 f[x]
        for (const y of g[x]) { // 遍历 x 的邻居 y
            f[y] = Math.max(f[y], f[x]); // 更新 f[y] 的所有先修课程耗时的最大值
            if (--deg[y] === 0) // y 的先修课已上完
                q.enqueue(y);
        }
    }
    return Math.max(...f);
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $m$ 为 $\textit{relations}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 相似题目

- [1857. 有向图中最大颜色值](https://leetcode.cn/problems/largest-color-value-in-a-directed-graph/)

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
