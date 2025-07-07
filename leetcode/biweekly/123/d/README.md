将 $\textit{points}$ 按照横坐标**从小到大**排序，横坐标相同的，按照纵坐标**从大到小**排序。

如此一来，在枚举 $\textit{points}[i]$ 和 $\textit{points}[j]$ 时（$i<j$），就只需要关心纵坐标的大小。

固定 $\textit{points}[i]$，然后枚举 $\textit{points}[j]$：

- 如果 $\textit{points}[j]$ 的纵坐标比之前枚举的点的纵坐标都大，那么矩形内没有其它点，符合要求，答案加一。
- 如果 $\textit{points}[j]$ 的纵坐标小于等于之前枚举的某个点的纵坐标，那么矩形内有其它点，不符合要求。

所以在枚举 $\textit{points}[j]$ 的同时，需要维护纵坐标的最大值 $\textit{maxY}$。这也解释了为什么横坐标相同的，按照纵坐标**从大到小**排序。这保证了横坐标相同时，我们总是优先枚举更靠上的点，不会误把包含其它点的矩形也当作符合要求的矩形。

[视频讲解](https://www.bilibili.com/video/BV14C411r7nN/) 第四题。

```py [sol-Python3]
class Solution:
    def numberOfPairs(self, points: List[List[int]]) -> int:
        points.sort(key=lambda p: (p[0], -p[1]))
        ans = 0
        for i, (_, y0) in enumerate(points):
            max_y = -inf
            for (_, y) in points[i + 1:]:
                if max_y < y <= y0:
                    max_y = y
                    ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int numberOfPairs(int[][] points) {
        Arrays.sort(points, (p, q) -> p[0] != q[0] ? p[0] - q[0] : q[1] - p[1]);
        int ans = 0;
        for (int i = 0; i < points.length; i++) {
            int y0 = points[i][1];
            int maxY = Integer.MIN_VALUE;
            for (int j = i + 1; j < points.length; j++) {
                int y = points[j][1];
                if (y <= y0 && y > maxY) {
                    maxY = y;
                    ans++;
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfPairs(vector<vector<int>>& points) {
        ranges::sort(points, {}, [](auto& p) -> pair<int, int> { return {p[0], -p[1]}; });
        int ans = 0, n = points.size();
        for (int i = 0; i < n; i++) {
            int y0 = points[i][1];
            int max_y = INT_MIN;
            for (int j = i + 1; j < n; j++) {
                int y = points[j][1];
                if (y <= y0 && y > max_y) {
                    max_y = y;
                    ans++;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfPairs(points [][]int) (ans int) {
	slices.SortFunc(points, func(a, b []int) int { return cmp.Or(a[0]-b[0], b[1]-a[1]) })
	for i, p := range points {
		y0 := p[1]
		maxY := math.MinInt
		for _, q := range points[i+1:] {
			y := q[1]
			if y <= y0 && y > maxY {
				maxY = y
				ans++
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{points}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销和 Python 切片开销。

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
