## 转化

假设划分因子 $\ge \textit{low}$。

$\textit{low}$ 越小，要求就越**宽松**，越能找到一个合法划分。

$\textit{low}$ 越大，要求就越**苛刻**，越不能找到一个合法划分。

据此，可以**二分猜答案**。关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

现在问题变成一个判定性问题：

- 给定 $\textit{low}$，能否把点集分成两个非空集合，使得每个集合内部的任意点对的曼哈顿距离都 $\ge \textit{low}$？

## 思路

对于曼哈顿距离 $< \textit{low}$ 的点对（非法点对），连一条边，我们可以得到一个无向图。

我们需要把这些点分成两个集合，每个集合内部不能有边（否则违背要求）。换句话说，每条边的两个端点，一定来自不同集合。

这和 [785. 判断二分图](https://leetcode.cn/problems/is-graph-bipartite/) 是完全一样的，做法见[【图解】交替染色法](https://leetcode.cn/problems/is-graph-bipartite/solutions/3803670/tu-jie-jiao-ti-ran-se-fa-pythonjavaccgoj-ov27/)。

## 细节

下面代码采用开区间二分，这仅仅是二分的一种写法，使用闭区间或者半闭半开区间都是可以的，喜欢哪种写法就用哪种。

- 开区间左端点初始值：$0$。可以随意划分，一定满足要求。
- 开区间右端点初始值：最大曼哈顿距离加一。此时一定无法满足要求。

> 对于开区间写法，简单来说 `check(mid) == true` 时更新的是谁，最后就返回谁。相比其他二分写法，开区间写法不需要思考加一减一等细节，更简单。推荐使用开区间写二分。

注意特判 $n=2$ 的情况，此时划分因子恒为 $0$。

## 答疑

**问**：为什么二分结束后，答案 $\textit{ans}$ 一定来自 $\textit{points}$ 的某两个点的曼哈顿距离？

**答**：反证法。假设 $\textit{ans}$ 不来自 $\textit{points}$ 的某两个点的曼哈顿距离，这意味着最小曼哈顿距离 $> \textit{ans}$，也就是 $\ge \textit{ans}+1$。换句话说，$\text{check}(\textit{ans}+1)=\texttt{true}$。但根据循环不变量，二分结束后 $\text{check}(\textit{ans}+1)=\texttt{false}$，矛盾。故原命题成立。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def maxPartitionFactor(self, points: List[List[int]]) -> int:
        if len(points) == 2:
            return 0

        # 原理见 785. 判断二分图
        def check(low: int) -> bool:
            colors = [0] * len(points)

            def dfs(x: int, c: int) -> bool:
                colors[x] = c
                x1, y1 = points[x]
                for y, (x2, y2) in enumerate(points):
                    if y == x or abs(x1 - x2) + abs(y1 - y2) >= low:  # 符合要求
                        continue
                    if colors[y] == c or colors[y] == 0 and not dfs(y, -c):
                        return False  # 不是二分图
                return True

            # 可能有多个连通块
            for i, c in enumerate(colors):
                if c == 0 and not dfs(i, 1):
                    return False
            return True

        max_dis = max(abs(x1 - x2) + abs(y1 - y2)
                      for (x1, y1), (x2, y2) in combinations(points, 2))

        left, right = 0, max_dis + 1
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid):
                left = mid
            else:
                right = mid
        return left
```

```py [sol-Python3 库函数]
class Solution:
    def maxPartitionFactor(self, points: List[List[int]]) -> int:
        if len(points) == 2:
            return 0

        # 原理见 785. 判断二分图
        def check(low: int) -> bool:
            low += 1
            colors = [0] * len(points)

            def dfs(x: int, c: int) -> bool:
                colors[x] = c
                x1, y1 = points[x]
                for y, (x2, y2) in enumerate(points):
                    if y == x or abs(x1 - x2) + abs(y1 - y2) >= low:  # 符合要求
                        continue
                    if colors[y] == c or colors[y] == 0 and not dfs(y, -c):
                        return False  # 不是二分图
                return True

            # 可能有多个连通块
            for i, c in enumerate(colors):
                if c == 0 and not dfs(i, 1):
                    return True
            return False

        max_dis = max(abs(x1 - x2) + abs(y1 - y2)
                      for (x1, y1), (x2, y2) in combinations(points, 2))
        return bisect_left(range(max_dis), True, key=check)
```

```java [sol-Java]
class Solution {
    public int maxPartitionFactor(int[][] points) {
        int n = points.length;
        if (n == 2) {
            return 0;
        }

        int maxDis = 0;
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                int d = Math.abs(points[i][0] - points[j][0]) + Math.abs(points[i][1] - points[j][1]);
                maxDis = Math.max(maxDis, d);
            }
        }

        int left = 0, right = maxDis + 1;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            if (check(mid, points)) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return left;
    }

    // 原理见 785. 判断二分图
    private boolean check(int low, int[][] points) {
        int n = points.length;
        int[] colors = new int[n];
        for (int i = 0; i < n; i++) {
            if (colors[i] == 0 && !dfs(i, 1, low, points, colors)) {
                return false;
            }
        }
        return true;
    }

    private boolean dfs(int x, int c, int low, int[][] points, int[] colors) {
        colors[x] = c;
        int x1 = points[x][0], y1 = points[x][1];
        for (int y = 0; y < points.length; y++) {
            if (y == x || Math.abs(x1 - points[y][0]) + Math.abs(y1 - points[y][1]) >= low) { // 符合要求
                continue;
            }
            if (colors[y] == c || colors[y] == 0 && !dfs(y, -c, low, points, colors)) {
                return false; // 不是二分图
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxPartitionFactor(vector<vector<int>>& points) {
        int n = points.size();
        if (n == 2) {
            return 0;
        }

        // 原理见 785. 判断二分图
        auto check = [&](int low) -> bool {
            vector<int8_t> colors(n);

            auto dfs = [&](this auto&& dfs, int x, int8_t c) -> bool {
                colors[x] = c;
                auto& p = points[x];
                for (int y = 0; y < n; y++) {
                    auto& q = points[y];
                    if (y == x || abs(p[0] - q[0]) + abs(p[1] - q[1]) >= low) { // 符合要求
                        continue;
                    }
                    if (colors[y] == c || colors[y] == 0 && !dfs(y, -c)) {
                        return false; // 不是二分图
                    }
                }
                return true;
            };

            // 可能有多个连通块
            for (int i = 0; i < n; i++) {
                if (colors[i] == 0 && !dfs(i, 1)) {
                    return false;
                }
            }
            return true;
        };

        int max_dis = 0;
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                max_dis = max(max_dis, abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1]));
            }
        }

        int left = 0, right = max_dis + 1;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            (check(mid) ? left : right) = mid;
        }
        return left;
    }
};
```

```go [sol-Go]
// 原理见 785. 判断二分图
func isBipartite(points [][]int, low int) bool {
	colors := make([]int8, len(points))

	var dfs func(int, int8) bool
	dfs = func(x int, c int8) bool {
		colors[x] = c
		p := points[x]
		for y, q := range points {
			if y == x || abs(p[0]-q[0])+abs(p[1]-q[1]) >= low { // 符合要求
				continue
			}
			if colors[y] == c || colors[y] == 0 && !dfs(y, -c) {
				return false // 不是二分图
			}
		}
		return true
	}

	// 可能有多个连通块
	for i, c := range colors {
		if c == 0 && !dfs(i, 1) {
			return false
		}
	}
	return true
}

func maxPartitionFactor(points [][]int) int {
	n := len(points)
	if n == 2 {
		return 0
	}

	// 不想算的话可以写 maxDis := int(4e8)
	maxDis := 0
	for i, p := range points {
		for _, q := range points[:i] {
			maxDis = max(maxDis, abs(p[0]-q[0])+abs(p[1]-q[1]))
		}
	}

	return sort.Search(maxDis, func(low int) bool {
		// 二分最小的不满足要求的 low+1，就可以得到最大的满足要求的 low
		return !isBipartite(points, low+1)
	})
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2\log U)$，其中 $n$ 是 $\textit{points}$ 的长度，$U\le 4\times 10^8$ 是曼哈顿距离的最大值。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面图论题单的「**七、二分图染色**」。

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
