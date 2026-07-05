枚举 $j$，那么 $\textit{nums}[i]$ 越大越好，问题变成：

- 计算下标 $i$ 在 $[0,j-k]$ 中的最大的 $\textit{nums}[i]$（记作 $\textit{mx}$）。

$\textit{mx}$ 可以在枚举 $j$ 的同时维护。$j$ 每增加 $1$，范围 $[0,j-k]$ 的右边界就增加 $1$。所以只需考虑用新增的元素 $\textit{nums}[j-k]$ 去更新 $\textit{mx}$ 的最大值。

[本题视频讲解](https://www.bilibili.com/video/BV1qXTC63EQa/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxValidPairSum(self, nums: list[int], k: int) -> int:
        ans = mx = 0
        for j in range(k, len(nums)):
            mx = max(mx, nums[j - k])  # nums[i] 的最大值
            ans = max(ans, mx + nums[j])
        return ans
```

```java [sol-Java]
class Solution {
    public int maxValidPairSum(int[] nums, int k) {
        int ans = 0;
        int mx = 0;
        for (int j = k; j < nums.length; j++) {
            mx = Math.max(mx, nums[j - k]); // nums[i] 的最大值
            ans = Math.max(ans, mx + nums[j]);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxValidPairSum(vector<int>& nums, int k) {
        int ans = 0, mx = 0;
        for (int j = k; j < nums.size(); j++) {
            mx = max(mx, nums[j - k]); // nums[i] 的最大值
            ans = max(ans, mx + nums[j]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxValidPairSum(nums []int, k int) (ans int) {
	mx := 0
	for j := k; j < len(nums); j++ {
		mx = max(mx, nums[j-k]) // nums[i] 的最大值
		ans = max(ans, mx+nums[j])
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n-k)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
