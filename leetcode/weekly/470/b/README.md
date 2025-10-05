如果 $\textit{nums}$ 所有元素的异或和不为 $0$，那么全选即可，返回 $n$。

否则，如果 $\textit{nums}$ 中有非零元素，那么去掉这个非零元素，就可以使异或和不为零，返回 $n-1$。

否则 $\textit{nums}$ 全为 $0$，无解，返回 $0$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def longestSubsequence(self, nums: List[int]) -> int:
        if all(x == 0 for x in nums):
            return 0  # nums 全为 0，无解

        xor_sum = reduce(xor, nums)
        # 如果 xor_sum 为 0，那么去掉 nums 的一个非零元素，就可以使 xor_sum 不为零
        return len(nums) - (xor_sum == 0)
```

```java [sol-Java]
class Solution {
    public int longestSubsequence(int[] nums) {
        boolean hasNonZero = false;
        int xor = 0;
        for (int x : nums) {
            hasNonZero = hasNonZero || x != 0;
            xor ^= x;
        }
        if (!hasNonZero) {
            return 0; // nums 全为 0，无解
        }

        int ans = nums.length;
        if (xor == 0) {
            ans--; // 去掉 nums 的一个非零元素，就可以使 xor 不为零
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestSubsequence(vector<int>& nums) {
        if (ranges::all_of(nums, [](int x) { return x == 0; })) {
            return 0; // nums 全为 0，无解
        }

        int xor_sum = reduce(nums.begin(), nums.end(), 0, bit_xor<>());
        // 如果 xor_sum 为 0，那么去掉 nums 的一个非零元素，就可以使 xor_sum 不为零
        return nums.size() - (xor_sum == 0);
    }
};
```

```go [sol-Go]
func longestSubsequence(nums []int) int {
	sum, xor := 0, 0
	for _, x := range nums {
		sum += x
		xor ^= x
	}
	if sum == 0 {
		return 0 // nums 全为 0，无解
	}

	ans := len(nums)
	if xor == 0 {
		ans-- // 去掉 nums 的一个非零元素，就可以使 xor 不为零
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面贪心与思维题单的「**§5.2 脑筋急转弯**」。

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
