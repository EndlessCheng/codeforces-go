本题 [视频讲解](https://www.bilibili.com/video/BV1pW4y1r7xs) 已出炉，欢迎点赞三连~

---

#### 何为时间戳？

我们可以在 DFS 一棵树的过程中，维护一个全局的时间戳 $\textit{clock}$，每访问一个新的节点，就将 $\textit{clock}$ 加一。同时，记录进入节点 $x$ 时的时间戳 $\textit{in}[x]$，和离开（递归结束）这个节点时的时间戳 $\textit{out}[x]$。

#### 时间戳有什么性质？

根据 DFS 的性质，当我们递归以 $x$ 为根的子树时，设 $y$ 是 $x$ 的子孙节点，那么我们必须先递归完以 $y$ 为根的子树，之后才能递归完以 $x$ 为根的子树。

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

由于 $n$ 比较小，我们可以用 $O(n^2)$ 的时间枚举要删除的两条边，这会产生以下三种情况：

1. 删除的两条边在同一棵子树内，且 $y_1$ 是 $x_2$ 的祖先节点（或重合）。
    
   如下图所示，这三个连通块的异或和分别为 $\textit{xor}[y_2]$、$\textit{xor}[y_1]\oplus\textit{xor}[y_2]$ 和 $\textit{xor}[0]\oplus\textit{xor}[y_1]$（$\oplus$ 表示异或运算）。

  ![same.png](https://pic.leetcode-cn.com/1656215504-THYbIW-same.png)

2. 删除的两条边在同一棵子树内，且 $y_2$ 是 $x_1$ 的祖先节点（或重合）。
    
   同上，这三个连通块的异或和分别为 $\textit{xor}[y_1]$、$\textit{xor}[y_2]\oplus\textit{xor}[y_1]$ 和 $\textit{xor}[0]\oplus\textit{xor}[y_2]$。
   
3. 删除的两条边分别属于两棵不相交的子树。

   如下图所示，这三个连通块的异或和分别为 $\textit{xor}[y_1]$、$\textit{xor}[y_2]$ 和 $\textit{xor}[0]\oplus\textit{xor}[y_1]\oplus\textit{xor}[y_2]$。

   ![diff.png](https://pic.leetcode-cn.com/1656215343-RsPtkr-diff.png)

因此关键之处在于判断这两条边的关系，这可以用上文提到的时间戳的性质 $O(1)$ 地判断出来。

代码实现时，我们可以改为枚举不是根的两个点，删除这两个点及其父节点形成的边。

#### 复杂度分析

- 时间复杂度：$O(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

```Python [sol1-Python3]
class Solution:
    def minimumScore(self, nums: List[int], edges: List[List[int]]) -> int:
        n = len(nums)
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        xor, in_, out, clock = [0] * n, [0] * n, [0] * n, 0
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
                if ans == 0: return 0  # 提前退出
        return ans
```

```java [sol1-Java]
class Solution {
    List<Integer>[] g;
    int[] nums, xor, in, out;
    int clock;

    public int minimumScore(int[] nums, int[][] edges) {
        var n = nums.length;
        g = new ArrayList[n];
        for (var i = 0; i < n; i++) g[i] = new ArrayList<>();
        for (var e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }
        this.nums = nums;
        xor = new int[n];
        in = new int[n];
        out = new int[n];
        dfs(0, -1);

        var ans = Integer.MAX_VALUE;
        for (var i = 2; i < n; ++i)
            for (int j = 1; j < i; ++j) {
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
                if (ans == 0) return 0; // 提前退出
            }
        return ans;
    }

    void dfs(int x, int fa) {
        in[x] = ++clock;
        xor[x] = nums[x];
        for (var y : g[x])
            if (y != fa) {
                dfs(y, x);
                xor[x] ^= xor[y];
            }
        out[x] = clock;
    }

    boolean isAncestor(int x, int y) {
        return in[x] < in[y] && in[y] <= out[x];
    }
}
```

```C++ [sol1-C++]
class Solution {
public:
    int minimumScore(vector<int> &nums, vector<vector<int>> &edges) {
        int n = nums.size();
        vector<vector<int>> g(n);
        for (auto &e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        int xr[n], in[n], out[n], clock = 0;
        function<void(int, int)> dfs = [&](int x, int fa) {
            in[x] = ++clock;
            xr[x] = nums[x];
            for (int y : g[x])
                if (y != fa) {
                    dfs(y, x);
                    xr[x] ^= xr[y];
                }
            out[x] = clock;
        };
        dfs(0, -1);
        auto is_ancestor = [&](int x, int y) -> bool { return in[x] < in[y] && in[y] <= out[x]; };

        int ans = INT_MAX;
        for (int i = 2, x, y, z; i < n; ++i)
            for (int j = 1; j < i; ++j) {
                if (is_ancestor(i, j)) x = xr[j], y = xr[i] ^ x, z = xr[0] ^ xr[i]; // i 是 j 的祖先节点
                else if (is_ancestor(j, i)) x = xr[i], y = xr[j] ^ x, z = xr[0] ^ xr[j]; // j 是 i 的祖先节点
                else x = xr[i], y = xr[j], z = xr[0] ^ x ^ y; // 删除的两条边分别属于两棵不相交的子树
                ans = min(ans, max({x, y, z}) - min({x, y, z}));
                if (ans == 0) return 0; // 提前退出
            }
        return ans;
    }
};
```

```go [sol1-Go]
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

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
```
