根据**排序不等式**，取 $\textit{nums}$ 最大的 $k$ 个数，按照从大到小的顺序依次操作。如果 $\textit{mul}\le 1$，则不与 $\textit{mul}$ 相乘。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def maxSum(self, nums: list[int], k: int, mul: int) -> int:
        nums.sort(reverse=True)
        ans = 0
        for x in nums[:k]:
            ans += x * mul if mul > 1 else x
            mul -= 1
        return ans
```

```py [sol-Python3 写法二]
class Solution:
    def maxSum(self, nums: list[int], k: int, mul: int) -> int:
        ans = 0
        for x in nlargest(k, nums):
            ans += x * mul if mul > 1 else x
            mul -= 1
        return ans
```

```java [sol-Java]
public class Solution {
    public static long maxSum(int[] nums, int k, int mul) {
        Arrays.sort(nums);
        int n = nums.length;
        long ans = 0;
        for (int i = n - 1; i >= n - k; i--) {
            ans += (long) nums[i] * Math.max(mul, 1);
            mul--;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSum(vector<int>& nums, int k, int mul) {
        ranges::sort(nums, greater());
        long long ans = 0;
        for (int i = 0; i < k; i++) {
            ans += 1LL * nums[i] * max(mul, 1);
            mul--;
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxSum(nums []int, k int, mul int) (ans int64) {
	slices.SortFunc(nums, func(a, b int) int { return b - a })
	for _, x := range nums[:k] {
		ans += int64(x) * int64(max(mul, 1))
		mul--
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n\log k)$，其中 $n$ 是 $\textit{nums}$ 的长度。瓶颈在排序上。如果用堆维护前 $k$ 大，可以做到 $\mathcal{O}(n\log k)$ 时间。
- 空间复杂度：$\mathcal{O}(1)$ 或 $\mathcal{O}(k)$。忽略排序的栈开销。堆的做法需要 $\mathcal{O}(k)$ 的空间。

## 专题训练

见下面贪心题单的「**§4.3 排序不等式**」。

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
