## 方法一：排列型回溯+最优性剪枝

排列型回溯，枚举所有开锁顺序（全排列），具体见[【基础算法精讲 16】](https://www.bilibili.com/video/BV1mY411D7f6/)。

设当前开了 $i$ 把锁，那么 $x=1+k\cdot i$。

枚举开第 $j$ 把锁，需要的时间为 $\left\lceil\dfrac{strength[j]}{x}\right\rceil$。

最优性剪枝：由于用时 $\textit{time}$ 在递归过程中只会增大，如果发现 $\textit{time}\ge \textit{ans}$，那么直接返回，不再往下递归。

关于上取整的计算，当 $a$ 和 $b$ 均为正整数时，我们有

$$
\left\lceil\dfrac{a}{b}\right\rceil = \left\lfloor\dfrac{a-1}{b}\right\rfloor + 1
$$

讨论 $a$ 被 $b$ 整除，和不被 $b$ 整除两种情况，可以证明上式的正确性。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1YeqHYSEXv/?t=3m54s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def findMinimumTime(self, strength: List[int], k: int) -> int:
        n = len(strength)
        ans = inf
        done = [False] * n
        def dfs(i: int, time: int) -> None:
            nonlocal ans
            # 最优性剪枝：答案不可能变小
            if time >= ans:
                return
            if i == n:
                ans = time
                return
            x = 1 + i * k
            for j, s in enumerate(strength):
                if not done[j]:
                    done[j] = True  # 已开锁
                    dfs(i + 1, time + (s - 1) // x + 1)
                    done[j] = False  # 恢复现场
        dfs(0, 0)
        return ans
```

```java [sol-Java]
class Solution {
    public int findMinimumTime(List<Integer> strength, int k) {
        boolean[] done = new boolean[strength.size()];
        dfs(0, 0, strength.toArray(Integer[]::new), k, done);
        return ans;
    }

    private int ans = Integer.MAX_VALUE;

    private void dfs(int i, int time, Integer[] strength, int k, boolean[] done) {
        // 剪枝：答案不可能变小
        if (time >= ans) {
            return;
        }
        if (i == strength.length) {
            ans = time;
            return;
        }
        int x = 1 + k * i;
        for (int j = 0; j < strength.length; j++) {
            if (!done[j]) {
                done[j] = true; // 已开锁
                dfs(i + 1, time + (strength[j] - 1) / x + 1, strength, k, done);
                done[j] = false; // 恢复现场
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findMinimumTime(vector<int>& strength, int k) {
        int ans = numeric_limits<int>::max();
        int n = strength.size();
        vector<int> done(n);
        auto dfs = [&](auto&& dfs, int i, int time) -> void {
            // 剪枝：答案不可能变小
            if (time >= ans) {
                return;
            }
            if (i == n) {
                ans = time;
                return;
            }
            int x = 1 + k * i;
            for (int j = 0; j < n; j++) {
                if (!done[j]) {
                    done[j] = 1; // 已开锁
                    dfs(dfs, i + 1, time + (strength[j] - 1) / x + 1);
                    done[j] = 0; // 恢复现场
                }
            }
        };
        dfs(dfs, 0, 0);
        return ans;
    }
};
```

```go [sol-Go]
func findMinimumTime(strength []int, k int) int {
	ans := math.MaxInt
	n := len(strength)
	done := make([]bool, n)
	var dfs func(int, int)
	dfs = func(i, time int) {
		// 最优性剪枝：答案不可能变小
		if time >= ans {
			return
		}
		if i == n {
			ans = time
			return
		}
		x := 1 + k*i
		for j, d := range done {
			if !d {
				done[j] = true // 已开锁
				dfs(i+1, time+(strength[j]-1)/x+1)
				done[j] = false // 恢复现场
			}
		}
	}
	dfs(0, 0)
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n!)$，其中 $n$ 是 $\textit{strength}$ 的长度。证明见[【基础算法精讲 16】](https://www.bilibili.com/video/BV1mY411D7f6/)。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：状压 DP

如果你没做过状压 DP，请先完成 [526. 优美的排列](https://leetcode.cn/problems/beautiful-arrangement/)，并阅读我的题解 [教你一步步思考状压 DP：从记忆化搜索到递推](https://leetcode.cn/problems/beautiful-arrangement/solution/jiao-ni-yi-bu-bu-si-kao-zhuang-ya-dpcong-c6kd/)。

定义 $\textit{dfs}(i)$ 表示尚未开锁的下标集合为 $i$ 时，打开所有锁需要的最少时间。

当前的 $x = 1 + k\cdot (n-|i|)$，其中 $|i|$ 是集合 $i$ 的大小。

枚举开第 $j$ 把锁，那么问题变成：尚未开锁的下标集合为 $i\setminus\{j\}$ 时，打开所有锁需要的最少时间，即 $\textit{dfs}(i\setminus\{j\})$。

取最小值，得

$$
\textit{dfs}(i) = \min\limits_{j\in i}  \textit{dfs}(i\setminus\{j\}) + \left\lceil\dfrac{strength[j]}{x}\right\rceil
$$

递归边界：$\textit{dfs}(\varnothing) = 0$。

递归入口：$\textit{dfs}(U)$，其中全集 $U=\{0,1,2,\ldots,n-1\}$。

代码实现时，用二进制表示集合，用位运算实现集合操作，具体请看 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

### 记忆化搜索

```py [sol-Python3]
class Solution:
    def findMinimumTime(self, strength: List[int], k: int) -> int:
        n = len(strength)
        @cache
        def dfs(i: int) -> int:
            if i == 0:
                return 0
            x = 1 + k * (n - i.bit_count())
            return min(dfs(i ^ (1 << j)) + (s - 1) // x
                       for j, s in enumerate(strength) if i >> j & 1) + 1
        return dfs((1 << n) - 1)
```

```java [sol-Java]
class Solution {
    public int findMinimumTime(List<Integer> strength, int k) {
        int n = strength.size();
        int[] memo = new int[1 << n];
        Arrays.fill(memo, -1);
        return dfs((1 << n) - 1, strength.toArray(Integer[]::new), k, n, memo);
    }

    private int dfs(int i, Integer[] strength, int k, int n, int[] memo) {
        if (i == 0) {
            return 0;
        }
        if (memo[i] != -1) { // 之前计算过
            return memo[i];
        }
        int x = 1 + k * (n - Integer.bitCount(i));
        int res = Integer.MAX_VALUE;
        for (int j = 0; j < n; j++) {
            if ((i >> j & 1) > 0) {
                res = Math.min(res, dfs(i ^ (1 << j), strength, k, n, memo) + (strength[j] - 1) / x + 1);
            }
        }
        return memo[i] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findMinimumTime(vector<int>& strength, int k) {
        int n = strength.size();
        vector<int> memo(1 << n, INT_MAX);
        auto dfs = [&](auto&& dfs, int i) -> int {
            if (i == 0) {
                return 0;
            }
            int& res = memo[i]; // 注意这里是引用
            if (memo[i] != INT_MAX) {
                return memo[i];
            }
            int x = 1 + k * (n - popcount((unsigned) i));
            for (int j = 0; j < n; j++) {
                if (i >> j & 1) {
                    res = min(res, dfs(dfs, i ^ (1 << j)) + (strength[j] - 1) / x + 1);
                }
            }
            return res;
        };
        return dfs(dfs, (1 << n) - 1);
    }
};
```

```go [sol-Go]
func findMinimumTime(strength []int, k int) int {
	n := len(strength)
	memo := make([]int, 1<<n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i == 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}
		x := 1 + k*(n-bits.OnesCount(uint(i)))
		res := math.MaxInt
		for j, s := range strength {
			if i>>j&1 > 0 {
				res = min(res, dfs(i^1<<j)+(s-1)/x+1)
			}
		}
		*p = res // 记忆化
		return res
	}
	return dfs(1<<n - 1)
}
```

### 递推

```py [sol-Python3]
class Solution:
    def findMinimumTime(self, strength: List[int], k: int) -> int:
        n = len(strength)
        f = [0] * (1 << n)
        for i in range(1, 1 << n):
            x = 1 + k * (n - i.bit_count())
            f[i] = min(f[i ^ (1 << j)] + (s - 1) // x
                       for j, s in enumerate(strength) if i >> j & 1) + 1
        return f[-1]
```

```java [sol-Java]
class Solution {
    public int findMinimumTime(List<Integer> strength, int k) {
        Integer[] a = strength.toArray(Integer[]::new);
        int n = a.length;
        int m = 1 << n;
        int[] f = new int[m];
        Arrays.fill(f, 1, m, Integer.MAX_VALUE);
        for (int i = 1; i < m; i++) {
            int x = 1 + k * (n - Integer.bitCount(i));
            for (int j = 0; j < n; j++) {
                if ((i >> j & 1) > 0) {
                    f[i] = Math.min(f[i], f[i ^ (1 << j)] + (a[j] - 1) / x + 1);
                }
            }
        }
        return f[m - 1];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findMinimumTime(vector<int>& strength, int k) {
        int n = strength.size();
        int m = 1 << n;
        vector<int> f(m, INT_MAX);
        f[0] = 0;
        for (unsigned i = 1; i < m; i++) {
            int x = 1 + k * (n - popcount(i));
            for (int j = 0; j < n; j++) {
                if (i >> j & 1) {
                    f[i] = min(f[i], f[i ^ (1 << j)] + (strength[j] - 1) / x + 1);
                }
            }
        }
        return f[m - 1];
    }
};
```

```go [sol-Go]
func findMinimumTime(strength []int, k int) int {
	n := len(strength)
	m := 1 << n
	f := make([]int, m)
	for i := 1; i < m; i++ {
		x := 1 + k*(n-bits.OnesCount(uint(i)))
		f[i] = math.MaxInt
		for j, s := range strength {
			if i>>j&1 > 0 {
				f[i] = min(f[i], f[i^1<<j]+(s-1)/x+1)
			}
		}
	}
	return f[m-1]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n2^n)$，其中 $n$ 是 $\textit{strength}$ 的长度。
- 空间复杂度：$\mathcal{O}(2^n)$。

## 方法三：最小费用流

如果 $n=100$ 呢？

创建一个**二分图**，左部为锁的编号 $i=0,1,2,\ldots,n-1$，右部表示这个锁是第 $j=0,1,2,\ldots,n-1$ 次开的。

问题相当于计算这个二分图的**最小带权匹配**。

建图：

- 从 $i$ 向 $j$ 连边，容量为 $\infty$，费用为 $\left\lceil\dfrac{\textit{strength}[i]}{1 + k\cdot j}\right\rceil$。
- 从超级源点 $S=2n$ 向每个 $i$ 连边，容量为 $1$，费用 $0$。
- 从每个 $j$ 向超级汇点 $T=2n+1$ 连边，容量为 $1$，费用为 $0$。

计算从 $S$ 到 $T$ 的最小费用流，满流时的费用即为答案。

```go [sol-Go]
func findMinimumTime(strength []int, k int) int {
	n := len(strength)
	S := n * 2
	T := S + 1

	// rid 为反向边在邻接表中的下标
	type neighbor struct{ to, rid, cap, cost int }
	g := make([][]neighbor, T+1)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost})
	}
	for i, s := range strength {
		// 枚举这个锁是第几次开的
		for j := range n {
			x := 1 + k*j
			addEdge(i, n+j, math.MaxInt, (s-1)/x+1)
		}
		addEdge(S, i, 1, 0)
	}
	for i := n; i < n*2; i++ {
		addEdge(i, T, 1, 0)
	}

	// 下面是最小费用最大流模板
	dis := make([]int, len(g))
	type vi struct{ v, i int }
	fa := make([]vi, len(g))
	inQ := make([]bool, len(g))
	spfa := func() bool {
		for i := range dis {
			dis[i] = math.MaxInt
		}
		dis[S] = 0
		inQ[S] = true
		q := []int{S}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			inQ[v] = false
			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				newD := dis[v] + e.cost
				if newD < dis[w] {
					dis[w] = newD
					fa[w] = vi{v, i}
					if !inQ[w] {
						inQ[w] = true
						q = append(q, w)
					}
				}
			}
		}
		// 循环结束后所有 inQ[v] 都为 false，无需重置
		return dis[T] < math.MaxInt
	}

	minCost := 0
	for spfa() {
		// 沿 st-end 的最短路尽量增广
		// 特别地，如果建图时所有边的容量都设为 1，那么 minF 必然为 1，下面第一个 for 循环可以省略
		minF := math.MaxInt
		for v := T; v != S; {
			p := fa[v]
			minF = min(minF, g[p.v][p.i].cap)
			v = p.v
		}
		for v := T; v != S; {
			p := fa[v]
			e := &g[p.v][p.i]
			e.cap -= minF
			g[v][e.rid].cap += minF
			v = p.v
		}
		minCost += dis[T] * minF
	}
	return minCost
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^3)$，其中 $n$ 是 $\textit{strength}$ 的长度。由于二分图的特殊性，算法跑至多 $n$ 次 $\mathcal{O}(n^2)$ 的 SPFA 就结束了。
- 空间复杂度：$\mathcal{O}(n^2)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. 【本题相关】[图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. 【本题相关】[链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
