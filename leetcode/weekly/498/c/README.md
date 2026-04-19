本题是多源 BFS 问题，关键是题目的这个要求：如果多个颜色在同一时间步到达同一个未着色单元格，该单元格将采用具有**最大值**的颜色。

我们可以先把 $\textit{sources}$ 按照 $\textit{color}$ **从大到小排序**，然后把 $\textit{sources}$ 中的所有元素入队。这样在 BFS 的过程中，我们会优先扩散颜色大的单元格。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def colorGrid(self, n: int, m: int, sources: list[list[int]]) -> list[list[int]]:
        ans = [[0] * m for _ in range(n)]
        for x, y, c in sources:
            ans[x][y] = c  # 初始颜色

        sources.sort(key=lambda s: -s[2])
        q = deque(sources)

        while q:
            i, j, c = q.popleft()
            for x, y in (i, j - 1), (i, j + 1), (i - 1, j), (i + 1, j):  # 左右上下
                if 0 <= x < n and 0 <= y < m and ans[x][y] == 0:  # (x, y) 未着色
                    ans[x][y] = c  # 着色
                    q.append((x, y, c))  # 继续扩散

        return ans
```

```java [sol-Java]
class Solution {
    private static final int[][] DIRS = {{0, -1}, {0, 1}, {-1, 0}, {1, 0}}; // 左右上下

    public int[][] colorGrid(int n, int m, int[][] sources) {
        Arrays.sort(sources, (a, b) -> b[2] - a[2]);

        int[][] ans = new int[n][m];
        // ArrayDeque 比较慢，更快的写法见【Java 数组】
        Queue<int[]> q = new ArrayDeque<>();
        for (int[] p : sources) {
            ans[p[0]][p[1]] = p[2]; // 初始颜色
            q.offer(p);
        }

        while (!q.isEmpty()) {
            int[] p = q.poll();
            int i = p[0];
            int j = p[1];
            int c = p[2];
            for (int[] dir : DIRS) { // 向四个方向扩散
                int x = i + dir[0];
                int y = j + dir[1];
                if (0 <= x && x < n && 0 <= y && y < m && ans[x][y] == 0) { // (x, y) 未着色
                    ans[x][y] = c; // 着色
                    q.offer(new int[]{x, y, c}); // 继续扩散
                }
            }
        }

        return ans;
    }
}
```

```java [sol-Java 数组]
class Solution {
    private static final int[][] DIRS = {{0, -1}, {0, 1}, {-1, 0}, {1, 0}}; // 左右上下

    public int[][] colorGrid(int n, int m, int[][] sources) {
        Arrays.sort(sources, (a, b) -> b[2] - a[2]);

        int[][] ans = new int[n][m];
        int[][] q = new int[n * m][];
        int head = 0;
        int tail = 0;
        for (int[] p : sources) {
            ans[p[0]][p[1]] = p[2]; // 初始颜色
            q[tail++] = p;
        }

        while (head < tail) {
            int[] p = q[head++];
            int i = p[0];
            int j = p[1];
            int c = p[2];
            for (int[] dir : DIRS) { // 向四个方向扩散
                int x = i + dir[0];
                int y = j + dir[1];
                if (0 <= x && x < n && 0 <= y && y < m && ans[x][y] == 0) { // (x, y) 未着色
                    ans[x][y] = c; // 着色
                    q[tail++] = new int[]{x, y, c}; // 继续扩散
                }
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int DIRS[4][2] = {{0, -1}, {0, 1}, {-1, 0}, {1, 0}}; // 左右上下

public:
    vector<vector<int>> colorGrid(int n, int m, vector<vector<int>>& sources) {
        ranges::sort(sources, {}, [](auto& a) { return -a[2]; });

        vector ans(n, vector<int>(m));
        queue<tuple<int, int, int>> q;
        for (auto& p : sources) {
            ans[p[0]][p[1]] = p[2]; // 初始颜色
            q.emplace(p[0], p[1], p[2]);
        }

        while (!q.empty()) {
            auto [i, j, c] = q.front();
            q.pop();
            for (auto [dx, dy] : DIRS) { // 向四个方向扩散
                int x = i + dx, y = j + dy;
                if (0 <= x && x < n && 0 <= y && y < m && ans[x][y] == 0) { // (x, y) 未着色
                    ans[x][y] = c; // 着色
                    q.emplace(x, y, c); // 继续扩散
                }
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
var dirs = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // 左右上下

func colorGrid(n, m int, sources [][]int) [][]int {
	slices.SortFunc(sources, func(a, b []int) int { return b[2] - a[2] })

	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}
	for _, p := range sources {
		ans[p[0]][p[1]] = p[2] // 初始颜色
	}

	q := sources
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		x, y, c := p[0], p[1], p[2]
		for _, d := range dirs { // 向四个方向扩散
			i, j := x+d.x, y+d.y
			if 0 <= i && i < n && 0 <= j && j < m && ans[i][j] == 0 { // (i, j) 未着色
				ans[i][j] = c // 着色
				q = append(q, []int{i, j, c}) // 继续扩散
			}
		}
	}

	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(k\log k + nm)$，其中 $k$ 是 $\textit{sources}$ 的长度。
- 空间复杂度：$\mathcal{O}(nm)$。

## 专题训练

见下面网格图题单的「**二、网格图 BFS**」。

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
