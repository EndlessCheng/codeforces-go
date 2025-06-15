题目意思是首尾元素下标相差至少是 $m-1$，这样子序列长度才能是 $m$（否则长度就小于 $m$ 了）。

我们只需关注首尾两个数，于是问题转化成：

- $\textit{nums}$ 的任意下标相差至少为 $m-1$ 的**两数之积**的最大值。

枚举 $\textit{nums}[i]$，为了让乘积最大，贪心地，维护 $[0,i-m+1]$ 中的最小值和最大值。维护最小值是因为负负得正，答案可以来自两个负数相乘。

注意最终答案可能是负数，要把答案初始化成 $-\infty$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def maximumProduct(self, nums: List[int], m: int) -> int:
        ans = mx = -inf
        mn = inf
        for i in range(m - 1, len(nums)):
            # 维护左边 [0,i-m+1] 中的最小值和最大值
            y = nums[i - m + 1]
            mn = min(mn, y)
            mx = max(mx, y)
            # 枚举右
            x = nums[i]
            ans = max(ans, x * mn, x * mx)
        return ans
```

```java [sol-Java]
class Solution {
    public long maximumProduct(int[] nums, int m) {
        long ans = Long.MIN_VALUE;
        int mn = Integer.MAX_VALUE;
        int mx = Integer.MIN_VALUE;
        for (int i = m - 1; i < nums.length; i++) {
            // 维护左边 [0,i-m+1] 中的最小值和最大值
            int y = nums[i - m + 1];
            mn = Math.min(mn, y);
            mx = Math.max(mx, y);
            // 枚举右
            long x = nums[i];
            ans = Math.max(ans, Math.max(x * mn, x * mx));
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumProduct(vector<int>& nums, int m) {
        long long ans = LLONG_MIN;
        int mn = INT_MAX, mx = INT_MIN;
        for (int i = m - 1; i < nums.size(); i++) {
            // 维护左边 [0,i-m+1] 中的最小值和最大值
            int y = nums[i - m + 1];
            mn = min(mn, y);
            mx = max(mx, y);
            // 枚举右
            long long x = nums[i];
            ans = max({ans, x * mn, x * mx});
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumProduct(nums []int, m int) int64 {
	ans := math.MinInt
	mn, mx := math.MaxInt, math.MinInt
	for i := m - 1; i < len(nums); i++ {
		// 维护左边 [0,i-m+1] 中的最小值和最大值
		y := nums[i-m+1]
		mn = min(mn, y)
		mx = max(mx, y)
		// 枚举右
		x := nums[i]
		ans = max(ans, x*mn, x*mx)
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n-m)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

见下面数据结构题单的「**§0.1 枚举右，维护左**」。

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
