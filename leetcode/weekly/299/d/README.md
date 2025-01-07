本题 [视频讲解](https://www.bilibili.com/video/BV1pW4y1r7xs) 已出炉，欢迎点赞三连~

---

#### 何为时间戳？

我们可以在 DFS 一棵树的过程中，维护一个全局的时间戳 $\textit{clock}$，每访问一个新的节点，就将 $\textit{clock}$ 加一。同时，记录进入节点 $x$ 时的时间戳 $\textit{in}[x]$，和离开（递归结束）这个节点时的时间戳 $\textit{out}[x]$。

#### 时间戳有什么性质？

根据 DFS 的性质，当我们递归以 $x$ 为根的子树时，设 $y$ 是 $x$ 的子孙节点，我们必须先递归完以 $y$ 为根的子树，之后才能递归完以 $x$ 为根的子树。

从时间戳上看，如果 $y$ 是 $x$ 的子孙节点，那么区间 $[\textit{in}[y],\textit{out}[y]]$ 必然被区间 $[\textit{in}[x],\textit{out}[x]]$ 所包含。

反之，如果区间 $[\textit{in}[y],\textit{out}[y]]$ 被区间 $[\textit{in}[x],\textit{out}[x]]$ 所包含，那么 $y$ 必然是 $x$ 的子孙节点（换句话说 $x$ 是 $y$ 的祖先节点）。因此我们可以通过

$$
\textit{in}[x]<\textit{in}[y]\le\textit{out}[y]\le\textit{out}[x]
$$

来判断 $x$ 是否为 $y$ 的祖先节点，由于 $\textit{in}[y]\le\textit{out}[y]$ 恒成立，上式可以简化为

$$
\textit{in}[x]<\textit{in}[y]\le\textit{out}[x]
$$

---

回到本题。由于需要求出子树的异或和，不妨以 $0$ 为根，DFS 这棵树，在求出时间戳的同时，求出每棵以 $x$ 为根的子树的异或和 $\textit{xor}[x]$。

由于 $n$ 比较小，我们可以用 $\mathcal{O}(n^2)$ 的时间枚举要删除的两条边 $x_1\text{-}y_1$ 和 $x_2\text{-}y_2$，并假设 $x$ 是 $y$ 的父节点，这会产生以下三种情况：

1. 删除的两条边在同一棵子树内，且 $y_1$ 是 $x_2$ 的祖先节点（或重合）。
   
   如下图所示，这三个连通块的异或和分别为 $\textit{xor}[y_2]$、$\textit{xor}[y_1]\oplus\textit{xor}[y_2]$ 和 $\textit{xor}[0]\oplus\textit{xor}[y_1]$（$\oplus$ 表示异或运算）。
   
   ![same.png](https://pic.leetcode-cn.com/1656215504-THYbIW-same.png)
   
2. 删除的两条边在同一棵子树内，且 $y_2$ 是 $x_1$ 的祖先节点（或重合）。
   
   同上，这三个连通块的异或和分别为 $\textit{xor}[y_1]$、$\textit{xor}[y_2]\oplus\textit{xor}[y_1]$ 和 $\textit{xor}[0]\oplus\textit{xor}[y_2]$。
   
3. 删除的两条边分别属于两棵不相交的子树。
   
   如下图所示，这三个连通块的异或和分别为 $\textit{xor}[y_1]$、$\textit{xor}[y_2]$ 和 $\textit{xor}[0]\oplus\textit{xor}[y_1]\oplus\textit{xor}[y_2]$。
   
   ![diff.png](https://pic.leetcode-cn.com/1656215343-RsPtkr-diff.png)

因此关键之处在于判断这两条边的关系，这可以用上文提到的时间戳的性质 $O(1)$ 地判断出来。

代码实现时，由于不知道 $\textit{edges}[i]$ 两个点的父子关系，枚举边的写法需要额外的判断。我们可以改为枚举不是根的两个点，删除这两个点到其父节点的边，这样代码更简洁，效率也略优于枚举边的写法。

```Python [sol-Python3]
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
            in_[x] = clock
            xor[x] = nums[x]
            for y in g[x]:
                if y != fa:
                    dfs(y, x)
                    xor[x] ^= xor[y]
            out[x] = clock
        dfs(0, -1)

        def is_ancestor(x: int, y: int) -> bool:
            return in_[x] < in_[y] <= out[x]

        ans = inf
        for i in range(2, n):
            for j in range(1, i):
                if is_ancestor(i, j):  # i 是 j 的祖先节点
                    x, y, z = xor[j], xor[i] ^ xor[j], xor[0] ^ xor[i]
                elif is_ancestor(j, i):  # j 是 i 的祖先节点
                    x, y, z = xor[i], xor[i] ^ xor[j], xor[0] ^ xor[j]
                else:  # 删除的两条边分别属于两棵不相交的子树
                    x, y, z = xor[i], xor[j], xor[0] ^ xor[i] ^ xor[j]
                ans = min(ans, max(x, y, z) - min(x, y, z))
                if ans == 0:
                    return 0  # 提前退出
        return ans
```

```java [sol-Java]
class Solution {
    private List<Integer>[] g;
    private int[] nums, xor, in, out;
    private int clock;

    public int minimumScore(int[] nums, int[][] edges) {
        int n = nums.length;
        g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }
        this.nums = nums;
        xor = new int[n];
        in = new int[n];
        out = new int[n];
        dfs(0, -1);

        int ans = Integer.MAX_VALUE;
        for (int i = 2; i < n; i++) {
            for (int j = 1; j < i; j++) {
                int x, y, z;
                if (isAncestor(i, j)) { // i 是 j 的祖先节点
                    x = xor[j];
                    y = xor[i] ^ x;
                    z = xor[0] ^ xor[i];
                } else if (isAncestor(j, i)) { // j 是 i 的祖先节点
                    x = xor[i];
                    y = xor[j] ^ x;
                    z = xor[0] ^ xor[j];
                } else { // 删除的两条边分别属于两棵不相交的子树
                    x = xor[i];
                    y = xor[j];
                    z = xor[0] ^ x ^ y;
                }
                ans = Math.min(ans, Math.max(Math.max(x, y), z) - Math.min(Math.min(x, y), z));
                if (ans == 0) {
                    return 0; // 提前退出
                }
            }
        }
        return ans;
    }

    private void dfs(int x, int fa) {
        in[x] = ++clock;
        xor[x] = nums[x];
        for (int y : g[x]) {
            if (y != fa) {
                dfs(y, x);
                xor[x] ^= xor[y];
            }
        }
        out[x] = clock;
    }

    private boolean isAncestor(int x, int y) {
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
            in[x] = ++clock;
            xr[x] = nums[x];
            for (int y : g[x]) {
                if (y != fa) {
                    dfs(y, x);
                    xr[x] ^= xr[y];
                }
            }
            out[x] = clock;
        };
        dfs(0, -1);

        auto is_ancestor = [&](int x, int y) -> bool { return in[x] < in[y] && in[y] <= out[x]; };

        int ans = INT_MAX;
        for (int i = 2, x, y, z; i < n; i++) {
            for (int j = 1; j < i; j++) {
                if (is_ancestor(i, j)) x = xr[j], y = xr[i] ^ x, z = xr[0] ^ xr[i]; // i 是 j 的祖先节点
                else if (is_ancestor(j, i)) x = xr[i], y = xr[j] ^ x, z = xr[0] ^ xr[j]; // j 是 i 的祖先节点
                else x = xr[i], y = xr[j], z = xr[0] ^ x ^ y; // 删除的两条边分别属于两棵不相交的子树
                ans = min(ans, max({x, y, z}) - min({x, y, z}));
                if (ans == 0) return 0; // 提前退出
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
		in[x] = clock
		xor[x] = nums[x]
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x)
				xor[x] ^= xor[y]
			}
		}
		out[x] = clock
	}
	dfs(0, -1)
	isAncestor := func(x, y int) bool { return in[x] < in[y] && in[y] <= out[x] }

	ans := math.MaxInt32
	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			var x, y, z int
			if isAncestor(i, j) { // i 是 j 的祖先节点
				x, y, z = xor[j], xor[i]^xor[j], xor[0]^xor[i]
			} else if isAncestor(j, i) { // j 是 i 的祖先节点
				x, y, z = xor[i], xor[i]^xor[j], xor[0]^xor[j]
			} else { // 删除的两条边分别属于两棵不相交的子树
				x, y, z = xor[i], xor[j], xor[0]^xor[i]^xor[j]
			}
			ans = min(ans, max(max(x, y), z)-min(min(x, y), z))
			if ans == 0 {
				return 0 // 提前退出
			}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
