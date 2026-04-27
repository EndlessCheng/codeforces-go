本题 $n\le 13$，我们可以枚举节点集合 $U = \{0,1,2,...,n-1\}$ 的所有非空子集 $S$，这只有 $2^n-1\le 8191$ 个。

对于每个 $S$，如果其节点值总和是偶数，且连通，那么答案增加一。

如何判断 $S$ 是否连通？随便选一个在 $S$ 中的节点，作为 DFS 的起点。在 DFS 这张图的过程中，只访问在 $S$ 中的节点。DFS 结束后，如果访问过的节点集合恰好等于 $S$，说明 $S$ 是连通的。

代码实现时：

1. 由于节点值只有 $0$ 和 $1$，节点值之和是偶数，等价于节点值是异或和为 $0$。
2. 用二进制表示集合，用位运算实现集合操作，具体请看 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

[本题视频讲解](https://www.bilibili.com/video/BV15pZcBzEmR/?t=4m47s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def evenSumSubgraphs(self, nums: list[int], edges: list[list[int]]) -> int:
        n = len(nums)
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        # 枚举节点集合 U = {0,1,2,...,n-1} 的非空子集 sub
        u = (1 << n) - 1
        ans = 0

        for sub in range(1, u + 1):
            # 计算子图的点权异或和
            xor_sum = 0
            for i, x in enumerate(nums):
                if sub >> i & 1:  # i 在 sub 中
                    xor_sum ^= x
            if xor_sum:
                continue

            def dfs(x: int) -> None:
                nonlocal vis
                vis |= 1 << x  # 标记 x 已访问
                for y in g[x]:
                    if (vis >> y & 1) == 0:  # y 没有访问过
                        dfs(y)

            # 判断子图是否连通
            vis = u ^ sub  # 技巧：把不在子图中的节点都标记为已访问
            dfs(sub.bit_length() - 1)  # 随便选一个在子图中的节点，开始 DFS
            if vis == u:  # 所有节点都已访问，子图是连通的
                ans += 1

        return ans
```

```java [sol-Java]
class Solution {
    private int vis;

    public int evenSumSubgraphs(int[] nums, int[][] edges) {
        int n = nums.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        // 枚举节点集合 U = {0,1,2,...,n-1} 的非空子集 sub
        int u = (1 << n) - 1;
        int ans = 0;
        for (int sub = 1; sub <= u; sub++) {
            // 计算子图的点权异或和
            int xor = 0;
            for (int i = 0; i < n; i++) {
                if ((sub >> i & 1) > 0) { // i 在 sub 中
                    xor ^= nums[i];
                }
            }
            if (xor != 0) {
                continue;
            }

            // 判断子图是否连通
            vis = u ^ sub; // 技巧：把不在子图中的节点都标记为已访问
            dfs(Integer.numberOfTrailingZeros(sub), g);
            if (vis == u) { // 所有节点都已访问，子图是连通的
                ans++;
            }
        }
        return ans;
    }

    private void dfs(int x, List<Integer>[] g) {
        vis |= 1 << x; // 标记 x 已访问
        for (int y : g[x]) {
            if ((vis >> y & 1) == 0) { // y 没有访问过
                dfs(y, g);
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int evenSumSubgraphs(vector<int>& nums, vector<vector<int>>& edges) {
        int n = nums.size();
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        // 枚举节点集合 U = {0,1,2,...,n-1} 的非空子集 sub
        int u = (1 << n) - 1;
        int ans = 0;
        for (int sub = 1; sub <= u; sub++) {
            // 计算子图的点权异或和
            int xor_sum = 0;
            for (int i = 0; i < n; i++) {
                if (sub >> i & 1) { // i 在 sub 中
                    xor_sum ^= nums[i];
                }
            }
            if (xor_sum) {
                continue;
            }

            // 判断子图是否连通
            int vis = u ^ sub; // 技巧：把不在子图中的节点都标记为已访问
            auto dfs = [&](this auto&& dfs, int x) -> void {
                vis |= 1 << x; // 标记 x 已访问
                for (int y : g[x]) {
                    if ((vis >> y & 1) == 0) { // y 没有访问过
                        dfs(y);
                    }
                }
            };
            dfs(countr_zero((uint32_t) sub)); // 随便选一个在子图中的节点，开始 DFS
            ans += vis == u; // 所有节点都已访问，子图是连通的
        }
        return ans;
    }
};
```

```go [sol-Go]
func evenSumSubgraphs(nums []int, edges [][]int) (ans int) {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// 枚举节点集合 U = {0,1,2,...,n-1} 的非空子集 sub
	u := 1<<n - 1
	for sub := 1; sub <= u; sub++ {
		// 计算子图的点权异或和
		xor := 0
		for i, x := range nums {
			if sub>>i&1 > 0 { // i 在 sub 中
				xor ^= x
			}
		}
		if xor != 0 {
			continue
		}

		// 判断子图是否连通
		vis := u ^ sub // 技巧：把不在子图中的节点都标记为已访问
		var dfs func(int)
		dfs = func(x int) {
			vis |= 1 << x // 标记 x 已访问
			for _, y := range g[x] {
				if vis>>y&1 == 0 { // y 没有访问过
					dfs(y)
				}
			}
		}
		dfs(bits.TrailingZeros(uint(sub))) // 随便选一个在子图中的节点，开始 DFS

		if vis == u { // 所有节点都已访问，子图是连通的
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(2^n(n+m))$，其中 $n$ 是 $\textit{nums}$ 的长度，$m$ 是 $\textit{edges}$ 的长度。每次 DFS 需要 $\mathcal{O}(n+m)$ 的时间。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 优化：位运算 BFS

1. 既然 $\textit{nums}$ 只有 $0$ 和 $1$，可以将其压缩成一个二进制数 $\textit{ones}$，这样可以 $\mathcal{O}(1)$ 计算子集的点权和。
2. 用二进制数保存 $g$ 的邻居节点。
3. 把 DFS 换成 BFS，用一个二进制数 $q$ 代替队列，表示在队列中的节点。

> **注**：严格来说这不是 BFS，只是遍历图的一种方法。

```py [sol-Python3]
class Solution:
    def evenSumSubgraphs(self, nums: list[int], edges: list[list[int]]) -> int:
        n = len(nums)
        g = [0] * n
        for x, y in edges:
            g[x] |= 1 << y
            g[y] |= 1 << x

        ones = 0
        for i, x in enumerate(nums):
            ones |= x << i

        # 枚举节点集合 U = {0,1,2,...,n-1} 的非空子集 sub
        u = (1 << n) - 1
        ans = 0
        for sub in range(1, u + 1):
            # 计算子图的点权和
            s = (sub & ones).bit_count()
            if s % 2:
                continue

            # 判断子图是否连通
            vis = u ^ sub  # 技巧：把不在子图中的节点都标记为已访问
            q = sub & -sub  # 随便选一个在子图中的节点，开始 BFS
            vis |= q
            while q > 0:
                x = q & -q  # 出队
                q ^= x
                to = g[x.bit_length() - 1] & ~vis  # 访问 x 的（尚未访问过的）邻居
                q |= to  # x 的邻居入队
                vis |= to
            if vis == u:  # 所有节点都已访问，子图是连通的
                ans += 1

        return ans
```

```java [sol-Java]
class Solution {
    public int evenSumSubgraphs(int[] nums, int[][] edges) {
        int n = nums.length;
        int[] g = new int[n];
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x] |= 1 << y;
            g[y] |= 1 << x;
        }

        int ones = 0;
        for (int i = 0; i < nums.length; i++) {
            ones |= nums[i] << i;
        }

        // 枚举节点集合 U = {0,1,2,...,n-1} 的非空子集 sub
        int u = (1 << n) - 1;
        int ans = 0;
        for (int sub = 1; sub <= u; sub++) {
            // 计算子图的点权和
            int sum = Integer.bitCount(sub & ones);
            if (sum % 2 != 0) {
                continue;
            }

            // 判断子图是否连通
            int vis = u ^ sub; // 技巧：把不在子图中的节点都标记为已访问
            int q = sub & -sub; // 随便选一个在子图中的节点，开始 BFS
            vis |= q;
            while (q > 0) {
                int x = q & -q; // 出队
                q ^= x;
                int to = g[Integer.numberOfTrailingZeros(x)] & ~vis; // 访问 x 的（尚未访问过的）邻居
                q |= to; // x 的邻居入队
                vis |= to;
            }
            if (vis == u) { // 所有节点都已访问，子图是连通的
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int evenSumSubgraphs(vector<int>& nums, vector<vector<int>>& edges) {
        int n = nums.size();
        vector<int> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x] |= 1 << y;
            g[y] |= 1 << x;
        }

        int ones = 0;
        for (int i = 0; i < nums.size(); i++) {
            ones |= nums[i] << i;
        }

        // 枚举节点集合 U = {0,1,2,...,n-1} 的非空子集 sub
        int u = (1 << n) - 1;
        int ans = 0;
        for (int sub = 1; sub <= u; sub++) {
            // 计算子图的点权和
            int sum = popcount((uint32_t) sub & ones);
            if (sum % 2) {
                continue;
            }

            // 判断子图是否连通
            int vis = u ^ sub; // 技巧：把不在子图中的节点都标记为已访问
            int q = sub & -sub; // 随便选一个在子图中的节点，开始 BFS
            vis |= q;
            while (q > 0) {
                int x = q & -q; // 出队
                q ^= x;
                int to = g[countr_zero((uint32_t) x)] & ~vis; // 访问 x 的（尚未访问过的）邻居
                q |= to; // x 的邻居入队
                vis |= to;
            }
            ans += vis == u; // 所有节点都已访问，子图是连通的
        }
        return ans;
    }
};
```

```go [sol-Go]
func evenSumSubgraphs(nums []int, edges [][]int) (ans int) {
	n := len(nums)
	g := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] |= 1 << y
		g[y] |= 1 << x
	}

	ones := 0
	for i, x := range nums {
		ones |= x << i
	}

	// 枚举节点集合 U = {0,1,2,...,n-1} 的非空子集 sub
	u := 1<<n - 1
	for sub := 1; sub <= u; sub++ {
		// 计算子图的点权和
		sum := bits.OnesCount(uint(sub & ones))
		if sum%2 != 0 {
			continue
		}

		// 判断子图是否连通
		vis := u ^ sub // 技巧：把不在子图中的节点都标记为已访问
		q := sub & -sub // 随便选一个在子图中的节点，开始 BFS
		vis |= q
		for q > 0 {
			x := q & -q // 出队
			q ^= x
			to := g[bits.TrailingZeros(uint(x))] &^ vis // 访问 x 的（尚未访问过的）邻居
			q |= to // x 的邻居入队
			vis |= to
		}

		if vis == u { // 所有节点都已访问，子图是连通的
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m + n2^n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$m$ 是 $\textit{edges}$ 的长度。每次 BFS 至多出队 $n$ 个点，需要 $\mathcal{O}(n)$ 的时间。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

1. 回溯题单的「**§4.2 子集型回溯**」。
2. 图论题单的「**§1.1 深度优先搜索（DFS）**」。
3. 数据结构题单的「**七、并查集**」。

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
