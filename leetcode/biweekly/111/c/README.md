## 前言

正难则反，计算最多**保留**多少个数。这些保留的数必须是非递减的。

## 方法一：最长非递减子序列

**前置题目**：[300. 最长递增子序列](https://leetcode.cn/problems/longest-increasing-subsequence/)

计算 $\textit{nums}$ 的最长非递减子序列的长度。

做法请看 [最长递增子序列【基础算法精讲 20】](https://www.bilibili.com/video/BV1ub411Q7sB/)，视频末尾讲了如何处理非递减的情况。

```py [sol-Python3]
class Solution:
    def minimumOperations(self, nums: List[int]) -> int:
        g = []
        for x in nums:
            j = bisect_right(g, x)
            if j == len(g):
                g.append(x)
            else:
                g[j] = x
        return len(nums) - len(g)
```

```java [sol-Java]
class Solution {
    public int minimumOperations(List<Integer> nums) {
        List<Integer> g = new ArrayList<>();
        for (int x : nums) {
            int j = upperBound(g, x);
            if (j == g.size()) {
                g.add(x);
            } else {
                g.set(j, x);
            }
        }
        return nums.size() - g.size();
    }

    // 开区间写法
    private int upperBound(List<Integer> g, int target) {
        int left = -1, right = g.size(); // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] <= target
            // nums[right] > target
            int mid = (left + right) >>> 1;
            if (g.get(mid) > target) {
                right = mid; // 范围缩小到 (left, mid)
            } else {
                left = mid; // 范围缩小到 (mid, right)
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumOperations(vector<int>& nums) {
        vector<int> g;
        for (int x : nums) {
            auto it = ranges::upper_bound(g, x);
            if (it == g.end()) {
                g.push_back(x);
            } else {
                *it = x;
            }
        }
        return nums.size() - g.size();
    }
};
```

```go [sol-Go]
func minimumOperations(nums []int) int {
	g := []int{}
	for _, x := range nums {
		p := sort.SearchInts(g, x+1)
		if p < len(g) {
			g[p] = x
		} else {
			g = append(g, x)
		}
	}
	return len(nums) - len(g)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：多维 DP

由于值域很小，考虑把值域当作 DP 的一个维度。

定义 $f[i+1][j]$ 表示 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 的最长非递减子序列的长度，其中子序列最后一个数 $\le j$。

$f[i+1][j]$ 怎么算？设 $x=\textit{nums}[i]$，分类讨论：

- 不选 $x$，问题变成 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$ 的最长非递减子序列的长度，其中子序列最后一个数 $\le j$，即 $f[i+1][j] = f[i][j]$。
- 选 $x$，也就是把 $x$ 当作子序列最后一个数，需要满足 $x\le j$。问题变成 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$ 的最长非递减子序列的长度，其中子序列最后一个数 $\le x$，即 $f[i+1][j] = f[i][x] + 1$。

两种情况取最大值，有

$$
f[i+1][j] =
\begin{cases}
f[i][j], & j < x     \\
\max(f[i][j], f[i][x] + 1), & j\ge x     \\
\end{cases}
$$

初始值 $f[0][j] = 0$。

答案为 $n - f[n][3]$。

### 优化前

```py [sol-Python3]
class Solution:
    def minimumOperations(self, nums: List[int]) -> int:
        n = len(nums)
        f = [[0] * 4 for _ in range(n + 1)]
        for i, x in enumerate(nums):
            for j in range(1, 4):
                if j < x:
                    f[i + 1][j] = f[i][j]
                else:
                    f[i + 1][j] = max(f[i][j], f[i][x] + 1)
        return n - f[n][3]
```

```java [sol-Java]
class Solution {
    public int minimumOperations(List<Integer> nums) {
        int n = nums.size();
        int[][] f = new int[n + 1][4];
        for (int i = 0; i < n; i++) {
            int x = nums.get(i);
            for (int j = 1; j <= 3; j++) {
                if (j < x) {
                    f[i + 1][j] = f[i][j];
                } else {
                    f[i + 1][j] = Math.max(f[i][j], f[i][x] + 1);
                }
            }
        }
        return n - f[n][3];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumOperations(vector<int>& nums) {
        int n = nums.size();
        vector<array<int, 4>> f(n + 1);
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            for (int j = 1; j <= 3; j++) {
                if (j < x) {
                    f[i + 1][j] = f[i][j];
                } else {
                    f[i + 1][j] = max(f[i][j], f[i][x] + 1);
                }
            }
        }
        return n - f[n][3];
    }
};
```

```go [sol-Go]
func minimumOperations(nums []int) int {
	n := len(nums)
	f := make([][4]int, n+1)
	for i, x := range nums {
		for j := 1; j <= 3; j++ {
			if j < x {
				f[i+1][j] = f[i][j]
			} else {
				f[i+1][j] = max(f[i][j], f[i][x]+1)
			}
		}
	}
	return n - f[n][3]
}
```

### 空间优化

去掉第一个维度，有

$$
f[j] =
\begin{cases}
f[j], & j < x     \\
\max(f[j], f[x] + 1), & j\ge x     \\
\end{cases}
$$

把 $\texttt{for}$ 循环展开，就是：

1. 先把 $f[x]$ 加一。
2. 然后更新 $f[2] = \max(f[2], f[1])$。如果 $x>1$，那么 $f[2]$ 不变。
3. 然后更新 $f[3] = \max(f[3], f[2])$。如果 $x>2$，那么 $f[3]$ 不变。

```py [sol-Python3]
class Solution:
    def minimumOperations(self, nums: List[int]) -> int:
        f = [0] * 4
        for x in nums:
            f[x] += 1
            f[2] = max(f[2], f[1])
            f[3] = max(f[3], f[2])
        return len(nums) - f[3]
```

```java [sol-Java]
class Solution {
    public int minimumOperations(List<Integer> nums) {
        int[] f = new int[4];
        for (int x : nums) {
            f[x]++;
            f[2] = Math.max(f[2], f[1]);
            f[3] = Math.max(f[3], f[2]);
        }
        return nums.size() - f[3];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumOperations(vector<int>& nums) {
        int f[4]{};
        for (int x: nums) {
            f[x]++;
            f[2] = max(f[2], f[1]);
            f[3] = max(f[3], f[2]);
        }
        return nums.size() - f[3];
    }
};
```

```go [sol-Go]
func minimumOperations(nums []int) int {
	f := [4]int{}
	for _, x := range nums {
		f[x]++
		f[2] = max(f[2], f[1])
		f[3] = max(f[3], f[2])
	}
	return len(nums) - f[3]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nU)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})=3$。
- 空间复杂度：$\mathcal{O}(U)$。

## 方法三：合法子序列 DP

这是一个固定的套路，见动态规划题单中的「§7.2 合法子序列 DP」。

一般定义 $f[x]$ 表示以元素 $x$ 结尾的合法子序列的最长长度/个数/元素和，从子序列的倒数第二个数转移过来。

本题倒数第二个数记作 $j$，那么必须满足 $j\le x$。

转移方程为

$$
f[x] = \max_{j=1}^{x} f[j] + 1
$$

其中 $+1$ 表示在以 $j$ 结尾的子序列的末尾添加一个 $x$，得到以 $x$ 结尾的子序列。

初始值 $f[x] = 0$。

答案为 $n - \max(f)$。

```py [sol-Python3]
class Solution:
    def minimumOperations(self, nums: List[int]) -> int:
        f = [0] * 4
        for x in nums:
            f[x] = max(f[1: x + 1]) + 1
        return len(nums) - max(f)
```

```java [sol-Java]
class Solution {
    public int minimumOperations(List<Integer> nums) {
        int[] f = new int[4];
        for (int x : nums) {
            int mx = 0;
            for (int j = 1; j <= x; j++) {
                mx = Math.max(mx, f[j]);
            }
            f[x] = mx + 1;
        }
        return nums.size() - Arrays.stream(f).max().getAsInt();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumOperations(vector<int>& nums) {
        int f[4]{};
        for (int x : nums) {
            f[x] = *max_element(f + 1, f + x + 1) + 1;
        }
        return nums.size() - ranges::max(f);
    }
};
```

```go [sol-Go]
func minimumOperations(nums []int) int {
	f := [4]int{}
	for _, x := range nums {
		f[x] = slices.Max(f[1:x+1]) + 1
	}
	return len(nums) - slices.Max(f[:])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nU)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})=3$。
- 空间复杂度：$\mathcal{O}(U)$。

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
