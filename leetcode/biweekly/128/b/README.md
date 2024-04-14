由于矩形的高没有限制，所以我们只需考虑 $\textit{points}$ 的横坐标。

矩形越大越好，所以 $x_2$ 应该恰好等于 $x_1+w$。

算法如下：

1. 把横坐标按照从小到大的顺序排序。
2. 为方便计算，假设第一个矩形左边还有一个矩形，初始化 $x_2 = -1$，因为所有横坐标都是非负数。
3. 遍历横坐标 $x=\textit{points}[i][0]$，如果 $x > x_2$，我们需要一个新的 $x_1=x$ 的矩形，答案加一，然后把 $x_2$ 更新为 $x+w$。
4. 最后，返回答案。

请看 [视频讲解](https://www.bilibili.com/video/BV1et42177VM/) 第二题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minRectanglesToCoverPoints(self, points: List[List[int]], w: int) -> int:
        points.sort(key=lambda p: p[0])
        ans = 0
        x2 = -1
        for x, _ in points:
            if x > x2:
                ans += 1
                x2 = x + w
        return ans
```

```java [sol-Java]
class Solution {
    public int minRectanglesToCoverPoints(int[][] points, int w) {
        Arrays.sort(points, (p, q) -> p[0] - q[0]);
        int ans = 0;
        int x2 = -1;
        for (int[] p : points) {
            if (p[0] > x2) {
                ans++;
                x2 = p[0] + w;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minRectanglesToCoverPoints(vector<vector<int>>& points, int w) {
        ranges::sort(points, [](const auto& p, const auto& q) {
            return p[0] < q[0];
        });
        int ans = 0;
        int x2 = -1;
        for (auto& p : points) {
            if (p[0] > x2) {
                ans++;
                x2 = p[0] + w;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minRectanglesToCoverPoints(points [][]int, w int) (ans int) {
	slices.SortFunc(points, func(p, q []int) int { return p[0] - q[0] })
	x2 := -1
	for _, p := range points {
		if p[0] > x2 {
			ans++
			x2 = p[0] + w
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{points}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
