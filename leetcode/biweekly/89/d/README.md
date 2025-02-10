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

[视频讲解](https://www.bilibili.com/video/BV1cV4y157BY) 第四题。

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
        for (int[] e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }
        this.nums = nums;

        int total = Arrays.stream(nums).sum();
        int max = Arrays.stream(nums).max().getAsInt();
        for (int i = total / max; ; i--) {
            if (total % i == 0) {
                target = total / i;
                if (dfs(0, -1) == 0) {
                    return i - 1;
                }
            }
        }
    }

    private int dfs(int x, int fa) {
        int sum = nums[x];
        for (int y : g[x]) {
            if (y != fa) {
                var res = dfs(y, x);
                if (res < 0) return -1;
                sum += res;
            }
        }
        if (sum > target) {
            return -1;
        }
        return sum < target ? sum : 0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int componentValue(vector<int>& nums, vector<vector<int>>& edges) {
        vector<vector<int>> g(nums.size());
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        int target;
        auto dfs = [&](this auto&& dfs, int x, int fa) -> int {
            int sum = nums[x];
            for (int y : g[x]) {
                if (y != fa) {
                    int res = dfs(y, x);
                    if (res < 0) {
                        return -1;
                    }
                    sum += res;
                }
            }
            if (sum > target) {
                return -1;
            }
            return sum < target ? sum : 0;
        };

        int total = reduce(nums.begin(), nums.end());
        int mx = ranges::max(nums);
        for (int i = total / mx; ; i--) {
            if (total % i == 0) {
                target = total / i;
                if (dfs(0, -1) == 0) {
                    return i - 1;
                }
            }
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
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\cdot d(s))$，其中 $n$ 为 $\textit{nums}$ 的长度，$s$ 为所有 $\textit{nums}[i]$ 的和，$d(s)$ 为 $s$ 的因子个数。根据本题的数据范围，$d(s)\le 240$，例如 $s=720720$ 时可以取到等号。
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
