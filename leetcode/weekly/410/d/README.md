看示例 1，$\textit{nums}=[2,3,2]$。

如果 $\textit{arr}_1[2]=2$，那么 $\textit{arr}_2[2]=\textit{nums}[2] - \textit{arr}_1[2]=2-2= 0$。考虑枚举 $\textit{arr}_1[1]$ 是多少：

- 如果 $\textit{arr}_1[1]=0$，那么问题变成计算下标 $0$ 到 $1$ 中的单调数组对的个数，且 $\textit{arr}_1[1]=0$。
- 如果 $\textit{arr}_1[1]=1$，那么问题变成计算下标 $0$ 到 $1$ 中的单调数组对的个数，且 $\textit{arr}_1[1]=1$。
- 如果 $\textit{arr}_1[1]=2$，那么问题变成计算下标 $0$ 到 $1$ 中的单调数组对的个数，且 $\textit{arr}_1[1]=2$。
- 注意 $\textit{arr}_1[1]$ 不能超过 $\textit{arr}_1[2]$，且 $\textit{arr}_2[1] = \textit{nums}[1] - \textit{arr}_1[1]$ 不能低于 $\textit{arr}_2[2]$。

累加这些方案数，我们就得到了下标 $0$ 到 $2$ 中的单调数组对的个数，且 $\textit{arr}_1[2]=2$。

上面的讨论说明本题的**子问题**是「下标 $0$ 到 $i$ 中的单调数组对的个数，且 $\textit{arr}_1[i]=j$」，将其记作 $f[i][j]$。

考虑枚举 $\textit{arr}_1[i-1] = k$，那么必须满足 $\textit{arr}_1[i-1]\le \textit{arr}_1[i]$ 且 $\textit{arr}_2[i-1]\ge \textit{arr}_2[i]$，即 $k\le j$ 且 $\textit{nums}[i-1]-k\ge \textit{nums}[i] - j$。

整理得

$$
k \le \min(j,\textit{nums}[i-1] - \textit{nums}[i] + j) = j + \min(\textit{nums}[i-1] - \textit{nums}[i], 0)
$$

⚠**注意**：无需判断 $k\le \textit{nums}[i-1]$，因为 $\textit{nums}[i-1] - \textit{nums}[i] + j = \textit{nums}[i-1] - (\textit{nums}[i] - j)\le \textit{nums}[i-1]$，由数学归纳法可以证明 $k\le \textit{nums}[i-1]$。

记 $\textit{maxK} = j + \min(\textit{nums}[i-1] - \textit{nums}[i], 0)$，那么有

$$
f[i][j] = 
\begin{cases} 
0, & \textit{maxK} < 0     \\
\sum\limits_{k=0}^{\textit{maxK}} f[i-1][k], &  \textit{maxK} \ge 0     \\
\end{cases}
$$

其中和式可以用 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 优化。

设 $f[i-1]$ 的前缀和 $s[j] = \sum\limits_{k=0}^{j} f[i-1][k]$，状态转移方程化简为

$$
f[i][j] =
\begin{cases}
0, & \textit{maxK} < 0     \\
s[\textit{maxK}], &  \textit{maxK} \ge 0     \\
\end{cases}
$$

初始值：$f[0][j] = 1$，其中 $j=0,1,2,\cdots,\textit{nums}[0]$。

答案：枚举 $\textit{arr}_1[n-1] = 0,1,2,\cdots,\textit{nums}[n-1]$，累加得 $\sum\limits_{j=0}^{\textit{nums}[n-1]} f[n-1][j]$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Cf421v7Ky/)，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def countOfPairs(self, nums: List[int]) -> int:
        MOD = 1_000_000_007
        n = len(nums)
        m = max(nums)
        f = [[0] * (m + 1) for _ in range(n)]
        for j in range(nums[0] + 1):
            f[0][j] = 1
        for i in range(1, n):
            s = list(accumulate(f[i - 1]))  # f[i-1] 的前缀和
            for j in range(nums[i] + 1):
                max_k = j + min(nums[i - 1] - nums[i], 0)
                f[i][j] = s[max_k] % MOD if max_k >= 0 else 0
        return sum(f[-1][:nums[-1] + 1]) % MOD
```

```java [sol-Java]
class Solution {
    public int countOfPairs(int[] nums) {
        final int MOD = 1_000_000_007;
        int n = nums.length;
        int m = Arrays.stream(nums).max().getAsInt();
        long[][] f = new long[n][m + 1];
        long[] s = new long[m + 1];

        Arrays.fill(f[0], 0, nums[0] + 1, 1);
        for (int i = 1; i < n; i++) {
            s[0] = f[i - 1][0];
            for (int k = 1; k <= m; k++) {
                s[k] = (s[k - 1] + f[i - 1][k]) % MOD; // f[i-1] 的前缀和
            }
            for (int j = 0; j <= nums[i]; j++) {
                int maxK = j + Math.min(nums[i - 1] - nums[i], 0);
                f[i][j] = maxK >= 0 ? s[maxK] % MOD : 0;
            }
        }

        return (int) (Arrays.stream(f[n - 1], 0, nums[n - 1] + 1).sum() % MOD);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countOfPairs(vector<int>& nums) {
        const int MOD = 1'000'000'007;
        int n = nums.size();
        int m = ranges::max(nums);
        vector<vector<long long>> f(n, vector<long long>(m + 1));
        vector<long long> s(m + 1);

        fill(f[0].begin(), f[0].begin() + nums[0] + 1, 1);
        for (int i = 1; i < n; i++) {
            partial_sum(f[i - 1].begin(), f[i - 1].end(), s.begin()); // f[i-1] 的前缀和
            for (int j = 0; j <= nums[i]; j++) {
                int max_k = j + min(nums[i - 1] - nums[i], 0);
                f[i][j] = max_k >= 0 ? s[max_k] % MOD : 0;
            }
        }

        return reduce(f[n - 1].begin(), f[n - 1].begin() + nums[n - 1] + 1, 0LL) % MOD;
    }
};
```

```go [sol-Go]
func countOfPairs(nums []int) (ans int) {
	const mod = 1_000_000_007
	n := len(nums)
	m := slices.Max(nums)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	s := make([]int, m+1)

	for j := 0; j <= nums[0]; j++ {
		f[0][j] = 1
	}
	for i := 1; i < n; i++ {
		s[0] = f[i-1][0]
		for k := 1; k <= m; k++ {
			s[k] = s[k-1] + f[i-1][k] // f[i-1] 的前缀和
		}
		for j := 0; j <= nums[i]; j++ {
			maxK := j + min(nums[i-1]-nums[i], 0)
			if maxK >= 0 {
				f[i][j] = s[maxK] % mod
			}
		}
	}

	for _, v := range f[n-1][:nums[n-1]+1] {
		ans += v
	}
	return ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 是 $\textit{nums}$ 的长度，$m=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(nm)$。

## 空间优化

进一步分析，$\textit{maxK} \ge 0$ 即

$$
j + \min(\textit{nums}[i-1] - \textit{nums}[i], 0) \ge 0
$$

变形得

$$
j\ge \max(\textit{nums}[i] - \textit{nums}[i-1], 0)
$$

记 $j_0 = \max(\textit{nums}[i] - \textit{nums}[i-1], 0)$，那么我们只需要计算在 $j$ 区间 $[j_0, \textit{nums}[i]]$ 中的 $f[i][j]$，其余 $f[i][j]$ 均为 $0$。

设 $s[j] = \sum\limits_{k=0}^{j} f[i-1][k]$，状态转移方程化简为

$$
f[i][j] =
\begin{cases}
0, & j < j_0     \\
s[j-j_0], &  j\ge j_0    \\
\end{cases}
$$

代码实现时，$f[i][j]$ 可以用一维数组滚动计算：先把前缀和直接保存在 $f$ 数组中，然后倒序更新 $f[j] = f[j-j_0]$（倒序更新的理由同 0-1 背包）。

此外，由于在计算答案时只考虑 $j\le \textit{nums}[n-1]$ 的状态，所以 $f$ 数组只需要开 $\textit{nums}[n-1]+1$ 的大小。

```py [sol-Python3]
class Solution:
    def countOfPairs(self, nums: List[int]) -> int:
        MOD = 1_000_000_007
        m = nums[-1]
        f = [0] * (m + 1)
        for j in range(min(nums[0], m) + 1):
            f[j] = 1
        for pre, cur in pairwise(nums):
            for j in range(1, m + 1):
                f[j] += f[j - 1]  # 计算前缀和
            j0 = max(cur - pre, 0)
            for j in range(min(cur, m), j0 - 1, -1):
                f[j] = f[j - j0] % MOD
            for j in range(min(j0, m + 1)):
                f[j] = 0
        return sum(f) % MOD
```

```java [sol-Java]
class Solution {
    public int countOfPairs(int[] nums) {
        final int MOD = 1_000_000_007;
        int n = nums.length;
        int m = nums[n - 1];
        int[] f = new int[m + 1];
        Arrays.fill(f, 0, Math.min(nums[0], m) + 1, 1);

        for (int i = 1; i < n; i++) {
            for (int j = 1; j <= m; j++) {
                f[j] = (f[j] + f[j - 1]) % MOD; // 计算前缀和
            }
            int j0 = Math.max(nums[i] - nums[i - 1], 0);
            for (int j = Math.min(nums[i], m); j >= j0; j--) {
                f[j] = f[j - j0];
            }
            Arrays.fill(f, 0, Math.min(j0, m + 1), 0);
        }

        long ans = 0;
        for (int v : f) {
            ans += v;
        }
        return (int) (ans % MOD);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countOfPairs(vector<int>& nums) {
        const int MOD = 1'000'000'007;
        int n = nums.size(), m = nums.back();
        vector<int> f(m + 1);
        fill(f.begin(), f.begin() + min(nums[0], m) + 1, 1);
        for (int i = 1; i < n; i++) {
            for (int j = 1; j <= m; j++) {
                f[j] = (f[j] + f[j - 1]) % MOD; // 计算前缀和
            }
            int j0 = max(nums[i] - nums[i - 1], 0);
            for (int j = min(nums[i], m); j >= j0; j--) {
                f[j] = f[j - j0];
            }
            fill(f.begin(), f.begin() + min(j0, m + 1), 0);
        }
        return reduce(f.begin(), f.end(), 0LL) % MOD;
    }
};
```

```go [sol-Go]
func countOfPairs(nums []int) (ans int) {
	const mod = 1_000_000_007
	n := len(nums)
	m := nums[n-1]
	f := make([]int, m+1)
	for j := range f[:min(nums[0], m)+1] {
		f[j] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= m; j++ {
			f[j] += f[j-1] // 计算前缀和
		}
		j0 := max(nums[i]-nums[i-1], 0)
		for j := min(nums[i], m); j >= j0; j-- {
			f[j] = f[j-j0] % mod
		}
		clear(f[:min(j0, m+1)])
	}
	for _, v := range f {
		ans += v
	}
	return ans % mod
}
```

## 进一步优化

如果 $j_0 > \min(\textit{nums}[i],m)$，那么后面计算出的 $f[j]$ 均为 $0$，我们可以直接返回 $0$。

此外，前缀和只需要计算到 $j=\min(\textit{nums}[i],m) - j_0$ 为止。

```py [sol-Python3]
class Solution:
    def countOfPairs(self, nums: List[int]) -> int:
        MOD = 1_000_000_007
        m = nums[-1]
        f = [0] * (m + 1)
        for j in range(min(nums[0], m) + 1):
            f[j] = 1
        for pre, cur in pairwise(nums):
            j0 = max(cur - pre, 0)
            m2 = min(cur, m)
            if j0 > m2:
                return 0
            for j in range(1, m2 - j0 + 1):
                f[j] = (f[j] + f[j - 1]) % MOD  # 计算前缀和
            f[j0: m2 + 1] = f[:m2 - j0 + 1]
            f[:j0] = [0] * j0
        return sum(f) % MOD
```

```java [sol-Java]
class Solution {
    public int countOfPairs(int[] nums) {
        final int MOD = 1_000_000_007;
        int n = nums.length;
        int m = nums[n - 1];
        int[] f = new int[m + 1];
        Arrays.fill(f, 0, Math.min(nums[0], m) + 1, 1);

        for (int i = 1; i < n; i++) {
            int j0 = Math.max(nums[i] - nums[i - 1], 0);
            int m2 = Math.min(nums[i], m);
            if (j0 > m2) {
                return 0;
            }
            for (int j = 1; j <= m2 - j0; j++) {
                f[j] = (f[j] + f[j - 1]) % MOD; // 计算前缀和
            }
            System.arraycopy(f, 0, f, j0, m2 - j0 + 1);
            Arrays.fill(f, 0, j0, 0);
        }

        long ans = 0;
        for (int v : f) {
            ans += v;
        }
        return (int) (ans % MOD);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countOfPairs(vector<int>& nums) {
        const int MOD = 1'000'000'007;
        int n = nums.size(), m = nums[n - 1];
        vector<int> f(m + 1);
        fill(f.begin(), f.begin() + min(nums[0], m) + 1, 1);
        for (int i = 1; i < n; i++) {
            int j0 = max(nums[i] - nums[i - 1], 0);
            int m2 = min(nums[i], m);
            if (j0 > m2) {
                return 0;
            }
            for (int j = 1; j <= m2 - j0; j++) {
                f[j] = (f[j] + f[j - 1]) % MOD; // 计算前缀和
            }
            copy(f.begin(), f.begin() + m2 - j0 + 1, f.begin() + j0);
            fill(f.begin(), f.begin() + j0, 0);
        }
        return reduce(f.begin(), f.end(), 0LL) % MOD;
    }
};
```

```go [sol-Go]
func countOfPairs(nums []int) (ans int) {
	const mod = 1_000_000_007
	n := len(nums)
	m := nums[n-1]
	f := make([]int, m+1)
	for j := range f[:min(nums[0], m)+1] {
		f[j] = 1
	}
	for i := 1; i < n; i++ {
		j0 := max(nums[i]-nums[i-1], 0)
		m2 := min(nums[i], m)
		if j0 > m2 {
			return 0
		}
		for j := 1; j <= m2-j0; j++ {
			f[j] = (f[j] + f[j-1]) % mod // 计算前缀和
		}
		copy(f[j0:m2+1], f)
		clear(f[:j0])
	}
	for _, v := range f {
		ans += v
	}
	return ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 是 $\textit{nums}$ 的长度，$m=\textit{nums}[n-1]$。
- 空间复杂度：$\mathcal{O}(m)$。

更多相似题目，见下面动态规划题单中的「**§11.1 前缀和优化 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
