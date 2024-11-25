请看 [视频讲解](https://www.bilibili.com/video/BV18j411b7v4/) 第二题。

## 方法一：枚举 j

枚举 $j$，为了让 $(\textit{nums}[i] - \textit{nums}[j]) * \textit{nums}[k]$ 尽量大，我们需要知道 $j$ 左侧元素的最大值，和 $j$ 右侧元素的最大值。

也就是 $\textit{nums}$ 的前缀最大值 $\textit{preMax}$ 和后缀最大值 $\textit{sufMax}$，这都可以用递推预处理出来：

- $\textit{preMax}[i] = \max(\textit{preMax}[i-1], \textit{nums}[i])$
- $\textit{sufMax}[i] = \max(\textit{sufMax}[i+1], \textit{nums}[i])$

代码实现时，可以只预处理 $\textit{sufMax}$ 数组，$\textit{preMax}$ 可以在计算答案的同时算出来。

```py [sol-Python3]
class Solution:
    def maximumTripletValue(self, nums: List[int]) -> int:
        n = len(nums)
        suf_max = [0] * (n + 1)
        for i in range(n - 1, 1, -1):
            suf_max[i] = max(suf_max[i + 1], nums[i])
        ans = pre_max = 0
        for j, x in enumerate(nums):
            ans = max(ans, (pre_max - x) * suf_max[j + 1])
            pre_max = max(pre_max, x)
        return ans
```

```java [sol-Java]
class Solution {
    public long maximumTripletValue(int[] nums) {
        int n = nums.length;
        int[] sufMax = new int[n + 1];
        for (int i = n - 1; i > 1; i--) {
            sufMax[i] = Math.max(sufMax[i + 1], nums[i]);
        }
        long ans = 0;
        int preMax = nums[0];
        for (int j = 1; j < n - 1; j++) {
            ans = Math.max(ans, (long) (preMax - nums[j]) * sufMax[j + 1]);
            preMax = Math.max(preMax, nums[j]);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumTripletValue(vector<int>& nums) {
        int n = nums.size();
        vector<int> suf_max(n + 1, 0);
        for (int i = n - 1; i > 1; i--) {
            suf_max[i] = max(suf_max[i + 1], nums[i]);
        }
        long long ans = 0;
        int pre_max = nums[0];
        for (int j = 1; j < n - 1; j++) {
            ans = max(ans, (long long) (pre_max - nums[j]) * suf_max[j + 1]);
            pre_max = max(pre_max, nums[j]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumTripletValue(nums []int) int64 {
	ans := 0
	n := len(nums)
	sufMax := make([]int, n+1)
	for i := n - 1; i > 1; i-- {
		sufMax[i] = max(sufMax[i+1], nums[i])
	}
	preMax := 0
	for j, x := range nums {
		ans = max(ans, (preMax-x)*sufMax[j+1])
		preMax = max(preMax, x)
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：枚举 k

枚举 $k$，我们需要知道 $k$ 左边 $\textit{nums}[i] - \textit{nums}[j]$ 的最大值。

类似 [121. 买卖股票的最佳时机](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/)，我们可以在遍历的过程中，维护 $\textit{nums}[i]$ 的最大值 $\textit{preMax}$，同时维护 $\textit{preMax}$ 减当前元素的最大值 $\textit{maxDiff}$，这就是 $k$ 左边 $\textit{nums}[i] - \textit{nums}[j]$ 的最大值。

```py [sol-Python3]
class Solution:
    def maximumTripletValue(self, nums: List[int]) -> int:
        ans = max_diff = pre_max = 0
        for x in nums:
            ans = max(ans, max_diff * x)
            max_diff = max(max_diff, pre_max - x)
            pre_max = max(pre_max, x)
        return ans
```

```java [sol-Java]
class Solution {
    public long maximumTripletValue(int[] nums) {
        long ans = 0;
        int maxDiff = 0;
        int preMax = 0;
        for (int x : nums) {
            ans = Math.max(ans, (long) maxDiff * x);
            maxDiff = Math.max(maxDiff, preMax - x);
            preMax = Math.max(preMax, x);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumTripletValue(vector<int>& nums) {
        long long ans = 0;
        int max_diff = 0, pre_max = 0;
        for (int x : nums) {
            ans = max(ans, (long long) max_diff * x);
            max_diff = max(max_diff, pre_max - x);
            pre_max = max(pre_max, x);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumTripletValue(nums []int) int64 {
	ans, maxDiff, preMax := 0, 0, 0
	for _, x := range nums {
		ans = max(ans, maxDiff*x)
		maxDiff = max(maxDiff, preMax-x)
		preMax = max(preMax, x)
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

如果 $\textit{nums}$ 中有负数要怎么做？

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
