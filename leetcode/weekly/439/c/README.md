按照**划分型 DP** 的套路（见 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/) §5.3 节），定义 $f[i][j]$ 表示从长为 $j$ 的前缀 $\textit{nums}[0]$ 到 $\textit{nums}[j-1]$ 取出 $i$ 个连续子数组所得到的最大和。

> 注：这里用左闭右开区间 $[0,j)$ 表示前缀，与前缀和定义相匹配，方便后面做公式变形。

分类讨论：

- 不选 $\textit{nums}[j-1]$，问题变成从长为 $j-1$ 的前缀 $\textit{nums}[0]$ 到 $\textit{nums}[j-2]$ 取出 $i$ 个连续子数组所得到的最大和，即 $f[i][j-1]$。
- 选 $\textit{nums}[j-1]$，也就是子数组 $\textit{nums}[L]$ 到 $\textit{nums}[j-1]$（$L$ 是我们枚举的值），问题变成从长为 $L$ 的前缀 $\textit{nums}[0]$ 到 $\textit{nums}[L-1]$ 取出 $i-1$ 个连续子数组所得到的最大和，即 $f[i-1][L]$。

二者取最大值，得

$$
f[i][j] = \max\left(f[i][j-1], \max_{L=(i-1)\cdot m}^{j-m} f[i-1][L] + s[j] - s[L]\right)
$$

其中 $s$ 是 $\textit{nums}$ 的前缀和数组（长为 $n+1$），原理和定义请看 [前缀和讲解](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

其中 $L$ 最小是 $(i-1)\cdot m$，因为左边有 $i-1$ 个长度至少为 $m$ 的子数组。

**初始值**：

- $f[0][j] = 0$，不选任何子数组，元素和为 $0$。
- 如果 $j < i\cdot m$，那么 $f[i][j]=-\infty$，因为没有足够的数选取 $i$ 个长度至少为 $m$ 的子数组。这里用 $-\infty$ 表示不合法的状态，这样计算 $\max$ 不会取到不合法的状态。

**答案**：$f[k][n]$。

这样做的时间复杂度为 $\mathcal{O}(n^2k)$，会超时。

由于

$$
\max_{L=(i-1)\cdot m}^{j-m} f[i-1][L] + s[j] - s[L] = s[j] + \max_{L=(i-1)\cdot m}^{j-m} f[i-1][L] - s[L]
$$

故定义 $d[L] = f[i-1][L] - s[L]$，并用一个变量 $\textit{mx}$ 在遍历 $j$ 的同时维护从 $L=(i-1)\cdot m$ 到 $j-m$ 的最大 $d[L]$，那么转移方程优化成

$$
f[i][j] = \max(f[i][j-1], \textit{mx} + s[j])
$$

这样就可以做到 $\mathcal{O}(nk)$ 时间了。

代码实现时，$f$ 的第一个维度可以优化掉。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1QP9bY3EL6/?t=16m57s)，欢迎点赞关注~

## 写法一

```py [sol-Python3]
class Solution:
    def maxSum(self, nums: List[int], k: int, m: int) -> int:
        n = len(nums)
        s = list(accumulate(nums, initial=0))  # 前缀和
        f = [0] * (n + 1)
        for i in range(1, k + 1):
            nf = [-inf] * (n + 1)
            mx = -inf
            # 左右两边留出足够空间给其他子数组
            for j in range(i * m, n - (k - i) * m + 1):
                # mx 表示最大的 f[L]-s[L]，其中 L 在区间 [(i-1)*m, j-m] 中
                mx = max(mx, f[j - m] - s[j - m])
                nf[j] = max(nf[j - 1], mx + s[j])  # 不选 vs 选
            f = nf
        return f[n]
```

```java [sol-Java]
class Solution {
    public int maxSum(int[] nums, int k, int m) {
        int n = nums.length;
        int[] s = new int[n + 1];
        for (int i = 0; i < n; i++) {
            s[i + 1] = s[i] + nums[i]; // 前缀和
        }

        int[] f = new int[n + 1];
        for (int i = 1; i <= k; i++) {
            int[] nf = new int[n + 1];
            Arrays.fill(nf, Integer.MIN_VALUE / 2);
            int mx = Integer.MIN_VALUE;
            // 左右两边留出足够空间给其他子数组
            for (int j = i * m; j <= n - (k - i) * m; j++) {
                // mx 表示最大的 f[L]-s[L]，其中 L 在区间 [(i-1)*m, j-m] 中
                mx = Math.max(mx, f[j - m] - s[j - m]);
                nf[j] = Math.max(nf[j - 1], mx + s[j]); // 不选 vs 选
            }
            f = nf;
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxSum(vector<int>& nums, int k, int m) {
        int n = nums.size();
        vector<int> s(n + 1), f(n + 1);
        partial_sum(nums.begin(), nums.end(), s.begin() + 1); // 前缀和
        
        for (int i = 1; i <= k; i++) {
            vector<int> nf(n + 1, INT_MIN / 2);
            int mx = INT_MIN;
            // 左右两边留出足够空间给其他子数组
            for (int j = i * m; j <= n - (k - i) * m; j++) {
                // mx 表示最大的 f[L]-s[L]，其中 L 在区间 [(i-1)*m, j-m] 中
                mx = max(mx, f[j - m] - s[j - m]);
                nf[j] = max(nf[j - 1], mx + s[j]); // 不选 vs 选
            }
            f = move(nf);
        }
        return f[n];
    }
};
```

```go [sol-Go]
func maxSum(nums []int, k, m int) int {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}

	f := make([]int, n+1)
	for i := 1; i <= k; i++ {
		nf := make([]int, n+1)
		for j := range nf {
			nf[j] = math.MinInt / 2
		}
		mx := math.MinInt
		// 左右两边留出足够空间给其他子数组
		for j := i * m; j <= n-(k-i)*m; j++ {
			// mx 表示最大的 f[L]-s[L]，其中 L 在区间 [(i-1)*m, j-m] 中
			mx = max(mx, f[j-m]-s[j-m])
			nf[j] = max(nf[j-1], mx+s[j]) // 不选 vs 选
		}
		f = nf
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 写法二

```py [sol-Python3]
class Solution:
    def maxSum(self, nums: List[int], k: int, m: int) -> int:
        n = len(nums)
        s = list(accumulate(nums, initial=0))  # 前缀和
        f = [0] * (n + 1)
        d = [0] * (n + 1)
        for i in range(1, k + 1):
            for j in range(i * m - m, i * m):
                d[j] = f[j] - s[j]
                f[j] = -inf  # 即使 [0,j) 全选，也没有 i 个长为 m 的子数组
            mx = -inf
            # 左右两边留出足够空间给其他子数组
            for j in range(i * m, n - (k - i) * m + 1):
                # mx 表示最大的 f[L]-s[L]，其中 L 在区间 [(i-1)*m, j-m] 中
                mx = max(mx, d[j - m])
                d[j] = f[j] - s[j]
                f[j] = max(f[j - 1], mx + s[j])  # 不选 vs 选
        return f[n]
```

```java [sol-Java]
class Solution {
    public int maxSum(int[] nums, int k, int m) {
        int n = nums.length;
        int[] s = new int[n + 1];
        for (int i = 0; i < n; i++) {
            s[i + 1] = s[i] + nums[i]; // 前缀和
        }

        int[] f = new int[n + 1];
        int[] d = new int[n + 1];
        for (int i = 1; i <= k; i++) {
            for (int j = i * m - m; j < i * m; j++) {
                d[j] = f[j] - s[j];
                f[j] = Integer.MIN_VALUE / 2; // 即使 [0,j) 全选，也没有 i 个长为 m 的子数组
            }
            int mx = Integer.MIN_VALUE;
            // 左右两边留出足够空间给其他子数组
            for (int j = i * m; j <= n - (k - i) * m; j++) {
                // mx 表示最大的 f[L]-s[L]，其中 L 在区间 [(i-1)*m, j-m] 中
                mx = Math.max(mx, d[j - m]);
                d[j] = f[j] - s[j];
                f[j] = Math.max(f[j - 1], mx + s[j]); // 不选 vs 选
            }
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxSum(vector<int>& nums, int k, int m) {
        int n = nums.size();
        vector<int> s(n + 1), f(n + 1), d(n + 1);
        partial_sum(nums.begin(), nums.end(), s.begin() + 1); // 前缀和

        for (int i = 1; i <= k; i++) {
            for (int j = i * m - m; j < i * m; j++) {
                d[j] = f[j] - s[j];
                f[j] = INT_MIN / 2; // 即使 [0,j) 全选，也没有 i 个长为 m 的子数组
            }
            int mx = INT_MIN;
            // 左右两边留出足够空间给其他子数组
            for (int j = i * m; j <= n - (k - i) * m; j++) {
                // mx 表示最大的 f[L]-s[L]，其中 L 在区间 [(i-1)*m, j-m] 中
                mx = max(mx, d[j - m]);
                d[j] = f[j] - s[j];
                f[j] = max(f[j - 1], mx + s[j]); // 不选 vs 选
            }
        }
        return f[n];
    }
};
```

```go [sol-Go]
func maxSum(nums []int, k, m int) int {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}

	f := make([]int, n+1)
	d := make([]int, n+1)
	for i := 1; i <= k; i++ {
		for j := i*m - m; j < i*m; j++ {
			d[j] = f[j] - s[j]
			f[j] = math.MinInt / 2 // 即使 [0,j) 全选，也没有 i 个长为 m 的子数组
		}
		mx := math.MinInt
		// 左右两边留出足够空间给其他子数组
		for j := i * m; j <= n-(k-i)*m; j++ {
			// mx 表示最大的 f[L]-s[L]，其中 L 在区间 [(i-1)*m, j-m] 中
			mx = max(mx, d[j-m])
			d[j] = f[j] - s[j]
			f[j] = max(f[j-1], mx+s[j]) // 不选 vs 选
		}
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n-km)\cdot k)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面动态规划题单中的「**§11.1 前缀和优化 DP**」。

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
