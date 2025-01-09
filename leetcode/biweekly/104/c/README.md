## 方法一：前后缀分解

### 提示 1

要让答案最大，首先应当最大化答案的二进制的**长度**。

### 提示 2

把「乘 $2$」分配给多个数，不如只分配给一个数，这样更有可能得到更大的答案。

### 提示 3

枚举把哪个 $\textit{nums}[i]$ 乘 $k$ 次 $2$（左移 $k$ 次）。

代码实现时，可以仿照 [238. 除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/)，预处理每个 $\textit{nums}[i]$ 左侧元素的或值 $\textit{pre}$，以及右侧元素的或值 $\textit{suf}$，从而快速计算出把 $\textit{nums}[i]$ 乘 $k$ 次 $2$ 后的所有元素的或值。

进一步地，只需要预处理右侧元素的或值，左侧的或值可以一边枚举一边计算。

[视频讲解](https://www.bilibili.com/video/BV1fV4y1r7e6/)

```py [sol1-Python3]
class Solution:
    def maximumOr(self, nums: List[int], k: int) -> int:
        n = len(nums)
        suf = [0] * (n + 1)
        for i in range(n - 1, 0, -1):
            suf[i] = suf[i + 1] | nums[i]

        ans = pre = 0
        for i, x in enumerate(nums):
            ans = max(ans, pre | (x << k) | suf[i + 1])
            pre |= x
        return ans
```

```java [sol1-Java]
class Solution {
    public long maximumOr(int[] nums, int k) {
        int n = nums.length;
        int[] suf = new int[n + 1];
        for (int i = n - 1; i > 0; i--)
            suf[i] = suf[i + 1] | nums[i];

        long ans = 0;
        for (int i = 0, pre = 0; i < n; i++) {
            ans = Math.max(ans, pre | ((long) nums[i] << k) | suf[i + 1]);
            pre |= nums[i];
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long maximumOr(vector<int>& nums, int k) {
        int n = nums.size(), suf[n + 1];
        suf[n] = 0;
        for (int i = n - 1; i; i--)
            suf[i] = suf[i + 1] | nums[i];

        long long ans = 0;
        for (int i = 0, pre = 0; i < n; i++) {
            ans = max(ans, pre | ((long long) nums[i] << k) | suf[i + 1]);
            pre |= nums[i];
        }
        return ans;
    }
};
```

```go [sol1-Go]
func maximumOr(nums []int, k int) int64 {
	n := len(nums)
	suf := make([]int, n+1)
	for i, x := range slices.Backward(nums) {
		suf[i] = suf[i+1] | x
	}

	ans, pre := 0, 0
	for i, x := range nums {
		ans = max(ans, pre|x<<k|suf[i+1])
		pre |= x
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二

设 $\textit{nums}$ 所有数的 OR 的结果为 $\textit{allOr}$。

从 $\textit{allOr}$ 中去掉 $x=\textit{nums}[i]$ 后的数是多少？

1. 先通过 XOR，直接去掉 $x$，即代码中的 `allOr ^ x`。
2. 然后，如果有多个 $\textit{nums}[i]$ 在同一个比特位上都是 $1$，那么去掉 $x$ 后其余 $n-1$ 个数的 OR，在这个比特位上也是 $1$。换句话说，无论去掉哪个 $x$，这些比特位恒为 $1$。把这些恒为 $1$ 的比特位记到二进制数 $\textit{multi}$ 中。

那么去掉 $x$ 后的其余 $n-1$ 个数的 OR 等于 `(allOr ^ x) | multi`。

如何计算 $\textit{multi}$？

用「枚举右，维护左」的思想，在遍历 $\textit{nums}$ 计算 $\textit{allOr}$ 的过程中，`allOr & x` 中的 $1$ 必然出现了两次，将其 OR 到 $\textit{multi}$ 中。

```py [sol-Python3]
class Solution:
    def maximumOr(self, nums: List[int], k: int) -> int:
        all_or = multi = 0
        for x in nums:
            multi |= all_or & x
            all_or |= x
        return max((x << k) | (all_or ^ x) | multi for x in nums)
```

```java [sol-Java]
class Solution {
    public long maximumOr(int[] nums, int k) {
        int allOr = 0, multi = 0;
        for (int x : nums) {
            multi |= allOr & x;
            allOr |= x;
        }

        long ans = 0;
        for (int x : nums) {
            ans = Math.max(ans, ((long) x << k) | (allOr ^ x) | multi);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumOr(vector<int>& nums, int k) {
        int all_or = 0, multi = 0;
        for (int x : nums) {
            multi |= all_or & x;
            all_or |= x;
        }

        long long ans = 0;
        for (int x : nums) {
            ans = max(ans, ((long long) x << k) | (all_or ^ x) | multi);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumOr(nums []int, k int) int64 {
	allOr, multi := 0, 0
	for _, x := range nums {
		multi |= allOr & x
		allOr |= x
	}

	ans := 0
	for _, x := range nums {
		ans = max(ans, x<<k|(allOr^x)|multi)
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

### 思考题

把 $2$ 换成其它数，方法是否一样？

如果题目限制乘法的结果不能超过 $\textit{limit}$ 呢？

把乘法改成除法，把 OR 改成 AND，要怎么做？

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
