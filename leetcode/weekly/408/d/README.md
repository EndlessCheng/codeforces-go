## 方法一：并查集

**等价转换**：如果从矩形【左边界/上边界】到矩形【下边界/右边界】的路被圆堵死，则无法从矩形左下角移动到矩形右上角。

怎么判断呢？

如果圆和圆相交或相切，则相当于在两个圆之间架起了一座桥。如果圆和矩形边界相交或相切，则相当于在矩形边界和圆之间架起了一座桥。如果可以从矩形【左边界/上边界】通过桥到达矩形【下边界/右边界】，则说明路被堵死，无法从矩形左下角移动到矩形右上角。

也可以把桥理解成切割线，如果能把从矩形左下角到矩形右上角的路径**切断**，则无法从矩形左下角移动到矩形右上角。

注意题目保证圆心均在矩形内部。

### 具体做法

抽象成一个图论题：

- 把第 $i$ 个圆视作节点 $i$，其中 $0\le i \le n-1$。
- 把矩形【左边界/上边界】视作节点 $n$。
- 把矩形【下边界/右边界】视作节点 $n+1$。

遍历每个圆 $i$：

- 判断是否与矩形【左边界/上边界】相交或相切，如果是，则合并节点 $i$ 和 $n$。
- 判断是否与矩形【下边界/右边界】相交或相切，如果是，则合并节点 $i$ 和 $n+1$。
- 判断是否与其他圆 $j$ 相交或相切，如果是，则合并节点 $i$ 和 $j$。

最后，如果节点 $n$ 和 $n+1$ 不在并查集的同一个连通块中，则返回 $\texttt{true}$，否则返回 $\texttt{false}$。也可以在遍历每个圆的过程中判断。

两圆是否相交相切，可以判断圆心距离与两圆半径之和的大小关系。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Mi421a7cZ/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def canReachCorner(self, x: int, y: int, circles: List[List[int]]) -> bool:
        n = len(circles)
        # 并查集中的 n 表示左边界或上边界，n+1 表示下边界或右边界
        fa = list(range(n + 2))
        # 非递归并查集
        def find(x: int) -> int:
            rt = x
            while fa[rt] != rt:
                rt = fa[rt]
            while fa[x] != rt:
                fa[x], x = rt, fa[x]
            return rt
        def merge(x: int, y: int) -> None:
            fa[find(x)] = find(y)

        for i, (ox, oy, r) in enumerate(circles):
            if ox <= r or oy + r >= y:  # 圆 i 和左边界或上边界有交集
                merge(i, n)
            if oy <= r or ox + r >= x:  # 圆 i 和下边界或右边界有交集
                merge(i, n + 1)
            for j, (qx, qy, qr) in enumerate(circles[:i]):
                if (ox - qx) * (ox - qx) + (oy - qy) * (oy - qy) <= (r + qr) * (r + qr):
                    merge(i, j)  # 圆 i 和圆 j 有交集
            if find(n) == find(n + 1):  # 无法到达终点
                return False
        return True
```

```java [sol-Java]
class UnionFind {
    private final int[] fa;

    public UnionFind(int size) {
        fa = new int[size];
        for (int i = 1; i < size; i++) {
            fa[i] = i;
        }
    }

    public int find(int x) {
        if (fa[x] != x) {
            fa[x] = find(fa[x]);
        }
        return fa[x];
    }

    public void merge(int x, int y) {
        fa[find(x)] = find(y);
    }
}

class Solution {
    public boolean canReachCorner(int x, int y, int[][] circles) {
        int n = circles.length;
        // 并查集中的 n 表示左边界或上边界，n+1 表示下边界或右边界
        UnionFind uf = new UnionFind(n + 2);
        for (int i = 0; i < n; i++) {
            int[] c = circles[i];
            int ox = c[0], oy = c[1], r = c[2];
            if (ox <= r || oy + r >= y) { // 圆 i 和左边界或上边界有交集
                uf.merge(i, n);
            }
            if (oy <= r || ox + r >= x) { // 圆 i 和下边界或右边界有交集
                uf.merge(i, n + 1);
            }
            for (int j = 0; j < i; j++) {
                int[] q = circles[j];
                if ((long) (ox - q[0]) * (ox - q[0]) + (long) (oy - q[1]) * (oy - q[1]) <= (long) (r + q[2]) * (r + q[2])) {
                    uf.merge(i, j); // 圆 i 和圆 j 有交集
                }
            }
            if (uf.find(n) == uf.find(n + 1)) { // 无法到达终点
                return false;
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool canReachCorner(int x, int y, vector<vector<int>>& circles) {
        int n = circles.size();
        // 并查集中的 n 表示左边界或上边界，n+1 表示下边界或右边界
        vector<int> fa(n + 2);
        iota(fa.begin(), fa.end(), 0);
        // 非递归并查集
        auto find = [&](int x) {
            int rt = x;
            while (fa[rt] != rt) {
                rt = fa[rt];
            }
            while (fa[x] != rt) {
                int temp = fa[x];
                fa[x] = rt;
                x = temp;
            }
            return rt;
        };
        auto merge = [&](int x, int y) {
            fa[find(x)] = find(y);
        };

        for (int i = 0; i < circles.size(); i++) {
            int ox = circles[i][0], oy = circles[i][1], r = circles[i][2];
            if (ox <= r || oy + r >= y) { // 圆 i 和左边界或上边界有交集
                merge(i, n);
            }
            if (oy <= r || ox + r >= x) { // 圆 i 和下边界或右边界有交集
                merge(i, n + 1);
            }
            for (int j = 0; j < i; j++) {
                int qx = circles[j][0], qy = circles[j][1], qr = circles[j][2];
                if ((long long) (ox - qx) * (ox - qx) + (long long) (oy - qy) * (oy - qy) <= (long long) (r + qr) * (r + qr)) {
                    merge(i, j); // 圆 i 和圆 j 有交集
                }
            }
            if (find(n) == find(n + 1)) { // 无法到达终点
                return false;
            }
        }
        return true;
    }
};
```

```go [sol-Go]
func canReachCorner(x, y int, circles [][]int) bool {
	n := len(circles)
	// 并查集中的 n 表示左边界或上边界，n+1 表示下边界或右边界
	fa := make([]int, n+2)
	for i := range fa {
		fa[i] = i
	}
	// 非递归并查集
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}
	merge := func(x, y int) {
		fa[find(x)] = find(y)
	}

	for i, c := range circles {
		ox, oy, r := c[0], c[1], c[2]
		if ox <= r || oy+r >= y { // 圆 i 和左边界或上边界有交集
			merge(i, n)
		}
		if oy <= r || ox+r >= x { // 圆 i 和下边界或右边界有交集
			merge(i, n+1)
		}
		for j, q := range circles[:i] {
			if (ox-q[0])*(ox-q[0])+(oy-q[1])*(oy-q[1]) <= (r+q[2])*(r+q[2]) {
				merge(i, j) // 圆 i 和圆 j 有交集
			}
		}
		if find(n) == find(n+1) { // 无法到达终点
			return false
		}
	}
	return true
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2\log n)$，其中 $n$ 是 $\textit{circles}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：DFS

从与左边界或上边界相交相切的圆出发，DFS 这张图，如果可以到达下边界或右边界，则说明路被堵死。

代码实现时，无需建图，直接判断圆心距离与两圆半径之和的大小关系。

```py [sol-Python3]
class Solution:
    def canReachCorner(self, x: int, y: int, circles: List[List[int]]) -> bool:
        vis = [False] * len(circles)
        def dfs(i: int) -> bool:
            ox, oy, r = circles[i]
            if oy <= r or ox + r >= x:
                return True
            vis[i] = True
            for j, (qx, qy, qr) in enumerate(circles):
                if not vis[j] and (ox - qx) * (ox - qx) + (oy - qy) * (oy - qy) <= (r + qr) * (r + qr) and dfs(j):
                    return True
            return False
        for i, (ox, oy, r) in enumerate(circles):
            if (ox <= r or oy + r >= y) and not vis[i] and dfs(i):
                return False
        return True
```

```java [sol-Java]
class Solution {
    public boolean canReachCorner(int x, int y, int[][] circles) {
        boolean[] vis = new boolean[circles.length];
        for (int i = 0; i < circles.length; i++) {
            int ox = circles[i][0], oy = circles[i][1], r = circles[i][2];
            if ((ox <= r || oy + r >= y) && !vis[i] && dfs(i, x, circles, vis)) {
                return false;
            }
        }
        return true;
    }

    private boolean dfs(int i, int x, int[][] circles, boolean[] vis) {
        int ox = circles[i][0], oy = circles[i][1], r = circles[i][2];
        if (oy <= r || ox + r >= x) {
            return true;
        }
        vis[i] = true;
        for (int j = 0; j < circles.length; j++) {
            if (!vis[j]) {
                int qx = circles[j][0], qy = circles[j][1], qr = circles[j][2];
                if ((long) (ox - qx) * (ox - qx) + (long) (oy - qy) * (oy - qy) <= (long) (r + qr) * (r + qr) && dfs(j, x, circles, vis)) {
                    return true;
                }
            }
        }
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool canReachCorner(int x, int y, vector<vector<int>>& circles) {
        int n = circles.size();
        vector<int> vis(n);
        auto dfs = [&](auto&& dfs, int i) -> bool {
            int ox = circles[i][0], oy = circles[i][1], r = circles[i][2];
            if (oy <= r || ox + r >= x) {
                return true;
            }
            vis[i] = true;
            for (int j = 0; j < n; j++) {
                if (!vis[j]) {
                    int qx = circles[j][0], qy = circles[j][1], qr = circles[j][2];
                    if ((long long) (ox - qx) * (ox - qx) + (long long) (oy - qy) * (oy - qy) <= (long long) (r + qr) * (r + qr) && dfs(dfs, j)) {
                        return true;
                    }
                }
            }
            return false;
        };
        for (int i = 0; i < n; i++) {
            int ox = circles[i][0], oy = circles[i][1], r = circles[i][2];
            if ((ox <= r || oy + r >= y) && !vis[i] && dfs(dfs, i)) {
                return false;
            }
        }
        return true;
    }
};
```

```go [sol-Go]
func canReachCorner(x, y int, circles [][]int) bool {
	vis := make([]bool, len(circles))
	var dfs func(int) bool
	dfs = func(i int) bool {
		ox, oy, r := circles[i][0], circles[i][1], circles[i][2]
		if oy <= r || ox+r >= x {
			return true
		}
		vis[i] = true
		for j, b := range vis {
			if !b {
				qx, qy, qr := circles[j][0], circles[j][1], circles[j][2]
				if (ox-qx)*(ox-qx)+(oy-qy)*(oy-qy) <= (r+qr)*(r+qr) && dfs(j) {
					return true
				}
			}
		}
		return false
	}
	for i, c := range circles {
		ox, oy, r := c[0], c[1], c[2]
		if (ox <= r || oy+r >= y) && !vis[i] && dfs(i) {
			return false
		}
	}
	return true
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{circles}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面图论题单中的「**DFS**」和数据结构题单中的「**并查集**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
