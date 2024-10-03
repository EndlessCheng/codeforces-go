从最大的元素开始思考：

- 数组的最大值 $m$ 不变是最好的。反证：如果把 $m$ 变小是最优的，那么把 $m$ 恢复成其原来的值，仍然满足题目要求，且我们得到了更优的答案，矛盾。
- 数组的次大值呢？如果它等于 $m$，那么它必须变成 $m-1$，否则不变。
- 依此类推。

为了方便计算，先把数组从大到小排序，那么 $\textit{maximumHeight}[i]$ 的实际值为

$$
\min(\textit{maximumHeight}[i], \textit{maximumHeight}[i-1] - 1)
$$

如果元素值 $\le 0$，不符合题目要求，返回 $-1$。

最终答案为 $\textit{maximumHeight}$ 的元素之和。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1bjxyewEQV/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maximumTotalSum(self, maximumHeight: List[int]) -> int:
        maximumHeight.sort(reverse=True)
        for i in range(1, len(maximumHeight)):
            maximumHeight[i] = min(maximumHeight[i], maximumHeight[i - 1] - 1)
            if maximumHeight[i] <= 0:
                return -1
        return sum(maximumHeight)
```

```java [sol-Java]
class Solution {
    public long maximumTotalSum(int[] maximumHeight) {
        Arrays.sort(maximumHeight); // 从小到大排序，下面倒着遍历
        int n = maximumHeight.length;
        long ans = maximumHeight[n - 1];
        for (int i = n - 2; i >= 0; i--) {
            maximumHeight[i] = Math.min(maximumHeight[i], maximumHeight[i + 1] - 1);
            if (maximumHeight[i] <= 0) {
                return -1;
            }
            ans += maximumHeight[i];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumTotalSum(vector<int>& maximumHeight) {
        ranges::sort(maximumHeight, greater()); // 从大到小排序
        for (int i = 1; i < maximumHeight.size(); i++) {
            maximumHeight[i] = min(maximumHeight[i], maximumHeight[i - 1] - 1);
            if (maximumHeight[i] <= 0) {
                return -1;
            }
        }
        return reduce(maximumHeight.begin(), maximumHeight.end(), 0LL);
    }
};
```

```go [sol-Go]
func maximumTotalSum(maximumHeight []int) int64 {
	slices.SortFunc(maximumHeight, func(a, b int) int { return b - a })
	ans := maximumHeight[0]
	for i := 1; i < len(maximumHeight); i++ {
		maximumHeight[i] = min(maximumHeight[i], maximumHeight[i-1]-1)
		if maximumHeight[i] <= 0 {
			return -1
		}
		ans += maximumHeight[i]
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{maximumHeight}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 思考题

推荐做做 [1840. 最高建筑高度](https://leetcode.cn/problems/maximum-building-height/)，作为本题的思考题。

更多相似题目，见下面贪心题单中的「**§1.1 从最小/最大开始贪心**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
