评论区有人给出了如下数据：

```
11
[[0,1],[1,2],[2,3],[5,6],[6,7]]
```

这个数据预期结果是 $366$，但实际上可以构造 $9\text{-}11\text{-}10\text{-}7$ 和 $5\text{-}8\text{-}6$ 两条链，算出结果是 $367$。

题目应该会改，可能的修改方向是把数据范围改小，变成一个子集状压题目？

下面的不用看了，题目修改后会重写题解。

---

## DFS

每个节点最多与其他两个节点相连，意味着图中的每个连通块要么是链，要么是环。

用 DFS 求出每个链和环的大小。

## 贪心法则一

考虑从大到小分配数字。

对于环来说，由于首尾节点可以多乘一次，相比链来说可以得到更大的乘积，所以应该优先分配环，然后再分配链。

## 贪心法则二

对于不同的环，大小越小，首尾节点的乘积越大，所以应该优先分配更小的环。

对于不同的链，比如大小为 $3$ 和大小为 $2$ 的两条链，有两种分配方案：

- $5,4,3$ 和 $2,1$。
- $5,4$ 和 $3,2,1$。

区别在于 $3$ 和谁相乘，那么和 $4$ 相乘更优。

一般地，优先分配更大的链，这样元素乘积更大。

所以，按照从小到大的顺序排序环（连通块大小），按照从大到小的顺序排序链（连通块大小）。

## 贪心法则三

假设分配给链的数字是 $1$ 到 $6$。

对于 $6$ 来说，与之相乘的两个数越大越好，所以 $6$ 的邻居是 $5$ 和 $4$。

对于 $5$ 来说，与之相乘的两个数越大越好，所以 $5$ 的另一个邻居是 $3$（注意 $4$ 已经是 $6$ 的邻居了）。

对于 $4$ 来说，另一个邻居是 $2$。

对于 $3$ 来说，另一个邻居是 $1$。

所以这条链是 $1\text{-}3\text{-}5\text{-}6\text{-}4\text{-}2$。

一般地，设分配给链的数字是 $l$ 到 $r$，那么乘积之和为

$$
r(r-1) + \sum_{i=l}^{r-2} i(i+2)
$$

> 注：上式可以继续化简成 $\mathcal{O}(1)$ 的公式。

对于环来说，上式额外加上首尾的乘积，即 $l(l+1)$。

```py [sol-Python3]
class Solution:
    def maxScore(self, n: int, edges: List[List[int]]) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        cycle = []
        chain = []
        vis = [False] * n

        def dfs(x: int):
            nonlocal cnt_v, cnt_e
            vis[x] = True
            cnt_v += 1
            cnt_e += len(g[x])
            for y in g[x]:
                if not vis[y]:
                    dfs(y)

        for i, b in enumerate(vis):
            if b:
                continue
            cnt_v = cnt_e = 0
            dfs(i)
            if cnt_v * 2 == cnt_e:  # 环
                cycle.append(cnt_v)
            elif cnt_e > 0:  # 链，但不考虑孤立点
                chain.append(cnt_v)

        ans = 0
        cur = n

        def f(sz: int, is_cycle: bool) -> None:
            nonlocal ans, cur
            l, r = cur - sz + 1, cur
            for i in range(l, r - 1):
                ans += i * (i + 2)
            ans += r * (r - 1)
            if is_cycle:
                ans += l * (l + 1)
            cur -= sz

        cycle.sort()
        for sz in cycle:
            f(sz, True)

        chain.sort(reverse=True)
        for sz in chain:
            f(sz, False)

        return ans
```

```java [sol-Java]
class Solution {
    private int cntV, cntE, cur;

    public long maxScore(int n, int[][] edges) {
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        List<Integer> cycle = new ArrayList<>();
        List<Integer> chain = new ArrayList<>();
        boolean[] vis = new boolean[n];

        for (int i = 0; i < n; i++) {
            if (vis[i]) {
                continue;
            }
            cntV = 0;
            cntE = 0;
            dfs(i, g, vis);
            if (cntV * 2 == cntE) { // 环
                cycle.add(cntV);
            } else if (cntE > 0) { // 链，但不考虑孤立点
                chain.add(cntV);
            }
        }

        long ans = 0;
        cur = n;

        cycle.sort(null);
        for (int sz : cycle) {
            ans += calc(sz, true);
        }

        chain.sort(Comparator.reverseOrder());
        for (int sz : chain) {
            ans += calc(sz, false);
        }

        return ans;
    }

    private void dfs(int x, List<Integer>[] g, boolean[] vis) {
        vis[x] = true;
        cntV++;
        cntE += g[x].size();
        for (int y : g[x]) {
            if (!vis[y]) {
                dfs(y, g, vis);
            }
        }
    }

    private long calc(int sz, boolean isCycle) {
        long res = 0;
        int l = cur - sz + 1;
        int r = cur;
        for (int i = l; i < r - 1; i++) {
            res += (long) i * (i + 2);
        }
        res += (long) r * (r - 1);
        if (isCycle) {
            res += (long) l * (l + 1);
        }
        cur -= sz;
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxScore(int n, vector<vector<int>>& edges) {
        vector g(n, vector<int>{});
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        vector<int> cycle, chain;
        vector<uint8_t> vis(n);
        int cnt_v, cnt_e;

        auto dfs = [&](this auto&& dfs, int x) -> void {
            vis[x] = 1;
            cnt_v++;
            cnt_e += g[x].size();
            for (int y : g[x]) {
                if (!vis[y]) {
                    dfs(y);
                }
            }
        };

        for (int i = 0; i < n; i++) {
            if (vis[i]) {
                continue;
            }
            cnt_v = 0;
            cnt_e = 0;
            dfs(i);
            if (cnt_v * 2 == cnt_e) { // 环
                cycle.push_back(cnt_v);
            } else if (cnt_e > 0) { // 链，但不考虑孤立点
                chain.push_back(cnt_v);
            }
        }

        long long ans = 0;
        int cur = n;

        auto f = [&](int sz, bool is_cycle) -> void {
            int l = cur - sz + 1, r = cur;
            for (int i = l; i < r - 1; i++) {
                ans += 1LL * i * (i + 2);
            }
            ans += 1LL * r * (r - 1);
            if (is_cycle) {
                ans += 1LL * l * (l + 1);
            }
            cur -= sz;
        };

        ranges::sort(cycle);
        for (int sz : cycle) {
            f(sz, true);
        }

        ranges::sort(chain, ranges::greater());
        for (int sz : chain) {
            f(sz, false);
        }

        return ans;
    }
};
```

```go [sol-Go]
func maxScore(n int, edges [][]int) int64 {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var cycle, chain []int
	var cntV, cntE int
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		cntV++
		cntE += len(g[x])
		for _, y := range g[x] {
			if !vis[y] {
				dfs(y)
			}
		}
	}
	for i, b := range vis {
		if b {
			continue
		}
		cntV, cntE = 0, 0
		dfs(i)
		if cntV*2 == cntE { // 环
			cycle = append(cycle, cntV)
		} else if cntE > 0 { // 链，但不考虑孤立点
			chain = append(chain, cntV)
		}
	}

	ans := 0
	cur := n
	f := func(sz int, isCycle bool) {
		l, r := cur-sz+1, cur
		for i := l; i < r-1; i++ {
			ans += i * (i + 2)
		}
		ans += r * (r - 1)
		if isCycle {
			ans += l * (l + 1)
		}
		cur -= sz
	}

	slices.Sort(cycle)
	for _, sz := range cycle {
		f(sz, true)
	}

	slices.SortFunc(chain, func(a, b int) int { return b - a })
	for _, sz := range chain {
		f(sz, false)
	}

	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
