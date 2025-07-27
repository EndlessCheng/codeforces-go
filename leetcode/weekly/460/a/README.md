设 $m = \dfrac{n}{3}$，我们要从 $\textit{nums}$ 中选一个长为 $m$ 的子序列，最大化这个子序列的元素和。

把 $\textit{nums}$ 从大到小排序，贪心地，取的数字越靠前越好。

我们每次拿出三个数，取第二大的数，所以前面得垫一个数。能取到的数的下标为 $1,3,5,7,\ldots,2m-1$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1pm8vzAEXx/?t=51m15s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maximumMedianSum(self, nums: List[int]) -> int:
        nums.sort(reverse=True)
        m = len(nums) // 3
        return sum(nums[1: m * 2: 2])
```

```java [sol-Java]
class Solution {
    public long maximumMedianSum(int[] nums) {
        Arrays.sort(nums);
        int n = nums.length;
        long ans = 0;
        for (int i = n - 2; i >= n / 3; i -= 2) {
            ans += nums[i];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumMedianSum(vector<int>& nums) {
        ranges::sort(nums, greater());
        int m = nums.size() / 3;
        long long ans = 0;
        for (int i = 1; i < m * 2; i += 2) {
            ans += nums[i];
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumMedianSum(nums []int) (ans int64) {
	slices.SortFunc(nums, func(a, b int) int { return b - a })
	m := len(nums) / 3
	for i := 1; i < m*2; i += 2 {
		ans += int64(nums[i])
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。Python 可以用普通循环代替切片。

## 相似题目

- [1561. 你可以获得的最大硬币数目](https://leetcode.cn/problems/maximum-number-of-coins-you-can-get/)
- [3457. 吃披萨](https://leetcode.cn/problems/eat-pizzas/)

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
