由于从左上角出发的小朋友只能移动 $n-1$ 次，所以他的走法有且仅有一种：主对角线。其余走法一定会超过 $n-1$ 步。

对于从右上角出发的小朋友，他不能穿过主对角线走到另一侧（不然就没法走到右下角），且同一个格子的水果不能重复收集。于是问题变成：

- 从右上角 $(0,n-1)$ 出发，在不访问主对角线的情况下，走到 $(n-2,n-1)$，也就是右下角的上面那个格子，所能收集到的水果总数的最大值。

做法类似 [931. 下降路径最小和](https://leetcode.cn/problems/minimum-falling-path-sum/)，请看 [我的题解](https://leetcode.cn/problems/minimum-falling-path-sum/solutions/2341851/cong-di-gui-dao-di-tui-jiao-ni-yi-bu-bu-2cwkb/)。

对于从左下角出发的小朋友，我们可以把矩阵按照主对角线翻转，就可以复用同一套代码逻辑了。

代码实现时，由于我们是倒着走的（为了方便翻译成递推），小朋友不能一直往左上走，不然没法走到右上角。所以要限制小朋友不能太靠左，即保证 $j\ge n-1-i$。这是因为从 $(0,n-1)$ 往左下的这条线满足 $i+j=n-1$，不能越过这条线，即 $i+j\ge n-1$，也就是 $j\ge n-1-i$。

本题由于元素值均非负，可以在出界时返回 $0$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1uzBxYoEJC/?t=13m12s)，欢迎点赞关注~

## 一、记忆化搜索

```py [sol-Python3]
class Solution:
    def maxCollectedFruits(self, fruits: List[List[int]]) -> int:
        n = len(fruits)

        @cache
        def dfs(i: int, j: int) -> int:
            if not (n - 1 - i <= j < n):
                return -inf
            if i == 0:
                return fruits[i][j]
            return max(dfs(i - 1, j - 1), dfs(i - 1, j), dfs(i - 1, j + 1)) + fruits[i][j]

        # 从 (0, 0) 出发的小朋友
        ans = sum(row[i] for i, row in enumerate(fruits))

        # 从 (0, n - 1) 出发的小朋友（倒着走）
        ans += dfs(n - 2, n - 1)  # 从下往上走，方便 1:1 翻译成递推
        dfs.cache_clear()

        # 从 (n - 1, 0) 出发的小朋友（按照主对角线翻转，然后倒着走）
        fruits = list(zip(*fruits))  # 按照主对角线翻转
        return ans + dfs(n - 2, n - 1)
```

```java [sol-Java]
class Solution {
    public int maxCollectedFruits(int[][] fruits) {
        int n = fruits.length;
        int[][] memo = new int[n][n];
        for (int[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }

        int ans = 0;
        // 从 (0, 0) 出发的小朋友
        for (int i = 0; i < n; i++) {
            ans += fruits[i][i];
        }

        // 从 (0, n - 1) 出发的小朋友（倒着走）
        // 从下往上走，方便 1:1 翻译成递推
        ans += dfs(n - 2, n - 1, fruits, memo);

        // 从 (n - 1, 0) 出发的小朋友（按照主对角线翻转，然后倒着走）
        // 把下三角形中的数据填到上三角形中
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < i; j++) {
                fruits[j][i] = fruits[i][j];
            }
        }
        for (int[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }
        return ans + dfs(n - 2, n - 1, fruits, memo);
    }

    private int dfs(int i, int j, int[][] fruits, int[][] memo) {
        int n = fruits.length;
        if (j < n - 1 - i || j >= n) {
            return 0;
        }
        if (i == 0) {
            return fruits[i][j];
        }
        if (memo[i][j] != -1) { // 之前计算过
            return memo[i][j];
        }
        int res1 = dfs(i - 1, j - 1, fruits, memo);
        int res2 = dfs(i - 1, j, fruits, memo);
        int res3 = dfs(i - 1, j + 1, fruits, memo);
        int res = Math.max(Math.max(res1, res2), res3) + fruits[i][j];
        return memo[i][j] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxCollectedFruits(vector<vector<int>>& fruits) {
        int n = fruits.size();
        vector memo(n, vector<int>(n, -1)); // -1 表示没有计算过
        auto dfs = [&](this auto&& dfs, int i, int j) -> int {
            if (j < n - 1 - i || j >= n) {
                return 0;
            }
            if (i == 0) {
                return fruits[i][j];
            }
            int& res = memo[i][j]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }
            return res = max({dfs(i - 1, j - 1), dfs(i - 1, j), dfs(i - 1, j + 1)}) + fruits[i][j];
        };

        int ans = 0;
        // 从 (0, 0) 出发的小朋友
        for (int i = 0; i < n; i++) {
            ans += fruits[i][i];
        }

        // 从 (0, n - 1) 出发的小朋友（倒着走）
        ans += dfs(n - 2, n - 1); // 从下往上走，方便 1:1 翻译成递推

        // 从 (n - 1, 0) 出发的小朋友（按照主对角线翻转，然后倒着走）
        // 把下三角形中的数据填到上三角形中
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < i; j++) {
                fruits[j][i] = fruits[i][j];
            }
        }
        ranges::fill(memo, vector(n, -1));
        return ans + dfs(n - 2, n - 1);
    }
};
```

```go [sol-Go]
func maxCollectedFruits(fruits [][]int) (ans int) {
	n := len(fruits)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if j < n-1-i || j >= n {
			return 0
		}
		if i == 0 {
			return fruits[i][j]
		}
		p := &memo[i][j]
		if *p < 0 { // 没有计算过
			*p = max(dfs(i-1, j-1), dfs(i-1, j), dfs(i-1, j+1)) + fruits[i][j]
		}
		return *p
	}

	// 从 (0, 0) 出发的小朋友
	for i, row := range fruits {
		ans += row[i]
	}

	// 从 (0, n-1) 出发的小朋友（倒着走）
	ans += dfs(n-2, n-1) // 从下往上走，方便 1:1 翻译成递推

	// 从 (n-1, 0) 出发的小朋友（按照主对角线翻转，然后倒着走）
	// 把下三角形中的数据填到上三角形中
	for i := range fruits {
		for j := range i {
			fruits[j][i] = fruits[i][j]
		}
	}
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	return ans + dfs(n-2, n-1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{fruits}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

## 二、1:1 翻译成递推

由于起点是 $(0,n-1)$，即使每一步都往左下走，$i+j$ 也不会低于 $n-1$，所以 $j\ge n-1-i$。

由于终点是 $(n-2,n-1)$，即使从终点倒着，每一步都往左上走，$j$ 也始终大于 $i$。

所以 $j$ 可以从

$$
\max(n-1-i,i+1)
$$

开始枚举。

```py [sol-Python3]
class Solution:
    def maxCollectedFruits(self, fruits: List[List[int]]) -> int:
        def dp(fruits: List[List[int]]) -> int:
            n = len(fruits)
            f = [[0] * (n + 1) for _ in range(n - 1)]
            f[0][n - 1] = fruits[0][-1]
            for i in range(1, n - 1):
                for j in range(max(n - 1 - i, i + 1), n):
                    f[i][j] = max(f[i - 1][j - 1], f[i - 1][j], f[i - 1][j + 1]) + fruits[i][j]
            return f[-1][n - 1]

        ans = sum(row[i] for i, row in enumerate(fruits))
        return ans + dp(fruits) + dp(list(zip(*fruits)))
```

```java [sol-Java]
class Solution {
    public int maxCollectedFruits(int[][] fruits) {
        int n = fruits.length;
        int ans = 0;
        // 从 (0, 0) 出发的小朋友
        for (int i = 0; i < n; i++) {
            ans += fruits[i][i];
        }

        // 从 (0, n - 1) 出发的小朋友
        ans += dp(fruits);

        // 从 (n - 1, 0) 出发的小朋友（按照主对角线翻转）
        // 把下三角形中的数据填到上三角形中
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < i; j++) {
                fruits[j][i] = fruits[i][j];
            }
        }
        return ans + dp(fruits);
    }

    int dp(int[][] fruits) {
        int n = fruits.length;
        int[][] f = new int[n - 1][n + 1];
        f[0][n - 1] = fruits[0][n - 1];
        for (int i = 1; i < n - 1; i++) {
            for (int j = Math.max(n - 1 - i, i + 1); j < n; j++) {
                f[i][j] = Math.max(Math.max(f[i - 1][j - 1], f[i - 1][j]), f[i - 1][j + 1]) + fruits[i][j];
            }
        }
        return f[n - 2][n - 1];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxCollectedFruits(vector<vector<int>>& fruits) {
        int n = fruits.size();
        auto dp = [&]() -> int {
            vector f(n - 1, vector<int>(n + 1));
            f[0][n - 1] = fruits[0][n - 1];
            for (int i = 1; i < n - 1; i++) {
                for (int j = max(n - 1 - i, i + 1); j < n; j++) {
                    f[i][j] = max({f[i - 1][j - 1], f[i - 1][j], f[i - 1][j + 1]}) + fruits[i][j];
                }
            }
            return f[n - 2][n - 1];
        };

        int ans = 0;
        // 从 (0, 0) 出发的小朋友
        for (int i = 0; i < n; i++) {
            ans += fruits[i][i];
        }

        // 从 (0, n - 1) 出发的小朋友
        ans += dp();

        // 从 (n - 1, 0) 出发的小朋友（按照主对角线翻转）
        // 把下三角形中的数据填到上三角形中
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < i; j++) {
                fruits[j][i] = fruits[i][j];
            }
        }
        return ans + dp();
    }
};
```

```go [sol-Go]
func maxCollectedFruits(fruits [][]int) (ans int) {
	n := len(fruits)
	f := make([][]int, n-1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	dp := func() int {
		f[0][n-1] = fruits[0][n-1]
		for i := 1; i < n-1; i++ {
			for j := max(n-1-i, i+1); j < n; j++ {
				f[i][j] = max(f[i-1][j-1], f[i-1][j], f[i-1][j+1]) + fruits[i][j]
			}
		}
		return f[n-2][n-1]
	}

	// 从 (0, 0) 出发的小朋友
	for i, row := range fruits {
		ans += row[i]
	}

	// 从 (0, n - 1) 出发的小朋友
	ans += dp()

	// 从 (n - 1, 0) 出发的小朋友（按照主对角线翻转）
	// 把下三角形中的数据填到上三角形中
	for i := range fruits {
		for j := range i {
			fruits[j][i] = fruits[i][j]
		}
	}
	return ans + dp()
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{fruits}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

**注**：用滚动数组可以做到 $\mathcal{O}(n)$ 的空间复杂度。也可以原地修改 $\textit{fruits}$，做到 $\mathcal{O}(1)$ 空间复杂度。

## 专题训练

见下面动态规划题单的「**二、网格图 DP**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
