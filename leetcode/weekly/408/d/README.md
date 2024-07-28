## ⚠观前必读

本题没有限制圆心必须在矩形内部，但测试表明，所有**测试数据的圆心均在矩形内部**，估计过段时间会修改题目描述，保证圆心均在矩形内部。所以**下面代码只考虑圆心一定在矩形内部的情况**。

## 思路

等价转换：

- 如果从矩形【左边界/上边界】到矩形【下边界/右边界】的路被圆堵死，则无法从矩形左下角移动到矩形右上角。

怎么判断呢？

如果圆和圆相交或相切，则相当于在两个圆之间架起了一座桥。如果圆和矩形边界相交或相切，则相当于在矩形边界和圆之间架起了一座桥。如果可以从矩形【左边界/上边界】通过桥到达矩形【下边界/右边界】，则说明路被堵死，无法从矩形左下角移动到矩形右上角。

也可以把桥理解成切割线，如果能把从矩形左下角到矩形右上角的路径**切断**，则无法从矩形左下角移动到矩形右上角。

## 具体做法

抽象成一个图论题：

- 把第 $i$ 个圆视作节点 $i$，其中 $0\le i \le n-1$。
- 把矩形【左边界/上边界】视作节点 $n$。
- 把矩形【下边界/右边界】视作节点 $n+1$。

遍历每个圆 $i$：

- 判断是否与矩形【左边界/上边界】相交或相切，如果是，则合并节点 $i$ 和 $n$。
- 判断是否与矩形【下边界/右边界】相交或相切，如果是，则合并节点 $i$ 和 $n+1$。
- 判断是否与其他圆 $j$ 相交或相切，如果是，则合并节点 $i$ 和 $j$。

最后，如果节点 $n$ 和 $n+1$ 不在并查集的同一个连通块中，则返回 $\texttt{true}$，否则返回 $\texttt{false}$。也可以在遍历每个圆的过程中判断。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Mi421a7cZ/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def canReachCorner(self, x: int, y: int, a: List[List[int]]) -> bool:
        n = len(a)
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

        for i, (ox, oy, r) in enumerate(a):
            if ox <= r or oy + r >= y:  # 圆 i 和左边界或上边界有交集
                merge(i, n)
            if oy <= r or ox + r >= x:  # 圆 i 和下边界或右边界有交集
                merge(i, n + 1)
            for j, (qx, qy, qr) in enumerate(a[:i]):
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
    bool canReachCorner(int x, int y, vector<vector<int>>& a) {
        int n = a.size();
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

        for (int i = 0; i < a.size(); i++) {
            int ox = a[i][0], oy = a[i][1], r = a[i][2];
            if (ox <= r || oy + r >= y) { // 圆 i 和左边界或上边界有交集
                merge(i, n);
            }
            if (oy <= r || ox + r >= x) { // 圆 i 和下边界或右边界有交集
                merge(i, n + 1);
            }
            for (int j = 0; j < i; j++) {
                int qx = a[j][0], qy = a[j][1], qr = a[j][2];
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
func canReachCorner(x, y int, a [][]int) bool {
	n := len(a)
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
	for i, c := range a {
		ox, oy, r := c[0], c[1], c[2]
		if ox <= r || oy+r >= y { // 圆 i 和左边界或上边界有交集
			merge(i, n)
		}
		if oy <= r || ox+r >= x { // 圆 i 和下边界或右边界有交集
			merge(i, n+1)
		}
		for j, q := range a[:i] {
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

注：如果建图 + DFS，可以做到 $\mathcal{O}(n^2)$ 的时间复杂度，但那样常数太大，比并查集慢。

更多相似题目，见下面数据结构题单中的「**并查集**」。

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
