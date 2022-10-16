#### 提示 1

枚举连通块的个数 $i$，则删除的边数为 $i-1$。

设 $\textit{total}$ 为所有 $\textit{nums}[i]$ 的和，如果 $\textit{total}$ 能被 $i$ 整除（$i$ 是 $\textit{total}$ 的因子），那么每个连通块的价值都应等于 $\dfrac{\textit{total}}{i}$，记作 $\textit{target}$。

如何判定存在这些连通块呢？

#### 提示 2

如果一颗子树的价值等于 $\textit{target}$，那么可以将其作为一个连通块，和其父节点断开，换句话说，它对其祖先节点的价值贡献是 $0$。

DFS 这棵树，统计子树的价值：

- 如果价值超过 $\textit{target}$，那么当前删边方案不合法，返回 $-1$。
- 如果价值等于 $\textit{target}$，找到了一个连通块，和其父节点断开，返回 $0$。
- 如果价值小于 $\textit{target}$，尚未找到一个完整的连通块，返回价值。

如果 DFS 完了没有返回 $-1$，则当前删边方案合法。如果从大到小枚举连通块的个数，则此时可以直接返回答案。

#### 优化

代码实现时，由于价值至少为 $\max(\textit{nums}[i])$，连通块的个数至多为 $\left\lfloor\dfrac{\textit{total}}{\max(\textit{nums}[i])}\right\rfloor$。因此若 $\left\lfloor\dfrac{\textit{total}}{\max(\textit{nums}[i])}\right\rfloor<n$，则可以从 $\left\lfloor\dfrac{\textit{total}}{\max(\textit{nums}[i])}\right\rfloor$ 开始枚举连通块的个数。

```py [sol1-Python3]
class Solution:
    def componentValue(self, nums: List[int], edges: List[List[int]]) -> int:
        n = len(nums)
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        def dfs(x: int, fa: int) -> int:
            s = nums[x]  # 价值
            for y in g[x]:
                if y != fa:
                    res = dfs(y, x)
                    if res < 0: return -1
                    s += res
            if s > target: return -1
            return s if s < target else 0

        total = sum(nums)
        for i in range(min(n, total // max(nums)), 1, -1):
            if total % i == 0:
                target = total // i
                if dfs(0, -1) == 0: return i - 1
        return 0
```

```java [sol1-Java]
class Solution {
    private List<Integer>[] g;
    private int[] nums;
    private int target;

    public int componentValue(int[] nums, int[][] edges) {
        var n = nums.length;
        g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (var e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }
        this.nums = nums;

        var total = Arrays.stream(nums).sum();
        var max = Arrays.stream(nums).max().orElseThrow();
        for (var i = Math.min(n, total / max); ; --i)
            if (total % i == 0) {
                target = total / i;
                if (dfs(0, -1) == 0) return i - 1;
            }
    }

    private int dfs(int x, int fa) {
        var sum = nums[x]; // 价值
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

```cpp [sol1-C++]
class Solution {
public:
    int componentValue(vector<int> &nums, vector<vector<int>> &edges) {
        int n = nums.size(), target;
        vector<vector<int>> g(n);
        for (auto &e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        function<int(int, int)> dfs = [&](int x, int fa) {
            int sum = nums[x]; // 价值
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
        for (int i = min(n, total / mx);; --i)
            if (total % i == 0) {
                target = total / i;
                if (dfs(0, -1) == 0) return i - 1;
            }
    }
};
```

```go [sol1-Go]
func componentValue(nums []int, edges [][]int) int {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var target int
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		sum := nums[x] // 价值
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
	for i := min(n, total/mx); ; i-- {
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

- 时间复杂度：$O(n\cdot d(s))$，其中 $n$ 为 $\textit{nums}$ 的长度，$s$ 为所有 $\textit{nums}[i]$ 的和，$d(s)$ 为 $s$ 的因子个数。根据本题的数据范围，$d(s)\le 240$，$s=720720$ 时取等号。
- 空间复杂度：$O(n)$。当树是一条链时，递归的深度最大为 $n$，需要的栈空间为 $O(n)$。
