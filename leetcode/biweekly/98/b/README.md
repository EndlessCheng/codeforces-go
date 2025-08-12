为方便计算，把 $\textit{nums}$ 从小到大排序。

如果什么也不改，那么最大得分为 $\textit{nums}[n-1] - \textit{nums}[0]$。

在允许修改两个数的情况下，最大得分可以是哪两个数的差？**哪些数字一定不可能参与最大得分的计算**？

贪心地，有如下修改方案：

- 修改最大的两个数为 $\textit{nums}[n-3]$，最大得分为 $\textit{nums}[n-3]-\textit{nums}[0]$。
- 修改最小的为 $\textit{nums}[1]$，最大的为 $\textit{nums}[n-2]$，最大得分为 $\textit{nums}[n-2]-\textit{nums}[1]$。
- 修改最小的两个数为 $\textit{nums}[2]$，最大得分为 $\textit{nums}[n-1]-\textit{nums}[2]$。

对于更内部的元素，比如 $\textit{nums}[3]$，一定不可能参与最大得分的计算。反证法：如果 $\textit{nums}[3]$ 参与最大得分的计算，这意味着我们把 $\textit{nums}[0],\textit{nums}[1],\textit{nums}[2]$ 都改了，不符合题目要求，矛盾。

所以最大得分就是上述三种情况的最小值。

此外可以发现，上述三种方案产生了相同的元素，所以最小得分等于 $0$。

综上所述，答案为

$$
\min(\textit{nums}[n-3]-\textit{nums}[0], \textit{nums}[n-2]-\textit{nums}[1],\textit{nums}[n-1]-\textit{nums}[2])
$$

[视频讲解](https://www.bilibili.com/video/BV15D4y1G7ms/)

```py [sol-Python3]
class Solution:
    def minimizeSum(self, a: List[int]) -> int:
        a.sort()
        return min(a[-3] - a[0], a[-2] - a[1], a[-1] - a[2])
```

```java [sol-Java]
class Solution {
    public int minimizeSum(int[] a) {
        Arrays.sort(a);
        int n = a.length;
        return Math.min(Math.min(a[n - 3] - a[0], a[n - 2] - a[1]), a[n - 1] - a[2]);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimizeSum(vector<int>& a) {
        ranges::sort(a);
        int n = a.size();
        return min({a[n - 3] - a[0], a[n - 2] - a[1], a[n - 1] - a[2]});
    }
};
```

```go [sol-Go]
func minimizeSum(a []int) int {
	slices.Sort(a)
	n := len(a)
	return min(a[n-3]-a[0], a[n-2]-a[1], a[n-1]-a[2])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。手动维护或者用快速选择可以做到 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序时栈的开销。

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
