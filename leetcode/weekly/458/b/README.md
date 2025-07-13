上场周赛出过一个类似的题目 [3608. 包含 K 个连通分量需要的最小时间](https://leetcode.cn/problems/minimum-time-for-k-connected-components/)，代码几乎一样。

如果 $k=n$，无需合并，返回 $0$。

否则，为了最小化答案，我们可以用**并查集**从小到大合并，直到连通块个数等于 $k$ 为止。此时，合并边之前有 $k+1$ 个连通块，合并后有 $k$ 个连通块，且参与合并的边权都是尽量小的，返回当前边的边权。合并更大的边权是没有意义的，因为小于等于最大边权的边我们都能合并，合并更大边权不会得到更优解。

```py [sol-Python3]
# 完整的并查集模板，见我的数据结构题单
class UnionFind:
    def __init__(self, n: int):
        # 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        # 集合 i 的代表元是自己，大小为 1
        self._fa = list(range(n))  # 代表元
        self.cc = n  # 连通块个数

    # 返回 x 所在集合的代表元
    # 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
    def find(self, x: int) -> int:
        # 如果 fa[x] == x，则表示 x 是代表元
        if self._fa[x] != x:
            self._fa[x] = self.find(self._fa[x])  # fa 改成代表元
        return self._fa[x]

    # 把 from 所在集合合并到 to 所在集合中
    def merge(self, from_: int, to: int) -> None:
        x, y = self.find(from_), self.find(to)
        if x == y:  # from 和 to 在同一个集合，不做合并
            return
        self._fa[x] = y  # 合并集合。修改后就可以认为 from 和 to 在同一个集合了
        self.cc -= 1  # 成功合并，连通块个数减一


class Solution:
    def minCost(self, n: int, edges: List[List[int]], k: int) -> int:
        if k == n:
            return 0
    
        edges.sort(key=lambda e: e[2])
        u = UnionFind(n)
        for x, y, w in edges:
            u.merge(x, y)
            if u.cc <= k:
                return w
```

```java [sol-Java]
// 完整的并查集模板，见我的数据结构题单
class UnionFind {
    private final int[] fa; // 代表元
    public int cc; // 连通块个数

    UnionFind(int n) {
        // 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        // 集合 i 的代表元是自己，大小为 1
        fa = new int[n];
        for (int i = 0; i < n; i++) {
            fa[i] = i;
        }
        cc = n;
    }

    // 返回 x 所在集合的代表元
    // 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
    public int find(int x) {
        // 如果 fa[x] == x，则表示 x 是代表元
        if (fa[x] != x) {
            fa[x] = find(fa[x]); // fa 改成代表元
        }
        return fa[x];
    }

    // 把 from 所在集合合并到 to 所在集合中
    public void merge(int from, int to) {
        int x = find(from);
        int y = find(to);
        if (x == y) { // from 和 to 在同一个集合，不做合并
            return;
        }
        fa[x] = y; // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
        cc--; // 成功合并，连通块个数减一
    }
}

class Solution {
    int minCost(int n, int[][] edges, int k) {
        if (k == n) {
            return 0;
        }

        Arrays.sort(edges, (a, b) -> a[2] - b[2]);
        UnionFind u = new UnionFind(n);
        for (int[] e : edges) {
            u.merge(e[0], e[1]);
            if (u.cc <= k) {
                return e[2];
            }
        }
        throw new RuntimeException("impossible");
    }
}
```

```cpp [sol-C++]
// 完整的并查集模板，见我的数据结构题单
class UnionFind {
    vector<int> fa; // 代表元

public:
    int cc; // 连通块个数

    UnionFind(int n) : fa(n), cc(n) {
        // 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        // 集合 i 的代表元是自己，大小为 1
        ranges::iota(fa, 0); // iota(fa.begin(), fa.end(), 0);
    }

    // 返回 x 所在集合的代表元
    // 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
    int find(int x) {
        // 如果 fa[x] == x，则表示 x 是代表元
        if (fa[x] != x) {
            fa[x] = find(fa[x]); // fa 改成代表元
        }
        return fa[x];
    }

    // 把 from 所在集合合并到 to 所在集合中
    void merge(int from, int to) {
        int x = find(from), y = find(to);
        if (x == y) { // from 和 to 在同一个集合，不做合并
            return;
        }
        fa[x] = y; // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
        cc--; // 成功合并，连通块个数减一
    }
};

class Solution {
public:
    int minCost(int n, vector<vector<int>>& edges, int k) {
        if (k == n) {
            return 0;
        }

        ranges::sort(edges, {}, [](auto& e) { return e[2]; });
        UnionFind u(n);
        for (auto& e : edges) {
            u.merge(e[0], e[1]);
            if (u.cc <= k) {
                return e[2];
            }
        }
        throw runtime_error("impossible");
    }
};
```

```go [sol-Go]
// 完整的并查集模板，见我的数据结构题单
type unionFind struct {
	fa []int // 代表元
	cc int   // 连通块个数
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	// 一开始有 n 个集合 {0}, {1}, ..., {n-1}
	// 集合 i 的代表元是自己
	for i := range fa {
		fa[i] = i
	}
	return unionFind{fa, n}
}

// 返回 x 所在集合的代表元
// 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
func (u unionFind) find(x int) int {
	// 如果 fa[x] == x，则表示 x 是代表元
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x]) // fa 改成代表元
	}
	return u.fa[x]
}

// 把 from 所在集合合并到 to 所在集合中
func (u *unionFind) merge(from, to int) {
	x, y := u.find(from), u.find(to)
	if x == y { // from 和 to 在同一个集合，不做合并
		return
	}
	u.fa[x] = y // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
	u.cc--      // 成功合并，连通块个数减一
}

func minCost(n int, edges [][]int, k int) int {
	if k == n {
		return 0
	}

	slices.SortFunc(edges, func(a, b []int) int { return a[2] - b[2] })
	u := newUnionFind(n)
	for _, e := range edges {
		u.merge(e[0], e[1])
		if u.cc <= k {
			return e[2]
		}
	}
	panic("impossible")
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log m + m\log n)$，其中 $m$ 是 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。忽略排序的栈开销。

## 专题训练

见下面数据结构题单的「**七、并查集**」。

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
