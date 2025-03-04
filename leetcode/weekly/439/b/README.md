**前置题目**：[516. 最长回文子序列](https://leetcode.cn/problems/longest-palindromic-subsequence/)，请看 [区间 DP【基础算法精讲 22】](https://www.bilibili.com/video/BV1Gs4y1E7EU/)。

本题在 516 题的基础上，增加一个参数 $k$，表示剩余操作次数。

## 写法一：记忆化搜索

定义 $\textit{dfs}(i,j,k)$ 表示执行最多 $k$ 次操作后，子串 $s[i]$ 到 $s[j]$ 中的最长回文子序列的长度。

和 516 题一样，分类讨论：

- 不选 $s[i]$，那么问题变成执行最多 $k$ 次操作后，子串 $s[i+1]$ 到 $s[j]$ 中的最长回文子序列的长度，即 $\textit{dfs}(i+1,j,k)$。
- 不选 $s[j]$，那么问题变成执行最多 $k$ 次操作后，子串 $s[i]$ 到 $s[j-1]$ 中的最长回文子序列的长度，即 $\textit{dfs}(i,j-1,k)$。
- 操作，把 $s[i]$ 和 $s[j]$ 都变成一样的，设操作 $\textit{op}$ 次且 $\textit{op}\le k$，那么问题变成执行最多 $k-\textit{op}$ 次操作后，子串 $s[i+1]$ 到 $s[j-1]$ 中的最长回文子序列的长度，即 $\textit{dfs}(i+1,j-1,k-\textit{op})+2$。其中 $+2$ 是因为 $s[i]$ 和 $s[j]$ 变成相同的了。

$\textit{op}$ 怎么算？设 $d = |s[i]-s[j]|$，那么 $\textit{op}=\min(d,26-d)$，前者没有跨过 $\texttt{z}$ 到 $\texttt{a}$，后者跨过。

三种情况取最大值，得

$$
\textit{dfs}(i,j,k) = \max(\textit{dfs}(i+1,j,k), \textit{dfs}(i,j-1,k), \textit{dfs}(i+1,j-1,k-\textit{op})+2)
$$

**递归边界**：$\textit{dfs}(i+1,i,k) = 0,\ \textit{dfs}(i,i,k) = 1$。注意只有一个字母的子串一定是回文串（回文子序列）。

**递归入口**：$\textit{dfs}(0,n-1,k)$，即答案。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1QP9bY3EL6/?t=8m38s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def longestPalindromicSubsequence(self, s: str, k: int) -> int:
        s = list(map(ord, s))  # 避免在 dfs 中频繁计算 ord
        n = len(s)
        @cache
        def dfs(i: int, j: int, k: int) -> int:
            if i >= j:
                return j - i + 1  # i=j+1 时返回 0，i=j 时返回 1
            res = max(dfs(i + 1, j, k), dfs(i, j - 1, k))
            d = abs(s[i] - s[j])
            op = min(d, 26 - d)
            if op <= k:
                res = max(res, dfs(i + 1, j - 1, k - op) + 2)
            return res
        ans = dfs(0, n - 1, k)
        dfs.cache_clear()  # 避免超出内存限制
        return ans
```

```java [sol-Java]
class Solution {
    public int longestPalindromicSubsequence(String S, int k) {
        char[] s = S.toCharArray();
        int n = s.length;
        int[][][] memo = new int[n][n][k + 1];
        for (int[][] mat : memo) {
            for (int[] row : mat) {
                Arrays.fill(row, -1);
            }
        }
        return dfs(0, n - 1, k, s, memo);
    }

    private int dfs(int i, int j, int k, char[] s, int[][][] memo) {
        if (i >= j) {
            return j - i + 1; // i=j+1 时返回 0，i=j 时返回 1
        }
        if (memo[i][j][k] != -1) {
            return memo[i][j][k];
        }
        int res = Math.max(dfs(i + 1, j, k, s, memo), dfs(i, j - 1, k, s, memo));
        int d = Math.abs(s[i] - s[j]);
        int op = Math.min(d, 26 - d);
        if (op <= k) {
            res = Math.max(res, dfs(i + 1, j - 1, k - op, s, memo) + 2);
        }
        return memo[i][j][k] = res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestPalindromicSubsequence(string s, int k) {
        int n = s.size();
        vector memo(n, vector(n, vector<int>(k + 1, -1)));
        auto dfs = [&](this auto&& dfs, int i, int j, int k) -> int {
            if (i >= j) {
                return j - i + 1; // i=j+1 时返回 0，i=j 时返回 1
            }
            int& res = memo[i][j][k]; // 注意这里是引用
            if (res != -1) {
                return res;
            }
            res = max(dfs(i + 1, j, k), dfs(i, j - 1, k));
            int d = abs(s[i] - s[j]);
            int op = min(d, 26 - d);
            if (op <= k) {
                res = max(res, dfs(i + 1, j - 1, k - op) + 2);
            }
            return res;
        };
        return dfs(0, n - 1, k);
    }
};
```

```go [sol-Go]
func longestPalindromicSubsequence(s string, k int) int {
	n := len(s)
	memo := make([][][]int, n)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, k+1)
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if i >= j {
			return j - i + 1 // i=j+1 时返回 0，i=j 时返回 1
		}
		p := &memo[i][j][k]
		if *p != -1 {
			return *p
		}
		res := max(dfs(i+1, j, k), dfs(i, j-1, k))
		d := abs(int(s[i]) - int(s[j]))
		op := min(d, 26-d)
		if op <= k {
			res = max(res, dfs(i+1, j-1, k-op)+2)
		}
		*p = res
		return res
	}
	return dfs(0, n-1, k)
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2k)$，其中 $n$ 是 $s$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n^2k)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(n^2k)$。
- 空间复杂度：$\mathcal{O}(n^2k)$。保存多少状态，就需要多少空间。

## 写法二：递推

和 516 题一样，1:1 翻译成递推。具体请看 [区间 DP【基础算法精讲 22】](https://www.bilibili.com/video/BV1Gs4y1E7EU/)。

为了提高访问缓存的效率，把 $k$ 放到第一个维度。这样第一个维度我们只会访问 $f[k-13],f[k-12],\ldots,f[k]$，减少 cache miss。

此外还有一个**优化**：如果 $s$ 可以在 $k$ 次操作内变成回文串，那么直接返回 $n$。

```py [sol-Python3]
class Solution:
    def longestPalindromicSubsequence(self, s: str, K: int) -> int:
        s = list(map(ord, s))  # 避免频繁计算 ord
        n = len(s)
        cnt = 0
        for i in range(n // 2):
            d = abs(s[i] - s[-1 - i])
            cnt += min(d, 26 - d)
        if cnt <= K:
            return n

        f = [[[0] * n for _ in range(n)] for _ in range(K + 1)]
        for k in range(K + 1):
            for i in range(n - 1, -1, -1):
                f[k][i][i] = 1
                for j in range(i + 1, n):
                    res = max(f[k][i + 1][j], f[k][i][j - 1])
                    d = abs(s[i] - s[j])
                    op = min(d, 26 - d)
                    if op <= k:
                        res = max(res, f[k - op][i + 1][j - 1] + 2)
                    f[k][i][j] = res
        return f[K][0][-1]
```

```java [sol-Java]
class Solution {
    public int longestPalindromicSubsequence(String S, int K) {
        char[] s = S.toCharArray();
        int n = s.length;
        int cnt = 0;
        for (int i = 0; i < n / 2; i++) {
            int d = Math.abs(s[i] - s[n - 1 - i]);
            cnt += Math.min(d, 26 - d);
        }
        if (cnt <= K) {
            return n;
        }

        int[][][] f = new int[K + 1][n][n];
        for (int k = 0; k <= K; k++) {
            for (int i = n - 1; i >= 0; i--) {
                f[k][i][i] = 1;
                for (int j = i + 1; j < n; j++) {
                    int res = Math.max(f[k][i + 1][j], f[k][i][j - 1]);
                    int d = Math.abs(s[i] - s[j]);
                    int op = Math.min(d, 26 - d);
                    if (op <= k) {
                        res = Math.max(res, f[k - op][i + 1][j - 1] + 2);
                    }
                    f[k][i][j] = res;
                }
            }
        }
        return f[K][0][n - 1];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestPalindromicSubsequence(string s, int K) {
        int n = s.size();
        int cnt = 0;
        for (int i = 0; i < n / 2; i++) {
            int d = abs(s[i] - s[n - 1 - i]);
            cnt += min(d, 26 - d);
        }
        if (cnt <= K) {
            return n;
        }

        vector f(K + 1, vector(n, vector<int>(n)));
        for (int k = 0; k <= K; k++) {
            for (int i = n - 1; i >= 0; i--) {
                f[k][i][i] = 1;
                for (int j = i + 1; j < n; j++) {
                    int res = max(f[k][i + 1][j], f[k][i][j - 1]);
                    int d = abs(s[i] - s[j]);
                    int op = min(d, 26 - d);
                    if (op <= k) {
                        res = max(res, f[k - op][i + 1][j - 1] + 2);
                    }
                    f[k][i][j] = res;
                }
            }
        }
        return f[K][0][n - 1];
    }
};
```

```go [sol-Go]
func longestPalindromicSubsequence(s string, K int) int {
	n := len(s)
	cnt := 0
	for i := range n / 2 {
		d := abs(int(s[i]) - int(s[n-1-i]))
		cnt += min(d, 26-d)
	}
	if cnt <= K {
		return n
	}

	f := make([][][]int, K+1)
	for k := range f {
		f[k] = make([][]int, n)
		for j := range f[k] {
			f[k][j] = make([]int, n)
		}
		for i := n - 1; i >= 0; i-- {
			f[k][i][i] = 1
			for j := i + 1; j < n; j++ {
				res := max(f[k][i+1][j], f[k][i][j-1])
				d := abs(int(s[i]) - int(s[j]))
				op := min(d, 26-d)
				if op <= k {
					res = max(res, f[k-op][i+1][j-1]+2)
				}
				f[k][i][j] = res
			}
		}
	}
	return f[K][0][n-1]
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2k)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2k)$。

**注**：$f$ 的第一维度可以用滚动数组优化，这样空间复杂度可以优化到 $\mathcal{O}(n^2|\Sigma|)$，其中 $|\Sigma|=26$ 是字符集合的大小。

更多相似题目，见下面动态规划题单中的「**§8.1 最长回文子序列**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
