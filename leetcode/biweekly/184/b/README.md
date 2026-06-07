一个灯泡可以照亮 $3$ 个位置。为了满足总照明度至少为 $\textit{brightness}$，我们需要开启

$$
\textit{bulbs} = \left\lceil\dfrac{\textit{brightness}}{3}\right\rceil = \left\lfloor\dfrac{\textit{brightness} + 2}{3}\right\rfloor
$$

个灯泡。

有多少个单位时间被至少一个区间覆盖？这等价于 [56. 合并区间](https://leetcode.cn/problems/merge-intervals/) 之后的区间长度之和 $\textit{sumLen}$，请看 [我的题解](https://leetcode.cn/problems/merge-intervals/solutions/2798138/jian-dan-zuo-fa-yi-ji-wei-shi-yao-yao-zh-f2b3/)。

答案为 $\textit{bulbs}\cdot \textit{sumLen}$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def minEnergy(self, _, brightness: int, intervals: list[list[int]]) -> int:
        intervals.sort(key=lambda p: p[0])  # 按照左端点从小到大排序

        # 56. 合并区间（只计算区间长度之和）
        sum_len = 0
        left, right = 0, -1
        for l, r in intervals:
            if l <= right:  # 左端点在合并区间内，可以合并
                right = max(right, r)  # 更新合并区间的右端点
            else:  # 不相交，无法合并
                sum_len += right - left + 1
                left, right = l, r  # 新的合并区间
        sum_len += right - left + 1

        bulbs = (brightness + 2) // 3  # 至少要开启 bulbs 个灯泡
        return bulbs * sum_len
```

```java [sol-Java]
class Solution {
    public long minEnergy(int n, int brightness, int[][] intervals) {
        Arrays.sort(intervals, (p, q) -> p[0] - q[0]); // 按照左端点从小到大排序

        // 56. 合并区间（只计算区间长度之和）
        int sumLen = 0;
        int left = 0;
        int right = -1;
        for (int[] p : intervals) {
            if (p[0] <= right) { // 左端点在合并区间内，可以合并
                right = Math.max(right, p[1]); // 更新合并区间的右端点
            } else { // 不相交，无法合并
                sumLen += right - left + 1;
                left = p[0];
                right = p[1]; // 新的合并区间
            }
        }
        sumLen += right - left + 1;

        int bulbs = (brightness + 2) / 3; // 至少要开启 bulbs 个灯泡
        return (long) bulbs * sumLen;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minEnergy(int, int brightness, vector<vector<int>>& intervals) {
        ranges::sort(intervals, {}, [](auto& p) { return p[0]; }); // 按照左端点从小到大排序

        // 56. 合并区间（只计算区间长度之和）
        int sum_len = 0;
        int left = 0, right = -1;
        for (auto& p : intervals) {
            if (p[0] <= right) { // 左端点在合并区间内，可以合并
                right = max(right, p[1]); // 更新合并区间的右端点
            } else { // 不相交，无法合并
                sum_len += right - left + 1;
                left = p[0];
                right = p[1]; // 新的合并区间
            }
        }
        sum_len += right - left + 1;

        int bulbs = (brightness + 2) / 3; // 至少要开启 bulbs 个灯泡
        return 1LL * bulbs * sum_len;
    }
};
```

```go [sol-Go]
func minEnergy(_, brightness int, intervals [][]int) int64 {
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] }) // 按照左端点从小到大排序

	// 56. 合并区间（只计算区间长度之和）
	sumLen := 0
	left, right := 0, -1
	for _, p := range intervals {
		if p[0] <= right { // 左端点在合并区间内，可以合并
			right = max(right, p[1]) // 更新合并区间的右端点
		} else { // 不相交，无法合并
			sumLen += right - left + 1
			left, right = p[0], p[1] // 新的合并区间
		}
	}
	sumLen += right - left + 1

	bulbs := (brightness + 2) / 3 // 至少要开启 bulbs 个灯泡
	return int64(bulbs * sumLen)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 专题训练

见下面贪心题单的「**§2.5 合并区间**」。

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
