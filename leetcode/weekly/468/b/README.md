由于同一个子数组可以重复选，所以最优解是，把差值最大的子数组重复选 $k$ 次。

最大差值等于 $\textit{nums}$ 的最大值减去 $\textit{nums}$ 的最小值。

```py [sol-Python3]
class Solution:
    def maxTotalValue(self, nums: List[int], k: int) -> int:
        return (max(nums) - min(nums)) * k
```

```java [sol-Java]
class Solution {
    public long maxTotalValue(int[] nums, int k) {
        int mn = Integer.MAX_VALUE;
        int mx = 0;
        for (int x : nums) {
            mn = Math.min(mn, x);
            mx = Math.max(mx, x);
        }
        return (long) (mx - mn) * k;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxTotalValue(vector<int>& nums, int k) {
        auto [mn, mx] = ranges::minmax(nums);
        return 1LL * (mx - mn) * k;
    }
};
```

```go [sol-Go]
func maxTotalValue(nums []int, k int) int64 {
	return int64(slices.Max(nums)-slices.Min(nums)) * int64(k)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

计算数组的最小值和最大值，最多需要多少次比大小？

C++ `std::minmax` 的 [文档](https://en.cppreference.com/w/cpp/algorithm/minmax.html) 写着：至多要 $\dfrac{3N}{2}$ 次。

如何实现？见 [CF730B](https://codeforces.com/problemset/problem/730/B)。

## 专题训练

见下面思维题单的「**§5.2 脑筋急转弯**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
