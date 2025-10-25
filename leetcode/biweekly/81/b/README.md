建图后，用 DFS 可以求出每个连通块的大小。

求连通块的大小的同时，用一个变量 $\textit{total}$ 维护前面求出的连通块的大小之和。

设当前连通块的大小为 $\textit{size}$，那么这个连通块中的每个点，与前面遍历过的连通块的每个点，都是无法互相到达的，根据乘法原理，这有 $\textit{size}\cdot\textit{total}$ 个，加到答案中。

```py [sol-Python3]
class Solution:
    def countPairs(self, n: int, edges: List[List[int]]) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)  # 建图

        vis = [False] * n
        def dfs(x: int) -> int:
            vis[x] = True  # 避免重复访问同一个点
            size = 1
            for y in g[x]:
                if not vis[y]:
                    size += dfs(y)
            return size

        ans = total = 0
        for i in range(n):
            if not vis[i]:  # 未访问的点：说明找到了一个新的连通块
                size = dfs(i)
                ans += size * total
                total += size
        return ans
```

```java [sol-Java]
class Solution {
    public long countPairs(int n, int[][] edges) {
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x); // 建图
        }

        boolean[] vis = new boolean[n];
        long ans = 0;
        for (int i = 0, total = 0; i < n; i++) {
            if (!vis[i]) { // 未访问的点：说明找到了一个新的连通块
                int size = dfs(i, g, vis);
                ans += (long) size * total;
                total += size;
            }
        }
        return ans;
    }

    private int dfs(int x, List<Integer>[] g, boolean[] vis) {
        vis[x] = true; // 避免重复访问同一个点
        int size = 1;
        for (int y : g[x]) {
            if (!vis[y]) {
                size += dfs(y, g, vis);
            }
        }
        return size;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countPairs(int n, vector<vector<int>>& edges) {
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x); // 建图
        }

        vector<int8_t> vis(n);
        auto dfs = [&](this auto&& dfs, int x) -> int {
            vis[x] = true; // 避免重复访问同一个点
            int size = 1;
            for (int y: g[x]) {
                if (!vis[y]) {
                    size += dfs(y);
                }
            }
            return size;
        };

        long long ans = 0;
        for (int i = 0, total = 0; i < n; i++) {
            if (!vis[i]) { // 未访问的点：说明找到了一个新的连通块
                int size = dfs(i);
                ans += (long long) size * total;
                total += size;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func countPairs(n int, edges [][]int) (ans int64) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建图
	}

	vis := make([]bool, n)
	var dfs func(int) int
	dfs = func(x int) int {
		vis[x] = true // 避免重复访问同一个点
		size := 1
		for _, y := range g[x] {
			if !vis[y] {
				size += dfs(y)
			}
		}
		return size
	}

	total := 0
	for i, b := range vis {
		if !b { // 未访问的点：说明找到了一个新的连通块
			size := dfs(i)
			ans += int64(size) * int64(total)
			total += size
		}
	}
	return
}
```

```js [sol-JavaScript]
var countPairs = function(n, edges) {
    const g = new Array(n).fill(null).map(() => []);
    for (const [x, y] of edges) {
        g[x].push(y);
        g[y].push(x); // 建图
    }

    const vis = new Array(n).fill(false);
    function dfs(x) {
        vis[x] = true; // 避免重复访问同一个点
        let size = 1;
        for (let y of g[x]) {
            if (!vis[y]) {
                size += dfs(y);
            }
        }
        return size;
    }

    let ans = 0, total = 0;
    for (let i = 0; i < n; i++) {
        if (!vis[i]) { // 未访问的点：说明找到了一个新的连通块
            const size = dfs(i);
            ans += size * total;
            total += size;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_pairs(n: i32, edges: Vec<Vec<i32>>) -> i64 {
        let n = n as usize;
        let mut g = vec![vec![]; n];
        for e in &edges {
            let x = e[0] as usize;
            let y = e[1] as usize;
            g[x].push(y);
            g[y].push(x); // 建图
        }

        fn dfs(x: usize, g: &Vec<Vec<usize>>, vis: &mut Vec<bool>) -> i32 {
            vis[x] = true; // 避免重复访问同一个点
            let mut size = 1;
            for &y in &g[x] {
                if !vis[y] {
                    size += dfs(y, g, vis);
                }
            }
            size
        }

        let mut ans = 0i64;
        let mut total = 0;
        let mut vis = vec![false; n];
        for i in 0..n {
            if !vis[i] { // 未访问的点：说明找到了一个新的连通块
                let size = dfs(i, &g, &mut vis);
                ans += size as i64 * total as i64;
                total += size;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $m$ 为 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
