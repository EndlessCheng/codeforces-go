**前置知识**：[差分数组](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)。

定义 $d[i] = \textit{nums}[i+1] - \textit{nums}[i]$。

- 示例 1 的 $d = [0,-1,-1]$。
- 示例 2 的 $d = [-4,1,1]$。

> **注**：题目要求 $\textit{nums}$ 递增，增大 $\textit{nums}[0]$ 无意义，所以 $\textit{nums}[0]$ 不需要在差分数组 $d$ 中。

题目让我们把 $\textit{nums}$ 变成递增的，这等价于把每个 $d[i]$ 都变成**非负数**。

根据前置知识，把 $\textit{nums}$ 的某个子数组中的元素都增加 $x$，等价于：

- 把差分数组中的一个数增加 $x$。
- 如果子数组不是 $\textit{nums}$ 的后缀，那么还要把差分数组的另一个数减少 $x$。

由于我们只需把 $d$ 中的负数变成非负数，减少 $d[i]$ 无意义，所以每次**只需操作 $\textit{nums}$ 的后缀**，等价于把某个 $d[i]$ 增加 $x$。

要把每个 $d[i]$ 都变成非负数，只需把负数 $d[i]$ 增加 $-d[i]$。

所以答案为

$$
\sum_{i=0}^{n-2} \max(-d[i], 0) = \sum_{i=1}^{n-1} \max(\textit{nums}[i-1] - \textit{nums}[i], 0)
$$

[本题视频讲解](https://www.bilibili.com/video/BV1xzZcBZEpe/?t=9m44s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: list[int]) -> int:
        return sum(max(x - y, 0) for x, y in pairwise(nums))
```

```java [sol-Java]
class Solution {
    public long minOperations(int[] nums) {
        long ans = 0;
        for (int i = 1; i < nums.length; i++) {
            ans += Math.max(nums[i - 1] - nums[i], 0);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minOperations(vector<int>& nums) {
        long long ans = 0;
        for (int i = 1; i < nums.size(); i++) {
            ans += max(nums[i - 1] - nums[i], 0);
        }
        return ans;
    }
};
```

```go [sol-Go]
func minOperations(nums []int) (ans int64) {
	for i := 1; i < len(nums); i++ {
		ans += int64(max(nums[i-1]-nums[i], 0))
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

- [1526. 形成目标数组的子数组最少增加次数](https://leetcode.cn/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/)
- [3229. 使数组等于目标数组所需的最少操作次数](https://leetcode.cn/problems/minimum-operations-to-make-array-equal-to-target/)

## 专题训练

见下面数据结构题单的「**§2.1 一维差分**」。

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
