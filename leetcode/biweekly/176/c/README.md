**前置题目**：[198. 打家劫舍](https://leetcode.cn/problems/house-robber/)，[视频讲解](https://www.bilibili.com/video/BV1Xj411K7oF/)。

定义 $f[i+1]$ 表示在 $[0,i]$ 中能得到的最大总和，这里 $+1$ 是为了把 $f[0]$（没有房子可偷的情况）作为初始值。

分类讨论：

- 如果 $\textit{colors}[i-1] = \textit{colors}[i]$，那么和 198 题一样，用「选或不选」思考：
  - 偷 $\textit{nums}[i]$，那么 $\textit{nums}[i-1]$ 不能偷，问题变成在 $[0,i-2]$ 中能得到的最大总和，即 $f[i-1]$，加上 $\textit{nums}[i]$。
  - 不偷 $\textit{nums}[i]$，问题变成在 $[0,i-1]$ 中能得到的最大总和，即 $f[i]$。
  - 二者取最大值，$f[i+1] = \max(f[i-1] + \textit{nums}[i], f[i])$。
- 如果 $\textit{colors}[i-1] \ne \textit{colors}[i]$，那么 $\textit{nums}[i]$ 偷或不偷，都可以偷 $\textit{nums}[i-1]$。由于 $\textit{nums}[i] > 0$，所以偷更好，得 $f[i+1] = f[i] + \textit{nums}[i]$。

初始值：

- $f[0] = 0$，没有房子可偷。
- $f[1] = \textit{nums}[0]$，只有一个房子，偷。

答案为在 $[0,n-1]$ 中能得到的最大总和，即 $f[n]$。

[本题视频讲解](https://www.bilibili.com/video/BV15TZ4B1Eev/?t=5m8s)，欢迎点赞关注~

## 写法一

```py [sol-Python3]
class Solution:
    def rob(self, nums: List[int], colors: List[int]) -> int:
        n = len(nums)
        f = [0] * (n + 1)
        f[1] = nums[0]
        for i in range(1, n):
            if colors[i] != colors[i - 1]:
                f[i + 1] = f[i] + nums[i]
            else:
                f[i + 1] = max(f[i - 1] + nums[i], f[i])  # 选或不选
        return f[n]
```

```java [sol-Java]
class Solution {
    public long rob(int[] nums, int[] colors) {
        int n = nums.length;
        long[] f = new long[n + 1];
        f[1] = nums[0];
        for (int i = 1; i < n; i++) {
            if (colors[i] != colors[i - 1]) {
                f[i + 1] = f[i] + nums[i];
            } else {
                f[i + 1] = Math.max(f[i - 1] + nums[i], f[i]); // 选或不选
            }
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long rob(vector<int>& nums, vector<int>& colors) {
        int n = nums.size();
        vector<long long> f(n + 1);
        f[1] = nums[0];
        for (int i = 1; i < n; i++) {
            if (colors[i] != colors[i - 1]) {
                f[i + 1] = f[i] + nums[i];
            } else {
                f[i + 1] = max(f[i - 1] + nums[i], f[i]); // 选或不选
            }
        }
        return f[n];
    }
};
```

```go [sol-Go]
func rob(nums, colors []int) int64 {
	n := len(nums)
	f := make([]int, n+1)
	f[1] = nums[0]
	for i := 1; i < n; i++ {
		if colors[i] != colors[i-1] {
			f[i+1] = f[i] + nums[i]
		} else {
			f[i+1] = max(f[i-1]+nums[i], f[i]) // 选或不选
		}
	}
	return int64(f[n])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 写法二（空间优化）

```py [sol-Python3]
class Solution:
    def rob(self, nums: List[int], colors: List[int]) -> int:
        f0, f1 = 0, nums[0]
        for i in range(1, len(nums)):
            if colors[i] != colors[i - 1]:
                f0 = f1
                f1 += nums[i]
            else:
                f0, f1 = f1, max(f0 + nums[i], f1)
        return f1
```

```java [sol-Java]
class Solution {
    public long rob(int[] nums, int[] colors) {
        int n = nums.length;
        long f0 = 0;
        long f1 = nums[0];
        for (int i = 1; i < n; i++) {
            if (colors[i] != colors[i - 1]) {
                f0 = f1;
                f1 += nums[i];
            } else {
                long tmp = f1;
                f1 = Math.max(f0 + nums[i], f1);
                f0 = tmp;
            }
        }
        return f1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long rob(vector<int>& nums, vector<int>& colors) {
        int n = nums.size();
        long long f0 = 0, f1 = nums[0];
        for (int i = 1; i < n; i++) {
            if (colors[i] != colors[i - 1]) {
                f0 = f1;
                f1 += nums[i];
            } else {
                long long tmp = f1;
                f1 = max(f0 + nums[i], f1);
                f0 = tmp;
            }
        }
        return f1;
    }
};
```

```go [sol-Go]
func rob(nums, colors []int) int64 {
	n := len(nums)
	f0, f1 := 0, nums[0]
	for i := 1; i < n; i++ {
		if colors[i] != colors[i-1] {
			f0 = f1
			f1 += nums[i]
		} else {
			f0, f1 = f1, max(f0+nums[i], f1)
		}
	}
	return int64(f1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面动态规划题单的「**§1.2 打家劫舍**」。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
