## 题意

对于 $x=0,1,2,\ldots,k-1$，计算有多少个非空连续子数组的元素积模 $k$ 等于 $x$。

## 初步想法

**枚举**子数组右端点，计算右端点为 $0,1,2,\ldots,n-1$ 时，元素积模 $k$ 等于 $x$ 的子数组有多少个。

## 寻找子问题

回想一下 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)，要计算右端点为 $i$ 的最大子数组和，我们需要知道右端点为 $i-1$ 的最大子数组和。换句话说，右端点为 $i-1$ 的子数组，拼接上 $\textit{nums}[i]$，就得到了右端点为 $i$ 的子数组。

本题也一样，要计算右端点为 $i$ 的、元素积模 $k$ 等于 $x$ 的子数组的个数，需要知道右端点为 $i-1$ 的、元素积模 $k$ 等于 $y$ 的子数组的个数。

比如 $k=5$，右端点为 $i-1$ 的元素积模 $5$ 等于 $2$ 的子数组有 $100$ 个。假设 $\textit{nums}[i]=4$，那么计算 $2\cdot 4\bmod 5 = 3$，我们就得到了 $100$ 个右端点为 $i$ 的元素积模 $5$ 等于 $3$ 的子数组。

⚠**注意**：我们只需要知道右端点为 $i-1$ 的元素积模 $k$ 的结果，至于原始元素积是多少，我们并不关心。比如 $k=5$，原始元素积等于 $10$ 还是 $15$，取模后都是 $0$，这些数（无论是原始元素积还是取模后的元素积）乘以其他数，模 $k$ 后都是 $0$。换句话说，我们用到的是如下恒等式

$$
(a\cdot b) \bmod k=((a\bmod k)\cdot  (b\bmod k)) \bmod k
$$

证明见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

## 动态规划

一般地，定义 $f[i+1][x]$ 表示右端点为 $i$ 的、元素积模 $k$ 等于 $x$ 的子数组的个数。$+1$ 是为了方便用 $f[0]$ 表示初始值。 

设 $v=\textit{nums}[i]$。采用**刷表法**转移：

- 对于右端点为 $i-1$ 的子数组，枚举子数组元素积模 $k$ 等于 $y=0,1,2,\ldots,k-1$，更新右端点为 $i$ 的子数组的状态值，也就是把 $f[i+1][y\cdot v\bmod k]$ 增加 $f[i][y]$。
- 此外，$v$ 可以单独作为一个子数组（长度等于 $1$），把 $f[i+1][v\bmod k]$ 加一。

> **刷表法**：如果枚举 $x$，倒推 $x = y \cdot v\bmod k$ 中的 $y$，是困难的。但枚举 $y$，计算 $y\cdot v \bmod k$，是简单的。也就是说，对于状态 $f[i+1][x]$ 而言，其转移来源是谁不好计算，但从 $f[i][y]$ 转移到的目标状态 $f[i+1][y\cdot v\bmod k]$ 是好计算的。在动态规划中，根据转移来源计算状态叫查表法，用当前状态更新其他状态叫刷表法。

**初始值**：$f[0][0] = 0$，一开始没有任何元素。 

**答案**：$\textit{ans}[x] = \displaystyle\sum\limits_{i=1}^{n}f[i][x]$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1NALczNERr/?t=18m15s)，欢迎点赞关注~

## 空间优化前

```py [sol-Python3]
class Solution:
    def resultArray(self, nums: List[int], k: int) -> List[int]:
        ans = [0] * k
        f = [[0] * k for _ in range(len(nums) + 1)]
        for i, v in enumerate(nums):
            f[i + 1][v % k] = 1
            for y, c in enumerate(f[i]):
                f[i + 1][y * v % k] += c  # 刷表法
            for x, c in enumerate(f[i + 1]):
                ans[x] += c
        return ans
```

```java [sol-Java]
class Solution {
    public long[] resultArray(int[] nums, int k) {
        int n = nums.length;
        long[] ans = new long[k];
        int[][] f = new int[n + 1][k];
        for (int i = 0; i < n; i++) {
            int v = nums[i];
            f[i + 1][v % k] = 1;
            for (int y = 0; y < k; y++) {
                f[i + 1][(int) ((long) y * v % k)] += f[i][y]; // 刷表法
            }
            for (int x = 0; x < k; x++) {
                ans[x] += f[i + 1][x];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> resultArray(vector<int>& nums, int k) {
        int n = nums.size();
        vector<long long> ans(k);
        vector f(n + 1, vector<int>(k));
        for (int i = 0; i < n; i++) {
            int v = nums[i];
            f[i + 1][v % k] = 1;
            for (int y = 0; y < k; y++) {
                f[i + 1][1LL * y * v % k] += f[i][y]; // 刷表法
            }
            for (int x = 0; x < k; x++) {
                ans[x] += f[i + 1][x];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func resultArray(nums []int, k int) []int64 {
	ans := make([]int64, k)
	f := make([][]int, len(nums)+1)
	for i := range f {
		f[i] = make([]int, k)
	}
	for i, v := range nums {
		f[i+1][v%k] = 1
		for y, c := range f[i] {
			f[i+1][y*v%k] += c // 刷表法
		}
		for x, c := range f[i+1] {
			ans[x] += int64(c)
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(nk)$。

## 空间优化

```py [sol-Python3]
class Solution:
    def resultArray(self, nums: List[int], k: int) -> List[int]:
        ans = [0] * k
        f = [0] * k
        for v in nums:
            nf = [0] * k
            nf[v % k] = 1
            for y, c in enumerate(f):
                nf[y * v % k] += c
            f = nf
            for x, c in enumerate(f):
                ans[x] += c
        return ans
```

```java [sol-Java]
class Solution {
    public long[] resultArray(int[] nums, int k) {
        long[] ans = new long[k];
        int[] f = new int[k];
        for (int v : nums) {
            int[] nf = new int[k];
            nf[v % k] = 1;
            for (int y = 0; y < k; y++) {
                nf[(int) ((long) y * v % k)] += f[y];
            }
            f = nf;
            for (int x = 0; x < k; x++) {
                ans[x] += f[x];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> resultArray(vector<int>& nums, int k) {
        vector<long long> ans(k);
        vector<int> f(k);
        for (int v : nums) {
            vector<int> nf(k);
            nf[v % k] = 1;
            for (int y = 0; y < k; y++) {
                nf[1LL * y * v % k] += f[y];
            }
            f = move(nf);
            for (int x = 0; x < k; x++) {
                ans[x] += f[x];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func resultArray(nums []int, k int) []int64 {
	ans := make([]int64, k)
	f := make([]int, k)
	for _, v := range nums {
		nf := make([]int, k)
		nf[v%k] = 1
		for y, c := range f {
			nf[y*v%k] += c
		}
		f = nf
		for x, c := range f {
			ans[x] += int64(c)
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。

## 相似题目（刷表法）

- [2140. 解决智力问题](https://leetcode.cn/problems/solving-questions-with-brainpower/) 1709
- [3448. 统计可以被最后一个数位整除的子字符串数目](https://leetcode.cn/problems/count-substrings-divisible-by-last-digit/) 2387 这题和本题比较像

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
