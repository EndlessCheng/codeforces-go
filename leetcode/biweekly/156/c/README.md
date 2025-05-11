## 方法一：DFS

**枚举**路径起点 $x$，从 $x$ 开始，DFS 这个图。

定义 $\textit{dfs}(x,i,s)$ 表示当前移动到节点 $x$，路径长为 $i$，路径边权和为 $s$。

枚举 $x$ 的邻居 $y$，递归到 $\textit{dfs}(y,i+1,s+w)$，其中 $w$ 是有向边 $x\to y$ 的边权。如果 $s+w\ge t$ 则不递归。

递归边界：如果 $i=k$，那么用 $s$ 更新答案的最大值。

为避免重复访问同一个 $(x,i,s)$ 状态，可以用哈希集合记录访问过的状态。

> 注：其实这个图有环也可以，下面的算法仍然是正确的，因为参数 $i$ 是递增的。一条路径如果包含同一个节点多次，这个节点对应的 $i$ 是不同的。换言之，加上 $i$ 之后，没有后效性。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1m7EuzqEqr/?t=13m21s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxWeight(self, n: int, edges: List[List[int]], k: int, t: int) -> int:
        g = [[] for _ in range(n)]
        for x, y, wt in edges:
            g[x].append((y, wt))

        ans = -1
        @cache  # 也可以用 vis 哈希集合
        def dfs(x: int, i: int, s: int) -> None:
            if i == k:
                nonlocal ans
                ans = max(ans, s)
                return
            for y, wt in g[x]:
                if s + wt < t:
                    dfs(y, i + 1, s + wt)

        for x in range(n):  # 枚举起点
            dfs(x, 0, 0)
        return ans
```

```java [sol-Java]
class Solution {
    private int ans = -1;

    public int maxWeight(int n, int[][] edges, int k, int t) {
        List<int[]>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1], wt = e[2];
            g[x].add(new int[]{y, wt});
        }

        Set<Integer> vis = new HashSet<>();
        for (int x = 0; x < n; x++) { // 枚举起点
            dfs(x, 0, 0, g, k, t, vis);
        }
        return ans;
    }

    private void dfs(int x, int i, int s, List<int[]>[] g, int k, int t, Set<Integer> vis) {
        if (i == k) {
            ans = Math.max(ans, s);
            return;
        }
        int mask = x << 20 | i << 10 | s; // 每个参数存储在 10 个比特中
        if (!vis.add(mask)) { // 访问过
            return;
        }
        for (int[] e : g[x]) {
            int wt = e[1];
            if (s + wt < t) {
                dfs(e[0], i + 1, s + wt, g, k, t, vis);
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxWeight(int n, vector<vector<int>>& edges, int k, int t) {
        vector<vector<pair<int, int>>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1], wt = e[2];
            g[x].emplace_back(y, wt);
        }

        int ans = -1;
        unordered_set<int> vis;
        auto dfs = [&](this auto&& dfs, int x, int i, int s) -> void {
            if (i == k) {
                ans = max(ans, s);
                return;
            }
            int mask = x << 20 | i << 10 | s; // 每个参数存储在 10 个比特中
            if (!vis.insert(mask).second) { // 访问过
                return;
            }
            for (auto& [y, wt] : g[x]) {
                if (s + wt < t) {
                    dfs(y, i + 1, s + wt);
                }
            }
        };
        for (int x = 0; x < n; x++) { // 枚举起点
            dfs(x, 0, 0);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxWeight(n int, edges [][]int, k int, t int) int {
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
	}

	ans := -1
	type tuple struct{ x, i, s int }
	vis := map[tuple]bool{}
	var dfs func(int, int, int)
	dfs = func(x, i, s int) {
		if i == k {
			ans = max(ans, s)
			return
		}
		args := tuple{x, i, s}
		if vis[args] { // 访问过
			return
		}
		vis[args] = true
		for _, e := range g[x] {
			if s+e.wt < t {
				dfs(e.to, i+1, s+e.wt)
			}
		}
	}
	for x := range n { // 枚举起点
		dfs(x, 0, 0)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+m)kt)$，其中 $m$ 是 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(nkt)$。

## 方法二：拓扑序 DP（适用于 DAG）

如果你没有做过任何拓扑序 DP 的题目，请先完成 [2050. 并行课程 III](https://leetcode.cn/problems/parallel-courses-iii/)。

同方法一，定义 $f[x][i][s]$ 表示能否得到一个终点在 $x$，有 $i$ 条边且路径元素和为 $s$ 的路径。

用**刷表法**转移，遍历 $x$ 的邻居 $y$，如果 $f[x][i][s]=\texttt{true}$ 且 $s+w<t$，那么更新 $f[y][i+1][s+w]=\texttt{true}$。

> 注：在动态规划中，用转移来源更新当前状态叫**查表法**，用当前状态更新其他状态叫**刷表法**。

初始值：$f[x][0][0]=\texttt{true}$。没有边的时候，边权和为 $0$。

代码实现时，第三个维度可以用哈希集合维护，减少无效状态的遍历。

答案：$\max\limits_{x=0}^n\max(f[x][k])$。

如果图中有环怎么办？请看方法三。

```py [sol-Python3]
class Solution:
    def maxWeight(self, n: int, edges: List[List[int]], k: int, t: int) -> int:
        g = [[] for _ in range(n)]
        deg = [0] * n
        for x, y, wt in edges:
            g[x].append((y, wt))
            deg[y] += 1

        ans = -1
        f = [[set() for _ in range(k + 1)] for _ in range(n)]
        q = deque(i for i, d in enumerate(deg) if d == 0)
        while q:
            x = q.popleft()
            f[x][0].add(0)  # x 单独一个点，路径边权和为 0
            if f[x][k]:
                ans = max(ans, max(f[x][k]))  # 恰好 k 条边
            for y, wt in g[x]:
                for i in range(k):
                    for s in f[x][i]:
                        if s + wt < t:
                            f[y][i + 1].add(s + wt)
                deg[y] -= 1
                if deg[y] == 0:
                    q.append(y)
        return ans
```

```java [sol-Java]
class Solution {
    public int maxWeight(int n, int[][] edges, int k, int t) {
        List<int[]>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        int[] deg = new int[n];
        for (int[] e : edges) {
            int x = e[0], y = e[1], wt = e[2];
            g[x].add(new int[]{y, wt});
            deg[y]++;
        }

        int ans = -1;
        Set<Integer>[][] f = new HashSet[n][k + 1];
        for (Set<Integer>[] row : f) {
            Arrays.setAll(row, i -> new HashSet<>());
        }
        Queue<Integer> q = new ArrayDeque<>();
        for (int i = 0; i < n; i++) {
            if (deg[i] == 0) {
                q.add(i);
            }
        }
        while (!q.isEmpty()) {
            int x = q.poll();
            f[x][0].add(0); // x 单独一个点，路径边权和为 0
            for (int s : f[x][k]) { // 恰好 k 条边
                ans = Math.max(ans, s);
            }
            for (int[] e : g[x]) {
                int y = e[0], wt = e[1];
                for (int i = 0; i < k; i++) {
                    for (int s : f[x][i]) {
                        if (s + wt < t) {
                            f[y][i + 1].add(s + wt);
                        }
                    }
                }
                if (--deg[y] == 0) {
                    q.add(y);
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxWeight(int n, vector<vector<int>>& edges, int k, int t) {
        vector<vector<pair<int, int>>> g(n);
        vector<int> deg(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1], wt = e[2];
            g[x].emplace_back(y, wt);
            deg[y]++;
        }

        int ans = -1;
        vector f(n, vector<unordered_set<int>>(k + 1));
        queue<int> q;
        for (int i = 0; i < n; i++) {
            if (deg[i] == 0) {
                q.push(i);
            }
        }
        while (!q.empty()) {
            int x = q.front();
            q.pop();
            f[x][0].insert(0); // x 单独一个点，路径边权和为 0
            for (int s : f[x][k]) { // 恰好 k 条边
                ans = max(ans, s);
            }
            for (auto& [y, wt] : g[x]) {
                for (int i = 0; i < k; i++) {
                    for (int s : f[x][i]) {
                        if (s + wt < t) {
                            f[y][i + 1].insert(s + wt);
                        }
                    }
                }
                if (--deg[y] == 0) {
                    q.push(y);
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxWeight(n int, edges [][]int, k int, t int) int {
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	deg := make([]int, n)
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
	}

	ans := -1
	f := make([][]map[int]struct{}, n)
	for i := range f {
		f[i] = make([]map[int]struct{}, k+1)
		for j := range f[i] {
			f[i][j] = map[int]struct{}{}
		}
	}
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		f[x][0][0] = struct{}{} // x 单独一个点，路径边权和为 0
		for s := range f[x][k] { // 恰好 k 条边
			ans = max(ans, s)
		}
		for _, e := range g[x] {
			y, wt := e.to, e.wt
			for i, st := range f[x][:k] {
				for s := range st {
					if s+wt < t {
						f[y][i+1][s+wt] = struct{}{}
					}
				}
			}
			deg[y]--
			if deg[y] == 0 {
				q = append(q, y)
			}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+m)kt)$，其中 $m$ 是 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(nkt)$。

### 附：bitset 优化

把集合用二进制数表示，原理见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

本题集合太大，需要用 bitset 表示二进制数。

```cpp [sol-C++]
class Solution {
public:
    int maxWeight(int n, vector<vector<int>>& edges, int k, int t) {
        vector<vector<pair<int, int>>> g(n);
        vector<int> deg(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1], wt = e[2];
            g[x].emplace_back(y, wt);
            deg[y]++;
        }

        int ans = -1;
        vector f(n, vector<bitset<600>>(k + 1));
        queue<int> q;
        for (int i = 0; i < n; i++) {
            if (deg[i] == 0) {
                q.push(i);
            }
        }
        while (!q.empty()) {
            int x = q.front();
            q.pop();
            f[x][0].set(0); // x 单独一个点，路径边权和为 0
            for (int s = t - 1; s >= 0; s--) {
                if (f[x][k].test(s)) { // 恰好 k 条边
                    ans = max(ans, s);
                    break;
                }
            }
            for (auto& [y, wt] : g[x]) {
                for (int i = 0; i < k && !f[x][i].none(); i++) {
                    f[y][i + 1] |= f[x][i] << wt;
                }
                if (--deg[y] == 0) {
                    q.push(y);
                }
            }
        }
        return ans;
    }
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+m)kt/w)$，其中 $m$ 是 $\textit{edges}$ 的长度，$w=64$ 或 $32$。
- 空间复杂度：$\mathcal{O}(nkt/w)$。

## 方法三：一般图 DP

在方法一和方法二中，边的个数 $i$ 是递增的，我们可以从 $i$ 入手计算状态。

略微调整状态定义，用 $f[i][x][s]$ 表示能否得到一个终点在 $x$，有 $i$ 条边且路径元素和为 $s$ 的路径。同样地，第三个维度可以用哈希集合维护，减少无效状态的遍历。

**无需建图**，遍历 $\textit{edges}$，对于边权为 $w$ 的有向边 $x\to y$，如果 $f[i][x]$ 中有 $s$ 且 $s+w<t$，那么把 $s+w$ 加到 $f[i+1][y]$ 中。

初始值：每个 $f[0][i]$ 都添加 $0$。没有边的时候，边权和为 $0$。

答案：$\max\limits_{x=0}^n\max(f[k][x])$。

```py [sol-Python3]
class Solution:
    def maxWeight(self, n: int, edges: List[List[int]], k: int, t: int) -> int:
        f = [[set() for _ in range(n)] for _ in range(k + 1)]
        for i in range(n):
            f[0][i].add(0)
        for i in range(k):
            for x, y, wt in edges:
                for s in f[i][x]:
                    if s + wt < t:
                        f[i + 1][y].add(s + wt)
        return max((max(st) for st in f[k] if st), default=-1)
```

```java [sol-Java]
class Solution {
    public int maxWeight(int n, int[][] edges, int k, int t) {
        Set<Integer>[][] f = new HashSet[k + 1][n];
        for (Set<Integer>[] row : f) {
            Arrays.setAll(row, i -> new HashSet<>());
        }
        for (int i = 0; i < n; i++) {
            f[0][i].add(0);
        }
        for (int i = 0; i < k; i++) {
            for (int[] e : edges) {
                int x = e[0], y = e[1], wt = e[2];
                for (int s : f[i][x]) {
                    if (s + wt < t) {
                        f[i + 1][y].add(s + wt);
                    }
                }
            }
        }

        int ans = -1;
        for (Set<Integer> st : f[k]) {
            for (int s : st) {
                ans = Math.max(ans, s);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxWeight(int n, vector<vector<int>>& edges, int k, int t) {
        vector f(k + 1, vector<unordered_set<int>>(n));
        for (int i = 0; i < n; i++) {
            f[0][i].insert(0);
        }
        for (int i = 0; i < k; i++) {
            for (auto& e : edges) {
                int x = e[0], y = e[1], wt = e[2];
                for (int s : f[i][x]) {
                    if (s + wt < t) {
                        f[i + 1][y].insert(s + wt);
                    }
                }
            }
        }

        int ans = -1;
        for (auto& st : f[k]) {
            for (int s : st) {
                ans = max(ans, s);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxWeight(n int, edges [][]int, k int, t int) int {
	f := make([][]map[int]struct{}, k+1)
	for i := range f {
		f[i] = make([]map[int]struct{}, n)
		for j := range f[i] {
			f[i][j] = map[int]struct{}{}
		}
	}
	for i := range f[0] {
		f[0][i][0] = struct{}{}
	}
	for i, sets := range f[:k] {
		for _, e := range edges {
			x, y, wt := e[0], e[1], e[2]
			for s := range sets[x] {
				if s+wt < t {
					f[i+1][y][s+wt] = struct{}{}
				}
			}
		}
	}

	ans := -1
	for _, set := range f[k] {
		for s := range set {
			ans = max(ans, s)
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mkt)$，其中 $m$ 是 $\textit{edges}$ 的长度。忽略初始化 $f$ 的时间。
- 空间复杂度：$\mathcal{O}(nkt)$。注：可以用滚动数组优化至 $\mathcal{O}(nt)$。

### 附：bitset 优化

原理同方法二。

```py [sol-Python3]
class Solution:
    def maxWeight(self, n: int, edges: List[List[int]], k: int, t: int) -> int:
        MASK = (1 << t) - 1
        f = [[0] * n for _ in range(k + 1)]
        f[0] = [1] * n
        for i in range(k):
            for x, y, wt in edges:
                if f[i][x]:
                    f[i + 1][y] |= (f[i][x] << wt) & MASK
        return max(s.bit_length() for s in f[k]) - 1
```

```java [sol-Java]
import java.math.BigInteger;

class Solution {
    public int maxWeight(int n, int[][] edges, int k, int t) {
        BigInteger[][] f = new BigInteger[k + 1][n];
        Arrays.fill(f[0], BigInteger.ONE);
        for (int i = 1; i <= k; i++) {
            Arrays.fill(f[i], BigInteger.ZERO);
        }

        final BigInteger MASK = BigInteger.ONE.shiftLeft(t).subtract(BigInteger.ONE);
        for (int i = 0; i < k; i++) {
            for (int[] e : edges) {
                int x = e[0], y = e[1], wt = e[2];
                if (!f[i][x].equals(BigInteger.ZERO)) {
                    BigInteger shifted = f[i][x].shiftLeft(wt).and(MASK);
                    f[i + 1][y] = f[i + 1][y].or(shifted);
                }
            }
        }

        int ans = 0;
        for (int i = 0; i < n; i++) {
            ans = Math.max(ans, f[k][i].bitLength());
        }
        return ans - 1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxWeight(int n, vector<vector<int>>& edges, int k, int t) {
        vector f(k + 1, vector<bitset<600>>(n));
        for (int i = 0; i < n; i++) {
            f[0][i].set(0);
        }
        for (int i = 0; i < k; i++) {
            for (auto& e : edges) {
                int x = e[0], y = e[1], wt = e[2];
                if (!f[i][x].none()) {
                    f[i + 1][y] |= f[i][x] << wt;
                }
            }
        }

        int ans = -1;
        for (auto& bs : f[k]) {
            for (int s = t - 1; s >= 0; s--) {
                if (bs.test(s)) {
                    ans = max(ans, s);
                    break;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxWeight(n int, edges [][]int, k int, t int) int {
	f := make([][]*big.Int, k+1)
	for i := range f {
		f[i] = make([]*big.Int, n)
		for j := range f[i] {
			f[i][j] = big.NewInt(0)
		}
	}
	for i := range f[0] {
		f[0][i] = big.NewInt(1)
	}

	p := new(big.Int)
	mask := new(big.Int).Sub(p.Lsh(big.NewInt(1), uint(t)), big.NewInt(1))
	for i, fi := range f[:k] {
		for _, e := range edges {
			x, y, wt := e[0], e[1], e[2]
			if fi[x].Sign() != 0 {
				shifted := p.And(p.Lsh(fi[x], uint(wt)), mask)
				f[i+1][y].Or(f[i+1][y], shifted)
			}
		}
	}

	ans := 0
	for _, bi := range f[k] {
		ans = max(ans, bi.BitLen())
	}
	return ans - 1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mkt/w)$，其中 $m$ 是 $\textit{edges}$ 的长度，$w=64$ 或 $32$。忽略初始化 $f$ 的时间。
- 空间复杂度：$\mathcal{O}(nkt/w)$。注：可以用滚动数组优化至 $\mathcal{O}(nt/w)$。

## 相似题目

- [2050. 并行课程 III](https://leetcode.cn/problems/parallel-courses-iii/) 2084
- [1857. 有向图中最大颜色值](https://leetcode.cn/problems/largest-color-value-in-a-directed-graph/) 2313

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
