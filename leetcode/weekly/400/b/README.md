正难则反，答案等于 $\textit{days}$ 减「有会议安排的天数」。

要计算有会议安排的天数，做法同 [56. 合并区间的题解](https://leetcode.cn/problems/merge-intervals/solution/jian-dan-zuo-fa-yi-ji-wei-shi-yao-yao-zh-f2b3/)。累加每个区间的长度，即为有会议安排的天数。

由于本题只需计算合并区间的长度，所以只需记录当前合并区间的左右端点。

```py [sol-Python3]
class Solution:
    def countDays(self, days: int, meetings: List[List[int]]) -> int:
        meetings.sort(key=lambda p: p[0])  # 按照左端点从小到大排序
        start, end = 1, 0  # 当前合并区间的左右端点
        for s, e in meetings:
            if s > end:  # 不相交
                days -= end - start + 1  # 当前合并区间的长度
                start = s  # 下一个合并区间的左端点
            end = max(end, e)
        days -= end - start + 1  # 最后一个合并区间的长度
        return days
```

```java [sol-Java]
class Solution {
    public int countDays(int days, int[][] meetings) {
        Arrays.sort(meetings, (p, q) -> p[0] - q[0]); // 按照左端点从小到大排序
        int start = 1, end = 0; // 当前合并区间的左右端点
        for (int[] p : meetings) {
            if (p[0] > end) { // 不相交
                days -= end - start + 1; // 当前合并区间的长度
                start = p[0]; // 下一个合并区间的左端点
            }
            end = Math.max(end, p[1]);
        }
        days -= end - start + 1; // 最后一个合并区间的长度
        return days;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countDays(int days, vector<vector<int>>& meetings) {
        ranges::sort(meetings); // 按照左端点从小到大排序
        int start = 1, end = 0; // 当前合并区间的左右端点
        for (auto& p : meetings) {
            if (p[0] > end) { // 不相交
                days -= end - start + 1; // 当前合并区间的长度
                start = p[0]; // 下一个合并区间的左端点
            }
            end = max(end, p[1]);
        }
        days -= end - start + 1; // 最后一个合并区间的长度
        return days;
    }
};
```

```go [sol-Go]
func countDays(days int, meetings [][]int) int {
	slices.SortFunc(meetings, func(p, q []int) int { return p[0] - q[0] }) // 按照左端点从小到大排序
	start, end := 1, 0 // 当前合并区间的左右端点
	for _, p := range meetings {
		if p[0] > end { // 不相交
			days -= end - start + 1 // 当前合并区间的长度
			start = p[0] // 下一个合并区间的左端点
		}
		end = max(end, p[1])
	}
	days -= end - start + 1 // 最后一个合并区间的长度
	return days
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{meetings}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
