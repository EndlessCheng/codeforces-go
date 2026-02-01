由于 Alice 先手，那么 Alice 选择删除 $[1,n-1]$，可以剩下 $\textit{nums}[0]$；选择删除 $[0,n-2]$，可以剩下 $\textit{nums}[n-1]$。由于只剩下一个数，游戏结束。

Alice 可以取这两种情况的最大值。

所以答案 $\ge \max(\textit{nums}[0], \textit{nums}[n-1])$。

下面证明，**大于号是取不到的**。

**反证法**，假设答案严格大于 $\max(\textit{nums}[0], \textit{nums}[n-1])$。

设最终剩下的数为 $x = \textit{nums}[i]$。由于 $x > \max(\textit{nums}[0], \textit{nums}[n-1])$，所以 $x$ 不在 $\textit{nums}$ 的两端，即 $1\le i\le n-2$。

讨论前两轮的操作：

- 第一轮，Alice 操作。Alice 不能移除 $x$，只能移除其他不含 $x$ 的子数组。这个子数组要么位于 $x$ 的左侧，要么位于 $x$ 的右侧。由于 $1\le i\le n-2$，所以移除子数组后，一定剩下至少两个数（游戏没有结束），且至少剩下 $\textit{nums}[0]$ 和 $\textit{nums}[n-1]$ 中的一个数。
- 第二轮，Bob 操作。如果剩下的数包含 $\textit{nums}[0]$，那么 Bob 可以删除其他数，最终剩下 $\textit{nums}[0]$。如果剩下的数包含 $\textit{nums}[n-1]$，那么 Bob 可以删除其他数，最终剩下 $\textit{nums}[n-1]$。矛盾，所以答案不可能严格大于 $\max(\textit{nums}[0], \textit{nums}[n-1])$。

综上所述，Alice 的最佳策略是第一回合就结束游戏（Bob：请输入文字），答案为

$$
\max(\textit{nums}[0], \textit{nums}[n-1])
$$

[本题视频讲解](https://www.bilibili.com/video/BV1hd64BcEBQ/?t=5m15s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def finalElement(self, nums: List[int]) -> int:
        return max(nums[0], nums[-1])
```

```java [sol-Java]
class Solution {
    public int finalElement(int[] nums) {
        return Math.max(nums[0], nums[nums.length - 1]);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int finalElement(vector<int>& nums) {
        return max(nums[0], nums.back());
    }
};
```

```go [sol-Go]
func finalElement(nums []int) int {
	return max(nums[0], nums[len(nums)-1])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
