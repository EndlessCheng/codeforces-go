[视频讲解](https://www.bilibili.com/video/BV1cV4y157BY) 第四题。

## 提示 1

枚举连通块的个数 $i$，则删除的边数为 $i-1$。

设 $\textit{total}$ 为整棵树的点权和（即 $\textit{nums}$ 的元素和），如果 $\textit{total}$ 能被 $i$ 整除（$i$ 是 $\textit{total}$ 的因子），那么每个连通块的点权和都应等于 $\dfrac{\textit{total}}{i}$，记作 $\textit{target}$。

什么样的边可以删除呢？

## 提示 2

如果一条边左右两侧的点权和都是 $\textit{target}$ 的倍数，那么这条边就可以删除。由于要使删除的边数最多，这条边**必须**删除。

由于 $\textit{total}$ 可以被 $\textit{target}$ 整除，我们只需要看一侧的点权和是否为 $\textit{target}$ 的倍数。

换言之，可以从任意点出发 DFS，只要发现子树的点权和是 $\textit{target}$ 的倍数，就说明子树到上面父节点的这条边是可以删除的。

具体来说，不妨以 $0$ 为根，DFS 这棵树，统计子树的点权和：

- 如果点权和超过 $\textit{target}$，说明当前删边方案不合法，返回 $-1$。
- 如果点权和等于 $\textit{target}$，这条边必须删除，返回 $0$。
- 如果点权和小于 $\textit{target}$，尚未找到一个完整的连通块，返回点权和。

如果 DFS 最终没有返回 $-1$，则当前删边方案合法。

如果我们从大到小枚举连通块的个数，则此时删除的边数是最多的，直接返回 $i-1$。

## 答疑

**问**：为什么这样做可以保证分出**恰好** $i$ 个连通块？

**答**：第一，不会超过 $i$ 个连通块，因为我们的做法相当于用水杯接水，每次接满 $\textit{target}$ 水就换下一杯继续接水。总共就 $\textit{total}$ 的水，至多可以接 $i$ 杯水。

第二，不会低于 $i$ 个连通块，如果出现这样的情况，说明至少有一个连通块的点权和超过 $\textit{target}$，此时 DFS 会返回 $-1$。

## 优化

代码实现时，由于点权至少为 $mx=\max(\textit{nums})$，所以连通块的个数至多为 $\left\lfloor\dfrac{\textit{total}}{mx}\right\rfloor$。由于 $\left\lfloor\dfrac{\textit{total}}{mx}\right\rfloor\le n$，因此可以从 $\left\lfloor\dfrac{\textit{total}}{mx}\right\rfloor$ 开始枚举连通块的个数。

```py [sol-Python3]
class Solution:
    def componentValue(self, nums: List[int], edges: List[List[int]]) -> int:
        g = [[] for _ in nums]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        def dfs(x: int, fa: int) -> int:
            s = nums[x]
            for y in g[x]:
                if y != fa:
                    res = dfs(y, x)
                    if res < 0: return -1
                    s += res
            if s > target: return -1
            return s if s < target else 0

        total = sum(nums)
        for i in range(total // max(nums), 1, -1):
            if total % i == 0:
                target = total // i
                if dfs(0, -1) == 0: return i - 1
        return 0
```

```java [sol-Java]
class Solution {
    private List<Integer>[] g;
    private int[] nums;
    private int target;

    public int componentValue(int[] nums, int[][] edges) {
        g = new ArrayList[nums.length];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (var e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }
        this.nums = nums;

        var total = Arrays.stream(nums).sum();
        var max = Arrays.stream(nums).max().orElseThrow();
        for (var i = total / max; ; --i)
            if (total % i == 0) {
                target = total / i;
                if (dfs(0, -1) == 0) return i - 1;
            }
    }

    private int dfs(int x, int fa) {
        var sum = nums[x];
        for (var y : g[x])
            if (y != fa) {
                var res = dfs(y, x);
                if (res < 0) return -1;
                sum += res;
            }
        if (sum > target) return -1;
        return sum < target ? sum : 0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int componentValue(vector<int> &nums, vector<vector<int>> &edges) {
        vector<vector<int>> g(nums.size());
        for (auto &e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        int target;
        function<int(int, int)> dfs = [&](int x, int fa) {
            int sum = nums[x];
            for (int y : g[x])
                if (y != fa) {
                    int res = dfs(y, x);
                    if (res < 0) return -1;
                    sum += res;
                }
            if (sum > target) return -1;
            return sum < target ? sum : 0;
        };

        int total = accumulate(nums.begin(), nums.end(), 0);
        int mx = *max_element(nums.begin(), nums.end());
        for (int i = total / mx;; --i)
            if (total % i == 0) {
                target = total / i;
                if (dfs(0, -1) == 0) return i - 1;
            }
    }
};
```

```go [sol-Go]
func componentValue(nums []int, edges [][]int) int {
	g := make([][]int, len(nums))
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var target int
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		sum := nums[x]
		for _, y := range g[x] {
			if y != fa {
				res := dfs(y, x)
				if res < 0 {
					return -1
				}
				sum += res
			}
		}
		if sum > target {
			return -1
		}
		if sum == target {
			return 0
		}
		return sum
	}

	total, mx := 0, 0
	for _, x := range nums {
		total += x
		mx = max(mx, x)
	}
	for i := total / mx; ; i-- {
		if total%i == 0 {
			target = total / i
			if dfs(0, -1) == 0 {
				return i - 1
			}
		}
	}
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n\cdot d(s))$，其中 $n$ 为 $\textit{nums}$ 的长度，$s$ 为所有 $\textit{nums}[i]$ 的和，$d(s)$ 为 $s$ 的因子个数。根据本题的数据范围，$d(s)\le 240$，例如 $s=720720$ 时可以取到等号。
- 空间复杂度：$O(n)$。
