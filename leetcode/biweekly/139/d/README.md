设 $\textit{kx} = \textit{coordinates}[k][0],\ \textit{ky} = \textit{coordinates}[k][1]$。

仔细读题：我们要选的并不是 $\textit{coordinates}$ 的子序列，所以可以排序。

按照 $x$ 从小到大排序，问题变成计算 $y$ 的 LIS。

⚠**注意**：对于 $x$ 相同的点，要按照 $y$ **从大到小**排序。这可以保证在计算 LIS 时，对于相同的 $x$，我们至多选一个 $y$。

然后选择 $x<\textit{kx}$ 且 $y < \textit{ky}$ 或者 $x>\textit{kx}$ 且 $y > \textit{ky}$ 的 $y$ 计算 LIS。

关于 LIS 的二分算法，请看[【基础算法精讲 20】](https://www.bilibili.com/video/BV1ub411Q7sB/)。

[本题视频讲解](https://www.bilibili.com/video/BV1Ub4mekE1x/)（第四题），欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxPathLength(self, coordinates: List[List[int]], k: int) -> int:
        kx, ky = coordinates[k]
        coordinates.sort(key=lambda p: (p[0], -p[1]))

        g = []
        for x, y in coordinates:
            if x < kx and y < ky or x > kx and y > ky:
                j = bisect_left(g, y)
                if j < len(g):
                    g[j] = y
                else:
                    g.append(y)
        return len(g) + 1  # 算上 coordinates[k]
```

```java [sol-Java]
class Solution {
    public int maxPathLength(int[][] coordinates, int k) {
        int kx = coordinates[k][0];
        int ky = coordinates[k][1];
        Arrays.sort(coordinates, (a, b) -> a[0] != b[0] ? a[0] - b[0] : b[1] - a[1]);

        List<Integer> g = new ArrayList<>();
        for (int[] p : coordinates) {
            int x = p[0];
            int y = p[1];
            if (x < kx && y < ky || x > kx && y > ky) {
                int j = Collections.binarySearch(g, y); // g 没有重复元素，可以用 binarySearch
                if (j < 0) {
                    j = -j - 1;
                }
                if (j < g.size()) {
                    g.set(j, y);
                } else {
                    g.add(y);
                }
            }
        }
        return g.size() + 1; // 算上 coordinates[k]
    }
}
```

```java [sol-Java 数组]
class Solution {
    public int maxPathLength(int[][] coordinates, int k) {
        int kx = coordinates[k][0];
        int ky = coordinates[k][1];
        Arrays.sort(coordinates, (a, b) -> a[0] != b[0] ? a[0] - b[0] : b[1] - a[1]);

        int[] g = new int[coordinates.length];
        int m = 0; // g 的长度
        for (int[] p : coordinates) {
            int x = p[0];
            int y = p[1];
            if (x < kx && y < ky || x > kx && y > ky) {
                int j = Arrays.binarySearch(g, 0, m, y); // g 没有重复元素，可以用 binarySearch
                if (j < 0) {
                    j = -j - 1;
                }
                if (j < m) {
                    g[j] = y;
                } else {
                    g[m++] = y;
                }
            }
        }
        return m + 1; // 算上 coordinates[k]
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxPathLength(vector<vector<int>>& coordinates, int k) {
        int kx = coordinates[k][0], ky = coordinates[k][1];
        ranges::sort(coordinates, [](const auto& a, const auto& b) {
            return a[0] < b[0] || a[0] == b[0] && a[1] > b[1];
        });

        vector<int> g;
        for (auto& p : coordinates) {
            int x = p[0], y = p[1];
            if (x < kx && y < ky || x > kx && y > ky) {
                auto it = ranges::lower_bound(g, y);
                if (it != g.end()) {
                    *it = y;
                } else {
                    g.push_back(y);
                }
            }
        }
        return g.size() + 1; // 算上 coordinates[k]
    }
};
```

```go [sol-Go]
func maxPathLength(coordinates [][]int, k int) int {
	kx, ky := coordinates[k][0], coordinates[k][1]
	slices.SortFunc(coordinates, func(a, b []int) int {
		return cmp.Or(a[0]-b[0], b[1]-a[1])
	})

	g := []int{}
	for _, p := range coordinates {
		x, y := p[0], p[1]
		if x < kx && y < ky || x > kx && y > ky {
			j := sort.SearchInts(g, y)
			if j < len(g) {
				g[j] = y
			} else {
				g = append(g, y)
			}
		}
	}
	return len(g) + 1 // 算上 coordinates[k]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{coordinates}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

返回一个长为 $n$ 的数组，表示 $k=0,1,2,\cdots,n-1$ 时原问题的答案。

欢迎在评论区分享你的思路/代码。

更多相似题目，见下面动态规划题单中的「**§4.2 最长递增子序列（LIS）**」。

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
