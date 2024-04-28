## 方法一：记忆化搜索

**前置知识**：[动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含如何把记忆化搜索 1:1 翻译成递推的技巧。

看示例 3，$\textit{zero}=3,\ \textit{one}=3,\ \textit{limit}=2$。

先解释 $\textit{limit}$，意思是数组中至多有连续 $\textit{limit}$ 个 $0$，且至多有连续 $\textit{limit}$ 个 $1$。

考虑稳定数组的最后一个位置填 $0$ 还是 $1$：

- 填 $0$，问题变成剩下 $2$ 个 $0$ 和 $3$ 个 $1$ 怎么填。
- 填 $1$，问题变成剩下 $3$ 个 $0$ 和 $2$ 个 $1$ 怎么填。
- 这两个都是和原问题相似的子问题。

看上去，定义 $\textit{dfs}(i,j)$ 表示用 $i$ 个 $0$ 和 $j$ 个 $1$ 构造稳定数组的方案数？但这样定义不方便计算 $\textit{limit}$ 带来的影响。

改成定义 $\textit{dfs}(i,j,k)$ 表示用 $i$ 个 $0$ 和 $j$ 个 $1$ 构造稳定数组的方案数，其中第 $i+j$ 个位置要填 $k$，其中 $k$ 为 $0$ 或 $1$。

如果 $k=0$，考虑第 $i+j-1$ 个位置要填什么：

- 填 $0$，问题变成 $\textit{dfs}(i-1,j,0)$。
- 填 $1$，问题变成 $\textit{dfs}(i-1,j,1)$。

但是，$\textit{dfs}(i-1,j,0)$ 包含了「最后连续 $\textit{limit}$ 个位置填 $0$」的方案，如果在这个方案末尾再加一个 $0$，就有连续 $\textit{limit}+1$ 个 $0$ 了，这是不合法的。

$\textit{dfs}(i-1,j,0)$ 中的「最后连续 $\textit{limit}$ 个位置填 $0$」的方案有多少个？

因为 $\textit{dfs}$ 的定义是稳定数组的方案数，只包含合法方案，所以在最后连续 $\textit{limit}$ 个位置填 $0$ 的情况下，倒数第 $\textit{limit}+1$ 个位置一定要填 $1$，这有 $\textit{dfs}(i-\textit{limit}-1,j,1)$ 种方案。对于 $\textit{dfs}(i,j,0)$ 来说，这 $\textit{dfs}(i-\textit{limit}-1,j,1)$ 个方案就是不合法方案了，要减去，得

$$
\textit{dfs}(i,j,0) = \textit{dfs}(i-1,j,0) + \textit{dfs}(i-1,j,1) - \textit{dfs}(i-\textit{limit}-1,j,1)
$$

同理得

$$
\textit{dfs}(i,j,1) = \textit{dfs}(i,j-1,0) + \textit{dfs}(i,j-1,1) - \textit{dfs}(i,j-\textit{limit}-1,0)
$$

递归边界 1：如果 $i<0$ 或者 $j<0$，返回 $0$。也可以在计算 $\textit{dfs}(i-\textit{limit}-1,j,1)$ 前判断 $i>\textit{limit}$，在计算 $\textit{dfs}(i,j-\textit{limit}-1,0)$ 前判断 $j>\textit{limit}$。

递归边界 2：如果 $i=0$，那么当 $k=1$ 且 $j\le \textit{limit}$ 的情况下返回 $1$，否则返回 $0$；如果 $j=0$，那么当 $k=0$ 且 $i\le \textit{limit}$ 的情况下返回 $1$，否则返回 $0$。

递归入口：$\textit{dfs}(\textit{zero},\textit{one},0)+\textit{dfs}(\textit{zero},\textit{one},1)$，即答案。

请看 [视频讲解](https://www.bilibili.com/video/BV16t421c7GB/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def numberOfStableArrays(self, zero: int, one: int, limit: int) -> int:
        MOD = 1_000_000_007
        @cache
        def dfs(i: int, j: int, k: int) -> int:
            if i == 0:
                return 1 if k == 1 and j <= limit else 0
            if j == 0:
                return 1 if k == 0 and i <= limit else 0
            if k == 0:
                return (dfs(i - 1, j, 0) + dfs(i - 1, j, 1) - (dfs(i - limit - 1, j, 1) if i > limit else 0)) % MOD
            else:  # else 可以去掉，这里仅仅是为了代码对齐
                return (dfs(i, j - 1, 0) + dfs(i, j - 1, 1) - (dfs(i, j - limit - 1, 0) if j > limit else 0)) % MOD
        ans = (dfs(zero, one, 0) + dfs(zero, one, 1)) % MOD
        dfs.cache_clear()  # 防止爆内存
        return ans
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int numberOfStableArrays(int zero, int one, int limit) {
        int[][][] memo = new int[zero + 1][one + 1][2];
        for (int[][] m : memo) {
            for (int[] m2 : m) {
                Arrays.fill(m2, -1); // -1 表示没有计算过
            }
        }
        return (dfs(zero, one, 0, limit, memo) + dfs(zero, one, 1, limit, memo)) % MOD;
    }

    private int dfs(int i, int j, int k, int limit, int[][][] memo) {
        if (i == 0) { // 递归边界
            return k == 1 && j <= limit ? 1 : 0;
        }
        if (j == 0) { // 递归边界
            return k == 0 && i <= limit ? 1 : 0;
        }
        if (memo[i][j][k] != -1) { // 之前计算过
            return memo[i][j][k];
        }
        if (k == 0) {
            // + MOD 保证答案非负
            memo[i][j][k] = (int) (((long) dfs(i - 1, j, 0, limit, memo) + dfs(i - 1, j, 1, limit, memo) +
                    (i > limit ? MOD - dfs(i - limit - 1, j, 1, limit, memo) : 0)) % MOD);
        } else {
            memo[i][j][k] = (int) (((long) dfs(i, j - 1, 0, limit, memo) + dfs(i, j - 1, 1, limit, memo) +
                    (j > limit ? MOD - dfs(i, j - limit - 1, 0, limit, memo) : 0)) % MOD);
        }
        return memo[i][j][k];
    }
}
```

```cpp [sol-C++]
class Solution {
    int MOD = 1'000'000'007;
    vector<vector<array<int, 2>>> memo;

    int dfs(int i, int j, int k, int limit) {
        if (i == 0) { // 递归边界
            return k == 1 && j <= limit;
        }
        if (j == 0) { // 递归边界
            return k == 0 && i <= limit;
        }
        int& res = memo[i][j][k]; // 注意这里是引用
        if (res != -1) { // 之前计算过
            return res;
        }
        if (k == 0) {
            // + MOD 保证答案非负
            res = ((long long) dfs(i - 1, j, 0, limit) + dfs(i - 1, j, 1, limit) +
                   (i > limit ? MOD - dfs(i - limit - 1, j, 1, limit) : 0)) % MOD;
        } else {
            res = ((long long) dfs(i, j - 1, 0, limit) + dfs(i, j - 1, 1, limit) +
                   (j > limit ? MOD - dfs(i, j - limit - 1, 0, limit) : 0)) % MOD;
        }
        return res;
    }

public:
    int numberOfStableArrays(int zero, int one, int limit) {
        // -1 表示没有计算过
        memo.resize(zero + 1, vector<array<int, 2>>(one + 1, {-1, -1}));
        return (dfs(zero, one, 0, limit) + dfs(zero, one, 1, limit)) % MOD;
    }
};
```

```go [sol-Go]
func numberOfStableArrays(zero, one, limit int) int {
	const mod = 1_000_000_007
	memo := make([][][2]int, zero+1)
	for i := range memo {
		memo[i] = make([][2]int, one+1)
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) (res int) {
		if i == 0 { // 递归边界
			if k == 1 && j <= limit {
				return 1
			}
			return
		}
		if j == 0 { // 递归边界
			if k == 0 && i <= limit {
				return 1
			}
			return
		}
		p := &memo[i][j][k]
		if *p != -1 { // 之前计算过
			return *p
		}
		if k == 0 {
			// +mod 保证答案非负
			res = (dfs(i-1, j, 0) + dfs(i-1, j, 1)) % mod
			if i > limit {
				res = (res - dfs(i-limit-1, j, 1) + mod) % mod
			}
		} else {
			res = (dfs(i, j-1, 0) + dfs(i, j-1, 1)) % mod
			if j > limit {
				res = (res - dfs(i, j-limit-1, 0) + mod) % mod
			}
		}
		*p = res // 记忆化
		return
	}
	return (dfs(zero, one, 0) + dfs(zero, one, 1)) % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\textit{zero}\cdot \textit{one})$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(\textit{zero}\cdot \textit{one})$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}(\textit{zero}\cdot \textit{one})$。
- 空间复杂度：$\mathcal{O}(\textit{zero}\cdot \textit{one})$。有多少个状态，$\textit{memo}$ 数组的大小就是多少。

## 方法二：递推

根据 [视频](https://www.bilibili.com/video/BV1Xj411K7oF/) 中讲的，把方法一 1:1 地翻译成递推。

定义 $f[i][j][k]$ 表示用 $i$ 个 $0$ 和 $j$ 个 $1$ 构造稳定数组的方案数，其中第 $i+j$ 个位置要填 $k$，其中 $k$ 为 $0$ 或 $1$。

状态转移方程：

$$
\begin{aligned}
&f[i][j][0] = f[i-1][j][0] + f[i-1][j][1] - f[i-\textit{limit}-1][j][1]\\
&f[i][j][1] = f[i][j-1][0] + f[i][j-1][1] - f[i][j-\textit{limit}-1][0]
\end{aligned}
$$

如果 $i\le \textit{limit}$ 则 $f[i-\textit{limit}-1][j][1]$ 视作 $0$，如果 $j\le \textit{limit}$ 则 $f[i][j-\textit{limit}-1][0]$ 视作 $0$。

初始值：$f[i][0][0] = f[0][j][1] = 1$，其中 $1\le i \le \min(\textit{limit}, \textit{zero}),\ 1\le j \le \min(\textit{limit}, \textit{one})$。翻译自递归边界。

答案：$f[\textit{zero}][\textit{one}][0] + f[\textit{zero}][\textit{one}][1]$。翻译自递归入口。

```py [sol-Python3]
class Solution:
    def numberOfStableArrays(self, zero: int, one: int, limit: int) -> int:
        MOD = 1_000_000_007
        f = [[[0, 0] for _ in range(one + 1)] for _ in range(zero + 1)]
        for i in range(1, min(limit, zero) + 1):
            f[i][0][0] = 1
        for j in range(1, min(limit, one) + 1):
            f[0][j][1] = 1
        for i in range(1, zero + 1):
            for j in range(1, one + 1):
                f[i][j][0] = (f[i - 1][j][0] + f[i - 1][j][1] - (f[i - limit - 1][j][1] if i > limit else 0)) % MOD
                f[i][j][1] = (f[i][j - 1][0] + f[i][j - 1][1] - (f[i][j - limit - 1][0] if j > limit else 0)) % MOD
        return sum(f[-1][-1]) % MOD
```

```java [sol-Java]
class Solution {
    public int numberOfStableArrays(int zero, int one, int limit) {
        final int MOD = 1_000_000_007;
        int[][][] f = new int[zero + 1][one + 1][2];
        for (int i = 1; i <= Math.min(limit, zero); i++) {
            f[i][0][0] = 1;
        }
        for (int j = 1; j <= Math.min(limit, one); j++) {
            f[0][j][1] = 1;
        }
        for (int i = 1; i <= zero; i++) {
            for (int j = 1; j <= one; j++) {
                // + MOD 保证答案非负
                f[i][j][0] = (int) (((long) f[i - 1][j][0] + f[i - 1][j][1] + (i > limit ? MOD - f[i - limit - 1][j][1] : 0)) % MOD);
                f[i][j][1] = (int) (((long) f[i][j - 1][0] + f[i][j - 1][1] + (j > limit ? MOD - f[i][j - limit - 1][0] : 0)) % MOD);
            }
        }
        return (f[zero][one][0] + f[zero][one][1]) % MOD;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfStableArrays(int zero, int one, int limit) {
        const int MOD = 1'000'000'007;
        vector<vector<array<int, 2>>> f(zero + 1, vector<array<int, 2>>(one + 1));
        for (int i = 1; i <= min(limit, zero); i++) {
            f[i][0][0] = 1;
        }
        for (int j = 1; j <= min(limit, one); j++) {
            f[0][j][1] = 1;
        }
        for (int i = 1; i <= zero; i++) {
            for (int j = 1; j <= one; j++) {
                // + MOD 保证答案非负
                f[i][j][0] = ((long long) f[i - 1][j][0] + f[i - 1][j][1] + (i > limit ? MOD - f[i - limit - 1][j][1] : 0)) % MOD;
                f[i][j][1] = ((long long) f[i][j - 1][0] + f[i][j - 1][1] + (j > limit ? MOD - f[i][j - limit - 1][0] : 0)) % MOD;
            }
        }
        return (f[zero][one][0] + f[zero][one][1]) % MOD;
    }
};
```

```go [sol-Go]
func numberOfStableArrays(zero, one, limit int) (ans int) {
	const mod = 1_000_000_007
	f := make([][][2]int, zero+1)
	for i := range f {
		f[i] = make([][2]int, one+1)
	}
	for i := 1; i <= min(limit, zero); i++ {
		f[i][0][0] = 1
	}
	for j := 1; j <= min(limit, one); j++ {
		f[0][j][1] = 1
	}
	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			f[i][j][0] = (f[i-1][j][0] + f[i-1][j][1]) % mod
			if i > limit {
				// + mod 保证答案非负
				f[i][j][0] = (f[i][j][0] - f[i-limit-1][j][1] + mod) % mod
			}
			f[i][j][1] = (f[i][j-1][0] + f[i][j-1][1]) % mod
			if j > limit {
				f[i][j][1] = (f[i][j][1] - f[i][j-limit-1][0] + mod) % mod
			}
		}
	}
	return (f[zero][one][0] + f[zero][one][1]) % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\textit{zero}\cdot \textit{one})$。
- 空间复杂度：$\mathcal{O}(\textit{zero}\cdot \textit{one})$。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
