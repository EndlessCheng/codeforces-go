## 转化

题意相当于把 $\textit{time}$ 数组切 $n-k-1$ 刀，划分成恰好 $n-k$ 个连续子数组。特别地，在 $0$ 和 $1$ 之间一定要切一刀，把 $\textit{time}[0]$ 单独作为一个子数组。（注意合并操作的 $i$ 必须大于 $0$）

示例 1 把 $\textit{time}=[5,8,3,6]$ 划分成 $[5],[8,3],[6]$ 三个子数组。

用区间 $[i,j]$ 表示子数组的左右端点（下标）。对于两个相邻的子数组 $A=[i,j]$ 和 $B=[j+1,k]$，我们需要从 $\textit{position}[j]$ 移动到 $\textit{position}[k]$，每公里所需时间为 $A$ 的元素和。

设 $\textit{time}$ 的 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solutions/2693498/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 数组为 $s$。这对子数组 $(A,B)$ 对总旅行时间的贡献为

$$
(\textit{position}[k] - \textit{position}[j])\cdot (s[j+1] - s[i])
$$

示例 1 划分成 $[5],[8,3],[6]$ 三个子数组，其中子数组 $[8,3]$ 和 $[6]$，对应着从 $\textit{position}[2]=8$ 移动到 $\textit{position}[3]=10$，移动 $10-8=2$ 公里，每公里所需时间为子数组 $[8,3]$ 的元素和 $8+3=11$，所以路段旅行时间为 $2\times 11 = 22$。

## DP 模型：相邻相关划分型 DP

回想一下 [300. 最长递增子序列](https://leetcode.cn/problems/longest-increasing-subsequence/) 的做法：设当前下标为 $i$，枚举下一个下标 $j$，判断 $\textit{nums}[i]< \textit{nums}[j]$ 是否成立。这是标准的相邻相关子序列 DP 问题。本题也类似，设当前子数组为 $[i,j]$，枚举下一个子数组为 $[j+1,k]$。

从左到右模拟旅行的过程。我们需要知道如下信息：

- 当前子数组为 $[i,j]$。
- 还需执行恰好 $\textit{leftK}$ 次合并操作。

定义 $\textit{dfs}(i,j,\textit{leftK})$ 表示在上述情况下，完成剩余旅程需要的最小旅行时间。

枚举下一个子数组的右端点 $k=j+1,j+2,\ldots, \min(n-1, j+1+\textit{leftK})$，这意味着我们执行了 $k-j-1$ 次合并操作，接下来要解决的问题变成：

- 当前子数组为 $[j+1,k]$。
- 还需执行恰好 $\textit{leftK}-(k-j-1)$ 次合并操作。

在上述情况下，完成剩余旅程需要的最小旅行时间，即 $\textit{dfs}(j+1,k,\textit{leftK} - (k-j-1))$。这两个子数组 $[i,j]$ 和 $[j+1,k]$ 对总旅行时间的贡献为 $(\textit{position}[k] - \textit{position}[j])\cdot (s[j+1] - s[i])$。

取最小值，得

$$
\textit{dfs}(i,j,\textit{leftK}) = \min_{k=j+1}^{\min(n-1, j+1+\textit{leftK})} \textit{dfs}(j+1,k,\textit{leftK} - (k-j-1)) + (\textit{position}[k] - \textit{position}[j])\cdot (s[j+1] - s[i])
$$

递归边界：当 $j=n-1$ 时，必须有 $\textit{leftK}=0$，所以 $\textit{dfs}(i,n-1,0)=0$，其余 $\textit{dfs}(i,n-1,\textit{leftK})=\infty$。返回 $\infty$，转移方程中的 $\min$ 不会取到不合法的状态。

递归入口：$\textit{dfs}(0,0,k)$，即答案。注意第一个子数组一定是 $[0,0]$。

## 写法一：记忆化搜索

原理见 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

[本题视频讲解](https://www.bilibili.com/video/BV1avVwz5EbY/?t=12m34s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minTravelTime(self, _, n: int, k: int, position: List[int], time: List[int]) -> int:
        s = list(accumulate(time, initial=0))  # 计算 time 的前缀和
        @cache
        def dfs(i: int, j: int, left_k: int) -> int:
            if j == n - 1:  # 到达终点
                return inf if left_k else 0
            t = s[j + 1] - s[i]  # 合并到 time[j] 的时间
            # 枚举下一个子数组 [j+1, k]
            return min(dfs(j + 1, k, left_k - (k - j - 1)) + (position[k] - position[j]) * t
                       for k in range(j + 1, min(n, j + 2 + left_k)))
        return dfs(0, 0, k)  # 第一个子数组是 [0, 0]
```

```java [sol-Java]
class Solution {
    public int minTravelTime(int l, int n, int k, int[] position, int[] time) {
        int[] s = new int[n];
        for (int i = 0; i < n - 1; i++) { // time[n-1] 用不到
            s[i + 1] = s[i] + time[i]; // 计算 time 的前缀和
        }

        int[][][] memo = new int[n - 1][n - 1][k + 1];
        return dfs(0, 0, k, position, s, memo); // 第一个子数组是 [0, 0]
    }

    private int dfs(int i, int j, int leftK, int[] position, int[] s, int[][][] memo) {
        int n = position.length;
        if (j == n - 1) { // 到达终点
            return leftK > 0 ? Integer.MAX_VALUE / 2 : 0; // 除以 2，避免下面计算 r 的地方加法溢出
        }
        if (memo[i][j][leftK] > 0) { // 之前计算过
            return memo[i][j][leftK];
        }
        int res = Integer.MAX_VALUE;
        int t = s[j + 1] - s[i]; // 合并到 time[j] 的时间
        // 枚举下一个子数组 [j+1, k]
        for (int k = j + 1; k < Math.min(n, j + 2 + leftK); k++) {
            int r = dfs(j + 1, k, leftK - (k - j - 1), position, s, memo) + (position[k] - position[j]) * t;
            res = Math.min(res, r);
        }
        return memo[i][j][leftK] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minTravelTime(int, int n, int K, vector<int>& position, vector<int>& time) {
        vector<int> s(n); // time 的前缀和
        partial_sum(time.begin(), time.end() - 1, s.begin() + 1); // time[n-1] 用不到

        vector memo(n - 1, vector(n - 1, vector<int>(K + 1)));
        auto dfs = [&](this auto&& dfs, int i, int j, int left_k) -> int {
            if (j == n - 1) { // 到达终点
                return left_k ? INT_MAX / 2 : 0; // 除以 2，避免下面计算 r 的地方加法溢出
            }
            int& res = memo[i][j][left_k]; // 注意这里是引用
            if (res) { // 之前计算过
                return res;
            }
            res = INT_MAX;
            int t = s[j + 1] - s[i]; // 合并到 time[j] 的时间
            // 枚举下一个子数组 [j+1, k]
            for (int k = j + 1; k < min(n, j + 2 + left_k); k++) {
                int r = dfs(j + 1, k, left_k - (k - j - 1)) + (position[k] - position[j]) * t;
                res = min(res, r);
            }
            return res;
        };
        return dfs(0, 0, K); // 第一个子数组是 [0, 0]
    }
};
```

```go [sol-Go]
func minTravelTime(_, n, K int, position, time []int) int {
	s := make([]int, n)
	for i, t := range time[:n-1] { // time[n-1] 用不到
		s[i+1] = s[i] + t // 计算 time 的前缀和
	}

	memo := make([][][]int, n-1)
	for i := range memo {
		memo[i] = make([][]int, n-1)
		for j := range memo[i] {
			memo[i][j] = make([]int, K+1)
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, leftK int) int {
		if j == n-1 { // 到达终点
			if leftK > 0 { // 不合法
				return math.MaxInt / 2 // 避免下面计算 r 的地方加法溢出
			}
			return 0
		}
		p := &memo[i][j][leftK]
		if *p > 0 { // 之前计算过
			return *p
		}
		res := math.MaxInt
		t := s[j+1] - s[i] // 合并到 time[j] 的时间
		// 枚举下一个子数组 [j+1, k]
		for k := j + 1; k < min(n, j+2+leftK); k++ {
			r := dfs(j+1, k, leftK-(k-j-1)) + (position[k]-position[j])*t
			res = min(res, r)
		}
		*p = res // 记忆化
		return res
	}
	return dfs(0, 0, K) // 第一个子数组是 [0, 0]
}
```

注意到子数组的长度不会超过 $k+1$，我们可以把子数组 $[i,j]$ 用 $[j-\textit{sz},j]$ 表示，从而减少 $\textit{memo}$ 数组的大小。注意这里的 $\textit{sz}$ 是子数组的长度减一。

```py [sol-Python3]
class Solution:
    def minTravelTime(self, _, n: int, k: int, position: List[int], time: List[int]) -> int:
        s = list(accumulate(time, initial=0))  # 计算 time 的前缀和
        @cache
        def dfs(j: int, sz: int, left_k: int) -> int:
            if j == n - 1:  # 到达终点
                return inf if left_k else 0
            t = s[j + 1] - s[j - sz]  # 合并到 time[j] 的时间
            # 枚举下一个子数组 [j+1, k]
            return min(dfs(k, k - j - 1, left_k - (k - j - 1)) + (position[k] - position[j]) * t
                       for k in range(j + 1, min(n, j + 2 + left_k)))
        return dfs(0, 0, k)  # 第一个子数组是 [0, 0]
```

```java [sol-Java]
class Solution {
    public int minTravelTime(int l, int n, int k, int[] position, int[] time) {
        int[] s = new int[n];
        for (int i = 0; i < n - 1; i++) { // time[n-1] 用不到
            s[i + 1] = s[i] + time[i]; // 计算 time 的前缀和
        }

        int[][][] memo = new int[n - 1][k + 1][k + 1];
        return dfs(0, 0, k, position, s, memo); // 第一个子数组是 [0, 0]
    }

    private int dfs(int j, int sz, int leftK, int[] position, int[] s, int[][][] memo) {
        int n = position.length;
        if (j == n - 1) { // 到达终点
            return leftK > 0 ? Integer.MAX_VALUE / 2 : 0; // 除以 2，避免下面计算 r 的地方加法溢出
        }
        if (memo[j][sz][leftK] > 0) { // 之前计算过
            return memo[j][sz][leftK];
        }
        int res = Integer.MAX_VALUE;
        int t = s[j + 1] - s[j - sz]; // 合并到 time[j] 的时间
        // 枚举下一个子数组 [j+1, k]
        for (int k = j + 1; k < Math.min(n, j + 2 + leftK); k++) {
            int r = dfs(k, k - j - 1, leftK - (k - j - 1), position, s, memo) + (position[k] - position[j]) * t;
            res = Math.min(res, r);
        }
        return memo[j][sz][leftK] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minTravelTime(int, int n, int K, vector<int>& position, vector<int>& time) {
        vector<int> s(n); // time 的前缀和
        partial_sum(time.begin(), time.end() - 1, s.begin() + 1); // time[n-1] 用不到

        vector memo(n - 1, vector(K + 1, vector<int>(K + 1)));
        auto dfs = [&](this auto&& dfs, int j, int sz, int left_k) -> int {
            if (j == n - 1) { // 到达终点
                return left_k ? INT_MAX / 2 : 0; // 除以 2，避免下面计算 r 的地方加法溢出
            }
            int& res = memo[j][sz][left_k]; // 注意这里是引用
            if (res) { // 之前计算过
                return res;
            }
            res = INT_MAX;
            int t = s[j + 1] - s[j - sz]; // 合并到 time[j] 的时间
            // 枚举下一个子数组 [j+1, k]
            for (int k = j + 1; k < min(n, j + 2 + left_k); k++) {
                int r = dfs(k, k - j - 1, left_k - (k - j - 1)) + (position[k] - position[j]) * t;
                res = min(res, r);
            }
            return res;
        };
        return dfs(0, 0, K); // 第一个子数组是 [0, 0]
    }
};
```

```go [sol-Go]
func minTravelTime(_, n, K int, position, time []int) int {
	s := make([]int, n)
	for i, t := range time[:n-1] { // time[n-1] 用不到
		s[i+1] = s[i] + t // 计算 time 的前缀和
	}

	memo := make([][][]int, n-1)
	for i := range memo {
		memo[i] = make([][]int, K+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, K+1)
		}
	}
	var dfs func(int, int, int) int
	dfs = func(j, sz, leftK int) int {
		if j == n-1 { // 到达终点
			if leftK > 0 { // 不合法
				return math.MaxInt / 2 // 避免下面计算 r 的地方加法溢出
			}
			return 0
		}
		p := &memo[j][sz][leftK]
		if *p > 0 {
			return *p
		}
		res := math.MaxInt
		t := s[j+1] - s[j-sz] // 合并到 time[j] 的时间
		// 枚举下一个子数组 [j+1, k]
		for k := j + 1; k < min(n, j+2+leftK); k++ {
			r := dfs(k, k-j-1, leftK-(k-j-1)) + (position[k]-position[j])*t
			res = min(res, r)
		}
		*p = res
		return res
	}
	return dfs(0, 0, K) // 第一个子数组是 [0, 0]
}
```

## 写法二：递推

把记忆化搜索 1:1 翻译成递推。

DP 数组的初始值怎么写？就是记忆化搜索的递归边界。

```py [sol-Python3]
class Solution:
    def minTravelTime(self, _, n: int, K: int, position: List[int], time: List[int]) -> int:
        s = list(accumulate(time, initial=0))  # 计算 time 的前缀和
        f = [[[inf] * (K + 1) for _ in range(K + 1)] for _ in range(n)]
        for sz in range(K + 1):
            f[-1][sz][0] = 0
        for j in range(n - 2, -1, -1):  # 转移来源 k 比 j 大，所以要倒序
            for sz in range(min(K, j) + 1):
                t = s[j + 1] - s[j - sz]  # 合并到 time[j] 的时间
                for left_k in range(min(K, n - 2 - j) + 1):
                    # 枚举下一个子数组 [j+1, k]
                    f[j][sz][left_k] = min(f[k][k - j - 1][left_k - (k - j - 1)] + (position[k] - position[j]) * t
                                          for k in range(j + 1, j + 2 + left_k))
        return f[0][0][K]  # 第一个子数组是 [0, 0]
```

```java [sol-Java]
class Solution {
    public int minTravelTime(int l, int n, int K, int[] position, int[] time) {
        int[] s = new int[n];
        for (int i = 0; i < n - 1; i++) { // time[n-1] 用不到
            s[i + 1] = s[i] + time[i]; // 计算 time 的前缀和
        }

        int[][][] f = new int[n][K + 1][K + 1];
        for (int[][] mat : f) {
            for (int[] row : mat) {
                Arrays.fill(row, Integer.MAX_VALUE / 2);
            }
        }
        for (int sz = 0; sz <= K; sz++) {
            f[n - 1][sz][0] = 0;
        }

        for (int j = n - 2; j >= 0; j--) { // 转移来源 k 比 j 大，所以要倒序
            for (int sz = 0; sz <= Math.min(K, j); sz++) {
                int t = s[j + 1] - s[j - sz]; // 合并到 time[j] 的时间
                for (int leftK = 0; leftK <= Math.min(K, n - 2 - j); leftK++) {
                    int res = Integer.MAX_VALUE;
                    // 枚举下一个子数组 [j+1, k]
                    for (int k = j + 1; k <= j + 1 + leftK; k++) {
                        int r = f[k][k - j - 1][leftK - (k - j - 1)] + (position[k] - position[j]) * t;
                        res = Math.min(res, r);
                    }
                    f[j][sz][leftK] = res;
                }
            }
        }
        return f[0][0][K]; // 第一个子数组是 [0, 0]
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minTravelTime(int, int n, int K, vector<int>& position, vector<int>& time) {
        vector<int> s(n); // time 的前缀和
        partial_sum(time.begin(), time.end() - 1, s.begin() + 1); // time[n-1] 用不到

        vector f(n, vector(K + 1, vector<int>(K + 1, INT_MAX / 2)));
        for (int sz = 0; sz <= K; sz++) {
            f[n - 1][sz][0] = 0;
        }
        for (int j = n - 2; j >= 0; j--) { // 转移来源 k 比 j 大，所以要倒序
            for (int sz = 0; sz <= min(K, j); sz++) {
                int t = s[j + 1] - s[j - sz]; // 合并到 time[j] 的时间
                for (int left_k = 0; left_k <= min(K, n - 2 - j); left_k++) {
                    int res = INT_MAX;
                    // 枚举下一个子数组 [j+1, k]
                    for (int k = j + 1; k <= j + 1 + left_k; k++) {
                        int r = f[k][k - j - 1][left_k - (k - j - 1)] + (position[k] - position[j]) * t;
                        res = min(res, r);
                    }
                    f[j][sz][left_k] = res;
                }
            }
        }
        return f[0][0][K]; // 第一个子数组是 [0, 0]
    }
};
```

```go [sol-Go]
func minTravelTime(_, n, K int, position, time []int) int {
	s := make([]int, n)
	for i, t := range time[:n-1] { // time[n-1] 用不到
		s[i+1] = s[i] + t // 计算 time 的前缀和
	}

	f := make([][][]int, n)
	for j := range f {
		f[j] = make([][]int, K+1)
		for sz := range f[j] {
			f[j][sz] = make([]int, K+1)
			for leftK := range f[j][sz] {
				f[j][sz][leftK] = math.MaxInt / 2
			}
		}
	}
	for sz := range K + 1 {
		f[n-1][sz][0] = 0
	}

	for j := n - 2; j >= 0; j-- { // 转移来源 k 比 j 大，所以要倒序
		for sz := range min(K, j) + 1 {
			t := s[j+1] - s[j-sz] // 合并到 time[j] 的时间
			for leftK := range min(K, n-2-j) + 1 {
				res := math.MaxInt
				// 枚举下一个子数组 [j+1, k]
				for k := j + 1; k <= j+1+leftK; k++ {
					r := f[k][k-j-1][leftK-(k-j-1)] + (position[k]-position[j])*t
					res = min(res, r)
				}
				f[j][sz][leftK] = res
			}
		}
	}
	return f[0][0][K] // 第一个子数组是 [0, 0]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk^3)$。注意最内层的循环是 $\mathcal{O}(k)$。
- 空间复杂度：$\mathcal{O}(nk^2)$。

更多相似题目，见下面动态规划题单的「**五、划分型 DP**」。

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
