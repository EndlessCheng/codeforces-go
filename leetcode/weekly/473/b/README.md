把 $\textit{nums}$ 每一项平方，然后从小到大排序。

设 $m = \left\lfloor\dfrac{n}{2}\right\rfloor$。贪心地，把前 $m$ 小作为奇数项（减去），后 $n-m$ 大作为偶数项（加上），这样算出来的交替和是最大的。

[本题视频讲解](https://www.bilibili.com/video/BV1eqxNzXE8v/?t=4m35s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxAlternatingSum(self, nums: List[int]) -> int:
        nums = sorted(x * x for x in nums)  # 原地写法见另一份代码

        m = len(nums) // 2
        # 交替和：减去小的，加上大的
        return sum(nums[m:]) - sum(nums[:m])
```

```py [sol-Python3 原地]
class Solution:
    def maxAlternatingSum(self, nums: List[int]) -> int:
        for i, x in enumerate(nums):
            nums[i] *= x
        nums.sort()

        m = len(nums) // 2
        # 交替和：减去小的，加上大的
        return sum(nums[m:]) - sum(nums[:m])
```

```java [sol-Java]
class Solution {
    public long maxAlternatingSum(int[] nums) {
        int n = nums.length;
        for (int i = 0; i < n; i++) {
            nums[i] *= nums[i];
        }
        Arrays.sort(nums);

        int m = n / 2;
        // 交替和：减去小的，加上大的
        long ans = 0;
        for (int i = 0; i < m; i++) {
            ans -= nums[i];
        }
        for (int i = m; i < n; i++) {
            ans += nums[i];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxAlternatingSum(vector<int>& nums) {
        for (int& x : nums) {
            x *= x;
        }
        int m = nums.size() / 2;
        ranges::nth_element(nums, nums.begin() + m);

        // 交替和：减去小的，加上大的
        return -accumulate(nums.begin(), nums.begin() + m, 0LL)
               +accumulate(nums.begin() + m, nums.end(), 0LL);
    }
};
```

```go [sol-Go]
func maxAlternatingSum(nums []int) (ans int64) {
	for i, x := range nums {
		nums[i] *= x
	}
	slices.Sort(nums)

	// 交替和：减去小的，加上大的
	m := len(nums) / 2
	for _, x := range nums[:m] {
		ans -= int64(x)
	}
	for _, x := range nums[m:] {
		ans += int64(x)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。瓶颈在排序上。用**快速选择算法**可以做到 $\mathcal{O}(n)$，见 C++ 代码。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 专题训练

见下面贪心题单的「**§1.1 从最小/最大开始贪心**」。

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
