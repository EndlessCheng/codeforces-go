## 分类讨论

在以 $0$ 为根的前提下，设子树 $x$ 的点权异或和为 $\textit{xor}[x]$。

本题我们要计算连通块的点权异或和。由于 $n\le 1000$，可以考虑 $\mathcal{O}(n^2)$ 枚举要删除的两条边 $x_1\text{-}y_1$ 和 $x_2\text{-}y_2$（其中 $x_i$ 是 $y_i$ 的父节点），这会产生以下三种情况：

1. 删除的两条边在同一棵子树内，且 $y_1$ 在上，是 $x_2$ 的祖先（或重合）。
   
   如下图所示，这三个连通块从下到上，异或和分别为 $\textit{xor}[y_2]$、$\textit{xor}[y_1]\oplus\textit{xor}[y_2]$ 和 $\textit{xor}[0]\oplus\textit{xor}[y_1]$（$\oplus$ 表示异或运算）。**注**：把异或理解为减法，从子树中去掉节点的值，在异或运算下，就是异或这个节点的值。
   
   ![same.png](https://pic.leetcode-cn.com/1656215504-THYbIW-same.png){:width=260px}
   
2. 删除的两条边在同一棵子树内，且 $y_2$ 在上，是 $x_1$ 的祖先（或重合）。
   
   同上，这三个连通块的异或和分别为 $\textit{xor}[y_1]$、$\textit{xor}[y_2]\oplus\textit{xor}[y_1]$ 和 $\textit{xor}[0]\oplus\textit{xor}[y_2]$。
   
3. 删除的两条边分别属于两棵不相交的子树。
   
   如下图所示，这三个连通块的异或和分别为 $\textit{xor}[y_1]$、$\textit{xor}[y_2]$ 和 $\textit{xor}[0]\oplus\textit{xor}[y_1]\oplus\textit{xor}[y_2]$。
   
   ![diff.png](https://pic.leetcode-cn.com/1656215343-RsPtkr-diff.png){:width=330px}

现在要解决的问题是，如何判断 $x_1\text{-}y_1$ 和 $x_2\text{-}y_2$ 这两条边的**位置关系**？谁在上，谁在下，还是分别属于两棵不相交的子树？

## 分析递归过程

我们把往下深入的过程叫做「递」，把往上返回的过程叫做「归」。

在递归过程中，对于两个节点 $x$ 和 $y$：

- 如果顺序是 $递\ x\to 递\ y\to 归\ y\to 归\ x$，那么 $x$ 是 $y$ 的祖先。
- 如果顺序是 $递\ y\to 递\ x\to 归\ x\to 归\ y$，那么 $y$ 是 $x$ 的祖先。
- 如果顺序是 $递\ x\to 归\ x\to 递\ y\to 归\ y$，或者 $递\ y\to 归\ y\to 递\ x\to 归\ x$，那么 $x$ 和 $y$ 就分别属于两棵不相交的子树。

## 引入时间戳

怎么知道上述「递」和「归」发生的先后顺序？

在 DFS 一棵树的过程中，维护一个全局的时间戳 $\textit{clock}$，每访问一个新的节点，就把 $\textit{clock}$ 加一。

对于每个节点 $x$，记录进入这个节点的时间戳 $\textit{in}[x]$，和从这个节点往上返回时的时间戳 $\textit{out}[x]$。

根据时间戳，分类讨论：

- 如果 $x$ 是 $y$ 的祖先，那么区间 $[\textit{in}[x],\textit{out}[x]]$ 包含区间 $[\textit{in}[y],\textit{out}[y]]$。
- 如果 $y$ 是 $x$ 的祖先，那么区间 $[\textit{in}[y],\textit{out}[y]]$ 包含区间 $[\textit{in}[x],\textit{out}[x]]$。
- 如果没有区间包含关系，那么 $x$ 和 $y$ 就分别属于两棵不相交的子树。

具体地，当 $x\ne y$ 时，如果

$$
\textit{in}[x]<\textit{in}[y]\le\textit{out}[y]\le\textit{out}[x]
$$

那么 $x$ 是 $y$ 的祖先。其中 $\textit{out}[y]\le\textit{out}[x]$ 是因为递归结束的时候，时间戳是不变的，所以可以相等。

由于 $\textit{in}[y]\le\textit{out}[y]$ 恒成立，判断方式可以简化为

$$
\textit{in}[x]<\textit{in}[y]\le\textit{out}[x]
$$

代码实现时，由于不知道 $\textit{edges}[i]$ 两个端点的父子关系，枚举边的写法需要额外的判断。我们规定 $0$ 是树的根（没有父节点），枚举不是 $0$ 的两个节点，删除这两个节点到其父节点的边，这样写更简单。

```py [sol-Python3]
class Solution:
    def minimumScore(self, nums: List[int], edges: List[List[int]]) -> int:
        n = len(nums)
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        xor, in_, out = [0] * n, [0] * n, [0] * n
        clock = 0
        def dfs(x: int, fa: int) -> None:
            nonlocal clock
            clock += 1
            in_[x] = clock  # 递
            xor[x] = nums[x]
            for y in g[x]:
                if y != fa:
                    dfs(y, x)
                    xor[x] ^= xor[y]
            out[x] = clock  # 归
        dfs(0, -1)

        # 判断 x 是否为 y 的祖先
        def is_ancestor(x: int, y: int) -> bool:
            return in_[x] < in_[y] <= out[x]

        ans = inf
        # 枚举：删除 x 与 x 父节点之间的边，删除 y 与 y 父节点之间的边
        for x in range(2, n):
            for y in range(1, x):
                if is_ancestor(x, y):  # x 是 y 的祖先
                    a, b, c = xor[y], xor[x] ^ xor[y], xor[0] ^ xor[x]
                elif is_ancestor(y, x):  # y 是 x 的祖先
                    a, b, c = xor[x], xor[x] ^ xor[y], xor[0] ^ xor[y]
                else:  # x 和 y 分别属于两棵不相交的子树
                    a, b, c = xor[x], xor[y], xor[0] ^ xor[x] ^ xor[y]
                ans = min(ans, max(a, b, c) - min(a, b, c))
                if ans == 0:  # 不可能变小
                    return 0  # 提前返回
        return ans
```

```java [sol-Java]
class Solution {
    public int minimumScore(int[] nums, int[][] edges) {
        int n = nums.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        int[] xor = new int[n];
        int[] in = new int[n];
        int[] out = new int[n];
        dfs(0, -1, g, nums, xor, in, out);

        int ans = Integer.MAX_VALUE;
        // 枚举：删除 x 与 x 父节点之间的边，删除 y 与 y 父节点之间的边
        for (int x = 2; x < n; x++) {
            for (int y = 1; y < x; y++) {
                int a, b, c;
                if (isAncestor(x, y, in, out)) { // x 是 y 的祖先
                    a = xor[y];
                    b = xor[x] ^ a;
                    c = xor[0] ^ xor[x];
                } else if (isAncestor(y, x, in, out)) { // y 是 x 的祖先
                    a = xor[x];
                    b = xor[y] ^ a;
                    c = xor[0] ^ xor[y];
                } else { // x 和 y 分别属于两棵不相交的子树
                    a = xor[x];
                    b = xor[y];
                    c = xor[0] ^ a ^ b;
                }
                ans = Math.min(ans, Math.max(Math.max(a, b), c) - Math.min(Math.min(a, b), c));
                if (ans == 0) { // 不可能变小
                    return 0; // 提前返回
                }
            }
        }
        return ans;
    }

    private int clock = 0;

    private void dfs(int x, int fa, List<Integer>[] g, int[] nums, int[] xor, int[] in, int[] out) {
        in[x] = ++clock; // 递
        xor[x] = nums[x];
        for (int y : g[x]) {
            if (y != fa) {
                dfs(y, x, g, nums, xor, in, out);
                xor[x] ^= xor[y];
            }
        }
        out[x] = clock; // 归
    }

    // 判断 x 是否为 y 的祖先
    private boolean isAncestor(int x, int y, int[] in, int[] out) {
        return in[x] < in[y] && in[y] <= out[x];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumScore(vector<int>& nums, vector<vector<int>>& edges) {
        int n = nums.size();
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        vector<int> xr(n), in(n), out(n);
        int clock = 0;
        auto dfs = [&](this auto&& dfs, int x, int fa) -> void {
            in[x] = ++clock; // 递
            xr[x] = nums[x];
            for (int y : g[x]) {
                if (y != fa) {
                    dfs(y, x);
                    xr[x] ^= xr[y];
                }
            }
            out[x] = clock; // 归
        };
        dfs(0, -1);

        // 判断 x 是否为 y 的祖先
        auto is_ancestor = [&](int x, int y) -> bool {
            return in[x] < in[y] && in[y] <= out[x];
        };

        int ans = INT_MAX;
        // 枚举：删除 x 与 x 父节点之间的边，删除 y 与 y 父节点之间的边
        for (int x = 2; x < n; x++) {
            for (int y = 1; y < x; y++) {
                int a, b, c;
                if (is_ancestor(x, y)) { // x 是 y 的祖先
                    a = xr[y], b = xr[x] ^ a, c = xr[0] ^ xr[x];
                } else if (is_ancestor(y, x)) { // y 是 x 的祖先
                    a = xr[x], b = xr[y] ^ a, c = xr[0] ^ xr[y];
                } else { // x 和 y 分别属于两棵不相交的子树
                    a = xr[x], b = xr[y], c = xr[0] ^ a ^ b;
                }
                ans = min(ans, max({a, b, c}) - min({a, b, c}));
                if (ans == 0) { // 不可能变小
                    return 0; // 提前返回
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumScore(nums []int, edges [][]int) int {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	xor := make([]int, n)
	in := make([]int, n)
	out := make([]int, n)
	clock := 0
	var dfs func(int, int)
	dfs = func(x, fa int) {
		clock++
		in[x] = clock // 递
		xor[x] = nums[x]
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x)
				xor[x] ^= xor[y]
			}
		}
		out[x] = clock // 归
	}
	dfs(0, -1)

	// 判断 x 是否为 y 的祖先
	isAncestor := func(x, y int) bool {
		return in[x] < in[y] && in[y] <= out[x]
	}

	ans := math.MaxInt
	// 枚举：删除 x 与 x 父节点之间的边，删除 y 与 y 父节点之间的边
	for x := 2; x < n; x++ {
		for y := 1; y < x; y++ {
			var a, b, c int
			if isAncestor(x, y) { // x 是 y 的祖先
				a, b, c = xor[y], xor[x]^xor[y], xor[0]^xor[x]
			} else if isAncestor(y, x) { // y 是 x 的祖先
				a, b, c = xor[x], xor[x]^xor[y], xor[0]^xor[y]
			} else { // x 和 y 分别属于两棵不相交的子树
				a, b, c = xor[x], xor[y], xor[0]^xor[x]^xor[y]
			}
			ans = min(ans, max(a, b, c)-min(a, b, c))
			if ans == 0 { // 不可能变小
				return 0 // 提前返回
			}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{nums}$ 的长度，也是树的节点个数。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面树题单的「§3.7 DFS 时间戳」。

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
11. 【本题相关】[链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
