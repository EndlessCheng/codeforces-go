移除次数尽量小，等价于保留的元素个数尽量多。

问题相当于：

- 给你 $n$ 个数，每个数要么选，要么不选。选的数的异或和恰好等于 $\textit{target}$，最多能选多少个数？

这是恰好装满型 0-1 背包，原理见视频讲解：[0-1 背包 完全背包【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

类似题目是 [2915. 和为目标值的最长子序列的长度](https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target/)，[我的题解](https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target/solutions/2502839/mo-ban-qia-hao-zhuang-man-xing-0-1-bei-b-0nca/)。由于 2915 题用的是大家熟悉的加法，推荐先完成 2915 题，再做本题。

和 2915 题一样，定义 $f[i+1][j]$ 表示在 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 中的，异或和为 $j$ 的子序列的最长长度。

由于异或运算本质是模 $2$ 加法，把 $v \oplus \textit{nums}[i] = j$ 移项得 $v = j\oplus \textit{nums}[i]$。我们只需把 2915 转移方程中的 $j - \textit{nums}[i]$ 改成 $j\oplus  \textit{nums}[i]$ 即可，其中 $\oplus$ 是异或运算。

$$
f[i+1][j] = \max(f[i][j],f[i][j\oplus\textit{nums}[i]] + 1)
$$

初始值 $f[0][0]=0$，其余为 $f[0][j] = -\infty$。

答案为 $f[n][\textit{target}]$。

**特殊情况**：设 $m$ 为 $\max(\textit{nums})$ 的二进制长度。如果 $m$ 小于 $\textit{target}$ 的二进制长度，那么 XOR 的二进制长度也小于 $\textit{target}$ 的二进制长度，必然无解，返回 $-1$。否则可以计算 DP，看看是否有解。由于 XOR 最大是 $2^m-1$，所以数组第二维的大小为 $2^m$。

## 答疑

**问**：为什么数组第二维的大小不能是 $\textit{target}+1$？

**答**：比如 $\textit{target}$ 二进制是 $100$，我们可能先异或得到 $110$，再异或一个等于 $10$ 的数，得到 $100$。一般地，在计算过程中，可能先算出比 $\textit{target}$ 大的数，再减小到 $\textit{target}$。所以要用 XOR 的最大值加一作为数组大小。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

## 优化前

```py [sol-Python3]
class Solution:
    def minRemovals(self, nums: List[int], target: int) -> int:
        m = max(nums).bit_length()
        if m < target.bit_length():
            return -1

        n = len(nums)
        f = [[-inf] * (1 << m) for _ in range(n + 1)]
        f[0][0] = 0

        for i, x in enumerate(nums):
            for j in range(1 << m):
                f[i + 1][j] = max(f[i][j], f[i][j ^ x] + 1)  # x 不选 or 选

        if f[n][target] < 0:
            return -1
        return len(nums) - f[n][target]
```

```java [sol-Java]
class Solution {
    public int minRemovals(int[] nums, int target) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        int m = 32 - Integer.numberOfLeadingZeros(mx);
        if (m < 32 - Integer.numberOfLeadingZeros(target)) {
            return -1;
        }

        int n = nums.length;
        int[][] f = new int[n + 1][1 << m];
        for (int[] row : f) {
            Arrays.fill(row, Integer.MIN_VALUE);
        }
        f[0][0] = 0;

        for (int i = 0; i < n; i++) {
            int x = nums[i];
            for (int j = 0; j < (1 << m); j++) {
                f[i + 1][j] = Math.max(f[i][j], f[i][j ^ x] + 1); // x 不选 or 选
            }
        }

        if (f[n][target] < 0) {
            return -1;
        }
        return nums.length - f[n][target];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minRemovals(vector<int>& nums, int target) {
        int m = bit_width((uint32_t) ranges::max(nums));
        if (m < bit_width((uint32_t) target)) {
            return -1;
        }

        int n = nums.size();
        vector f(n + 1, vector<int>(1 << m, INT_MIN));
        f[0][0] = 0;

        for (int i = 0; i < n; i++) {
            int x = nums[i];
            for (int j = 0; j < (1 << m); j++) {
                f[i + 1][j] = max(f[i][j], f[i][j ^ x] + 1); // x 不选 or 选
            }
        }

        if (f[n][target] < 0) {
            return -1;
        }
        return nums.size() - f[n][target];
    }
};
```

```go [sol-Go]
func minRemovals(nums []int, target int) int {
	m := bits.Len(uint(slices.Max(nums)))
	if m < bits.Len(uint(target)) {
		return -1
	}

	n := len(nums)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, 1<<m)
		for j := range f[i] {
			f[i][j] = math.MinInt
		}
	}
	f[0][0] = 0

	for i, x := range nums {
		for j := range 1 << m {
			f[i+1][j] = max(f[i][j], f[i][j^x]+1) // x 不选 or 选
		}
	}

	if f[n][target] < 0 {
		return -1
	}
	return len(nums) - f[n][target]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nU)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(nU)$。

## 空间优化（查表法）

```py [sol-Python3]
# 手写 max 更快
fmax = lambda a, b: b if b > a else a

class Solution:
    def minRemovals(self, nums: List[int], target: int) -> int:
        m = max(nums).bit_length()
        if m < target.bit_length():
            return -1

        f = [-inf] * (1 << m)
        f[0] = 0

        nf = [0] * (1 << m)
        for x in nums:
            for j in range(1 << m):
                nf[j] = max(f[j], f[j ^ x] + 1)  # x 不选 or 选
            f, nf = nf, f

        if f[target] < 0:
            return -1
        return len(nums) - f[target]
```

```java [sol-Java]
class Solution {
    public int minRemovals(int[] nums, int target) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        int m = 32 - Integer.numberOfLeadingZeros(mx);
        if (m < 32 - Integer.numberOfLeadingZeros(target)) {
            return -1;
        }

        int[] f = new int[1 << m];
        Arrays.fill(f, Integer.MIN_VALUE);
        f[0] = 0;

        int[] nf = new int[1 << m];
        for (int x : nums) {
            for (int j = 0; j < (1 << m); j++) {
                nf[j] = Math.max(f[j], f[j ^ x] + 1); // x 不选 or 选
            }
            int[] tmp = f;
            f = nf;
            nf = tmp;
        }

        if (f[target] < 0) {
            return -1;
        }
        return nums.length - f[target];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minRemovals(vector<int>& nums, int target) {
        int m = bit_width((uint32_t) ranges::max(nums));
        if (m < bit_width((uint32_t) target)) {
            return -1;
        }

        vector<int> f(1 << m, INT_MIN);
        f[0] = 0;

        vector<int> nf(1 << m);
        for (int x : nums) {
            for (int j = 0; j < (1 << m); j++) {
                nf[j] = max(f[j], f[j ^ x] + 1); // x 不选 or 选
            }
            swap(f, nf);
        }

        if (f[target] < 0) {
            return -1;
        }
        return nums.size() - f[target];
    }
};
```

```go [sol-Go]
func minRemovals(nums []int, target int) int {
	m := bits.Len(uint(slices.Max(nums)))
	if m < bits.Len(uint(target)) {
		return -1
	}

	f := make([]int, 1<<m)
	for i := range f {
		f[i] = math.MinInt
	}
	f[0] = 0

	nf := make([]int, 1<<m)
	for _, x := range nums {
		for j := range 1 << m {
			nf[j] = max(f[j], f[j^x]+1) // x 不选 or 选
		}
		f, nf = nf, f
	}

	if f[target] < 0 {
		return -1
	}
	return len(nums) - f[target]
}
```

## 空间优化（刷表法）

如果修改问题，把 XOR 改成没有逆运算的 AND 或者 OR，用**刷表法**更合适。也就是用当前状态更新其他状态。

```py [sol-Python3]
# 手写 max 更快
fmax = lambda a, b: b if b > a else a

class Solution:
    def minRemovals(self, nums: List[int], target: int) -> int:
        m = max(nums).bit_length()
        if m < target.bit_length():
            return -1

        f = [-inf] * (1 << m)
        f[0] = 0

        for x in nums:
            nf = f[:]
            for j, fj in enumerate(f):
                nf[j ^ x] = fmax(nf[j ^ x], fj + 1)  # x 不选 or 选
            f = nf

        if f[target] < 0:
            return -1
        return len(nums) - f[target]
```

```java [sol-Java]
class Solution {
    public int minRemovals(int[] nums, int target) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        int m = 32 - Integer.numberOfLeadingZeros(mx);
        if (m < 32 - Integer.numberOfLeadingZeros(target)) {
            return -1;
        }

        int[] f = new int[1 << m];
        Arrays.fill(f, Integer.MIN_VALUE);
        f[0] = 0;

        for (int x : nums) {
            int[] nf = f.clone();
            for (int j = 0; j < (1 << m); j++) {
                nf[j ^ x] = Math.max(nf[j ^ x], f[j] + 1); // x 不选 or 选
            }
            f = nf;
        }

        if (f[target] < 0) {
            return -1;
        }
        return nums.length - f[target];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minRemovals(vector<int>& nums, int target) {
        int m = bit_width((uint32_t) ranges::max(nums));
        if (m < bit_width((uint32_t) target)) {
            return -1;
        }

        vector<int> f(1 << m, INT_MIN);
        f[0] = 0;

        for (int x : nums) {
            auto nf = f;
            for (int j = 0; j < (1 << m); j++) {
                nf[j ^ x] = max(nf[j ^ x], f[j] + 1); // x 不选 or 选
            }
            f = nf;
        }

        if (f[target] < 0) {
            return -1;
        }
        return nums.size() - f[target];
    }
};
```

```go [sol-Go]
func minRemovals(nums []int, target int) int {
	m := bits.Len(uint(slices.Max(nums)))
	if m < bits.Len(uint(target)) {
		return -1
	}

	f := make([]int, 1<<m)
	for i := range f {
		f[i] = math.MinInt
	}
	f[0] = 0

	for _, x := range nums {
		nf := slices.Clone(f)
		for j, fj := range f {
			nf[j^x] = max(nf[j^x], fj+1) // x 不选 or 选
		}
		f = nf
	}

	if f[target] < 0 {
		return -1
	}
	return len(nums) - f[target]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nU)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(U)$。

## 专题训练

见下面动态规划题单的「**§3.1 0-1 背包**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)
