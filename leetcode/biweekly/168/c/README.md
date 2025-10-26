## 理解题意

由于 $\textit{nums}_2$ 的长度恰好比 $\textit{nums}_1$ 多一，所以恰好有一个元素要追加到 $\textit{nums}_1$ 的末尾。

其余元素必须一一对应，对于 $[0,n-1]$ 中的 $i$，$x = \textit{nums}_1[i]$ 要变成 $y = \textit{nums}_2[i]$。

## 思路

设 $\textit{target}$ 是 $\textit{nums}_2$ 的最后一个数。

贪心地想，能不能在从 $x$ 变成 $y$ 的过程中，**顺带**把 $\textit{target}$ 追加到 $\textit{nums}_1$ 的末尾？

不妨设 $x\le y$，分类讨论：

- 如果 $\textit{target}$ 在 $x$ 和 $y$ 之间，由于每次操作只能把 $x$ 增大/减少 $1$，只能一步一步挪动，所以可以在 $x$ 变成 $\textit{target}$ 的那一刻，把 $\textit{target}$ 追加到 $\textit{nums}_1$ 的末尾。
- 如果 $\textit{target} < x$，那么先把 $x$ 追加到 $\textit{nums}_1$ 的末尾，再花费 $x-\textit{target}$ 的操作次数。
- 如果 $\textit{target} > y$，那么先把 $x$ 变成 $y$，把 $y$ 追加到 $\textit{nums}_1$ 的末尾，再花费 $\textit{target}-y$ 的操作次数。

对于每一对 $(x,y)$，计算处理 $\textit{target}$ 操作次数，取最小值。

[本题视频讲解](https://www.bilibili.com/video/BV1zxxNzcERu/?t=9m22s)，欢迎点赞关注~

```py [sol-Python3]
# 手写 min max 更快
min = lambda a, b: b if b < a else a
max = lambda a, b: b if b > a else a

class Solution:
    def minOperations(self, nums1: List[int], nums2: List[int]) -> int:
        target = nums2[-1]
        ans = 1  # 把元素追加到 nums1 的末尾需要一次操作
        mn = inf
        for x, y in zip(nums1, nums2):
            if x > y:
                x, y = y, x  # 保证 x <= y，简化后续逻辑
            ans += y - x
            # 如果 target 在 [x,y] 中，那么在从 x 变成 y 的过程中，可以顺带把 target 追加到 nums1 的末尾，代价为 0
            # 如果 target < x，代价为 x-target
            # 如果 target > y，代价为 target-y
            if mn > 0:  # 如果 target 还不在任何 [x,y] 中，则计算
                mn = min(mn, max(x - target, target - y))
        return ans + max(mn, 0)  # 如果 target 在 [x,y] 中，上面可能会算出负数
```

```java [sol-Java]
class Solution {
    public long minOperations(int[] nums1, int[] nums2) {
        int n = nums1.length;
        int target = nums2[n];
        long ans = 1; // 把元素追加到 nums1 的末尾需要一次操作
        int mn = Integer.MAX_VALUE;
        for (int i = 0; i < n; i++) {
            // 保证 x <= y，简化后续逻辑
            int x = Math.min(nums1[i], nums2[i]);
            int y = Math.max(nums1[i], nums2[i]);
            ans += y - x;
            // 如果 target 在 [x,y] 中，那么在从 x 变成 y 的过程中，可以顺带把 target 追加到 nums1 的末尾，代价为 0
            // 如果 target < x，代价为 x-target
            // 如果 target > y，代价为 target-y
            mn = Math.min(mn, Math.max(x - target, target - y));
        }
        return ans + Math.max(mn, 0); // 如果 target 在 [x,y] 中，上面可能会算出负数
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minOperations(vector<int>& nums1, vector<int>& nums2) {
        int target = nums2.back();
        long long ans = 1; // 把元素追加到 nums1 的末尾需要一次操作
        int mn = INT_MAX;
        for (int i = 0; i < nums1.size(); i++) {
            int x = nums1[i], y = nums2[i];
            if (x > y) {
                swap(x, y); // 保证 x <= y，简化后续逻辑
            }
            ans += y - x;
            // 如果 target 在 [x,y] 中，那么在从 x 变成 y 的过程中，可以顺带把 target 追加到 nums1 的末尾，代价为 0
            // 如果 target < x，代价为 x-target
            // 如果 target > y，代价为 target-y
            mn = min(mn, max(x - target, target - y));
        }
        return ans + max(mn, 0); // 如果 target 在 [x,y] 中，上面可能会算出负数
    }
};
```

```go [sol-Go]
func minOperations(nums1, nums2 []int) int64 {
	target := nums2[len(nums2)-1]
	ans := 1 // 把元素追加到 nums1 的末尾需要一次操作
	mn := math.MaxInt
	for i, x := range nums1 {
		y := nums2[i]
		if x > y {
			x, y = y, x // 保证 x <= y，简化后续逻辑
		}
		ans += y - x
		// 如果 target 在 [x,y] 中，那么在从 x 变成 y 的过程中，可以顺带把 target 追加到 nums1 的末尾，代价为 0
		// 如果 target < x，代价为 x-target
		// 如果 target > y，代价为 target-y
		mn = min(mn, max(x-target, target-y))
	}
	return int64(ans + max(mn, 0)) // 如果 target 在 [x,y] 中，上面可能会算出负数
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}_1$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

1. 如果 $\textit{nums}_2$ 的长度是 $m\ (m>n)$，怎么做？
2. 如果 $\textit{nums}_2$ 的长度是 $n+1$，但操作三改成可以插在 $\textit{nums}_1$ 的任意位置，怎么做？

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
